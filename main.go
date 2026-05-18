package main

import (
	"api-mongo-go/config"
	"api-mongo-go/handlers"
	"api-mongo-go/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitLogger()
	config.Logger.Info("Servidor iniciado")

	err := config.Connect()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.POST("/productos", handlers.CrearProducto)
	router.GET("/productos", handlers.ObtenerProductos)

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

	router.POST("/integrantes", handlers.CrearIntegrante)
	router.POST("/integrantes/login", handlers.LoginIntegrante)
	auth.GET("/integrantes", handlers.ListarIntegrantes)
	auth.DELETE("/integrantes/:id", handlers.EliminarIntegrante)
	auth.GET("/integrantes/:id", handlers.ObtenerIntegrantePorID)
	auth.PUT("/integrantes/:id", handlers.ActualizarIntegrante)

	router.Run(":443")
}
