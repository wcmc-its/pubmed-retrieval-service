package xmlretriever

import (
	"encoding/xml"
	"net/http"
	"fmt"
	"io/ioutil"
	
	"github.com/wcmc-its/pubmed-retrieval-service/model"
)

/**
 * Given an url as a string, perform an ESearch on PubMed.
 *
 * @param url string used to query PubMed for ESearch
 * @return an ESearchResult object
 */
func EsearchRetrieve(rawUrl string) model.ESearchResult {
	resp, err := http.Get(rawUrl)
	fmt.Println("esearch query=" + rawUrl)
	if err != nil {
		fmt.Println("err occurred while retrieving " + rawUrl)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return parseEsearch(body)
}

/**
 * Given an url as a string, perofrm an EFetch on PubMed.
 *
 *
 */
func EfetchRetrieve(rawUrl string) model.PubmedArticleSet {
	resp, err := http.Get(rawUrl)
	fmt.Println("efetch query=" + rawUrl)
	if err != nil {
		fmt.Println("err occurred while retrieving " + rawUrl)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return parseEfetch(body)
}

/**
 * Given an array of bytes, unmarshall the result and return an ESearchResult
 */
func parseEsearch(bytes []byte) model.ESearchResult {
	var esearchResult model.ESearchResult
	xml.Unmarshal(bytes, &esearchResult)
	return esearchResult
}

/**
 * Given an array of bytes, unmarshall the result and return an PubmedArticleSet
 */
func parseEfetch(bytes[] byte) model.PubmedArticleSet {
	var pubmedArticleSet model.PubmedArticleSet
	xml.Unmarshal(bytes, &pubmedArticleSet)
	for _, element := range pubmedArticleSet.PubmedArticles {
		fmt.Printf("%+v\n", element)
		break
	}
	return pubmedArticleSet
}
