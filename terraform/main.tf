 terraform {
  required_version = "1.0.10"

  required_providers {
    keycloak = {
      source  = "mrparkers/keycloak"
      version = "3.5.1"
    }
  }
}

provider "keycloak" {
  url       = "http://keycloak:8080"
  client_id = "admin-cli"
  username  = "admin"
  password  = "admin"
}

resource "keycloak_realm" "master" {
  realm = "master"

  internationalization {
    supported_locales = [
      "ja",
      "en",
    ]
    default_locale = "ja"
  }
}

resource "keycloak_openid_client" "openid_client" {
  realm_id            = keycloak_realm.master.id
  client_id           = "test-client"

  name                = "test client"
  enabled             = true

  access_type         = "CONFIDENTIAL"

  login_theme = "keycloak"
}