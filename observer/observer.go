package observer

const (
	// controllerAnnotation target annotations for deployments to be maintained for monarch election
	controllerAnnotationKey = "monarch.alpha.kubernetes.io/controller"
)

// Observer observes the cluster for the potential clients who want to get a monarch established among them
// observer consistenly queries API server for more info and correspondingly update the storage
// each set of clients is run in a separate goroutine with short sync periods to handle fast failover
type Observer struct {
}
