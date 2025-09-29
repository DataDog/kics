---
title: "Kubelet not managing IP tables"
group_id: "rules/k8s"
meta:
  name: "k8s/kubelet_not_managing_ip_tables"
  id: "5f89001f-6dd9-49ff-9b15-d8cd71b617f4"
  display_name: "Kubelet not managing IP tables"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `5f89001f-6dd9-49ff-9b15-d8cd71b617f4`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Networking and Firewall

### Description

 The kubelet argument `--make-iptables-util-chains` should be set to `true`.
