<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="人事编号:">
          <el-input v-model="formData.personnelId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="教务编号:">
          <el-input v-model="formData.officeId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="财务编号:">
          <el-input v-model="formData.financialId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户编号:">
          <el-input v-model.number="formData.uId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="昵称:">
          <el-input v-model="formData.nickname" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="邮箱:">
          <el-input v-model="formData.email" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头像:">
          <el-input v-model="formData.avatar" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:">
          <el-input v-model="formData.realName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="性别:">
          <el-select v-model="formData.gender" placeholder="请选择" clearable>
            <el-option
              v-for="(item, key) in genderOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="学院:">
          <el-input v-model="formData.department" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="专业方向:">
          <el-input v-model="formData.major" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="职称:">
          <el-input v-model="formData.position" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="特长:">
          <el-input v-model="formData.specialty" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="学历:">
          <el-input v-model="formData.degree" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="银行名称:">
          <el-input v-model="formData.bankName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="银行卡号:">
          <el-input v-model="formData.bankCardNumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="介绍:">
          <el-input v-model="formData.introduction" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'TeacherInfo'
}
</script>

<script setup>
import {
  createTeacherInfo,
  updateTeacherInfo,
  findTeacherInfo
} from '@/api/teacherInfo'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, defineProps, watch } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const genderOptions = ref([])
const formData = ref({
  personnelId: '',
  officeId: '',
  financialId: '',
  uId: 0,
  nickname: '',
  email: '',
  avatar: '',
  realName: '',
  gender: undefined,
  department: '',
  major: '',
  position: '',
  specialty: '',
  degree: '',
  bankName: '',
  bankCardNumber: '',
  introduction: '',
})

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findTeacherInfo({ uId: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.reteacherInfo
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }

  genderOptions.value = await getDictFunc('gender')
}

init()
// 保存按钮
const save = async () => {
  let res
  switch (type.value) {
    case 'create':
      res = await createTeacherInfo(formData.value)
      break
    case 'update':
      res = await updateTeacherInfo(formData.value)
      break
    default:
      res = await createTeacherInfo(formData.value)
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

// =========== 自定义部分 ===========
const props = defineProps({
  uid: {
    type: Number,
    default: 0
  }
})
// 监听count
watch(
  () => props.uid,
  async (newVal, oldVal) => {
    console.log('%c 🥔 oldVal: ', 'font-size:20px;background-color: #E41A6A;color:#fff;', oldVal)
    console.log('%c 🍉 newVal: ', 'font-size:20px;background-color: #4b4b4b;color:#fff;', newVal)
    if (newVal > 0) {
      const res = await findTeacherInfo({ uId: newVal })
      if (res.code === 0) {
        formData.value = res.data.reteacherInfo
        type.value = 'update'
      }
    }
  },
  {
    immediate: true, // 立即执行
    deep: true // 深度监听
  }
)
</script>

<style>
</style>
