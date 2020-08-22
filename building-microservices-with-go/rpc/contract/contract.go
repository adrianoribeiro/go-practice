package contract

type CalcRequest struct {
	Number1   int
	Number2   int
	Operation string
}

type CalcResponse struct {
	Message string
}
