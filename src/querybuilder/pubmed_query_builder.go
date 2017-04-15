package querybuilder

import (
	"strconv"
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
		"&term=" + term +
		"&retmode=" + retmode
}

const efetchBaseUrl string = "https://www.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi"

func BuildEFetchQuery(term string, retstart int, webenv string) string {
	return efetchBaseUrl +
		"?db=" + db +
		"&retstart=" + strconv.Itoa(retstart) +
		"&retmax=" + strconv.Itoa(retmax) +
		"&query_key=1" +
		"&webenv=" + webenv +
		"&retmode=" + retmode
}