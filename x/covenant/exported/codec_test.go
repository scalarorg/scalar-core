package exported_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"

	"github.com/scalarorg/scalar-core/x/covenant/exported"
)

func TestCodec(t *testing.T) {
	mockKeyXOnly := exported.KeyXOnly(bytes.Repeat([]byte{0x01}, 32))
	mockSignature := exported.Signature(bytes.Repeat([]byte{0x02}, 64))
	mockLeafHash := exported.LeafHash(bytes.Repeat([]byte{0x03}, 32))

	keyXOnlyBuf := make([]byte, 32)
	_, err := mockKeyXOnly.MarshalTo(keyXOnlyBuf)
	if err != nil {
		t.Fatalf("failed to marshal key x only: %v", err)
	}

	fmt.Printf("XOnlyKey: %+v\n", hex.EncodeToString(keyXOnlyBuf))

	var v exported.KeyXOnly
	err = v.Unmarshal(keyXOnlyBuf)
	if err != nil {
		t.Fatalf("failed to unmarshal key x only: %v", err)
	}
	fmt.Printf("Unmarshaled KeyXOnly: %+v\n", v)

	_, err = mockKeyXOnly.MarshalTo(keyXOnlyBuf)
	if err != nil {
		t.Fatalf("failed to marshal key x only: %v", err)
	}

	leafHashBuf := make([]byte, 32)
	_, err = mockLeafHash.MarshalTo(leafHashBuf)
	if err != nil {
		t.Fatalf("failed to marshal leaf hash: %v", err)
	}

	tapScriptList := exported.TapScriptSigsMap{
		Inner: map[uint64]*exported.TapScriptSigsList{
			0: {
				List: []*exported.TapScriptSig{
					{
						KeyXOnly:  &mockKeyXOnly,
						Signature: &mockSignature,
						LeafHash:  &mockLeafHash,
					},
				},
			},
		},
	}

	data := make([]byte, 145)

	_, err = tapScriptList.MarshalTo(data)
	if err != nil {
		t.Fatalf("failed to marshal tap script list: %v", err)
	}

	fmt.Println(hex.EncodeToString(data))

	unmarshaled := exported.TapScriptSigsMap{}
	err = unmarshaled.Unmarshal(data)
	if err != nil {
		t.Fatalf("failed to unmarshal tap script list: %v", err)
	}

	for _, tapScriptList := range unmarshaled.Inner {
		for _, tapScriptSig := range tapScriptList.List {
			fmt.Printf("%+v\n", tapScriptSig.KeyXOnly)
			fmt.Printf("%+v\n", tapScriptSig.Signature)
			fmt.Printf("%+v\n", tapScriptSig.LeafHash)
		}
	}

	if !reflect.DeepEqual(tapScriptList, unmarshaled) {
		t.Fatalf("failed to deep equal key x only")
	}

}
