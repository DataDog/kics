---
title: "RDS Database Cluster not Encrypted"
meta:
  name: "terraform/rds_database_cluster_not_encrypted"
  id: "656880aa-1388-488f-a6d4-8f73c23149b2"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata
**Name:** `terraform/rds_database_cluster_not_encrypted`
**Id:** `656880aa-1388-488f-a6d4-8f73c23149b2`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Encryption
## Description
Amazon RDS Database Clusters should be encrypted at rest to protect sensitive data from unauthorized access if storage is compromised. When not properly configured with the 'storage_encrypted' attribute set to true, database contents remain in plaintext, potentially exposing sensitive information to attackers who gain access to the underlying storage. To secure your RDS cluster, ensure encryption is enabled as shown below:

```
resource "aws_rds_cluster" "secure_example" {
  // Required configuration
  cluster_identifier   = "example"
  engine               = "aurora"
  master_password      = "securepassword"
  master_username      = "admin"
  
  // Security configuration
  storage_encrypted    = true  // Enables encryption at rest
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_cluster_snapshot)

## Non-Compliant Code Examples
```aws
resource "aws_db_cluster_snapshot" "positive1" {
  db_cluster_identifier          = aws_rds_cluster.example2.id 
  db_cluster_snapshot_identifier = "resourcetestsnapshot1234"
}

resource "aws_rds_cluster" "example2" {
  cluster_identifier   = "example"
  db_subnet_group_name = aws_db_subnet_group.example.name
  engine_mode          = "multimaster"
  master_password      = "barbarbarbar"
  master_username      = "foo"
  skip_final_snapshot  = true
}

```

```aws
resource "aws_db_cluster_snapshot" "positive2" {
  db_cluster_identifier          = aws_rds_cluster.example3.id 
  db_cluster_snapshot_identifier = "resourcetestsnapshot1234"
}

resource "aws_rds_cluster" "example3" {
  cluster_identifier   = "example"
  db_subnet_group_name = aws_db_subnet_group.example.name
  engine_mode          = "multimaster"
  master_password      = "barbarbarbar"
  master_username      = "foo"
  skip_final_snapshot  = true
  storage_encrypted    = false
}

```

## Compliant Code Examples
```aws
resource "aws_db_cluster_snapshot" "negative" {
  db_cluster_identifier          = aws_rds_cluster.example.id 
  db_cluster_snapshot_identifier = "resourcetestsnapshot1234"
}

resource "aws_rds_cluster" "example" {
  cluster_identifier   = "example"
  db_subnet_group_name = aws_db_subnet_group.example.name
  engine_mode          = "multimaster"
  master_password      = "barbarbarbar"
  master_username      = "foo"
  skip_final_snapshot  = true
  storage_encrypted    = true
}

```