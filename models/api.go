package model

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

type HomeGetAll struct {
	TCreate    string `db:"t_create" json:"t_create"`
	TNumber    string `db:"t_number" json:"t_number"`
	TName      string `db:"t_name" json:"t_name"`
	TType      string `db:"t_type" json:"t_type"`
	TOrg       string `db:"t_org" json:"t_org"`
	TUsername string `db:"t_username" json:"t_username"`
	TPhone string `db:"t_phone" json:"t_phone"`
	TStatus    string    `db:"t_status" json:"t_status"`
	TEndtime   string `db:"t_endtime" json:"t_endtime"`
	TExtend string `db:"t_extend" json:"t_extend"`
}

type OrderPage struct {
	TNumber   string `db:"t_number" json:"t_number"`
	TName     string `db:"t_name" json:"t_name"`
	TType     string `db:"t_type" json:"t_type"`
	TOrg      string `db:"t_org" json:"t_org"`
	TUsername string `db:"t_username" json:"t_username"`
	TPhone    string `db:"t_phone" json:"t_phone"`
	TAddr     string `db:"t_addr" json:"t_addr"`

	TDusername string `db:"t_dusername" json:"t_dusername"`
	TDphone    string `db:"t_dphone" json:"t_dphone"`
	TStatus    string `db:"t_status" json:"t_status"`
	TEndtime   string `db:"t_endtime" json:"t_endtime"`
	TExtend    string `db:"t_extend" json:"t_extend"`
}