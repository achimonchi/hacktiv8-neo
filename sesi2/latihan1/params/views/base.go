package views

type GeneralSuccessPayload struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

type GeneralErrorPayload struct {
	Status         int         `json:"status"`
	Message        string      `json:"message"`
	Error          string      `json:"error"`
	AdditionalInfo interface{} `json:"additional_info"`
}
