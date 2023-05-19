package models

type TagName string

const (
	ACCOUNT       TagName = "Account"
	PAYMENT       TagName = "Payment"
	RETURNPRODUCT TagName = "Return Product"
	SENDFEEDBACK  TagName = "Send Feedback"
)

type FAQ struct {
	ID          int
	question    string
	answer      string
	questionTag TagName
}
