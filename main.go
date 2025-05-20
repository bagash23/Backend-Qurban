package main

import (
	"fmt"
	"log"
	"masjid/auth"
	"masjid/handler"
	"masjid/helper"
	"masjid/pengurus"
	"masjid/qurban"
	"masjid/user"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {	
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

	// koneksi ke database PLN
	dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    // Create the DSN string
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	autoMigrate(db)

	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// repository
	userRepository := user.NewRepository(db)
	pengurusRepository := pengurus.NewRepository(db)
	qurbanRepository := qurban.NewRepository(db)
	// service
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	pengurusService := pengurus.NewService(pengurusRepository)
	qurbanService := qurban.NewService(qurbanRepository)
	// handler
	userHandler := handler.NewUserHandler(userService, authService)
	pengurusHandler := handler.NewPengurusHandler(pengurusService)
	qurbanHandler := handler.NewQurbanHandler(qurbanService)


	// router
	router := gin.Default()
	api := router.Group("/api/v1")
	private := router.Group("/api/private/v1")

	// Serve folder statis "images"
	router.Static("/images", "./images")

	// api General
	api.POST("/user", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)

	// api Private
	private.Use(authMiddleware(authService, userService))
	private.GET("/me", userHandler.Me)
	private.POST("/daftar-masjid",  pengurusHandler.RegisterPengurus)
	private.GET("/info-masjid", pengurusHandler.GetPengurusByUserID)
	private.GET("/search-masjid", pengurusHandler.SearchMasjid)
	// qurban
	private.PATCH("/qurban/:qurbanID", qurbanHandler.UpdateQurban) 
	private.POST("/qurban", qurbanHandler.RegisterQurban)   
	private.GET("/qurban-me", qurbanHandler.GetQurbanByPengurus)
	private.GET("/qurbans/search", qurbanHandler.GetQurbanByMasjidName)
	private.DELETE("/delete-qurban/:id", qurbanHandler.DeleteQurbanByID)



	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid{
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := fmt.Sprintf("%v", claim["id_user"])		
		user, err := userService.GetUserByID(userID)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{},
		&pengurus.Pengurus{},
		&qurban.Qurban{},
	)
	if err != nil {
		log.Fatal("Gagal melakukan auto-migrate:", err)
	}
	log.Println("✅ AutoMigrate berhasil dijalankan")
}