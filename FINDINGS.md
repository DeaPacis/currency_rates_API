# Security Findings

## Vulnerability Scan

- The container image was scanned using Trivy. <br>
Report file:
trivy-report.json

- SBOM generated using Syft. <br>
File:
sbom.json

## Findings

No critical vulnerabilities were found in the base image according to the scan files.

## Security Enhancement

1. Multi-stage build used to reduce attack surface.
2. Runtime image uses minimal Alpine.
3. Container runs as non-root user.
4. Read-only filesystem enabled.
5. Linux capabilities dropped.