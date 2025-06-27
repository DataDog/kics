---
title: "DOCDB Cluster Not Encrypted"
meta:
  name: "aws/docdb_cluster_not_encrypted"
  id: "bc1f9009-84a0-490f-ae09-3e0ea6d74ad6"
  display_name: "DOCDB Cluster Not Encrypted"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata

**Name:** `aws/docdb_cluster_not_encrypted`

**Query Name** `DOCDB Cluster Not Encrypted`

**Id:** `bc1f9009-84a0-490f-ae09-3e0ea6d74ad6`

**Cloud Provider:** aws

**Platform** Terraform

**Severity:** High

**Category:** Encryption

## Description
This check verifies whether Amazon DocumentDB cluster storage encryption is enabled. DocumentDB clusters store sensitive data and should have storage encryption enabled to protect data at rest. When storage encryption is disabled or not configured, data stored in the cluster is vulnerable to unauthorized access if the underlying storage is compromised.

To properly secure a DocumentDB cluster, ensure the 'storage_encrypted' attribute is explicitly set to 'true' as shown in the example below:
```terraform
resource "aws_docdb_cluster" "docdb" {
  cluster_identifier = "my-docdb-cluster"
  // ... other configuration ...
  storage_encrypted = true
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/docdb_cluster#storage_encrypted)


## Compliant Code Examples
```terraform
resource "aws_docdb_cluster" "docdb" {
  cluster_identifier      = "my-docdb-cluster"
  engine                  = "docdb"
  master_username         = "foo"
  master_password         = "mustbeeightchars"
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
  skip_final_snapshot     = true
  storage_encrypted = true
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_docdb_cluster" "docdb" {
  cluster_identifier      = "my-docdb-cluster"
  engine                  = "docdb"
  master_username         = "foo"
  master_password         = "mustbeeightchars"
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
  skip_final_snapshot     = true
}

resource "aws_docdb_cluster" "docdb_2" {
  cluster_identifier      = "my-docdb-cluster"
  engine                  = "docdb"
  master_username         = "foo"
  master_password         = "mustbeeightchars"
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
  skip_final_snapshot     = true
  storage_encrypted = false
}

```