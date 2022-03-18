package webhookv3

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncidentEvent_GetIncidentEvent(t *testing.T) {
	var wp WebhookPayload

	data := `{ "event": { "id": "5ac64822-4adc-4fda-ade0-410becf0de4f", "event_type": "incident.priority_updated", "resource_type": "incident", "occurred_at": "2020-10-02T18:45:22.169Z", "agent": { "html_url": "https://acme.pagerduty.com/users/PLH1HKV", "id": "PLH1HKV", "self": "https://api.pagerduty.com/users/PLH1HKV", "summary": "Tenex Engineer", "type": "user_reference" }, "client": { "name": "PagerDuty" }, "data": { "id": "PGR0VU2", "type": "incident", "self": "https://api.pagerduty.com/incidents/PGR0VU2", "html_url": "https://acme.pagerduty.com/incidents/PGR0VU2", "number": 2, "status": "triggered", "title": "A little bump in the road", "service": { "html_url": "https://acme.pagerduty.com/services/PF9KMXH", "id": "PF9KMXH", "self": "https://api.pagerduty.com/services/PF9KMXH", "summary": "API Service", "type": "service_reference" }, "assignees": [ { "html_url": "https://acme.pagerduty.com/users/PTUXL6G", "id": "PTUXL6G", "self": "https://api.pagerduty.com/users/PTUXL6G", "summary": "User 123", "type": "user_reference" } ], "escalation_policy": { "html_url": "https://acme.pagerduty.com/escalation_policies/PUS0KTE", "id": "PUS0KTE", "self": "https://api.pagerduty.com/escalation_policies/PUS0KTE", "summary": "Default", "type": "escalation_policy_reference" }, "teams": [ { "html_url": "https://acme.pagerduty.com/teams/PFCVPS0", "id": "PFCVPS0", "self": "https://api.pagerduty.com/teams/PFCVPS0", "summary": "Engineering", "type": "team_reference" } ], "priority": { "html_url": "https://acme.pagerduty.com/account/incident_priorities", "id": "PSO75BM", "self": "https://api.pagerduty.com/priorities/PSO75BM", "summary": "P1", "type": "priority_reference" }, "urgency": "high", "conference_bridge": { "conference_number": "+1 1234123412,,987654321#", "conference_url": "https://example.com" }, "resolve_reason": null } } }`

	err := json.Unmarshal([]byte(data), &wp)
	assert.NoError(t, err)

	oe := wp.Event
	incident, err := oe.GetIncident()
	assert.NoError(t, err)
	assert.Equal(t, "PGR0VU2", incident.ID)
	assert.Equal(t, Incident, incident.Type)
	assert.Equal(t, "A little bump in the road", incident.Title)
	assert.Equal(t, "https://api.pagerduty.com/incidents/PGR0VU2", incident.Self)
	assert.Equal(t, "https://acme.pagerduty.com/incidents/PGR0VU2", incident.HtmlUrl)

}

func TestIncidentNote_GetIncidentNote(t *testing.T) {
	var wp WebhookPayload

	data := `{ "event": { "id": "01BRB6ZP4M6T8ZG4X6BP63ZB9O", "event_type": "incident.annotated", "resource_type": "incident", "occurred_at": "2021-03-02T13:35:11.682Z", "agent": null, "client": null, "data": { "incident": { "html_url": "https://acme.pagerduty.com/incidents/PGR0VU2", "id": "PGR0VU2", "self": "https://api.pagerduty.com/incidents/PGR0VU2", "summary": "A little bump in the road", "type": "incident_reference" }, "id": "P2LA89X", "content": "I sure am glad we are using PagerDuty!", "trimmed": false, "type": "incident_note" } } }`

	err := json.Unmarshal([]byte(data), &wp)
	assert.NoError(t, err)

	oe := wp.Event
	note, err := oe.GetIncidentNote()
	assert.NoError(t, err)
	assert.Equal(t, "P2LA89X", note.ID)
	assert.Equal(t, "PGR0VU2", note.Incident.ID)
	assert.Equal(t, "incident_reference", note.Incident.Type)
	assert.Equal(t, "I sure am glad we are using PagerDuty!", note.Content)
	assert.Equal(t, false, note.Trimmed)

	_, err = oe.GetIncident()
	assert.Error(t, err)
}
