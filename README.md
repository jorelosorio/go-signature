# Signatures playground

This project is considered as a test for packing and signing messages that may be sent over to a -third party- to validate the authenticity and trust the content inside.

> NOTE: This test does not cover sending the data to the third party, both roles are assumed in this solution.

## Tools

- Go 1.17
- Docker
- Visual Studio Code
    > It requires a `Remote - Containers` extension. for more information please refers to: https://code.visualstudio.com/docs/remote/containers#_getting-started

## Data structure

Message data will contains the following attributes:

```
sender (String)
payload (String)
timestamp (Timestamp)
```