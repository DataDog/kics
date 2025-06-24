---
title: "Neptune Cluster Snapshot Not Encrypted"
meta:
  name: "aws/neptune_snapshots_not_encrypted"
  id: "g3l20gd0k-e5f6-7890-ab12-cd34ef567890"
  cloud_provider: "aws"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata
**Name:** `aws/neptune_snapshots_not_encrypted`
**Id:** `g3l20gd0k-e5f6-7890-ab12-cd34ef567890`
**Cloud Provider:** aws
**Severity:** High
**Category:** Encryption
## Description
AWS Neptune is a fully managed graph database service that stores and queries highly connected data. When Neptune Cluster Snapshots are not encrypted, sensitive data stored in these snapshots could be vulnerable to unauthorized access, potentially exposing proprietary information, personal data, or other confidential content. Enabling encryption for Neptune Cluster Snapshots adds an additional layer of security that helps protect your data at rest.

Secure configuration example:
```terraform
resource "aws_neptune_cluster_snapshot" "good_example" {
  db_cluster_identifier          = "example-cluster"
  db_cluster_snapshot_identifier = "example-snapshot"
  storage_encrypted              = true
}
```

Vulnerable configuration example:
```terraform
resource "aws_neptune_cluster_snapshot" "bad_example" {
  db_cluster_identifier          = "example-cluster"
  db_cluster_snapshot_identifier = "example-snapshot"
  storage_encrypted              = false
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/neptune_cluster_snapshot#storage_encrypted)


## Compliant Code Examples
```terraform
resource "aws_neptune_cluster_snapshot" "good_example" {
  db_cluster_identifier          = "example-cluster"
  db_cluster_snapshot_identifier = "example-snapshot"
  storage_encrypted              = true # ✅ Neptune snapshot encryption is enabled
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_neptune_cluster_snapshot" "bad_example" {
  db_cluster_identifier          = "example-cluster"
  db_cluster_snapshot_identifier = "example-snapshot"
  storage_encrypted              = false # ❌ Neptune snapshot encryption is disabled
}

```