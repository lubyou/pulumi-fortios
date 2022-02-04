// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system fortisandbox
func LookupSystemFortisandbox(ctx *pulumi.Context, args *LookupSystemFortisandboxArgs, opts ...pulumi.InvokeOption) (*LookupSystemFortisandboxResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
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

func LookupSystemFortisandboxOutput(ctx *pulumi.Context, args LookupSystemFortisandboxOutputArgs, opts ...pulumi.InvokeOption) LookupSystemFortisandboxResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemFortisandboxResult, error) {
			args := v.(LookupSystemFortisandboxArgs)
			r, err := LookupSystemFortisandbox(ctx, &args, opts...)
			return *r, err
		}).(LookupSystemFortisandboxResultOutput)
}

// A collection of arguments for invoking GetSystemFortisandbox.
type LookupSystemFortisandboxOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemFortisandboxOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemFortisandboxArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemFortisandbox.
type LookupSystemFortisandboxResultOutput struct{ *pulumi.OutputState }

func (LookupSystemFortisandboxResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemFortisandboxResult)(nil)).Elem()
}

func (o LookupSystemFortisandboxResultOutput) ToLookupSystemFortisandboxResultOutput() LookupSystemFortisandboxResultOutput {
	return o
}

func (o LookupSystemFortisandboxResultOutput) ToLookupSystemFortisandboxResultOutputWithContext(ctx context.Context) LookupSystemFortisandboxResultOutput {
	return o
}

// Notifier email address.
func (o LookupSystemFortisandboxResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.Email }).(pulumi.StringOutput)
}

// Configure the level of SSL protection for secure communication with FortiSandbox.
func (o LookupSystemFortisandboxResultOutput) EncAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.EncAlgorithm }).(pulumi.StringOutput)
}

// Enable/disable FortiSandbox Cloud.
func (o LookupSystemFortisandboxResultOutput) Forticloud() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.Forticloud }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemFortisandboxResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specify outgoing interface to reach server.
func (o LookupSystemFortisandboxResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server.
func (o LookupSystemFortisandboxResultOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// IPv4 or IPv6 address of the remote FortiSandbox.
func (o LookupSystemFortisandboxResultOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.Server }).(pulumi.StringOutput)
}

// Source IP address for communications to FortiSandbox.
func (o LookupSystemFortisandboxResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
func (o LookupSystemFortisandboxResultOutput) SslMinProtoVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.SslMinProtoVersion }).(pulumi.StringOutput)
}

// Enable/disable FortiSandbox.
func (o LookupSystemFortisandboxResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSystemFortisandboxResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemFortisandboxResultOutput{})
}
