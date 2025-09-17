---
title: "Unpinned actions full length commit SHA"
group_id: "rules/cicd/github"
meta:
  name: "github/unpinned_actions_full_length_commit_sha"
  id: "555ab8f9-2001-455e-a077-f2d0f41e2fb9"
  display_name: "Unpinned actions full length commit SHA"
  cloud_provider: "github"
  framework: "CICD"
  severity: "LOW"
  category: "Supply-Chain"
---
## Metadata

**Id:** `555ab8f9-2001-455e-a077-f2d0f41e2fb9`

**Cloud Provider:** github

**Framework:** CICD

**Severity:** Low

**Category:** Supply-Chain

#### Learn More

 - [Provider Reference](https://docs.github.com/en/actions/security-guides/security-hardening-for-github-actions#using-third-party-actions)

### Description

 Pinning an action to a full-length commit SHA is currently the only way to use it as an immutable release. This helps mitigate the risk of a bad actor introducing a backdoor, as doing so would require generating a SHA-1 collision for a valid Git object. When choosing a SHA, ensure it comes from the action's original repository and not a fork.
