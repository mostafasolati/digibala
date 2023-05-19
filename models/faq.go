package models

type TagType int

const (
	ACCOUNT int = iota
	PAYMENT
	RETURNPRODUCT
	SENDFEEDBACK
)

type FAQ struct {
	ID          int
	Question    string
	Answer      string
	QuestionTag []TagType
}
