$ORIGIN    {{ .Domain.Name }}.
$TTL       {{ .Domain.TTL }}
@          86400        IN        SOA        {{ .Domain.Provider }} {{ .Domain.Contact }} {{ .Domain.Serial }} {{ .Domain.Refresh }} {{ .Domain.Retry }} {{ .Domain.Expire }} {{ .Domain.TTL }}
@          86400        IN        NS         ns1.{{ .Domain.Name }}.
@          86400        IN        NS         ns2.{{ .Domain.Name }}.
;records
{{- range .Records}}
; {{.ID}}. {{.Comment}}
{{- if .Disabled }}
; {{.Name}} {{.TTL}} IN {{.Type}} {{.Content}}
{{- else }}
{{.Name}} {{.TTL}} IN {{.Type}} {{.Content}}
{{- end }}
{{- end }}