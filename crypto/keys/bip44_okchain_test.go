package keys

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys/hd"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"testing"
)

func TestBip44(t *testing.T) {

	mnemonic := "depart neither they audit pen exile fire smart tongue express " +
		"blanket burden culture shove curve address " +
		"together pottery suggest lady sell clap seek whisper"
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	fmt.Printf("seed[%x]\n", seed)

	if err != nil {
		t.Fatalf("%s", err)
	}

	for accountIndex := 0; accountIndex < 6; accountIndex++ {
		for addressIndex := 0; addressIndex < 6; addressIndex++ {

			_, err := createAccount(seed, uint32(accountIndex), uint32(addressIndex))

			if err != nil {
				t.Fatalf("%s", err)
			}
		}
	}
}

func createAccount(seed []byte, accountIdx uint32, addressIdx uint32) (Info, error) {

	coinType := types.GetConfig().GetCoinType()
	hdPath := hd.NewFundraiserParams(accountIdx, coinType, addressIdx)

	fmt.Printf("HD Path[%+v]\n", hdPath)

	fullHdPath := hdPath.String()
	masterPriv, ch := hd.ComputeMastersFromSeed(seed)
	derivedPriv, err := hd.DerivePrivateKeyForPath(masterPriv, ch, fullHdPath)
	if err != nil {
		return nil, err
	}
	fmt.Printf("	PriKey  [%x]\n", derivedPriv)

	pubk := secp256k1.PrivKeySecp256k1(derivedPriv).PubKey()
	info := newOfflineInfo("", pubk)

	fmt.Printf("	PubKey  [%s]\n", info.GetPubKey())
	fmt.Printf("	Address [%s]\n", info.GetAddress())
	return info, err
}
