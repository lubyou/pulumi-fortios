// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemDnsServer(ctx *pulumi.Context, args *LookupSystemDnsServerArgs, opts ...pulumi.InvokeOption) (*LookupSystemDnsServerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSystemDnsServerResult
	err := ctx.Invoke("fortios:index/getSystemDnsServer:GetSystemDnsServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemDnsServer.
type LookupSystemDnsServerArgs struct {
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemDnsServer.
type LookupSystemDnsServerResult struct {
	DnsfilterProfile string `pulumi:"dnsfilterProfile"`
	Doh              string `pulumi:"doh"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Mode      string  `pulumi:"mode"`
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemDnsServerOutput(ctx *pulumi.Context, args LookupSystemDnsServerOutputArgs, opts ...pulumi.InvokeOption) LookupSystemDnsServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemDnsServerResult, error) {
			args := v.(LookupSystemDnsServerArgs)
			r, err := LookupSystemDnsServer(ctx, &args, opts...)
			var s LookupSystemDnsServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemDnsServerResultOutput)
}

// A collection of arguments for invoking GetSystemDnsServer.
type LookupSystemDnsServerOutputArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemDnsServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemDnsServerArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemDnsServer.
type LookupSystemDnsServerResultOutput struct{ *pulumi.OutputState }

func (LookupSystemDnsServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemDnsServerResult)(nil)).Elem()
}

func (o LookupSystemDnsServerResultOutput) ToLookupSystemDnsServerResultOutput() LookupSystemDnsServerResultOutput {
	return o
}

func (o LookupSystemDnsServerResultOutput) ToLookupSystemDnsServerResultOutputWithContext(ctx context.Context) LookupSystemDnsServerResultOutput {
	return o
}

func (o LookupSystemDnsServerResultOutput) DnsfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsServerResult) string { return v.DnsfilterProfile }).(pulumi.StringOutput)
}

func (o LookupSystemDnsServerResultOutput) Doh() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsServerResult) string { return v.Doh }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemDnsServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemDnsServerResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsServerResult) string { return v.Mode }).(pulumi.StringOutput)
}

func (o LookupSystemDnsServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsServerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSystemDnsServerResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemDnsServerResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemDnsServerResultOutput{})
}
