<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="学号:">
          <el-input v-model="formData.studentId" clearable placeholder="请输入" />
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
        <el-form-item label="学历:">
          <el-input v-model="formData.degree" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="年级:">
          <el-input v-model="formData.grade" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="学院:">
          <el-input v-model="formData.department" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="专业:">
          <el-input v-model="formData.major" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="班级:">
          <el-input v-model="formData.classNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="特长:">
          <el-input v-model="formData.specialty" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="QQ号:">
          <el-input v-model="formData.QQ" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="微信号:">
          <el-input v-model="formData.wechat" clearable placeholder="请输入" />
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
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'StudentInfo'
}
</script>

<script setup>
import {
  updateStudentInfo,
  findStudentInfo
} from '@/api/studentInfo'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const genderOptions = ref([])
const formData = ref({
  studentId: '',
  uId: 0,
  nickname: '',
  email: '',
  avatar: '',
  realName: '',
  gender: undefined,
  degree: '',
  grade: '',
  department: '',
  major: '',
  classNum: '',
  specialty: '',
  QQ: '',
  wechat: '',
  bankName: '',
  bankCardNumber: '',
  introduction: '',
})

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findStudentInfo({ uId: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.restudentInfo
      type.value = 'update'
    }
  }
  genderOptions.value = await getDictFunc('gender')
}

init()
// 保存按钮
const save = async () => {
  let res
  if (type.value === 'update') {
    res = await updateStudentInfo(formData.value)
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '更改成功'
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
