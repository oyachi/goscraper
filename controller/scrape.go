package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/oyachi/goscraper/model"
)

func Scrape() ([]byte, error){
	wets := make([]model.Info, 2, 2)

	url := "https://tenki.jp/forecast/3/16/4410/13112/"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	doc.Find("div.forecast-days-wrap").Each(func(_ int , s *goquery.Selection) {
		// today weather information
		s.Find("section.today-weather").Each(func(_ int, tags *goquery.Selection) {
			wets[0].Date     = tags.Find("h3.left-style").Text()
			wets[0].Weather  = tags.Find("p.weather-telop").Text()
			wets[0].HighTemp = tags.Find("dd.high-temp > span.value").Text()
			wets[0].LowTemp  = tags.Find("dd.low-temp > span.value").Text()
			tags.Find("div.precip-table > table > tbody > tr").Each(func(_ int, tr *goquery.Selection) {
				attr, exists := tr.Attr("class")
				if !exists {
					tr.Find("th").Each(func(i int, th *goquery.Selection) {
						if i == 0{
							return
						}
						wets[0].Times = append(wets[0].Times, th.Text())
					})
					return
				}
				if attr == "rain-probability" {
					tr.Find("td").Each(func(_ int, td *goquery.Selection) {
						wets[0].Precipitations = append(wets[0].Precipitations, td.Text())
					})	
				}
				if attr == "wind-wave" {
					wets[0].Wind = tr.Find("td").Text()
				}
				
			})
		})
		
		// tomorrow weather information
		s.Find("section.tomorrow-weather").Each(func(_ int, tags *goquery.Selection) {
			wets[1].Date     = tags.Find("h3.left-style").Text()
			wets[1].Weather  = tags.Find("p.weather-telop").Text()
			wets[1].HighTemp = tags.Find("dd.high-temp > span.value").Text()
			wets[1].LowTemp  = tags.Find("dd.low-temp > span.value").Text()
			tags.Find("div.precip-table > table > tbody > tr").Each(func(_ int, tr *goquery.Selection) {
				attr, exists := tr.Attr("class")
				if !exists {
					tr.Find("th").Each(func(i int, th *goquery.Selection) {
						if i == 0{
							return
						}
						wets[1].Times = append(wets[1].Times, th.Text())
					})
					return
				}
				if attr == "rain-probability" {
					tr.Find("td").Each(func(_ int, td *goquery.Selection) {
						wets[1].Precipitations = append(wets[1].Precipitations, td.Text())
					})	
				}
				if attr == "wind-wave" {
					wets[1].Wind = tr.Find("td").Text()
				}
				
			})
		})
	})

	jBytes, err := json.Marshal(wets)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return jBytes, nil
}

