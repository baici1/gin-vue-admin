<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="竞赛编号:">
          <el-input v-model.number="formData.cId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="竞赛级别:">
          <el-select v-model="formData.level" placeholder="请选择" clearable>
            <el-option
              v-for="(item, key) in competitionLevelOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="比赛的届数:">
          <el-input v-model.number="formData.version" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="报名开始日期:">
          <el-date-picker v-model="formData.startTime" type="date" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="报名截止日期:">
          <el-date-picker v-model="formData.endTime" type="date" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="比赛开始时间:">
          <el-date-picker v-model="formData.rStartTime" type="date" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="比赛结束时间:">
          <el-date-picker v-model="formData.rEndTime" type="date" placeholder="选择日期" clearable />
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
  name: 'CompetitionSche'
}
</script>

<script setup>
import {
  createCompetitionSche,
  updateCompetitionSche,
  findCompetitionSche
} from '@/api/competitionSche'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const competitionLevelOptions = ref([])
const formData = ref({
  cId: 0,
  level: undefined,
  version: 0,
  startTime: new Date(),
  endTime: new Date(),
  rStartTime: new Date(),
  rEndTime: new Date(),
})

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findCompetitionSche({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.recompetitionSche
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
  competitionLevelOptions.value = await getDictFunc('competitionLevel')
}

init()
// 保存按钮
const save = async () => {
  let res
  switch (type.value) {
    case 'create':
      res = await createCompetitionSche(formData.value)
      break
    case 'update':
      res = await updateCompetitionSche(formData.value)
      break
    default:
      res = await createCompetitionSche(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
  }
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
