---
title: "Root Containers Admitted"
meta:
  name: "terraform/root_containers_admitted"
  id: "4c415497-7410-4559-90e8-f2c8ac64ee38"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Best Practices"
---
## Metadata
**Name:** `terraform/root_containers_admitted`
**Id:** `4c415497-7410-4559-90e8-f2c8ac64ee38`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Best Practices
## Description
Containers must not be allowed to run with root privileges, which means the attributes 'privileged' and 'allow_privilege_escalation' must be set to false, 'run_as_user.rule' must be set to 'MustRunAsNonRoot', and adding the root group must be forbidden

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/pod_security_policy#run_as_user)

## Non-Compliant Code Examples
```kubernetes
resource "kubernetes_pod_security_policy" "example" {
  metadata {
    name = "terraform-example"
  }
  spec {
    privileged                 = true
    allow_privilege_escalation = true

    volumes = [
      "configMap",
      "emptyDir",
      "projected",
      "secret",
      "downwardAPI",
      "persistentVolumeClaim",
    ]

    run_as_user {
      rule = "RunAsAny"
    }

    se_linux {
      rule = "RunAsAny"
    }

    supplemental_groups {
      rule = "RunAsAny"
      range {
        min = 1
        max = 65535
      }
    }

    fs_group {
      rule = "MustRunAs"
      range {
        min = 0
        max = 65535
      }
    }
  }
}

```

## Compliant Code Examples
```kubernetes
resource "kubernetes_pod_security_policy" "example2" {
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