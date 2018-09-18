package model

import (
	"obi/master/utils"
	"sync"
	"sync/atomic"
)

// Scalable is the interface that must be implemented from a scalable cluster
type Scalable interface {
	Scale(nodes int32) bool
}

// ClusterBase is the base class for any type of cluster
type ClusterBase struct {
	Name string
	WorkerNodes int32
	ServiceType string
	HeartbeatHost string
	HeartbeatPort int
	AssignedJobs int32
	metrics *utils.ConcurrentSlice // not available outside package to prevent race conditions, get and set must be used
	sync.Mutex
}

// ClusterBaseInterface defines the primitive methods that must be implemented for any type of cluster
type ClusterBaseInterface interface {
	GetName() string
	SubmitJob(Job) error
	GetMetricsWindow() *utils.ConcurrentSlice
	AddMetricsSnapshot(message HeartbeatMessage)
	AllocateResources() error
	FreeResources() error
	AddJob()
	RemoveJob()
}


// NewClusterBase is the constructor of ClusterBase struct
// @param clusterName is the name of the cluster
// @param size is the number of nodes in the cluster
// @param platform is the cloud service environment name
// return the pointer to the ClusterBase instance
func NewClusterBase(clusterName string, workers int32, platform string, hbHost string, hbPort int) *ClusterBase {
	return &ClusterBase{
		Name:  clusterName,
		WorkerNodes: workers,
		ServiceType: platform,
		HeartbeatHost: hbHost,
		HeartbeatPort: hbPort,
		metrics: utils.NewConcurrentSlice(6, true),
	}
}

// GetMetrics is the getter of status field inside ClusterBase
// thread-safe
func (c *ClusterBase) GetMetrics() *utils.ConcurrentSlice {
	return c.metrics
}

// SetMetrics is the setter of status field inside ClusterBase
// thread-safe
func (c *ClusterBase) SetMetrics(newStatus HeartbeatMessage) {
	c.metrics.Append(newStatus)
}

func (c *ClusterBase) AddJob() {
	c.AssignedJobs = atomic.AddInt32(&c.AssignedJobs, 1)
}

func (c *ClusterBase) RemoveJob() {
	c.AssignedJobs = atomic.AddInt32(&c.AssignedJobs, -1)
}
