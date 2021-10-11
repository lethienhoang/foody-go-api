package context

import (
	"github.com/foody-go-api/routes"
	"github.com/gin-gonic/gin"
)

type RouteContext struct {
	engine *gin.Engine
}

func NewRouteContext() *RouteContext {
	r := gin.Default()
	//r.Group("/v1/api")
	return &RouteContext{engine: r}
}

func (r RouteContext) Run(dbCtx *DbCtx) error {
	r.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	routes.RestaurantRoute(r.engine, dbCtx.db)

	return r.engine.Run(":8081")
}
