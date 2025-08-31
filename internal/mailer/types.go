package mailer

type Message struct {
	To      string
	Subject string
	Body    string
}

type EmailAddress struct {
	Address string `json:"address"`
}

type GraphRecipient struct {
	EmailAddress EmailAddress `json:"emailAddress"`
}

type GraphMessage struct {
	Message struct {
		Subject string `json:"subject"`
		Body    struct {
			ContentType string `json:"contentType"`
			Content     string `json:"content"`
		} `json:"body"`
		ToRecipients []GraphRecipient `json:"toRecipients"`
	} `json:"message"`
}
