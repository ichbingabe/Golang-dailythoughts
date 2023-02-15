package model

type Thought struct {
	ID      uint64 `json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var query string
