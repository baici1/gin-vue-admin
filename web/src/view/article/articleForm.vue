<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="100px">
        <el-form-item label="标题:">
          <el-input v-model="formData.title" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="引用:">
          <el-input
            v-model="formData.description"
            clearable
            placeholder="请输入"
            maxlength="150"
            show-word-limit
            autosize
            type="textarea"
          />
        </el-form-item>

        <el-form-item label="浏览量:">
          <el-input-number v-model="formData.views" :step="1" />
        </el-form-item>
        <el-form-item label="作者:">
          <el-input v-model="formData.author" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="文章类型:">
          <el-select v-model="formData.type" placeholder="请选择" clearable>
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
        <el-form-item label="是否发布:">
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
            placeholder="选择日期"
            clearable
          />
        </el-form-item>
        <div
          v-if="isEditorShow"
          v-loading="!isEditorShow"
          style="border: 1px solid #ccc"
          class="mb-10"
        >
          <Toolbar
            :editor-id="editorId"
            :default-config="toolbarConfig"
            mode="default"
            style="border-bottom: 1px solid #ccc"
          />
          <Editor
            :editor-id="editorId"
            :default-config="editorConfig"
            :default-html="defaultHtml"
            mode="default"
            style="height: 500px; overflow-y: hidden"
          />
          <!-- 注意: defaultContent (JSON 格式) 和 defaultHtml (HTML 格式) ，二选一 -->
        </div>
        <div v-else>loading</div>
        <!-- <el-form-item label="内容:">
          <el-input v-model="formData.content" clearable placeholder="请输入" />
        </el-form-item>-->
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
  name: 'Article',
}
</script>

<script setup>
import {
  createArticle,
  updateArticle,
  findArticle,
  uploadFile,
} from '@/api/article'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, onBeforeUnmount } from 'vue'
import {
  Editor,
  Toolbar,
  getEditor,
  removeEditor,
} from '@wangeditor/editor-for-vue'
import { useUserStore } from '../../pinia/modules/user'

const route = useRoute()
const router = useRouter()
const type = ref('')
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

// 存在自定义
const isEditorShow = ref(false)
const defaultHtml = ref('')
// 初始化方法
const init = async () => {
  isEditorShow.value = false
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findArticle({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.rearticle
      type.value = 'update'
      defaultHtml.value = '<p>' + formData.value.content + '<p>'
    }
  } else {
    type.value = 'create'
  }
  isEditorShow.value = true
  articleTypeOptions.value = await getDictFunc('articleType')
}

init()
// 保存按钮
const save = async () => {
  const editor = getEditor(editorId) // 获取 editor ，必须等待它渲染完之后
  formData.value.content = editor.getHtml()
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
      message: '创建/更改成功',
    })
  }
}

// 返回按钮
const back = () => {
  router.go(-1)
}
// ============================================ 自定义部分开始=========================================
const userStore = useUserStore()
const editorId = `w-e-${Math.random().toString().slice(-5)}` // 【注意】编辑器 id ，要全局唯一

// defaultContent (JSON 格式) 和 defaultHtml (HTML 格式) ，二选一

// const defaultContent = [
//   { type: 'paragraph', children: [{ text: '一行文字' }] }
// ]
// 工具栏配置
const toolbarConfig = {
  excludeKeys: ['fullScreen'],
}
// 自定义校验链接
function customCheckLinkFn(text, url) {
  if (!url) {
    return
  }
  if (url.indexOf('http') !== 0) {
    return '链接必须以 http/https 开头'
  }
  return true

  // 返回值有三种选择：
  // 1. 返回 true ，说明检查通过，编辑器将正常插入链接
  // 2. 返回一个字符串，说明检查未通过，编辑器会阻止插入。会 alert 出错误信息（即返回的字符串）
  // 3. 返回 undefined（即没有任何返回），说明检查未通过，编辑器会阻止插入。但不会提示任何信息
}

const editorConfig = {
  autoFocus: false,
  // 自定义alert
  customAlert: (s, t) => {
    console.log(
      '%c 🍕 t: ',
      'font-size:20px;background-color: #33A5FF;color:#fff;',
      t
    )
    console.log(
      '%c 🥥 s: ',
      'font-size:20px;background-color: #FCA650;color:#fff;',
      s
    )
    switch (t) {
      case 'success':
        ElMessage.success(s)
        break
      case 'info':
        ElMessage.info(s)
        break
      case 'warning':
        ElMessage.warning(s)
        break
      case 'error':
        ElMessage.error(s)
        break
      default:
        ElMessage.info(s)
        break
    }
  },
  MENU_CONF: {
    placeholder: '请输入内容...',

    // 插入链接
    insertLink: {
      checkLink: customCheckLinkFn, // 也支持 async 函数
    },
    // 更新链接
    editLink: {
      checkLink: customCheckLinkFn, // 也支持 async 函数
    },
    insertImage: {
      onInsertedImage(imageNode) {
        if (imageNode == null) return

        const { src, alt, url, href } = imageNode
        console.log('inserted image', src, alt, url, href)
      },
    },
    uploadImage: {
      // 将 meta 拼接到 url 参数中，默认 false
      metaWithUrl: false,
      server: 'http://127.0.0.1:8888/fileUploadAndDownload/upload',
      fieldName: 'file',
      // 自定义增加 http  header
      headers: {
        xToken: userStore.token,
      },
      customInsert(res, insertFn) {
        // res 即服务端的返回结果
        // 从 res 中找到 url alt href ，然后插图图片
        insertFn(res.data.file.url, res.data.file.name)
      },
      customUpload: async (file, insertFn) => {
        const param = new FormData() // 创建form对象
        param.append('file', file) // 通过append向form对象添加数据
        // file 即选中的文件
        // 自己实现上传，并得到图片 url alt href
        const data = await uploadFile(param)
        const url = data.data.file.url
        const alt = data.data.file.name
        // 最后插入图片
        insertFn(url, alt)
      },
    },
  },
}

// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = getEditor(editorId)
  if (editor == null) return

  editor.destroy()
  removeEditor(editorId)
})
</script>
<style lang="scss" scoped>
.w-e-text-container {
  p {
    strong {
      font-weight: bolder !important;
    }
    b {
      font-weight: bold !important;
    }
    i {
      font-style: italic !important;
    }
    em {
      font-style: italic !important;
    }
  }
}
h2 {
  font-weight: bold !important;
  font-size: 1.5em !important;
}
</style>
<style src="@wangeditor/editor/dist/css/style.css"></style>
