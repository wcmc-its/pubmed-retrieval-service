package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wcmc-its/pubmed-retrieval-service/model"
    "github.com/wcmc-its/pubmed-retrieval-service/querybuilder"
    "github.com/wcmc-its/pubmed-retrieval-service/xmlretriever"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.GET("/articles/:term", GetArticles)
        v1.GET("/count/:term", GetNumberOfArticles)
	}
	router.Run(":5000")
}

func GetArticles(c *gin.Context) {
	term := c.Params.ByName("term")
	var esearchResult model.ESearchResult = xmlretriever.EsearchRetrieve(querybuilder.BuildESearchQuery(term))
    var efetchurl string = querybuilder.BuildEFetchQuery(esearchResult.WebEnv, 0)
    var result model.PubmedArticleSet = xmlretriever.EfetchRetrieve(efetchurl)
	c.JSON(200, result)
}

func GetNumberOfArticles(c *gin.Context) {
    term := c.Params.ByName("term")
    var esearchResult model.ESearchResult = xmlretriever.EsearchRetrieve(querybuilder.BuildESearchQuery(term))
    c.JSON(200, esearchResult.Count)
}