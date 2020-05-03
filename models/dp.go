package model

/*
idint(11) NOT NULL
t_namevarchar(32) NULL接收人姓名
t_phoneint(13) NULL接收人联系方式
t_numbervarchar(32) NOT NULL接收人编号
t_functionvarchar(32) NULL接收人的职位*/

type TDperson struct {
	Id         int    `db:"id" json:"id"`
	TDusername string `db:"t_dusername" json:"t_dusername"`
	TDphone    string    `db:"t_dphone" json:"t_dphone"`
	TNumber    string `db:"t_number" json:"t_number"`
	TFunction  string `db:"t_function" json:"t_function"`
	TLock      int    `db:"t_lock" json:"t_lock"`
	TPasswd    string `db:"t_passwd" json:"t_passwd"`
	TOrg       string `db:"t_org" json:"t_org"`
	TType      string `db:"t_type" json:"t_type"`
	IsDis 	   string `db:"is_dis" json:"is_dis"`
}

// t_dperson 表名绑定
func (t TDperson) TableName() string {
	return "t_dperson"
}
