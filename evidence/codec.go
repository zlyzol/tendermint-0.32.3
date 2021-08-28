package evidence

import (
	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/zlyzol/tendermint-0.32.3/crypto/encoding/amino"
	"github.com/zlyzol/tendermint-0.32.3/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterEvidenceMessages(cdc)
	cryptoAmino.RegisterAmino(cdc)
	types.RegisterEvidences(cdc)
}

// For testing purposes only
func RegisterMockEvidences() {
	types.RegisterMockEvidences(cdc)
}
