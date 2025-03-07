package zendesk

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nukosuke/go-zendesk/zendesk"
	"strconv"
)

func dataSourceZendeskTicketField() *schema.Resource {
	return &schema.Resource{
		Read: func(data *schema.ResourceData, i interface{}) error {
			zd := i.(*zendesk.Client)
			return readTicketFieldDataSource(data, zd)
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"title": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"position": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"required": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"collapsed_for_agents": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"regexp_for_validation": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"title_in_portal": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"visible_in_portal": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"editable_in_portal": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"required_in_portal": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"tag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_field_options": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				Computed: true,
			},
			// Required only for "tagger" type
			// https://developer.zendesk.com/rest_api/docs/support/ticket_fields#updating-drop-down-field-options
			"custom_field_option": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
				Computed: true,
				//TODO: empty is invalid form
			},
			// "priority" and "status" fields only
			"sub_type_id": {
				Type:     schema.TypeInt,
				Computed: true,
				//TODO: validation
			},
			//TODO: this is not necessary because it's only for system field
			"removable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"agent_description": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readTicketFieldDataSource(d identifiableGetterSetter, zd zendesk.TicketFieldAPI) error {
	id := d.Get("id").(int)
	d.SetId(strconv.Itoa(id))
	return readTicketField(d, zd)
}
