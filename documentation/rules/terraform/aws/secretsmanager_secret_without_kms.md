---
title: "Secrets Manager secret without KMS"
group_id: "rules/terraform/aws"
meta:
  name: "aws/secretsmanager_secret_without_kms"
  id: "a2f548f2-188c-4fff-b172-e9a6acb216bd"
  display_name: "Secrets Manager secret without KMS"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Encryption"
---
## Metadata

**Id:** `a2f548f2-188c-4fff-b172-e9a6acb216bd`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Encryption

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/secretsmanager_secret#kms_key_id)

### Description

 By default, AWS Secrets Manager encrypts secrets using the default AWS-managed key, which may not provide the desired level of control over key rotation, access policies, or auditability. Without explicitly specifying a `kms_key_id` in your Terraform resource, as shown below, secrets will not use a customer-managed AWS KMS key (CMK) for encryption:

```
resource "aws_secretsmanager_secret" "example" {
  name = "example"
}
```

This misconfiguration can increase the exposure of sensitive data and limit your ability to implement strict access controls. To reduce risk, explicitly provide a `kms_key_id` attribute referencing a CMK:

```
resource "aws_secretsmanager_secret" "example" {
  name       = "example"
  kms_key_id = "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
}
```


## Compliant Code Examples
```terraform
resource "aws_secretsmanager_secret" "example" {
  name = "example"
  kms_key_id = "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
}

```

```terraform
module "complete" {
  source  = "terraform-aws-modules/secrets-manager/aws"
  version = "5.30.1"

  create_role             = true
  role_name               = "complete-example"
  role_requires_mfa       = true
  role_description        = "My awesome role"
  role_permissions_boundary_arn = "arn:aws:iam::aws:policy/AdministratorAccess"
  custom_role_policy_arns = [
    "arn:aws:iam::aws:policy/AdministratorAccess",
  ]
  role_policy_arns = [
    "arn:aws:iam::aws:policy/AdministratorAccess",
  ]

  tags = {
    Role = "role"
  }
  kms_key_id = aws_kms_key.example.arn
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_secretsmanager_secret" "example" {
  name = "example"
}

```

```terraform
module "complete" {
  source  = "terraform-aws-modules/secrets-manager/aws"
  version = "5.30.1"

  create_role             = true
  role_name               = "complete-example"
  role_requires_mfa       = true
  role_description        = "My awesome role"
  role_permissions_boundary_arn = "arn:aws:iam::aws:policy/AdministratorAccess"
  custom_role_policy_arns = [
    "arn:aws:iam::aws:policy/AdministratorAccess",
  ]
  role_policy_arns = [
    "arn:aws:iam::aws:policy/AdministratorAccess",
  ]

  tags = {
    Role = "role"
  }
}

```