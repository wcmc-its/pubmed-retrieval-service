package xmlretriever

import (
	"encoding/xml"
	"net/http"
	"fmt"
	"io/ioutil"
	
	"github.com/wcmc-its/pubmed-retrieval-service/model"
)

func parseEsearch(bytes []byte) model.ESearchResult {
	var esearchResult model.ESearchResult
	xml.Unmarshal(bytes, &esearchResult)
	return esearchResult
}

func EsearchRetrieve(url string) model.ESearchResult {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err occurred while retrieving " + url)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return parseEsearch(body)
}

func parseEfetch(bytes[] byte) model.PubmedArticleSet {
	var pubmedArticleSet model.PubmedArticleSet
	xml.Unmarshal(bytes, &pubmedArticleSet)
	fmt.Println(pubmedArticleSet)
	return pubmedArticleSet
}

func EfetchRetrieve(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err occurred while retrieving " + url)
	}
	fmt.Println("hello")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	parseEfetch(body)
}