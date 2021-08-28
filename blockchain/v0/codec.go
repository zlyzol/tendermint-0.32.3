package v0

import (
	amino "github.com/tendermint/go-amino"
	"github.com/zlyzol/tendermint-0.32.3/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterBlockchainMessages(cdc)
	types.RegisterBlockAmino(cdc)
}
