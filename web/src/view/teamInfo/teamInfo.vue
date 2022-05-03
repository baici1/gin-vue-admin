<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="团队名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
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
          label="团队名称"
          prop="name"
          width="120"
        />
        <el-table-column align="left" label="注册公司" width="120">
          <template #default="scope">
            <el-button
              type="text"
              size="small"
              class="table-button"
              @click="openDialogCompany(scope.row)"
            >
              {{ scope.row.companyId == 0 ? '创建公司' : '查看公司' }}
            </el-button>
          </template>
        </el-table-column>
        <el-table-column align="left" label="负责人" prop="uId" width="120">
          <template #default="scope">
            <el-button
              type="text"
              size="small"
              class="table-button"
              @click="openDialogStudent(scope.row)"
            >
              查看负责人
            </el-button>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="团队介绍"
          prop="introduction"
          width="120"
        />
        <el-table-column
          align="left"
          label="知识产权"
          prop="intellectualProperty"
          width="120"
        />
        <el-table-column align="left" label="备注" prop="remark" width="120" />
        <el-table-column align="left" label="审核" prop="check" width="120">
          <template #default="scope">
            <el-tag
              class="ml-2"
              :type="ChooseTagType(scope.row.check, checkOptions)"
            >
              {{ filterDict(scope.row.check, checkOptions) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button
              type="text"
              icon="edit"
              size="small"
              class="table-button"
              @click="updateTeamInfoFunc(scope.row)"
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
            <!-- <el-button
              type="text"
              icon="view"
              size="small"
              @click="goTeamMember(scope.row)"
            >
              人员详情
            </el-button> -->
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
        <el-form-item label="团队名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="注册公司:">
          <el-input
            v-model.number="formData.companyId"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="团队介绍:">
          <el-input
            v-model="formData.introduction"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="知识产权:">
          <el-input
            v-model="formData.intellectualProperty"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核:">
          <el-select
            v-model="formData.check"
            placeholder="请选择"
            style="width: 100%"
            clearable
          >
            <el-option
              v-for="(item, key) in checkOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
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
    <!-- 公司弹框 -->
    <el-dialog
      v-model="dialogFormVisibleCompany"
      :before-close="closeDialogCompany"
      title="弹窗操作"
    >
      <el-form
        :model="formDataCompany"
        label-position="right"
        label-width="80px"
      >
        <el-form-item label="公司名称:">
          <el-input
            v-model="formDataCompany.name"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="公司地址:">
          <el-input
            v-model="formDataCompany.address"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="公司介绍:">
          <el-input
            v-model="formDataCompany.introduction"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="审核:">
          <el-select
            v-model="formDataCompany.check"
            placeholder="请选择"
            style="width: 100%"
            clearable
          >
            <el-option
              v-for="(item, key) in checkOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialogCompany">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialogCompany">
            确 定
          </el-button>
        </div>
      </template>
    </el-dialog>
    <!-- 学生弹框 -->
    <el-dialog
      v-model="dialogFormVisibleStudent"
      :before-close="closeDialogStudent"
      title="弹窗操作"
    >
      <el-form
        :model="formDataStudent"
        label-position="right"
        label-width="80px"
      >
        <el-form-item label="学号:">
          <el-input
            v-model="formDataStudent.studentId"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="手机号:">
          <el-input
            v-model="formDataStudent.phone"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="用户身份:">
          <el-input
            v-model="formDataStudent.authorityId"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="昵称:">
          <el-input
            v-model="formDataStudent.nickname"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="邮箱:">
          <el-input
            v-model="formDataStudent.email"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="头像:">
          <el-input
            v-model="formDataStudent.avatar"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="姓名:">
          <el-input
            v-model="formDataStudent.realName"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="性别:">
          <el-select
            v-model="formDataStudent.gender"
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
          <el-input
            v-model="formDataStudent.degree"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="年级:">
          <el-input
            v-model="formDataStudent.grade"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="学院:">
          <el-input
            v-model="formDataStudent.department"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="专业:">
          <el-input
            v-model="formDataStudent.major"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="班级:">
          <el-input
            v-model="formDataStudent.classNum"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="特长:">
          <el-input
            v-model="formDataStudent.specialty"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="QQ号:">
          <el-input
            v-model="formDataStudent.QQ"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="微信号:">
          <el-input
            v-model="formDataStudent.wechat"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="银行名称:">
          <el-input
            v-model="formDataStudent.bankName"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="银行卡号:">
          <el-input
            v-model="formDataStudent.bankCardNumber"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="介绍:">
          <el-input
            v-model="formDataStudent.introduction"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialogStudent">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'TeamInfo',
}
</script>

<script setup>
import { ChooseTagType } from '@/utils/tools'
import {
  createTeamInfo,
  deleteTeamInfo,
  deleteTeamInfoByIds,
  updateTeamInfo,
  findTeamInfo,
  getTeamInfoList,
} from '@/api/teamInfo'
import {
  createCompanyInfo,
  updateCompanyInfo,
  findCompanyInfo,
} from '@/api/companyInfo'
import { findStudentInfo } from '@/api/studentInfo'
// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
// 自动化生成的字典（可能为空）以及字段
const checkOptions = ref([])
const genderOptions = ref([])
const formData = ref({
  name: '',
  companyId: 0,
  introduction: '',
  intellectualProperty: '',
  remark: '',
  check: undefined,
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
  const table = await getTeamInfoList({
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
  checkOptions.value = await getDictFunc('check')
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
    deleteTeamInfoFunc(row)
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
  const res = await deleteTeamInfoByIds({ ids })
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
const updateTeamInfoFunc = async (row) => {
  const res = await findTeamInfo({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reteamInfo
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteTeamInfoFunc = async (row) => {
  const res = await deleteTeamInfo({ ID: row.ID })
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
    name: '',
    companyId: 0,
    introduction: '',
    intellectualProperty: '',
    remark: '',
    check: undefined,
  }
}
// 弹窗确定
const enterDialog = async () => {
  let res
  switch (type.value) {
    case 'create':
      res = await createTeamInfo(formData.value)
      break
    case 'update':
      res = await updateTeamInfo(formData.value)
      break
    default:
      res = await createTeamInfo(formData.value)
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
// ============== 自定义部分 ===============
// ===========公司弹框==============
const formDataCompany = ref({
  name: '',
  address: '',
  introduction: '',
  check: undefined,
})
// 行为控制标记（弹窗内部需要增还是改）
const typeCompany = ref('')
// 弹窗控制标记
const dialogFormVisibleCompany = ref(false)

const getCompanyInfoFunc = async (row) => {
  const res = await findCompanyInfo({ ID: row.companyId })
  if (res.code === 0) {
    formDataCompany.value = res.data.recompanyInfo
    dialogFormVisibleCompany.value = true
  }
}
// 打开弹窗
const openDialogCompany = (row) => {
  if (row.companyId === 0) {
    typeCompany.value = 'create'
  } else {
    typeCompany.value = 'update'
    getCompanyInfoFunc(row)
  }
  formData.value = row
  dialogFormVisibleCompany.value = true
}

// 关闭弹窗
const closeDialogCompany = () => {
  dialogFormVisibleCompany.value = false
  formDataCompany.value = {
    name: '',
    address: '',
    introduction: '',
    check: undefined,
  }
}
// 弹窗确定
const enterDialogCompany = async () => {
  let res
  switch (typeCompany.value) {
    case 'create':
      res = await createCompanyInfo(formDataCompany.value)
      break
    case 'update':
      res = await updateCompanyInfo(formDataCompany.value)
      break
    default:
      res = await createTeamInfo(formDataCompany.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功',
    })
    if (typeCompany.value === 'create') {
      formData.value.companyId = res.data
      const r = await updateTeamInfo(formData.value)
      if (r.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建成功',
        })
      }
    }
    closeDialogCompany()
    getTableData()
  }
}
// =============学生弹框==============
// 弹窗控制标记
const dialogFormVisibleStudent = ref(false)
const formDataStudent = ref({
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
// 打开弹窗
const openDialogStudent = async (row) => {
  const res = await findStudentInfo({ ID: row.uId })
  if (res.code === 0) {
    formDataStudent.value = res.data.restudentInfo
    dialogFormVisibleStudent.value = true
  }
}

// 关闭弹窗
const closeDialogStudent = () => {
  dialogFormVisibleStudent.value = false
  formDataStudent.value = {
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
</script>

<style></style>
