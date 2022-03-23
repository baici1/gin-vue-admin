<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="竞赛名称:">
          <el-input v-model="formData.cName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="竞赛类型:">
          <el-select v-model="formData.cType" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in competitionTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="主办单位:">
          <el-input v-model="formData.organizer" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="简介:">
          <el-input v-model="formData.introduction" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="竞赛官网:">
          <el-input v-model="formData.url" clearable placeholder="请输入" />
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
  name: 'CompetitionInfo'
}
</script>

<script setup>
import {
  createCompetitionInfo,
  updateCompetitionInfo,
  findCompetitionInfo
} from '@/api/competitionInfo'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const competitionTypeOptions = ref([])
const formData = ref({
        cName: '',
        cType: undefined,
        organizer: '',
        introduction: '',
        url: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCompetitionInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recompetitionInfo
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    competitionTypeOptions.value = await getDictFunc('competitionType')
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createCompetitionInfo(formData.value)
          break
        case 'update':
          res = await updateCompetitionInfo(formData.value)
          break
        default:
          res = await createCompetitionInfo(formData.value)
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
