<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="手机号:">
          <el-input v-model="formData.phone" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密码:">
          <el-input v-model="formData.password" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.slat" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="身份:">
          <el-select v-model="formData.identity" placeholder="请选择" clearable>
            <el-option
              v-for="(item, key) in identityOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="审核:">
          <el-select v-model="formData.check" placeholder="请选择" clearable>
            <el-option
              v-for="(item, key) in checkOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
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
  name: 'UserInfo',
}
</script>

<script setup>
import { createUserInfo, updateUserInfo, findUserInfo } from '@/api/userInfo'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const identityOptions = ref([])
const checkOptions = ref([])
const formData = ref({
  password: '',
  phone: '',
  slat: '',
  identity: undefined,
  check: undefined,
})

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findUserInfo({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.reuserInfo
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
  identityOptions.value = await getDictFunc('identity')
  checkOptions.value = await getDictFunc('check')
}

init()
// 保存按钮
const save = async () => {
  let res
  switch (type.value) {
    case 'create':
      res = await createUserInfo(formData.value)
      break
    case 'update':
      res = await updateUserInfo(formData.value)
      break
    default:
      res = await createUserInfo(formData.value)
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
