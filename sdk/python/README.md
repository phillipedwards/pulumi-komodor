# Komodor Resource Provider

The Komodor Resource Provider lets you manage [Komodor](https://komodor.com) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumi/komodor
```

or `yarn`:

```bash
yarn add @pulumi/komodor
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_komodor
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/phillipedwards/pulumi-komodor/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Komodor
```

## Resources

This provider exposes the following resources:

- `komodor.Action`
- `komodor.Kubernetes`
- `komodor.Monitor`
- `komodor.Policy`
- `komodor.PolicyRoleAttachment`
- `komodor.Role`

And the following data sources:

- `komodor.getKubernetes`
- `komodor.getPolicy`
- `komodor.getRole`

## Configuration

The following configuration points are available for the `komodor` provider:

- `komodor:apiKey` (environment: `KOMODOR_API_KEY`) - the API key used to authenticate with Komodor

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/komodor/api-docs/).
