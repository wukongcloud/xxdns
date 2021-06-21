$ORIGIN {{ .Domain.Name }}.
$TTL    {{ .Domain.TTL }}
@     86400 IN SOA    {{ .Domain.Provider }} {{ .Domain.Contact }} {{ .Domain.Serial }} {{ .Domain.Refresh }} {{ .Domain.Retry }} {{ .Domain.Expire }} {{ .Domain.TTL }}
@     86400 IN NS     ns1.{{ .Domain.Name }}.
;@     86400 IN NS     ns2.{{ .Domain.Name }}.
ns1   86400 IN A      60.29.244.151
;ns2   86400 IN A      172.16.42.59

;records
{{- range .Records}}
; {{.ID}}. {{.Comment}}
{{- if .Disabled }}
; {{.Name}} {{.TTL}} IN {{.Type}} {{.Content}}
{{- else }}
{{.Name}} {{.TTL}} IN {{.Type}} {{.Content}}
{{- end }}
{{- end }}