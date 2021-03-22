{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "saas-org-x.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "saas-org-x.fullname" -}}
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
{{- define "saas-org-x.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "saas-org-x.labels" -}}
helm.sh/chart: {{ include "saas-org-x.chart" . }}
{{ include "saas-org-x.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "saas-org-x.selectorLabels" -}}
app.kubernetes.io/name: {{ include "saas-org-x.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "saas-org-x.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "saas-org-x.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Return the name of the nlx-management-api image
*/}}
{{- define "saas-org-x.managementAPI.image" -}}
{{- $registryName := default .Values.managementAPI.image.registry .Values.global.imageRegistry -}}
{{- $repositoryName := .Values.managementAPI.image.repository -}}
{{- $tag := default (printf "v%s" .Chart.AppVersion) (default .Values.managementAPI.image.tag .Values.global.imageTag) -}}

{{- printf "%s/%s:%s" $registryName $repositoryName $tag -}}
{{- end -}}

{{/*
Return the image name for transaction log database job
*/}}
{{- define "saas-org-x.transactionLog.image" -}}
{{- $registryName := default .Values.transactionLog.image.registry .Values.global.imageRegistry -}}
{{- $repositoryName := .Values.transactionLog.image.repository -}}
{{- $tag := default (printf "v%s" .Chart.AppVersion) (default .Values.transactionLog.image.tag .Values.global.imageTag) -}}

{{- printf "%s/%s:%s" $registryName $repositoryName $tag -}}
{{- end -}}

{{/*
Return the secret name of the transaction log database
*/}}
{{- define "saas-org-x.transactionLog.secret" -}}
{{- default (printf "%s-postgresql" (include "saas-org-x.fullname" .)) .Values.postgresql.existingSecret -}}
{{- end -}}
