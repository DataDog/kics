---
title: "Request timeout not properly set"
group_id: "rules/k8s"
meta:
  name: "k8s/request_timeout_not_properly_set"
  id: "d89a15bb-8dba-4c71-9529-bef6729b9c09"
  display_name: "Request timeout not properly set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Availability"
---
## Metadata

**Id:** `d89a15bb-8dba-4c71-9529-bef6729b9c09`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Availability

### Description

 When using `kube-apiserver`, the `--request-timeout` value should not be excessively long.
