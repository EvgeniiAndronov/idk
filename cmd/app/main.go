package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var (
	SortedBy = []string{"relevancy", "popularity", "publishedAt"}
)

type Response struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []Article
}

type Article struct {
	Sources     Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type Source struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PreviewArticle struct {
	Title      string `json:"title"`
	Url        string `json:"url"`
	UrlToImage string `json:"urlToImage"`
}

func createUrl(question string, sorted_by int) (string, error) {
	resultUrl := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&from=2024-11-18&sortBy=%s&apiKey=%s", question, SortedBy[sorted_by], API_KEY)
	return resultUrl, nil
}

func sendRequest(q string, sortedBy int) Response {
	crurl, err := createUrl(q, sortedBy)
	if err != nil {
		log.Fatal(err)
		return Response{}
	}
	resp, err := http.Get(crurl)
	if err != nil {
		log.Fatal(err)
		return Response{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return Response{}
	}

	var response Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
		return Response{}
	}

	return response
}

func makeTitles(articles []Article) []string {
	var titles []string
	for _, article := range articles {
		if article.Title != "[Removed]" {
			titles = append(titles, article.Title)
		}
	}
	return titles
}

func takeStorys(c *gin.Context) {
	q := c.Query("q")
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(
			200,
			gin.H{
				"error": "id is not int",
			},
		)
	}
	titles := makeTitles(sendRequest(q, id).Articles)
	c.JSON(200, titles)
}

func createPreviewArticle(article []Article) []PreviewArticle {
	var res []PreviewArticle
	//res := make([]PreviewArticle, len(article))
	for _, ar := range article {
		if ar.Title != "[Removed]" && ar.Title != "" {
			res = append(res, PreviewArticle{
				Title:      ar.Title,
				Url:        ar.Url,
				UrlToImage: ar.UrlToImage,
			})
		}
	}
	return res
}

func previewArticles(c *gin.Context) {
	q := c.Query("q")
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(
			200,
			gin.H{
				"error": "id is not int",
			},
		)
	}
	res := createPreviewArticle(sendRequest(q, id).Articles)
	c.JSON(200, res)
}

func setupRoutes(router *gin.Engine) {
	router.GET("/titles", takeStorys)
	router.GET("/prev", previewArticles)
}

func main() {
	router := gin.Default()

	setupRoutes(router)

	router.Run(":8080")
}
