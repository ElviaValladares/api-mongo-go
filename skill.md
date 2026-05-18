# Skill: Login Integrante con JWT en Go

## Descripción
Implementa un flujo de login para un modelo de usuario (integrante) en Go, usando JWT para autenticación. El proceso incluye:
- Servicio con método `Login` que valida credenciales.
- Handler HTTP que recibe credenciales, valida y responde con JWT.
- Utilidad para generación de JWT reutilizable.
- Ejemplos de requests para pruebas.

## Estructura recomendada

### 1. Servicio (services/integrante_service.go)
```go
func (s IntegranteService) Login(id string, secretPass string) (*models.IntegranteLiga, error)
```
- Busca el integrante por ID y secret_pass.
- Retorna error si no coincide.

### 2. Handler (handlers/integrante_handler.go)
```go
func LoginIntegrante(c *gin.Context)
```
- Recibe JSON con id_integrante_liga y secret_pass.
- Llama a `Login` del servicio.
- Si es válido, genera JWT y responde con el token y datos del usuario.

### 3. Utilidad para JWT (utils/jwt.go o middleware/jwt.go)
```go
func GenerateJWT(integrante *models.IntegranteLiga) (string, error)
```
- Usa github.com/golang-jwt/jwt/v5.
- Incluye claims como user_id, nombre, exp.

### 4. Ejemplo de requests (integrantes.http)
```http
##Login Integrante
curl -X POST http://localhost:8080/integrantes/login \
  -H "Content-Type: application/json" \
  -d '{
    "id_integrante_liga": "juan123",
    "secret_pass": "123456"
  }'
```

## Referencias de archivos
- services/integrante_service.go
- handlers/integrante_handler.go
- middleware/jwt.go o utils/jwt.go
- integrantes.http

## Buenas prácticas
- Centralizar la generación de JWT en una función reutilizable.
- Proteger rutas con middleware JWT.
- Documentar los endpoints y ejemplos de uso.

---

## Plantilla de documento de referencia

### [Archivo] services/integrante_service.go
```go
func (s IntegranteService) Login(id string, secretPass string) (*models.IntegranteLiga, error) {
	// Buscar por ID y secret_pass
}
```

### [Archivo] handlers/integrante_handler.go
```go
func LoginIntegrante(c *gin.Context) {
	// Recibe JSON, valida, genera JWT y responde
}
```

### [Archivo] utils/jwt.go o middleware/jwt.go
```go
func GenerateJWT(integrante *models.IntegranteLiga) (string, error) {
	// Implementación JWT
}
```

### [Archivo] integrantes.http
```http
##Login Integrante
curl -X POST http://localhost:8080/integrantes/login \
  -H "Content-Type: application/json" \
  -d '{
    "id_integrante_liga": "juan123",
    "secret_pass": "123456"
  }'
```

---

## Repetición autónoma
1. Implementa el método Login en el servicio.
2. Implementa el handler que llama a Login y genera JWT.
3. Centraliza la generación de JWT.
4. Agrega ejemplos de requests en el archivo .http.
5. Protege rutas con middleware JWT.
