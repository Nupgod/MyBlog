package handler

import (
	"ex01"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

type Pagination struct {
	CurrentPage  int
	TotalPages   int
	NextPage     int
	PreviousPage int
	Posts        ex01.PostList
}

func (h *Handler) getPosts(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	postsPerPage := 3
	offset := (page - 1) * postsPerPage
	mdPosts, err := h.services.GetPosts(offset)
	var posts ex01.PostList
	if len(mdPosts) > 0 {
		posts = MdToHtml(mdPosts)
	} else {
		posts = mdPosts
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	totalPosts, err := h.services.CountPosts()
	log.Printf("Total posts: %d", totalPosts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	totalPages := (totalPosts + postsPerPage - 1) / postsPerPage

	// Create the pagination struct.
	pagination := Pagination{
		CurrentPage:  page,
		TotalPages:   totalPages,
		NextPage:     page + 1,
		PreviousPage: page - 1,
		Posts:        posts,
	}

	c.HTML(http.StatusOK, "blog.html", gin.H{
		"pagination": pagination,
	})

}

func MdToHtml(mdposts ex01.PostList) ex01.PostList {
	for i, _ := range mdposts {
		mdposts[i].Text = string(blackfriday.MarkdownCommon([]byte(mdposts[i].Text))[:])
	}
	log.Printf("Text after:\n%s\n\n", mdposts[0].Text)
	return mdposts
}
