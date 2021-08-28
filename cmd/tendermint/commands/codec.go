package commands

import (
	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/zlyzol/tendermint-0.32.3/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
}
