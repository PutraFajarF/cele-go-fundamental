package main

import (
	"log"
	"net/http"
	"project-go/auth"
	"project-go/handler"
	"project-go/helper"
	"project-go/user"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dataConn := "host=localhost user=postgres password=ktl123 dbname=cele_go_project port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dataConn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// call repository
	userRepository := user.NewRepository(db)

	// call service
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	// call handler
	userHandler := handler.NewUserHandler(userService, authService)

	// gin router
	router := gin.Default()

	// api versioning
	api := router.Group("/api/v1")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/check-email", userHandler.CheckAvailabilityEmail)

	router.Run()
}

// Syarat gin handler parameter dalam fungsi hanya 1, jadi jika ada 2 sudah tidak memenuhi syarat seperti func authMiddleware(c *gin.Context, authService auth.Service)
// Solusinya kita buat func yg akan menjalankan func (c *gin.Context) yang akan mengembalikan func handlerFunc (func yg mengembalikan *gin.Context)
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ambil nilai header Authorization: Bearer tokentokentoken
		// dari header Authorization, kita ambil nilai tokennya saja
		// kita validasi token
		// ambil user_id
		// ambil user dari db berdasarkan user_id lewat service
		// set context isinya user

		authHeader := ctx.GetHeader("Authorization")
		// cek apakah dalam string authHeader ada "Bearer"
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.JsonResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// ambil tokennya saja
		var tokenString string
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.JsonResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		jwtClaim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.JsonResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// ubah format dari jwt.MapClaims ke float64 lalu ubah lg ke int agar sama seperti parameter GetUserByID di service.go user
		userID := int(jwtClaim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.JsonResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// set context yg isinya user dan akan dipakai pada handler yg membutuhkan context ini
		ctx.Set("currentuser", user)
	}
}
