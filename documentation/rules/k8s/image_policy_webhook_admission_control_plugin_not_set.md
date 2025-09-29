---
title: "Image policy webhook admission control plugin not set"
group_id: "rules/k8s"
meta:
  name: "k8s/image_policy_webhook_admission_control_plugin_not_set"
  id: "14abda69-8e91-4acb-9931-76e2bee90284"
  display_name: "Image policy webhook admission control plugin not set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Build Process"
---
## Metadata

**Id:** `14abda69-8e91-4acb-9931-76e2bee90284`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Build Process

### Description

 When using `kube-apiserver`, the `--enable-admission-plugins` flag should include `ImagePolicyWebhook`, and it should be correctly configured in the admission control config file.
