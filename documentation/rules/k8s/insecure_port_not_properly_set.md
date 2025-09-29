---
title: "Insecure port not properly set"
group_id: "rules/k8s"
meta:
  name: "k8s/insecure_port_not_properly_set"
  id: "fa4def8c-1898-4a35-a139-7b76b1acdef0"
  display_name: "Insecure port not properly set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `fa4def8c-1898-4a35-a139-7b76b1acdef0`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Networking and Firewall

### Description

 When using `kube-apiserver`, the `--insecure-port` flag should be set to `0`.
