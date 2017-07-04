package main

import (
	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jl987-Jie/structomap"
	"github.com/pjebs/restgate"
	"github.com/wcmc-its/pubmed-retrieval-service/model"
	"github.com/wcmc-its/pubmed-retrieval-service/querybuilder"
	"github.com/wcmc-its/pubmed-retrieval-service/xmlretriever"
	"net/http"
	//"strings"
)

var r *gin.Engine

func main() {
	r = gin.Default()

	// Initialize Restgate
	rg := restgate.New("X-Auth-Key", "X-Auth-Secret", restgate.Static, restgate.Config{
		Key:                []string{"12345"},
		Secret:             []string{"secret"},
		HTTPSProtectionOff: true})

	// Create Gin middleware - integrate Restgate with Gin
	rgAdapter := func(c *gin.Context) {
		nextCalled := false
		nextAdapter := func(http.ResponseWriter, *http.Request) {
			nextCalled = true
			c.Next()
		}
		rg.ServeHTTP(c.Writer, c.Request, nextAdapter)
		if nextCalled == false {
			c.AbortWithStatus(401)
		}
	}

	// Use Restgate with Gin
	r.Use(rgAdapter)
	v1 := r.Group("v1/article/pmid=:pmid/")
	{
		// v1.GET("/articles/*term", GetArticlesByTerm)
		//       v1.GET("/count/:term", GetCountByTerm)
		//       v1.GET("/article/pmid/:pmid", GetArticleByPmid)
		//       v1.GET("/article/grantlist/:pmid", GetArticleGrantList)
		//       v1.GET("/article/meshheadinglist/:pmid", GetArticleMeshHeadingList)
		//       v1.GET("/article/otherid/:pmid", GetArticleOtherID)
		//       v1.GET("/article/medlinecitationstatus/:pmid", GetArticleMedlineCitationStatus)
		//       v1.GET("/article/publicationtypelist/:pmid", GetArticlePublicationTypeList)
		v1.GET("fields=:fields", GetArticleByFields)
	}
	r.Run(":5000")
}

// localhost:5000/api/v1/article/pmid/1234567/fields=pubmedarticle(medlinecitation(pmid))
func GetArticleByFields(c *gin.Context) {
	term := c.Params.ByName("pmid")
	fields := c.Params.ByName("fields")
	fmt.Println("fields=" + fields)
	var esearchResult model.ESearchResult = xmlretriever.EsearchRetrieve(querybuilder.BuildESearchQuery(term))
	var efetchurl string = querybuilder.BuildEFetchQuery(esearchResult.WebEnv, 0)
	var result model.PubmedArticleSet = xmlretriever.EfetchRetrieve(efetchurl)
	if len(result.PubmedArticles) > 0 {
		userSerializer := structomap.New().Pick(fields)
		userMap := userSerializer.Transform(result.PubmedArticles[0])
		c.JSON(200, userMap)
	}
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
