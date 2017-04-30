package querybuilder

import (
	"strconv"
	"net/url"
)

const esearchBaseUrl string = "https://www.ncbi.nlm.nih.gov/entrez/eutils/esearch.fcgi"
const db string = "pubmed"
const retmax int = 10000
const usehistory string = "y"
const retmode = "xml"

func BuildESearchQuery(term string) string {
	return esearchBaseUrl + 
		"?db=" + db +
		"&retmax=" + strconv.Itoa(retmax) +
		"&usehistory=" + usehistory +
		"&term=" + url.QueryEscape(term) +
		"&retmode=" + retmode
}

const efetchBaseUrl string = "https://www.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi"

func BuildEFetchQuery(webenv string, retstart int) string {
	return efetchBaseUrl +
		"?db=" + db +
		"&retstart=" + strconv.Itoa(retstart) +
		"&retmax=" + strconv.Itoa(retmax) +
		"&query_key=1" +
		"&webenv=" + webenv +
		"&retmode=" + retmode
}