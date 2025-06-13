---
title: "Missing Cluster Log Types"
meta:
  name: "aws/missing_cluster_log_types"
  id: "66f130d9-b81d-4e8e-9b08-da74b9c891df"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `aws/missing_cluster_log_types`

**Id:** `66f130d9-b81d-4e8e-9b08-da74b9c891df`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Observability

## Description
Amazon EKS control plane logging don't enabled for all log types

#### Learn More

 - [Provider Reference](https://www.terraform.io/docs/providers/aws/r/eks_cluster.html)

## Non-Compliant Code Examples
```terraform
variable "cluster_name" {
  default = "example"
  type    = string
}

resource "aws_eks_cluster" "positive1" {
  depends_on = [aws_cloudwatch_log_group.example]

  enabled_cluster_log_types = ["api", "audit"]
  name                      = var.cluster_name

  # ... other configuration ...
}

resource "aws_cloudwatch_log_group" "positive2" {
  name              = "/aws/eks/${var.cluster_name}/cluster"
  retention_in_days = 7

  # ... potentially other configuration ...
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

  # ... other configuration ...
}

resource "aws_cloudwatch_log_group" "negative2" {
  name              = "/aws/eks/${var.cluster_name}/cluster"
  retention_in_days = 7

  # ... potentially other configuration ...
}

```