---
title: "Redshift Cluster Without VPC"
meta:
  name: "aws/redshift_cluster_without_vpc"
  id: "0a494a6a-ebe2-48a0-9d77-cf9d5125e1b3"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/redshift_cluster_without_vpc`

**Id:** `0a494a6a-ebe2-48a0-9d77-cf9d5125e1b3`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Insecure Configurations

## Description
Redshift Cluster should be configured in VPC (Virtual Private Cloud)

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/redshift_cluster#vpc_security_group_ids)

## Non-Compliant Code Examples
```terraform
resource "aws_redshift_cluster" "positive1" {
  cluster_identifier = "tf-redshift-cluster"
  database_name      = "mydb"
  master_username    = "foo"
  master_password    = "Mustbe8characters"
  node_type          = "dc1.large"
  cluster_type       = "single-node"
  logging {
      enable = true
      bucket_name = "nameOfAnExistingS3Bucket"
  }
}

```

## Compliant Code Examples
```terraform
resource "aws_redshift_cluster" "negative1" {
  cluster_identifier = "tf-redshift-cluster"
  database_name      = "mydb"
  master_username    = "foo"
  master_password    = "Mustbe8characters"
  node_type          = "dc1.large"
  cluster_type       = "single-node"
  logging {
      enable = true
      bucket_name = "nameOfAnExistingS3Bucket"
  }
  vpc_security_group_ids = [
    aws_security_group.redshift.id
  ]
  cluster_subnet_group_name = aws_redshift_subnet_group.redshift_subnet_group.id
}

```