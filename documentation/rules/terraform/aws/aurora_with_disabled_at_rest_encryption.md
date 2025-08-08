---
title: "Aurora with disabled at rest encryption"
group_id: "rules/terraform/aws"
meta:
  name: "aws/aurora_with_disabled_at_rest_encryption"
  id: "1a690d1d-0ae7-49fa-b2db-b75ae0dd1d3e"
  display_name: "Aurora with disabled at rest encryption"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata

**Id:** `1a690d1d-0ae7-49fa-b2db-b75ae0dd1d3e`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** High

**Category:** Encryption

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/rds_cluster#storage_encrypted)

### Description

 Amazon Aurora clusters should have encryption at rest enabled to protect sensitive data stored in the database. When `storage_encrypted` is set to `false` or omitted, data stored in the underlying storage is vulnerable to unauthorized access if the storage media is compromised or improperly disposed of. Enabling encryption at rest helps maintain data confidentiality and supports compliance with regulatory requirements for data protection. To address this vulnerability, set the `storage_encrypted` attribute to true in your `aws_rds_cluster` resource, as shown in the secure example:

```terraform
resource "aws_rds_cluster" "my_cluster" {
  // other configuration
  storage_encrypted = true
}
```


## Compliant Code Examples
```terraform
module "rds_cluster" {
  source = "terraform-aws-modules/rds-aurora/aws"
  version = "~> 7.0"

  name            = "aurora-example"
  engine          = "aurora-mysql"
  engine_version  = "5.7.12"
  instance_type   = "r5.large"
  storage_encrypted = true

  username = "root"
  password = "123SecurePassword123"
  port     = "3306"

  vpc_id  = "vpc-12345678"
  subnets = ["subnet-12345678", "subnet-87654321"]

  allowed_cidr_blocks = ["10.20.0.0/16"]

  monitoring_interval = 10
  instance_monitoring_interval = 60

  apply_immediately   = true
  skip_final_snapshot = true

  # Database Deletion Protection
  # deletion_protection = true

  tags = {
    Owner       = "user"
    Environment = "dev"
  }
}

```

```terraform
provider "aws" {
  region = "us-west-2"  # Replace with your desired AWS region
}

resource "aws_rds_cluster" "my_cluster" {
  cluster_identifier            = "my-cluster"
  engine                       = "aurora-mysql"
  engine_version               = "5.7.mysql_aurora.2.08.0"
  master_username              = "admin"
  master_password              = "password"
  backup_retention_period      = 7
  preferred_backup_window      = "02:00-03:00"
  deletion_protection          = false
  skip_final_snapshot          = true
  apply_immediately            = true
  storage_encrypted            = true
}

resource "aws_rds_cluster_instance" "my_cluster_instance" {
  identifier                    = "my-cluster-instance"
  cluster_identifier            = aws_rds_cluster.my_cluster.id
  engine                        = "aurora-mysql"
  instance_class                = "db.r5.large"
  publicly_accessible           = false
  availability_zone             = "us-west-2a"  # Replace with your desired availability zone
}

output "cluster_endpoint" {
  value = aws_rds_cluster.my_cluster.endpoint
}

```
## Non-Compliant Code Examples
```terraform
provider "aws" {
  region = "us-west-2"  # Replace with your desired AWS region
}

resource "aws_rds_cluster" "my_cluster" {
  cluster_identifier            = "my-cluster"
  engine                       = "aurora-mysql"
  engine_version               = "5.7.mysql_aurora.2.08.0"
  master_username              = "admin"
  master_password              = "password"
  backup_retention_period      = 7
  preferred_backup_window      = "02:00-03:00"
  deletion_protection          = false
  skip_final_snapshot          = true
  apply_immediately            = true
  storage_encrypted            = false
}

resource "aws_rds_cluster_instance" "my_cluster_instance" {
  identifier                    = "my-cluster-instance"
  cluster_identifier            = aws_rds_cluster.my_cluster.id
  engine                        = "aurora-mysql"
  instance_class                = "db.r5.large"
  publicly_accessible           = false
  availability_zone             = "us-west-2a"  # Replace with your desired availability zone
}

output "cluster_endpoint" {
  value = aws_rds_cluster.my_cluster.endpoint
}

```

```terraform
module "rds_cluster" {
  source = "terraform-aws-modules/rds-aurora/aws"
  version = "~> 7.0"

  name            = "aurora-example"
  engine          = "aurora-mysql"
  engine_version  = "5.7.12"
  instance_type   = "r5.large"
  storage_encrypted = false

  username = "root"
  password = "123SecurePassword123"
  port     = "3306"

  vpc_id  = "vpc-12345678"
  subnets = ["subnet-12345678", "subnet-87654321"]

  allowed_cidr_blocks = ["10.20.0.0/16"]

  monitoring_interval = 10
  instance_monitoring_interval = 60

  apply_immediately   = true
  skip_final_snapshot = true

  # Database Deletion Protection
  # deletion_protection = true

  tags = {
    Owner       = "user"
    Environment = "dev"
  }
}
```

```terraform
provider "aws" {
  region = "us-west-2"  # Replace with your desired AWS region
}

resource "aws_rds_cluster" "my_cluster" {
  cluster_identifier            = "my-cluster"
  engine                       = "aurora-mysql"
  engine_version               = "5.7.mysql_aurora.2.08.0"
  master_username              = "admin"
  master_password              = "password"
  backup_retention_period      = 7
  preferred_backup_window      = "02:00-03:00"
  deletion_protection          = false
  skip_final_snapshot          = true
  apply_immediately            = true
}

resource "aws_rds_cluster_instance" "my_cluster_instance" {
  identifier                    = "my-cluster-instance"
  cluster_identifier            = aws_rds_cluster.my_cluster.id
  engine                        = "aurora-mysql"
  instance_class                = "db.r5.large"
  publicly_accessible           = false
  availability_zone             = "us-west-2a"  # Replace with your desired availability zone
}

output "cluster_endpoint" {
  value = aws_rds_cluster.my_cluster.endpoint
}

```