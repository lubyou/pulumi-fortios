// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemVxlan(ctx *pulumi.Context, args *LookupSystemVxlanArgs, opts ...pulumi.InvokeOption) (*LookupSystemVxlanResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSystemVxlanResult
	err := ctx.Invoke("fortios:index/getSystemVxlan:GetSystemVxlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemVxlan.
type LookupSystemVxlanArgs struct {
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemVxlan.
type LookupSystemVxlanResult struct {
	Dstport int `pulumi:"dstport"`
	EvpnId  int `pulumi:"evpnId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                    `pulumi:"id"`
	Interface        string                    `pulumi:"interface"`
	IpVersion        string                    `pulumi:"ipVersion"`
	LearnFromTraffic string                    `pulumi:"learnFromTraffic"`
	MulticastTtl     int                       `pulumi:"multicastTtl"`
	Name             string                    `pulumi:"name"`
	RemoteIp6s       []GetSystemVxlanRemoteIp6 `pulumi:"remoteIp6s"`
	RemoteIps        []GetSystemVxlanRemoteIp  `pulumi:"remoteIps"`
	Vdomparam        *string                   `pulumi:"vdomparam"`
	Vni              int                       `pulumi:"vni"`
}

func LookupSystemVxlanOutput(ctx *pulumi.Context, args LookupSystemVxlanOutputArgs, opts ...pulumi.InvokeOption) LookupSystemVxlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemVxlanResult, error) {
			args := v.(LookupSystemVxlanArgs)
			r, err := LookupSystemVxlan(ctx, &args, opts...)
			var s LookupSystemVxlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemVxlanResultOutput)
}

// A collection of arguments for invoking GetSystemVxlan.
type LookupSystemVxlanOutputArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemVxlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemVxlanArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemVxlan.
type LookupSystemVxlanResultOutput struct{ *pulumi.OutputState }

func (LookupSystemVxlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemVxlanResult)(nil)).Elem()
}

func (o LookupSystemVxlanResultOutput) ToLookupSystemVxlanResultOutput() LookupSystemVxlanResultOutput {
	return o
}

func (o LookupSystemVxlanResultOutput) ToLookupSystemVxlanResultOutputWithContext(ctx context.Context) LookupSystemVxlanResultOutput {
	return o
}

func (o LookupSystemVxlanResultOutput) Dstport() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemVxlanResult) int { return v.Dstport }).(pulumi.IntOutput)
}

func (o LookupSystemVxlanResultOutput) EvpnId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemVxlanResult) int { return v.EvpnId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemVxlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVxlanResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemVxlanResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVxlanResult) string { return v.Interface }).(pulumi.StringOutput)
}

func (o LookupSystemVxlanResultOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVxlanResult) string { return v.IpVersion }).(pulumi.StringOutput)
}

func (o LookupSystemVxlanResultOutput) LearnFromTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVxlanResult) string { return v.LearnFromTraffic }).(pulumi.StringOutput)
}

func (o LookupSystemVxlanResultOutput) MulticastTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemVxlanResult) int { return v.MulticastTtl }).(pulumi.IntOutput)
}

func (o LookupSystemVxlanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVxlanResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSystemVxlanResultOutput) RemoteIp6s() GetSystemVxlanRemoteIp6ArrayOutput {
	return o.ApplyT(func(v LookupSystemVxlanResult) []GetSystemVxlanRemoteIp6 { return v.RemoteIp6s }).(GetSystemVxlanRemoteIp6ArrayOutput)
}

func (o LookupSystemVxlanResultOutput) RemoteIps() GetSystemVxlanRemoteIpArrayOutput {
	return o.ApplyT(func(v LookupSystemVxlanResult) []GetSystemVxlanRemoteIp { return v.RemoteIps }).(GetSystemVxlanRemoteIpArrayOutput)
}

func (o LookupSystemVxlanResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemVxlanResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LookupSystemVxlanResultOutput) Vni() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemVxlanResult) int { return v.Vni }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemVxlanResultOutput{})
}
