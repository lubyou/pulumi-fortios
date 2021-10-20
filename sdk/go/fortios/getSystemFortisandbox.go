// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system fortisandbox
func LookupSystemFortisandbox(ctx *pulumi.Context, args *LookupSystemFortisandboxArgs, opts ...pulumi.InvokeOption) (*LookupSystemFortisandboxResult, error) {
	var rv LookupSystemFortisandboxResult
	err := ctx.Invoke("fortios:index/getSystemFortisandbox:GetSystemFortisandbox", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemFortisandbox.
type LookupSystemFortisandboxArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemFortisandbox.
type LookupSystemFortisandboxResult struct {
	// Notifier email address.
	Email string `pulumi:"email"`
	// Configure the level of SSL protection for secure communication with FortiSandbox.
	EncAlgorithm string `pulumi:"encAlgorithm"`
	// Enable/disable FortiSandbox Cloud.
	Forticloud string `pulumi:"forticloud"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specify outgoing interface to reach server.
	Interface string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server.
	InterfaceSelectMethod string `pulumi:"interfaceSelectMethod"`
	// IPv4 or IPv6 address of the remote FortiSandbox.
	Server string `pulumi:"server"`
	// Source IP address for communications to FortiSandbox.
	SourceIp string `pulumi:"sourceIp"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
	SslMinProtoVersion string `pulumi:"sslMinProtoVersion"`
	// Enable/disable FortiSandbox.
	Status    string  `pulumi:"status"`
	Vdomparam *string `pulumi:"vdomparam"`
}