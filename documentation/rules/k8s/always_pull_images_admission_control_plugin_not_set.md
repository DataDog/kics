---
title: "Always pull images admission control plugin not set"
group_id: "rules/k8s"
meta:
  name: "k8s/always_pull_images_admission_control_plugin_not_set"
  id: "a77f4d07-c6e0-4a48-8b35-0eeb51576f4f"
  display_name: "Always pull images admission control plugin not set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Build Process"
---
## Metadata

**Id:** `a77f4d07-c6e0-4a48-8b35-0eeb51576f4f`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Build Process

### Description

 When using `kube-apiserver`, the `--enable-admission-plugins` flag should include `AlwaysPullImages`, and it should be correctly configured in the admission control config file.
