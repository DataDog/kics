---
title: "Secure Ciphers Disabled"
meta:
  name: "terraform/secure_ciphers_disabled"
  id: "5c0003fb-9aa0-42c1-9da3-eb0e332bef21"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Encryption"
---
## Metadata
**Name:** `terraform/secure_ciphers_disabled`
**Id:** `5c0003fb-9aa0-42c1-9da3-eb0e332bef21`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Encryption
## Description
This check verifies whether Amazon CloudFront distributions are configured to use secure TLS protocols and ciphers by examining the `viewer_certificate` block and its `minimum_protocol_version` attribute. CloudFront distributions that specify weak or outdated protocol versions, such as `SSLv3` (e.g., `minimum_protocol_version = "SSLv3"`), expose transmitted data to vulnerabilities like man-in-the-middle attacks and eavesdropping, because these protocols have known security flaws and are no longer considered safe. Failing to enforce secure cipher suites can allow attackers to decrypt, tamper with, or intercept sensitive information in transit between CloudFront and clients. To mitigate this risk, the attribute should be set to a secure value (such as "TLSv1.2_2019" or higher), or by using the default CloudFront certificate which enforces modern standards.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudfront_distribution)

## Non-Compliant Code Examples
```aws
#this is a problematic code where the query should report a result(s)
resource "aws_cloudfront_distribution" "positive1" {
  origin {
    domain_name = "mybucket"
    origin_id   = "myS3Origin"

    s3_origin_config {
      origin_access_identity = "origin-access-identity/cloudfront/ABCDEFG1234567"
    }
  }

  enabled             = true

  default_cache_behavior {
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "myS3Origin"

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "allow-all"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  restrictions {
    geo_restriction {
      restriction_type = "whitelist"
      locations        = ["US", "CA", "GB", "DE"]
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = false
    minimum_protocol_version = "SSLv3"
  }
}
```

## Compliant Code Examples
```aws
#this code is a correct code for which the query should not find any result
resource "aws_cloudfront_distribution" "negative1" {
  origin {
    domain_name = "mybucket"
    origin_id   = "myS3Origin"

    s3_origin_config {
      origin_access_identity = "origin-access-identity/cloudfront/ABCDEFG1234567"
    }
  }

  enabled             = true

  default_cache_behavior {
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "myS3Origin"

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "allow-all"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  restrictions {
    geo_restriction {
      restriction_type = "whitelist"
      locations        = ["US", "CA", "GB", "DE"]
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }
}
```

```aws
#this code is a correct code for which the query should not find any result
resource "aws_cloudfront_distribution" "negative1" {
  origin {
    domain_name = "mybucket"
    origin_id   = "myS3Origin"

    s3_origin_config {
      origin_access_identity = "origin-access-identity/cloudfront/ABCDEFG1234567"
    }
  }

  enabled = true

  default_cache_behavior {
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "myS3Origin"

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "allow-all"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  restrictions {
    geo_restriction {
      restriction_type = "whitelist"
      locations        = ["US", "CA", "GB", "DE"]
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = true
    minimum_protocol_version       = "TLSv1.2_2019"
  }
}

```