package autocode

import (
	"encoding/json"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/rabbitmq/hello_world"
	"go.uber.org/zap"
	"os"
)

type StudentRecruitService struct {
}

// CreateStudentRecruit 创建StudentRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentRecruitService *StudentRecruitService) CreateStudentRecruit(studentRecruit autocode.StudentRecruit) (err error) {
	err = global.GVA_DB.Create(&studentRecruit).Error
	return err
}

// DeleteStudentRecruit 删除StudentRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentRecruitService *StudentRecruitService) DeleteStudentRecruit(studentRecruit autocode.StudentRecruit) (err error) {
	err = global.GVA_DB.Delete(&studentRecruit).Error
	return err
}

// DeleteStudentRecruitByIds 批量删除StudentRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentRecruitService *StudentRecruitService) DeleteStudentRecruitByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.StudentRecruit{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateStudentRecruit 更新StudentRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentRecruitService *StudentRecruitService) UpdateStudentRecruit(studentRecruit autocode.StudentRecruit) (err error) {
	err = global.GVA_DB.Save(&studentRecruit).Error
	return err
}

// GetStudentRecruit 根据id获取StudentRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentRecruitService *StudentRecruitService) GetStudentRecruit(id uint) (err error, studentRecruit autocode.StudentRecruit) {
	err = global.GVA_DB.Where("id = ?", id).First(&studentRecruit).Error
	return
}

// GetStudentRecruitInfoList 分页获取StudentRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentRecruitService *StudentRecruitService) GetStudentRecruitInfoList(info autoCodeReq.StudentRecruitSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.StudentRecruit{})
	var studentRecruits []autocode.StudentRecruit
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ComId != nil {
		db = db.Where("com_id = ?", info.ComId)
	}
	if info.UId != nil {
		db = db.Where("u_id = ?", info.UId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&studentRecruits).Error
	return err, studentRecruits, total
}

func (studentRecruitService *StudentRecruitService) ProduceStudentRecruitInfomation(stu autocode.StudentInfo, info autocode.StudentRecruit) (err error) {
	//首先选择简单的工作模式 hello_world
	/*
		思路：
		传递消息：
			1.学生id
			2.学生联系方式qq//微信
		传递对象：
			发布消息的负责人
		程序步骤
			1.创建生产者
			2.根据需要传递消息生成str 序列化
			3.关闭传送
	*/
	helloProducer, err := hello_world.CreateProducer()
	if err != nil {
		global.GVA_LOG.Error("创建生产者失败", zap.Error(err))
		os.Exit(1)
	}
	//将生产者信息序列化，作为消息进行发布
	infomation := autocode.ProduceComsumeInfomation{}
	infomation.Producer = int(stu.ID)
	infomation.Comsumer = *info.UId
	infomation.FormId = *info.EntryId
	data, err := json.Marshal(infomation)
	if err != nil {
		global.GVA_LOG.Error("序列化失败", zap.Error(err))
		return
	}
	//开始发送消息
	res := helloProducer.Send(data)
	helloProducer.Close() // 消息投递结束，必须关闭连接
	if res {
		return nil
	} else {
		return errors.New("发送失败")
	}
}
