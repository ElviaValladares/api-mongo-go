package config

import (
	"log"
	"os"
)

func LoadConfig() *MongoConfig {
	//localmente, para desarrollo, se puede usar un archivo .env para cargar las variables de entorno
	/*err := godotenv.Load()

	if err != nil {
		log.Println(".env no encontrado")
	}*/
	//localmente, para desarrollo, se puede usar un archivo .env para cargar las variables de entorno

	usuario := os.Getenv("USUARIO_MONGO")

	if usuario == "" {
		log.Fatal("USUARIO_MONGO no definida")
	}

	pass := os.Getenv("PASS_MONGO_DB")

	if pass == "" {
		log.Fatal("PASS_MONGO_DB no definida")
	}

	return &MongoConfig{
		Usuario: usuario,
		Pass:    pass,
	}
}
