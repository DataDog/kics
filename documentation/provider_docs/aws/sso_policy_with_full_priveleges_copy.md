---
title: "SSO Identity User Unsafe Creation"
meta:
  name: "aws/sso_policy_with_full_priveleges_copy"
  id: "4003118b-046b-4640-b200-b8c7a4c8b89f"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/sso_policy_with_full_priveleges_copy`

**Id:** `4003118b-046b-4640-b200-b8c7a4c8b89f`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
The use of AWS SSO for creating users may pose a security risk as it does not synchronize with external Identity Providers (IdP) or Active Directory (AD). This can lead to inconsistencies and potential unauthorized access to resources. It is recommended to review and update user creation processes to ensure proper security protocols are in place.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/identitystore_user)

## Non-Compliant Code Examples
```terraform
resource "aws_identitystore_user" "example" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.example.identity_store_ids)[0]

  display_name = "John Doe"
  user_name    = "johndoe"

  name {
    given_name  = "John"
    family_name = "Doe"
  }

  emails {
    value = "john@example.com"
  }
}

```

## Compliant Code Examples
```terraform
resource "aws_ssoadmin_permission_set_inline_policy" "neg1" {
  instance_arn       = aws_ssoadmin_permission_set.example.instance_arn
  permission_set_arn = aws_ssoadmin_permission_set.example.arn
  inline_policy = <<POLICY
{
  "Statement": [
    {
      "Action": [
        "s3:ListBucket*",
        "s3:HeadBucket",
        "s3:Get*"
      ],
      "Effect": "Allow",
      "Resource": [
        "arn:aws:s3:::b1",
        "arn:aws:s3:::b1/*",
        "arn:aws:s3:::b2",
        "arn:aws:s3:::b2/*"
      ],
      "Sid": ""
    },
    {
      "Action": "s3:PutObject*",
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::b1/*",
      "Sid": ""
    }
  ],
  "Version": "2012-10-17"
}
POLICY
}

```