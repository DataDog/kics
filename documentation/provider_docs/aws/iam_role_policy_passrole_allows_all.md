---
title: "IAM Role Policy passRole Allows All"
meta:
  name: "aws/iam_role_policy_passrole_allows_all"
  id: "e39bee8c-fe54-4a3f-824d-e5e2d1cca40a"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/iam_role_policy_passrole_allows_all`

**Id:** `e39bee8c-fe54-4a3f-824d-e5e2d1cca40a`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
Using the iam:passrole action with wildcards (*) in the resource can be overly permissive because it allows iam:passrole permissions on multiple resources

#### Learn More

 - [Provider Reference](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-reference-policy-checks.html#access-analyzer-reference-policy-checks-security-warning-pass-role-with-star-in-resource)

## Non-Compliant Code Examples
```terraform
resource "aws_iam_role_policy" "test_policy" {
  name = "test_policy"
  role = aws_iam_role.test_role.id

  policy = <<-EOF
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": [
          "iam:passrole"
        ],
        "Effect": "Allow",
        "Resource": "*"
      }
    ]
  }
  EOF
}




```

## Compliant Code Examples
```terraform
resource "aws_iam_role_policy" "test_policy" {
  name = "test_policy"
  role = aws_iam_role.test_role.id

  policy = <<-EOF
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": [
          "iam:passrole"
        ],
        "Effect": "Allow",
        "Resource": "arn:aws:sqs:us-east-2:account-ID-without-hyphens:queue1"
      }
    ]
  }
  EOF
}


```