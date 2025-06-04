---
title: "DOCDB Cluster Encrypted With AWS Managed Key"
meta:
  name: "aws/docdb_cluster_encrypted_with_aws_managed_key"
  id: "2134641d-30a4-4b16-8ffc-2cd4c4ffd15d"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Encryption"
---

## Metadata
**Name:** `aws/docdb_cluster_encrypted_with_aws_managed_key`

**Id:** `2134641d-30a4-4b16-8ffc-2cd4c4ffd15d`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Encryption

## Description
DOCDB Cluster should be encrypted with customer-managed KMS keys instead of AWS managed keys

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/docdb_cluster#kms_key_id)

## Non-Compliant Code Examples
```terraform
provider "aws" {
  region = "us-east-1"
}

data "aws_kms_key" "test" {
  key_id = "alias/aws/rds"
}

resource "aws_docdb_cluster" "test2" {
  cluster_identifier  = "my-docdb-cluster-test2"
  engine              = "docdb"
  master_username     = "foo"
  master_password     = "mustbeeightchars"
  skip_final_snapshot = true
  storage_encrypted   = true
  kms_key_id          = data.aws_kms_key.test.arn
}

```

## Compliant Code Examples
```terraform
provider "aws" {
  region = "us-east-1"
}

data "aws_kms_key" "test2" {
  key_id = "alias/myAlias"
}

resource "aws_docdb_cluster" "test22" {
  cluster_identifier  = "my-docdb-cluster-test2"
  engine              = "docdb"
  master_username     = "foo"
  master_password     = "mustbeeightchars"
  skip_final_snapshot = true
  storage_encrypted   = true
  kms_key_id          = data.aws_kms_key.test2.arn
}

```