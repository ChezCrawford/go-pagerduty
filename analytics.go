package pagerduty

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

type AnalyticsRequest struct {
	AnalyticsFilter *AnalyticsFilter `json:"filters,omitempty"`
	AggregateUnit   string           `json:"aggregate_unit,omitempty"`
	TimeZone        string           `json:"time_zone,omitempty"`
}
type AnalyticsResponse struct {
	Data            []AnalyticsData  `json:"data,omitempty"`
	AnalyticsFilter *AnalyticsFilter `json:"filters,omitempty"`
	AggregateUnit   string           `json:"aggregate_unit,omitempty"`
	TimeZone        string           `json:"time_zone,omitempty"`
}

type AnalyticsFilter struct {
	CreatedAtStart string   `json:"created_at_start,omitempty"`
	CreatedAtEnd   string   `json:"created_at_end,omitempty"`
	Urgency        string   `json:"urgency,omitempty"`
	Major          bool     `json:"major,omitempty"`
	ServiceIds     []string `json:"service_ids,omitempty"`
	TeamIds        []string `json:"team_ids,omitempty"`
	PriorityIds    []string `json:"priority_ids,omitempty"`
	PriorityName   []string `json:"priority_name,omitempty"`
}

type AnalyticsData struct {
	ServiceId                      string `json:"service_id,omitempty"`
	ServiceName                    string `json:"service_name,omitempty"`
	TeamId                         string `json:"team_id,omitempty"`
	TeamName                       string `json:"team_name,omitempty"`
	MeanSecondsToResolve           int    `json:"mean_seconds_to_resolve,omitempty"`
	MeanSecondsToFirstAck          int    `json:"mean_seconds_to_first_ack,omitempty"`
	MeanSecondsToEngage            int    `json:"mean_seconds_to_engage,omitempty"`
	MeanSecondsToMobilize          int    `json:"mean_seconds_to_mobilize,omitempty"`
	MeanEngagedSeconds             int    `json:"mean_engaged_seconds,omitempty"`
	MeanEngagedUserCount           int    `json:"mean_engaged_user_count,omitempty"`
	TotalEscalationCount           int    `json:"total_escalation_count,omitempty"`
	MeanAssignmentCount            int    `json:"mean_assignment_count,omitempty"`
	TotalBusinessHourInterruptions int    `json:"total_business_hour_interruptions,omitempty"`
	TotalSleepHourInterruptions    int    `json:"total_sleep_hour_interruptions,omitempty"`
	TotalOffHourInterruptions      int    `json:"total_off_hour_interruptions,omitempty"`
	TotalSnoozedSeconds            int    `json:"total_snoozed_seconds,omitempty"`
	TotalEngagedSeconds            int    `json:"total_engaged_seconds,omitempty"`
	TotalIncidentCount             int    `json:"total_incident_count,omitempty"`
	UpTimePct                      int    `json:"up_time_pct,omitempty"`
	UserDefinedEffortSeconds       int    `json:"user_defined_effort_seconds,omitempty"`
	RangeStart                     string `json:"range_start,omitempty"`
}

func (c *Client) GetAggregatedIncidentData(analytics AnalyticsRequest) (AnalyticsResponse, error) {
	var analyticsResponse AnalyticsResponse
	headers := make(map[string]string)
	headers["X-EARLY-ACCESS"] = "analytics-v2"
	resp, err := c.post("/analytics/metrics/incidents/all", analytics, &headers)
	if err != nil {
		return analyticsResponse, err
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		bytesBody, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		log.Debug(string(bytesBody))
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(bytesBody))
	}
	err = c.decodeJSON(resp, &analyticsResponse)
	return analyticsResponse, err
}
