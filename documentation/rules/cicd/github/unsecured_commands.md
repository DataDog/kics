---
title: "Unsecured commands"
group_id: "rules/cicd/github"
meta:
  name: "github/unsecured_commands"
  id: "60fd272d-15f4-4d8f-afe4-77d9c6cc0453"
  display_name: "Unsecured commands"
  cloud_provider: "github"
  framework: "CICD"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `60fd272d-15f4-4d8f-afe4-77d9c6cc0453`

**Cloud Provider:** github

**Framework:** CICD

**Severity:** Medium

**Category:** Insecure Configurations

#### Learn More

 - [Provider Reference](https://0xn3va.gitbook.io/cheat-sheets/ci-cd/github/actions#misuse-of-the-events-related-to-incoming-prs)

### Description

 The deprecated `set-env` and `add-path` commands can still be explicitly enabled by setting the `ACTIONS_ALLOW_UNSECURE_COMMANDS` environment variable to true. Depending on how this variable is used, an attacker could potentially modify the system path to run unintended commands, which may lead to arbitrary code execution.
