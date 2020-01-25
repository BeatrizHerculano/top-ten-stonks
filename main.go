package main

import (
	"fmt"
	"sort"
	. "top-ten-stonks/dao"
	. "top-ten-stonks/models"

	"github.com/gocolly/colly"
)

var dao = PapersDAO{}

func init() {

	dao.URI = "your Atlas connection URI"
	dao.Connect()
}

func main() {
	allPapers := make([]Paper, 0)

	papersColector := colly.NewCollector(
		colly.AllowedDomains("www.fundamentus.com.br"),
		colly.MaxDepth(1),
	)

	detailsCollector := colly.NewCollector()

	papersColector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	papersColector.OnHTML("td > a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		detailsCollector.Visit(e.Request.AbsoluteURL(link))
	})

	detailsCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting detail", r.URL.String())
	})
	detailsCollector.OnHTML("body", func(body *colly.HTMLElement) {
		paper := Paper{}
		body.ForEach("table tr", func(_ int, tr *colly.HTMLElement) {

			tr.ForEach("td", func(_ int, td *colly.HTMLElement) {
				if td.ChildText("span.txt") == "Papel" {
					paper.Name = tr.ChildText("td:nth-child(2) > span.txt")
				}

				if td.ChildText("span.txt") == "Valor de mercado" {
					paper.Value = tr.ChildText("td:nth-child(2) > span.txt")
				}
				if td.ChildText("span.txt") == "Empresa" {
					paper.Corp = tr.ChildText("td:nth-child(2) > span.txt")
				}
				if td.ChildText("span.txt") == "Dia" {
					paper.DayPerCent = tr.ChildText("td:nth-child(2) > span.oscil > font")
				}

			})

		})
		if paper.Value != "" {
			allPapers = append(allPapers, paper)
		}

	})
	papersColector.Visit("https://www.fundamentus.com.br/detalhes.php")
	papersColector.Wait()
	detailsCollector.Wait()
	sortedPapers := sortPapers(allPapers)
	topTenPapers := sortedPapers[0:10]
	fmt.Printf("#%v", topTenPapers)
	writeTopTenPapers(topTenPapers)

}

func sortPapers(allPapers []Paper) []Paper {
	sort.Slice(allPapers, func(i, j int) bool { return allPapers[i].Value > allPapers[j].Value })
	return allPapers
}

func writeTopTenPapers(sortedPapers []Paper) {
	for index := 0; index < 10; index++ {
		dao.Create(sortedPapers[index])
	}

}
