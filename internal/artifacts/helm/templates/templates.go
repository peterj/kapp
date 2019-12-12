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
{{- end -}}

{{/*
Common labels
*/}}
{{- define "[[ .ApplicationName ]].labels" -}}
helm.sh/chart: {{ include "[[ .ApplicationName ]].chart" . }}
{{ include "[[ .ApplicationName ]].selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "[[ .ApplicationName ]].selectorLabels" -}}
app.kubernetes.io/name: {{ include "[[ .ApplicationName ]].name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "[[ .ApplicationName ]].serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "[[ .ApplicationName ]].fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}
`

// DeploymentYaml holds the contents of the deployment.yaml file
const DeploymentYaml = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ template "[[ .ApplicationName ]].fullname" . }}
  labels:
    {{- include "[[ .ApplicationName ]].labels" . | nindent 4 }}
spec:
  replicas: {{ .Values.replicaCount }}
  selector:
    matchLabels:
      {{- include "[[ .ApplicationName ]].selectorLabels" . | nindent 6 }}
  template:
    metadata:
      labels:
        {{- include "[[ .ApplicationName ]].selectorLabels" . | nindent 8 }}
    spec:
    {{- with .Values.imagePullSecrets }}
      imagePullSecrets:
        {{- toYaml . | nindent 8 }}
    {{- end }}
      serviceAccountName: {{ include "[[ .ApplicationName ]].serviceAccountName" . }}
      securityContext:
        {{- toYaml .Values.podSecurityContext | nindent 8 }}
      containers:
      - name: {{ .Chart.Name }}
        securityContext:
          {{- toYaml .Values.securityContext | nindent 12 }}
        image: "{{ .Values.image.repository }}:{{ .Chart.AppVersion }}"
        imagePullPolicy: {{ .Values.image.pullPolicy }}
        ports:
        - name: http
          containerPort: 80
          protocol: TCP
        livenessProbe:
          httpGet:
            path: /health
            port: http
        readinessProbe:
          httpGet:
            path: /health
            port: http
        resources:
          {{- toYaml .Values.resources | nindent 12 }}
    {{- with .Values.nodeSelector }}
    nodeSelector:
      {{- toYaml . | nindent 8 }}
    {{- end }}
  {{- with .Values.affinity }}
    affinity:
      {{- toYaml . | nindent 8 }}
  {{- end }}
  {{- with .Values.tolerations }}
    tolerations:
      {{- toYaml . | nindent 8 }}
  {{- end }}
`

// ServiceYaml holds the contents of the service.yaml file
const ServiceYaml = `apiVersion: v1
kind: Service
metadata:
  name: {{ include "[[ .ApplicationName ]].fullname" . }}
  labels:
    {{- include "[[ .ApplicationName ]].labels" . | nindent 4 }}
spec:
  type: {{ .Values.service.type }}
  ports:
    - port: {{ .Values.service.port }}
      targetPort: http
      protocol: TCP
      name: http
  selector:
    {{- include "[[ .ApplicationName ]].selectorLabels" . | nindent 4 }}
`

// ServiceAccountYaml holds the contents of the serviceaccount.yaml file
const ServiceAccountYaml = `{{- if .Values.serviceAccount.create -}}
apiVersion: v1
kind: ServiceAccount
metadata:
  name: {{ include "[[ .ApplicationName ]].serviceAccountName" . }}
  labels:
{{ include "[[ .ApplicationName ]].labels" . | nindent 4 }}
{{- end -}}
`

// TestConnectionYaml holds the contents of the test-connection.yaml file
const TestConnectionYaml = `apiVersion: v1
kind: Pod
metadata:
  name: "{{ include "[[ .ApplicationName ]].fullname" . }}-test-connection"
  labels:
{{ include "[[ .ApplicationName ]].labels" . | nindent 4 }}
  annotations:
    "helm.sh/hook": test-success
spec:
  containers:
    - name: wget
      image: busybox
      command: ['wget']
      args:  ['{{ include "[[ .ApplicationName ]].fullname" . }}:{{ .Values.service.port }}']
  restartPolicy: Never
`

// HelmIgnore holds the contents of the .helmignore file
const HelmIgnore = `# Patterns to ignore when building packages.
# This supports shell glob matching, relative path matching, and
# negation (prefixed with !). Only one pattern per line.
.DS_Store
# Common VCS dirs
.git/
.gitignore
.bzr/
.bzrignore
.hg/
.hgignore
.svn/
# Common backup files
*.swp
*.bak
*.tmp
*~
# Various IDEs
.project
.idea/
*.tmproj
.vscode/
`

// ChartYaml holds the contents of the Chart.yaml file
const ChartYaml = `apiVersion: v2
name: [[ .ApplicationName ]]
description: A Helm chart for [[ .ApplicationName ]] service
type: application
appVersion: 0.1.0
version: 0.1.0`

// ValuesYaml holds the contents of the values.yaml file
const ValuesYaml = `replicaCount: 1
image:
  repository: [[ .DockerRepository ]]
  pullPolicy: Always

imagePullSecrets: []

serviceAccount:
    create: true
    name:

podSecurityContext: {}
securityContext: {}

service:
  type: ClusterIP
  port: 80

resources: {}
nodeSelector: {}
tolerations: []
affinity: {}
`
