package types

import (
	"testing"

	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"
)

func TestMarshalYAML(t *testing.T) {
	coins := NewCoins(NewCoin("okb", NewInt(1024)))
	out, err := yaml.Marshal(&coins)
	require.NoError(t, err)
	expectantStr := `- denom: okb
  amount: "1024.00000000"
`
	require.Equal(t, expectantStr, string(out))
}
