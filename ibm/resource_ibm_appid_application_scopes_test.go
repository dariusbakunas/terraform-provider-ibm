package ibm

import (
	"fmt"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

func TestAccIBMAppIDApplicationScopes_basic(t *testing.T) {
	appName := fmt.Sprintf("tf_testacc_app_scopes_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAppIDApplicationScopesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMApplicationScopesConfig(appIDTenantID, appName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_application_scopes.scopes", "scopes.#", "3"),
					resource.TestCheckResourceAttr("ibm_appid_application_scopes.scopes", "scopes.0", "scope1"),
					resource.TestCheckResourceAttr("ibm_appid_application_scopes.scopes", "scopes.1", "scope2"),
					resource.TestCheckResourceAttr("ibm_appid_application_scopes.scopes", "scopes.2", "scope3"),
				),
			},
		},
	})
}

func testAccCheckIBMApplicationScopesConfig(tenantID string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_application" "test_app" {
			tenant_id = "%s"
			name = "%s"
		}

		resource "ibm_appid_application_scopes" "scopes" {
		  tenant_id = ibm_appid_application.test_app.tenant_id
		  client_id = ibm_appid_application.test_app.client_id
		  scopes = ["scope1", "scope2", "scope3"]
		}
	`, tenantID, name)
}

func testAccCheckIBMAppIDApplicationScopesDestroy(s *terraform.State) error {
	appIDClient, err := testAccProvider.Meta().(ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_application_scopes" {
			continue
		}

		id := rs.Primary.ID
		idParts := strings.Split(id, "/")

		tenantID := idParts[0]
		clientID := idParts[1]

		_, _, err := appIDClient.GetApplicationScopes(&appid.GetApplicationScopesOptions{
			TenantID: &tenantID,
			ClientID: &clientID,
		})

		if err == nil {
			return fmt.Errorf("error checking if AppID application scopes resource (%s) has been destroyed", rs.Primary.ID)
		}
	}

	return nil
}
