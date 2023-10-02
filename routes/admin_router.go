package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"rookieCode/config"
	"rookieCode/middleware"
)

// 后台管理页面的接口路由
func AdminRouter() http.Handler {
	gin.SetMode(config.Cfg.Server.AppMode)

	r := gin.New()
	//SetTrustedProxies 方法通常用于告诉 web 框架或库哪些 IP 地址可以被信任，以便从请求头中提取正确的客户端 IP 地址。
	r.SetTrustedProxies([]string{"*"})

	// 使用本地文件上传, 需要静态文件服务, 使用七牛云不需要
	if config.Cfg.Upload.OssType == "local" {
		r.Static("/public", "./public")
		r.StaticFS("/dir", http.Dir("./public")) // 将 public 目录内的文件列举展示
	}

	r.Use(middleware.Logger())             // 自定义的 zap 日志中间件
	r.Use(middleware.ErrorRecovery(false)) // 自定义错误处理中间件
	r.Use(middleware.Cors())               // 跨域中间件

	// 基于 cookie 存储 session
	store := cookie.NewStore([]byte(config.Cfg.Session.Salt)) //Cookie 存储提供的密钥

	// session 存储时间跟 JWT 过期时间一致
	store.Options(sessions.Options{MaxAge: int(config.Cfg.JWT.Expire) * 3600})
	r.Use(sessions.Sessions(config.Cfg.Session.Name, store)) // Session 中间件

	// 无需鉴权的接口
	base := r.Group("/api")
	{
		// TODO: 用户注册 和 后台登录 应该记录到 日志
		base.POST("/login", userAuthAPI.Login) // 后台登录
		//base.POST("/report", blogInfoAPI.Report) // 上报信息
	}

	// 需要鉴权的接口
	auth := base.Group("") // "/admin"
	// !注意使用中间件的顺序
	auth.Use(middleware.JWTAuth())      // JWT 鉴权中间件
	auth.Use(middleware.RBAC())         // casbin 权限中间件
	auth.Use(middleware.ListenOnline()) // 监听在线用户
	auth.Use(middleware.OperationLog()) // 记录操作日志
	{
		//auth.GET("/home", blogInfoAPI.GetHomeInfo) // 后台首页信息
		auth.GET("/logout", userAuthAPI.Logout) // 退出登录
		//auth.POST("/upload", uploadAPI.UploadFile) // 文件上传

		// 用户模块
		user := auth.Group("/user")
		{
			user.GET("/list", userAPI.GetList)                           // 用户列表
			user.PUT("", userAPI.Update)                                 // 更新用户信息
			user.PUT("/disable", userAPI.UpdateDisable)                  // 修改用户禁用状态
			user.PUT("/password", userAPI.UpdatePassword)                // 修改普通用户密码
			user.PUT("/current/password", userAPI.UpdateCurrentPassword) // 修改管理员密码
			user.GET("/info", userAPI.GetInfo)                           // 获取当前用户信息
			user.PUT("/current", userAPI.UpdateCurrent)                  // 修改当前用户信息
			user.GET("/online", userAPI.GetOnlineList)                   // 获取在线用户
			user.DELETE("/offline", userAPI.ForceOffline)                // 强制用户下线
		}
		// 角色模块
		role := auth.Group("/role")
		{
			role.GET("/list", roleAPI.GetTreeList) // 角色列表(树形)
			role.POST("", roleAPI.SaveOrUpdate)    // 新增/编辑菜单
			role.DELETE("", roleAPI.Delete)        // 删除角色
			role.GET("/option", roleAPI.GetOption) // 角色选项列表(树形)
		}
	}
	return r
}
