package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	// providerConfig is a shared configuration to combine with the actual
	// test configuration so the client is properly configured.
	// It is also possible to use environment variables instead,
	// such as updating the Makefile and running the testing through that tool.
	providerConfig = `
provider "waypoint" {}

`
)

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"waypoint": providerserver.NewProtocol6WithError(New("test")()),
	}
)

// not completely needed but following the example set in
// hashicorp/terraform-provider-scaffolding-framework
func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}
