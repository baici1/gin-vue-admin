<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="轮播图名称:">
          <el-input v-model="formData.swiperName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="轮播图照片:">
          <el-input v-model="formData.swiperPicture" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否展示:">
          <el-switch v-model="formData.isShow" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="分组:">
          <el-select v-model="formData.group" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in imgGropuOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="前往路径:">
          <el-input v-model="formData.goToUrl" clearable placeholder="请输入" />
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
  name: 'Swiper'
}
</script>

<script setup>
import {
  createSwiper,
  updateSwiper,
  findSwiper
} from '@/api/swiper'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const boolOptions = ref([])
const imgGropuOptions = ref([])
const formData = ref({
        swiperName: '',
        swiperPicture: '',
        isShow: false,
        group: undefined,
        goToUrl: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSwiper({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reswiper
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    boolOptions.value = await getDictFunc('bool')
    imgGropuOptions.value = await getDictFunc('imgGropu')
}

init()
// 保存按钮
const save = async() => {
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
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
