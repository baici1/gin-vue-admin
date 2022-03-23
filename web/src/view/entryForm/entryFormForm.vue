<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="参赛表名:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="竞赛编号:">
          <el-input v-model.number="formData.cmpId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目编号:">
          <el-input v-model.number="formData.pId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="获奖级别:">
          <el-select v-model="formData.rank" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in awardOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="获奖名称:">
          <el-input v-model="formData.achName" clearable placeholder="请输入" />
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
  name: 'EntryForm'
}
</script>

<script setup>
import {
  createEntryForm,
  updateEntryForm,
  findEntryForm
} from '@/api/entryForm'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const awardOptions = ref([])
const formData = ref({
        name: '',
        cmpId: 0,
        pId: 0,
        rank: undefined,
        achName: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEntryForm({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reentryForm
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    awardOptions.value = await getDictFunc('award')
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createEntryForm(formData.value)
          break
        case 'update':
          res = await updateEntryForm(formData.value)
          break
        default:
          res = await createEntryForm(formData.value)
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
