package api


import (
    "singo/service"

    "github.com/gin-gonic/gin"
)

// ListDbinstances 视频列表接口
func ListDbinstances(c *gin.Context) {
    service := service.ListDbinstancesService{}
    if err := c.ShouldBind(&service); err == nil {
        res := service.List()
        c.JSON(200, res)
    } else {
        c.JSON(200, ErrorResponse(err))
    }
}