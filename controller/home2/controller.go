package packageB
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func Index2(ctx *gin.Context) {
	ctx.JSON( http.StatusOK, gin.H{
		"message": "Index2",
	})
}