package xmlretriever

import (
	"testing"
)

func TestRetrieve(t *testing.T) {
	Retrieve("https://www.ncbi.nlm.nih.gov/entrez/eutils/esearch.fcgi?db=pubmed&retmax=1&usehistory=y&term=Kukafka%20R[au]")
}