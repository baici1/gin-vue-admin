package initialize

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/es"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"log"
	"os"
)

var CreateIndexExists = errors.New("已经创建")

type EsIndex struct {
	Index   string
	Mapping string
}

func GetEsClient() *elastic.Client {
	es := global.GVA_CONFIG.Elasticsearch
	cli, err := elastic.NewSimpleClient(
		elastic.SetURL(es.Dsn()),
		elastic.SetSniff(false),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "", log.LstdFlags)), // 设置错误日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),  // 设置info日志输出
	)
	if err != nil {
		panic("new es client failed, err=" + err.Error())
	}
	return cli
}

// ESIndexExists 索引是否存在。
func ESIndexExists(ctx context.Context, index string) (bool, error) {
	return global.GVA_ES.IndexExists(index).Do(ctx)
}

// CrtESIndex 创建 ES 索引。
func CrtESIndex(ctx context.Context, index, desc string) error {
	exist, err := ESIndexExists(ctx, index)
	if err != nil {
		return err
	}
	// 已经创建
	if exist {
		return CreateIndexExists
	}
	// 重复创建会报错
	_, err = global.GVA_ES.CreateIndex(index).BodyString(desc).Do(ctx)
	return err
}

func InitEsAllIndex() {
	var index = []EsIndex{
		{
			es.ArticleIndex,
			es.ArticleMapping,
		},
	}
	for _, v := range index {
		ctx := context.Background()
		err := CrtESIndex(ctx, v.Index, v.Mapping)
		if err != nil {
			if errors.Is(err, CreateIndexExists) {
				continue
			}
			global.GVA_LOG.Error("create Index "+v.Index+" failed", zap.Error(err))
		}
	}
}
