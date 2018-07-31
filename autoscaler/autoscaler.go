package autoscaler

import (
	"time"
	"obi/model"
)

// ScalingAlgorithm is the enum type to specify different scaling algorithms
type ScalingAlgorithm int
const (
	// BacklogBased scales the cluster to meet Time Of Completion constraints
	BacklogBased ScalingAlgorithm = iota
	// WorkloadBased scales the cluster when the resource utilization is too high
	WorkloadBased
)

// Autoscaler class with properties
type Autoscaler struct {
	Algorithm ScalingAlgorithm
	Timeout int16
	SustainedTimeout int16
	quit chan struct{}
	managedCluster model.Scalable
}

// New is the constructor of Autoscaler struct
// @param algorithm is the algorithm to follow during scaling policy execution
// @param timeoutInterval is the time interval to wait before triggering the scaling-check action again
// @param sustainedTimeoutInterval is the time interval to wait before triggering the scaling action again, when a
// 	`scale-up` or `scale-down` was triggered
// @param cluster is the scalable cluster to be managed
// return the pointer to the instance
func New(algorithm ScalingAlgorithm, timeout int16, sustainedTimeout int16, cluster model.Scalable) *Autoscaler {
	return &Autoscaler{
		algorithm,
		timeout,
		sustainedTimeout,
		make(chan struct{}),
		cluster,
	}
}


// StartMonitoringScale starts the execution of the autoscaler
func (as *Autoscaler) StartMonitoringScale() {
	go autoscalerRoutine(as)
}

// StopMonitoringScale stops the execution of the autoscaler
func (as *Autoscaler) StopMonitoringScale() {
	close(as.quit)
}

// goroutine which apply the scaling policy at each time interval. It will be stop when an empty object is inserted in
// the `quit` channel
// @param as is the autoscaler
func autoscalerRoutine(as *Autoscaler) {
	var shouldScaleUp, shouldScaleDown bool
	for {
		select {
		case <-as.quit:
			return
		default:
			shouldScaleUp, shouldScaleDown = applyPolicy(
					as.managedCluster.(model.ClusterBaseInterface).GetMetricsSnapshot(),
					as.Algorithm,
			)

			var nodes int16 = 1
			for shouldScaleUp && nodes < 128 {
				as.managedCluster.Scale(nodes, false)
				time.Sleep(time.Duration(as.SustainedTimeout) * time.Second)
				shouldScaleUp, shouldScaleDown = applyPolicy(
					as.managedCluster.(model.ClusterBaseInterface).GetMetricsSnapshot(),
					as.Algorithm,
				)
				nodes = nodes << 1
			}

			for shouldScaleDown {
				as.managedCluster.Scale(1, true)
				time.Sleep(time.Duration(as.SustainedTimeout) * time.Second)
				_, shouldScaleDown = applyPolicy(
					as.managedCluster.(model.ClusterBaseInterface).GetMetricsSnapshot(),
					as.Algorithm,
				)
			}
			time.Sleep(time.Duration(as.Timeout) * time.Second)
		}
	}
}

func applyPolicy(currentStatus model.Metrics, algorithm ScalingAlgorithm) (bool, bool) {
	switch algorithm {
	case WorkloadBased:
		// TODO
	case BacklogBased:
		// TODO
	}
	return true, true
}

