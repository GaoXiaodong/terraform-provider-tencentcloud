package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccTencentCloudDataSourceDnsPodRecords(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_PREPAY) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsPodRecordsBasic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.tencentcloud_dnspod_records.record", "result.#"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_dnspod_records.record", "result.0.record_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_dnspod_records.record", "result.0.value"),
				),
			},
		},
	})
}

const testAccDnsPodRecordsBasic = `
data "tencentcloud_domains" "domains" {}

locals {
  domains = data.tencentcloud_domains.domains.list.*.domain_name
  target = [for i in local.domains: i if substr(i, 0, 7) == "tencent"]
}

data "tencentcloud_dnspod_records" "record" {
  domain = local.target[0]
  sort_field = "updated_on"
  sort_type  = "DESC"
}
`
