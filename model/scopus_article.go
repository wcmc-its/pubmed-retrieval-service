package model

type ScopusSearchResult struct {
    Entries []Entry `xml:"entry"`
}

type Entry struct {
    Prism struct {
        Url string `xml:",chardata"`
    } `xml:"prism"`
}