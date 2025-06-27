---
title: "Neptune Logging Is Disabled"
meta:
  name: "aws/neptune_logging_disabled"
  id: "45cff7b6-3b80-40c1-ba7b-2cf480678bb8"
  display_name: "Neptune Logging Is Disabled"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata

**Name:** `aws/neptune_logging_disabled`

**Query Name** `Neptune Logging Is Disabled`

**Id:** `45cff7b6-3b80-40c1-ba7b-2cf480678bb8`

**Cloud Provider:** aws

**Platform** Terraform

**Severity:** Medium

**Category:** Observability

## Description
Enabling Neptune logging ensures that audit and error logs are exported to Amazon CloudWatch, which is critical for monitoring, troubleshooting, and security auditing of Neptune database activity. If the `enable_cloudwatch_logs_exports` attribute is not set with values such as `["audit"]` or `["audit", "error"]`, as shown below:

```
resource "aws_neptune_cluster" "example" {
  ...
  enable_cloudwatch_logs_exports = ["audit", "error"]
}
```

no logs will be exported by default, leaving potentially malicious or unauthorized database actions undetected. Without these logs, it becomes challenging to investigate incidents, meet compliance requirements, or identify operational issues, increasing the risk of undetected attacks or data breaches.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/neptune_cluster#enable_cloudwatch_logs_exports)


## Compliant Code Examples
```terraform
resource "aws_neptune_cluster" "negative1" {
  cluster_identifier                  = "neptune-cluster"
  engine                              = "neptune"
  backup_retention_period             = 5
  preferred_backup_window             = "10:10-11:11"
  skip_final_snapshot                 = true
  iam_database_authentication_enabled = true
  apply_immediately                   = true
  enable_cloudwatch_logs_exports      = ["audit"]
}

resource "aws_neptune_cluster" "negative2" {
  cluster_identifier                  = "neptune-cluster"
  engine                              = "neptune"
  backup_retention_period             = 5
  preferred_backup_window             = "10:10-11:11"
  skip_final_snapshot                 = true
  iam_database_authentication_enabled = true
  apply_immediately                   = true
  enable_cloudwatch_logs_exports      = ["audit", "error"]
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_neptune_cluster" "postive2" {
  cluster_identifier                  = "neptune-cluster"
  engine                              = "neptune"
  backup_retention_period             = 5
  preferred_backup_window             = "10:10-11:11"
  skip_final_snapshot                 = true
  iam_database_authentication_enabled = true
  apply_immediately                   = true
  enable_cloudwatch_logs_exports      = []
}

```

```terraform
resource "aws_neptune_cluster" "postive3" {
  cluster_identifier                  = "neptune-cluster"
  engine                              = "neptune"
  backup_retention_period             = 5
  preferred_backup_window             = "10:10-11:11"
  skip_final_snapshot                 = true
  iam_database_authentication_enabled = true
  apply_immediately                   = true
  enable_cloudwatch_logs_exports      = ["error"]
}

```

```terraform
resource "aws_neptune_cluster" "postive1" {
  cluster_identifier                  = "neptune-cluster"
  engine                              = "neptune"
  backup_retention_period             = 5
  preferred_backup_window             = "10:10-11:11"
  skip_final_snapshot                 = true
  iam_database_authentication_enabled = true
  apply_immediately                   = true
}

```