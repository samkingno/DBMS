package serializer

import 	"singo/model"

// Duty 用户序列化器
type Duty struct {
	Id   int `json:"Id"`
	Begindate        string   `json:"Begindate"`
	Enddate      string`json:"Enddate"`
	Request_person      string  `json:"Request_person"`
	Alert_person   string  `json:"Alert_person"`
}

// BuildDuty 序列化值班
func BuildDuty(item model.Duty) Duty {
	return Duty{
		Id: item.Id,
		Begindate:        item.Begindate,
		Enddate:     item.Enddate,
		Request_person:      item.Request_person,
		Alert_person: item.Alert_person,
	}
}

// BuildDutys 序列化值班列表
func BuildDutys(items []model.Duty) (dutys []Duty) {
	for _, item := range items {
		duty := BuildDuty(item)
		dutys = append(dutys, duty)
	}
	return dutys
}