{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "nlx-outway.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "nlx-outway.fullname" -}}
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
{{- define "nlx-outway.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "nlx-outway.labels" -}}
helm.sh/chart: {{ include "nlx-outway.chart" . }}
{{ include "nlx-outway.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "nlx-outway.selectorLabels" -}}
app.kubernetes.io/name: {{ include "nlx-outway.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "nlx-outway.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "nlx-outway.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Return the image name for the NLX outway
*/}}
{{- define "nlx-outway.image" -}}
{{- $registryName := default .Values.image.registry .Values.global.imageRegistry -}}
{{- $repositoryName := .Values.image.repository -}}
{{- $tag := default (printf "v%s" .Chart.AppVersion) (default .Values.image.tag .Values.global.imageTag) -}}

{{- printf "%s/%s:%s" $registryName $repositoryName $tag -}}
{{- end -}}

{{/*
Return the secret name of the PostgreSQL username/password
*/}}
{{- define "nlx-outway.transactionLog.secret" -}}
{{- default (printf "%s-postgresql" (include "nlx-outway.fullname" .)) .Values.transactionLog.existingSecret -}}
{{- end -}}

{{/*
Return the image name for the unsafe ca
*/}}
{{- define "nlx-outway.unsafeCA.image" -}}
{{- $registryName := default .Values.unsafeCA.image.registry .Values.global.imageRegistry -}}
{{- $repositoryName := .Values.unsafeCA.image.repository -}}
{{- $tag := default (printf "v%s" .Chart.AppVersion) (default .Values.unsafeCA.image.tag .Values.global.imageTag) -}}

{{- printf "%s/%s:%s" $registryName $repositoryName $tag -}}
{{- end -}}

{{/*
Return the address of the CFSSL unsafe CA
*/}}
{{- define "nlx-outway.unsafeCA.cfsslHostname" -}}
{{- default .Values.unsafeCA.cfsslHostname .Values.global.unsafeCA.cfsslHostname -}}
{{- end -}}

{{/*
Return the organization name for the certificated to be genereate by the CFSSL unsafe CA
*/}}
{{- define "nlx-outway.unsafeCA.organizationName" -}}
{{- default .Values.unsafeCA.organizationName .Values.global.unsafeCA.organizationName -}}
{{- end -}}
