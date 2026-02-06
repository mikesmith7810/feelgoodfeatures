package service

import (
	"feelgoodfeatures/internal/domain"
	"feelgoodfeatures/internal/service/ai_processor"
	"feelgoodfeatures/internal/service/parser"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func ScrapeNews() ([]domain.Article, error) {
	const url = "https://www.bbc.co.uk/news"

	links, _ := parser.ExtractArticleLinks(url)

	var articles []domain.Article

	//temporarily limit results to first 3 links found
	if len(links) > 10 {
		links = links[:10]
	}

	for i, link := range links {

		response, _ := http.Get(link)

		newsDocument, _ := goquery.NewDocumentFromReader(response.Body)
		articleText := parser.ExtractArticleText(newsDocument)

		prompt := fmt.Sprintf("Summarize this articleText article in a short paragraph - just give me your summary and no other text. At the end of the summary, can rate the article as good news or bad news and add this as a colon followed by true or false (true for good, false for bad). if the articale is funny or is about someone being stupid, then rate it as a good news story :\n\n%s", articleText)

		summary, err := ai_processor.GenerateArticleSummary(prompt)
		if err != nil {
			log.Printf("Error summarizing article ")
			articles = append(articles, domain.Article{ID: i, Summary: "Error during generation"})
			continue
		}

		articles = append(articles, domain.Article{ID: i, Summary: summary, FeelingGoodArticle: ExtractRating(summary)})
	}

	log.Println("Number of news articles : ", len(articles))

	return articles, nil
}

func ExtractRating(summary string) bool {
	parts := strings.Split(summary, ":")
	if len(parts) < 2 {
		return false
	}
	valueStr := strings.TrimSpace(parts[len(parts)-1])
	valueStr = strings.ToLower(valueStr)
	return valueStr == "true"
}
