generate_hcl "_generated_backend.tf" {
  content {
    terraform {
      backend "azurerm" {
        resource_group_name  = "rg-test"
        storage_account_name = "satest"
        container_name       = "tfstate"
        key                  = "terraform/stacks/by-id/${terramate.stack.id}/terraform.tfstate"
        use_oidc             = true
      }
    }
  }
}
