---
title: "EMR Without VPC"
meta:
  name: "aws/emr_without_vpc"
  id: "2b3c8a6d-9856-43e6-ab1d-d651094f03b4"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/emr_without_vpc`

**Id:** `2b3c8a6d-9856-43e6-ab1d-d651094f03b4`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Networking and Firewall

## Description
Elastic MapReduce Cluster (EMR) should be launched in a Virtual Private Cloud (VPC)

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/emr_cluster#subnet_id)

## Non-Compliant Code Examples
```terraform
resource "aws_emr_cluster" "positive1" {
  name          = "emr-test-arn"
  release_label = "emr-4.6.0"
}

```

## Compliant Code Examples
```terraform
resource "aws_emr_cluster" "negative1" {
  name          = "emr-test-arn"
  release_label = "emr-4.6.0"
  subnet_id = aws_subnet.main.id
}

```