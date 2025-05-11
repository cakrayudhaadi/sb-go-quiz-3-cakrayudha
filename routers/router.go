package routers

import "github.com/gin-gonic/gin"

func StartServer() {
	router := gin.Default()

	bookInitiator(router)
	categoryInitiator(router)
	userInitiator(router)

	router.Run(":8080")
}
