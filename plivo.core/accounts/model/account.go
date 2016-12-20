package model

type Account struct {
	Id       int
	AuthId   string `orm:"column(auth_id)"`
	UserName string `orm:"column(username)"`
}

func (acc *Account) TableName() string {
	return "account"
}
