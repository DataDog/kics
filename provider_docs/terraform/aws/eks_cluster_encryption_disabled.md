---
title: "EKS Cluster Encryption Disabled"
meta:
  name: "terraform/eks_cluster_encryption_disabled"
  id: "63ebcb19-2739-4d3f-aa5c-e8bbb9b85281"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata
**Name:** `terraform/eks_cluster_encryption_disabled`
**Id:** `63ebcb19-2739-4d3f-aa5c-e8bbb9b85281`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Encryption
## Description
Amazon EKS clusters store sensitive information including certificate authorities and service account tokens. When encryption is disabled, this sensitive data is stored in plaintext, potentially exposing it to unauthorized access and data breaches. Enabling encryption using KMS keys for EKS clusters adds an essential layer of security by encrypting Kubernetes secrets stored in etcd.

Insecure example without encryption:
```
resource "aws_eks_cluster" "positive1" {
  depends_on = [aws_cloudwatch_log_group.example]
  name = var.cluster_name
  // Missing encryption_config block
}
```

Secure example with encryption enabled:
```
resource "aws_eks_cluster" "negative1" {
  depends_on = [aws_cloudwatch_log_group.example]
  name = var.cluster_name

  encryption_config {
    resources = ["secrets"]
    provider {
      key_arn = "your-kms-key-arn"
    }
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eks_cluster#encryption_config)

## Non-Compliant Code Examples
```aws
variable "cluster_name" {
  default = "example"
  type    = string
}

resource "aws_eks_cluster" "positive1" {
  depends_on = [aws_cloudwatch_log_group.example]
  name                      = var.cluster_name
}

```

```aws
variable "cluster_name" {
  default = "example"
  type    = string
}

resource "aws_eks_cluster" "positive2" {
  depends_on = [aws_cloudwatch_log_group.example]
  name                      = var.cluster_name

  encryption_config {
    resources = ["s"]
    provider {
      key_arn = "test"
    }
  }
}

```

## Compliant Code Examples
```aws
variable "cluster_name" {
  default = "example"
  type    = string
}

resource "aws_eks_cluster" "negative1" {
  depends_on = [aws_cloudwatch_log_group.example]
  name                      = var.cluster_name

  encryption_config {
    resources = ["secrets"]
    provider {
      key_arn = "test"
    }
  }
}

```