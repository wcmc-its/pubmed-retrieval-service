package xmlretriever

import (
	"encoding/xml"
	"net/http"
	"fmt"
	"io/ioutil"
	"model"
)

func Retrieve(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err occurred while retrieving " + url)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var pubmedarticles PubmedArticleSet
	xml.Unmarshal(body, pubmedarticles)
	fmt.Println(pubmedarticles)
}