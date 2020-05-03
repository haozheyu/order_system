package model

type TSperson struct {
	Id        int    `db:"id" json:"id"`
	TUsername string `db:"t_username" json:"t_username"`
	TPhone    string `db:"t_phone" json:"t_phone"`
	TAddr     string `db:"t_addr" json:"t_addr"`
	TNumber   string `db:"t_number" json:"t_number"`
	TLock     int    `db:"t_lock" json:"t_lock"`
	TPasswd   string `db:"t_passwd" json:"t_passwd"`
	TOrg      string `db:"t_org" json:"t_org"`
	TType     string `db:"t_type" json:"t_type"`
	IsSrc     string `db:"is_src" json:"is_src"`
}

// t_sperson 表名绑定
func (t TSperson) TableName() string {
	return "t_sperson"
}
