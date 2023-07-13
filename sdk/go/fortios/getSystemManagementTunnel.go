// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemManagementTunnel(ctx *pulumi.Context, args *LookupSystemManagementTunnelArgs, opts ...pulumi.InvokeOption) (*LookupSystemManagementTunnelResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSystemManagementTunnelResult
	err := ctx.Invoke("fortios:index/getSystemManagementTunnel:GetSystemManagementTunnel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemManagementTunnel.
type LookupSystemManagementTunnelArgs struct {
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemManagementTunnel.
type LookupSystemManagementTunnelResult struct {
	AllowCollectStatistics string `pulumi:"allowCollectStatistics"`
	AllowConfigRestore     string `pulumi:"allowConfigRestore"`
	AllowPushConfiguration string `pulumi:"allowPushConfiguration"`
	AllowPushFirmware      string `pulumi:"allowPushFirmware"`
	AuthorizedManagerOnly  string `pulumi:"authorizedManagerOnly"`
	// The provider-assigned unique ID for this managed resource.
	Id           string  `pulumi:"id"`
	SerialNumber string  `pulumi:"serialNumber"`
	Status       string  `pulumi:"status"`
	Vdomparam    *string `pulumi:"vdomparam"`
}

func LookupSystemManagementTunnelOutput(ctx *pulumi.Context, args LookupSystemManagementTunnelOutputArgs, opts ...pulumi.InvokeOption) LookupSystemManagementTunnelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemManagementTunnelResult, error) {
			args := v.(LookupSystemManagementTunnelArgs)
			r, err := LookupSystemManagementTunnel(ctx, &args, opts...)
			var s LookupSystemManagementTunnelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemManagementTunnelResultOutput)
}

// A collection of arguments for invoking GetSystemManagementTunnel.
type LookupSystemManagementTunnelOutputArgs struct {
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemManagementTunnelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemManagementTunnelArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemManagementTunnel.
type LookupSystemManagementTunnelResultOutput struct{ *pulumi.OutputState }

func (LookupSystemManagementTunnelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemManagementTunnelResult)(nil)).Elem()
}

func (o LookupSystemManagementTunnelResultOutput) ToLookupSystemManagementTunnelResultOutput() LookupSystemManagementTunnelResultOutput {
	return o
}

func (o LookupSystemManagementTunnelResultOutput) ToLookupSystemManagementTunnelResultOutputWithContext(ctx context.Context) LookupSystemManagementTunnelResultOutput {
	return o
}

func (o LookupSystemManagementTunnelResultOutput) AllowCollectStatistics() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementTunnelResult) string { return v.AllowCollectStatistics }).(pulumi.StringOutput)
}

func (o LookupSystemManagementTunnelResultOutput) AllowConfigRestore() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementTunnelResult) string { return v.AllowConfigRestore }).(pulumi.StringOutput)
}

func (o LookupSystemManagementTunnelResultOutput) AllowPushConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementTunnelResult) string { return v.AllowPushConfiguration }).(pulumi.StringOutput)
}

func (o LookupSystemManagementTunnelResultOutput) AllowPushFirmware() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementTunnelResult) string { return v.AllowPushFirmware }).(pulumi.StringOutput)
}

func (o LookupSystemManagementTunnelResultOutput) AuthorizedManagerOnly() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementTunnelResult) string { return v.AuthorizedManagerOnly }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemManagementTunnelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementTunnelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemManagementTunnelResultOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementTunnelResult) string { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o LookupSystemManagementTunnelResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementTunnelResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSystemManagementTunnelResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemManagementTunnelResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemManagementTunnelResultOutput{})
}
