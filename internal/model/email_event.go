package model

type EmailEvent struct {
	Email   string `json:"email,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
	Subject string `json:"subject,omitempty"`
	Body    string `json:"body,omitempty"`
}

func (e *EmailEvent) GetEmail() string {
	return e.Email
}
