---
title: "DOCDB Cluster Encrypted With AWS Managed Key"
meta:
  name: "aws/docdb_cluster_encrypted_with_aws_managed_key"
  id: "2134641d-30a4-4b16-8ffc-2cd4c4ffd15d"
  display_name: "DOCDB Cluster Encrypted With AWS Managed Key"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "LOW"
  category: "Encryption"
---
## Metadata
**Name:** `aws/docdb_cluster_encrypted_with_aws_managed_key`
**Query Name** `DOCDB Cluster Encrypted With AWS Managed Key`
**Id:** `2134641d-30a4-4b16-8ffc-2cd4c4ffd15d`
**Cloud Provider:** aws
**Platform** Terraform
**Severity:** Low
**Category:** Encryption
## Description
Amazon DocumentDB (DocDB) clusters should be encrypted using customer-managed AWS Key Management Service (KMS) keys rather than default AWS-managed keys. When the `kms_key_id` attribute in the `aws_docdb_cluster` resource is set to an AWS-managed key such as `alias/aws/rds`, like in the configuration below:

```
data "aws_kms_key" "test" {
  key_id = "alias/aws/rds"
}

resource "aws_docdb_cluster" "test2" {
  ...
  storage_encrypted   = true
  kms_key_id          = data.aws_kms_key.test.arn
}
```

it limits the organization's ability to control, rotate, and audit the usage of encryption keys, potentially weakening the security posture. If this misconfiguration is left unaddressed, sensitive data stored in the DocDB cluster may be at risk because AWS retains full control of the KMS key, making it difficult to restrict or monitor access and respond to compliance requirements. Using customer-managed KMS keys allows for finer-grained access control, detailed key usage logging, and key lifecycle management, helping protect data from unauthorized access and achieving regulatory compliance.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/docdb_cluster#kms_key_id)


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