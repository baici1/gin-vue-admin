package autocode

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type EsService struct {
}

func (e *EsService) Create(index, id, json string) error {
	ctx := context.Background()
	_, err := global.GVA_ES.Index().Index(index).OpType("create").Id(id).BodyJson(json).Refresh("true").Do(ctx)
	return err
}

// Delete 通过 ID 删除文档
func (e *EsService) Delete(index, id string) error {
	ctx := context.Background()
	_, err := global.GVA_ES.Delete().Index(index).Id(id).Refresh("true").Do(ctx)
	return err
}

// Update 修改文档
// param: index 索引; id 文档ID; m 待更新的字段键值结构
func (e *EsService) Update(index, id string, doc interface{}) error {
	ctx := context.Background()
	_, err := global.GVA_ES.
		Update().
		Index(index).
		Id(id).
		Doc(doc).
		Refresh("true").
		Do(ctx)
	return err
}
