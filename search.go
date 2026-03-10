package main

import (
	"strings"

	"golang.org/x/net/html"
)

func extractText(htmlStr string) string {
	doc, _ := html.Parse(strings.NewReader(htmlStr))

	var buf strings.Builder

	var walk func(n *html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.TextNode { //node
			buf.WriteString(n.Data) //field
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling { // traversing from first child to next
			walk(c)
		}
	}

	walk(doc)
	return buf.String()
}

func extractSnippet(text, query string, radius int) string {
	/*1. Lowercase both text and query, use strings.Index to find where the match starts — call that idx
	2. Compute start = idx - radius, clamp to 0 if negative
		3. Compute end = idx + len(query) + radius, clamp to len(text) if too large
	4. Slice text[start:end] and strings.TrimSpace it
	5. If start > 0, prepend "..."
	6. If end < len(text), append "..."
	*/

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
