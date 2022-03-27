<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="是否可评论:">
          <el-switch v-model="formData.commentabled" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="是否发表:">
          <el-switch v-model="formData.published" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="发布时间:">
          <el-date-picker v-model="formData.publishedTime" type="date" placeholder="选择日期" clearable></el-date-picker>
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
          <el-select v-model="formData.type" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in articleTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="置顶:">
          <el-switch v-model="formData.orderNum" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  name: 'Article'
}
</script>

<script setup>
import {
  createArticle,
  updateArticle,
  findArticle
} from '@/api/article'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const boolOptions = ref([])
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findArticle({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rearticle
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    boolOptions.value = await getDictFunc('bool')
    articleTypeOptions.value = await getDictFunc('articleType')
}

init()
// 保存按钮
const save = async() => {
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
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
