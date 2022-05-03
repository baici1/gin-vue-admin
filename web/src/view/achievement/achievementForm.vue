<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="队伍:">
          <el-input
            v-model.number="formData.formId"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="比赛:">
          <el-input
            v-model.number="formData.comId"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="赛事级别:">
          <el-select v-model="formData.match" placeholder="请选择" clearable>
            <el-option
              v-for="(item, key) in competitionLevelOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="获奖级别:">
          <el-select v-model="formData.rank" placeholder="请选择" clearable>
            <el-option
              v-for="(item, key) in awardOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="成果附件:">
          <el-input v-model="formData.url" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核:">
          <el-switch
            v-model="formData.check"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            clearable
          />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Achievement',
}
</script>

<script setup>
import {
  createAchievement,
  updateAchievement,
  findAchievement,
} from '@/api/achievement'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const awardOptions = ref([])
const competitionLevelOptions = ref([])
const formData = ref({
  formId: 0,
  comId: 0,
  match: undefined,
  rank: undefined,
  remark: '',
  url: '',
  check: false,
})

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findAchievement({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.reachievement
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
  awardOptions.value = await getDictFunc('award')
  competitionLevelOptions.value = await getDictFunc('competitionLevel')
}

init()
// 保存按钮
const save = async () => {
  let res
  switch (type.value) {
    case 'create':
      res = await createAchievement(formData.value)
      break
    case 'update':
      res = await updateAchievement(formData.value)
      break
    default:
      res = await createAchievement(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功',
    })
  }
}

// 返回按钮
const back = () => {
  router.go(-1)
}
</script>

<style></style>
