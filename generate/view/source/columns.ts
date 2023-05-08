import {h, reactive} from "vue";
import {NButton, NTag} from "naive-ui";

//class

type actionFun = (editAction: Permission) => void

const GetColumns = (editConfig: actionFun, deleteConfig: actionFun, detailConfig: actionFun) => {
    return [
//showField
        {
            title: 'Action',
            key: 'actions',
            render(row: Permission) {
                return [
                    h(
                        NButton,
                        {
                            tertiary: true,
                            size: 'small',
                            type: 'info',
                            style: "margin:5px",
                            onClick: () => editConfig(row)
                        },
                        {default: () => "编辑"}
                    ),
                    h(
                        NButton,
                        {
                            tertiary: true,
                            size: 'small',
                            type: 'error',
                            style: "margin:5px",
                            onClick: () => deleteConfig(row)
                        },
                        {default: () => "删除"}
                    ),
                    h(
                        NButton,
                        {
                            tertiary: true,
                            size: 'small',
                            type: 'default',
                            style: "margin:5px",
                            onClick: () => detailConfig(row)
                        },
                        {default: () => "详情"}
                    )
                ];
            }
        }
    ]
}


const rules = reactive({
//rules
})
export {GetColumns, rules}
export {Permission}
