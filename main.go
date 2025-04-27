package main

import (
	"TrellGo/src/config"
	"TrellGo/src/database"
	"TrellGo/src/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		db, err := database.Connect()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar no banco de dados"})
			return
		}

		sqlDB, err := db.DB() // transforma *gorm.DB em *sql.DB
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao pegar instancia SQL"})
			return
		}

		err = sqlDB.Ping()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Banco de dados inacessível"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Conexão bem sucedida com o banco de dados!"})
	})

	return r
}

func main() {
	config.Load()
	r := setupRouter()
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	// Agora use os modelos importados
	err = db.AutoMigrate(&models.User{}, &models.Project{}, &models.Task{}, &models.Comment{})
	if err != nil {
		log.Fatalf("Falha na migração: %v", err)
	}

	// Listen and Server in 0.0.0.0:8080
	fmt.Println(config.StringConnection)
	r.Run(fmt.Sprintf(":%d", config.Port))
}
