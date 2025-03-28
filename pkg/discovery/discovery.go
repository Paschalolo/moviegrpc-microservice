package discovery

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Registery defines a service interface
type Registery interface {
	//Register creates a service instance record in the registy
	Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error

	//Deregister removes a service instaamce from a record
	// from the registry
	Deregister(ctx context.Context, instanceID string, serviceName string) error

	// ServiceAddresses return list of address of active instances of given service
	ServiceAddresses(ctx context.Context, serviceName string) ([]string, error)
	// ReportHealthyState is a push mechanism for reporting
	//healthy state to the registry
	ReportHealthyState(instanceID string, serviceName string) error
}

// ErrNotFound is returnes when no service addresses are found
var ErrNotFound = errors.New("no service found")

// GenerateInstanceID generates a pseudo-random service
// instance identifiers , using a service name
// suffixed by dahs and a random number
func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
