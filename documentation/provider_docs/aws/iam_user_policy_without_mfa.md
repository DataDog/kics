---
title: "IAM User Policy Without MFA"
meta:
  name: "aws/iam_user_policy_without_mfa"
  id: "b5681959-6c09-4f55-b42b-c40fa12d03ec"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/iam_user_policy_without_mfa`

**Id:** `b5681959-6c09-4f55-b42b-c40fa12d03ec`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Insecure Configurations

## Description
Check if the root user is authenticated with MFA

#### Learn More

 - [Provider Reference](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_mfa_configure-api-require.html)

## Non-Compliant Code Examples
```terraform
resource "aws_iam_user" "positive1" {
  name = "root"
  path = "/system/"

  tags = {
    tag-key = "tag-value"
  }
}

resource "aws_iam_access_key" "positive2" {
  user = aws_iam_user.lb.name
}

resource "aws_iam_user_policy" "positive3" {
  name = "test"
  user = aws_iam_user.lb.name

  policy = <<EOF
{
   "Version": "2012-10-17",
   "Statement": [
     {
       "Effect": "Allow",
       "Principal": {
         "AWS": "arn:aws:iam::111122223333:root"
       },
       "Action": "sts:AssumeRole"
     }
   ]
}
EOF
}
```

## Compliant Code Examples
```terraform
resource "aws_iam_user" "negative1" {
  name = "root"
  path = "/system/"

  tags = {
    tag-key = "tag-value"
  }
}

resource "aws_iam_access_key" "negative2" {
  user = aws_iam_user.lb.name
}

resource "aws_iam_user_policy" "negative3" {
  name = "test"
  user = aws_iam_user.lb.name

  policy = <<EOF
{
   "Version": "2012-10-17",
   "Statement": [
     {
       "Effect": "Allow",
       "Principal": {
         "AWS": "arn:aws:iam::111122223333:root"
       },
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

resource "aws_iam_user_policy" "negative4" {
  name = "test"
  user = aws_iam_user.lb.name

  policy = <<EOF
{
   "Version": "2012-10-17",
   "Statement": [
     {
       "Effect": "Allow",
       "Principal": {
         "AWS": "arn:aws:iam::mfa/111122223333:root"
       },
       "Action": "sts:AssumeRole"
     }
   ]
}
EOF
}
```