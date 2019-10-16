package service

import (
    "singo/model"
	"singo/serializer"
	"github.com/jinzhu/gorm"
	"singo/util"
	//
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


var RQDB *gorm.DB


func RQDatabase(connString string)  {
	db, err := gorm.Open("postgres", connString)
	db.LogMode(true)
	if err != nil {
		util.Log().Error("连接数据库不成功", err)
}
RQDB =  db
}



// Pg_requestService1 发布pg_request服务
type Pg_requestService struct {
	Hostname      string  `form:"Hostname" json:"hostname"`
	Port       string  `form:"Port" json:"port"`
	Database    string `form:"Database" json:"database"`
	Request    string  `form:"Request"  json:"Request"`
	Result    string  `form:"Result"  json:"Result"`
}


// PgRequest request发布	
func (service *Pg_requestService) PgRequest() serializer.Response {
	request := model.Pg_request{
		Hostname:service.Hostname,
		Port:service.Port,
		Database:service.Database,
		Request: service.Request,
		Result:service.Result,
	}
	config_hostname := request.Hostname
	config_port := request.Port
	config_database := request.Database
	config_sql := request.Request

	connect_config := " host=" 
	connect_config += config_hostname
	connect_config += " port="
	connect_config += config_port
	connect_config += " user=zhiyinwen"
	connect_config += " dbname="
	connect_config +=config_database
	connect_config +=" sslmode=disable"
	
    //RQDatabase(connect_config)
    //Dosql(config_sql)

	connString := string(connect_config)
	requestsql := string(config_sql)



	RQDatabase(connString)

	err := RQDB.Exec(requestsql).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "sql执行失败",
			Error:  err.Error(),
		}
     }

	return serializer.Response{
		Data: serializer.BuildPgrequest(request),
	}	
}
