package p2pub

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceLoadBalancer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLoadBalancerRead,
		Timeouts: &schema.ResourceTimeout{
			Default: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			// D10M, D100M, D150M, D1000M
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"redundant": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			// Global, PrivateStandard, Private
			"external_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"external_trafficip_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"external_servicecode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_trafficip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_masterhost_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_slavehost_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			// PrivateStandard, Private
			"internal_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"internal_trafficip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internal_masterhost_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internal_slavehost_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internal_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internal_servicecode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trafficip_list": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv4_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv4_domainname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6_domainname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"host_list": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"master": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"external_ipv4_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"external_ipv6_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"internal_ipv4_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"static_route_list": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_route_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"destination": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"servicecode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				Optional: true,
			},
		},
	}
}

func dataSourceLoadBalancerRead(d *schema.ResourceData, m interface{}) error {
	if d.Get("service_code") == nil {
		return fmt.Errorf("service_code is required")
	}

	d.SetId(d.Get("service_code").(string))

	return resourceLoadBalancerRead(d, m)
}
