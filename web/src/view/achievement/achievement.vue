<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="赛事级别">
          <el-input v-model="searchInfo.match" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="获奖级别">
          <el-input v-model="searchInfo.rank" placeholder="搜索条件" />
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
          label="队伍id"
          prop="formId"
          width="120"
        />
        <el-table-column align="left" label="赛事级别" prop="match" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.match, competitionLevelOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="获奖级别" prop="rank" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.rank, awardOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="备注" prop="remark" width="120" />
        <el-table-column align="left" label="成果附件" prop="url" width="120" />
        <el-table-column align="left" label="审核" prop="check" width="120">
          <template #default="scope">{{
            formatBoolean(scope.row.check)
          }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button
              type="text"
              icon="edit"
              size="small"
              class="table-button"
              @click="updateAchievementFunc(scope.row)"
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
        <el-form-item label="队伍id:">
          <el-input
            v-model.number="formData.formId"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="赛事级别:">
          <el-select
            v-model="formData.match"
            placeholder="请选择"
            style="width: 100%"
            clearable
          >
            <el-option
              v-for="(item, key) in competitionLevelOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="获奖级别:">
          <el-select
            v-model="formData.rank"
            placeholder="请选择"
            style="width: 100%"
            clearable
          >
            <el-option
              v-for="(item, key) in awardOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="成果附件:">
          <el-input v-model="formData.url" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核:">
          <el-switch
            v-model="formData.check"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            clearable
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
  name: 'Achievement',
}
</script>

<script setup>
import {
  createAchievement,
  deleteAchievement,
  deleteAchievementByIds,
  updateAchievement,
  findAchievement,
  getAchievementList,
} from '@/api/achievement'

// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
} from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const competitionLevelOptions = ref([])
const awardOptions = ref([])
const formData = ref({
  formId: 0,
  match: undefined,
  rank: undefined,
  remark: '',
  url: '',
  check: false,
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
  if (searchInfo.value.check === '') {
    searchInfo.value.check = null
  }
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
  const table = await getAchievementList({
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
  competitionLevelOptions.value = await getDictFunc('competitionLevel')
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
    deleteAchievementFunc(row)
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
  const res = await deleteAchievementByIds({ ids })
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
const updateAchievementFunc = async (row) => {
  const res = await findAchievement({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reachievement
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteAchievementFunc = async (row) => {
  const res = await deleteAchievement({ ID: row.ID })
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
    formId: 0,
    match: undefined,
    rank: undefined,
    remark: '',
    url: '',
    check: false,
  }
}
// 弹窗确定
const enterDialog = async () => {
  let res
  switch (type.value) {
    case 'create':
      res = await createAchievement(formData.value)
      break
    case 'update':
      res = await updateAchievement(formData.value)
      break
    default:
      res = await createAchievement(formData.value)
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
