---
title: "DNSSEC Using RSASHA1"
meta:
  name: "gcp/dnssec_using_rsasha1"
  id: "ccc3100c-0fdd-4a5e-9908-c10107291860"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Encryption"
---

## Metadata
**Name:** `gcp/dnssec_using_rsasha1`

**Id:** `ccc3100c-0fdd-4a5e-9908-c10107291860`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Encryption

## Description
DNSSEC should not use the RSASHA1 algorithm, which means if, within the 'dnssec_config' block, the 'default_key_specs' block exists with the 'algorithm' field is 'rsasha1' which is bad.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/dns_managed_zone#algorithm)

## Non-Compliant Code Examples
```terraform
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
```terraform
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