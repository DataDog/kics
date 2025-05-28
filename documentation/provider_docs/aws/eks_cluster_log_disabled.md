---
title: "EKS cluster logging is not enabled"
meta:
  name: "aws/eks_cluster_log_disabled"
  id: "37304d3f-f852-40b8-ae3f-725e87a7cedf"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `aws/eks_cluster_log_disabled`

**Id:** `37304d3f-f852-40b8-ae3f-725e87a7cedf`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Observability

## Description
Amazon EKS control plane logging is not enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eks_cluster#enabled_cluster_log_types)

## Non-Compliant Code Examples
```terraform
variable "cluster_name" {
  default = "example"
  type    = string
}

resource "aws_eks_cluster" "positive1" {
  depends_on = [aws_cloudwatch_log_group.example]
  name                      = var.cluster_name
}

```

## Compliant Code Examples
```terraform
variable "cluster_name" {
  default = "example"
  type    = string
}

resource "aws_eks_cluster" "negative1" {
  depends_on = [aws_cloudwatch_log_group.example]

  enabled_cluster_log_types = ["api", "audit", "authenticator", "controllerManager", "scheduler"]
  name                      = var.cluster_name
}

```