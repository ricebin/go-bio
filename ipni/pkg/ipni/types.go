package ipni

type AuthorRecord struct {
	AlternativeAbbreviations string `json:"alternativeAbbreviations"`
	AlternativeNames         string `json:"alternativeNames"`
	BhlPageLink              string `json:"bhlPageLink"`
	Dates                    string `json:"dates"`
	FqId                     string `json:"fqId"`
	Forename                 string `json:"forename"`
	Id                       string `json:"id"`
	IsoCountries             string `json:"isoCountries"`
	RecordType               string `json:"recordType"`
	Source                   string `json:"source"`
	Summary                  string `json:"summary"`
	Surname                  string `json:"surname"`
	TaxonGroups              string `json:"taxonGroups"`
	Url                      string `json:"url"`
}

type AuthorTeam struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Summary *string `json:"summary,omitempty"`
	Type    string  `json:"type"`
	Url     string  `json:"url"`
}

type CitationRecord struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	FqId       string `json:"fqId"`
	RecordType string `json:"recordType"`
	Url        string `json:"url"`

	Family  *string `json:"family,omitempty"`
	Genus   *string `json:"genus,omitempty"`
	Species *string `json:"species,omitempty"`

	Authors             *string      `json:"authors"`
	AuthorTeam          []AuthorTeam `json:"authorTeam"`
	Publication         *string      `json:"publication,omitempty"`
	PublicationId       *string      `json:"publicationId,omitempty"`
	PublicationYear     *uint32      `json:"publicationYear,omitempty"`
	PublishingAuthor    *string      `json:"publishingAuthor,omitempty"`
	PublicationYearNote *string      `json:"publicationYearNote,omitempty"`

	NomenclaturalSynonym []CitationRecord `json:"nomenclaturalSynonym,omitempty"`
	Parent               []CitationRecord `json:"parent"`
	Child                []CitationRecord `json:"child"`
}
