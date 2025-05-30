---
title: "CodeBuild Project Encrypted With AWS Managed Key"
meta:
  name: "aws/codebuild_project_encrypted_with_aws_managed_key"
  id: "3deec14b-03d2-4d27-9670-7d79322e3340"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Encryption"
---

## Metadata
**Name:** `aws/codebuild_project_encrypted_with_aws_managed_key`

**Id:** `3deec14b-03d2-4d27-9670-7d79322e3340`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Encryption

## Description
CodeBuild Project should be encrypted with customer-managed KMS keys instead of AWS managed keys

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/codebuild_project#encryption_key)

## Non-Compliant Code Examples
```terraform
provider "aws" {
  region = "us-east-1"
}

data "aws_kms_key" "by_alias" {
  key_id = "alias/aws/s3"
}

# No policy attached to this role because it is for testing purposes
resource "aws_iam_role" "codebuild" {
  name = "codebuild-cloudrail-test"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "codebuild.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_codebuild_project" "project-cloudrail-test" {
  name           = "project-cloudrail-test"
  description    = "project-cloudrail-test"
  build_timeout  = "5"
  queued_timeout = "5"
  service_role   = aws_iam_role.codebuild.arn
  encryption_key = data.aws_kms_key.by_alias.arn

  artifacts {
    type = "NO_ARTIFACTS"
  }

  cache {
    type  = "LOCAL"
    modes = ["LOCAL_DOCKER_LAYER_CACHE", "LOCAL_SOURCE_CACHE"]
  }

  environment {
    compute_type                = "BUILD_GENERAL1_SMALL"
    image                       = "aws/codebuild/standard:1.0"
    type                        = "LINUX_CONTAINER"
    image_pull_credentials_type = "CODEBUILD"

    environment_variable {
      name  = "SOME_KEY1"
      value = "SOME_VALUE1"
    }
  }

  source {
    type            = "GITHUB"
    location        = "https://github.com/foo/bar.git"
    git_clone_depth = 1
  }
}

```

## Compliant Code Examples
```terraform
provider "aws" {
  region = "us-east-1"
}

data "aws_kms_key" "by_alias2" {
  key_id = "alias/myAlias"
}

# No policy attached to this role because it is for testing purposes
resource "aws_iam_role" "codebuild2" {
  name = "codebuild-cloudrail-test"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "codebuild.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_codebuild_project" "project-cloudrail-test2" {
  name           = "project-cloudrail-test"
  description    = "project-cloudrail-test"
  build_timeout  = "5"
  queued_timeout = "5"
  service_role   = aws_iam_role.codebuild2.arn
  encryption_key = data.aws_kms_key.by_alias2.arn

  artifacts {
    type = "NO_ARTIFACTS"
  }

  cache {
    type  = "LOCAL"
    modes = ["LOCAL_DOCKER_LAYER_CACHE", "LOCAL_SOURCE_CACHE"]
  }

  environment {
    compute_type                = "BUILD_GENERAL1_SMALL"
    image                       = "aws/codebuild/standard:1.0"
    type                        = "LINUX_CONTAINER"
    image_pull_credentials_type = "CODEBUILD"

    environment_variable {
      name  = "SOME_KEY1"
      value = "SOME_VALUE1"
    }
  }

  source {
    type            = "GITHUB"
    location        = "https://github.com/foo/bar.git"
    git_clone_depth = 1
  }
}

```