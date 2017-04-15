package model

type ESearchResult struct {
    Count int `xml:"Count"`
    // RetMax int `xml:"RetMax"`
    // RetStart int `xml:"RetStart"`
    // QueryKey int `xml:"QueryKey"`
    WebEnv string `xml:"WebEnv"`
    /*
    IdList []struct {
        Id int `xml:"Id"`
    } `xml:"IdList>Id"`
    */
}