package handlers

import (
	"net/http"

	"api-mongo-go/dto"
	"api-mongo-go/services"
	"api-mongo-go/config"
	"api-mongo-go/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var rolService = services.RolService{}

func CrearRol(c *gin.Context) {
	config.Logger.WithFields(log.Fields{
	"endpoint": "/roles",
	"user": "admin",
	}).Info("Creando rol")

	var req dto.RolDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	err := rolService.Crear(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"mensaje": "Rol creado"})
}

func ListarRoles(c *gin.Context) {

	roles, err := rolService.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//utils.Error(c, 500, "error": err.Error())
		return
	}

	//c.JSON(http.StatusOK, roles)
	utils.Success(c, "Lista de roles", roles)
}



func ObtenerRolPorID(c *gin.Context) {

	id := c.Param("id")

	result, err := rolService.ObtenerPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func ActualizarRol(c *gin.Context) {

	id := c.Param("id")

	var req dto.RolDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	err := rolService.Actualizar(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Actualizada correctamente"})
}

func EliminarRol(c *gin.Context) {

	id := c.Param("id")

	err := rolService.Eliminar(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Eliminada correctamente (soft delete)"})
}
