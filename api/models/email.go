package models

//Email ...
type Email struct {
	ReceiverEmail string
	ReceiverName string
	Subject string
	FileName string
	SenderEmail string
	Replacer map[string]string
}