package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccAppIDMFAChannelDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupIBMMFAChannelDataSourceConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_mfa_channel.mf", "tenant_id", appIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_mfa_channel.mf", "active", "sms"),
					resource.TestCheckResourceAttr("data.ibm_appid_mfa_channel.mf", "sms_config.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_appid_mfa_channel.mf", "sms_config.0.key", "api_key"),
					resource.TestCheckResourceAttr("data.ibm_appid_mfa_channel.mf", "sms_config.0.secret", "api_secret"),
					resource.TestCheckResourceAttr("data.ibm_appid_mfa_channel.mf", "sms_config.0.from", "+12223334444"),
				),
			},
		},
	})
}

func setupIBMMFAChannelDataSourceConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_mfa_channel" "mf" {
			tenant_id = "%s"
			active = "sms"

			sms_config {
			  key = "api_key"
			  secret = "api_secret"
			  from = "+12223334444"
		    }
		}
		
		data "ibm_appid_mfa_channel" "mf" {
			tenant_id = ibm_appid_mfa_channel.mf.tenant_id
			depends_on = [
				ibm_appid_mfa_channel.mf
			]
		}
	`, tenantID)
}
