package main

import (
	"api-mongo-go/config"
	"api-mongo-go/handlers"
	"api-mongo-go/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitLogger()
	config.Logger.Info("Servidor iniciado")

	err := config.Connect(*config.LoadConfig())
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	// Configuración CORS: permitir todo origen
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// grupo protegido
	auth := router.Group("/")
	auth.Use(middleware.JWTAuth())

	//Temporadas
	auth.POST("/temporadas", handlers.CrearTemporada)
	auth.GET("/temporadas", handlers.ListarTemporadas)
	auth.GET("/temporadas/:id", handlers.ObtenerTemporadaPorID)
	auth.PUT("/temporadas/:id", handlers.ActualizarTemporada)
	auth.DELETE("/temporadas/:id", handlers.EliminarTemporada)

	//roles
	auth.POST("/roles", handlers.CrearRol)
	auth.GET("/roles", handlers.ListarRoles)
	auth.GET("/roles/:id", handlers.ObtenerRolPorID)
	auth.PUT("/roles/:id", handlers.ActualizarRol)
	auth.DELETE("/roles/:id", handlers.EliminarRol)

	auth.POST("/ramas", handlers.CrearRama)
	auth.GET("/ramas", handlers.ListarRamas)
	auth.DELETE("/ramas/:id", handlers.EliminarRama)
	auth.GET("/ramas/:id", handlers.ObtenerRamaPorID)
	auth.PUT("/ramas/:id", handlers.ActualizarRama)

	router.GET("/integrantes", handlers.ListarIntegrantes)
	router.POST("/integrantes", handlers.CrearIntegrante)
	router.POST("/integrantes/login", handlers.LoginIntegrante)
	auth.DELETE("/integrantes/:id", handlers.EliminarIntegrante)
	auth.GET("/integrantes/:id", handlers.ObtenerIntegrantePorID)
	auth.PUT("/integrantes/:id", handlers.ActualizarIntegrante)

	router.Run(":8081")
}
