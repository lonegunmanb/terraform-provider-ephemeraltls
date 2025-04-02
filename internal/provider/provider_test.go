// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func protoV5ProviderFactories() map[string]func() (tfprotov5.ProviderServer, error) {
	return map[string]func() (tfprotov5.ProviderServer, error){
		"ephemeraltls": providerserver.NewProtocol5WithError(New()),
	}
}

func TestProvider_InvalidProxyConfig(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: protoV5ProviderFactories(),

		Steps: []resource.TestStep{
			{
				Config: `
					provider "ephemeraltls" {
						proxy {
							url = "https://proxy.host.com"
							from_env = true
						}
					}
					ephemeral "ephemeraltls_private_key" "test" {
						algorithm = "ED25519"
					}
				`,
				ExpectError: regexp.MustCompile(`Invalid Attribute Combination`),
			},
			{
				Config: `
					provider "ephemeraltls" {
						proxy {
							username = "user"
						}
					}
					ephemeral "ephemeraltls_private_key" "test" {
						algorithm = "ED25519"
					}
				`,
				ExpectError: regexp.MustCompile(`Invalid Attribute Combination`),
			},
			{
				Config: `
					provider "ephemeraltls" {
						proxy {
							password = "pwd"
						}
					}
					ephemeral "ephemeraltls_private_key" "test" {
						algorithm = "ED25519"
					}
				`,
				ExpectError: regexp.MustCompile(`Invalid Attribute Combination`),
			},
			{
				Config: `
					provider "ephemeraltls" {
						proxy {
							username = "user"
							password = "pwd"
						}
					}
					ephemeral "ephemeraltls_private_key" "test" {
						algorithm = "ED25519"
					}
				`,
				ExpectError: regexp.MustCompile(`Invalid Attribute Combination`),
			},
			{
				Config: `
					provider "ephemeraltls" {
						proxy {
							username = "user"
							from_env = true
						}
					}
					ephemeral "ephemeraltls_private_key" "test" {
						algorithm = "ED25519"
					}
				`,
				ExpectError: regexp.MustCompile(`Invalid Attribute Combination`),
			},
		},
	})
}
