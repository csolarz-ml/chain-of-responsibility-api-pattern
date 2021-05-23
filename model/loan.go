package model

type Loan struct {
	Form     LoanForm `json:"form"`
	Message  string   `json:"message"`
	SignedBy string   `json:"signed_by"`
}
