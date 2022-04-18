<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="响应者:">
          <el-input v-model.number="formData.producer" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发布者:">
          <el-input v-model.number="formData.comsumer" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="参赛表:">
          <el-input v-model.number="formData.formId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否查看:">
          <el-switch v-model="formData.check" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  name: 'ProduceComsumeInfomation'
}
</script>

<script setup>
import {
  createProduceComsumeInfomation,
  updateProduceComsumeInfomation,
  findProduceComsumeInfomation
} from '@/api/produceComsumeInfomation'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        producer: 0,
        comsumer: 0,
        formId: 0,
        check: false,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findProduceComsumeInfomation({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reproduceComsumeInfomation
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
          res = await createProduceComsumeInfomation(formData.value)
          break
        case 'update':
          res = await updateProduceComsumeInfomation(formData.value)
          break
        default:
          res = await createProduceComsumeInfomation(formData.value)
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
