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

resource "keycloak_user" "test_user" {
  realm_id = keycloak_realm.master.id
  username = "mendako"
  enabled = true

  email = "Mendako@sample.email.address.io"
  first_name = "Opisthoteuthis"
  last_name = "Depressa"
}

resource "keycloak_oidc_identity_provider" "realm_identity_provider" {
  realm             = keycloak_realm.master.id
  alias             = "my-idp"
  authorization_url = "https://authorizationurl.com"
  client_id         = "clientID"
  client_secret     = "clientSecret"
  token_url         = "https://tokenurl.com"
}