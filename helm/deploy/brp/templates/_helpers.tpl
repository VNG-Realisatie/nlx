{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "brp.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "brp.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "brp.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "brp.labels" -}}
helm.sh/chart: {{ include "brp.chart" . }}
{{ include "brp.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "brp.selectorLabels" -}}
app.kubernetes.io/name: {{ include "brp.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "brp.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "brp.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Return the name of the nlxctl image
*/}}
{{- define "brp.nlxctl.image" -}}
{{- $registryName := default .Values.nlxctl.image.registry .Values.global.imageRegistry -}}
{{- $repositoryName := .Values.nlxctl.image.repository -}}
{{- $tag := default (printf "v%s" .Chart.AppVersion) (default .Values.nlxctl.image.tag .Values.global.imageTag) -}}

{{- printf "%s/%s:%s" $registryName $repositoryName $tag -}}
{{- end -}}

{{/*
Return the image name for the unsafe CA
*/}}
{{- define "brp.unsafeCA.image" -}}
{{- $registryName := default .Values.unsafeCA.image.registry .Values.global.imageRegistry -}}
{{- $repositoryName := .Values.unsafeCA.image.repository -}}
{{- $tag := default (printf "v%s" .Chart.AppVersion) (default .Values.unsafeCA.image.tag .Values.global.imageTag) -}}

{{- printf "%s/%s:%s" $registryName $repositoryName $tag -}}
{{- end -}}

{{/*
Return the image name for transaction log database job
*/}}
{{- define "brp.transactionLog.image" -}}
{{- $registryName := default .Values.transactionLog.image.registry .Values.global.imageRegistry -}}
{{- $repositoryName := .Values.transactionLog.image.repository -}}
{{- $tag := default (printf "v%s" .Chart.AppVersion) (default .Values.transactionLog.image.tag .Values.global.imageTag) -}}

{{- printf "%s/%s:%s" $registryName $repositoryName $tag -}}
{{- end -}}

{{/*
Return the secret name of the transaction log database
*/}}
{{- define "brp.transactionLog.secret" -}}
{{- default (printf "%s-postgresql" (include "brp.fullname" .)) .Values.postgresql.existingSecret -}}
{{- end -}}
