package model

/*
idint(11) NOT NULL
t_numbervarchar(32) NULL公告编号
t_titlevarchar(32) NULL公告标题
t_centertext NULL公告内容
t_initiatoridvarchar(32) NULL公告发起人id
t_createvarchar(64) NOT NULL公告创建时间
t_updatevarchar(64) NOT NULL公告更新时间*/

type TAnnouncement struct {
  TTitle string `db:"t_title" json:"t_title"`
  TCenter string `db:"t_center" json:"t_center"`
}
// t_announcement 表名绑定
func (t TAnnouncement) TableName() string {
	return "t_announcement"
}

