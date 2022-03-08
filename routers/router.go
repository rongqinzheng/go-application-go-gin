package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rongqinzheng/go-application-go-gin/middleware/jwt"
	"github.com/rongqinzheng/go-application-go-gin/pkg/setting"
	"github.com/rongqinzheng/go-application-go-gin/routers/api"
	"github.com/rongqinzheng/go-application-go-gin/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("test/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取单个文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//新增文章
		apiv1.POST("/articles", v1.AddArticle)
		//编辑文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return r
}
