# supersecret Terraform Provider

The `supersecret` provider is a hello-world example demonstrating the private-source / public-release provider publishing pattern.

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- Go >= 1.21 (only needed for development)

## Using the Provider

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
```

## Resources

| Resource | Description |
|---|---|
| `supersecret_hello` | Creates a greeting from a name. |

### Example

```hcl
resource "supersecret_hello" "example" {
  name = "world"
}

output "greeting" {
  value = supersecret_hello.example.greeting
  # => "Hello, world!"
}
```

## Documentation

Full provider documentation is available on the [Terraform Registry](https://registry.terraform.io/providers/l2fprod/supersecret/latest/docs).
