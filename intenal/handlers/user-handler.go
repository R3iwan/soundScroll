package handlers

import (
	"net/http"

	"github.com/R3iwan/soundScroll/intenal/models"
	"github.com/R3iwan/soundScroll/intenal/services"
	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(router *gin.Context) {
	var req models.RegisterUser
	if err := router.ShouldBindBodyWithJSON(&req); err != nil {
		router.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := services.RegisterUserService(req); err != nil {
		router.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	router.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
