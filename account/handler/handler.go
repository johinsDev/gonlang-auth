package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johinsDev/authentication/app"
	"github.com/johinsDev/authentication/lib/hash"
)

// Handler struct holds required services for handler to function
type Handler struct{}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R *gin.Engine
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{} // currently has no properties

	// Create an account group
	g := c.R.Group("/api/v1")

	g.GET("/me", h.Me)
	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.POST("/signout", h.Signout)
	g.POST("/tokens", h.Tokens)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)

	g.GET("/mail", h.Mail)
}

type CustomDriverConfig struct {
	hash.BaseConfig
}

type CustomDriver struct {
	Config *CustomDriverConfig
}

func (hassher *CustomDriver) Make(value string) (string, error) {
	return value, nil
}

func (hasher *CustomDriver) Verify(hashedValue string, plainValue string) (bool, error) {
	return hashedValue == plainValue, nil
}

type CustomBcrypt struct {
	Config hash.BcryptConfig
}

func (hassher *CustomBcrypt) Make(value string) (string, error) {
	return value, nil
}

func (hasher *CustomBcrypt) Verify(hashedValue string, plainValue string) (bool, error) {
	return hashedValue == plainValue, nil
}

func (h *Handler) Mail(c *gin.Context) {

	// user.Name = "johan"
	// user.Email = "johandbz@hotmail.com"

	// hasher.Extend("custom", func(data interface{}, hasher *hash.Hash) hash.HashDriverContract {
	// 	config := &CustomDriverConfig{}

	// 	hasher.BindConfig(data, config)

	// 	return &CustomDriver{
	// 		Config: config,
	// 	}
	// })

	// Mail
	// hasher
	// log
	// config
	// dig
	// folders
	// providers
	// gorm
	// default baseMiodel functions
	// auth jwt session
	// redis
	// cors
	// trotlher
	// social login
	// storage
	// mail queeue
	// queues
	// sms
	// push
	// notification
	// state

	// fmt.Println(hasher.Use("custom").Make("Holi"))
	// fmt.Println(hasher.Use("bcrypt").Make("Holi"))

	// mailer := mail.NewMailer()
	// TODO
	// migrate to gomail
	// review how handle layout
	// and data
	// mailer.To(user.Email, user.Name).Send(&mails.Welcome{
	// 	Mailable: mail.Mailable{},
	// 	User:     &user,
	// })

	// mailer.Send([]string{"layout.html", "template.html"}, struct {
	// 	Name string
	// 	URL  string
	// }{
	// 	Name: user.Name,
	// 	URL:  "Holi",
	// }, func(message *mail.Message, template *template.Template) {
	// 	message.To(user.Email, user.Name).Subject("Testing golang")
	// })

	value, _ := app.Hash().Make("bcryptme")

	c.JSON(http.StatusOK, gin.H{
		"hello": value,
	})
}

// Me handler calls services for getting
// a user's details
func (h *Handler) Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}

// Signup handler
func (h *Handler) Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signup",
	})
}

// Signin handler
func (h *Handler) Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signin",
	})
}

// Signout handler
func (h *Handler) Signout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signout",
	})
}

// Tokens handler
func (h *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's tokens",
	})
}

// Image handler
func (h *Handler) Image(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's image",
	})
}

// DeleteImage handler
func (h *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's deleteImage",
	})
}

// Details handler
func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's details",
	})
}
