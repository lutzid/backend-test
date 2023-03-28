package controllers

import (
	"auth/auth"
	"auth/database"
	"auth/models"
	"auth/utils"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LoginPayload login body
type LoginPayload struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginResponse struct {
	Token string `json:"token"`
}

type UserData struct {
	Phone     string    `json:"phone" gorm:"unique"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// Creates a user in db
func Register(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)

		ctx.JSON(400, gin.H{
			"msg": "invalid json",
		})
		ctx.Abort()

		return
	}

	generatedPassword := utils.RandString(4)

	err = user.HashPassword(generatedPassword)
	if err != nil {
		log.Println(err.Error())

		ctx.JSON(500, gin.H{
			"msg": "error hashing password",
		})
		ctx.Abort()

		return
	}

	err = user.CreateUserRecord()
	if err != nil {
		log.Println(err)

		ctx.JSON(500, gin.H{
			"msg": "error creating user",
		})
		ctx.Abort()

		return
	}

	data := UserData{
		Name:      user.Name,
		Phone:     user.Phone,
		Role:      user.Role,
		Password:  generatedPassword,
		CreatedAt: user.CreatedAt,
	}
	data.Password = generatedPassword

	utils.WriteSuccess(*ctx, "Account successfully created", data)
}

// Login logs users in
func Login(ctx *gin.Context) {
	var payload LoginPayload
	var user models.User

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(400, gin.H{
			"msg": "invalid json",
		})
		ctx.Abort()
		return
	}

	result := database.GlobalDB.Where("phone = ?", payload.Phone).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		ctx.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		ctx.Abort()
		return
	}

	err = user.CheckPassword(payload.Password)
	if err != nil {
		log.Println(err)
		ctx.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		ctx.Abort()
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	additionalClaim := auth.AdditionalJwtClaim{
		Name:      user.Name,
		Phone:     user.Phone,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.String(),
	}

	signedToken, err := jwtWrapper.GenerateToken(additionalClaim)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"msg": "error signing token",
		})
		ctx.Abort()
		return
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
	}

	ctx.JSON(200, tokenResponse)

	return
}

func JwtCheck(ctx *gin.Context) {
	clientToken := ctx.Request.Header.Get("Authorization")
	if clientToken == "" {
		ctx.JSON(403, "No Authorization header provided")
		ctx.Abort()
		return
	}

	extractedToken := strings.Split(clientToken, "Bearer ")

	if len(extractedToken) == 2 {
		clientToken = strings.TrimSpace(extractedToken[1])
	} else {
		ctx.JSON(400, "Incorrect Format of Authorization Token")
		ctx.Abort()
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey: "verysecretkey",
		Issuer:    "AuthService",
	}

	claims, err := jwtWrapper.ValidateToken(clientToken)
	if err != nil {
		ctx.JSON(401, err.Error())
		ctx.Abort()
		return
	}

	utils.WriteSuccess(*ctx, "JWT Check Success", claims)
	return
}
