package handler

import (
	"html/template"
	"ex01/pkg/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

type Handler struct {
	services *services.Service
}

var limiter = rate.NewLimiter(100, 100)

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func safeHTML(s string) template.HTML {
	return template.HTML(s)
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(rateLimitMiddleware())
	router.SetFuncMap(template.FuncMap{
		"safeHTML": safeHTML,
	})
	router.LoadHTMLGlob("html/*")
	router.Static("/assets", "assets")
	auth := router.Group("/admin")
	{
		auth.GET("/log-in", h.showPage)
		auth.POST("/log-in", h.logIn)
		auth.GET("/create-post", h.showForm)
		auth.POST("/create-post", h.createPost)
	}
	blog := router.Group("/blog")
	{
		blog.GET("/", h.getPosts)
	}
	router.GET("/main", h.mainPage)
	return router
}

func rateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		c.Next()
	}
}