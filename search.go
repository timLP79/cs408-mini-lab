package main

import (
	"strings"

	"golang.org/x/net/html"
)

type SearchResult struct {
	ModuleName string
	ItemTitle  string
	MatchType  string
	Snippet    string
}

func searchModules(token, baseURL string, courseID int, modules []Module, moduleItems map[int][]ModuleItem, query string) ([]SearchResult, error) {
	lower := strings.ToLower(query)
	var results []SearchResult

	for _, m := range modules {
		for _, item := range moduleItems[m.ID] {
			titleMatched := strings.Contains(strings.ToLower(item.Title), lower)
			if titleMatched {
				results = append(results, SearchResult{
					ModuleName: m.Name,
					ItemTitle:  item.Title,
					MatchType:  "title",
				})
			}

			if item.Type == "Page" && item.PageURL != "" && !titleMatched {
				content, err := fetchPageContent(token, baseURL, courseID, item.PageURL)
				if err != nil {
					continue
				}
				text := extractText(content.Body)
				if strings.Contains(strings.ToLower(text), lower) {
					results = append(results, SearchResult{
						ModuleName: m.Name,
						ItemTitle:  item.Title,
						Snippet:    extractSnippet(text, query, 80),
						MatchType:  "body",
					})
				}
			}
		}
	}

	return results, nil
}

func extractText(htmlStr string) string {
	doc, _ := html.Parse(strings.NewReader(htmlStr))

	var buf strings.Builder

	var walk func(n *html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.TextNode {
			buf.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling { // traversing from first child to next
			walk(c)
		}
	}

	walk(doc)
	return buf.String()
}

func extractSnippet(text, query string, radius int) string {
	lower := strings.ToLower(text)
	idx := strings.Index(lower, strings.ToLower(query))
	if idx < 0 {
		return ""
	}

	start := idx - radius
	if start < 0 {
		start = 0
	}

	end := idx + len(query) + radius
	if end > len(text) {
		end = len(text)
	}

	snippet := strings.TrimSpace(text[start:end])

	if start > 0 {
		snippet = "..." + snippet
	}

	if end < len(text) {
		snippet = snippet + "..."
	}

	return snippet
}
