package memory

import (
	"context"
	"errors"
	"sync"
	"time"

	"movieexample.com/pkg/discovery"
)

type serviceNameType string
type instanceIDType string

// Registry defines an inmmemory service registry
type Registry struct {
	sync.RWMutex
	ServiceAddrs map[serviceNameType]map[instanceIDType]*serviceInstance
}
type serviceInstance struct {
	hostPort   string
	lastActive time.Time
}

// New registry creates a new in memoery service
func New() *Registry {
	return &Registry{ServiceAddrs: map[serviceNameType]map[instanceIDType]*serviceInstance{}}
}

// Register creates a service record in the registry
func (r *Registry) Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.ServiceAddrs[serviceNameType(serviceName)]; !ok {
		r.ServiceAddrs[serviceNameType(serviceName)] = map[instanceIDType]*serviceInstance{}
		r.ServiceAddrs[serviceNameType(serviceName)][instanceIDType(instanceID)] = &serviceInstance{
			hostPort:   hostPort,
			lastActive: time.Now(),
		}
	}
	return nil
}

// Deregister removes a service record from the registry
func (r *Registry) Deregister(ctx context.Context, instanceID string, serviceName string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.ServiceAddrs[serviceNameType(serviceName)]; !ok {
		delete(r.ServiceAddrs[serviceNameType(serviceName)], instanceIDType(instanceID))
	}
	return nil
}

// ReportHealthyState is a push mechanism for  reporting healthy state to the registry.
func (r *Registry) ReportHealthyState(instanceID string, serviceName string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.ServiceAddrs[serviceNameType(serviceName)]; !ok {
		return errors.New("service is not registered yet ")
	}
	if _, ok := r.ServiceAddrs[serviceNameType(serviceName)][instanceIDType(instanceID)]; !ok {
		return errors.New("service instance  is not registered yet ")
	}
	r.ServiceAddrs[serviceNameType(serviceName)][instanceIDType(instanceID)].lastActive = time.Now()

	return nil
}

// ServiceAddresses returns the list of addresses of
// active instances of the given service
func (r *Registry) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
	r.RLock()
	defer r.RUnlock()
	if len(r.ServiceAddrs[serviceNameType(serviceName)]) == 0 {
		return nil, discovery.ErrNotFound
	}
	var res []string
	for _, i := range r.ServiceAddrs[serviceNameType(serviceName)] {
		if i.lastActive.Before(time.Now().Add(-5 * time.Second)) {
			continue
		}
		res = append(res, i.hostPort)
	}
	return res, nil
}
