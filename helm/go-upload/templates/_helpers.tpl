{{- define "go-upload.name" -}}
go-upload
{{- end }}

{{- define "go-upload.fullname" -}}
{{ .Release.Name }}-{{ include "go-upload.name" . }}
{{- end }}
