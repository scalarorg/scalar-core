package types

import (
	"errors"
)

func (req *ProtocolRequest) ValidateBasic() error {
	if req.Symbol != "" && req.Address != "" && len(req.Sender) != 0 {
		return errors.New("symbol or address or sender cannot be set together")
	}

	if req.Symbol == "" && req.Address == "" && len(req.Sender) == 0 {
		return errors.New("symbol or address or sender is required")
	}

	if req.Address != "" && req.MinorChain == "" && req.OriginChain == "" {
		return errors.New("minor chain and origin chain are required when address is provided")
	}

	return nil
}
