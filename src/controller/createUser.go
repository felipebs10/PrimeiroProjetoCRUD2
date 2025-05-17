package controller

import (
	"fmt"

	"github.com/felipebs10/PrimeiroProjetoCRUD2/src/configuration/rest_err"
	"github.com/felipebs10/PrimeiroProjetoCRUD2/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorret fields, error=%s", err.Error()))

		c.JSON(restErr.Code, restErr)
	}
}
