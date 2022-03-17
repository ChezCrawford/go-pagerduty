package webhookv3

import (
	"encoding/json"
	"fmt"

	"github.com/PagerDuty/go-pagerduty"
)

type EventDataType string

const (
	Incident     EventDataType = "incident"
	IncidentNote EventDataType = "incident_note"
)

type IncidentEvent struct {
	ID      string        `json:"id"`
	Type    EventDataType `json:"type"`
	Self    string        `json:"self"`
	HtmlUrl string        `json:"html_url"`
	Number  uint          `json:"number"`
	Title   string        `json:"title"`
}

func (e OutboundEvent) GetIncident() (*IncidentEvent, error) {
	if e.Data.Type != "incident" {
		return nil, fmt.Errorf("incorrect event data type")
	}

	var ie IncidentEvent
	if err := json.Unmarshal(e.Data.RawData, &ie); err != nil {
		return nil, err
	}

	return &ie, nil
}

type IncidentNoteEvent struct {
	ID       string              `json:"id"`
	Type     EventDataType       `json:"type"`
	Incident pagerduty.APIObject `json:"incident"`
	Content  string              `json:"content"`
	Trimmed  bool                `json:"trimmed"`
}

func (e OutboundEvent) GetIncidentNote() (*IncidentNoteEvent, error) {
	if e.Data.Type != "incident_note" {
		return nil, fmt.Errorf("incorrect event data type")
	}

	var ine IncidentNoteEvent
	if err := json.Unmarshal(e.Data.RawData, &ine); err != nil {
		return nil, err
	}

	return &ine, nil
}
