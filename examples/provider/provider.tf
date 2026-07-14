terraform {
  required_providers {
    supersecret = {
      source  = "l2fprod/supersecret"
      version = "~> 0.1"
    }
  }
}

provider "supersecret" {}
