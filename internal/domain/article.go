package domain

type Article struct {
	ID      int    `json:"id"`
	Summary string `json:"content"`
}
