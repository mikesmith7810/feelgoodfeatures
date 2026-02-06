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
	"sync"
)

func ScrapeNews() ([]domain.Article, error) {
	const url = "https://www.bbc.co.uk/news"
	links, _ := parser.ExtractArticleLinks(url)

	if len(links) > 30 {
		links = links[:30]
	}

	// 1. Create a buffered channel to collect results safely from multiple threads
	articleChan := make(chan domain.Article, len(links))
	var wg sync.WaitGroup

	log.Printf("✅ Starting articles")
	for i, link := range links {
		wg.Add(1)

		go func(id int, articleLink string) {
			defer wg.Done()

			response, err := http.Get(articleLink)
			if err != nil {
				articleChan <- domain.Article{ID: id, Summary: "Network Error"}
				return
			}
			defer response.Body.Close()

			newsDocument, _ := goquery.NewDocumentFromReader(response.Body)
			articleText := parser.ExtractArticleText(newsDocument)

			prompt := fmt.Sprintf("Summarize this article passed at the end of this prompt in a short paragraph. Please do not prepend the summary with somehting liek - here is a summary.. - just the acutal news details.There is a marker of XX before the text. just give me your summary and no other text apart from at the end of the summary, can rate the article as good news or bad news and add this as a colon followed by true or false (true for good, false for bad). if the articale is funny or is about someone being stupid, then rate it as a good news story.XX\n\n%s", articleText)

			summary, err := ai_processor.GenerateArticleSummary(prompt)
			log.Printf("✅ Finished article %d", id)
			if err != nil {
				articleChan <- domain.Article{ID: id, Summary: "AI Error"}
				return
			}

			articleChan <- domain.Article{
				ID:                 id,
				Summary:            summary,
				FeelingGoodArticle: ExtractRating(summary),
			}
		}(i, link)
	}

	go func() {
		wg.Wait()
		close(articleChan)
	}()

	var articles []domain.Article
	for a := range articleChan {
		articles = append(articles, a)
	}

	log.Println("Number of news articles processed: ", len(articles))
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
