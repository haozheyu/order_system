package controllers

import (
	model "OrdsOrder/models"
	"OrdsOrder/util"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strconv"
	"time"
)

//主页所有的工单
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}
	// 判断是否可以发布工单  （只能是发单人）
	sp := model.TSperson{}
	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sp)
	switch sp.IsSrc == "1" {
	case true:
		c.Data["is_pushorder"] = "layui-show"
	case false:
		c.Data["is_pushorder"] = "layui-hide"
	}

	dperson := []model.TDperson{}
	db.Table("t_dperson").Where("is_dis = ?",1).Scan(&dperson)
	c.Data["dperson"] = dperson

	c.TplName = "index.html"
}

//执行中的工单
type ExecController struct {
	beego.Controller
}

func (c *ExecController) Get() {
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}

	sp := model.TSperson{}
	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sp)
	switch sp.IsSrc == "1" {
	case true:
		c.Data["is_pushorder"] = "layui-show"
	case false:
		c.Data["is_pushorder"] = "layui-hide"
	}

	dperson := []model.TDperson{}
	db.Table("t_dperson").Where("is_dis = ?",1).Scan(&dperson)
	c.Data["dperson"] = dperson
	c.TplName = "exec.html"
}

//挂起的工单
type WaitExecController struct {
	beego.Controller
}

func (c *WaitExecController) Get() {
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}

	sp := model.TSperson{}
	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sp)
	switch sp.IsSrc == "1" {
	case true:
		c.Data["is_pushorder"] = "layui-show"
	case false:
		c.Data["is_pushorder"] = "layui-hide"
	}

	dperson := []model.TDperson{}
	db.Table("t_dperson").Where("is_dis = ?",1).Scan(&dperson)
	c.Data["dperson"] = dperson
	c.TplName = "waitexec.html"
}

// 结束的工单
type EndOrderController struct {
	beego.Controller
}

func (c *EndOrderController) Get() {
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}

	sp := model.TSperson{}
	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sp)
	switch sp.IsSrc == "1" {
	case true:
		c.Data["is_pushorder"] = "layui-show"
	case false:
		c.Data["is_pushorder"] = "layui-hide"
	}

	dperson := []model.TDperson{}
	db.Table("t_dperson").Where("is_dis = ?",1).Scan(&dperson)
	c.Data["dperson"] = dperson
	c.TplName = "endorder.html"
}

//超时工单
type TimeOutController struct {
	beego.Controller
}

func (c *TimeOutController) Get() {
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}

	sp := model.TSperson{}
	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sp)
	switch sp.IsSrc == "1" {
	case true:
		c.Data["is_pushorder"] = "layui-show"
	case false:
		c.Data["is_pushorder"] = "layui-hide"
	}

	dperson := []model.TDperson{}
	db.Table("t_dperson").Where("is_dis = ?",1).Scan(&dperson)
	c.Data["dperson"] = dperson
	c.TplName = "timeoutorder.html"
}

//发布工单
type PushOrderController struct {
	beego.Controller
}

func (c *PushOrderController) Get() {
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}

	sp := model.TSperson{}
	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sp)
	switch sp.IsSrc == "1" {
	case true:
		c.Data["is_pushorder"] = "layui-show"
	case false:
		c.Data["is_pushorder"] = "layui-hide"
	}

	var atype = []model.TType{}
	var atypelist = []string{}
	db.Table("t_type").Find(&atype)
	for _, v := range atype {
		atypelist = append(atypelist, v.TName)
	}

	var aorg = []model.TOrg{}
	var aorglist = []string{}
	db.Table("t_org").Find(&aorg)
	for _, v := range aorg {
		aorglist = append(aorglist, v.TName)
	}
	c.Data["aorglist"] = aorglist
	c.Data["atypelist"] = atypelist

	dperson := []model.TDperson{}
	db.Table("t_dperson").Where("is_dis = ?",1).Scan(&dperson)
	c.Data["dperson"] = dperson
	c.TplName = "pushorder.html"
}
func (c *PushOrderController) Post() {
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}

	s1 := c.GetString("ordername")
	s2 := c.GetString("ordertype")
	s3 := c.GetString("orderorg")
	s4 := c.GetString("desc")

	info := model.TInfo{}
	typp := model.TType{}
	org := model.TOrg{}
	sperson := model.TSperson{}

	db.Table("t_type").Where("id = ?", s2).Scan(&typp)
	db.Table("t_org").Where("id = ?", s3).Scan(&org)
	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sperson)
	//插入数据准备
	info.TName = s1
	info.TType = typp.TName
	info.TOrg = org.TName
	uuid, err := util.GetUUID()
	if err != nil {
		log.Println("插入工单获取uuid失败", err)
	}
	info.TNumber = uuid
	util.GetTime()
	info.TCreate = util.GetTime()
	info.TSrcnameid = sperson.TNumber
	info.TStatus = "待接手"
	info.TExtend = s4
	db.Table("t_info").Create(&info)

	c.Ctx.Redirect(302, "/")
}

//人员授权 发单人员list
type GetPushPersonController struct {
	beego.Controller
}
func (c *GetPushPersonController) Get() {
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}

	sp := model.TSperson{}
	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sp)
	switch sp.IsSrc == "1" {
	case true:
		c.Data["is_pushorder"] = "layui-show"
	case false:
		c.Data["is_pushorder"] = "layui-hide"
	}

	dperson := []model.TDperson{}
	db.Table("t_dperson").Where("is_dis = ?",1).Scan(&dperson)
	c.Data["dperson"] = dperson
	c.TplName = "accredit.html"
}

//查询 工单详情
type OrderidController struct {
	beego.Controller
}
func (c *OrderidController) Get() {
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}

	sp := model.TSperson{}
	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sp)
	switch sp.IsSrc == "1" {
	case true:
		c.Data["is_pushorder"] = "layui-show"
	case false:
		c.Data["is_pushorder"] = "layui-hide"
	}

	orderid := c.GetString("uid") //工单的id
	lenorderid := len(orderid)
	if lenorderid != 32 {
		log.Println("uid 格式有误")
	}

	tinfo := model.TInfo{}
	sperson := model.TSperson{}
	dperson := model.TDperson{}
	db.Table("t_info").Where("t_number = ?", orderid).Scan(&tinfo)
	db.Table("t_sperson").Where("t_number = ?", tinfo.TSrcnameid).Scan(&sperson)
	db.Table("t_dperson").Where("t_number = ?", tinfo.TDisnameid).Scan(&dperson)
	log.Println(dperson)
	c.Data["order_nmber"] = tinfo.TNumber
	c.Data["order_name"] = tinfo.TName
	c.Data["order_type"] = tinfo.TType
	c.Data["order_org"] = tinfo.TOrg
	c.Data["order_srcname"] = sperson.TUsername
	c.Data["order_srcphone"] = sperson.TPhone
	c.Data["order_srcaddr"] = sperson.TAddr

	c.Data["order_disname"] = dperson.TDusername
	c.Data["order_disphone"] = dperson.TDphone
	c.Data["order_status"] = tinfo.TStatus
	c.Data["order_start"] = tinfo.TCreate
	c.Data["order_end"] = tinfo.TEndtime
	c.Data["order_extend"] = tinfo.TExtend
	c.Data["uid"] = tinfo.TNumber

	c.Data["dperson"] = dperson
	c.TplName = "queryorder.html"
}

//编辑 发单人信息
type EditStartP struct {
	beego.Controller
}

func (c *EditStartP) Get() {
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}

	sp := model.TSperson{}
	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sp)
	switch sp.IsSrc == "1" {
	case true:
		c.Data["is_pushorder"] = "layui-show"
	case false:
		c.Data["is_pushorder"] = "layui-hide"
	}

	orderid := c.GetString("uid")
	lenorderid := len(orderid)
	if lenorderid != 32 {
		log.Println("uid 格式有误")
	}
	sp1 := model.TSperson{}
	db.Table("t_sperson").Where("t_number = ? ", orderid).Scan(&sp1)
	c.Data["titile"] = "编辑人员信息"
	c.Data["t_username"] = sp1.TUsername
	c.Data["spasswd"] = sp1.TPasswd
	c.Data["sphone"] = sp1.TPhone
	c.Data["saddr"] = sp1.TAddr
	c.Data["is_look"] = sp1.TLock
	c.Data["sorg"] = sp1.TOrg
	c.Data["stype"] = sp1.TType

	var orglist = []model.TOrg{}
	var olist []string
	db.Table("t_org").Find(&orglist)
	for _, v := range orglist {
		olist = append(olist, v.TName)
	}
	c.Data["olist"] = olist

	var typelist = []model.TSptype{}
	var tlist []string
	db.Table("t_sptype").Find(&typelist)
	for _, v := range typelist {
		tlist = append(tlist, v.TName)
	}
	c.Data["tlist"] = tlist

	dperson := []model.TDperson{}
	db.Table("t_dperson").Where("is_dis = ?",1).Scan(&dperson)
	c.Data["dperson"] = dperson
	c.TplName = "editstaff.html"
}
func (c *EditStartP) Post() {
	db := model.Init()
	defer db.Close()

	uid := c.GetString("uid")
	name := c.GetString("username")
	password := c.GetString("passwrod")
	addr := c.GetString("addr")
	is_start := c.GetString("is_start")
	orgid := c.GetString("orgid")
	typeid := c.GetString("typeid")
	lock, err := strconv.Atoi(is_start)
	if err != nil {
		log.Println(err, "default 502")
	}
	sperson := model.TSperson{}
	sperson.TUsername = name
	sperson.TPasswd = password
	sperson.TAddr = addr
	sperson.TLock = lock
	org := model.TOrg{}
	db.Table("t_org").Where("id = ?", orgid).Scan(&org)
	sperson.TOrg = org.TName
	tpp := model.SpType{}
	db.Table("t_sptype").Where("id = ?", typeid).Scan(&tpp)
	sperson.TType = tpp.TName
	db.Table("t_sperson").Where("t_number = ?", uid).Update(&sperson)
	c.Ctx.Redirect(302, "/startp")
}

//执行 工单人员
type ExecPController struct {
	beego.Controller
}

func (c *ExecPController) Get() {
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}

	sp := model.TSperson{}
	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sp)
	switch sp.IsSrc == "1" {
	case true:
		c.Data["is_pushorder"] = "layui-show"
	case false:
		c.Data["is_pushorder"] = "layui-hide"
	}

	dperson := []model.TDperson{}
	db.Table("t_dperson").Where("is_dis = ?",1).Scan(&dperson)
	c.Data["dperson"] = dperson
	c.TplName = "execp.html"
}

//编辑 收单人信息
type ExecEditController struct {
	beego.Controller
}
func (c *ExecEditController) Get() { //18：52 10:30
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}

	sp := model.TSperson{}
	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sp)
	switch sp.IsSrc == "1" {
	case true:
		c.Data["is_pushorder"] = "layui-show"
	case false:
		c.Data["is_pushorder"] = "layui-hide"
	}

	uid := c.GetString("uid")
	var dp = model.TDperson{}
	db.Table("t_dperson").Where("t_number = ? ", uid).First(&dp)

	c.Data["titile"] = "收单人信息编辑"
	c.Data["t_username"] = dp.TDusername
	c.Data["spasswd"] = dp.TPasswd
	c.Data["sphone"] = dp.TDphone
	c.Data["t_function"] = dp.TFunction
	c.Data["is_look"] = dp.TLock
	c.Data["sorg"] = dp.TOrg
	c.Data["stype"] = dp.TType

	var orglist = []model.TOrg{}
	var olist []string
	db.Table("t_org").Find(&orglist)
	for _, v := range orglist {
		olist = append(olist, v.TName)
	}
	c.Data["olist"] = olist

	var typelist = []model.TSptype{}
	var tlist []string
	db.Table("t_sptype").Find(&typelist)
	for _, v := range typelist {
		tlist = append(tlist, v.TName)
	}
	c.Data["tlist"] = tlist

	dperson := []model.TDperson{}
	db.Table("t_dperson").Where("is_dis = ?",1).Scan(&dperson)
	c.Data["dperson"] = dperson
	c.TplName = "editexecp.html"

}
func (c *ExecEditController) Post() { //19:01
	db := model.Init()
	defer db.Close()
	uid := c.GetString("uid")
	name := c.GetString("username")
	password := c.GetString("passwrod")
	addr := c.GetString("function")
	is_start := c.GetString("is_start")
	orgid := c.GetString("orgid")
	typeid := c.GetString("typeid")
	is_dis := c.GetString("is_dis")
	lock, err := strconv.Atoi(is_start)
	if err != nil {
		log.Println(err, "default 502")
	}
	dperson := model.TDperson{}
	dperson.TDusername = name
	dperson.TPasswd = password
	dperson.TFunction = addr
	dperson.TLock = lock
	dperson.IsDis = is_dis
	org := model.TOrg{}
	db.Table("t_org").Where("id = ?", orgid).Scan(&org)
	dperson.TOrg = org.TName
	tpp := model.SpType{}
	db.Table("t_sptype").Where("id = ?", typeid).Scan(&tpp)
	dperson.TType = tpp.TName

	db.Table("t_dperson").Where("t_number = ?", uid).Update(&dperson)

	c.Ctx.Redirect(302, "/execp")
}

// add 人员添加
type AddPController struct {
	beego.Controller
}

func (c *AddPController) Get() {
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}

	sp := model.TSperson{}
	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sp)
	switch sp.IsSrc == "1" {
	case true:
		c.Data["is_pushorder"] = "layui-show"
	case false:
		c.Data["is_pushorder"] = "layui-hide"
	}

	org := []model.TOrg{}
	var orgarry  []string
	db.Table("t_org").Find(&org)
	for _, v := range org {
		orgarry = append(orgarry, v.TName)
	}
	c.Data["orglist"] = orgarry

	tpp := []model.SpType{}
	var typearry []string
	db.Table("t_sptype").Find(&tpp)
	for _, v := range tpp {
		typearry = append(typearry, v.TName)
	}
	c.Data["typelist"] = typearry

	dperson := []model.TDperson{}
	db.Table("t_dperson").Where("is_dis = ?",1).Scan(&dperson)
	c.Data["dperson"] = dperson
	c.TplName = "addperson.html"
}
func (c *AddPController) Post() {
	db := model.Init()
	defer db.Close()
	switch c.GetString("fname") != "" {
	case true:
		fname := c.GetString("fname")
		fpassword := c.GetString("fpassword")
		fphone := c.GetString("fphone")
		faddr := c.GetString("faddr")
		forg := c.GetString("forg")
		ftype := c.GetString("ftype")

		sperson := model.TSperson{}
		sperson.TUsername = fname
		sperson.TPasswd = fpassword
		sperson.TPhone = fphone
		uuid, _ := util.GetUUID()
		sperson.TNumber = uuid
		sperson.TLock = 0
		org := model.TOrg{}
		db.Table("t_org").Where("id = ?",forg).Scan(&org)
		sperson.TOrg = org.TName
		typp := model.SpType{}
		db.Table("t_sptype").Where("id = ?",ftype).Scan(&typp)
		sperson.TType = typp.TName
		sperson.TAddr = faddr
		sperson.IsSrc = "1"
		db.Table("t_sperson").Save(&sperson)
	case false:
		sname:=c.GetString("sname")
		spassword:=c.GetString("spassword")
		sphone:=c.GetString("sphone")
		sfunction:=c.GetString("sfunction")
		sorg:=c.GetString("sorg")
		stype:=c.GetString("stype")

		sperson := model.TDperson{}
		sperson.TDusername = sname
		sperson.TPasswd = spassword
		sperson.TDphone = sphone
		uuid, _ := util.GetUUID()
		sperson.TNumber = uuid
		sperson.TLock = 0
		org := model.TOrg{}
		db.Table("t_org").Where("id = ?",sorg).Scan(&org)
		sperson.TOrg = org.TName
		typp := model.SpType{}
		db.Table("t_sptype").Where("id = ?",stype).Scan(&typp)
		sperson.TType = typp.TName
		sperson.TFunction = sfunction
		sperson.IsDis = "0"
		db.Table("t_dperson").Save(&sperson)
	}

	time.Sleep(time.Second * 1)
	c.Ctx.Redirect(302,"/addp")

} //19.26

//登录页面
type LoginController struct {
	beego.Controller
}
func (c *LoginController) Get() {
	db := model.Init()
	defer db.Close()

	ann := []model.TAnnouncement{}
	db.Table("t_announcement").Order("id desc").Limit(3).Scan(&ann)
	c.Data["ann"]= ann
	c.TplName = "login.html"
}
func (c *LoginController) Errot() {
	c.TplName = "usererr.html"
}
func (c *LoginController) AddAnn(){ //添加公告
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}
	// 判断是否可以发布工单  （只能是发单人）
	sp := model.TSperson{}
	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sp)
	switch sp.IsSrc == "1" {
	case true:
		c.Data["is_pushorder"] = "layui-show"
	case false:
		c.Data["is_pushorder"] = "layui-hide"
	}
	dperson := []model.TDperson{}
	db.Table("t_dperson").Where("is_dis = ?",1).Scan(&dperson)
	c.Data["dperson"] = dperson
	c.TplName = "addann.html"
}
func (c *LoginController) PostAddAnn(){
	title := c.GetString("title")
	desc := c.GetString("desc")
	ann := model.TAnnouncement{}
	ann.TTitle = title
	ann.TCenter = desc
	db := model.Init()
	defer db.Close()
	db.Table("t_announcement").Save(&ann)
	c.Ctx.Redirect(302,"/")
}
func (c *LoginController) OneSelf() {
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}

	dp := model.TDperson{}
	sp := model.TSperson{}
	db.Table("t_dperson").Where("t_dphone = ?", phone).First(&dp)
	db.Table("t_sperson").Where("t_phone = ?", phone).First(&sp)
	switch dp.TDphone == phone {
	case false:
		c.Data["is_src"] = "发布工单人"
		c.Data["src"] = true
		c.Data["name"] = sp.TUsername
		c.Data["passwd"] = sp.TPasswd
		c.Data["phon"] = sp.TPhone
		c.Data["addr"] = sp.TAddr
		c.Data["is_true"] = sp.TLock
		c.Data["org"] = sp.TOrg
		c.Data["type"] = sp.TType
	case true:
		c.Data["src"] = false
		c.Data["is_src"] = "工单处理人"
		c.Data["name"] = dp.TDusername
		c.Data["passwd"] = dp.TPasswd
		c.Data["phon"] = dp.TDphone
		c.Data["function"] = dp.TFunction
		c.Data["is_true"] = dp.TLock
		c.Data["org"] = dp.TOrg
		c.Data["type"] = dp.TType
	}
	c.TplName = "oneself.html"
}
func (c *LoginController) Post() {
	phone := c.GetString("phone")
	passwd := c.GetString("password")
	is_true := c.GetString("is_true")

	db := model.Init()
	defer db.Close()
	rep := model.ResponseMsg{}
	sp := model.TSperson{}
	dp := model.TDperson{}

	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sp)
	db.Table("t_dperson").Where("t_dphone = ?", phone).Scan(&dp)

	switch is_true == "on" {
	case true: //发单人
		switch sp.TPasswd == passwd {
		case true: //密码正确
			switch sp.TType == "管理员" {
			case true: //是否为管理员
				token, err := util.CreateToken([]byte(util.MySigningKey), "zmenh", phone, true)
				if err != nil {
					log.Println("token 获取失败", err)
				}
				c.Ctx.SetCookie("token", token)
				c.Ctx.Redirect(302, "/")
			case false: //不是管理员
				token, err := util.CreateToken([]byte(util.MySigningKey), "zmenh", phone, false)
				if err != nil {
					log.Println("token 获取失败", err)
				}
				c.Ctx.SetCookie("token", token)
				c.Ctx.Redirect(302, "/")
			}
		case false: //密码错误
			rep.Status = 601
			rep.Data = "发单人用户或密码不存在"
			c.Data["msg"] = rep
			c.TplName = "usererr.html"
		}
	case false: //收单人
		switch dp.TPasswd == passwd {
		case true: //密码正确
			switch dp.TType == "管理员" {
			case true: //是否为管理员
				token, err := util.CreateToken([]byte(util.MySigningKey), "zmenh", phone, true)
				if err != nil {
					log.Println("token 获取失败", err)
				}
				c.Ctx.SetCookie("token", token)
				c.Ctx.Redirect(302, "/")
			case false: //不是管理员
				token, err := util.CreateToken([]byte(util.MySigningKey), "zmenh", phone, false)
				if err != nil {
					log.Println("token 获取失败", err)
				}
				c.Ctx.SetCookie("token", token)
				c.Ctx.Redirect(302, "/")
			}
		case false: //密码错误
			rep.Status = 701
			rep.Data = "收单人用户或密码不存在"
			c.Data["msg"] = rep
			c.TplName = "usererr.html"
		}
	}
}

//发单人编辑工单详情/收单人接手
func (c *LoginController) EditOrder() {
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}

	sp := model.TSperson{}
	db.Table("t_sperson").Where("t_phone = ?", phone).Scan(&sp)
	switch sp.IsSrc == "1" {
	case true:
		c.Data["is_pushorder"] = "layui-show"
	case false:
		c.Data["is_pushorder"] = "layui-hide"
	}

	orderid := c.GetString("uid")
	lenorderid := len(orderid)
	if lenorderid != 32 {
		log.Println("uid 格式有误")
	}

	//处理编辑工单的数据持久化
	tinfo := model.TInfo{}
	sperson := model.TSperson{}
	dperson := model.TDperson{}
	db.Table("t_info").Where("t_number = ?", orderid).First(&tinfo)
	db.Table("t_sperson").Where("t_number = ?", tinfo.TSrcnameid).First(&sperson)
	db.Table("t_dperson").Where("t_number = ?", tinfo.TDisnameid).First(&dperson)

	//后端传给前端的展示数据
	c.Data["order_nmber"] = tinfo.TNumber
	c.Data["order_name"] = tinfo.TName
	c.Data["order_type"] = tinfo.TType
	c.Data["order_org"] = tinfo.TOrg
	c.Data["order_srcname"] = sperson.TUsername
	c.Data["order_srcphone"] = sperson.TPhone
	c.Data["order_srcaddr"] = sperson.TAddr
	c.Data["order_status"] = tinfo.TStatus
	c.Data["order_start"] = tinfo.TCreate
	c.Data["order_extend"] = tinfo.TExtend

	// 收单人值班查询
	db.Table("t_dperson").Where("is_dis = ?",1).Scan(&dperson)
	c.Data["dperson"] = dperson
	c.TplName = "editoder.html"
}
func (c *LoginController) EditOrderPost() {
	db := model.Init()
	defer db.Close()
	token := c.Ctx.GetCookie("token")
	pmap, err := util.ParseToken(token, []byte(util.MySigningKey))
	if err != nil {
		log.Println("没有获取到token 信息")
	}
	phone := pmap.(jwt.MapClaims)["phone"]
	is_admin := pmap.(jwt.MapClaims)["admin"]
	c.Data["phone"] = phone
	if is_admin != true {
		c.Data["is_admin"] = "layui-hide"
	} else {
		c.Data["is_admin"] = "layui-show"
	}

	desc := c.GetString("desc")
	orderstatus := c.GetString("orderstatus")
	uid := c.GetString("uid")
	dperson := model.TDperson{}
	t_info := model.TInfo{}
	db.Table("t_dperson").Where("t_dphone = ?", phone).Scan(&dperson)
	db.Table("t_info").Where("t_number = ?", uid).Scan(&t_info)

	if orderstatus == "1" {
		temp := "待接手"
		t_info.TStatus = temp
		t_info.TDisnameid = dperson.TNumber
		t_info.TExtend = desc
		t_info.TEndtime = ""
		db.Table("t_info").Where("t_number = ?", uid).Update(&t_info)
	} else if orderstatus == "2" {
		temp := "处理中"
		t_info.TStatus = temp
		t_info.TDisnameid = dperson.TNumber
		t_info.TExtend = desc
		t_info.TEndtime = ""
		db.Table("t_info").Where("t_number = ?", uid).Update(&t_info)
	} else if orderstatus == "3" {
		temp := "已挂起"
		t_info.TStatus = temp
		t_info.TDisnameid = dperson.TNumber
		t_info.TExtend = desc
		t_info.TEndtime = ""
		db.Table("t_info").Where("t_number = ?", uid).Update(&t_info)
	} else if orderstatus == "4" {
		temp := "完成"
		t_info.TStatus = temp
		t_info.TDisnameid = dperson.TNumber
		t_info.TExtend = desc
		time := util.GetTime()
		t_info.TEndtime = time
		db.Table("t_info").Where("t_number = ?", uid).Update(&t_info)
	} else {
		c.Data["msg"] = "编辑工单上传失败，请联系管理员"
		c.TplName = "usererr.html"
	}
	c.Ctx.Redirect(302, "/")
}
