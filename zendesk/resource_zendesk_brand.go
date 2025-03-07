package zendesk

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	client "github.com/nukosuke/go-zendesk/zendesk"
)

// https://developer.zendesk.com/rest_api/docs/support/brands
func resourceZendeskBrand() *schema.Resource {
	return &schema.Resource{
		Create: func(d *schema.ResourceData, meta interface{}) error {
			return nil
		},
		Read: func(d *schema.ResourceData, meta interface{}) error {
			return nil
		},
		Update: func(d *schema.ResourceData, meta interface{}) error {
			return nil
		},
		Delete: func(d *schema.ResourceData, meta interface{}) error {
			return nil
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"brand_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"has_help_center": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"help_center_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"logo_attachment_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ticket_form_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Computed: true,
			},
			"subdomain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"host_mapping": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"signature_template": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func marshalBrand(brand client.Brand, d identifiableGetterSetter) error {
	fields := map[string]interface{}{
		"url":                brand.URL,
		"name":               brand.Name,
		"brand_url":          brand.BrandURL,
		"has_help_center":    brand.HasHelpCenter,
		"help_center_state":  brand.HelpCenterState,
		"active":             brand.Active,
		"default":            brand.Default,
		"logo_attachment_id": brand.Logo.ID,
		"ticket_form_ids":    brand.TicketFieldIDs,
		"subdomain":          brand.Subdomain,
		"host_mapping":       brand.HostMapping,
		"signature_template": brand.SignatureTemplate,
	}

	return setSchemaFields(d, fields)
}

func unmarshalBrand(d identifiableGetterSetter) (client.Brand, error) {
	brand := client.Brand{}

	if v := d.Id(); v != "" {
		id, err := atoi64(v)
		if err != nil {
			return brand, fmt.Errorf("could not parse brand id %s: %v", v, err)
		}
		brand.ID = id
	}

	if v, ok := d.GetOk("url"); ok {
		brand.URL = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		brand.Name = v.(string)
	}

	if v, ok := d.GetOk("brand_url"); ok {
		brand.BrandURL = v.(string)
	}

	if v, ok := d.GetOk("has_help_center"); ok {
		brand.HasHelpCenter = v.(bool)
	}

	if v, ok := d.GetOk("help_center_state"); ok {
		brand.HelpCenterState = v.(string)
	}

	if v, ok := d.GetOk("active"); ok {
		brand.Active = v.(bool)
	}

	if v, ok := d.GetOk("default"); ok {
		brand.Default = v.(bool)
	}

	if v, ok := d.GetOk("logo_attachment_id"); ok {
		brand.Logo.ID = v.(int64)
	}

	if v, ok := d.GetOk("ticket_form_ids"); ok {
		ticketFormIDs := v.(*schema.Set).List()
		for _, ticketFormID := range ticketFormIDs {
			brand.TicketFieldIDs = append(brand.TicketFieldIDs, int64(ticketFormID.(int)))
		}
	}

	if v, ok := d.GetOk("subdomain"); ok {
		brand.Subdomain = v.(string)
	}

	if v, ok := d.GetOk("host_mapping"); ok {
		brand.HostMapping = v.(string)
	}

	if v, ok := d.GetOk("signature_template"); ok {
		brand.SignatureTemplate = v.(string)
	}

	return brand, nil
}
