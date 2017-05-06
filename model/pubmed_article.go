package model

type PubmedArticleSet struct {
	PubmedArticles []PubmedArticle `xml:"PubmedArticle"`
}

type PMID struct {
	Value int `xml:",chardata"`
	Version int `xml:"Version,attr"`
}

// https://www.nlm.nih.gov/bsd/licensee/elements_descriptions.html
type PubmedArticle struct {
	MedlineCitation struct {
		Owner string `xml:"Owner,attr"`
		Status string `xml:"Status,attr"`
		VersionDate string `xml:"VersionDate,attr"`
		VersionID int `xml:"VersionID,attr"`
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
					ValidYN string `xml:"ValidYN,attr"`
					LastName string `xml:"LastName"`
					ForeName string `xml:"ForeName"`
					Initials string `xml:"Initials"`
					AffiliationInfo struct {
						Affiliation string `xml:"Affiliation"`
					} `xml:"AffiliationInfo"`
				} `xml:"Author"`
			} `xml:"AuthorList"`
			Language string `xml:"Language"`
			GrantList struct {
				CompleteYN string `xml:"CompleteYN"`
				Grant []struct {
					GrantID string `xml:"GrantID"`
					Acronym string `xml:"Acronym"`
					Agency string `xml:"Agency"`
					Country string `xml:"Country"`
				} `xml:"Grant"`
			} `xml:"GrantList"`
			PublicationTypeList struct {
				PublicationType []struct {
					UI string `xml:"UI,attr"`
					Value string `xml:",chardata"`
				} `xml:"PublicationType"`
			} `xml:"PublicationTypeList"`
		} `xml:"Article"`
		CommentsCorrectionsList struct {
			CommentsCorrections []struct {
				RefType string `xml:"RefType,attr"`
				RefSource string `xml:"RefSource"`
				PMID PMID `xml:"PMID"`
			}
		} `xml:"CommentsCorrectionsList"`
		MeshHeadingList struct {
			MeshHeading []struct {
				DescriptorName struct {
					UI string `xml:"UI,attr"`
					MajorTopicYN string `xml:"MajorTopicYN,attr"`
					Value string `xml:",chardata"`
				} `xml: "DescriptorName"`
			} `xml:"MeshHeading"`
		} `xml:"MeshHeadingList"`
		OtherID struct {
			Source string `xml:"Source,attr"`
			Value string `xml:",chardata"`
		} `xml:"OtherID"`
		PublicationStatus string `xml:"PublicationStatus"`
	} `xml:"MedlineCitation"`
	PubmedData struct {
		History struct {
			PubMedPubDate []struct {
				PubStatus string `xml:PubStatus,attr"`
				Year string `xml:"Year"`
				Month string `xml:"Month"`
				Day string `xml:"Day"`
			} `xml:"PubMedPubDate"`
		} `xml:"History"`
		PublicationStatus string `xml:"PublicationStatus"`
		ArticleIdList struct {
			ArticleId []struct {
				IdType string `xml:"IdType,attr"`
				Value string `xml:",chardata"`
			} `xml:"ArticleId"`
		} `xml:"ArticleIdList"`
	} `xml:"PubmedData"`
}
