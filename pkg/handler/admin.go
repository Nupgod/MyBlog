package handler

import (
	"ex01"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) showPage(c *gin.Context) {
	c.HTML(http.StatusOK, "auto.html", gin.H{
		"title": "Log-in page",
	})
}

func (h *Handler) logIn(c *gin.Context) {
	username := ""
	password := ""
	username = c.PostForm("username")
	password = c.PostForm("password")
	if username == "" || password == "" {
		c.HTML(http.StatusBadRequest, "auto.html", gin.H{
			"title": "Login",
			"error": "Username and password are required.",
		})
		return
	}
	if username == "admin" && password == "admin" {
		// Redirect to the panal
		c.Redirect(http.StatusFound, "/admin/create-post")
	} else {
		// Show an error message
		c.HTML(http.StatusUnauthorized, "auto.html", gin.H{
			"title": "Login",
			"Error": "Invalid username or password.",
		})
	}
}

func (h *Handler) showForm(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-panel.html", gin.H{
		"title": "Creating post page",
	})
}

func (h *Handler) createPost(c *gin.Context) {
	var data ex01.Post
	data.Title = c.PostForm("title")
	data.Text = c.PostForm("content")
	if data.Title != "" && data.Text != "" {
		data.Time = time.Now()
		if err := h.services.CreatePost(data); err != nil {
			c.HTML(http.StatusOK, "admin-panel.html", gin.H{
				"title": "Creating post page",
				"Error": err,
			})
		} else {
			c.HTML(http.StatusOK, "admin-panel.html", gin.H{
				"title": "Creating post page",
				"Error": "Post was created",
			})
		}

	} else {
		c.HTML(http.StatusOK, "admin-panel.html", gin.H{
			"title": "Creating post page",
			"Error": "All filds must be filled in",
		})
	}
}
