package autocode

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type NoAuthRouter struct {
}

// InitNoAuthRouter 存放不需要鉴权的路由
func (s *NoAuthRouter) InitNoAuthRouter(Router *gin.RouterGroup) {

	NoAuthRouterWithoutRecord := Router.Group("api")
	{
		var articleApi = v1.ApiGroupApp.AutoCodeApiGroup.ArticleApi
		NoAuthArticle := NoAuthRouterWithoutRecord.Group("article")
		{
			NoAuthArticle.GET("getArticleList", articleApi.GetArticleList) // 获取Article列表
			NoAuthArticle.GET("findArticle", articleApi.FindArticle)       // 根据ID获取Article
		}
		var swiperApi = v1.ApiGroupApp.AutoCodeApiGroup.SwiperApi
		NoAuthSwiper := NoAuthRouterWithoutRecord.Group("swiper")
		{
			NoAuthSwiper.GET("getSwiperList", swiperApi.GetSwiperList) // 获取Swiper列表
		}
		var competitionScheApi = v1.ApiGroupApp.AutoCodeApiGroup.CompetitionScheApi
		NoAuthcompetitionSche := NoAuthRouterWithoutRecord.Group("sche")
		{
			NoAuthcompetitionSche.GET("getList", competitionScheApi.GetCompetitionScheDetailList) // 获取Swiper列表
		}
	}

}
