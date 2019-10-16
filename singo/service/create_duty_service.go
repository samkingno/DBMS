package service

import (
    "singo/model"
    "singo/serializer"
)

// CreateDutyService 新建值班服务
type CreateDutyService struct {
	Begindate   string `form:"Begindate" json:"Begindate"`
	Enddate    string `form:"Enddate" json:"Enddate"`
	Request_person string `form:"Request_person" json:"Request_person"`
	Alert_person   string `form:"Alert_person" json:"Alert_person"`
}


// CreateDuty 新建值班
func (service *CreateDutyService) CreateDuty() serializer.Response {
	duty := model.Duty{
		Begindate:   service.Begindate,
		Enddate:    service.Enddate,
		Request_person: service.Request_person,
		Alert_person: service.Alert_person,
	}

	//err := model.DB.Create(&video).Error
	err := model.DB.Exec("Insert into pgdba_schedule  (Begindate,Enddate,Request_person,Alert_person) values (?,?,?,?)",&duty.Begindate,&duty.Enddate,&duty.Request_person,&duty.Alert_person).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "视频保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildDuty(duty),
	}
}
