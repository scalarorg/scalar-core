package types

import "errors"

func (req *ProtocolRequest) ValidateBasic() error {
	if req.OriginChain == "" {
		return errors.New("origin chain is required")
	}
	if req.MinorChain == "" {
		return errors.New("minor chain is required")
	}
	if req.Symbol == "" && req.Address == "" {
		return errors.New("symbol or address is required")
	}

	if req.Symbol != "" && req.Address != "" {
		return errors.New("symbol and address cannot be set together")
	}

	return nil
}
