---
title: "RDS Cluster With Backup Disabled"
meta:
  name: "terraform/rds_cluster_with_backup_disabled"
  id: "e542bd46-58c4-4e0f-a52a-1fb4f9548e02"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Backup"
---
## Metadata
**Name:** `terraform/rds_cluster_with_backup_disabled`
**Id:** `e542bd46-58c4-4e0f-a52a-1fb4f9548e02`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Backup
## Description
RDS Cluster backup retention period should be specifically defined. When creating an AWS RDS Cluster using Terraform, omitting the `backup_retention_period` parameter allows the database to default to the minimum backup retention, which could be zero or just one day depending on the engine. This short or undefined retention window risks losing the ability to restore data to a specific point in time, potentially resulting in irreversible data loss in the event of accidental deletion, corruption, or ransomware attacks. Explicitly setting a sufficient retention period ensures backups are available for recovery as required by business continuity or compliance requirements.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/rds_cluster#backup_retention_period)

## Non-Compliant Code Examples
```aws
resource "aws_rds_cluster" "postgresql" {
  cluster_identifier      = "aurora-cluster-demo"
  engine                  = "aurora-postgresql"
  availability_zones      = ["us-west-2a", "us-west-2b", "us-west-2c"]
  database_name           = "mydb"
  master_username         = "foo"
  master_password         = "bar"
  preferred_backup_window = "07:00-09:00"
}


```

## Compliant Code Examples
```aws
resource "aws_rds_cluster" "postgresql" {
  cluster_identifier      = "aurora-cluster-demo"
  engine                  = "aurora-postgresql"
  availability_zones      = ["us-west-2a", "us-west-2b", "us-west-2c"]
  database_name           = "mydb"
  master_username         = "foo"
  master_password         = "bar"
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
}

```