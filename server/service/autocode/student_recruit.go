package autocode

import (
	"encoding/json"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/rabbitmq/hello_world"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/rabbitmq/routing"
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
	err = global.GVA_DB.Where("id = ?", id).Preload("Competition").Preload("Competition.BaseInfo").Preload("Member").First(&studentRecruit).Error
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
	err = db.Limit(limit).Offset(offset).Preload("Competition").Preload("Competition.BaseInfo").Preload("Member").Find(&studentRecruits).Error
	return err, studentRecruits, total
}
func (studentRecruitService *StudentRecruitService) ProduceStudentRecruitInfomationByHello(stu autoCodeReq.StudentRequestInfo, info autocode.StudentRecruit) (err error) {
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
	//对信息进行校验，保证后续功能能正常进行下去。
	if stu.QQ == "" || stu.Wechat == "" || *info.Num <= 0 {
		return errors.New("很抱歉！不能继续申请。")
	}
	helloProducer, err := hello_world.CreateProducer()
	if err != nil {
		global.GVA_LOG.Error("创建生产者失败", zap.Error(err))
		os.Exit(1)
	}
	//将生产者信息序列化，作为消息进行发布
	infomation := autoCodeReq.StudentRecruitToRabbitmq{
		Producer: stu,
		Comsumer: info,
	}
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

func (studentRecruitService *StudentRecruitService) ProduceRecruitInfomationByRouting(stu autoCodeReq.StudentRequestInfo, info autoCodeReq.ComsumerRequestInfo) (err error) {
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
	//对信息进行校验，保证后续功能能正常进行下去。
	if stu.QQ == "" || stu.Wechat == "" || stu.RealName == "" {
		return errors.New("很抱歉！不能继续申请。")
	}
	producer, err := routing.CreateProducer()
	if err != nil {
		global.GVA_LOG.Error("创建生产者失败", zap.Error(err))
		os.Exit(1)
	}
	//将生产者信息序列化，作为消息进行发布
	infomation := autoCodeReq.RecruitToRabbitmqInfo{
		Producer: stu,
		Comsumer: info,
	}
	data, err := json.Marshal(infomation)
	if err != nil {
		global.GVA_LOG.Error("序列化失败", zap.Error(err))
		return
	}
	//开始发送消息
	res := producer.Send(info.Key, data, 10000)
	producer.Close() // 消息投递结束，必须关闭连接
	if res {
		return nil
	} else {
		return errors.New("发送失败")
	}
}
