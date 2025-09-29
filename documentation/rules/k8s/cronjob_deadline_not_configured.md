---
title: "CronJob deadline not configured"
group_id: "rules/k8s"
meta:
  name: "k8s/cronjob_deadline_not_configured"
  id: "192fe40b-b1c3-448a-aba2-6cc19a300fe3"
  display_name: "CronJob deadline not configured"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Resource Management"
---
## Metadata

**Id:** `192fe40b-b1c3-448a-aba2-6cc19a300fe3`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Resource Management

### Description

 CronJobs must have a configured deadline. The `startingDeadlineSeconds` attribute must be defined.
