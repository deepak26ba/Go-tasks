package routes

import (
	"brevo/api/handler"
	"brevo/common/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(g *gin.Engine) {

	api := g.Group("/api/v1")
	{
		g.NoRoute(func(path *gin.Context) {
			errorRespone := dto.Error{
				Message:        "Invalid Path",
				Warn:           "Use valid Path",
				HttpStatusCode: http.StatusBadRequest}
			path.JSON(http.StatusBadRequest, gin.H{
				"Error": errorRespone})
		})

		api.POST("/create-email-template", handler.CreateTemplate)
		api.GET("/get-email-templates", handler.GetTemplate)
		api.GET("/get-email-template/:id", handler.GetByIdTemplate)
		api.PUT("/update-email-template/:id", handler.UpdateTemplate)
		api.DELETE("/delete-email-template/:id", handler.DeleteTemplate)

	}

}
