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
		var competitionInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.CompetitionInfoApi
		NoAuthcompetitionInfo := NoAuthRouterWithoutRecord.Group("info")
		{
			NoAuthcompetitionInfo.GET("tree_list", competitionInfoApi.GetComSelectList)
		}
		var entryFormApi = v1.ApiGroupApp.AutoCodeApiGroup.EntryFormApi
		NoAuthentryForm := NoAuthRouterWithoutRecord.Group("entry")
		{
			NoAuthentryForm.GET("all", entryFormApi.GetAllEntryFormDetailInfo)
			NoAuthentryForm.GET("first", entryFormApi.GetEntryFormDetailInfo)
			NoAuthentryForm.POST("create", entryFormApi.CreateEntryFormByUser)
			NoAuthentryForm.POST("update", entryFormApi.UpdateEntryFormByUser)
		}
		var teamMemberApi = v1.ApiGroupApp.AutoCodeApiGroup.TeamMemberApi
		NoAuthteamMember := NoAuthRouterWithoutRecord.Group("member")
		{
			NoAuthteamMember.GET("list", teamMemberApi.GetTeamMemberList)
			NoAuthteamMember.POST("create", teamMemberApi.CreateOwnTeamMember)
			NoAuthteamMember.PUT("update", teamMemberApi.UpdateTeamMember)    // 更新TeamMember
			NoAuthteamMember.DELETE("delete", teamMemberApi.DeleteTeamMember) // 批量删除TeamMember
		}
		var companyInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.CompanyInfoApi
		NoAuthcompanyInfo := NoAuthRouterWithoutRecord.Group("company")
		{
			NoAuthcompanyInfo.POST("create", companyInfoApi.CreateCompanyInfo)   // 新建CompanyInfo
			NoAuthcompanyInfo.DELETE("delete", companyInfoApi.DeleteCompanyInfo) // 删除CompanyInfo
			NoAuthcompanyInfo.PUT("update", companyInfoApi.UpdateCompanyInfo)    // 更新CompanyInfo
			NoAuthcompanyInfo.GET("find", teamMemberApi.FindTeamMember)          // 根据ID获取TeamMember
		}
		var teamInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.TeamInfoApi
		NoAuthteamInfo := NoAuthRouterWithoutRecord.Group("team")
		{
			NoAuthteamInfo.POST("createTeamInfo", teamInfoApi.CreateTeamInfo)   // 新建TeamInfo
			NoAuthteamInfo.DELETE("deleteTeamInfo", teamInfoApi.DeleteTeamInfo) // 删除TeamInfo
			NoAuthteamInfo.PUT("updateTeamInfo", teamInfoApi.UpdateTeamInfo)    // 更新TeamInfo
			NoAuthteamInfo.GET("findTeamInfo", teamInfoApi.FindTeamInfo)        // 根据ID获取TeamInfo
		}

	}

}
