package events

type Event struct {
	EventKind string `json:"event_kind"`
	MsgBody   string `json:"msg_body"`
}
