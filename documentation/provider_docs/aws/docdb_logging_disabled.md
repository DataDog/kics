---
title: "DocDB Logging Is Disabled"
meta:
  name: "aws/docdb_logging_disabled"
  id: "56f6a008-1b14-4af4-b9b2-ab7cf7e27641"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `aws/docdb_logging_disabled`

**Id:** `56f6a008-1b14-4af4-b9b2-ab7cf7e27641`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Observability

## Description
DocDB logging should be enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/docdb_cluster#enabled_cloudwatch_logs_exports)

## Non-Compliant Code Examples
```terraform
resource "aws_docdb_cluster" "positive2" {
  cluster_identifier      = "my-docdb-cluster"
  engine                  = "docdb"
  master_username         = "foo"
  master_password         = "mustbeeightchars"
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
  skip_final_snapshot     = true

  enabled_cloudwatch_logs_exports = []
}

```

```terraform
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

```terraform
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

```terraform
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

## Compliant Code Examples
```terraform
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