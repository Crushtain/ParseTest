package parse

import (
	"ParseTest/iternal/clean"
	"ParseTest/iternal/newUrl"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
)

type Product struct {
	Number string `json:"Номер"`
	Name   string `json:"Название"`
}

func Parse() {
	var result []Product
	var key string

	url := newUrl.NewURL(key)

	coll := colly.NewCollector()

	coll.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %s\n", r.URL)
	})

	coll.OnHTML(".marketplace-unit.ready", func(coll *colly.HTMLElement) {
		number := coll.ChildText(".marketplace-unit__info__name")
		cleanNumber := strings.Split(number, "№")
		name := coll.ChildText(".marketplace-unit__title")

		prod := Product{}

		prod.Number = cleanNumber[1]
		prod.Number = strings.TrimSpace(prod.Number)
		prod.Number = clean.Number(prod.Number)
		prod.Name = name
		prod.Name = clean.Name(prod.Name)

		result = append(result, prod)

	})

	err := coll.Visit(url)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for i, productInfo := range result {
		fmt.Println(i+1, productInfo)
	}

	resultJSON, _ := json.MarshalIndent(result, "", "")

	err = os.WriteFile("Products.json", []byte(resultJSON), 0644)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
