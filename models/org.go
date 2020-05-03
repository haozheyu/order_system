package model

/*

idint(11) NOT NULL
t_numbervarchar(32) NOT NULL
t_namevarchar(32) NULL*/

type TOrg struct {
	TName   string `db:"t_name" json:"t_name"`
}
// t_org 表名绑定
func (t TOrg) TableName() string {
	return "t_org"
}
