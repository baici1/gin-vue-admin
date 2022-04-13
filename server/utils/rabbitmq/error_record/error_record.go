package error_record

import "github.com/flipped-aurora/gin-vue-admin/server/global"

func ErrorDeal(err error) error {
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}
	return err
}
