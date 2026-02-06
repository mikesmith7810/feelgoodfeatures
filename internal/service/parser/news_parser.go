package parser

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func ExtractArticleLinks(url string) ([]string, error) {
	var articlePattern = regexp.MustCompile(`^/news/articles/[a-zA-Z0-9]+$`)

	newsPageResponse, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(newsPageResponse.Body)

	jsonNewsDocument, err := goquery.NewDocumentFromReader(newsPageResponse.Body)
	if err != nil {
		return nil, err
	}

	var links []string

	jsonNewsDocument.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && articlePattern.MatchString(href) {
			links = append(links, "https://www.bbc.co.uk"+href)
		}
	})

	return links, nil
}

func ExtractArticleText(doc *goquery.Document) string {
	var articleTextParts []string

	doc.Find("div[data-block='text'] p").Each(func(i int, s *goquery.Selection) {
		// Clean the text: trim whitespace and handle encoded characters
		text := strings.TrimSpace(s.Text())
		if text != "" {
			articleTextParts = append(articleTextParts, text)
		}
	})

	return strings.Join(articleTextParts, "\n\n")
}
