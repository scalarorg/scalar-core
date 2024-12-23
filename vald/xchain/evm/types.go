package evm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Header struct {
	ParentHash    common.Hash    `json:"parentHash"       gencodec:"required"`
	Number        *hexutil.Big   `json:"number"           gencodec:"required"`
	Time          hexutil.Uint64 `json:"timestamp"        gencodec:"required"`
	Hash          common.Hash    `json:"hash"`
	Transactions  []common.Hash  `json:"transactions"     gencodec:"required"`
	L1BlockNumber *hexutil.Big   `json:"l1BlockNumber"`
}

type FinalityOverride int

const (
	NoOverride FinalityOverride = iota
	Confirmation
)

func ParseFinalityOverride(s string) (FinalityOverride, error) {
	switch strings.ToLower(s) {
	case "":
		return NoOverride, nil
	case strings.ToLower(Confirmation.String()):
		return Confirmation, nil
	default:
		return -1, fmt.Errorf("invalid finality override option")
	}
}

var FinalityOverrideName = []string{"NoOverride", "Confirmation"}

func (f FinalityOverride) String() string {
	if f < 0 || f >= FinalityOverride(len(FinalityOverrideName)) {
		return "FinalityOverride(" + strconv.FormatInt(int64(f), 10) + ")"
	}
	return FinalityOverrideName[f]
}
