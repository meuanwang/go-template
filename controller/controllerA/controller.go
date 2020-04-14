package packageA
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func Index1(ctx *gin.Context) {
	ctx.JSON( http.StatusOK, gin.H{
		"message": "Index1",
	})
}