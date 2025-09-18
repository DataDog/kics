---
title: "Script block injection"
group_id: "rules/cicd/github"
meta:
  name: "github/script_block_injection"
  id: "62ff6823-927a-427f-acf9-f1ea2932d616"
  display_name: "Script block injection"
  cloud_provider: "github"
  framework: "CICD"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `62ff6823-927a-427f-acf9-f1ea2932d616`

**Cloud Provider:** github

**Framework:** CICD

**Severity:** High

**Category:** Insecure Configurations

#### Learn More

 - [Provider Reference](https://securitylab.github.com/research/github-actions-untrusted-input/)

### Description

 GitHub Actions workflows can be triggered by a variety of events. Each trigger includes a GitHub context with details about the event, such as the user who triggered it, the branch name, and other relevant information. Some of this data, like the base repository name, changeset hash, or pull request number, is typically not controlled by the user and is unlikely to be used for injection.
