---
title: "NET_RAW Capabilities Disabled for PSP"
meta:
  name: "kubernetes/net_raw_capabilities_disabled_for_psp"
  id: "9aa32890-ac1a-45ee-81ca-5164e2098556"
  cloud_provider: "kubernetes"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `kubernetes/net_raw_capabilities_disabled_for_psp`

**Id:** `9aa32890-ac1a-45ee-81ca-5164e2098556`

**Cloud Provider:** kubernetes

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Containers need to have NET_RAW or All as drop capabilities

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/pod_security_policy#required_drop_capabilities)

## Non-Compliant Code Examples
```terraform
resource "kubernetes_pod_security_policy" "example" {
  metadata {
    name = "terraform-example"
  }
  spec {
    privileged                 = false
    allow_privilege_escalation = false

    volumes = [
      "configMap",
      "emptyDir",
      "projected",
      "secret",
      "downwardAPI",
      "persistentVolumeClaim",
    ]
    required_drop_capabilities = [
      "KILL",
      "SYS_TIME",
    ]

    run_as_user {
      rule = "MustRunAsNonRoot"
    }

    se_linux {
      rule = "RunAsAny"
    }

    supplemental_groups {
      rule = "MustRunAs"
      range {
        min = 1
        max = 65535
      }
    }

    fs_group {
      rule = "MustRunAs"
      range {
        min = 1
        max = 65535
      }
    }

    read_only_root_filesystem = true
  }
}

```

## Compliant Code Examples
```terraform
resource "kubernetes_pod_security_policy" "example" {
  metadata {
    name = "terraform-example"
  }
  spec {
    privileged                 = false
    allow_privilege_escalation = false

    volumes = [
      "configMap",
      "emptyDir",
      "projected",
      "secret",
      "downwardAPI",
      "persistentVolumeClaim",
    ]
    required_drop_capabilities = [
      "ALL"
    ]

    run_as_user {
      rule = "MustRunAsNonRoot"
    }

    se_linux {
      rule = "RunAsAny"
    }

    supplemental_groups {
      rule = "MustRunAs"
      range {
        min = 1
        max = 65535
      }
    }

    fs_group {
      rule = "MustRunAs"
      range {
        min = 1
        max = 65535
      }
    }

    read_only_root_filesystem = true
  }
}

```