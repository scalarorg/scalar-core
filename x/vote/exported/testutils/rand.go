package testutils

import (
	rand2 "github.com/scalarorg/scalar-core/testutils/rand"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/utils/test/rand"
	"github.com/scalarorg/scalar-core/x/vote/exported"
)

// RandomPollID generates a random PollID
func RandomPollID() exported.PollID {
	return exported.PollID(rand.PosI64())
}

// RandomPollParticipants generates random PollParticipants
func RandomPollParticipants() exported.PollParticipants {
	return exported.PollParticipants{
		PollID:       RandomPollID(),
		Participants: slices.Expand2(rand2.ValAddr, int(rand.I64Between(1, 20))),
	}
}
