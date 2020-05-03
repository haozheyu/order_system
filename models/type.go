package model

/*
idint(6) NOT NULL
t_typevarchar(32) NULL工单类型
t_numbervarchar(32) NOT NULL工单类型编号
*/

type TType struct {
	TName   string `db:"t_name" json:"t_name"`
}
// t_type 表名绑定
func (t TType) TableName() string {
	return "t_type"
}

type TSptype struct {
	TName   string `db:"t_name" json:"t_name"`
}
// t_type 表名绑定
func (t TSptype) TableName() string {
	return "t_sptype"
}

