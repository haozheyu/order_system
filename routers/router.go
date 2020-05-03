package routers

import (
	"OrdsOrder/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

)


//路由过滤器设置获取token 解密
var fileterFunc = func(ctx *context.Context){
	url := ctx.Request.URL
	if url.Path != "/login" && url.Path != "/error" {
		token := ctx.GetCookie("token")
		if len(token) < 40 {
			ctx.Redirect(302,"/login")
		}
	}
}


func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, fileterFunc)
    beego.Router("/", &controllers.MainController{})  //主页
	beego.Router("/exec", &controllers.ExecController{}) //执行工单
	beego.Router("/suspensive", &controllers.WaitExecController{}) //挂起工单
	beego.Router("/pushorder", &controllers.PushOrderController{}) //上传工单
	beego.Router("/endorder", &controllers.EndOrderController{}) //结束工单
	beego.Router("/timeout", &controllers.TimeOutController{}) //超时工单
	beego.Router("/startp", &controllers.GetPushPersonController{}) //发单人列表
	beego.Router("/execp", &controllers.ExecPController{}) //执行工单的列表
	beego.Router("/editstartp", &controllers.EditStartP{}) //管理员编辑发单人的
	beego.Router("/addp", &controllers.AddPController{}) //添加人员
	beego.Router("/epedit", &controllers.ExecEditController{}) // 管理员编辑收单人的
	beego.Router("/login", &controllers.LoginController{}) //登录
	beego.Router("/addann", &controllers.LoginController{},"get:AddAnn;post:PostAddAnn") //登录
	beego.Router("/staff", &controllers.LoginController{},"get:OneSelf") //个人主页
	beego.Router("/error", &controllers.LoginController{},"get:Errot") //错误页面
	beego.Router("/editorder", &controllers.LoginController{},"get:EditOrder;post:EditOrderPost") //工单处理页面（发单人编辑/收单人接收）




    //api
    beego.Router("/v1/api/allinfo", &controllers.ApiController{}) //home 工单列表jsoN
	beego.Router("/v1/api/exec/", &controllers.ExecApiController{}) //执行工单列表json
	beego.Router("/v1/api/suspensive/", &controllers.WaitExecApiController{}) //挂起工单列表json
	beego.Router("/v1/api/endorder/", &controllers.EndOrderApiController{}) //结束工单列表json
	beego.Router("/v1/api/timeout/", &controllers.TimeOutApiController{}) //超时工单列表json
	beego.Router("/v1/api/Order", &controllers.OrderidController{}) //查询工单详情页
	beego.Router("/v1/api/srcperson", &controllers.SrcPersonController{}) //发单人列表 json
	beego.Router("/v1/api/execplist", &controllers.ExecPersonController{}) //收单人列表json
	beego.Router("/v1/api/logout", &controllers.ApiLogOutCotroller{}) //等出

}
