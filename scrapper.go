package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://saturn-imax.ru"},
		ParseFunc: parseMovies,
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()
}

func parseMovies(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("event-list").Each(func(i int, s *goquery.Selection) {
		var sessions = strings.Split(s.Find(".kLyqEz").Text(), " \n ")
		sessions = sessions[:len(sessions)-1]

		for i := 0; i < len(sessions); i++ {
			sessions[i] = strings.Trim(sessions[i], "\n ")
		}

		var description string

		if href, ok := s.Find("event-rental").Attr("href"); ok {
			g.Get(r.JoinURL(href), func(_g *geziyor.Geziyor, _r *client.Response) {
				description = _r.HTMLDoc.Find("event-info").Text()

				description = strings.ReplaceAll(description, "\t", "")
				description = strings.ReplaceAll(description, "\n", "")
				description = strings.TrimSpace(description)

				g.Exports <- map[string]interface{}{
					"title":       strings.TrimSpace(s.Find("span.movie_card_header.title").Text()),
					"subtitle":    strings.TrimSpace(s.Find("event-name").Text()),
					"sessions":    sessions,
					"description": description,
				}
			})
		}
	})
}
