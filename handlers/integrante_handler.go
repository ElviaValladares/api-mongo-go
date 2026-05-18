package handlers

import (
	"net/http"
	"time"

	"api-mongo-go/dto"
	"api-mongo-go/services"

	"github.com/gin-gonic/gin"
	  "github.com/golang-jwt/jwt/v5"

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


var jwtSecret = []byte("mi_secret_super_seguro")

func LoginIntegrante(c *gin.Context) {
    var req struct {
        ID         string `json:"id_integrante_liga"`
        SecretPass string `json:"secret_pass"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
        return
    }
    integrante, err := integranteService.Login(req.ID, req.SecretPass)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }
    // Generar JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": integrante.ID,
        "nombre":  integrante.NombreCompleto,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "mensaje":    "Login exitoso",
        "token":      tokenString,
		"integrante": integrante,
    })
}