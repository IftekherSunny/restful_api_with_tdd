package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iftekhersunny/restful_api_with_tdd/controllers"
)

////////////////////////////////////////////
// Register api routes
////////////////////////////////////////////
func Api() *gin.Engine {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	v1 := r.Group("/v1")
	{
		// todos routes...
		todosController := new(controllers.TodosController)

		v1.GET("/todos", todosController.Index)
		v1.POST("/todos", todosController.Create)
		v1.GET("/todos/:id", todosController.Get)
		v1.PUT("/todos/:id", todosController.Update)
		v1.DELETE("/todos/:id", todosController.Delete)
	}

	return r
}
