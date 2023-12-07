package models

type ZincSearchResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Hits     struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Index  string      `json:"_index"`
			Type   string      `json:"_type"`
			ID     string      `json:"_id"`
			Score  float64     `json:"_score"`
			Source EmailDetail `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type EmailDetail struct {
	Timestamp string                 `json:"@timestamp"`
	Body      map[string]interface{} `json:"Body"`
	Date      string                 `json:"Date"`
	From      string                 `json:"From"`
	MessageID string                 `json:"MessageID"`
	Subject   string                 `json:"Subject"`
	To        string                 `json:"To"`
}
