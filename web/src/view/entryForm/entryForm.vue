<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="参赛表名">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="竞赛编号">
          <el-input v-model="searchInfo.cmpId" placeholder="搜索条件" />
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
          label="参赛表名"
          prop="name"
          width="120"
        />
        <el-table-column
          align="left"
          label="竞赛编号"
          prop="competitionName"
          width="120"
        />
        <el-table-column align="left" label="项目编号" prop="pId" width="120">
          <template #default="scope">
            <el-button
              type="text"
              size="small"
              class="table-button"
              @click="OpenDialogProject(scope.row)"
            >
              {{ scope.row.pId == 0 ? '创建项目' : '查看项目' }}
            </el-button>
          </template>
        </el-table-column>

        <el-table-column align="left" label="按钮组" width="240">
          <template #default="scope">
            <el-button
              type="text"
              icon="edit"
              size="small"
              class="table-button"
              @click="updateEntryFormFunc(scope.row)"
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
            <el-button
              type="text"
              icon="view"
              size="small"
              @click="goEntryMember(scope.row)"
            >
              人员详情
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
        <el-form-item label="参赛表名:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="竞赛编号:">
          <el-input
            v-model.number="formData.cmpId"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="项目编号:">
          <el-input
            v-model.number="formData.pId"
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
    <!-- 项目弹出框：用作查看与创建 -->
    <el-dialog
      v-model="dialogProjectFormVisible"
      :before-close="closeDialogProject"
      title="弹窗操作"
    >
      <el-form
        :model="ProjectformData"
        label-position="right"
        label-width="80px"
      >
        <el-form-item label="项目编号:">
          <el-input
            v-model="ProjectformData.projectCode"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="项目名称:">
          <el-input
            v-model="ProjectformData.projectName"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="项目简介:">
          <el-input
            v-model="ProjectformData.introduction"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="备注:">
          <el-input
            v-model="ProjectformData.remark"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialogProject">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialogProject">
            确 定
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'EntryForm',
}
</script>

<script setup>
import {
  createEntryForm,
  deleteEntryForm,
  deleteEntryFormByIds,
  updateEntryForm,
  findEntryForm,
  getEntryFormList,
} from '@/api/entryForm'
import {
  createProjectInfo,
  updateProjectInfo,
  findProjectInfo,
} from '@/api/projectInfo'
import { findCompetitionSche } from '@/api/competitionSche'
import { findCompetitionInfo } from '@/api/competitionInfo'
// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
const router = useRouter()
// 自动化生成的字典（可能为空）以及字段
const awardOptions = ref([])
const formData = ref({
  name: '',
  cmpId: 0,
  pId: 0,
  rank: undefined,
  achName: '',
  competitionName: '无',
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
  const table = await getEntryFormList({
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
  getFormCompetitionInfo()
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
  awardOptions.value = await getDictFunc('award')
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
    deleteEntryFormFunc(row)
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
  const res = await deleteEntryFormByIds({ ids })
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
const updateEntryFormFunc = async (row) => {
  const res = await findEntryForm({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reentryForm
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteEntryFormFunc = async (row) => {
  const res = await deleteEntryForm({ ID: row.ID })
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
    cmpId: 0,
    pId: 0,
    rank: undefined,
    achName: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
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
      message: '创建/更改成功',
    })
    closeDialog()
    getTableData()
  }
}

// ============== 自定义部分 ===============
const goEntryMember = async (param) => {
  router.push({ name: 'entryMember', query: { formId: param.ID } })
}

// ========= 项目弹出框部分=============
const ProjectformData = ref({
  projectCode: '',
  projectName: '',
  introduction: '',
  remark: '',
})

// 弹窗控制标记
const dialogProjectFormVisible = ref(false)

// 获取项目的详情信息
const getProjectInfoFunc = async (id) => {
  const res = await findProjectInfo({ ID: id })
  if (res.code === 0) {
    ProjectformData.value = res.data.reprojectInfo
    dialogProjectFormVisible.value = true
  }
}
// 打开弹窗
const OpenDialogProject = (row) => {
  if (row.pId === 0) {
    typeProject.value = 'create'
  } else {
    typeProject.value = 'update'
    getProjectInfoFunc(row.pId)
  }
  formData.value = row
  dialogProjectFormVisible.value = true
}
// 行为控制标记（弹窗内部需要增还是改）
const typeProject = ref('')
// 关闭弹窗
const closeDialogProject = () => {
  dialogProjectFormVisible.value = false
  ProjectformData.value = {
    projectCode: '',
    projectName: '',
    introduction: '',
    remark: '',
  }
}
// 弹窗确定
const enterDialogProject = async () => {
  let res
  switch (typeProject.value) {
    case 'create':
      res = await createProjectInfo(ProjectformData.value)
      break
    case 'update':
      res = await updateProjectInfo(ProjectformData.value)
      break
    default:
      res = await createProjectInfo(ProjectformData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功',
    })
    if (typeProject.value === 'create') {
      formData.value.pId = res.data
      const r = await updateEntryForm(formData.value)
      if (r.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建成功',
        })
      }
    }
    closeDialogProject()
  }
}
// ====================比赛信息======================
const getFormCompetitionInfo = async () => {
  for (let i = 0; i < tableData.value.length; i++) {
    const res = await findCompetitionSche({
      ID: tableData.value[i].cmpId,
    })
    if (res.code === 0) {
      const ans = await findCompetitionInfo({
        ID: res.data.recompetitionSche.cId,
      })
      if (ans.code === 0) {
        tableData.value[i].competitionName = ans.data.recompetitionInfo.cName
      }
    }
  }
}
</script>

<style></style>
