---
title: "Lambda Function Publicly Accessible"
meta:
  name: "terraform/lambda_function_publicly_accessible"
  id: "1btsf2f9-af8c-4dfc-a0f1-a03adb70deb2"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/lambda_function_publicly_accessible`
**Id:** `1btsf2f9-af8c-4dfc-a0f1-a03adb70deb2`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Access Control
## Description
AWS Lambda permissions with a principal value of '*' allow any AWS account or user to invoke the function, making it publicly accessible. This creates a significant security risk as unauthorized parties can trigger your Lambda function, potentially accessing sensitive data or consuming your AWS resources. To secure Lambda functions, specify the exact AWS account ARN in the principal field (e.g., 'arn:aws:iam::123456789012:root') rather than using the wildcard '*'. Compare the secure configuration:
```
principal = "arn:aws:iam::123456789012:root"
```
versus the insecure configuration:
```
principal = "*"
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_function)

## Non-Compliant Code Examples
```aws
resource "aws_lambda_permission" "public_lambda" {
  statement_id  = "AllowPublicAccess"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.my_lambda.function_name
  principal     = "*"
}

```

## Compliant Code Examples
```aws
resource "aws_lambda_permission" "restricted_lambda" {
  statement_id  = "AllowSpecificAccount"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.my_lambda.function_name
  principal     = "arn:aws:iam::123456789012:root"
}

```