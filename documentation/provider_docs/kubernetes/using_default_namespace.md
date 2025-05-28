---
title: "Using Default Namespace"
meta:
  name: "kubernetes/using_default_namespace"
  id: "abcb818b-5af7-4d72-aba9-6dd84956b451"
  cloud_provider: "kubernetes"
  severity: "LOW"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `kubernetes/using_default_namespace`

**Id:** `abcb818b-5af7-4d72-aba9-6dd84956b451`

**Cloud Provider:** kubernetes

**Severity:** Low

**Category:** Insecure Configurations

## Description
The default namespace should not be used

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/pod#namespace)

## Non-Compliant Code Examples
```terraform
resource "kubernetes_pod" "test" {
  metadata {
    name = "terraform-example"
    namespace = "default"
  }
}

resource "kubernetes_cron_job" "test2" {
  metadata {
    name = "terraform-example"
  }
}

```

## Compliant Code Examples
```terraform
resource "kubernetes_pod" "test3" {
  metadata {
    name = "terraform-example"
    namespace = "terraform-namespace"
  }
}

resource "kubernetes_cron_job" "test4" {
  metadata {
    name = "terraform-example"
    namespace = "terraform-namespace"
  }
}

```