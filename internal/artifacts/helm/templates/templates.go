package templates

// HelpersTpl holds contents of the _helpers.tpl file
const HelpersTpl = `
{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "[[ .ApplicationName ]].name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "[[ .ApplicationName ]].fullname" -}}
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
{{- define "[[ .ApplicationName ]].chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}`

// DeploymentYaml holds the contents of the deployment.yaml file
const DeploymentYaml = `apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: {{ template "[[ .ApplicationName ]].fullname" . }}
  labels:
    app: {{ template "[[ .ApplicationName ]].name" . }}
    chart: {{ template "[[ .ApplicationName ]].chart" . }}
    version: {{ .Release.Name }}
    heritage: {{ .Release.Service }}
spec:
  replicas: {{ .Values.replicaCount }}
  selector:
    matchLabels:
      app: {{ template "[[ .ApplicationName ]].name" . }}
      release: {{ .Release.Name }}
  template:
    metadata:
      labels:
        app: {{ template "[[ .ApplicationName ]].name" . }}
        release: {{ .Release.Name }}
    spec:
      containers:
      - name: {{ .Chart.Name }}
        image: "{{ .Values.image.repository }}:{{ .Values.image.tag }}"
        imagePullPolicy: {{ .Values.image.pullPolicy }}
        ports:
        - name: http
          containerPort: 8080
{{- if .Values.image.imagePullSecret }}
      imagePullSecrets:
      - name: {{ .Values.image.imagePullSecret }}
{{- end }}`

// ServiceYaml holds the contents of the service.yaml file
const ServiceYaml = `apiVersion: v1
kind: Service
metadata:
  name: {{ template "[[ .ApplicationName ]].fullname" . }}
  labels:
    app: {{ template "[[ .ApplicationName ]].name" . }}
    chart: {{ template "[[ .ApplicationName ]].chart" . }}
    release: {{ .Release.Name }}
    heritage: {{ .Release.Service }}
spec:
  type: {{ .Values.service.type }}
  ports:
    - port: {{ .Values.service.port }}
      targetPort: http
      protocol: TCP
      name: http
  selector:
    app: {{ template "[[ .ApplicationName ]].name" . }}
    release: {{ .Release.Name }}`

// ChartYaml holds the contents of the Chart.yaml file
const ChartYaml = `apiVersion: v1
appVersion: "1.0"
description: A Helm chart for [[ .ApplicationName ]] service
name: [[ .ApplicationName ]]
version: 0.1.0`

// ValuesYaml holds the contents of the values.yaml file
const ValuesYaml = `replicaCount: 1
image:
  repository: [[ .DockerRepository ]]
  tag: 0.1.0
  pullPolicy: Always
  imagePullSecret: ""

service:
  type: ClusterIP
  port: 80`
