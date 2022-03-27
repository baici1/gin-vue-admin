<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="轮播图名称">
          <el-input v-model="searchInfo.swiperName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="是否展示" prop="isShow">
          <el-select v-model="searchInfo.isShow" clearable placeholder="请选择">
            <el-option key="true" label="是" value="true" />
            <el-option key="false" label="否" value="false" />
          </el-select>
        </el-form-item>
        <el-form-item label="分组">
          <el-input v-model="searchInfo.group" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
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
        <el-table-column align="left" label="轮播图名称" prop="swiperName" width="120" />
        <el-table-column align="left" label="轮播图照片" prop="swiperPicture" width="120">
          <template #default="scope">
            <el-image
              :src="scope.row.swiperPicture"
              fit="cover"
              :preview-src-list="[scope.row.swiperPicture]"
              hide-on-click-modal
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="是否展示" prop="isShow" width="120">
          <template #default="scope">
            <el-switch
              v-model="scope.row.isShow"
              inline-prompt
              active-text="是"
              inactive-text="否"
              disabled
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="分组" prop="group" width="120">
          <template #default="scope">
            <el-tag>{{ filterDict(scope.row.group, imgGropuOptions) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="前往路径" prop="goToUrl" width="180" />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button
              type="text"
              icon="edit"
              size="small"
              class="table-button"
              @click="updateSwiperFunc(scope.row)"
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
        <el-form-item label="轮播图名称:">
          <el-input v-model="formData.swiperName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="轮播图照片:">
          <!-- <el-input v-model="formData.swiperPicture" clearable placeholder="请输入" /> -->
          <el-upload
            class="avatar-uploader"
            action="http://127.0.0.1:8888/fileUploadAndDownload/upload"
            :show-file-list="false"
            :on-success="handleAvatarSuccess"
            :before-upload="beforeAvatarUpload"
            :headers="{
              'x-token': token
            }"
          >
            <img v-if="formData.swiperPicture" :src="formData.swiperPicture" class="avatar" />
            <el-icon v-else class="avatar-uploader-icon">
              <Plus />
            </el-icon>
          </el-upload>
        </el-form-item>
        <el-form-item label="是否展示:">
          <el-switch
            v-model="formData.isShow"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            clearable
          />
        </el-form-item>
        <el-form-item label="分组:">
          <el-select v-model="formData.group" placeholder="请选择" style="width:100%" clearable>
            <el-option
              v-for="(item, key) in imgGropuOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="前往路径:">
          <el-input v-model="formData.goToUrl" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Swiper'
}
</script>

<script setup>
import {
  createSwiper,
  deleteSwiper,
  deleteSwiperByIds,
  updateSwiper,
  findSwiper,
  getSwiperList
} from '@/api/swiper'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { useUserStore } from '@/pinia/modules/user'

// 自动化生成的字典（可能为空）以及字段
const boolOptions = ref([])
const imgGropuOptions = ref([])
const formData = ref({
  swiperName: '',
  swiperPicture: '',
  isShow: false,
  group: undefined,
  goToUrl: '',
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
  if (searchInfo.value.isShow === '') {
    searchInfo.value.isShow = null
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
  const table = await getSwiperList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  boolOptions.value = await getDictFunc('bool')
  imgGropuOptions.value = await getDictFunc('imgGropu')
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
    deleteSwiperFunc(row)
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
  const res = await deleteSwiperByIds({ ids })
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
const updateSwiperFunc = async (row) => {
  const res = await findSwiper({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reswiper
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteSwiperFunc = async (row) => {
  const res = await deleteSwiper({ ID: row.ID })
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

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    swiperName: '',
    swiperPicture: '',
    isShow: false,
    group: undefined,
    goToUrl: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  let res
  switch (type.value) {
    case 'create':
      res = await createSwiper(formData.value)
      break
    case 'update':
      res = await updateSwiper(formData.value)
      break
    default:
      res = await createSwiper(formData.value)
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

// ============== 自定义部分开始 ===============

const userStore = useUserStore()
const token = ref(userStore.token)
const handleAvatarSuccess = (
  response,
  uploadFile
) => {
  formData.value.swiperPicture = response.data.file.url
  // imageUrl.value = URL.createObjectURL(uploadFile.raw)
}

const beforeAvatarUpload = (rawFile) => {
  if (rawFile.type !== 'image/jpeg') {
    ElMessage.error('Avatar picture must be JPG format!')
    return false
  } else if (rawFile.size / 1024 / 1024 > 2) {
    ElMessage.error('Avatar picture size can not exceed 2MB!')
    return false
  }
  return true
}
</script>

<style lang="scss">
.el-table .el-table__cell {
  z-index: auto;
}
.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
}
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}
</style>
