package api


import (
    "singo/service"

    "github.com/gin-gonic/gin"
)

// Pg_request 接口
func Pg_request(c *gin.Context) {
    service := service.Pg_requestService{}
    if err := c.ShouldBind(&service); err == nil {
        res := service.PgRequest()
        c.JSON(200, res)
    } else {
        c.JSON(200, ErrorResponse(err))
    }
}