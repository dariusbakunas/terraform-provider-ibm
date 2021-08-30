package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppIDIDPSAMLMetadataDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDIDPSAMLMetadataDataSourceConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_idp_saml_metadata.meta", "tenant_id", appIDTenantID),
					resource.TestCheckResourceAttrSet("data.ibm_appid_idp_saml_metadata.meta", "metadata"),
				),
			},
		},
	})
}

func setupAppIDIDPSAMLMetadataDataSourceConfig(tenantID string) string {
	return fmt.Sprintf(`	
		data "ibm_appid_idp_saml_metadata" "meta" {
			tenant_id = "%s"
		}
	`, tenantID)
}
