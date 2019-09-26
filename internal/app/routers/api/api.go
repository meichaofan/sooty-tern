package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"sooty-tern/internal/app/ginplus"
	"sooty-tern/internal/app/routers/api/ctl"
	"sooty-tern/internal/app/schema"
)

// RegisterRouter 注册/api路由
func RegisterRouter(app *gin.Engine, container *dig.Container) error {
	err := ctl.Inject(container)
	if err != nil {
		return err
	}
	return container.Invoke(func() error {
		g := app.Group("/api")
		g.GET("/version", func(c *gin.Context) {
			data := make(map[string]string)
			data["version"] = "v1"
			ginplus.ResData(c, data)
		})

		g.GET("/listUser", func(c *gin.Context) {
			d := map[string]string{
				"name": "meichaofan",
				"sex":  "boy",
			}
			pr := &schema.PaginationResult{
				Total: 12,
			}
			ginplus.ResList(c, d, pr)
		})

		return nil
	})
}
