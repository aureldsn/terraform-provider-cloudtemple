package provider

import (
	"context"

	"github.com/cloud-temple/terraform-provider-cloudtemple/internal/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHosts() *schema.Resource {
	return &schema.Resource{
		Description: "",

		ReadContext: readFullResource(func(ctx context.Context, client *client.Client, d *schema.ResourceData, sw *stateWriter) (interface{}, error) {
			hosts, err := client.Compute().Host().List(ctx, "", "", "", "")
			return map[string]interface{}{
				"id":    "hosts",
				"hosts": hosts,
			}, err
		}),

		Schema: map[string]*schema.Schema{
			// Out
			"hosts": {
				Type:     schema.TypeList,
				Computed: true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"moref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"machine_manager_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"metrics": {
							Type:     schema.TypeList,
							Computed: true,

							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"esx": {
										Type:     schema.TypeList,
										Computed: true,

										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"version": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"build": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"full_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"cpu": {
										Type:     schema.TypeList,
										Computed: true,

										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"overall_cpu_usage": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"cpu_mhz": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"cpu_cores": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"cpu_threads": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"memory": {
										Type:     schema.TypeList,
										Computed: true,

										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"memory_size": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"memory_usage": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"maintenance_status": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"uptime": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"connected": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"virtual_machines": {
							Type:     schema.TypeList,
							Computed: true,

							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
