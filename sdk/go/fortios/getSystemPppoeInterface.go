// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system pppoeinterface
func LookupSystemPppoeInterface(ctx *pulumi.Context, args *LookupSystemPppoeInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupSystemPppoeInterfaceResult, error) {
	var rv LookupSystemPppoeInterfaceResult
	err := ctx.Invoke("fortios:index/getSystemPppoeInterface:GetSystemPppoeInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemPppoeInterface.
type LookupSystemPppoeInterfaceArgs struct {
	// Specify the name of the desired system pppoeinterface.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemPppoeInterface.
type LookupSystemPppoeInterfaceResult struct {
	// PPPoE AC name.
	AcName string `pulumi:"acName"`
	// PPP authentication type to use.
	AuthType string `pulumi:"authType"`
	// Name for the physical interface.
	Device string `pulumi:"device"`
	// Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface.
	DialOnDemand string `pulumi:"dialOnDemand"`
	// PPPoE discovery init timeout value in (0-4294967295 sec).
	DiscRetryTimeout int `pulumi:"discRetryTimeout"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// PPPoE auto disconnect after idle timeout (0-4294967295 sec).
	IdleTimeout int `pulumi:"idleTimeout"`
	// PPPoE unnumbered IP.
	Ipunnumbered string `pulumi:"ipunnumbered"`
	// Enable/disable IPv6 Control Protocol (IPv6CP).
	Ipv6 string `pulumi:"ipv6"`
	// PPPoE LCP echo interval in (0-4294967295 sec, default = 5).
	LcpEchoInterval int `pulumi:"lcpEchoInterval"`
	// Maximum missed LCP echo messages before disconnect (0-4294967295, default = 3).
	LcpMaxEchoFails int `pulumi:"lcpMaxEchoFails"`
	// Name of the PPPoE interface.
	Name string `pulumi:"name"`
	// PPPoE terminate timeout value in (0-4294967295 sec).
	PadtRetryTimeout int `pulumi:"padtRetryTimeout"`
	// Enter the password.
	Password string `pulumi:"password"`
	// Enable/disable PPPoE unnumbered negotiation.
	PppoeUnnumberedNegotiate string `pulumi:"pppoeUnnumberedNegotiate"`
	// PPPoE service name.
	ServiceName string `pulumi:"serviceName"`
	// User name.
	Username  string  `pulumi:"username"`
	Vdomparam *string `pulumi:"vdomparam"`
}
