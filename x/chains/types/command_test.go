package types_test

import (
	fmt "fmt"
	"strconv"
	"testing"
)

func TestDecodeRedeemTokenParams(t *testing.T) {

}
func TestConverter(t *testing.T) {
	phase := uint8(1)
	phaseStr := strconv.FormatUint(uint64(phase), 10)
	fmt.Println(phaseStr)
}
