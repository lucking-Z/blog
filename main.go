package main

import (
	"blogs/pkg/db"
	"blogs/pkg/log"
	"blogs/pkg/resp"
	"blogs/pkg/uuid"
	"blogs/src/controller/article"
	"blogs/src/controller/user"
	"blogs/src/model/muser"
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func initSessionMiddleware(ginObj *gin.Engine) gin.HandlerFunc {
	gob.Register(muser.User{})
	store := cookie.NewStore([]byte("test"))
	return sessions.Sessions("blog", store)
}

func initRouter(ginObj *gin.Engine) {
	v1 := ginObj.Group("/", AuthMiddleware())
	{
		v1.GET("/article/get", article.CArticle{}.Get)
		v1.POST("/article/del", article.CArticle{}.Delete)
		v1.POST("/article/update", article.CArticle{}.Update)
		v1.POST("/article/add", article.CArticle{}.Add)
		v1.POST("/user/create", user.Cuser{}.Create)
	}

	ginObj.POST("/user/login", user.Cuser{}.Login)
	ginObj.POST("/user/logout", user.Cuser{}.Logout)
}

func BaseMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuidStr := uuid.GenerateUUID()
		ctx.Set("uuid", uuidStr)
		ctx.Set("ip", ctx.ClientIP())
		ctx.Next()
		//c := context.WithValue(ctx, "uuid", uuidStr)
		//ctx.Set("ctx", c)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//todo 临时用session，后续可以变更jwt
		sess := sessions.Default(ctx)
		if u := sess.Get("login_user"); u == nil {
			resp.ResultRender(ctx, resp.ErrLogin, nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func main() {
	log.Init("./log/")
	ginObj := gin.Default()
	ginObj.Use(initSessionMiddleware(ginObj))
	ginObj.Use(BaseMiddleware())
	//db初始化
	configDb := db.ConfigDB{
		//Username: "root",
		//Password: "kuangzhc",
		//Port:     0,
		//Host:     "",
		//DbName:   "blog",
	}
	db.InitDB(configDb)
	//路由初始化
	initRouter(ginObj)

	ginObj.Run(":80")

}
