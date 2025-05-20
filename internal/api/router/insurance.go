package router

import (
	"github.com/gin-gonic/gin"
	"github.com/his-vita/dictionary-service/internal/api/middleware"
	v1 "github.com/his-vita/dictionary-service/internal/api/v1"
)

func Insurance(rg *gin.RouterGroup, insurance *v1.InsuranceHandler) {
	g := rg.Group("/insurance")
	{
		g.GET("/:id", middleware.ValidateUUIDParam("id"), insurance.GetById)
		g.GET("/list", insurance.GetList)
		g.POST("/", insurance.Create)
		g.PUT("/", insurance.Update)
		g.PATCH("/mark_deleted/:id", insurance.MarkDelete)
		g.PATCH("/unmark_deleted/:id", insurance.UnMarkDelete)
	}
}
