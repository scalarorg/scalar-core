package approve

import (
	"fmt"

	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
)

func ProcessDestCallApproved(event chainsTypes.DestCallApproved) error {
	fmt.Printf("Processing DestCallApproved event: %+v\n", event)

	return nil
}
