package model

type PubmedArticleSet struct {
	PubmedArticles []PubmedArticle `xml:"PubmedArticle"`
}

type PMID struct {
	Value int `xml:",chardata"`
	// Version int `xml:"Version,attr"`
}

// https://www.nlm.nih.gov/bsd/licensee/elements_descriptions.html
type PubmedArticle struct {
	MedlineCitation struct {
		// Owner string `xml:"Owner,attr"`
		// Status string `xml:"Status,attr"`
		// VersionDate string `xml:"VersionDate,attr"`
		// VersionID int `xml:"VersionID,attr"`
		PMID PMID `xml:"PMID"`
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
				ISSN struct {
					Value string `xml:",chardata"`
					// IssnType string `xml:"IssnType,attr"`
				}
				JournalIssue struct {
					CitedMedium string `xml:"CitedMedium,attr"`
					Volume int `xml:"Volume"`
					PubDate struct {
						Year string `xml:"Year"`
						Month string `xml:"Month"`
						Day string `xml:"Day"`
					} `xml:"PubDate"`
				} `xml:"JournalIssue"`
				Title string `xml:"Title"`
				ISOAbbreviation string `xml:"ISOAbbreviation"`
			} `xml:"Journal"`
			ArticleTitle string `xml:"ArticleTitle"`
			AuthorList struct {
				Author []struct {
					ValidYN rune `xml:"ValidYN,attr"`
					LastName string `xml:"LastName"`
					ForeName string `xml:"ForeName"`
					Initials string `xml:"Initials"`
					AffiliationInfo struct {
						Affiliation string `xml:"Affiliation"`
					} `xml:"AffiliationInfo"`
				} `xml:"Author"`
			} `xml:"AuthorList"`
		} `xml:"Article"`
		CommentsCorrectionsList struct {
			CommentsCorrections []struct {
				RefType string `xml:"RefType,attr"`
				RefSource string `xml:"RefSource"`
				PMID PMID `xml:"PMID"`
			}
		}
	} `xml:"MedlineCitation"`
}
