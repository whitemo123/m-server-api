import { IPage } from "@/apis/core/types.ts";

// 查询参数
export interface I{{.FName}}Params extends IPage {

}

export interface I{{ .FName }} {
{{- range .Columns }}
  // {{ .Label }}
  {{- if eq .Type ""}}
  {{ .Prop }}: string;
  {{- else if eq .Type "datetimerange"}}
  {{ .Prop }}: string;
  {{- else}}
  {{ .Prop }}: number;
  {{- end }}
{{- end }}
}
