package service


import (
    "singo/model"
    "singo/serializer"
)

// ListDbinstancesService 视频列表服务
type ListDbinstancesService struct {

}

// List 视频列表
func (service *ListDbinstancesService) List() serializer.Response {
    var dbinstances []model.Db_instance

    err := model.DB.Find(&dbinstances).Error
    if err != nil {
        return serializer.Response{
            Status: 50000,
            Msg:    "数据库连接错误",
            Error:  err.Error(),
        }
    }

    return serializer.Response{
        Data: serializer.BuildDbinstances(dbinstances),
    }
}