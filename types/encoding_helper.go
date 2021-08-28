package types

import (
	cmn "github.com/zlyzol/tendermint-0.32.3/libs/common"
)

// cdcEncode returns nil if the input is nil, otherwise returns
// cdc.MustMarshalBinaryBare(item)
func cdcEncode(item interface{}) []byte {
	if item != nil && !cmn.IsTypedNil(item) && !cmn.IsEmpty(item) {
		return cdc.MustMarshalBinaryBare(item)
	}
	return nil
}
