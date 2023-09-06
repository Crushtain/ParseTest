package main

import (
	"fmt"
	colly "github.com/gocolly/colly"
	"log"
	"strings"
)

type Product struct {
	Number string `json:"Номер"`
}

func main() {

	var result []Product

	coll := colly.NewCollector()

	coll.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %s\n", r.URL)
	})

	coll.OnHTML(".marketplace-unit.ready", func(coll *colly.HTMLElement) {
		number := coll.ChildText(".marketplace-unit__info__name")

		cleanNumber := strings.Split(number, "№")

		prod := Product{}
		prod.Number = cleanNumber[1]
		prod.Number = strings.TrimSpace(prod.Number)
		result = append(result, prod)
	})

	err := coll.Visit("https://www.fabrikant.ru/trades/procedure/search/?type=0&procedure_stage=0&price_from=&price_to=&currency=0&date_type=date_publication&date_from=&date_to=&ensure=all&count_on_page=10&order_direction=1&type_hash=1561441166&query=usb")
	if err != nil {
		log.Fatalf(err.Error())
	}

	for i, productNumber := range result {
		fmt.Println(i+1, productNumber)
	}
}
