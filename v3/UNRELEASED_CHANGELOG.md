# Unreleased Changes

<!-- 
This file is used to collect changelog entries for the next v3 alpha release.
Add your changes under the appropriate sections below.

Guidelines:
- Follow the "Keep a Changelog" format (https://keepachangelog.com/)
- Write clear, concise descriptions of changes
- Include the impact on users when relevant
- Use present tense ("Add feature" not "Added feature")
- Reference issue/PR numbers when applicable

This file is automatically processed by the nightly release workflow.
After processing, the content will be moved to the main changelog and this file will be reset.
-->

## Added
<!-- New features, capabilities, or enhancements -->

## Changed
- Bump `webview2` to v1.0.24.
  - ci(webview2): fix release build (cross-compile Windows + complete go.sum) (#5671)
  - fix(webview2): recover from transient runtime COM errors instead of exiting (#5658)
  - fix(webview2): never treat a failed SetSize/PutBounds as fatal (#5597)
  - feat: add go workspace for cross-module development
  - fix(webview2): guard Focus() until controller initialisation completes (#5568)
  - fix(v3/webview2): use log instead of fmt for error and stack trace output (#5453)
  - chore(webview2): import scripts/ from go-webview2
  - feat: migrate go-webview2 into wails monorepo (#5317)
<!-- Changes in existing functionality -->

## Fixed
<!-- Bug fixes -->

## Deprecated
<!-- Soon-to-be removed features -->

## Removed
<!-- Features removed in this release -->

## Security
<!-- Security-related changes -->

---

### Example Entries:

**Added:**
- Add support for custom window icons in application options
- Add new `SetWindowIcon()` method to runtime API (#1234)

**Changed:**
- Update minimum Go version requirement to 1.21
- Improve error messages for invalid configuration files

**Fixed:**
- Fix memory leak in event system during window close operations (#5678)
- Fix crash when using context menus on Linux with Wayland

**Security:**
- Update dependencies to address CVE-2024-12345 in third-party library
