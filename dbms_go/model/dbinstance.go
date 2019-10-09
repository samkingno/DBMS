package model

import (
	//"github.com/jinzhu/gorm"
)


// Db_instance 实例模型
type Db_instance struct {
	ID  uint
	Db_host_id     uint 
	Hostname      string 
	Port       string 
	Cluster_product_name    string 
	Cluster_product_owner      uint64 
	Pg_version string  
	Master_slave	string 
	Status         string
}