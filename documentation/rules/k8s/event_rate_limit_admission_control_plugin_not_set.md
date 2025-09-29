---
title: "Event rate limit admission control plugin not set"
group_id: "rules/k8s"
meta:
  name: "k8s/event_rate_limit_admission_control_plugin_not_set"
  id: "e0099af2-fe17-411f-9991-0de28fe15f3c"
  display_name: "Event rate limit admission control plugin not set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Availability"
---
## Metadata

**Id:** `e0099af2-fe17-411f-9991-0de28fe15f3c`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Availability

### Description

 When using `kube-apiserver`, the `--enable-admission-plugins` flag should include `EventRateLimit`, and it should be correctly configured in the admission control config file.
