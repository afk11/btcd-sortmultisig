package sortutil

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"fmt"
)

func FormatPublicKey(publicKey *btcec.PublicKey, format btcutil.PubKeyFormat) ([]byte, error) {
	if btcutil.PKFHybrid == format {
		return publicKey.SerializeHybrid(), nil
	} else if btcutil.PKFCompressed == format {
		return publicKey.SerializeCompressed(), nil
	} else if btcutil.PKFUncompressed == format {
		return publicKey.SerializeUncompressed(), nil
	} else {
		return nil, fmt.Errorf("Unknown public key format value")
	}
}
