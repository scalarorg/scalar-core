package exported

import (
	"encoding/binary"
	"encoding/hex"
	fmt "fmt"
	"strings"

	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

func (p *ProtocolInfo) GetKeyID(pk []byte) multisig.KeyID {
	return multisig.KeyID(hex.EncodeToString(pk))
}

func (p *ProtocolInfo) IsSupportedChain(chain nexus.ChainName) bool {
	for _, c := range p.MinorAddresses {
		if c.ChainName == chain {
			return true
		}
	}
	return false
}

func FormatContractCallWithTokenToBTCKeyID(bitcoinPubKey []byte, model LiquidityModel) (multisig.KeyID, error) {
	if _, ok := LiquidityModel_name[int32(model)]; !ok {
		return "", fmt.Errorf("FormatContractCallWithTokenToBTCKeyID > invalid model")
	}

	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(model))

	return multisig.KeyID(hex.EncodeToString(bytes) + "|" + hex.EncodeToString(bitcoinPubKey)), nil
}

func ParseContractCallWithTokenToBTCKeyID(keyID multisig.KeyID) (bitcoinPubKey []byte, model LiquidityModel, err error) {
	parts := strings.Split(string(keyID), "|")
	if len(parts) != 2 {
		return nil, LiquidityModel(0), fmt.Errorf("ParseContractCallWithTokenToBTCKeyID > invalid keyID")
	}

	modelBytes, err := hex.DecodeString(parts[0])
	if err != nil {
		return nil, LiquidityModel(0), fmt.Errorf("ParseContractCallWithTokenToBTCKeyID > invalid model")
	}

	modelInt := binary.BigEndian.Uint32(modelBytes)
	if _, ok := LiquidityModel_name[int32(modelInt)]; !ok {
		return nil, LiquidityModel(0), fmt.Errorf("ParseContractCallWithTokenToBTCKeyID > invalid model")
	}

	bitcoinPubKey, err = hex.DecodeString(parts[1])
	if err != nil {
		return nil, LiquidityModel(0), fmt.Errorf("ParseContractCallWithTokenToBTCKeyID > invalid bitcoinPubKey")
	}

	return bitcoinPubKey, LiquidityModel(modelInt), nil
}
