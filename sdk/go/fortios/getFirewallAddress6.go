// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall address6
func LookupFirewallAddress6(ctx *pulumi.Context, args *LookupFirewallAddress6Args, opts ...pulumi.InvokeOption) (*LookupFirewallAddress6Result, error) {
	var rv LookupFirewallAddress6Result
	err := ctx.Invoke("fortios:index/getFirewallAddress6:GetFirewallAddress6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallAddress6.
type LookupFirewallAddress6Args struct {
	// Specify the name of the desired firewall address6.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallAddress6.
type LookupFirewallAddress6Result struct {
	// Minimal TTL of individual IPv6 addresses in FQDN cache.
	CacheTtl int `pulumi:"cacheTtl"`
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color int `pulumi:"color"`
	// Comment.
	Comment string `pulumi:"comment"`
	// IPv6 addresses associated to a specific country.
	Country string `pulumi:"country"`
	// Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
	EndIp string `pulumi:"endIp"`
	// Last MAC address in the range.
	EndMac string `pulumi:"endMac"`
	// Fully qualified domain name.
	Fqdn string `pulumi:"fqdn"`
	// Host Address.
	Host string `pulumi:"host"`
	// Host type.
	HostType string `pulumi:"hostType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
	Ip6 string `pulumi:"ip6"`
	// IP address list. The structure of `list` block is documented below.
	Lists []GetFirewallAddress6ListType `pulumi:"lists"`
	// Name.
	Name string `pulumi:"name"`
	// Object ID for NSX.
	ObjId string `pulumi:"objId"`
	// SDN.
	Sdn string `pulumi:"sdn"`
	// First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
	StartIp string `pulumi:"startIp"`
	// First MAC address in the range.
	StartMac string `pulumi:"startMac"`
	// IPv6 subnet segments. The structure of `subnetSegment` block is documented below.
	SubnetSegments []GetFirewallAddress6SubnetSegment `pulumi:"subnetSegments"`
	// Config object tagging The structure of `tagging` block is documented below.
	Taggings []GetFirewallAddress6Tagging `pulumi:"taggings"`
	// IPv6 address template.
	Template string `pulumi:"template"`
	// Subnet segment type.
	Type string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid      string  `pulumi:"uuid"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable the visibility of the object in the GUI.
	Visibility string `pulumi:"visibility"`
}