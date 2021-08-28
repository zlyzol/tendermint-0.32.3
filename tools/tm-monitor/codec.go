package main

import (
	amino "github.com/tendermint/go-amino"
	ctypes "github.com/zlyzol/tendermint-0.32.3/rpc/core/types"
)

var cdc = amino.NewCodec()

func init() {
	ctypes.RegisterAmino(cdc)
}
