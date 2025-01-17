package types

import "errors"

func (req *ProtocolAssetRequest) ValidateBasic() error {
	if req.DestinationChain == "" {
		return errors.New("destination chain is required")
	}
	if req.TokenAddress == "" {
		return errors.New("token address is required")
	}
	if req.SourceChain == "" {
		return errors.New("source chain is required")
	}
	return nil
}
