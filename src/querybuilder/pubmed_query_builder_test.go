package querybuilder

import (
	"testing"
)

func TestBuildESearchQuery(t *testing.T) {
	var term string = "kukafka"
	var query string = BuildESearchQuery(term)
	var expected string = "https://www.ncbi.nlm.nih.gov/entrez/eutils/esearch.fcgi?db=pubmed&retmax=10000&usehistory=y&term=kukafka&retmode=xml"
	if query != expected {
		t.Error("Expected " + expected + ", got ", query)
	}
}

func TestBuildEFetchQuery(t *testing.T) {
	var term string = "kukafka"
	var retstart int = 0
	var webenv string = "NCID_1_163338131_130.14.22.215_9001_1492219886_1613725473_0MetA0_S_MegaStore_F_1"
	var query string = BuildEFetchQuery(term, retstart, webenv)
	var expected string = "https://www.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi?db=pubmed&retstart=0&retmax=10000&query_key=1&webenv=NCID_1_163338131_130.14.22.215_9001_1492219886_1613725473_0MetA0_S_MegaStore_F_1&retmode=xml"
	if query != expected {
		t.Error("Expected " + expected + ", got ", query)
	}
}