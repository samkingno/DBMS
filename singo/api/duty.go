package api


import (
    "singo/service"

    "github.com/gin-gonic/gin"
)

// CreateDuty 创建值班借口
func CreateDuty(c *gin.Context) {
	service := service.CreateDutyService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateDuty()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}


// ListDuty 值班列表接口
func ListDuty(c *gin.Context) {
    service := service.ListDutysService{}
    if err := c.ShouldBind(&service); err == nil {
        res := service.Listzhiban()
        c.JSON(200, res)
    } else {
        c.JSON(200, ErrorResponse(err))
    }
}