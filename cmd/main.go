package main

import (
	"fmt"
	"github.com/bektosh/fiber-app/api"
	"github.com/bektosh/fiber-app/api/handlers"
	"github.com/bektosh/fiber-app/api/middleware"
	cfg "github.com/bektosh/fiber-app/config"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func initStorage(config cfg.Config) (*sqlx.DB, *gormadapter.Adapter) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDB,
	)
	db := sqlx.MustConnect("postgres", dsn)
	gormDB, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Println("Could not connect to db with gorm")
		panic(err)
	}

	adapter, err := gormadapter.NewAdapterByDB(gormDB)
	if err != nil {
		log.Println("Could not create new adapter")
		panic(err)
	}

	return db, adapter
}

func main() {
	config := cfg.Load()
	app := fiber.New()
	app.Use(recover.New())
	db, adapter := initStorage(config)
	handler := handlers.New(db)
	jwtRoleAuthorizer, err := middleware.NewJWTRoleAuthorizer(config, handler.Logger, adapter)
	app.Use(middleware.NewAuthorizer(jwtRoleAuthorizer))

	api.SetUpRoutes(app, handler)

	err = app.Listen(config.HTTPPort)
	if err != nil {
		panic(err)
	}
}
