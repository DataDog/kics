---
title: "SageMaker notebook instance without KMS"
group_id: "rules/terraform/aws"
meta:
  name: "aws/sagemaker_notebook_instance_without_kms"
  id: "f3674e0c-f6be-43fa-b71c-bf346d1aed99"
  display_name: "SageMaker notebook instance without KMS"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata

**Id:** `f3674e0c-f6be-43fa-b71c-bf346d1aed99`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** High

**Category:** Encryption

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sagemaker_notebook_instance#kms_key_id)

### Description

 AWS SageMaker notebook instances should be configured with a KMS key for encryption at rest to protect sensitive data and machine learning artifacts. Without proper encryption, confidential information stored in these notebooks may be exposed to unauthorized access, potentially leading to data breaches and compliance violations. To secure your SageMaker notebook instance, specify the `kms_key_id` attribute in your Terraform configuration, as shown below:

```
resource "aws_sagemaker_notebook_instance" "ni" {
  name          = "my-notebook-instance"
  role_arn      = aws_iam_role.role.arn
  instance_type = "ml.t2.medium"
  kms_key_id    = "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
}
```


## Compliant Code Examples
```terraform
resource "aws_sagemaker_notebook_instance" "ni" {
  name          = "my-notebook-instance"
  role_arn      = aws_iam_role.role.arn
  instance_type = "ml.t2.medium"
  kms_key_id = "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"

  tags = {
    Name = "foo"
  }
}

```

```terraform
module "notebook" {
  source = "github.com/SebastianUA/terraform-aws-sagemaker"
  version = "~> 2.0"

  name      = "notebook"
  role_arn  = module.notebook_role.iam_role_arn
  subnet_id = module.vpc.private_subnets[0]

  instance_type = "ml.t2.medium"
  kms_key_id        = aws_kms_key.this.id
  root_access = "Enabled"

  policy_arns = [
    "arn:aws:iam::aws:policy/AmazonSageMakerFullAccess"
  ]

  lifecycle_config {
    jupyter_server_app_settings_lifecycle_config_arn = aws_sagemaker_app_image_config.jupyter.arn
  }

  vpc_security_group_ids = [
    module.security_group.security_group_id
  ]

  tags = {
    Owner   = "user"
    Project = "foo"
  }

  depends_on = [module.vpc]

  instance_metadata_service_configuration {
    minimum_instance_metadata_service_version = "2"
  }
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_sagemaker_notebook_instance" "ni" {
  name          = "my-notebook-instance"
  role_arn      = aws_iam_role.role.arn
  instance_type = "ml.t2.medium"

  tags = {
    Name = "foo"
  }
}

```

```terraform
module "notebook" {
  source = "github.com/SebastianUA/terraform-aws-sagemaker"
  version = "~> 2.0"

  name      = "notebook"
  role_arn  = module.notebook_role.iam_role_arn
  subnet_id = module.vpc.private_subnets[0]

  instance_type = "ml.t2.medium"
  # kms_key_id        = aws_kms_key.this.id
  root_access = "Enabled"

  policy_arns = [
    "arn:aws:iam::aws:policy/AmazonSageMakerFullAccess"
  ]

  lifecycle_config {
    jupyter_server_app_settings_lifecycle_config_arn = aws_sagemaker_app_image_config.jupyter.arn
  }

  vpc_security_group_ids = [
    module.security_group.security_group_id
  ]

  tags = {
    Owner   = "user"
    Project = "foo"
  }

  depends_on = [module.vpc]

  instance_metadata_service_configuration {
    minimum_instance_metadata_service_version = "2"
  }
}

```