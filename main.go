package main

import (
	"flag"
	"log"
	authApi "nom/api/routes/auth"
	"nom/database"
	"nom/middleware"
	"nom/models"
	"nom/repository"
	"nom/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type application struct {
	DSN          string
	Domain       string
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
	APIKey       string
}

type DbInstance struct {
	DB *gorm.DB
}

var Database DbInstance

func main() {
	var appConfig application

	// read from the command line
	flag.StringVar(&appConfig.DSN, "dsn", "host=localhost user=postgres password=mysecretpassword dbname=foodini-101 port=5432 sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.StringVar(&appConfig.JWTSecret, "jwt-secret", "12u8319sjhb1293bsd123i8@@@isbfsbdjb#@!@#", "signing secret")
	flag.StringVar(&appConfig.JWTIssuer, "jwt-issuer", "example.com", "signing issuer")
	flag.StringVar(&appConfig.JWTAudience, "jwt-audience", "example.com", "signing audience")
	flag.StringVar(&appConfig.CookieDomain, "cookie-domain", "localhost", "cockie domain")
	flag.StringVar(&appConfig.Domain, "domain", "example.com", "domain")
	flag.StringVar(&appConfig.APIKey, "api-key", "d094730457cbaa387c86dd3b85880a66", "api key")
	flag.Parse()

	dbConn, err := database.ConnectDB(appConfig.DSN)
	if err != nil {
		panic(err)
	}

	Database = DbInstance{
		DB: dbConn,
	}

	dbConn.AutoMigrate(
		&models.User{},
	)
	sqlDB, err := dbConn.DB()

	if err != nil {
		panic("failed to get sql.DB")
	}
	defer sqlDB.Close()

	app := fiber.New()
	app.Use(middleware.LoggerMiddleware)
	apiMain := fiber.New()

	// protected_apps.Use(middleware.AuthMiddleware())

	app.Mount("/api", apiMain)
	// app.Mount("/api", protected_apps)
	// protected_apps.Mount("/", res_app)
	// auth.SetupAuthRouts(auth_app)
	// restaurents.SetupRestaurentRouts(res_app)

	userRepo := &repository.DbRepo{DB: dbConn}
	userService := &services.AuthService{AuthRepo: userRepo}
	authApi.SetupAuthRouts(apiMain, userService)
	log.Fatal(app.Listen(":5005"))
}
