package controllers

import (
	model "OrdsOrder/models"
	"OrdsOrder/util"
	"github.com/astaxie/beego"
	"log"
)

//home 页面的json api
type ApiController struct {
	beego.Controller
}

var (
	allinfo []model.HomeGetAll
	rep     model.Response
)

func (c *ApiController) Get() {
	//分页和url的详细信息
	pageindex, err := c.GetInt("page")
	if err != nil {
		log.Println("没有输入 page")
	}
	pagesize, err := c.GetInt("limit")
	if err != nil {
		log.Println("没有输入 limit")
	}

	db := model.Init()
	defer db.Close()
	count := 0
	db.Table("t_info").Count(&count)

	db.Table("t_info").Select("t_info.t_create,t_info.t_number,t_info.t_name,t_info.t_type,t_info.t_org,t_sperson.t_username, t_sperson.t_phone,t_info.t_status,t_info.t_endtime,t_info.t_extend").
		Order("t_info.t_create desc").Joins("left join t_sperson on t_info.t_srcnameid = t_sperson.t_number").Offset((pageindex - 1) * pagesize).Limit(pagesize).Scan(&allinfo)

	rep.Code = 200
	rep.Count = count
	rep.Msg = "home 页面的全部数据"
	rep.Data = allinfo
	c.Data["json"] = rep
	c.ServeJSON()

}

//处理工单的json api
type ExecApiController struct {
	beego.Controller
}

func (c *ExecApiController) Get() {
	//分页和url的详细信息
	pageindex, err := c.GetInt("page")
	if err != nil {
		log.Println("没有输入 page")
	}
	pagesize, err := c.GetInt("limit")
	if err != nil {
		log.Println("没有输入 limit")
	}

	db := model.Init()
	defer db.Close()
	count := 0
	db.Table("t_info").Where("t_status = ?", "处理中").Count(&count)

	db.Table("t_info").Select("t_info.t_create,t_info.t_number,t_info.t_name,t_info.t_type,t_info.t_org,t_sperson.t_username, t_sperson.t_phone,t_info.t_status,t_info.t_endtime,t_info.t_extend").
		Order("t_info.t_create desc").Joins("left join t_sperson on t_info.t_srcnameid = t_sperson.t_number").Where("t_status = ?", "处理中").Offset((pageindex - 1) * pagesize).Limit(pagesize).Scan(&allinfo)

	rep.Code = 200
	rep.Count = count
	rep.Msg = "执行 页面的全部数据"
	rep.Data = allinfo
	c.Data["json"] = rep
	c.ServeJSON()

}

//挂起工单的json api
type WaitExecApiController struct {
	beego.Controller
}

func (c *WaitExecApiController) Get() {
	//分页和url的详细信息
	pageindex, err := c.GetInt("page")
	if err != nil {
		log.Println("没有输入 page")
	}
	pagesize, err := c.GetInt("limit")
	if err != nil {
		log.Println("没有输入 limit")
	}

	db := model.Init()
	defer db.Close()
	count := 0
	db.Table("t_info").Where("t_status = ?", "已挂起").Count(&count)

	db.Table("t_info").Select("t_info.t_create,t_info.t_number,t_info.t_name,t_info.t_type,t_info.t_org,t_sperson.t_username, t_sperson.t_phone,t_info.t_status,t_info.t_endtime,t_info.t_extend").
		Order("t_info.t_create desc").Joins("left join t_sperson on t_info.t_srcnameid = t_sperson.t_number").Where("t_status = ?", "已挂起").Offset((pageindex - 1) * pagesize).Limit(pagesize).Scan(&allinfo)

	rep.Code = 200
	rep.Count = count
	rep.Msg = "挂起 页面的全部数据"
	rep.Data = allinfo
	c.Data["json"] = rep
	c.ServeJSON()

}

//结束工单的 json api
type EndOrderApiController struct {
	beego.Controller
}

func (c *EndOrderApiController) Get() {
	//分页和url的详细信息
	pageindex, err := c.GetInt("page")
	if err != nil {
		log.Println("没有输入 page")
	}
	pagesize, err := c.GetInt("limit")
	if err != nil {
		log.Println("没有输入 limit")
	}
	db := model.Init()
	defer db.Close()
	count := 0
	db.Table("t_info").Where("t_status = ?", "完成").Count(&count)
	db.Table("t_info").Select("t_info.t_create,t_info.t_number,t_info.t_name,t_info.t_type,t_info.t_org,t_sperson.t_username, t_sperson.t_phone,t_info.t_status,t_info.t_endtime,t_info.t_extend").
		Order("t_info.t_create desc").Joins("left join t_sperson on t_info.t_srcnameid = t_sperson.t_number").Where("t_status = ?", "完成").Offset((pageindex - 1) * pagesize).Limit(pagesize).Scan(&allinfo)

	rep.Code = 200
	rep.Count = count
	rep.Msg = "挂起 页面的全部数据"
	rep.Data = allinfo
	c.Data["json"] = rep
	c.ServeJSON()
}

//超时工单的 json api
type TimeOutApiController struct {
	beego.Controller
}

func (c *TimeOutApiController) Get() {
	//分页和url的详细信息
	pageindex, err := c.GetInt("page")
	if err != nil {
		log.Println("没有输入 page")
	}
	pagesize, err := c.GetInt("limit")
	if err != nil {
		log.Println("没有输入 limit")
	}

	db := model.Init()
	defer db.Close()
	count := 0
	db.Table("t_info").Where("t_status = ?", "完成").Count(&count)
	db.Table("t_info").Select("t_info.t_create,t_info.t_number,t_info.t_name,t_info.t_type,t_info.t_org,t_sperson.t_username, t_sperson.t_phone,t_info.t_status,t_info.t_endtime,t_info.t_extend").
		Order("t_info.t_create desc").Joins("left join t_sperson on t_info.t_srcnameid = t_sperson.t_number").Where("t_status = ?", "完成").Offset((pageindex - 1) * pagesize).Limit(pagesize).Scan(&allinfo)
	var replist []model.HomeGetAll
	for _, v := range allinfo {
		time, err := util.TransTime(v.TCreate, v.TEndtime)
		if err != nil {
			continue
			log.Println("时间搓获取失败", "api:174")
		}
		if time > 7200 {
			replist = append(replist, v)
		}
	}
	rep.Code = 200
	rep.Count = count
	rep.Msg = "超时 页面的全部数据"
	rep.Data = replist
	c.Data["json"] = rep
	c.ServeJSON()
}

//获取 发单人员的api
type SrcPersonController struct {
	beego.Controller
}

func (c *SrcPersonController) Get() {
	//分页和url的详细信息
	pageindex, err := c.GetInt("page")
	if err != nil {
		log.Println("没有输入 page")
	}
	pagesize, err := c.GetInt("limit")
	if err != nil {
		log.Println("没有输入 limit")
	}
	db := model.Init()
	defer db.Close()
	var count = 0
	db.Table("t_sperson").Count(&count)
	var srcpersonlist = []model.TSperson{}
	db.Table("t_sperson").Offset((pageindex - 1) * pagesize).Limit(pagesize).Scan(&srcpersonlist)

	rep.Code = 200
	rep.Count = count
	rep.Msg = "获取所有发单人的信息"
	rep.Data = srcpersonlist
	c.Data["json"] = rep
	c.ServeJSON()
}

//获取 执行人员列表的json
type ExecPersonController struct {
	beego.Controller
}
func (c *ExecPersonController) Get() {
	//分页和url的详细信息
	pageindex, err := c.GetInt("page")
	if err != nil {
		log.Println("没有输入 page")
	}
	pagesize, err := c.GetInt("limit")
	if err != nil {
		log.Println("没有输入 limit")
	}
	db := model.Init()
	defer db.Close()
	var count = 0
	db.Table("t_dperson").Count(&count)
	var tpserson = []model.TDperson{}
	var red = model.Response{}
	db.Table("t_dperson").Offset((pageindex - 1) * pagesize).Limit(pagesize).Scan(&tpserson)
	red.Code = 200
	red.Count = count
	red.Msg = "获取所有收单人的信息"
	red.Data = tpserson
	c.Data["json"] = red
	c.ServeJSON()
}

//logot
type ApiLogOutCotroller struct {
	beego.Controller
}
func (c *ApiLogOutCotroller) Get(){
	c.Ctx.SetCookie("token","")
	c.Ctx.Redirect(302,"/login")
}

