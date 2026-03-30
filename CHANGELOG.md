# v1.1.0
## Chores
- Bump oauth2 from 0.21.0 to 0.27.0 (requires Go >= 1.23.0)
## Fixes
- Fixes an issue where the OAuth token source did not use the provided CA trust bundle when fetching tokens, which could lead to TLS errors in environments with custom CAs.

# v1.0.2
## Chores
- Bump oauth2 from 0.20.0 to 0.21.0

# v1.0.1
## Chores
- Rename `GetHttpClient` to `GetHTTPClient`

# v1.0.0
- Initial release of the EJBCA Go Client SDK, providing a client for interacting with EJBCA's REST API.