package routers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"sooty-tern/internal/app/controller"
	"sooty-tern/internal/app/controller/api"
)

// RegisterRouter
func RegisterRouter(app *gin.Engine, container *dig.Container) error {
	// inject api
	err := controller.Inject(container)
	if err != nil {
		return err
	}
	return container.Invoke(func(cUser *api.User) error {
		v1 := app.Group("/api/v1")
		{
			// 注册/api/v1/users
			v1.GET("/users", cUser.Query)
			v1.GET("/users/:record_id", cUser.Get)
			v1.POST("/user", cUser.Create)
			v1.PATCH("/users/:record_id/enable", cUser.Enable)
			v1.PATCH("/users/:record_id/disable", cUser.Disable)
		}
		return nil
	})
}
