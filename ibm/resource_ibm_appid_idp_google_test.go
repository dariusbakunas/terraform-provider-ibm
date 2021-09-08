package ibm

import (
	"fmt"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccAppIDIDPGoogle_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAppIDIDPGoogleDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupIBMAppIDGoogleIDPConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_idp_google.gg", "tenant_id", appIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_idp_google.gg", "config.0.application_id", "test_id"),
					resource.TestCheckResourceAttr("ibm_appid_idp_google.gg", "config.0.application_secret", "test_secret"),
					resource.TestCheckResourceAttrSet("ibm_appid_idp_google.gg", "redirect_url"),
				),
			},
		},
	})
}

func setupIBMAppIDGoogleIDPConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_idp_google" "gg" {
			tenant_id = "%s"
			is_active = true
			
			config {
				application_id 		= "test_id"
				application_secret 	= "test_secret"
			}
		}
	`, tenantID)
}

func testAccCheckIBMAppIDIDPGoogleDestroy(s *terraform.State) error {
	appIDClient, err := testAccProvider.Meta().(ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_idp_google" {
			continue
		}

		tenantID := rs.Primary.ID

		config, _, err := appIDClient.GetGoogleIDP(&appid.GetGoogleIDPOptions{
			TenantID: &tenantID,
		})

		if err != nil {
			return fmt.Errorf("Error checking if AppID Google IDP configuration was reset: %s", err)
		}

		if config == nil || (config.IsActive != nil && *config.IsActive != false) {
			return fmt.Errorf("Error checking if AppID Google IDP configuration was reset")
		}
	}

	return nil
}
