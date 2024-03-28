package sales

import "github.com/gin-gonic/gin"

func SalesRoute(r *gin.Engine, controller HTTPController, basePath string) {
	salesRouter := r.Group(basePath + "/sales")
	salesRouter.POST("/import", controller.Add)
}
