package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"go-spider-study/database"
	"log"
	"os"
)

type News struct {
	Title string `json:"title"`
	Time  string `json:"time"`
}

var allnews = make([]News, 0)

func spider(url string) {
	collector := colly.NewCollector(
		colly.AllowedDomains("www.gdut.edu.cn", "gdut.edu.cn"),
	)
	collector.OnHTML(".list-search li", func(element *colly.HTMLElement) {
		newsTitle := element.ChildText("p")
		newsTime := element.ChildText("i")
		news := News{
			Title: newsTitle,
			Time:  newsTime,
		}

		allnews = append(allnews, news)
		//fmt.Println(news)
	})
	//collector.OnHTML(".list-search li i", func(element *colly.HTMLElement) {
	//	newsTime := element.Text
	//	news.Time = newsTime
	//	fmt.Println(news)
	//	allnews = append(allnews, news)
	//})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})
	collector.Visit(url)
	//fmt.Println(allnews)
}

func writeJSON2(data []News) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	_ = os.WriteFile("gdut.json", file, 0644)
	//fmt.Println(string(file))
}

func update() {
	data, err := os.ReadFile("gdut.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	var news []News
	err = json.Unmarshal(data, &news)
	if err != nil {
		fmt.Println(err)
		return
	}

	//if err != nil {
	//	print(err)
	//}
	//解析新数据
	fmt.Println(news)
	//fmt.Println(news)
	//upNews := News{}
	//err = json.Unmarshal(data, &upNews)
	//if err != nil {
	//	return
	//}
	////fmt.Println(os.Getwd())
	//fmt.Println(upNews)

}

func dbUpdate() {
}

func main() {
	db, err := database.DbInit()

	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}

	url := "https://www.gdut.edu.cn/index/tzgg.htm"
	spider(url)
	fmt.Println(allnews)
	//writeJSON2(allnews)
	db.Select("title", "time").Create(&allnews)
}
