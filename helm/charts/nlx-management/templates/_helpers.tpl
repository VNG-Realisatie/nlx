{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "nlx-management.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "nlx-management.fullname" -}}
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
{{- define "nlx-management.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "nlx-management.labels" -}}
helm.sh/chart: {{ include "nlx-management.chart" . }}
{{ include "nlx-management.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "nlx-management.selectorLabels" -}}
app.kubernetes.io/name: {{ include "nlx-management.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "nlx-management.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "nlx-management.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Return the image name for the API
*/}}
{{- define "nlx-management.apiImage" -}}
{{- $registryName := default .Values.image.registry .Values.global.imageRegistry -}}
{{- $repositoryName := .Values.image.apiRepository -}}
{{- $tag := default (printf "v%s" .Chart.AppVersion) (default .Values.image.tag .Values.global.imageTag) -}}

{{- printf "%s/%s:%s" $registryName $repositoryName $tag -}}
{{- end -}}

{{/*
Return the image name for the UI
*/}}
{{- define "nlx-management.uiImage" -}}
{{- $registryName := default .Values.image.registry .Values.global.imageRegistry -}}
{{- $repositoryName := .Values.image.uiRepository -}}
{{- $tag := default (printf "v%s" .Chart.AppVersion) (default .Values.image.tag .Values.global.imageTag) -}}

{{- printf "%s/%s:%s" $registryName $repositoryName $tag -}}
{{- end -}}

{{/*
Return the image name for the UI
*/}}
{{- define "nlx-management.configImage" -}}
{{- $registryName := default .Values.image.registry .Values.global.imageRegistry -}}
{{- $repositoryName := .Values.image.configRepository -}}
{{- $tag := default (printf "v%s" .Chart.AppVersion) (default .Values.image.tag .Values.global.imageTag) -}}

{{- printf "%s/%s:%s" $registryName $repositoryName $tag -}}
{{- end -}}

{{/*
Return the image name for the unsafe ca
*/}}
{{- define "nlx-management.unsafeCA.image" -}}
{{- $registryName := default .Values.unsafeCA.image.registry .Values.global.imageRegistry -}}
{{- $repositoryName := .Values.unsafeCA.image.repository -}}
{{- $tag := default (printf "v%s" .Chart.AppVersion) (default .Values.unsafeCA.image.tag .Values.global.imageTag) -}}

{{- printf "%s/%s:%s" $registryName $repositoryName $tag -}}
{{- end -}}

{{/*
Return the address of the CFSSL unsafe CA
*/}}
{{- define "nlx-management.unsafeCA.cfsslHostname" -}}
{{- default .Values.unsafeCA.cfsslHostname .Values.global.unsafeCA.cfsslHostname -}}
{{- end -}}

{{/*
Return the organization name for the certificated to be genereate by the CFSSL unsafe CA
*/}}
{{- define "nlx-management.unsafeCA.organizationName" -}}
{{- default .Values.unsafeCA.organizationName .Values.global.unsafeCA.organizationName -}}
{{- end -}}
