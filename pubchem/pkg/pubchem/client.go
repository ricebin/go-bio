package pubchem

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	molecularFormula = "MolecularFormula"
	molecularWeight  = "MolecularWeight"
	canonicalSMILES  = "CanonicalSMILES"

	jsonFormat = "JSON"

	baseCompoundUrlFormat = "https://pubchem.ncbi.nlm.nih.gov/rest/pug/compound/cid/%s/property/%s/%s"
)

// https://pubchem.ncbi.nlm.nih.gov/docs/pug-rest
type Client struct {
	hc *http.Client
}

func NewClient(hc *http.Client) *Client {
	return &Client{hc: hc}
}

type Properties struct {
}
type Compound struct {
	MolecularFormula string
	MolecularWeight  string
	CanonicalSMILES  string
}

func (c *Client) GetCompounds(ctx context.Context, cids []uint32) (map[uint32]Compound, error) {
	cidStrings := make([]string, len(cids))
	for i, v := range cids {
		cidStrings[i] = strconv.Itoa(int(v))
	}

	compoundUrl := fmt.Sprintf(
		baseCompoundUrlFormat,
		strings.Join(cidStrings, ","),
		strings.Join([]string{molecularFormula, molecularWeight, canonicalSMILES}, ","),
		jsonFormat)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, compoundUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}

	// {
	//  "PropertyTable": {
	//    "Properties": [
	//      {
	//        "CID": 6049,
	//        "MolecularFormula": "C10H16N2O8",
	//        "MolecularWeight": "292.24",
	//        "CanonicalSMILES": "C(CN(CC(=O)O)CC(=O)O)N(CC(=O)O)CC(=O)O"
	//      },
	//      {
	//        "CID": 1,
	//        "MolecularFormula": "C9H17NO4",
	//        "MolecularWeight": "203.24",
	//        "CanonicalSMILES": "CC(=O)OC(CC(=O)[O-])C[N+](C)(C)C"
	//      }
	//    ]
	//  }
	//}
	type respType struct {
		PropertyTable struct {
			Properties []struct {
				CID              uint32 `json:"CID"`
				MolecularFormula string `json:"MolecularFormula"`
				MolecularWeight  string `json:"MolecularWeight"`
				CanonicalSMILES  string `json:"CanonicalSMILES"`
			} `json:"Properties"`
		} `json:"PropertyTable"`
	}
	var respBody respType
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	out := make(map[uint32]Compound)
	for _, v := range respBody.PropertyTable.Properties {
		if v.MolecularFormula != "" {
			out[v.CID] = Compound{
				MolecularFormula: v.MolecularFormula,
				MolecularWeight:  v.MolecularWeight,
				CanonicalSMILES:  v.CanonicalSMILES,
			}
		}
	}

	return out, nil
}
