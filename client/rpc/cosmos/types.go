package cosmos

import "fmt"

type EventQuery struct {
	TmEvent   string
	Module    string
	Version   string
	Event     string
	Attribute string
	Operator  string
}

func (e *EventQuery) ToQuery() string {
	return fmt.Sprintf("tm.event='%s' AND %s.%s.%s.%s %s", e.TmEvent, e.Module, e.Version, e.Event, e.Attribute, e.Operator)
}
