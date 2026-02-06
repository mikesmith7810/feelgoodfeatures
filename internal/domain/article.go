package domain

type FeelGoodRating int

type Article struct {
	ID                 int    `json:"id"`
	Summary            string `json:"content"`
	FeelingGoodArticle bool   `json:"feel_good_article"`
}
