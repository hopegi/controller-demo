package recordsets

import "github.com/gophercloud/gophercloud"

func baseURL(c *gophercloud.ServiceClient, zoneID string) string {
	if zoneID!="" {
		return c.ServiceURL( "zones",zoneID, "recordsets")
	}else {
		return c.ServiceURL("recordsets")
	}
}

func rrsetURL(c *gophercloud.ServiceClient, zoneID string, rrsetID string) string {
	return c.ServiceURL("zones", zoneID, "recordsets", rrsetID)
}
