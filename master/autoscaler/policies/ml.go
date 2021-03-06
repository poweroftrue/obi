// Copyright 2018 Delivery Hero Germany
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
//     Unless required by applicable law or agreed to in writing, software
//     distributed under the License is distributed on an "AS IS" BASIS,
//     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//     See the License for the specific language governing permissions and
//     limitations under the License.

package policies

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"math"
	"obi/master/model"
	"obi/master/predictor"
	"obi/master/utils"
	"os"
)

// ScalingTrigger integer constant used to decide when to trigger autoscaler
const ScalingTrigger = 0

// MLPolicy contains all useful state-variable to apply the policy
type MLPolicy struct {
	scalingFactor int32
	record        *predictor.AutoscalerData
	client        predictor.ObiPredictorClient
}

// NewMLPolicy is the constructor of the MLPolicy struct
func NewMLPolicy() *MLPolicy {
	// Open predictor connection
	serverAddr := fmt.Sprintf("%s:%d",
		os.Getenv("PREDICTOR_SERVICE_DNS_NAME"),
		8080)
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure()) // TODO: encrypt communication
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return &MLPolicy{
		record: nil,
		client: predictor.NewObiPredictorClient(conn),
	}
}

// Apply is the implementation of the Policy interface
func (p *MLPolicy) Apply(metricsWindow *utils.ConcurrentSlice) int32 {
	var previousMetrics model.HeartbeatMessage
	var throughput float32
	var pendingGrowthRate float32
	var count int8
	var performance float32

	// Reset scaling factor
	p.scalingFactor = 0

	for obj := range metricsWindow.Iter() {
		if obj.Value == nil {
			continue
		}

		hb := obj.Value.(model.HeartbeatMessage)

		if previousMetrics.ClusterName != "" {
			throughput += float32(hb.AggregateContainersReleased - previousMetrics.AggregateContainersReleased)
			if hb.PendingContainers > 0 {
				memoryContainer := hb.PendingMB / hb.PendingContainers
				containersWillConsumed := hb.AvailableMB / memoryContainer
				pendingGrowth := float32(hb.PendingContainers - containersWillConsumed - previousMetrics.PendingContainers)
				if pendingGrowth > 0 {
					pendingGrowthRate += pendingGrowth
				}
			}

			count++
		}
		previousMetrics = hb
	}

	if count > 0 {
		throughput /= float32(count)
		pendingGrowthRate /= float32(count)

		performance = throughput - pendingGrowthRate // I want to maximize this

		if p.record != nil {
			// If I have scaled, send data point
			p.record.MetricsAfter = &previousMetrics
			p.record.PerformanceAfter = performance
			p.client.CollectAutoscalerData(context.Background(), p.record)
			// Clear data point
			p.record = nil
		}

		fmt.Printf("Throughput: %f\n", throughput)
		fmt.Printf("Pending rate: %f\n", pendingGrowthRate)

		// Decide whether to scale or not
		if math.Abs(float64(performance)) > ScalingTrigger {
			scalingResp, err := p.client.RequestAutoscaling(context.Background(),
				&predictor.AutoscalerRequest{
					Metrics: &previousMetrics,
					Performance: performance,
				},
			)
			logrus.WithField("response", scalingResp).Info("Asked ML autoscaler to scale")
			if err != nil {
				logrus.WithField("error", err).Error("MLAutoscaler could not generate predictions")
				p.scalingFactor = 0
			} else {
				p.scalingFactor = scalingResp.ScalingFactor
			}
		}

		// Never scale below the admitted threshold
		if previousMetrics.NumberOfNodes + p.scalingFactor < LowerBoundNodes {
			p.scalingFactor = 0
		}
	}

	if p.scalingFactor != 0 && p.record == nil {
		// Before scaling, save metrics
		p.record = &predictor.AutoscalerData{
			Nodes:             previousMetrics.NumberOfNodes,
			PerformanceBefore: performance,
			ScalingFactor:     p.scalingFactor,
			MetricsBefore:     &previousMetrics,
		}
	}

	// (Down)scaling factor should never be too low since we do not have
	// graceful decommission and jobs may crash
	//if p.scalingFactor < -5 {
	//	p.scalingFactor = -5
	//}

	// TODO: downscaling is disabled till graceful decommission is available
	if p.scalingFactor < 0 {
		p.scalingFactor = 0
	}

	return p.scalingFactor
}
