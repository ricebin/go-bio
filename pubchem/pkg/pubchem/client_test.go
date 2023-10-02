package pubchem_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/ricebin/pubchem/pkg/pubchem"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_Happy(t *testing.T) {
	c := pubchem.NewClient(http.DefaultClient)
	out, err := c.GetCompounds(context.Background(), []uint32{6049, 1})
	require.NoError(t, err)

	expected := map[uint32]pubchem.Compound{
		1: {
			MolecularFormula: "C9H17NO4",
			MolecularWeight:  "203.24",
			CanonicalSMILES:  "CC(=O)OC(CC(=O)[O-])C[N+](C)(C)C",
		},
		6049: {
			MolecularFormula: "C10H16N2O8",
			MolecularWeight:  "292.24",
			CanonicalSMILES:  "C(CN(CC(=O)O)CC(=O)O)N(CC(=O)O)CC(=O)O",
		},
	}

	assert.Equal(t, expected, out)
}
