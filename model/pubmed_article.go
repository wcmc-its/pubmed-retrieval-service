package model

type PubmedArticleSet struct {
	PubmedArticles []PubmedArticle `xml:"PubmedArticle"`
}

// https://www.nlm.nih.gov/bsd/licensee/elements_descriptions.html
type PubmedArticle struct {
	MedlineCitation struct {
		Owner string `xml:"Owner,attr"`
		Status string `xml:"Status,attr"`
		VersionDate string `xml:"VersionDate,attr"`
		VersionID int `xml:"VersionID,attr"`
		PMID struct {
			Value int `xml:",chardata"`
			Version int `xml:"Version,attr"`
		} `xml:"PMID"`
	} `xml:"MedlineCitation"`
}
