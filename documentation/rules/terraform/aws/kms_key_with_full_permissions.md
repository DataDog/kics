---
title: "KMS key with vulnerable policy"
group_id: "rules/terraform/aws"
meta:
  name: "aws/kms_key_with_full_permissions"
  id: "7ebc9038-0bde-479a-acc4-6ed7b6758899"
  display_name: "KMS key with vulnerable policy"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `7ebc9038-0bde-479a-acc4-6ed7b6758899`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** High

**Category:** Insecure Configurations

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms_key)

### Description

 This check identifies AWS KMS keys with overly permissive policies that grant full access to all AWS services or users. KMS keys with policies allowing `kms:*` actions to all principals (`AWS:*`) create a significant security risk, potentially allowing unauthorized users to access, modify, or delete encrypted data across your AWS environment.

Vulnerable policies typically include a statement with `Effect:Allow`, `Principal:{"AWS":"*"}`, and `Action:["kms:*"]`, as shown in this insecure example:
```
"Statement":[{
  "Effect":"Allow",
  "Principal": {"AWS":"*"},
  "Action":["kms:*"],
  "Resource":"*"
}]
```

Secure your KMS keys by using least privilege principless—restrict access to specific IAM roles/users, limit actions to only those needed, and use explicit `Deny` statements where appropriate as shown in the following example:
```
"Statement":[{
  "Effect":"Deny",
  "Principal": {"AWS": ["arn:aws:iam::111122223333:user/CMKUser"]},
  "Action": ["kms:Encrypt", "kms:Decrypt", "kms:ReEncrypt*", "kms:GenerateDataKey*", "kms:DescribeKey"],
  "Resource":"*"
}]
```


## Compliant Code Examples
```terraform
resource "aws_kms_key" "negative1" {
  description             = "KMS key 1"
  deletion_window_in_days = 10

  policy = <<POLICY
  {
    "Version": "2012-10-17",
    "Statement":[
      {
        "Sid":"AddCannedAcl",
        "Effect":"Deny",
        "Principal": {"AWS": [
          "arn:aws:iam::111122223333:user/CMKUser"
        ]},
        "Action": [
          "kms:Encrypt",
          "kms:Decrypt",
          "kms:ReEncrypt*",
          "kms:GenerateDataKey*",
          "kms:DescribeKey"
        ],
        "Resource":"*"
      }
    ]
  }
  POLICY
}


```

```terraform
module "kms_key" {
  source = "terraform-aws-modules/kms/aws"
  version = "1.3.0"

  description             = "KMS key 1"
  deletion_window_in_days = 7

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Id": "key-default-1",
  "Statement": [
    {
      "Sid": "Allow administration of the key",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::646842065146:root"
      },
      "Action": "kms:*",
      "Resource": "*",
      "Condition": {
        "StringEquals": {
          "kms:CallerAccount": "646842065146",
          "kms:ViaService": "ec2.eu-west-2.amazonaws.com"
        }
      }
    }
  ]
}
EOF
}
```
## Non-Compliant Code Examples
```terraform
resource "aws_kms_key" "positive1" {
  description             = "KMS key 1"
  deletion_window_in_days = 10

  policy = <<POLICY
  {
    "Version": "2012-10-17",
    "Statement":[
      {
        "Sid":"AddCannedAcl",
        "Effect":"Allow",
        "Principal": "*",
        "Action":["kms:*"],
        "Resource":"*"
      }
    ]
  }
  POLICY
}

```

```terraform
resource "aws_kms_key" "positive1" {
  description             = "KMS key 1"
  deletion_window_in_days = 10

  policy = <<POLICY
  {
    "Version": "2012-10-17",
    "Statement":[
      {
        "Sid":"AddCannedAcl",
        "Effect":"Allow",
        "Principal": {"AWS":"*"},
        "Action":["kms:*"],
        "Resource":"*"
      }
    ]
  }
  POLICY
}
```

```terraform
module "kms_key" {
  source = "terraform-aws-modules/kms/aws"
  version = "1.3.0"

  description             = "KMS key 1"
  deletion_window_in_days = 7

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Id": "key-default-1",
  "Statement": [
    {
      "Sid": "Allow administration of the key",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::646842065146:root"
      },
      "Action": "kms:*",
      "Resource": "*"
    }
  ]
}
EOF
}
```