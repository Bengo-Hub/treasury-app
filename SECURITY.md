# Security Policy

## Supported Versions

| Version | Supported |
| ------- | --------- |
| `main`  | ✅ |
| Latest release | ✅ |

## Reporting a Vulnerability

Please email [security@bengobox.com](mailto:security@bengobox.com) with subject `SECURITY: Treasury Service`. Include:

- Description of the vulnerability and potential impact
- Steps to reproduce (proof of concept preferred)
- Affected versions/commits
- Suggested remediation if available

We will acknowledge within 48 hours and provide an initial assessment within 5 business days.

## Responsible Disclosure

- Do not disclose vulnerability details publicly before a fix is available
- Avoid accessing or modifying data that is not yours
- Do not perform tests that degrade system availability or integrity

## Patch Process

1. Triage and assign severity/CVSS score
2. Develop fix on private branch, reviewed by security team
3. Deploy via CI/CD and ArgoCD
4. Publish advisory in [`CHANGELOG.md`](CHANGELOG.md) and internal security channels

Thank you for helping us keep our customers safe.
