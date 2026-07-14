# terraform-provider-supersecret-internal

Internal source repository for the `l2fprod/supersecret` Terraform provider.

This provider currently includes one example resource:

- `supersecret_hello`: creates a computed greeting for a supplied name

## Requirements

- Terraform >= 1.0
- Go >= 1.26

## Provider Usage

```hcl
terraform {
  required_providers {
    supersecret = {
      source  = "l2fprod/supersecret"
      version = "~> 0.1"
    }
  }
}

provider "supersecret" {}

resource "supersecret_hello" "example" {
  name = "world"
}

output "greeting" {
  value = supersecret_hello.example.greeting
}
```

## Development

Build and test locally:

```bash
make build
make test
```

Regenerate docs in `docs/`:

```bash
make generate
```

Other useful targets:

```bash
make lint
make snapshot
```

## Project Layout

- `internal/provider/`: provider and resource implementations
- `examples/`: example Terraform configurations
- `docs/`: generated Terraform provider/resource docs
- `public-repo/`: public-facing release artifacts (README/license mirror)
- `release-notes/`: release note files used by release tooling
