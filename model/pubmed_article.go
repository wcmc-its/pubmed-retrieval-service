package model

type PubmedArticleSet struct {
	PubmedArticles []PubmedArticle `xml:"PubmedArticle"`
}

// https://www.nlm.nih.gov/bsd/licensee/elements_descriptions.html
type PubmedArticle struct {
	MedlineCitation struct {
		// Owner string `xml:"Owner,attr"`
		// Status string `xml:"Status,attr"`
		// VersionDate string `xml:"VersionDate,attr"`
		// VersionID int `xml:"VersionID,attr"`
		PMID struct {
			Value int `xml:",chardata"`
			// Version int `xml:"Version,attr"`
		} `xml:"PMID"`
		DateCompleted struct {
			Year string `xml:"Year"`
			Month string `xml:"Month"`
			Day string `xml:"Day"`
		} `xml:"DateCompleted"`
		DateCreated struct {
			Year string `xml:"Year"`
			Month string `xml:"Month"`
			Day string `xml:"Day"`
		} `xml:"DateCreated"`
		DateRevised struct {
			Year string `xml:"Year"`
			Month string `xml:"Month"`
			Day string `xml:"Day"`
		} `xml:"DateRevised"`
		Article struct {
			Journal struct {
			} `xml:"Journal"`
		} `xml:"Article"`
	} `xml:"MedlineCitation"`
}
