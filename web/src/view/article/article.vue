<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="发布时间">
          <!-- <el-input v-model="searchInfo.publishedTime" placeholder="搜索条件" /> -->
          <el-date-picker
            v-model="searchInfo.publishedTime"
            type="date"
            style="width:100%"
            placeholder="选择日期"
            clearable
          />
        </el-form-item>
        <el-form-item label="标题">
          <el-input v-model="searchInfo.title" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="作者">
          <el-input v-model="searchInfo.author" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="文章类型">
          <el-select v-model="searchInfo.type" placeholder="搜索条件" style="width:100%" clearable>
            <el-option
              v-for="(item, key) in articleTypeOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="goCreate">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="text" @click="deleteVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              size="small"
              style="margin-left: 10px;"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
            >删除</el-button>
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
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="评论" prop="commentabled" width="60">
          <template #default="scope">{{ formatBoolean(scope.row.commentabled) }}</template>
        </el-table-column>
        <el-table-column align="left" label="发表" prop="published" width="60">
          <template #default="scope">{{ formatBoolean(scope.row.published) }}</template>
        </el-table-column>
        <el-table-column align="left" label="发布时间" prop="publishedTime" width="180">
          <template #default="scope">{{ formatDate(scope.row.publishedTime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="标题" prop="title" width="180" />
        <!-- <el-table-column align="left" label="引用" prop="description" width="120" /> -->
        <!-- <el-table-column
          align="left"
          label="内容"
          class="truncate"
          prop="content"
          width="220"
          style="max-height: 200px;"
        />-->
        <el-table-column align="left" label="文章内容" prop="views" width="100">
          <template #default="scope">
            <el-button type="text" @click="updateArticleFunc(scope.row)">查看文章</el-button>
          </template>
        </el-table-column>
        <el-table-column align="left" label="浏览量" prop="views" width="70" />
        <el-table-column align="left" label="作者" prop="author" width="120" />
        <el-table-column align="left" label="文章类型" prop="type" width="120">
          <template #default="scope">{{ filterDict(scope.row.type, articleTypeOptions) }}</template>
        </el-table-column>
        <el-table-column align="left" label="置顶" prop="orderNum" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.orderNum) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button
              type="text"
              icon="edit"
              size="small"
              class="table-button"
              @click="goUpdate(scope.row)"
            >变更</el-button>
            <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="是否可评论:">
          <el-switch
            v-model="formData.commentabled"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            clearable
          />
        </el-form-item>
        <el-form-item label="是否发表:">
          <el-switch
            v-model="formData.published"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            clearable
          />
        </el-form-item>
        <el-form-item label="发布时间:">
          <el-date-picker
            v-model="formData.publishedTime"
            type="date"
            style="width:100%"
            placeholder="选择日期"
            clearable
          />
        </el-form-item>
        <el-form-item label="标题:">
          <el-input v-model="formData.title" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="引用:">
          <el-input v-model="formData.description" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容:">
          <el-input v-model="formData.content" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="浏览量:">
          <el-input v-model.number="formData.views" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="作者:">
          <el-input v-model="formData.author" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文章类型:">
          <el-select v-model="formData.type" placeholder="请选择" style="width:100%" clearable>
            <el-option
              v-for="(item, key) in articleTypeOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="置顶:">
          <el-switch
            v-model="formData.orderNum"
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
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <el-drawer v-model="drawer" :title="formData.title" direction="rtl">
      <v-html>{{ formData.content }}</v-html>
    </el-drawer>
  </div>
</template>

<script>
export default {
  name: 'Article'
}
</script>

<script setup>
import {
  createArticle,
  deleteArticle,
  deleteArticleByIds,
  updateArticle,
  findArticle,
  getArticleList
} from '@/api/article'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
const router = useRouter()
const drawer = ref(false)

// 自动化生成的字典（可能为空）以及字段
const articleTypeOptions = ref([])
const formData = ref({
  commentabled: false,
  published: false,
  publishedTime: new Date(),
  title: '',
  description: '',
  content: '',
  views: 0,
  author: '',
  type: undefined,
  orderNum: false,
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
  if (searchInfo.value.commentabled === '') {
    searchInfo.value.commentabled = null
  }
  if (searchInfo.value.published === '') {
    searchInfo.value.published = null
  }
  if (searchInfo.value.orderNum === '') {
    searchInfo.value.orderNum = null
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
  const table = await getArticleList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  articleTypeOptions.value = await getDictFunc('articleType')
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
    type: 'warning'
  }).then(() => {
    deleteArticleFunc(row)
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
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.ID)
    })
  const res = await deleteArticleByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
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
const updateArticleFunc = async (row) => {
  const res = await findArticle({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.rearticle
    // dialogFormVisible.value = true
    drawer.value = true
  }
}

// 删除行
const deleteArticleFunc = async (row) => {
  const res = await deleteArticle({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// // 打开弹窗
// const openDialog = () => {
//   type.value = 'create'
//   dialogFormVisible.value = true
// }

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    commentabled: false,
    published: false,
    publishedTime: new Date(),
    title: '',
    description: '',
    content: '',
    views: 0,
    author: '',
    type: undefined,
    orderNum: false,
  }
}
// 弹窗确定
const enterDialog = async () => {
  let res
  switch (type.value) {
    case 'create':
      res = await createArticle(formData.value)
      break
    case 'update':
      res = await updateArticle(formData.value)
      break
    default:
      res = await createArticle(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
    closeDialog()
    getTableData()
  }
}

// ============自定义部分开始 ===============
const goUpdate = (param) => {
  router.push({
    name: 'articleForm', query: {
      id: param.ID
    }
  })
}
const goCreate = () => {
  router.push({
    name: 'articleForm'
  })
}
</script>

<style>
</style>
