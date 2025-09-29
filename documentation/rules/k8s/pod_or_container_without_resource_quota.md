---
title: "Pod or container without ResourceQuota"
group_id: "rules/k8s"
meta:
  name: "k8s/pod_or_container_without_resource_quota"
  id: "48a5beba-e4c0-4584-a2aa-e6894e4cf424"
  display_name: "Pod or container without ResourceQuota"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `48a5beba-e4c0-4584-a2aa-e6894e4cf424`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Insecure Configurations

### Description

 Each namespace should have a ResourceQuota policy associated with limiting the total resources that Pods, containers, and PersistentVolumeClaims can consume.
