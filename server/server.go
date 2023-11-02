package server

type ServerStatus bool

const (
	Up   ServerStatus = true
	Down ServerStatus = false
)

type Server struct {
	Hostname  string       `json:"hostname"`
	Ip        string       `json:"ip"`
	UrlHealth string       `json:"url"`
	Status    ServerStatus `json:"status"`
	MailTo    *string      `json:"mailTo"`
	Frequency uint64       `json:"frequency"`
}
