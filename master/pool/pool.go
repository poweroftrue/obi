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

package pool

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
	"obi/master/autoscaler"
	"obi/master/model"
	"sync"
	"time"
)

// Pool is the struct for clusters monitoring. Each created cluster is added in the pool
// in order to use them in all the different modules of the system.
// The pool is continuously updated, checking if all the clusters are actually alive.
type Pool struct {
	clusters sync.Map
	autoscalers sync.Map
	quit chan struct{}
	killTimeout int16
	sleepInterval int
}

// singleton instance
var poolInstance *Pool

// GetPool is for retrieving the singleton Pool struct
// return the pointer to the instance
func GetPool() *Pool {
	if poolInstance == nil {
		poolInstance = &Pool{
			sync.Map{},
			sync.Map{},
			make(chan struct{}),
			60,
			30,
		}
	}

	return poolInstance
}

// AddCluster is for adding a new cluster inside the pool
// @param cluster is a generic cluster struct
// @param autoscaler is the autoscaler object that will monitor the cluster
func (p *Pool) AddCluster(cluster model.ClusterBaseInterface, autoscaler *autoscaler.Autoscaler) {
	clusterName := cluster.GetName()
	p.clusters.Store(clusterName, cluster)
	p.autoscalers.Store(clusterName, autoscaler)
}

// RemoveCluster is for deleting a cluster from the pool, turning off its autoscaler
// @param clusterName is the name of the cluster
func (p *Pool) RemoveCluster(clusterName string) {
	// Delete cluster from pool
	p.clusters.Delete(clusterName)

	// Shutdown autoscaler
	obj, ok := p.autoscalers.Load(clusterName)
	if ok {
		obj.(*autoscaler.Autoscaler).StopMonitoring()
	}
	p.autoscalers.Delete(clusterName)
}

// LivelinessCheck is the the procedure to delete dead clusters from the pool
func (p *Pool) LivelinessCheck(timeout int16) {
	p.clusters.Range(func(key interface{}, value interface{}) bool {
		cluster := value.(model.ClusterBaseInterface)
		var lastHeartbeat model.HeartbeatMessage
		for hb := range cluster.GetMetricsWindow().Iter() {
			if hb.Value != nil {
				lastHeartbeat = hb.Value.(model.HeartbeatMessage)
			}
		}

		if lastHeartbeat.ClusterName != "" {
			lastTimestamp, _ := ptypes.Timestamp(lastHeartbeat.Timestamp)
			lastHeartbeatInterval := int16(time.Now().Sub(lastTimestamp).Seconds())
			if lastHeartbeatInterval > timeout {
				clusterName := cluster.GetName()
				logrus.WithField("Name", clusterName).Info("Deleting cluster.")
				p.RemoveCluster(clusterName)
			}
		}
		return true
	})
}

// GetCluster is for getting a specific cluster inside the pool
// @param clusterName is the name of the cluster
// return the optional object and a bool to check if it is present
func (p *Pool) GetCluster(clusterName string) (interface{}, bool) {
	return p.clusters.Load(clusterName)
}

// StartLivelinessMonitoring starts the execution of the liveliness monitor routine
func (p *Pool) StartLivelinessMonitoring() {
	logrus.Info("Starting cluster tracker routine.")
	go livelinessMonitorRoutine(poolInstance)
}

// StopLivelinessMonitoring stops the execution of the liveliness monitor routine
func (p *Pool) StopLivelinessMonitoring() {
	logrus.Info("Stopping cluster tracker routine.")
	close(p.quit)
}

// goroutine which periodically removes outdated/down clusters. It will be stop when the `quit` channel is closed
// @param pool contains all the clusters to track
// @param timeout is the time interval after which a cluster must be removed from the pool
func livelinessMonitorRoutine(pool *Pool) {

	for {
		select {
		case <-pool.quit:
			logrus.Info("Closing liveliness monitor routine.")
			return
		default:
			pool.LivelinessCheck(pool.killTimeout)
			time.Sleep(time.Duration(pool.sleepInterval) * time.Second)
		}
	}
}


