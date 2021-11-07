#!/bin/sh

set -e

##############################################################################
# Variables
##############################################################################

if [ -z "${KEYCLOAK_URL}" ]; then
	echo "Require Environment Variables: KEYCLOAK_URL"
	exit 1
fi

##############################################################################
# Wait
##############################################################################

dockerize -timeout 300s -wait "${KEYCLOAK_URL}"

##############################################################################
# Running
##############################################################################

if [ "$1" = 'provision' ]; then
	rm -fr /tmp/terraform
	cp -r /terraform /tmp/terraform
	ln -s /.terraform /tmp/terraform
	cd /tmp/terraform
	terraform init
	terraform import keycloak_realm.master master
	terraform apply -auto-approve
	exit 0
fi

##############################################################################
# Fallback
##############################################################################

exec "$@"