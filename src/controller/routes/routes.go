package routes

import (
	"github.com/felipebs10/PrimeiroProjetoCRUD2/tree/main/src/controller"
	"github.com/felipebs10/PrimeiroProjetoCRUD2/tree/main/src"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserByid/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("createuser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)

}
