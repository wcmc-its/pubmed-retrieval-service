package xmlretriever

import (
	"testing"

    "github.com/wcmc-its/pubmed-retrieval-service/model"
)

var esearchResultXml = []byte(`
<eSearchResult>
    <Count>72</Count>
    <RetMax>1</RetMax>
    <RetStart>0</RetStart>
    <QueryKey>1</QueryKey>
    <WebEnv>NCID_1_170191730_130.14.22.215_9001_1492267750_1234158498_0MetA0_S_MegaStore_F_1</WebEnv>
    <IdList>
        <Id>28269836</Id>
    </IdList>
    <TranslationSet />
    <TranslationStack>
        <TermSet>
            <Term>Kukafka R[Author]</Term>
            <Field>Author</Field>
            <Count>72</Count>
            <Explode>N</Explode>
        </TermSet>
        <OP>GROUP</OP>
    </TranslationStack>
    <QueryTranslation>Kukafka R[Author]</QueryTranslation>
</eSearchResult>
`)

func TestESearchRetrieve(t *testing.T) {
	var esearchResult model.ESearchResult = parseEsearch(esearchResultXml)
    if esearchResult.Count != 72 {
        t.Error("Expected 72, got ", esearchResult.Count)
    }
    if esearchResult.WebEnv != "NCID_1_170191730_130.14.22.215_9001_1492267750_1234158498_0MetA0_S_MegaStore_F_1" {
        t.Error("Expected NCID_1_170191730_130.14.22.215_9001_1492267750_1234158498_0MetA0_S_MegaStore_F_1, got ", esearchResult.WebEnv)
    }
}