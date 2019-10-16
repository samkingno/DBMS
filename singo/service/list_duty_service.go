package service


import (
    "singo/model"
	"singo/serializer"
	"time"
)

// ListDutysService 值班列表服务
type ListDutysService struct {

}

// Listzhiban 值班列表
func (service *ListDutysService) Listzhiban() serializer.Response {
    var dutys []model.Duty
	//err := model.DB.Find(&dbinstances).Error
    err := model.DB.Raw("SELECT * FROM pgdba_schedule where Begindate-interval'15 days'<=?  and enddate+interval'15 days'>=?",time.Now(),time.Now()).Scan(&dutys).Error
    if err != nil {
   return serializer.Response{
       Status: 50000,
       Msg:    "数据库连接错误",
       Error:  err.Error(),
   }
    }

    return serializer.Response{
        Data: serializer.BuildDutys(dutys),
    }
}