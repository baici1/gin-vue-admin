<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="参赛表编号:">
          <el-input v-model.number="formData.formId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户编号:">
          <el-input v-model.number="formData.uId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份:">
          <el-select v-model="formData.identify" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in teamIdentifyOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="排序级别:">
          <el-input v-model.number="formData.order" clearable placeholder="请输入" />
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
  name: 'EntryMember'
}
</script>

<script setup>
import {
  createEntryMember,
  updateEntryMember,
  findEntryMember
} from '@/api/entryMember'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const teamIdentifyOptions = ref([])
const formData = ref({
        formId: 0,
        uId: 0,
        identify: undefined,
        order: 0,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEntryMember({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reentryMember
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    teamIdentifyOptions.value = await getDictFunc('teamIdentify')
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createEntryMember(formData.value)
          break
        case 'update':
          res = await updateEntryMember(formData.value)
          break
        default:
          res = await createEntryMember(formData.value)
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
