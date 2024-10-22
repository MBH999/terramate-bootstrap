generate_hcl "_generated_terraform.tf" {
  content {
    # Terraform version constraints
    terraform {
      required_version = tm_try(global.terraform.version, "~> 1.8")
    }
  }
}
