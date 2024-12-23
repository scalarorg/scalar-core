package types

const (
	AttributeKeyChain      = "chain"
	AttributeKeyMessageID  = "messageID"
	AttributeKeyCommandsID = "commandID"
)

type ConfirmationEvent interface {
	isEvent_Event
	MarshalTo([]byte) (int, error)
	Size() int
}
