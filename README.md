# Terraform Provider: EphemeralTLS

Important notice: This provider is not intended to replace Terraform `tls` provider. It is forked from it, but removed all resources and data sources, only keeping the new `ephemeraltls_private_key` and `ephemeraltls_public_key` ephemeral.

The TLS provider project seems not actively maintained for a while, but I have to use ephemeral tls private key in my project. I've opened my pull request to contribute a new ephemeral resoruce, but I'm not quite sure whether if and when it could be merged, so I forked the `tls` provider, removed all other resources and data sources. Since the ephemeral resource won't be stored in Terraform state, you can use this ephemeral provider and resource in your Terraform config, once the `tls` provider has supported this feature, you can switch back to the original `tls` provider without any configuration drift.

Thanks for HashiCorp `tls` provider team for their great work, and I hope this fork could be merged back to the original `tls` provider in the future.

---

[`ephemeral tls_private_key`](https://registry.terraform.io/providers/hashicorp/tls/latest/docs/ephemeral-resources/private_key) has been merged thanks to HashiCorp, you SHOULD NOT use `ephemeraltls_private_key` anymore but `ephemeral tls_private_key` instead.

---

The TLS provider provides utilities for working with *Transport Layer Security*
keys and certificates. It provides resources that
allow private keys, certificates and certificate requests to be
created as part of a Terraform deployment.

## Documentation, questions and discussions

Official documentation on how to use this provider can be found on the 
[Terraform Registry](https://registry.terraform.io/providers/hashicorp/tls/latest/docs).
In case of specific questions or discussions, please use the
HashiCorp [Terraform Providers Discuss forums](https://discuss.hashicorp.com/c/terraform-providers/31),
in accordance with HashiCorp [Community Guidelines](https://www.hashicorp.com/community-guidelines).

We also provide:

* [Support](.github/SUPPORT.md) page for help when using the provider
* [Contributing](.github/CONTRIBUTING.md) guidelines in case you want to help this project
* [Design](DESIGN.md) documentation to understand the scope and maintenance decisions

The remainder of this document will focus on the development aspects of the provider.

## Compatibility

Compatibility table between this provider, the [Terraform Plugin Protocol](https://www.terraform.io/plugin/how-terraform-works#terraform-plugin-protocol)
version it implements, and Terraform:

| TLS Provider | Terraform Plugin Protocol | Terraform |
|:------------:|:-------------------------:|:---------:|
|   `>= 4.x`   |            `5`            | `>= 0.12` |
|   `>= 3.x`   |            `5`            | `>= 0.12` |
|   `>= 2.x`   |        `4` and `5`        | `<= 0.12` |
|   `>= 0.x`   |            `4`            | `<= 0.11` |

Details can be found querying the [Registry API](https://www.terraform.io/internals/provider-registry-protocol#list-available-versions)
that return all the details about which version are currently available for a particular provider.
[Here](https://registry.terraform.io/v1/providers/hashicorp/tls/versions) are the details for TLS (JSON response).

## Requirements

* [Terraform](https://www.terraform.io/downloads)
* [Go](https://go.dev/doc/install) (1.23)
* [GNU Make](https://www.gnu.org/software/make/)
* [golangci-lint](https://golangci-lint.run/usage/install/#local-installation) (optional)

## Development

### Building

1. `git clone` this repository and `cd` into its directory
2. `make` will trigger the Golang build

The provided `GNUmakefile` defines additional commands generally useful during development,
like for running tests, generating documentation, code formatting and linting.
Taking a look at it's content is recommended.

### Testing

In order to test the provider, you can run

* `make test` to run provider tests
* `make testacc` to run provider acceptance tests

It's important to note that acceptance tests (`testacc`) will actually spawn
`terraform` and the provider. Read more about they work on the
[official page](https://www.terraform.io/plugin/sdkv2/testing/acceptance-tests).

### Generating documentation

This provider uses [terraform-plugin-docs](https://github.com/hashicorp/terraform-plugin-docs/)
to generate documentation and store it in the `docs/` directory.
Once a release is cut, the Terraform Registry will download the documentation from `docs/`
and associate it with the release version. Read more about how this works on the
[official page](https://www.terraform.io/registry/providers/docs).

Use `make generate` to ensure the documentation is regenerated with any changes.

### Using a development build

If [running tests and acceptance tests](#testing) isn't enough, it's possible to set up a local terraform configuration
to use a development builds of the provider. This can be achieved by leveraging the Terraform CLI
[configuration file development overrides](https://www.terraform.io/cli/config/config-file#development-overrides-for-provider-developers).

First, use `make install` to place a fresh development build of the provider in your
[`${GOBIN}`](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies)
(defaults to `${GOPATH}/bin` or `${HOME}/go/bin` if `${GOPATH}` is not set). Repeat
this every time you make changes to the provider locally.

Then, setup your environment following [these instructions](https://www.terraform.io/plugin/debugging#terraform-cli-development-overrides)
to make your local terraform use your local build.

### Testing GitHub Actions

This project uses [GitHub Actions](https://docs.github.com/en/actions/automating-builds-and-tests) to realize its CI.

Sometimes it might be helpful to locally reproduce the behaviour of those actions,
and for this we use [act](https://github.com/nektos/act). Once installed, you can _simulate_ the actions executed
when opening a PR with:

```shell
# List of workflows for the 'pull_request' action
$ act -l pull_request

# Execute the workflows associated with the `pull_request' action 
$ act pull_request
```

## Releasing

The release process is automated via GitHub Actions, and it's defined in the Workflow
[release.yml](./.github/workflows/release.yml).

Each release is cut by pushing a [semantically versioned](https://semver.org/) tag to the default branch.

## License

[Mozilla Public License v2.0](./LICENSE)
