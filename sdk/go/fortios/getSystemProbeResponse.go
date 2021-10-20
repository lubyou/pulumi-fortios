// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system proberesponse
func LookupSystemProbeResponse(ctx *pulumi.Context, args *LookupSystemProbeResponseArgs, opts ...pulumi.InvokeOption) (*LookupSystemProbeResponseResult, error) {
	var rv LookupSystemProbeResponseResult
	err := ctx.Invoke("fortios:index/getSystemProbeResponse:GetSystemProbeResponse", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemProbeResponse.
type LookupSystemProbeResponseArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemProbeResponse.
type LookupSystemProbeResponseResult struct {
	// Value to respond to the monitoring server.
	HttpProbeValue string `pulumi:"httpProbeValue"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// SLA response mode.
	Mode string `pulumi:"mode"`
	// Twamp respondor password in authentication mode
	Password string `pulumi:"password"`
	// Port number to response.
	Port int `pulumi:"port"`
	// Twamp respondor security mode.
	SecurityMode string `pulumi:"securityMode"`
	// An inactivity timer for a twamp test session.
	Timeout int `pulumi:"timeout"`
	// Mode for TWAMP packet TTL modification.
	TtlMode   string  `pulumi:"ttlMode"`
	Vdomparam *string `pulumi:"vdomparam"`
}