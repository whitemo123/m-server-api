import api from "@/request"
import { IPageRes } from "@/apis/core/types.ts"
import { I{{.FName}}Params, I{{.FName}} } from "./types"

/**
 * 分页列表
 * @param params 查询参数
 */
export const {{.Name}}PageListApi = (params: I{{.FName}}Params) => api.get<IPageRes<I{{.FName}}>>("/admin/{{.Name}}/page", params)

/**
 * 列表
 */
export const {{.Name}}ListApi = () => api.get<I{{.FName}}[]>("/admin/{{.Name}}/list")

/**
 * 创建
 * @param data 请求数据
 */
export const {{.Name}}CreateApi = (data: Partial<I{{.FName}}>) => api.postJ("/admin/{{.Name}}/create", data)

/**
 * 编辑
 * @param data 请求数据
 */
export const {{.Name}}ModifyApi = (data: Partial<I{{.FName}}>) => api.postJ("/admin/{{.Name}}/modify", data)

/**
 * 详情
 * @param id 主键ID
 */
export const {{.Name}}DetailApi = (id: string) => api.get<I{{.FName}}>("/admin/{{.Name}}/detail/" + id)

/**
 * 删除
 * @param id 主键ID
 */
export const {{.Name}}DelApi = (id: string) => api.get("/admin/{{.Name}}/del/" + id)

/**
 * 导出
 * @param params 查询参数
 */
export const {{.Name}}ExportApi = (params: I{{.FName}}Params) => api.get("/admin/{{.Name}}/export", params)
