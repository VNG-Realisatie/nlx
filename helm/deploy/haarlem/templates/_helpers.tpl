{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "haarlem.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "haarlem.fullname" -}}
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
{{- define "haarlem.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "haarlem.labels" -}}
helm.sh/chart: {{ include "haarlem.chart" . }}
{{ include "haarlem.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "haarlem.selectorLabels" -}}
app.kubernetes.io/name: {{ include "haarlem.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "haarlem.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "haarlem.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Return the name of the nlxctl image
*/}}
{{- define "haarlem.nlxctl.image" -}}
{{- $registryName := default .Values.nlxctl.image.registry .Values.global.imageRegistry -}}
{{- $repositoryName := .Values.nlxctl.image.repository -}}
{{- $tag := default (printf "v%s" .Chart.AppVersion) (default .Values.nlxctl.image.tag .Values.global.imageTag) -}}

{{- printf "%s/%s:%s" $registryName $repositoryName $tag -}}
{{- end -}}

{{/*
Return the image name for transaction log database job
*/}}
{{- define "haarlem.transactionLog.image" -}}
{{- $registryName := default .Values.transactionLog.image.registry .Values.global.imageRegistry -}}
{{- $repositoryName := .Values.transactionLog.image.repository -}}
{{- $tag := default (printf "v%s" .Chart.AppVersion) (default .Values.transactionLog.image.tag .Values.global.imageTag) -}}

{{- printf "%s/%s:%s" $registryName $repositoryName $tag -}}
{{- end -}}

{{/*
Return the secret name of the transaction log database
*/}}
{{- define "haarlem.transactionLog.secret" -}}
{{- default (printf "%s-postgresql" (include "haarlem.fullname" .)) .Values.postgresql.existingSecret -}}
{{- end -}}
