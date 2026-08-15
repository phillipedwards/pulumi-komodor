# Komodor Resource Provider

The Komodor Resource Provider lets you manage [Komodor](https://komodor.com) resources with Pulumi.

It is a [Pulumi Terraform Bridge](https://github.com/pulumi/pulumi-terraform-bridge) provider wrapping the
[Komodor Terraform provider](https://github.com/komodorio/terraform-provider-komodor). This is a community
project — it is not published to the Pulumi Registry, npm, PyPI, or NuGet, and it is not an official
Komodor or Pulumi product.

## Installing

Because this provider is distributed only via GitHub Releases, you generate a local SDK for your project
rather than installing a published package. This works for every supported language.

First, install the provider plugin from this repo's releases:

```bash
pulumi plugin install resource komodor 0.1.0 \
  --server github://api.github.com/phillipedwards/pulumi-komodor
```

Then, from inside your Pulumi project, generate and wire up the SDK:

```bash
pulumi package add komodor@0.1.0
```

This writes the SDK into your project, records the dependency in `Pulumi.yaml`, and prints the import
path to use. Subsequent `pulumi up` runs fetch the plugin automatically — the release download URL is
baked into the generated SDK, so collaborators and CI do not need the `--server` flag.

The generated package is named per language:

| Language | Package | Import |
|---|---|---|
| Node.js | `@phillipedwards/pulumi-komodor` | `import * as komodor from "@phillipedwards/pulumi-komodor";` |
| Python | `pulumi_komodor` | `import pulumi_komodor as komodor` |
| Go | `github.com/phillipedwards/pulumi-komodor/sdk/go/komodor` | `import ".../sdk/go/komodor"` |
| .NET | `Pulumi.Komodor` | `using Pulumi.Komodor;` |

### Go

Go users can skip `pulumi package add` and depend on the committed SDK module directly:

```bash
go get github.com/phillipedwards/pulumi-komodor/sdk@v0.1.0
```

The plugin still has to be installed with the `pulumi plugin install` command above, or be resolvable
from the release URL embedded in the SDK.

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

- `komodor:apiKey` — the API key used to authenticate with Komodor. Can also be supplied via the
  `KOMODOR_API_KEY` or `KOMODOR_TOKEN` environment variables.

Set it as a secret:

```bash
pulumi config set --secret komodor:apiKey <your-api-key>
```

## Example

```typescript
import * as komodor from "@phillipedwards/pulumi-komodor";

const viewer = new komodor.Role("viewer", {
    name: "viewer",
});
```

## Reference

There is no Pulumi Registry documentation for this provider. For the available fields on each resource,
consult the generated SDK in your project, the [provider schema](provider/cmd/pulumi-resource-komodor/schema.json),
or the [upstream Terraform provider docs](https://registry.terraform.io/providers/komodorio/komodor/latest/docs).

## Contributing

See [SETUP.md](SETUP.md) for build, test, and development instructions.
