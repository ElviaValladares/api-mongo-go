package handlers

import (
	"net/http"

	"api-mongo-go/dto"
	"api-mongo-go/services"

	"github.com/gin-gonic/gin"
)

var ramaService = services.RamaService{}

func CrearRama(c *gin.Context) {

	var req dto.RamaDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	err := ramaService.Crear(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Rama creada"})
}

func ListarRamas(c *gin.Context) {

	ramas, err := ramaService.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ramas)
}

func EliminarRama(c *gin.Context) {

	id := c.Param("id")

	err := ramaService.Eliminar(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Rama eliminada"})
}




func ObtenerRamaPorID(c *gin.Context) {

	id := c.Param("id")

	result, err := ramaService.ObtenerPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func ActualizarRama(c *gin.Context) {

	id := c.Param("id")

	var req dto.RamaDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	err := ramaService.Actualizar(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Actualizada correctamente"})
}
