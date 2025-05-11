package main

import (
	"log"

	"github.com/felipebs10/PrimeiroProjetoCRUD2/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup) //Passando o ponteiro,

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	/*A instrução router := gin.New() criará um novo roteador gin. Os
	roteadores podem ser inicializados de duas maneiras, uma usando gin.New()
	e a outra usando gin.Default().

	A diferença é que gin.New() inicializa um roteador sem nenhum
	middleware enquanto gin.Default() inicializa o roteador com logger
	e middlewares de recovery.
	*/
}
