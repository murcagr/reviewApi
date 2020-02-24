package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/murcagr/reviewApi/controllers"
)

// SetupRouter blablabla
func SetupRouter(r *gin.Engine) *gin.Engine {

	v1 := r.Group("/v1")
	{
		v1.GET("status", controllers.Status)
		v1.GET("getReviewByProductId", controllers.UnrealizedMethod)
		v1.GET("getReviweByUser", controllers.UnrealizedMethod)
		v1.POST("addReview", controllers.UnrealizedMethod)
		v1.PATCH("editReview", controllers.UnrealizedMethod)
	}

	return r
}
