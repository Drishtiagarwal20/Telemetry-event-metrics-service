package model

type Events struct {
	Id   int   `json:"id"`
	Data Event `json:"data"`
}

type Event struct {
	ODataContext string       `json:"@odata.context"`
	ODataType    string       `json:"@odata.type"`
	Events       []EventEntry `json:"Events"`
	EventsCount  int          `json:"Events@odata.count"`
	ID           string       `json:"Id"`
	Name         string       `json:"Name"`
}

type EventEntry struct {
	EventID           string          `json:"EventId"`
	EventTimestamp    string          `json:"EventTimestamp"`
	EventType         string          `json:"EventType"`
	MemberID          string          `json:"MemberId"`
	MessageID         string          `json:"MessageId"`
	OriginOfCondition OriginCondition `json:"OriginOfCondition"`
}

type OriginCondition struct {
	ODataID string `json:"@odata.id"`
}
