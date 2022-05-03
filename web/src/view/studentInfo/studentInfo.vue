<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="学号">
          <el-input v-model="searchInfo.studentId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="姓名">
          <el-input v-model="searchInfo.realName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="性别">
          <el-input v-model="searchInfo.gender" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="年级">
          <el-input v-model="searchInfo.grade" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="学院">
          <el-input v-model="searchInfo.department" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="专业">
          <el-input v-model="searchInfo.major" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="班级">
          <el-input v-model="searchInfo.classNum" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button
            size="small"
            type="primary"
            icon="search"
            @click="onSubmit"
          >
            查询
          </el-button>
          <el-button size="small" icon="refresh" @click="onReset">
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">
          新增
        </el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button size="small" type="text" @click="deleteVisible = false">
              取消
            </el-button>
            <el-button size="small" type="primary" @click="onDelete">
              确定
            </el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              size="small"
              style="margin-left: 10px"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
            >
              删除
            </el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="学号"
          prop="studentId"
          width="120"
        />
        <el-table-column align="left" label="手机号" prop="phone" width="120" />

        <el-table-column
          align="left"
          label="用户身份"
          prop="authorityId"
          width="120"
        />
        <el-table-column
          align="left"
          label="昵称"
          prop="nickname"
          width="120"
        />
        <el-table-column align="left" label="邮箱" prop="email" width="120" />
        <el-table-column align="left" label="头像" prop="avatar" width="120">
          <template #default="scope">
            <el-image
              :src="scope.row.avatar"
              style="width: 50px; height: 50px"
            />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="姓名"
          prop="realName"
          width="120"
        />
        <el-table-column align="left" label="性别" prop="gender" width="120">
          <template #default="scope">{{
            filterDict(scope.row.gender, genderOptions)
          }}</template>
        </el-table-column>
        <el-table-column align="left" label="学历" prop="degree" width="120" />
        <el-table-column align="left" label="年级" prop="grade" width="120" />
        <el-table-column
          align="left"
          label="学院"
          prop="department"
          width="120"
        />
        <el-table-column align="left" label="专业" prop="major" width="120" />
        <el-table-column
          align="left"
          label="班级"
          prop="classNum"
          width="120"
        />
        <el-table-column
          align="left"
          label="特长"
          prop="specialty"
          width="120"
        />
        <el-table-column align="left" label="QQ号" prop="QQ" width="120" />
        <el-table-column
          align="left"
          label="微信号"
          prop="wechat"
          width="120"
        />
        <el-table-column
          align="left"
          label="银行名称"
          prop="bankName"
          width="120"
        />
        <el-table-column
          align="left"
          label="银行卡号"
          prop="bankCardNumber"
          width="120"
        />
        <el-table-column
          align="left"
          label="介绍"
          prop="introduction"
          width="120"
          show-overflow-tooltip
        />
        <el-table-column align="left" label="按钮组" width="120" fixed>
          <template #default="scope">
            <el-button
              type="text"
              icon="edit"
              size="small"
              class="table-button"
              @click="updateStudentInfoFunc(scope.row)"
            >
              变更
            </el-button>
            <el-button
              type="text"
              icon="delete"
              size="small"
              @click="deleteRow(scope.row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      title="弹窗操作"
    >
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="学号:">
          <el-input
            v-model="formData.studentId"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="手机号:">
          <el-input v-model="formData.phone" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密码:">
          <el-input
            v-model="formData.password"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="用户身份:">
          <el-input
            v-model="formData.authorityId"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="昵称:">
          <el-input
            v-model="formData.nickname"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="邮箱:">
          <el-input v-model="formData.email" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头像:">
          <el-input v-model="formData.avatar" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:">
          <el-input
            v-model="formData.realName"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="性别:">
          <el-select
            v-model="formData.gender"
            placeholder="请选择"
            style="width: 100%"
            clearable
          >
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
          <el-input
            v-model="formData.department"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="专业:">
          <el-input v-model="formData.major" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="班级:">
          <el-input
            v-model="formData.classNum"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="特长:">
          <el-input
            v-model="formData.specialty"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="QQ号:">
          <el-input v-model="formData.QQ" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="微信号:">
          <el-input v-model="formData.wechat" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="银行名称:">
          <el-input
            v-model="formData.bankName"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="银行卡号:">
          <el-input
            v-model="formData.bankCardNumber"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="介绍:">
          <el-input
            v-model="formData.introduction"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">
            确 定
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'StudentInfo',
}
</script>

<script setup>
import {
  createStudentInfo,
  deleteStudentInfo,
  deleteStudentInfoByIds,
  updateStudentInfo,
  findStudentInfo,
  getStudentInfoList,
} from '@/api/studentInfo'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const genderOptions = ref([])
const formData = ref({
  studentId: '',
  phone: '',
  password: '',
  authorityId: '',
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

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getStudentInfoList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
  genderOptions.value = await getDictFunc('gender')
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    deleteStudentInfoFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async () => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据',
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map((item) => {
      ids.push(item.ID)
    })
  const res = await deleteStudentInfoByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功',
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateStudentInfoFunc = async (row) => {
  const res = await findStudentInfo({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.restudentInfo
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteStudentInfoFunc = async (row) => {
  const res = await deleteStudentInfo({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功',
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    studentId: '',
    phone: '',
    password: '',
    authorityId: '',
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
  }
}
// 弹窗确定
const enterDialog = async () => {
  let res
  switch (type.value) {
    case 'create':
      res = await createStudentInfo(formData.value)
      break
    case 'update':
      res = await updateStudentInfo(formData.value)
      break
    default:
      res = await createStudentInfo(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功',
    })
    closeDialog()
    getTableData()
  }
}
</script>

<style></style>
