---
title: "Seccomp profile is not configured"
group_id: "rules/k8s"
meta:
  name: "k8s/seccomp_profile_is_not_configured"
  id: "f377b83e-bd07-4f48-a591-60c82b14a78b"
  display_name: "Seccomp profile is not configured"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `f377b83e-bd07-4f48-a591-60c82b14a78b`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Configurations

### Description

 Containers should be configured with a secure seccomp profile to restrict potentially dangerous syscalls.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: pod-test-1
  annotations:
    seccomp.security.alpha.kubernetes.io/defaultProfileName:  'runtime/default'
spec:
  containers:
  - name: foobar
    image: foo/bar:latest
```

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: securitydemo
  labels:
    app: web
spec:
  replicas: 2
  selector:
    matchLabels:
      app: web
  template:
    metadata:
      labels:
        app: web
    spec:
      securityContext:
        runAsUser: 19000
        seccompProfile:
            type: RuntimeDefault
      containers:
        - name: frontend
          image: nginx
          ports:
            - containerPort: 80
          securityContext:
            allowPrivilegeEscalation: false
        - name: echoserver
          image: k8s.gcr.io/echoserver:1.4
          ports:
            - containerPort: 8080
          securityContext:
            allowPrivilegeEscalation: false

```

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: securitydemo
  labels:
    app: web
spec:
  replicas: 2
  selector:
    matchLabels:
      app: web
  template:
    metadata:
      labels:
        app: web
    spec:
      securityContext:
        runAsUser: 19000
      containers:
        - name: frontend
          image: nginx
          ports:
            - containerPort: 80
          securityContext:
            allowPrivilegeEscalation: false
            seccompProfile:
                type: RuntimeDefault
        - name: echoserver
          image: k8s.gcr.io/echoserver:1.4
          ports:
            - containerPort: 8080
          securityContext:
            allowPrivilegeEscalation: false
            seccompProfile:
                type: RuntimeDefault

```
## Non-Compliant Code Examples
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: securitydemo
  labels:
    app: web
spec:
  replicas: 2
  selector:
    matchLabels:
      app: web
  template:
    metadata:
      labels:
        app: web
    spec:
      securityContext:
        runAsUser: 19000
      containers:
        - name: frontend
          image: nginx
          ports:
            - containerPort: 80
          securityContext:
            allowPrivilegeEscalation: false
        - name: echoserver
          image: k8s.gcr.io/echoserver:1.4
          ports:
            - containerPort: 8080
          securityContext:
            allowPrivilegeEscalation: false
            seccompProfile:
                type: RuntimeDefault

```

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: securitydemo
  labels:
    app: web
spec:
  replicas: 2
  selector:
    matchLabels:
      app: web
  template:
    metadata:
      labels:
        app: web
    spec:
      securityContext:
        runAsUser: 19000
      containers:
        - name: frontend
          image: nginx
          ports:
            - containerPort: 80
          securityContext:
            allowPrivilegeEscalation: false
        - name: echoserver
          image: k8s.gcr.io/echoserver:1.4
          ports:
            - containerPort: 8080
          securityContext:
            allowPrivilegeEscalation: false
            seccompProfile:
                type: Unconfined

```

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: securitydemo
  labels:
    app: web
spec:
  replicas: 2
  selector:
    matchLabels:
      app: web
  template:
    metadata:
      labels:
        app: web
    spec:
      securityContext:
        runAsUser: 19000
        seccompProfile:
            type: RuntimeDefault
      containers:
        - name: frontend
          image: nginx
          ports:
            - containerPort: 80
          securityContext:
            allowPrivilegeEscalation: false
        - name: echoserver
          image: k8s.gcr.io/echoserver:1.4
          ports:
            - containerPort: 8080
          securityContext:
            allowPrivilegeEscalation: false
            seccompProfile:
                type: Unconfined

```