<template>
    <n-spin :show="globalLoading" description="working...">
        <n-row gutter="12">
            <n-col :span="24" style="text-align: left;margin-bottom: 10px;margin-left: -10px">
                <n-button type="info" secondary @click="addPermission" style="margin-left: 10px;width: 100px">
                    新增@target@
                </n-button>
            </n-col>
        </n-row>
        <n-row gutter="12">
            <n-col :span="24" style="color: white;">
                <n-data-table
                        remote
                        ref="table"
                        :columns="columns"
                        :data="dataList"
                        :loading="dataLoading"
                        :pagination="pagination"
                        :bordered="false"
                        @update:page="getDataList"
                />
            </n-col>
        </n-row>
        <!--  新增 111-->
        <n-modal v-model:show="showAddModal" preset="dialog" title="Dialog" style="width: 760px"
                 :mask-closable="false"
                 transform-origin="center">
            <template #header>
                <div>{{ formTitle }}</div>
            </template>
            <n-spin :show="loading" description="working...">
                <div style="margin-top: 20px;">
                    <n-form ref="formRef" :show-label="false" :model="permissionForm" :rules="rules" size="large">
<!--addField-->

                    </n-form>
                </div>
            </n-spin>
            <template #action>
                <n-button type="primary" :disabled="disable" @click="actionFun" secondary class="action-button"
                          style="width: 100px">
                    确定
                </n-button>
                <n-button @click="cancel" :disabled="disable" secondary class="action-button" style="width: 100px">取 消
                </n-button>
            </template>
        </n-modal>
        <n-modal v-model:show="showDetailModal" preset="dialog" title="Dialog" style="width: 760px"
                 :mask-closable="false"
                 transform-origin="center">
            <template #header>
                <div>{{ formTitle }}</div>
            </template>
            <div style="margin-top: 20px;">
                <n-descriptions label-placement="top" bordered :column="3" size="large">
<!--showField-->
                </n-descriptions>
            </div>
            <template #action>
                <n-button type="primary" @click="cancel" secondary class="action-button"
                          style="width: 100px">
                    确定
                </n-button>
            </template>
        </n-modal>
    </n-spin>
</template>

<script lang="ts">
import {PermissionAdd, PermissionDelete, PermissionList, PermissionSave} from '@/api/services/PermissionService'
import {defineComponent, reactive, ref} from 'vue'
import {GetColumns, Permission, rules} from '@/element/columns/Permission'
import {DialogApi, FormInst, MessageApi, useDialog, useMessage} from "naive-ui";
import {validation} from "@/api/models";
import {paginationReactive} from '@/utils/attribute'

const formRef = ref<FormInst | null>(null)
let permissionForm = reactive(new Permission())

let actionFun = ref()
let dataList = ref()
let loading = ref(false)
let showAddModal = ref(false)
let formTitle = ref("新建@target@配置")
let dialog: DialogApi
let message: MessageApi
let showDetailModal = ref(false)
let disable = ref(false)
const dataLoading = ref(false)
const globalLoading = ref(false)


// 添加基础@target@
const addPermission = () => {
    showAddModal.value = true
    formTitle.value = "新建@target@配置"
    // 重置
    Object.assign(permissionForm, new Permission())
    actionFun.value = addPermissionAction
    console.log(actionFun.value)
}

// 添加@target@配置-操作
const addPermissionAction = () => {
    formRef.value?.validate((errors) => {
        if (errors) {
            console.log("errors:", errors)
        } else {
            disable.value = true
            loading.value = true
            let params = validation.PermissionAddParams.createFrom(permissionForm)
            PermissionAdd(params).then(() => {
                message.success("操作成功!")
                getDataList(1)
                showAddModal.value = false
                loading.value = false
                disable.value = false
            }).catch((err: any) => {
                message.error("操作失败")
                showAddModal.value = false
                loading.value = false
                disable.value = false
            })
        }
    })
}

const editAction = (row: Permission) => {
    showAddModal.value = true
    formTitle.value = "编辑@target@配置"
    Object.assign(permissionForm, row)
    actionFun.value = savePermissionAction
}
// 保存@target@配置操作
const savePermissionAction = () => {
    formRef.value?.validate((errors) => {
        if (errors) {
            console.log("errors:", errors)
        } else {
            disable.value = true
            loading.value = true
            let params = validation.PermissionSaveParams.createFrom(permissionForm)
            PermissionSave(params).then(() => {
                message.success("操作成功!")
                loading.value = false
                showAddModal.value = false
                disable.value = false
                getDataList(1)
            }).catch((err: any) => {
                message.error("操作失败")
                loading.value = false
                showAddModal.value = false
                disable.value = false
            })
        }
    })
}

const deleteAction = (row: Permission) => {
    dialog.warning({
        title: '警告',
        content: '确定删除当前配置吗？',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: () => {
            globalLoading.value=true
            let params = new validation.PermissionDeleteParams()
            params.Id = row.Id
            PermissionDelete(params).then(response => {
                message.success("操作成功!")
                getDataList(1)
                globalLoading.value=false
            }).catch(err => {
                globalLoading.value=false
                message.error("操作失败")
            })
        },
        onNegativeClick: () => {
        }
    })
}

const getDataList = (page: number) => {
    if (dataLoading.value) {
        return
    }
    dataLoading.value = true
    let params = new (validation.PermissionListParams)
    params.PageIndex = page
    PermissionList(params).then((response: any) => {
        dataList.value = response.list
        paginationReactive.page = response.pageIndex
        paginationReactive.itemCount = response.count
        console.log(response)
        dataLoading.value = false
    }).catch((err: any) => {
        console.log(err)
        dataLoading.value = false
    })
}

const detailAction = (row: Permission) => {
    console.log("detail:", row)
    showDetailModal.value = true
    // 重置
    Object.assign(permissionForm, row)
    console.log(permissionForm.Id)
}

const cancel = (): void => {
    showAddModal.value = false
    showDetailModal.value = false
}


export default defineComponent({
    setup() {
        getDataList(1)
        actionFun.value = addPermissionAction
        dialog = useDialog()
        message = useMessage()
        return {
            dataList,
            disable,
            actionFun,
            formRef,
            formTitle,
            cancel,
            showDetailModal,
            rules,
            showAddModal,
            getDataList,
            loading,
            globalLoading,
            dataLoading,
            permissionForm,
            addPermission,
            columns: GetColumns(editAction, deleteAction, detailAction),
            pagination: paginationReactive
        }
    }
})

</script>