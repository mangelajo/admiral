package federate

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// ClusterIDLabelKey is the key for a label that may be added to federated resources to hold the ID of the cluster from
// which the resource originated, allowing for filtering of resources emanating from the originating cluster.
const ClusterIDLabelKey = "submariner-io/clusterID"

// Federator provides methods for accessing federated resources.
type Federator interface {
	// Distribute distributes the given resource to all federated clusters.
	// The actual distribution may occur asynchronously in which case any returned error only indicates that the request
	// failed.
	//
	// If the resource was previously distributed and the given resource differs, each previous cluster will receive the
	// updated resource.
	Distribute(resource runtime.Object) error

	// Delete stops distributing the given resource and deletes it from all clusters to which it was distributed.
	// The actual deletion may occur asynchronously in which any returned error only indicates that the request
	// failed.
	Delete(resource runtime.Object) error
}
