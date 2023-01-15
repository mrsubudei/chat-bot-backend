package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type translationRoutes struct {
}

func newTranslationRoutes(handler *gin.RouterGroup) {
	r := &translationRoutes{}

	h := handler.Group("/appointment")
	{
		h.GET("/history", r.history)
		//h.POST("/do-translate", r.doTranslate)
	}
}

func (r *translationRoutes) history(c *gin.Context) {

	c.JSON(http.StatusOK, "pid")
}
