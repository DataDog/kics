---
title: "Authentication Without MFA"
meta:
  name: "aws/authentication_without_mfa"
  id: "3ddfa124-6407-4845-a501-179f90c65097"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Access Control"
---

## Metadata
**Name:** `aws/authentication_without_mfa`

**Id:** `3ddfa124-6407-4845-a501-179f90c65097`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Access Control

## Description
Users should authenticate with MFA (Multi-factor Authentication) to ensure an extra layer of protection when authenticating

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_user_policy)

## Non-Compliant Code Examples
```terraform
provider "aws" {
  region  = "us-east-1"
}

resource "aws_iam_user" "positive1" {
  name = "aws-foundations-benchmark-1-4-0-terraform-user"
  path = "/"
}

resource "aws_iam_user_login_profile" "positive2" {
  user = aws_iam_user.positive2.name
  pgp_key = "gpgkeybase64gpgkeybase64gpgkeybase64gpgkeybase64"
}

resource "aws_iam_user_policy" "positive2" {
  name = "aws-foundations-benchmark-1-4-0-terraform-user"
  user = aws_iam_user.positive2.name

  policy = <<EOF
{
   "Version": "2012-10-17",
   "Statement": [
     {
       "Effect": "Allow",
       "Resource": "${aws_iam_user.positive2.arn}",
       "Action": "sts:AssumeRole"
     }
   ]
}
EOF
}

```

```terraform
provider "aws" {
  region  = "us-east-1"
}

resource "aws_iam_user" "positive1" {
  name = "aws-foundations-benchmark-1-4-0-terraform-user"
  path = "/"
}

resource "aws_iam_user_login_profile" "positive1" {
  user = aws_iam_user.positive1.name
  pgp_key = "gpgkeybase64gpgkeybase64gpgkeybase64gpgkeybase64"
}

resource "aws_iam_access_key" "positive1" {
  user = aws_iam_user.positive1.name
}

resource "aws_iam_user_policy" "positive1" {
  name = "aws-foundations-benchmark-1-4-0-terraform-user"
  user = aws_iam_user.positive1.name

  policy = <<EOF
{
   "Version": "2012-10-17",
   "Statement": [
     {
       "Effect": "Allow",
       "Resource": "${aws_iam_user.positive1.arn}",
       "Action": "sts:AssumeRole",
       "Condition": {
         "BoolIfExists": {
           "aws:MultiFactorAuthPresent" : "false"
         }
       }
     }
   ]
}
EOF
}

```

## Compliant Code Examples
```terraform
provider "aws" {
  region  = "us-east-1"
}

resource "aws_iam_user" "negative1" {
  name = "aws-foundations-benchmark-1-4-0-terraform-user"
  path = "/"
}

resource "aws_iam_user_login_profile" "negative1" {
  user = aws_iam_user.negative1.name
  pgp_key = "gpgkeybase64gpgkeybase64gpgkeybase64gpgkeybase64"
}

resource "aws_iam_access_key" "negative1" {
  user = aws_iam_user.negative1.name
}

resource "aws_iam_user_policy" "negative1" {
  name = "aws-foundations-benchmark-1-4-0-terraform-user"
  user = aws_iam_user.negative1.name

  policy = <<EOF
{
   "Version": "2012-10-17",
   "Statement": [
     {
       "Effect": "Allow",
       "Resource": ${aws_iam_user.negative1.arn},
       "Action": "sts:AssumeRole",
       "Condition": {
         "BoolIfExists": {
           "aws:MultiFactorAuthPresent" : "true"
         }
       }
     }
   ]
}
EOF
}

```