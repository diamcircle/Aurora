---
title: Bad Request
---

If Aurora cannot understand a request due to invalid parameters, it will return a `bad_request`
error. This is analogous to the
[HTTP 400 Error](https://developer.mozilla.org/en-US/docs/Web/HTTP/Response_codes).

If you are encountering this error, check the `invalid_field` attribute on the `extras` object to
see what field is triggering the error.

## Attributes

As with all errors Aurora returns, `bad_request` follows the
[Problem Details for HTTP APIs](https://tools.ietf.org/html/draft-ietf-appsawg-http-problem-00)
draft specification guide and thus has the following attributes:

| Attribute   | Type   | Description                                                                     |
| ----------- | ------ | ------------------------------------------------------------------------------- |
| `type`      | URL    | The identifier for the error.  This is a URL that can be visited in the browser.|
| `title`     | String | A short title describing the error.                                             |
| `status`    | Number | An HTTP status code that maps to the error.                                     |
| `detail`    | String | A more detailed description of the error.                                       |

## Example

```shell
$ curl -X GET "https://aurora-testnet.hcnet.org/ledgers?limit=invalidlimit"
{
  "type": "https://hcnet.org/aurora-errors/bad_request",
  "title": "Bad Request",
  "status": 400,
  "detail": "The request you sent was invalid in some way",
  "extras": {
    "invalid_field": "limit",
    "reason": "unparseable value"
  }
}
```

## Related

- [Malformed Transaction](./transaction-malformed.md)