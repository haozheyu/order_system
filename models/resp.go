package model

type ResponseMsg struct {
	Status int  `json:"status"`
	Msg interface{} `json:"msg"`
	Data string `json:"data"`
}



