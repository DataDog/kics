---
title: "Certificate RSA Key Bytes Lower Than 256"
meta:
  name: "aws/certificate_rsa_key_bytes_lower_than_256"
  id: "874d68a3-bfbe-4a4b-aaa0-9e74d7da634b"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/certificate_rsa_key_bytes_lower_than_256`

**Id:** `874d68a3-bfbe-4a4b-aaa0-9e74d7da634b`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Insecure Configurations

## Description
The certificate should use a RSA key with a length equal to or higher than 256 bytes

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_rest_api)

## Non-Compliant Code Examples
```terraform
resource "aws_iam_server_certificate" "test_cert2" {
  name             = "some_test_cert"
  certificate_body = file("./rsa1024.pem")
  private_key      = <<EOF
-----BEGIN RSA PRIVATE KEY-----
[......] # cert contents
-----END RSA PRIVATE KEY-----
EOF
}

```

```terraform
resource "aws_api_gateway_domain_name" "example" {
  certificate_body = file("./rsa1024.pem")
  domain_name     = "api.example.com"
}

```

## Compliant Code Examples
```terraform
resource "aws_iam_server_certificate" "test_cert22" {
  name             = "some_test_cert"
  certificate_body = file("./rsa4096.pem")
  private_key      = <<EOF
-----BEGIN RSA PRIVATE KEY-----
[......] # cert contents
-----END RSA PRIVATE KEY-----
EOF
}


```

```terraform
resource "aws_api_gateway_domain_name" "example3" {
  certificate_body = file("./rsa4096.pem")
  domain_name     = "api.example.com"
}

```