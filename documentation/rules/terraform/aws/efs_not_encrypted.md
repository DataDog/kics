---
title: "EFS not encrypted"
group_id: "rules/terraform/aws"
meta:
  name: "aws/efs_not_encrypted"
  id: "48207659-729f-4b5c-9402-f884257d794f"
  display_name: "EFS not encrypted"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata

**Id:** `48207659-729f-4b5c-9402-f884257d794f`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** High

**Category:** Encryption

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/efs_file_system#encrypted)

### Description

 AWS Elastic File System (EFS) stores data in clear text by default, potentially exposing sensitive information if the storage system is compromised. When EFS is not encrypted, unauthorized users who gain access to the underlying storage could read file contents, leading to data breaches and compliance violations. To properly secure an EFS file system, set the `encrypted` attribute to `true` in your Terraform configuration, as shown below:

```hcl
resource "aws_efs_file_system" "secure_example" {
  creation_token = "my-product"
  encrypted = true
  
  tags = {
    Name = "MyProduct"
  }
}
```


## Compliant Code Examples
```terraform
resource "aws_efs_file_system" "negative1" {
  creation_token = "my-product"
  encrypted = true
  
  tags = {
    Name = "MyProduct"
  }
}
```

```terraform
module "efs" {
  source             = "terraform-aws-modules/efs/aws"
  version            = "1.0.0"
  name               = "my-efs"
  encrypted          = true
  creation_token     = "my-product"
  performance_mode   = "generalPurpose"
  throughput_mode    = "bursting"
  lifecycle_policy   = "AFTER_30_DAYS"
  provisioned_throughput_in_mibps = 0
  enabled            = true

  mount_targets = {
    ap-south-1a = {
      subnet_id = "subnet-12345678"
    }
    ap-south-1b = {
      subnet_id = "subnet-87654321"
    }
  }

  security_group_description = "EFS security group"

  security_group_rules = [
    {
      cidr_blocks = ["10.0.0.0/8"]
    }
  ]
}
```
## Non-Compliant Code Examples
```terraform
module "efs" {
  source             = "terraform-aws-modules/efs/aws"
  version            = "1.0.0"
  name               = "my-efs"
  creation_token     = "my-product"
  performance_mode   = "generalPurpose"
  throughput_mode    = "bursting"
  lifecycle_policy   = "AFTER_30_DAYS"
  provisioned_throughput_in_mibps = 0
  enabled            = true

  mount_targets = {
    ap-south-1a = {
      subnet_id = "subnet-12345678"
    }
    ap-south-1b = {
      subnet_id = "subnet-87654321"
    }
  }

  security_group_description = "EFS security group"

  security_group_rules = [
    {
      cidr_blocks = ["10.0.0.0/8"]
    }
  ]
}
```

```terraform
resource "aws_efs_file_system" "positive1" {
  creation_token = "my-product"

  tags = {
    Name = "MyProduct"
  }
}

resource "aws_efs_file_system" "positive2" {
  creation_token = "my-product"
  encrypted = false
  
  tags = {
    Name = "MyProduct"
  }
}
```