package handler

import (
	"net/http"
	"project-go/auth"
	"project-go/helper"
	"project-go/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (u *userHandler) RegisterUser(c *gin.Context) {
	// Tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct diatas kita passing sebagai parameter service
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal register akun", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := u.userService.RegisterUser(input)
	if err != nil {
		response := helper.JsonResponse("Gagal register akun", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := u.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.JsonResponse("Gagal register akun", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := user.FormatUser(newUser, token)
	response := helper.JsonResponse("Registrasi akun berhasil", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}

func (u *userHandler) Login(c *gin.Context) {
	// user masukan input (email & password)
	// input ditangkap handler
	// mapping input user ke input struct
	// input struct passing ke service
	// service mencari dengan bantuan repository user dengan email x
	// cocokan password
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Akun gagal login", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginUser, err := u.userService.LoginUser(input)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}

		response := helper.JsonResponse("Akun gagal login", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := u.authService.GenerateToken(loginUser.ID)
	if err != nil {
		response := helper.JsonResponse("Akun gagal login", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msgFormatResp := user.FormatUser(loginUser, token)
	response := helper.JsonResponse("Login sukses", http.StatusOK, "success", msgFormatResp)

	c.JSON(http.StatusOK, response)
}

func (u *userHandler) CheckAvailabilityEmail(c *gin.Context) {
	// ada input email dari user
	// input email dimapping ke struct input
	// struct input dipassing ke service
	// servie akan manggil repository - email sudah ada atau blm
	// repository akan melakukan query ke database
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.JsonResponse("Gagal check email", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := u.userService.IsEmailAvailable(input)
	if err != nil {
		errMessage := gin.H{"errors": "Server error"}
		response := helper.JsonResponse("Gagal check email", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	messageResp := "Email sudah terdaftar"

	if isEmailAvailable {
		messageResp = "Email tersedia"
	}

	response := helper.JsonResponse(messageResp, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
