acl {{ .Name }} {
{{- range .Acls}}
{{ .CIDR }};
{{- end }}
};