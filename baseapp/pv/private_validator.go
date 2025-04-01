package pv

import (
	cmttypes "github.com/cometbft/cometbft/types"
	pvm "github.com/cometbft/cometbft/privval"
)

type ArbitrarySignerPrivValidator interface {
	cmttypes.PrivValidator

	SignBytes(bytes []byte) ([]byte, error)
}

type ArbitrarySignerPrivValidatorFilePV struct {
	*pvm.FilePV
}

func (pVal *ArbitrarySignerPrivValidatorFilePV) SignBytes(bytes []byte) ([]byte, error) {
	return pVal.FilePV.Key.PrivKey.Sign(bytes)
}
