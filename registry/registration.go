package registry

type Registration struct {
	ServiceName      ServiceName   `json:"service_name"`
	ServiceURL       string        `json:"service_url"`
	RequiredServices []ServiceName `json:"required_services"`
	ServiceUpdateURL string        `json:"service_update_url"`
	HeartbeatURL     string        `json:"heartbeat"`
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
	PortalService  = ServiceName("Portald")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
