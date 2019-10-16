package serializer

import "singo/model"

// Dbinstance 用户序列化器
type Db_instance struct {
	ID        uint   `json:"id"`
	Db_host_id     uint `json:"db_host_id"`
	Hostname      string  `json:"hostname"`
	Port       string  `json:"port"`
	Status    string `json:"status"`
	Cluster_product_name    string  `json:"cluster_product_name"`
	Cluster_product_owner      uint64   `json:"cluster_product_owner"`
	Pg_version string   `json:"pg_version"`
	Master_slave	string   `json:"master_slave"`
}

func BuildDbinstance(item model.Db_instance) Db_instance {
	return Db_instance{

		ID:        item.ID,
		Db_host_id:     item.Db_host_id,
		Hostname:      item.Hostname,
		Port:       item.Port,
		Status:     item.Status,
		Cluster_product_name:    item.Cluster_product_name,
		Cluster_product_owner:      item.Cluster_product_owner,
		Pg_version: item.Pg_version,
		Master_slave :item.Master_slave,
	}
}

// BuildDbinstances 序列化视频列表
func BuildDbinstances(items []model.Db_instance) (dbinstances []Db_instance) {
	for _, item := range items {
		dbinstance := BuildDbinstance(item)
		dbinstances = append(dbinstances, dbinstance)
	}
	return dbinstances
}