package model

import (
	//"github.com/jinzhu/gorm"
)


// Duty 实例模型
type Duty struct {
	Id int
	Begindate  string
	Enddate      string 
	Request_person      string 
	Alert_person string
}