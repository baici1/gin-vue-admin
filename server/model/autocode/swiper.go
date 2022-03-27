// 自动生成模板Swiper
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Swiper 结构体
// 如果含有time.Time 请自行import time包
type Swiper struct {
      global.GVA_MODEL
      SwiperName  string `json:"swiperName" form:"swiperName" gorm:"column:swiper_name;comment:轮播图名称;size:255;"`
      SwiperPicture  string `json:"swiperPicture" form:"swiperPicture" gorm:"column:swiper_picture;comment:轮播图照片;size:255;"`
      IsShow  *bool `json:"isShow" form:"isShow" gorm:"column:is_show;comment:是否展示;"`
      Group  *int `json:"group" form:"group" gorm:"column:group;comment:分组;size:10;"`
      GoToUrl  string `json:"goToUrl" form:"goToUrl" gorm:"column:go_to_url;comment:前往路径;size:255;"`
}


// TableName Swiper 表名
func (Swiper) TableName() string {
  return "swiper"
}

