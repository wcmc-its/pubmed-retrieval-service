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

func ESearchRetrieve(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err occurred while retrieving " + url)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	parseEsearch(body)
}

func Retrieve(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err occurred while retrieving " + url)
	}
	fmt.Println("hello")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var pubmedarticles model.PubmedArticleSet
	xml.Unmarshal(body, &pubmedarticles)
	fmt.Println(pubmedarticles)
}