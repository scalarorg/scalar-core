package testutils

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/scalarorg/scalar-core/testutils/rand"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/x/snapshot/exported"
)

// Snapshot returns random snapshot based on the given parameters
func Snapshot(participantCount uint64, threshold utils.Threshold) exported.Snapshot {
	participantsWeight := sdk.ZeroUint()
	participants := slices.Expand(func(int) exported.Participant {
		weight := sdk.NewUint(uint64(rand.I64Between(1, 100)))
		participantsWeight = participantsWeight.Add(weight)

		return exported.NewParticipant(rand.ValAddr(), weight)
	},
		int(participantCount),
	)

	bondedWeight := sdk.NewUint(uint64(rand.I64Between(
		participantsWeight.BigInt().Int64(),
		participantsWeight.MulUint64(uint64(threshold.Denominator)).QuoUint64(uint64(threshold.Numerator)).BigInt().Int64()+1),
	))

	return exported.NewSnapshot(time.Now(), rand.I64Between(1, 1000), participants, bondedWeight)
}
