package context

import (
	"github.com/foody-go-api/common"
	"github.com/foody-go-api/middlewares"
	"github.com/foody-go-api/modules/restaurants/restauranthttps"
	"github.com/foody-go-api/modules/uploadfiles/uploadfilehttps"
	"github.com/gin-gonic/gin"
)

type RouteContext struct {
	Engine *gin.Engine
}

func NewRouteContext() *RouteContext {
	r := gin.Default()

	//r.Group("/v1/api")
	return &RouteContext{Engine: r}
}

func (r *AppCtx) RoutesMapping() {
	r.RouteContext.Engine.Use(middlewares.HttpResponseMiddleware())

	// routes
	r.RouteContext.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	restaurant := r.RouteContext.Engine.Group("/restaurants")
	{
		restaurant.POST("", restauranthttps.CreateRestaurantPath(r.DbCtx.DB))
		restaurant.GET("", restauranthttps.ListRestaurantPath(r.DbCtx.DB))
		restaurant.GET("/:id", restauranthttps.GetByIdRestaurantPath(r.DbCtx.DB))
		restaurant.PUT("/:id", restauranthttps.UpdateRestaurantPath(r.DbCtx.DB))
		restaurant.DELETE("/:id", restauranthttps.DeleteRestaurantPath(r.DbCtx.DB))
	}

	upload := r.RouteContext.Engine.Group("/upload")
	{
		upload.POST("", uploadfilehttps.Upload(r.s3Provider, r.DbCtx.DB))
	}

	if err := r.RouteContext.Engine.Run(":8081"); err != nil {
		panic(common.ErrInternal(err))
	}
}
