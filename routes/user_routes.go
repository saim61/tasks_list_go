package routes

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/saim61/tasks_list_go/db"
	userPkg "github.com/saim61/tasks_list_go/user"
	"github.com/saim61/tasks_list_go/utils"
	"golang.org/x/crypto/bcrypt"
)

var userArg userPkg.UserRequest

// RegisterUser create a new user
// @Summary Register a new user
// @Description Register yourself using your email and password
// @Tags User
// @Param user body userPkg.UserRequest true "Required user parameters"
// @Param X-CSRF-token header string true "Insert your CSRF token. Access the GET /protected route to get it"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /register [post]
func RegisterUser(g *gin.Context) {
	registerUserRequestCounter.Inc()
	log.Println("Request to register user")
	database := db.GetDatabaseObject()
	defer database.Close()

	if err := g.ShouldBindJSON(&userArg); err != nil {
		g.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	userArg.Email = strings.ToLower(userArg.Email)
	if !utils.IsValidEmail(userArg.Email) {
		errorResponse = utils.NewErrorResponse("000x28", "Failed to create user", "Invalid email format")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	_, _, _, status := userPkg.GetUser(userArg.Email, database)
	if status {
		errorResponse = utils.NewErrorResponse("000x29", "Failed to create user", "User already exists")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errorCode, errorString, status := userPkg.RegisterUser(userArg, database)
	if !status {
		errorResponse = utils.NewErrorResponse(errorCode, errorString, "Error while registering user")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	successResponse = utils.NewSuccessResponse("User registered successfully")
	log.Println(successResponse)
	g.JSON(http.StatusCreated, successResponse)
}

// GetUser get an existing user
// @Summary Get an existing user
// @Description Fetch your details by using your email and password
// @security bearerToken
// @scheme bearer
// @Tags User
// @Param user body userPkg.UserRequest true "Required user parameters"
// @Param X-CSRF-token header string true "Insert your CSRF token. Access the GET /protected route to get it"
// @Success 200 {object} userPkg.User
// @Failure 400 {object} utils.ErrorResponse
// @Router /user [post]
func GetUser(g *gin.Context) {
	getUserRequestCounter.Inc()
	log.Println("Request to get user")
	database := db.GetDatabaseObject()
	defer database.Close()

	if err := g.ShouldBindJSON(&userArg); err != nil {
		errorResponse = utils.NewErrorResponse("000x24", "Invalid parameters", "Invalid request")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	userArg.Email = strings.ToLower(userArg.Email)
	if !utils.IsValidEmail(userArg.Email) {
		errorResponse = utils.NewErrorResponse("000x30", "Failed to get user", "Invalid email format")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errorCode, errorString, user, status := userPkg.GetUser(userArg.Email, database)
	if !status {
		errorResponse = utils.NewErrorResponse(errorCode, errorString, "User doesnt exist. Please create a new user.")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	log.Println(utils.NewSuccessResponse("Successfully fetched user"))
	log.Println("User: ", user)
	g.JSON(http.StatusOK, user)
}

// Login user
// @Summary Login user
// @Description Login by using your email and password and get your token
// @Tags User
// @Param user body userPkg.UserRequest true "Required user parameters"
// @Param X-CSRF-token header string true "Insert your CSRF token. Access the GET /protected route to get it"
// @Success 200 {object} utils.LoginSuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /login [post]
func LoginUser(g *gin.Context) {
	loginRequestCounter.Inc()
	log.Println("Request to login user")
	database := db.GetDatabaseObject()
	defer database.Close()

	if err := g.ShouldBindJSON(&userArg); err != nil {
		errorResponse = utils.NewErrorResponse("000x25", "Invalid parameters", "Invalid request")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	userArg.Email = strings.ToLower(userArg.Email)
	if !utils.IsValidEmail(userArg.Email) {
		errorResponse = utils.NewErrorResponse("000x31", "Failed to login user", "Invalid email format")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errorCode, errorString, user, status := userPkg.GetUser(userArg.Email, database)
	if !status {
		errorResponse = utils.NewErrorResponse(errorCode, errorString, "User doesnt exist. Please create a new user.")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userArg.Password)) != nil {
		errorResponse = utils.NewErrorResponse("000x26", "Error while attempting login", "Invalid credentials")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	mySigningKey := []byte(os.Getenv("JWT_SECRET"))
	type MyCustomClaims struct {
		Email string
		jwt.RegisteredClaims
	}

	claims := MyCustomClaims{
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, _ := token.SignedString(mySigningKey)

	g.JSON(http.StatusOK, gin.H{"message": "Successfully login user. Copy your token to access your tasks.", "validity": "1 hour", "token": jwtToken})
}

// EditUser edit a user
// @Summary Edit a user
// @Description Edit your details
// @security bearerToken
// @scheme bearer
// @Tags User
// @Param user body userPkg.UserRequest true "Required user parameters"
// @Param X-CSRF-token header string true "Insert your CSRF token. Access the GET /protected route to get it"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /editUser [patch]
func EditUser(g *gin.Context) {
	editUserRequestCounter.Inc()
	log.Println("Request to edit user")
	database := db.GetDatabaseObject()
	defer database.Close()

	previousEmail := utils.GetUserEmailFromJWT(g.GetHeader("Authorization"))

	if err := g.ShouldBindJSON(&userArg); err != nil {
		errorResponse = utils.NewErrorResponse("000x27", "Invalid parameters", "Invalid request")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	userArg.Email = strings.ToLower(userArg.Email)
	if !utils.IsValidEmail(userArg.Email) {
		errorResponse = utils.NewErrorResponse("000x32", "Failed to login user", "Invalid email format")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	_, _, _, status := userPkg.GetUser(userArg.Email, database)
	if status {
		errorResponse = utils.NewErrorResponse("000x33", "Failed to edit user", "Email already taken")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errorCode, errorString, status := userPkg.EditUser(userArg, previousEmail, database)
	if !status {
		errorResponse = utils.NewErrorResponse(errorCode, errorString, "Failed to edit user")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	successResponse = utils.NewSuccessResponse("User details edited successfully")
	log.Println(successResponse)
	g.JSON(http.StatusOK, successResponse)
}
