package ipni_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/ricebin/go-bio/ipni/pkg/ipni"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetAuthor(t *testing.T) {
	c := ipni.NewClient(http.DefaultClient)
	out, err := c.GetAuthor(context.Background(), "12653-1")
	require.NoError(t, err)

	expected := ipni.AuthorRecord{
		AlternativeAbbreviations: "L. From TL2;L.or Linn. From Meikle",
		AlternativeNames:         "Linné, Carl von",
		BhlPageLink:              "http://www.biodiversitylibrary.org/openurl?ctx_ver=Z39.88-2004&rft_id=http://www.biodiversitylibrary.org/page/33355158&rft_val_fmt=info:ofi/fmt:kev:mtx:book&url_ver=z39.88-2004",
		Dates:                    "1707-1778",
		FqId:                     "urn:lsid:ipni.org:authors:12653-1",
		Forename:                 "Carl",
		Id:                       "12653-1",
		IsoCountries:             "Sweden",
		RecordType:               "author",
		Source:                   "CMI, Missouri Bryological list, Berkeley Algal list",
		Summary:                  "Linnaeus, Carl (1707-1778)",
		Surname:                  "Linnaeus",
		TaxonGroups:              "Pteridophytes, Bryophytes, Mycology, Algae, Spermatophytes",
		Url:                      "/a/12653-1",
	}
	assert.Equal(t, expected, *out)
}

func TestClient_GetIpni(t *testing.T) {
	c := ipni.NewClient(http.DefaultClient)
	out, err := c.GetCitation(context.Background(), "208258-2")
	require.NoError(t, err)

	expected := ipni.CitationRecord{
		Id:         "208258-2",
		Name:       "Populus balsamifera subsp. trichocarpa",
		FqId:       "urn:lsid:ipni.org:names:208258-2",
		RecordType: "citation",
		Url:        "/n/208258-2",
		Family:     anyPtr("Salicaceae"),
		Genus:      anyPtr("Populus"),
		Species:    anyPtr("balsamifera"),
		Authors:    anyPtr("(Torr. & A.Gray) Hultén"),
		AuthorTeam: []ipni.AuthorTeam{
			{
				Id:      "4219-1",
				Name:    "Hultén",
				Summary: anyPtr("Hultén, Oskar Eric Gunnar (1894-1981)"),
				Url:     "/a/4219-1",
				Type:    "aut",
			},
			{
				Id:      "19574-1",
				Name:    "A.Gray",
				Summary: anyPtr("Gray, Asa (1810-1888)"),
				Url:     "/a/19574-1",
				Type:    "bas",
			},
			{
				Id:      "10754-1",
				Name:    "Torr.",
				Summary: anyPtr("Torrey, John (1796-1873)"),
				Url:     "/a/10754-1",
				Type:    "bas",
			},
		},
		Publication:      anyPtr("Ark. Bot."),
		PublicationId:    anyPtr("536-2"),
		PublicationYear:  anyPtr(uint32(1968)),
		PublishingAuthor: anyPtr("Hultén"),
		NomenclaturalSynonym: []ipni.CitationRecord{
			{
				Id:         "208351-2",
				Name:       "Populus trichocarpa",
				FqId:       "urn:lsid:ipni.org:names:208351-2",
				Url:        "/n/208351-2",
				RecordType: "citation",
				Family:     anyPtr("Salicaceae"),
				Genus:      anyPtr("Populus"),
				Species:    anyPtr("trichocarpa"),
				Authors:    anyPtr("Torr. & A.Gray"),
				AuthorTeam: []ipni.AuthorTeam{
					{
						Id:   "10754-1",
						Name: "Torr.",
						Url:  "/a/10754-1",
						Type: "aut",
					},
					{
						Id:   "19574-1",
						Name: "A.Gray",
						Url:  "/a/19574-1",
						Type: "aut",
					},
				},
				PublishingAuthor: anyPtr("Torr. & A.Gray"),
			},
		},
		Parent: []ipni.CitationRecord{
			{
				Id:         "208245-2",
				Url:        "/n/208245-2",
				Name:       "Populus balsamifera",
				FqId:       "urn:lsid:ipni.org:names:208245-2",
				RecordType: "citation",
				Family:     anyPtr("Salicaceae"),
				Genus:      anyPtr("Populus"),
				Species:    anyPtr("balsamifera"),
				Authors:    anyPtr("L."),
				AuthorTeam: []ipni.AuthorTeam{
					{
						Id:   "12653-1",
						Name: "L.",
						Url:  "/a/12653-1",
						Type: "aut",
					},
				},
				Publication:         anyPtr("Sp. Pl."),
				PublicationId:       anyPtr("1071-2"),
				PublicationYear:     anyPtr(uint32(1753)),
				PublishingAuthor:    anyPtr("L."),
				PublicationYearNote: anyPtr("1 May 1753"),
			},
		},
	}

	assert.Equal(t, expected, *out)
}

func TestClient_GetIpni_WithChild(t *testing.T) {
	c := ipni.NewClient(http.DefaultClient)
	out, err := c.GetCitation(context.Background(), "208350-2")
	require.NoError(t, err)

	expected := ipni.CitationRecord{
		Id:         "208350-2",
		Name:       "Populus trichocarpa",
		FqId:       "urn:lsid:ipni.org:names:208350-2",
		RecordType: "citation",
		Url:        "/n/208350-2",
		Family:     anyPtr("Salicaceae"),
		Genus:      anyPtr("Populus"),
		Species:    anyPtr("trichocarpa"),
		Authors:    anyPtr("Torr. & A.Gray ex Hook."),
		AuthorTeam: []ipni.AuthorTeam{
			{
				Id:      "4086-1",
				Name:    "Hook.",
				Summary: anyPtr("Hooker, William Jackson (1785-1865)"),
				Url:     "/a/4086-1",
				Type:    "autEx",
			},
			{
				Id:      "19574-1",
				Name:    "A.Gray",
				Summary: anyPtr("Gray, Asa (1810-1888)"),
				Url:     "/a/19574-1",
				Type:    "aut",
			},
			{
				Id:      "10754-1",
				Name:    "Torr.",
				Summary: anyPtr("Torrey, John (1796-1873)"),
				Url:     "/a/10754-1",
				Type:    "aut",
			},
		},
		Publication:         anyPtr("Icon. Pl."),
		PublicationId:       anyPtr("16943-2"),
		PublicationYear:     anyPtr(uint32(1851)),
		PublishingAuthor:    anyPtr("Torr. & A.Gray ex Hook."),
		PublicationYearNote: anyPtr("Apr-Dec 1851"),
		NomenclaturalSynonym: []ipni.CitationRecord{
			{
				Id:         "208257-2",
				Name:       "Populus balsamifera subsp. trichocarpa",
				FqId:       "urn:lsid:ipni.org:names:208257-2",
				Url:        "/n/208257-2",
				RecordType: "citation",
				Family:     anyPtr("Salicaceae"),
				Genus:      anyPtr("Populus"),
				Species:    anyPtr("balsamifera"),
				Authors:    anyPtr("(Torr. & A.Gray) Brayshaw"),
				AuthorTeam: []ipni.AuthorTeam{
					{
						Id:   "10754-1",
						Name: "Torr.",
						Url:  "/a/10754-1",
						Type: "bas",
					},
					{
						Id:   "19574-1",
						Name: "A.Gray",
						Type: "bas",
						Url:  "/a/19574-1",
					},
					{
						Id:   "1098-1",
						Name: "Brayshaw",
						Type: "aut",
						Url:  "/a/1098-1",
					},
				},
				Publication:      anyPtr("Canad. Field-Naturalist"),
				PublicationId:    anyPtr("1982-2"),
				PublicationYear:  anyPtr(uint32(1965)),
				PublishingAuthor: anyPtr("Brayshaw"),
			},
			{
				Id:         "208258-2",
				Name:       "Populus balsamifera subsp. trichocarpa",
				FqId:       "urn:lsid:ipni.org:names:208258-2",
				RecordType: "citation",
				Url:        "/n/208258-2",
				Family:     anyPtr("Salicaceae"),
				Genus:      anyPtr("Populus"),
				Species:    anyPtr("balsamifera"),
				Authors:    anyPtr("(Torr. & A.Gray) Hultén"),
				AuthorTeam: []ipni.AuthorTeam{
					{
						Id:   "4219-1",
						Name: "Hultén",
						Url:  "/a/4219-1",
						Type: "aut",
					},
					{
						Id:   "19574-1",
						Name: "A.Gray",
						Url:  "/a/19574-1",
						Type: "bas",
					},
					{
						Id:   "10754-1",
						Name: "Torr.",
						Url:  "/a/10754-1",
						Type: "bas",
					},
				},
				Publication:      anyPtr("Ark. Bot."),
				PublicationId:    anyPtr("536-2"),
				PublicationYear:  anyPtr(uint32(1968)),
				PublishingAuthor: anyPtr("Hultén"),
			},
		},
		Parent: []ipni.CitationRecord{
			{
				Id:         "328417-2",
				Url:        "/n/328417-2",
				Name:       "Populus",
				FqId:       "urn:lsid:ipni.org:names:328417-2",
				RecordType: "citation",
				Family:     anyPtr("Salicaceae"),
				Genus:      anyPtr("Populus"),
				Authors:    anyPtr("L."),
				AuthorTeam: []ipni.AuthorTeam{
					{
						Id:   "12653-1",
						Name: "L.",
						Url:  "/a/12653-1",
						Type: "aut",
					},
				},
				Publication:         anyPtr("Sp. Pl."),
				PublicationId:       anyPtr("1071-2"),
				PublicationYear:     anyPtr(uint32(1753)),
				PublishingAuthor:    anyPtr("L."),
				PublicationYearNote: anyPtr("1 Mai 1753"),
			},
		},
		Child: []ipni.CitationRecord{
			{
				Id:         "1103964-2",
				Name:       "Populus trichocarpa var. hastata",
				FqId:       "urn:lsid:ipni.org:names:1103964-2",
				RecordType: "citation",
				Url:        "/n/1103964-2",
				Family:     anyPtr("Salicaceae"),
				Genus:      anyPtr("Populus"),
				Species:    anyPtr("trichocarpa"),
				Authors:    anyPtr("(Dode) A.Henry"),
				AuthorTeam: []ipni.AuthorTeam{
					{
						Id:   "2245-1",
						Name: "Dode",
						Type: "bas",
						Url:  "/a/2245-1",
					},
					{
						Id:   "3849-1",
						Name: "A.Henry",
						Type: "aut",
						Url:  "/a/3849-1",
					},
				},
				PublishingAuthor: anyPtr("A.Henry"),
			},
			{
				Id:         "1143851-2",
				Name:       "Populus trichocarpa subsp. hastata",
				FqId:       "urn:lsid:ipni.org:names:1143851-2",
				RecordType: "citation",
				Url:        "/n/1143851-2",
				Family:     anyPtr("Salicaceae"),
				Genus:      anyPtr("Populus"),
				Species:    anyPtr("trichocarpa"),
				Authors:    anyPtr("(Dode) Dode"),
				AuthorTeam: []ipni.AuthorTeam{
					{
						Id:   "2245-1",
						Name: "Dode",
						Type: "bas",
						Url:  "/a/2245-1",
					},
					{
						Id:   "2245-1",
						Name: "Dode",
						Type: "aut",
						Url:  "/a/2245-1",
					},
				},
				PublishingAuthor: anyPtr("Dode"),
			},
		},
	}

	//assert.Equal(t, expected, *out)
	//assert.Equal(t, spew.Sdump(expected), spew.Sdump(*out))
	assert.Equal(t, expected, *out)
}

func anyPtr[T any](in T) *T {
	return &in
}
