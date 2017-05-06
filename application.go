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
		v1.GET("/articles/*term", GetArticlesByTerm)
        v1.GET("/count/:term", GetCountByTerm)
        v1.GET("/article/pmid/:pmid", GetArticleByPmid)
        v1.GET("/article/grantlist/:pmid", GetArticleGrantList)
        v1.GET("/article/meshheadinglist/:pmid", GetArticleMeshHeadingList)
        v1.GET("/article/otherid/:pmid", GetArticleOtherID)
        v1.GET("/article/medlinecitationstatus/:pmid", GetArticleMedlineCitationStatus)
        v1.GET("/article/publicationtypelist/:pmid", GetArticlePublicationTypeList)
	}
	router.Run(":5000")
}

func GetArticlesByTerm(c *gin.Context) {
	term := c.Params.ByName("term")
	var esearchResult model.ESearchResult = xmlretriever.EsearchRetrieve(querybuilder.BuildESearchQuery(term))
    var efetchurl string = querybuilder.BuildEFetchQuery(esearchResult.WebEnv, 0)
    var result model.PubmedArticleSet = xmlretriever.EfetchRetrieve(efetchurl)
	c.JSON(200, result)
}

func GetCountByTerm(c *gin.Context) {
    term := c.Params.ByName("term")
    var esearchResult model.ESearchResult = xmlretriever.EsearchRetrieve(querybuilder.BuildESearchQuery(term))
    c.JSON(200, esearchResult.Count)
}

func GetArticleByPmid(c *gin.Context) {
    term := c.Params.ByName("pmid")
    var esearchResult model.ESearchResult = xmlretriever.EsearchRetrieve(querybuilder.BuildESearchQuery(term))
    var efetchurl string = querybuilder.BuildEFetchQuery(esearchResult.WebEnv, 0)
    var result model.PubmedArticleSet = xmlretriever.EfetchRetrieve(efetchurl)
    if len(result.PubmedArticles) > 0 {
        c.JSON(200, result.PubmedArticles[0])
    }
}

func GetArticleMeshHeadingList(c *gin.Context) {
    term := c.Params.ByName("pmid")
    var esearchResult model.ESearchResult = xmlretriever.EsearchRetrieve(querybuilder.BuildESearchQuery(term))
    var efetchurl string = querybuilder.BuildEFetchQuery(esearchResult.WebEnv, 0)
    var result model.PubmedArticleSet = xmlretriever.EfetchRetrieve(efetchurl)
    if len(result.PubmedArticles) > 0 {
        c.JSON(200, result.PubmedArticles[0].MedlineCitation.MeshHeadingList)
    }
}

func GetArticleGrantList(c *gin.Context) {
    term := c.Params.ByName("pmid")
    var esearchResult model.ESearchResult = xmlretriever.EsearchRetrieve(querybuilder.BuildESearchQuery(term))
    var efetchurl string = querybuilder.BuildEFetchQuery(esearchResult.WebEnv, 0)
    var result model.PubmedArticleSet = xmlretriever.EfetchRetrieve(efetchurl)
    if len(result.PubmedArticles) > 0 {
        c.JSON(200, result.PubmedArticles[0].MedlineCitation.Article.GrantList)
    }
}

func GetArticleOtherID(c *gin.Context) {
    term := c.Params.ByName("pmid")
    var esearchResult model.ESearchResult = xmlretriever.EsearchRetrieve(querybuilder.BuildESearchQuery(term))
    var efetchurl string = querybuilder.BuildEFetchQuery(esearchResult.WebEnv, 0)
    var result model.PubmedArticleSet = xmlretriever.EfetchRetrieve(efetchurl)
    if len(result.PubmedArticles) > 0 {
        c.JSON(200, result.PubmedArticles[0].MedlineCitation.OtherID)
    }
}

func GetArticleMedlineCitationStatus(c *gin.Context) {
    term := c.Params.ByName("pmid")
    var esearchResult model.ESearchResult = xmlretriever.EsearchRetrieve(querybuilder.BuildESearchQuery(term))
    var efetchurl string = querybuilder.BuildEFetchQuery(esearchResult.WebEnv, 0)
    var result model.PubmedArticleSet = xmlretriever.EfetchRetrieve(efetchurl)
    if len(result.PubmedArticles) > 0 {
        c.JSON(200, result.PubmedArticles[0].MedlineCitation.Status)
    }
}

func GetArticlePublicationTypeList(c *gin.Context) {
    term := c.Params.ByName("pmid")
    var esearchResult model.ESearchResult = xmlretriever.EsearchRetrieve(querybuilder.BuildESearchQuery(term))
    var efetchurl string = querybuilder.BuildEFetchQuery(esearchResult.WebEnv, 0)
    var result model.PubmedArticleSet = xmlretriever.EfetchRetrieve(efetchurl)
    if len(result.PubmedArticles) > 0 {
        c.JSON(200, result.PubmedArticles[0].MedlineCitation.Article.PublicationTypeList)
    }
}
