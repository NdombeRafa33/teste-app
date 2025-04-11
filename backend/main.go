package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// Estrutura do usuário
type User struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Banco de dados em memória
var (
	users = make(map[string]User)
	mu    sync.RWMutex
)

func main() {
	r := gin.Default()

	// Configuração CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// Rotas
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the API!"})
	})

	r.POST("/register", Register)
	r.POST("/login", Login)

	// Inicia o servidor
	fmt.Println("🚀 Servidor rodando na porta 8080...")
	log.Fatal(r.Run(":8080"))
}

// 🔹 Rota de Registro melhorada
func Register(c *gin.Context) {
	var user User

	// Validação dos dados
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	// Verifica se usuário já existe
	mu.RLock()
	_, exists := users[user.Email]
	mu.RUnlock()

	if exists {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email já cadastrado",
		})
		return
	}

	// Armazena o usuário
	mu.Lock()
	users[user.Email] = user
	mu.Unlock()

	fmt.Printf("✅ Usuário registrado: %+v\n", user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário registrado com sucesso!",
		"user":    user.Email,
	})
}

// 🔹 Rota de Login melhorada
func Login(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	// Verifica credenciais
	mu.RLock()
	storedUser, exists := users[user.Email]
	mu.RUnlock()

	if !exists || storedUser.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Credenciais inválidas",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login bem-sucedido!",
		"user":    user.Email,
	})
}