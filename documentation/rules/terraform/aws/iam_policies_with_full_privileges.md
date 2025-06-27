---
title: "IAM Policies With Full Privileges"
meta:
  name: "aws/iam_policies_with_full_privileges"
  id: "2f37c4a3-58b9-4afe-8a87-d7f1d2286f84"
  display_name: "IAM Policies With Full Privileges"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `aws/iam_policies_with_full_privileges`
**Query Name** `IAM Policies With Full Privileges`
**Id:** `2f37c4a3-58b9-4afe-8a87-d7f1d2286f84`
**Cloud Provider:** aws
**Platform** Terraform
**Severity:** Medium
**Category:** Access Control
## Description
IAM policies should never allow full administrative privileges (for all resources), as this is achieved by setting the `"Action"` to `"*"` and the `"Resource"` to `"*"`, as shown below:

```
"Statement": [
  {
    "Effect": "Allow",
    "Action": ["*"],
    "Resource": "*"
  }
]
```

Granting such broad permissions bypasses the principle of least privilege, enabling any user or service with this policy to perform any action on any AWS resource. If left unaddressed, this misconfiguration can lead to privilege escalation, data exfiltration, resource manipulation, or complete account compromise in the event of credential leakage.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy)


## Compliant Code Examples
```terraform
resource "aws_iam_role_policy" "negative1" {
  name = "apigateway-cloudwatch-logging"
  role = aws_iam_role.apigateway_cloudwatch_logging.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": ["some:action"],
      "Resource": "*"
    }
  ]
}
EOF
}
data "aws_iam_policy_document" "example" {
  statement {
    sid = "1"
    effect = "Allow"
    actions = [
      "*"
    ]
    resources = [
      "arn:aws:s3:::*",
    ]
  }
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_iam_role_policy" "positive1" {
  name = "apigateway-cloudwatch-logging"
  role = aws_iam_role.apigateway_cloudwatch_logging.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": ["*"],
      "Resource": "*"
    }
  ]
}
EOF
}

data "aws_iam_policy_document" "example" {
  statement {
    sid = "1"
    effect = "Allow"
    actions = [
      "*"
    ]
    resources = [
      "*",
    ]
  }
}

```