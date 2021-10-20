// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system wccp
func LookupSystemWccp(ctx *pulumi.Context, args *LookupSystemWccpArgs, opts ...pulumi.InvokeOption) (*LookupSystemWccpResult, error) {
	var rv LookupSystemWccpResult
	err := ctx.Invoke("fortios:index/getSystemWccp:GetSystemWccp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemWccp.
type LookupSystemWccpArgs struct {
	// Specify the serviceId of the desired system wccp.
	ServiceId string `pulumi:"serviceId"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemWccp.
type LookupSystemWccpResult struct {
	// Assignment bucket format for the WCCP cache engine.
	AssignmentBucketFormat string `pulumi:"assignmentBucketFormat"`
	// Assignment destination address mask.
	AssignmentDstaddrMask string `pulumi:"assignmentDstaddrMask"`
	// Hash key assignment preference.
	AssignmentMethod string `pulumi:"assignmentMethod"`
	// Assignment source address mask.
	AssignmentSrcaddrMask string `pulumi:"assignmentSrcaddrMask"`
	// Assignment of hash weight/ratio for the WCCP cache engine.
	AssignmentWeight int `pulumi:"assignmentWeight"`
	// Enable/disable MD5 authentication.
	Authentication string `pulumi:"authentication"`
	// Method used to forward traffic to the routers or to return to the cache engine.
	CacheEngineMethod string `pulumi:"cacheEngineMethod"`
	// IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
	CacheId string `pulumi:"cacheId"`
	// Method used to forward traffic to the cache servers.
	ForwardMethod string `pulumi:"forwardMethod"`
	// IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
	GroupAddress string `pulumi:"groupAddress"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Password for MD5 authentication.
	Password string `pulumi:"password"`
	// Service ports.
	Ports string `pulumi:"ports"`
	// Match method.
	PortsDefined string `pulumi:"portsDefined"`
	// Hash method.
	PrimaryHash string `pulumi:"primaryHash"`
	// Service priority.
	Priority int `pulumi:"priority"`
	// Service protocol.
	Protocol int `pulumi:"protocol"`
	// Method used to decline a redirected packet and return it to the FortiGate.
	ReturnMethod string `pulumi:"returnMethod"`
	// IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
	RouterId string `pulumi:"routerId"`
	// IP addresses of one or more WCCP routers.
	RouterList string `pulumi:"routerList"`
	// IP addresses and netmasks for up to four cache servers.
	ServerList string `pulumi:"serverList"`
	// Cache server type.
	ServerType string `pulumi:"serverType"`
	// Service ID.
	ServiceId string `pulumi:"serviceId"`
	// WCCP service type used by the cache server for logical interception and redirection of traffic.
	ServiceType string  `pulumi:"serviceType"`
	Vdomparam   *string `pulumi:"vdomparam"`
}