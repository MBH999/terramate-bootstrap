generate_hcl "_generated_backend.tf" {
  content {
    terraform {
      backend "azurerm" {
        resource_group_name  = global.backend.config.resource_group_name
        storage_account_name = global.backend.config.storage_account_name
        container_name       = "tfstate"
        key                  = "terraform/stacks/by-id/${terramate.stack.id}/terraform.tfstate"
        use_oidc             = true
      }
    }
  }
}
