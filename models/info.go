package model



type TInfo struct {
	TCreate    string `db:"t_create" json:"t_create"`
	TNumber    string `db:"t_number" json:"t_number"`
	TName      string `db:"t_name" json:"t_name"`
	TType      string `db:"t_type" json:"t_type"`
	TOrg       string `db:"t_org" json:"t_org"`
	TDisnameid string `db:"t_disnameid" json:"t_disnameid"`
	TSrcnameid string `db:"t_srcnameid" json:"t_srcnameid"`
	TStatus    string    `db:"t_status" json:"t_status"`
	TEndtime   string `db:"t_endtime" json:"t_endtime"`
	TExtend    string `db:"t_extend" json:"t_extend"`
}

// t_info 表名绑定
func (t TInfo) TableName() string {
	return "t_info"
}
