---
title: "Security context deny admission control plugin not set"
group_id: "rules/k8s"
meta:
  name: "k8s/security_context_deny_admission_control_plugin_not_set"
  id: "6a68bebe-c021-492e-8ddb-55b0567fb768"
  display_name: "Security context deny admission control plugin not set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `6a68bebe-c021-492e-8ddb-55b0567fb768`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Configurations

### Description

 When using `kube-apiserver`, the `--enable-admission-plugins` flag should include `SecurityContextDeny`, and it should be correctly configured in the admission control config file when `PodSecurityPolicy` is not set.
