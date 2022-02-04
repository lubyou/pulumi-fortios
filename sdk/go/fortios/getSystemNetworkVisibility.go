// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system networkvisibility
func LookupSystemNetworkVisibility(ctx *pulumi.Context, args *LookupSystemNetworkVisibilityArgs, opts ...pulumi.InvokeOption) (*LookupSystemNetworkVisibilityResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemNetworkVisibilityResult
	err := ctx.Invoke("fortios:index/getSystemNetworkVisibility:GetSystemNetworkVisibility", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemNetworkVisibility.
type LookupSystemNetworkVisibilityArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemNetworkVisibility.
type LookupSystemNetworkVisibilityResult struct {
	// Enable/disable logging of destination hostname visibility.
	DestinationHostnameVisibility string `pulumi:"destinationHostnameVisibility"`
	// Enable/disable logging of destination geographical location visibility.
	DestinationLocation string `pulumi:"destinationLocation"`
	// Enable/disable logging of destination visibility.
	DestinationVisibility string `pulumi:"destinationVisibility"`
	// Limit of the number of hostname table entries (0 - 50000).
	HostnameLimit int `pulumi:"hostnameLimit"`
	// TTL of hostname table entries (60 - 86400).
	HostnameTtl int `pulumi:"hostnameTtl"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable logging of source geographical location visibility.
	SourceLocation string  `pulumi:"sourceLocation"`
	Vdomparam      *string `pulumi:"vdomparam"`
}

func LookupSystemNetworkVisibilityOutput(ctx *pulumi.Context, args LookupSystemNetworkVisibilityOutputArgs, opts ...pulumi.InvokeOption) LookupSystemNetworkVisibilityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemNetworkVisibilityResult, error) {
			args := v.(LookupSystemNetworkVisibilityArgs)
			r, err := LookupSystemNetworkVisibility(ctx, &args, opts...)
			return *r, err
		}).(LookupSystemNetworkVisibilityResultOutput)
}

// A collection of arguments for invoking GetSystemNetworkVisibility.
type LookupSystemNetworkVisibilityOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemNetworkVisibilityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemNetworkVisibilityArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemNetworkVisibility.
type LookupSystemNetworkVisibilityResultOutput struct{ *pulumi.OutputState }

func (LookupSystemNetworkVisibilityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemNetworkVisibilityResult)(nil)).Elem()
}

func (o LookupSystemNetworkVisibilityResultOutput) ToLookupSystemNetworkVisibilityResultOutput() LookupSystemNetworkVisibilityResultOutput {
	return o
}

func (o LookupSystemNetworkVisibilityResultOutput) ToLookupSystemNetworkVisibilityResultOutputWithContext(ctx context.Context) LookupSystemNetworkVisibilityResultOutput {
	return o
}

// Enable/disable logging of destination hostname visibility.
func (o LookupSystemNetworkVisibilityResultOutput) DestinationHostnameVisibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNetworkVisibilityResult) string { return v.DestinationHostnameVisibility }).(pulumi.StringOutput)
}

// Enable/disable logging of destination geographical location visibility.
func (o LookupSystemNetworkVisibilityResultOutput) DestinationLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNetworkVisibilityResult) string { return v.DestinationLocation }).(pulumi.StringOutput)
}

// Enable/disable logging of destination visibility.
func (o LookupSystemNetworkVisibilityResultOutput) DestinationVisibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNetworkVisibilityResult) string { return v.DestinationVisibility }).(pulumi.StringOutput)
}

// Limit of the number of hostname table entries (0 - 50000).
func (o LookupSystemNetworkVisibilityResultOutput) HostnameLimit() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemNetworkVisibilityResult) int { return v.HostnameLimit }).(pulumi.IntOutput)
}

// TTL of hostname table entries (60 - 86400).
func (o LookupSystemNetworkVisibilityResultOutput) HostnameTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemNetworkVisibilityResult) int { return v.HostnameTtl }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemNetworkVisibilityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNetworkVisibilityResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable logging of source geographical location visibility.
func (o LookupSystemNetworkVisibilityResultOutput) SourceLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNetworkVisibilityResult) string { return v.SourceLocation }).(pulumi.StringOutput)
}

func (o LookupSystemNetworkVisibilityResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemNetworkVisibilityResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemNetworkVisibilityResultOutput{})
}
