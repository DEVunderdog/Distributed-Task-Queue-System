package api

import (
	"net/http"

	database "github.com/DEVunderdog/auth-service/database/sqlc"
	"github.com/DEVunderdog/auth-service/middleware"
	"github.com/DEVunderdog/auth-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config utils.Config
	store  database.Store
	router *gin.Engine
}

type CORSConfig struct {
	AllowOrigins []string
	AllowMethods []string
	AllowHeaders []string
	MaxAge       int
}

type ResponseData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type StandardResponse struct {
	Message  string       `json:"message"`
	Response ResponseData `json:"response"`
}

func NewServer(config utils.Config, store database.Store) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()
	server.setupValidator()

	return server, nil
}

func (server *Server) enhanceHTTPResponse(ctx *gin.Context, httpStatusCode int, message string, data interface{}) {
	response := StandardResponse{
		Message: message,
		Response: ResponseData{
			Status: httpStatusCode,
		},
	}

	if data != nil {
		response.Response.Data = data
	}

	ctx.JSON(httpStatusCode, response)
}

func (server *Server) setupValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("domain_email", utils.ValidEmailDomain)
		v.RegisterValidation("strong_password", utils.PasswordValidator)
	}
}

func (server *Server) Start(address string) *http.Server {
	srv := &http.Server{
		Addr:    address,
		Handler: server.router,
	}

	return srv
}

func validationErrorResponse(errors validator.ValidationErrors) map[string]string {
	errorMessages := make(map[string]string)

	for _, err := range errors {
		field := err.Field()
		errorMessage := utils.ValidationErrorToText(err)
		errorMessages[field] = errorMessage
	}

	return errorMessages
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.SetTrustedProxies(nil)
	router.ForwardedByClientIP = false

	if server.config.Environment == "DEVELOPMENT" {
		router.Use(middleware.CorsMiddlewareDev())
	} else {
		router.Use(middleware.CorsMiddlewareProd(server.config.Domain))
	}

	

	router.POST("/user-signup", server.signupUser)
	router.POST("/user-login", server.loginUser)

	authRoutes := router.Group("/auth").Use(middleware.Authenticate(server.config, server.store))
	authRoutes.GET("/user-logout", server.logoutUser)

	server.router = router
}
