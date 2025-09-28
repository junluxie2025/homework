package controllers

import (
	"homework/go-basic/task4/config"
	"homework/go-basic/task4/models"
	"homework/go-basic/task4/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"require"`
	Password string `json:"password" binding:"require"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func (ac *AuthController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	var existingUser models.User
	if err := config.DB.Where("username = ?", req.Username).Find(&existingUser).Error; err == nil {
		utils.BadRequest(c, "user already exists")
		return
	}

	if err := config.DB.Where("email = ?", req.Email).Find(&existingUser).Error; err == nil {
		utils.BadRequest(c, "email already exists")
		return
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.InternalServerError(c, "failed to create user")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		utils.InternalServerError(c, "failed to generate token")
	}

	utils.Success(c, AuthResponse{
		Token: token,
		User:  user,
	})
}

func (ac *AuthController) Login(c *gin.Context) {

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	//find user
	var user models.User
	if err := config.DB.Where("username=?", req.Username).Find(&user).Error; err != nil {
		utils.Unauthorized(c, "Invalid username")
	}

	//验证密码
	if user.CheckPassword(req.Password) {
		utils.Unauthorized(c, "Invalid password")
	}

	//生成Token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		utils.InternalServerError(c, "failed to generate token")
		return
	}

	utils.Success(c, AuthResponse{
		Token: token,
		User:  user,
	})

}

func (ac *AuthController) GetProfile(c *gin.Context) {

	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "user not authenticated")
		return
	}
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		utils.NotFound(c, "User not found")
	}
	utils.Success(c, user)

}
