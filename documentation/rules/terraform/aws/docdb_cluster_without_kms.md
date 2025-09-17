---
title: "DocumentDB cluster without KMS"
group_id: "rules/terraform/aws"
meta:
  name: "aws/docdb_cluster_without_kms"
  id: "4766d3ea-241c-4ee6-93ff-c380c996bd1a"
  display_name: "DocumentDB cluster without KMS"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata

**Id:** `4766d3ea-241c-4ee6-93ff-c380c996bd1a`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** High

**Category:** Encryption

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/docdb_cluster#kms_key_id)

### Description

 AWS DocumentDB clusters should be encrypted with a KMS encryption key to protect sensitive data at rest. Without proper encryption, your database contents could be exposed if unauthorized access to the storage occurs, potentially leading to data breaches and compliance violations. To secure your DocumentDB cluster, you must enable storage encryption and specify a KMS key ID, as shown in the following secure example:

```terraform
resource "aws_docdb_cluster" "docdb" {
  cluster_identifier = "my-docdb-cluster"
  storage_encrypted = true
  kms_key_id = "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
  // other configuration...
}
```


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
  kms_key_id = "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
}

```

```terraform
module "cluster" {
  source  = "terraform-aws-modules/docdb/aws"
  version = "~> 1.0"

  identifier    = "democlustor"
  cluster_size  = 2
  engine_version = "3.6.0"
  instance_class = "db.r5.large"
  admin_password = "Mustbe8-characters"
  kms_key_id     = "arn:aws:kms:us-west-2:xxxxxxxxxxxxx:key/xxxxxxxxxxxx"
  vpc_id         = "vpc-4e00c123"

  preferred_backup_window = "07:00-09:00"
  preferred_maintenance_window = "Mon:00:00-Mon:03:00"

  tags = {
    Owner       = "user"
    Environment = "dev"
  }
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

```

```terraform
module "cluster" {
  source  = "terraform-aws-modules/docdb/aws"
  version = "~> 1.0"

  identifier    = "democlustor"
  cluster_size  = 2
  engine_version = "3.6.0"
  instance_class = "db.r5.large"
  admin_password = "Mustbe8-characters"
  vpc_id         = "vpc-4e00c123"

  preferred_backup_window = "07:00-09:00"
  preferred_maintenance_window = "Mon:00:00-Mon:03:00"

  tags = {
    Owner       = "user"
    Environment = "dev"
  }
}
```