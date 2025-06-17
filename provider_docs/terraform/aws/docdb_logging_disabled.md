---
title: "DocDB Logging Is Disabled"
meta:
  name: "terraform/docdb_logging_disabled"
  id: "56f6a008-1b14-4af4-b9b2-ab7cf7e27641"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata
**Name:** `terraform/docdb_logging_disabled`
**Id:** `56f6a008-1b14-4af4-b9b2-ab7cf7e27641`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Observability
## Description
Enabling logging for Amazon DocumentDB clusters helps ensure that database activity is captured and monitored, allowing for the detection of anomalous behavior and aiding in incident investigations. If logging is not enabled by omitting the `enabled_cloudwatch_logs_exports` attribute, critical events and queries may go unrecorded, making it difficult to audit access or troubleshoot security events. To enforce secure configurations, the `aws_docdb_cluster` resource should specify the desired log exports, such as in the example below:

```
resource "aws_docdb_cluster" "example" {
  cluster_identifier      = "my-docdb-cluster"
  engine                  = "docdb"
  master_username         = "foo"
  master_password         = "mustbeeightchars"
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
  skip_final_snapshot     = true

  enabled_cloudwatch_logs_exports = ["profiler", "audit"]
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/docdb_cluster#enabled_cloudwatch_logs_exports)

## Non-Compliant Code Examples
```aws
resource "aws_docdb_cluster" "positive1" {
  cluster_identifier      = "my-docdb-cluster"
  engine                  = "docdb"
  master_username         = "foo"
  master_password         = "mustbeeightchars"
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
  skip_final_snapshot     = true
}

```

```aws
resource "aws_docdb_cluster" "positive4" {
  cluster_identifier      = "my-docdb-cluster"
  engine                  = "docdb"
  master_username         = "foo"
  master_password         = "mustbeeightchars"
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
  skip_final_snapshot     = true

  enabled_cloudwatch_logs_exports = ["audit"]
}

```

```aws
resource "aws_docdb_cluster" "positive3" {
  cluster_identifier      = "my-docdb-cluster"
  engine                  = "docdb"
  master_username         = "foo"
  master_password         = "mustbeeightchars"
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
  skip_final_snapshot     = true

  enabled_cloudwatch_logs_exports = ["profiler"]
}

```

## Compliant Code Examples
```aws
resource "aws_docdb_cluster" "negative1" {
  cluster_identifier      = "my-docdb-cluster"
  engine                  = "docdb"
  master_username         = "foo"
  master_password         = "mustbeeightchars"
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
  skip_final_snapshot     = true

  enabled_cloudwatch_logs_exports = ["profiler", "audit"]
}

```