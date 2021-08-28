package main

import (
	"fmt"
	"os"

	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/zlyzol/tendermint-0.32.3/crypto/encoding/amino"
)

func main() {
	cdc := amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.PrintTypes(os.Stdout)
	fmt.Println("")
}
