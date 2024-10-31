# terramate-bootstrap

terramate-bootstrap is a GoLang based CLI tool designed to create terramate projects and stacks.

It allows users to define their project structure and assign tags automatically to defined stacks.

## Installation

To install \ configure terramate-bootstrap, follow the below steps.

1. Clone this repo.
2. Switch into the repo directory.

Within the make file, there are two options for creating the binary.

`make build` : Build the binary into ./bin/terramate-bootstrap.

`make install` : Build the binary to /usr/local/bin/terramate-bootstrap (to allow for global usage).

## YAML Configuration

The recommended method of configuration is using a yaml configuration file.

The allows users to easily view the configuration that has been deployed. It also allows for re-usability of the configuration files for users that setup multiple platforms.

### YAML Example

```YAML
---
backend:
  azurerm:
    resource_group_name: "rg-test"
    storage_account_name: "satfstatetest"
    location: "uksouth"
stacks:
  environments:
    prd:
      tags:
        - "test_tag_prd"
    dev:
      tags:
        - "test_tag_dev"
        - "test_tag_dev2"
  regions:
    uks:
    ukw:
  resource_types:
    frontdoor:
      tags:
        - "frontdoor_tag"
      exclude_environments:
        - dev
    virtual_machines:
      tags:
        - "virtual_machines_tag"
      exclude_environments:
        - prd
```
