package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}

type CustDetails struct {
	AccID   int32  `json:"accid" gorm:"unique"`
	AccType string `json:"acctype"`
	BCode   string `json:"bcode"`
	Contact int64  `json:"contact"`
	Balance int32  `json:"balance"`
}

type DataToBeSent struct {
	CustDetails []CustDetails 
	Count int64 `json:"count"`
}
