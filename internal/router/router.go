package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// 1. Inicializa o Router utilizando as configurações default do Gin
	r := gin.Default()

	// 2. Definindo uma rota GET para o endpoint "/ping"
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 3. Vamos iniciar a nossa API, por padrão na porta 8080
	r.Run(":8080")
}
