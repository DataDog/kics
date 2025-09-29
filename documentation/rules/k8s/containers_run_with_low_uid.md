---
title: "Container with low UID"
group_id: "rules/k8s"
meta:
  name: "k8s/containers_run_with_low_uid"
  id: "02323c00-cdc3-4fdc-a310-4f2b3e7a1660"
  display_name: "Container with low UID"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Best Practices"
---
## Metadata

**Id:** `02323c00-cdc3-4fdc-a310-4f2b3e7a1660`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Best Practices

### Description

 Containers should not run with a low UID, as this may cause conflicts with the host's user table.
