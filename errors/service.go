package errors

const (
	// SrvUnknown : unknown service
	SrvUnknown string = "unknown service"
	// SrvBackend : backend service
	SrvBackend string = "backend service"
	// SrvEntryCache : entry cache
	SrvEntryCache string = "entry cache service"
	// SrvEntryStore : entry store
	SrvEntryStore string = "entry store service"
)

var service string

// SetService : set service name
func SetService(serviceName string) {
	service = serviceName
}
