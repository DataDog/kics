---
title: "DNSSEC Using RSASHA1"
meta:
  name: "terraform/dnssec_using_rsasha1"
  id: "ccc3100c-0fdd-4a5e-9908-c10107291860"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Encryption"
---
## Metadata
**Name:** `terraform/dnssec_using_rsasha1`
**Id:** `ccc3100c-0fdd-4a5e-9908-c10107291860`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Encryption
## Description
DNSSEC should not use the RSASHA1 algorithm, which is considered weak and vulnerable to cryptographic attacks. If a `dnssec_config` block contains a `default_key_specs` attribute with `algorithm = "rsasha1"`, attackers may be able to exploit known weaknesses in this algorithm to forge DNS records, potentially redirecting users to malicious sites or causing other security issues. Instead, use a stronger algorithm such as `rsasha256`:

```
dnssec_config {
    default_key_specs {
        algorithm = "rsasha256"
    }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/dns_managed_zone#algorithm)

## Non-Compliant Code Examples
```gcp
resource "google_dns_managed_zone" "positive1" {
    name        = "example-zone"
    dns_name    = "example-${random_id.rnd.hex}.com."
    description = "Example DNS zone"
    labels = {
        foo = "bar"
    }

    dnssec_config {
        default_key_specs{
            algorithm = "rsasha1"
        }
    }
}

```

## Compliant Code Examples
```gcp
resource "google_dns_managed_zone" "negative1" {
    name        = "example-zone"
    dns_name    = "example-${random_id.rnd.hex}.com."
    description = "Example DNS zone"
    labels = {
        foo = "bar"
    }

    dnssec_config {
        default_key_specs{
            algorithm = "rsasha256"
        }
    }
}


```