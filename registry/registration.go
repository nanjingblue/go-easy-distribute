package registry

type Registration struct {
	ServiceName      ServiceName   `json:"service_name"`
	ServiceURL       string        `json:"service_url"`
	RequiredServices []ServiceName `json:"required_services"`
	ServiceUpdateURL string
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
