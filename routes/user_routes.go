package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/saim61/tasks_list_go/db"
	"github.com/saim61/tasks_list_go/user"
	"github.com/saim61/tasks_list_go/utils"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_jwt_secret") // TODO: get jwt secret from env
var userArg user.UserRequest

// RegisterUser create a new user
// @Summary Register a new user
// @Description Register yourself using your email and password
// @Tags User
// @Param user body user.UserRequest true "Required user parameters"
// @Param X-CSRF-token header string true "Insert your CSRF token. Access the GET /protected route to get it"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /register [post]
func RegisterUser(g *gin.Context) {
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

	_, _, _, status := user.GetUser(strings.ToLower(userArg.Email), database)
	if status {
		errorResponse = utils.NewErrorResponse("000x29", "Failed to create user", "User already exists")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errorCode, errorString, status := user.RegisterUser(userArg, database)
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
// @Param user body user.UserRequest true "Required user parameters"
// @Param X-CSRF-token header string true "Insert your CSRF token. Access the GET /protected route to get it"
// @Success 200 {object} user.User
// @Failure 400 {object} utils.ErrorResponse
// @Router /user [post]
func GetUser(g *gin.Context) {
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

	errorCode, errorString, user, status := user.GetUser(userArg.Email, database)
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
// @Param user body user.UserRequest true "Required user parameters"
// @Param X-CSRF-token header string true "Insert your CSRF token. Access the GET /protected route to get it"
// @Success 200 {object} utils.LoginSuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /login [post]
func LoginUser(g *gin.Context) {
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

	errorCode, errorString, user, status := user.GetUser(userArg.Email, database)
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

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.Id,
		"email":    user.Email,
		"password": user.Password,
	})

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		errorResponse = utils.NewErrorResponse("000x27", "Error while creating JWT token", "Internal server error")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Successfully login user. Copy your token to access your tasks.", "token": jwtToken})

}
