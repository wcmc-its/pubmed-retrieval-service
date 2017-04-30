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
		v1.GET("/article/", GetArticle)
	}
	router.Run(":8080")
}

func GetArticle(c *gin.Context) {
	var esearchResult model.ESearchResult = xmlretriever.EsearchRetrieve(querybuilder.BuildESearchQuery("bales michael"))
    var efetchurl string = querybuilder.BuildEFetchQuery(esearchResult.WebEnv, 0)
    var result model.PubmedArticleSet = xmlretriever.EfetchRetrieve(efetchurl)
	c.JSON(200, result)
}