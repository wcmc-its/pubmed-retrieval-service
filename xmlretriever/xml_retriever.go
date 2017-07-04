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
func EsearchRetrieve(url string) model.ESearchResult {
	resp, err := http.Get(url)
	fmt.Println("esearch query=" + url)
	if err != nil {
		fmt.Println("err occurred while retrieving " + url)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return parseEsearch(body)
}

/**
 * Given an url as a string, perform an EFetch on PubMed.
 */
func EfetchRetrieve(url string) model.PubmedArticleSet {
	resp, err := http.Get(url)
	fmt.Println("efetch query=" + url)
	if err != nil {
		fmt.Println("err occurred while retrieving " + url)
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
	// for _, element := range pubmedArticleSet.PubmedArticles {
	// 	fmt.Printf("%+v\n", element)
	// 	break
	// }
	return pubmedArticleSet
}
