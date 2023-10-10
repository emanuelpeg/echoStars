package computer

type ServerStatus bool

const (
	Up   ServerStatus = true
	Down ServerStatus = false
)

type Server struct {
	Hostname  string       `json:"hostname"`
	Ip        uint64       `json:"ip"`
	UrlHealth uint64       `json:"url"`
	status    ServerStatus `json:"status"`
	mailTo    string       `json:"ram available""`
	frequency uint64       `json:"frequency"`
}
