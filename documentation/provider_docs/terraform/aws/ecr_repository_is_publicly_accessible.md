---
title: "ECR Repository Is Publicly Accessible"
meta:
  name: "aws/ecr_repository_is_publicly_accessible"
  id: "e86e26fc-489e-44f0-9bcd-97305e4ba69a"
  cloud_provider: "aws"
  severity: "CRITICAL"
  category: "Access Control"
---
## Metadata
**Name:** `aws/ecr_repository_is_publicly_accessible`
**Id:** `e86e26fc-489e-44f0-9bcd-97305e4ba69a`
**Cloud Provider:** aws
**Severity:** Critical
**Category:** Access Control
## Description
Amazon ECR repository policies that use wildcard '*' in the Principal field grant public access to your container images, potentially exposing sensitive data or proprietary code. When left unaddressed, this vulnerability allows unauthorized users to pull, and in some cases push, container images, compromising the integrity and confidentiality of your container deployments. To remediate this issue, always specify explicit IAM principals in your repository policies, such as `"Principal": { "AWS": "arn:aws:iam::account_number:root" }` instead of using `"Principal": "*"`.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecr_repository_policy)


## Compliant Code Examples
```terraform
resource "aws_ecr_repository" "negative1" {
  name = "bar"
}

resource "aws_ecr_repository_policy" "negative2" {
  repository = aws_ecr_repository.foo.name

  policy = <<EOF
{
    "Version": "2008-10-17",
    "Statement": [
        {
            "Sid": "new policy",
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::##account_number##:root"
            },
            "Action": [
                "ecr:GetDownloadUrlForLayer",
                "ecr:BatchGetImage",
                "ecr:BatchCheckLayerAvailability",
                "ecr:PutImage",
                "ecr:InitiateLayerUpload",
                "ecr:UploadLayerPart",
                "ecr:CompleteLayerUpload",
                "ecr:DescribeRepositories",
                "ecr:GetRepositoryPolicy",
                "ecr:ListImages",
                "ecr:DeleteRepository",
                "ecr:BatchDeleteImage",
                "ecr:SetRepositoryPolicy",
                "ecr:DeleteRepositoryPolicy"
            ]
        }
    ]
}
EOF
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_ecr_repository" "positive1" {
  name = "bar"
}

resource "aws_ecr_repository_policy" "positive2" {
  repository = aws_ecr_repository.foo.name

  policy = <<EOF
{
    "Version": "2008-10-17",
    "Statement": [
        {
            "Sid": "new policy",
            "Effect": "Allow",
            "Principal": "*",
            "Action": [
                "ecr:GetDownloadUrlForLayer",
                "ecr:BatchGetImage",
                "ecr:BatchCheckLayerAvailability",
                "ecr:PutImage",
                "ecr:InitiateLayerUpload",
                "ecr:UploadLayerPart",
                "ecr:CompleteLayerUpload",
                "ecr:DescribeRepositories",
                "ecr:GetRepositoryPolicy",
                "ecr:ListImages",
                "ecr:DeleteRepository",
                "ecr:BatchDeleteImage",
                "ecr:SetRepositoryPolicy",
                "ecr:DeleteRepositoryPolicy"
            ]
        }
    ]
}
EOF
}

```