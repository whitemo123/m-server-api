<script lang="ts" setup>
import { ref } from 'vue'
// 接口文件
import {
  {{ .Name }}PageListApi,
  {{ .Name }}CreateApi,
  {{ .Name }}ModifyApi,
  {{ .Name }}DetailApi,
  {{ .Name }}DelApi,
  {{ .Name }}ExportApi
} from '@/apis/{{ .Folder }}/{{ .Name }}'
// 类型文件
import type {
  I{{ .FName }},
  I{{ .FName }}Params
} from '@/apis/{{ .Folder }}/{{ .Name }}/types.ts'

import { CrudInstance, ICrudOption } from "m-eleplus-crud";
import { deepClone } from "@/utils/util.ts";
import { ElMessage } from "element-plus";

// crud ref
const crudRef = ref<CrudInstance>()
// 查询条件
const query = ref<I{{ .FName }}Params>({
  page: 1,
  limit: 10
})
// 表单数据
const modelForm = ref<Partial<I{{ .FName }}>>({})
// 列表数据
const tableData = ref<I{{ .FName }}[]>([])
// 列表总条数
const tableTotal = ref(0)
// 列表加载状态
const tableLoading = ref(false)

// crud配置
const crudOption: ICrudOption = {
  height: "auto",
  column: [
    {{- range .Columns }}
    {
      label: "{{ .Label }}",
      prop: "{{ .Prop }}",
      align: "center",
      {{- if eq .Required true }}
      formRules: [
        { required: true, message: "{{ .Label }}不能为空", trigger: "blur" }
      ],
      {{- end }}
      {{- if ne .Default "" }}
      formDefault: "{{ .Default }}",
      {{- end }}
      {{- if eq .Search true }}
      search: true,
      {{- end }}
      {{- if ne .Table true }}
      hide: true,
      {{- end }}
      {{- if ne .Edit true }}
      editHide: true,
      {{- end }}
      {{- if ne .Add true }}
      addHide: true,
      {{- end }}
    },
    {{- end }}
  ]
}

// 查询列表
const getList = () => {
  tableLoading.value = true

  const params: I{{ .FName }}Params = deepClone(query.value)
 
  {{ .Name }}PageListApi(params).then(e => {
    if (e && e.data) {
      tableTotal.value = e.data.total
      tableData.value = e.data.list
      tableLoading.value = false
    }
  })
}

// 保存
const rowSave = (form: I{{ .FName }}, done: () => void, loading: () => void) => {
  {{ .Name }}CreateApi(form).then(e => {
    if (e && e.data) {
      ElMessage.success("操作成功!")
      getList()
      done()
    }
  }).finally(() => {
    setTimeout(() => {
      loading()
    }, 500)
  })
}

// 打开编辑
const openEdit = (row: I{{ .FName }}, index: number) => {
  {{ .Name }}DetailApi(row.id).then(e => {
    if (e && e.data) {
      crudRef.value?.rowEdit(e.data, index)
    }
  })
}

// 编辑
const rowEdit = (form: I{{ .FName }}, done: () => void, loading: () => void) => {
  {{ .Name }}ModifyApi(form).then(e => {
    if (e && e.data) {
      ElMessage.success("操作成功!")
      getList()
      done()
    }
  }).finally(() => {
    setTimeout(() => {
      loading()
    }, 500)
  })
}

// 删除
const rowDel = (row: I{{ .FName }}) => {
  {{ .Name }}DelApi(row.id).then(e => {
    if (e) {
      ElMessage.success("操作成功!")
      getList()
    }
  })
}

getList()
</script>

<template>
  <page>
    <m-crud
        ref="crudRef"
        v-model="modelForm"
        v-model:search="query"
        :option="crudOption"
        :data="tableData"
        :total="tableTotal"
        :loading="tableLoading"
        @row-save="rowSave"
        @row-edit="rowEdit"
        @row-del="rowDel"
        @search="getList"
        @reset="getList"
    >
      <template #editBtn="{row, $index}">
        <el-link
            class="m-control-btns"
            type="primary"
            :underline="false"
            icon="Edit"
            @click="openEdit(row, $index)"
        >
          编辑
        </el-link>
      </template>
    </m-crud>
  </page>
</template>
