package handlers

import (
	"net/http"

	"api-mongo-go/dto"
	"api-mongo-go/services"

	"github.com/gin-gonic/gin"
)

var integranteService = services.IntegranteService{}

func CrearIntegrante(c *gin.Context) {

	var req dto.IntegranteDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	err := integranteService.Crear(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Integrante creado"})
}

func ListarIntegrantes(c *gin.Context) {

	integrantes, err := integranteService.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, integrantes)
}


func EliminarIntegrante(c *gin.Context) {

	id := c.Param("id")

	err := integranteService.Eliminar(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Integrante eliminada"})
}




func ObtenerIntegrantePorID(c *gin.Context) {

	id := c.Param("id")

	result, err := integranteService.ObtenerPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func ActualizarIntegrante(c *gin.Context) {

	id := c.Param("id")

	var req dto.IntegranteDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	err := integranteService.Actualizar(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Actualizada correctamente"})
}
