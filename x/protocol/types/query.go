package types

import "errors"

func (req *ProtocolRequest) ValidateBasic() error {
	if req.Symbol != "" && req.Address != "" {
		return errors.New("symbol and address cannot be set together")
	}

	if req.Symbol == "" && req.Address == "" {
		return errors.New("symbol or address is required")
	}

	if req.Address != "" && req.MinorChain == "" && req.OriginChain == "" {
		return errors.New("minor chain and origin chain are required when address is provided")
	}

	return nil
}
