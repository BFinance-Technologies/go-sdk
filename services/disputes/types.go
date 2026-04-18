package disputes

import "time"

type CreateDisputeRequest struct {
	CardID        string   `json:"cardId"`
	TransactionID string   `json:"transactionId"`
	TextEvidence  string   `json:"textEvidence"`
	Files         []string `json:"files,omitempty"`
}

type CreateDisputeData struct {
	ID           string    `json:"id"`
	Status       string    `json:"status"`
	TextEvidence string    `json:"textEvidence"`
	Files        []string  `json:"files"`
	CreatedAt    time.Time `json:"createdAt"`
}

type GetDisputeStatusData struct {
	ID       string  `json:"id"`
	Status   string  `json:"status"`
	Response *string `json:"response,omitempty"`
}

type CancelDisputeData struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
