package ex01

import "time"

type Post struct {
	Title string    `json:"title" db:"title"`
	Text  string    `json:"text" db:"p_text"` 
	Time  time.Time `json:"time" db:"p_time"`
}

type PostList []Post
