package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB() {
	// Carrega variáveis de ambiente com fallback para .env.local
	_ = godotenv.Load(".env.local") // Ignora se não existir
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Aviso: Arquivo .env não encontrado - usando variáveis de ambiente do sistema")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ Erro: DATABASE_URL não definido. Verifique seu .env")
	}

	// Configuração do logger do GORM
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                 true,
		},
	)

	// Conexão com configurações adicionais
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	
	if err != nil {
		log.Fatalf("❌ Falha na conexão com o banco de dados: %v", err)
	}

	// Configuração da pool de conexões
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("❌ Erro ao obter instância do DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = database
	log.Println("✅ Conexão com PostgreSQL estabelecida com sucesso!")
}