<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="项目编号:">
          <el-input v-model="formData.projectCode" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目名称:">
          <el-input v-model="formData.projectName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目简介:">
          <el-input v-model="formData.introduction" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
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
  name: 'ProjectInfo'
}
</script>

<script setup>
import {
  createProjectInfo,
  updateProjectInfo,
  findProjectInfo
} from '@/api/projectInfo'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        projectCode: '',
        projectName: '',
        introduction: '',
        remark: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findProjectInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reprojectInfo
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createProjectInfo(formData.value)
          break
        case 'update':
          res = await updateProjectInfo(formData.value)
          break
        default:
          res = await createProjectInfo(formData.value)
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
