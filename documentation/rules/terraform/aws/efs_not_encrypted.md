---
title: "EFS Not Encrypted"
meta:
  name: "aws/efs_not_encrypted"
  id: "48207659-729f-4b5c-9402-f884257d794f"
  display_name: "EFS Not Encrypted"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata
**Name:** `aws/efs_not_encrypted`
**Query Name** `EFS Not Encrypted`
**Id:** `48207659-729f-4b5c-9402-f884257d794f`
**Cloud Provider:** aws
**Platform** Terraform
**Severity:** High
**Category:** Encryption
## Description
AWS Elastic File System (EFS) stores data in clear text by default, potentially exposing sensitive information if the storage system is compromised. When EFS is not encrypted, unauthorized users who gain access to the underlying storage could read file contents, leading to data breaches and compliance violations. To properly secure an EFS file system, set the 'encrypted' attribute to 'true' in your Terraform configuration, as shown below:

```hcl
resource "aws_efs_file_system" "secure_example" {
  creation_token = "my-product"
  encrypted = true
  
  tags = {
    Name = "MyProduct"
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/efs_file_system#encrypted)


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
## Non-Compliant Code Examples
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