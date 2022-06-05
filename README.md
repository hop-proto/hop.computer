hop.computer
============

DNS is in Google Domains. David owns the domain in his personal account.

It is currently implemented as a Go program using `go:embed`. The /hop path
needs to return a special meta tag in order for `go.mod` to work correctly.

```html
  <meta name="go-import" content="hop.computer/hop git https://github.com/hop-proto/hop-go" />
```

## Running locally

```cmd
$ go run . -address "localhost:8080"
```

The default listen address is 0.0.0.0:8080 so that we don't have to pass any
CLI args when deploying.

## Deploying

The site runs on fly.io under the Hop organization.

```cmd
$ flyctl auth login
$ flyctl deploy
```
