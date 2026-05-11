package handlers

import (
	"net/http"

	"api-mongo-go/dto"
	"api-mongo-go/services"

	"github.com/gin-gonic/gin"
)

var service = services.TemporadaService{}

func CrearTemporada(c *gin.Context) {

	var req dto.TemporadaDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.Crear(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Temporada creada"})
}

func ListarTemporadas(c *gin.Context) {

	result, err := service.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func ObtenerTemporadaPorID(c *gin.Context) {

	id := c.Param("id")

	result, err := service.ObtenerPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func ActualizarTemporada(c *gin.Context) {

	id := c.Param("id")

	var req dto.TemporadaDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	err := service.Actualizar(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Actualizada correctamente"})
}

func EliminarTemporada(c *gin.Context) {

	id := c.Param("id")

	err := service.Eliminar(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Eliminada correctamente (soft delete)"})
}
