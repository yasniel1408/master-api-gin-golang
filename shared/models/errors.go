package shared_errors

type ErrorStruct struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
