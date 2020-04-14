package router
import (
	"project-demo/controller/controllerA"
	"project-demo/controller/controllerB"
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	r := gin.Default()
	r.GET("/packageA", packageA.Index1)
	r.GET("/packageB", packageB.Index2)
	return r
}