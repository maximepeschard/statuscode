# statuscode

A httpstatuses.com copycat as an API serving JSON, written in Go.

```console
$ curl -s "localhost:8080/503" | jq
{
  "code": 503,
  "phrase": "Service Unavailable",
  "class": "5XX Server Error",
  "description": "The 503 (Service Unavailable) status code indicates that the server is currently unable to handle the request due to a temporary overload or scheduled maintenance, which will likely be alleviated after some delay. The server MAY send a Retry-After header field to suggest an appropriate amount of time for the client to wait before retrying the request.",
  "reference": "https://tools.ietf.org/html/rfc7231#section-6.6.4"
}
```

```console
$ curl -s "localhost:8080/?class=1XX" | jq
[
  {
    "code": 103,
    "phrase": "Early Hints",
    "class": "1XX Informational",
    "description": "The 103 (Early Hints) informational status code indicates to the client that the server is likely to send a final response with the header fields included in the informational response.",
    "reference": "https://tools.ietf.org/html/rfc8297#section-2"
  },
  {
    "code": 101,
    "phrase": "Switching Protocols",
    "class": "1XX Informational",
    "description": "The 101 (Switching Protocols) status code indicates that the server understands and is willing to comply with the client's request, via the Upgrade header field, for a change in the application protocol being used on this connection. The server MUST generate an Upgrade header field in the response that indicates which protocol(s) will be switched to immediately after the empty line that terminates the 101 response.",
    "reference": "https://tools.ietf.org/html/rfc7231#section-6.2.2"
  },
  ...
]
```