package service

import (
	"feelgoodfeatures/internal/domain"
	"feelgoodfeatures/internal/service/ai_processor"
	"feelgoodfeatures/internal/service/parser"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func ScrapeNews() ([]domain.Article, error) {
	const url = "https://www.bbc.co.uk/news"

	links, _ := parser.ExtractArticleLinks(url)

	var articles []domain.Article

	//temporarily limit results to first 3 links found
	if len(links) > 3 {
		links = links[:3]
	}

	for i, link := range links {

		response, _ := http.Get(link)

		newsDocument, _ := goquery.NewDocumentFromReader(response.Body)
		articleText := parser.ExtractArticleText(newsDocument)

		prompt := fmt.Sprintf("Summarize this articleText article in a short paragraph - just give me your summary and no other text:\n\n%s", articleText)

		summary, err := ai_processor.GenerateArticleSummary(prompt)
		if err != nil {
			log.Printf("Error summarizing article ")
			articles = append(articles, domain.Article{ID: i, Summary: "Error during generation"})
			continue
		}

		articles = append(articles, domain.Article{ID: i, Summary: summary})
	}

	log.Println("Number of news articles :%n ", len(articles))

	return articles, nil
}
