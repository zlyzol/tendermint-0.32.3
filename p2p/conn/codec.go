package conn

import (
	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/zlyzol/tendermint-0.32.3/crypto/encoding/amino"
)

var cdc *amino.Codec = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
	RegisterPacket(cdc)
}
