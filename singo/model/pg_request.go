package model

import (
)


// Pg_request 实例模型
type Pg_request struct {
	Hostname      string  
	Port       string  
	Database    string 
	Request    string 
	Result    interface{}
}