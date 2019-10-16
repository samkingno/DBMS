package serializer

import "singo/model"

// Pg_request 用户序列化器
type Pg_request struct {
	Hostname      string  `json:"hostname"`
	Port       string  `json:"port"`
	Database    string `json:"database"`
	Request    string  `json:"Request"`
	Result      interface{}   `json:"Result"`
}

// BuildPgrequest 序列化pg_request
func BuildPgrequest(item model.Pg_request) Pg_request {
	return Pg_request{
		Hostname: item.Hostname,
		Port:item.Port,
		Database:item.Database,
		Request:    item.Request,
		Result: item.Result,
	}
}
