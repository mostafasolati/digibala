package models

type TagName string

const (
	ACCOUNT       TagName = "Account"
	PAYMENT       TagName = "Payment"
	RETURNPRODUCT TagName = "Return_Product"
	SENDFEEDBACK  TagName = "Send_Feedback"
)

type FAQ struct {
	ID          int
	Question    string
	Answer      string
	QuestionTag TagName
}
