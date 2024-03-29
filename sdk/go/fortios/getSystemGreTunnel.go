// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemGreTunnel(ctx *pulumi.Context, args *LookupSystemGreTunnelArgs, opts ...pulumi.InvokeOption) (*LookupSystemGreTunnelResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSystemGreTunnelResult
	err := ctx.Invoke("fortios:index/getSystemGreTunnel:GetSystemGreTunnel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemGreTunnel.
type LookupSystemGreTunnelArgs struct {
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemGreTunnel.
type LookupSystemGreTunnelResult struct {
	ChecksumReception    string `pulumi:"checksumReception"`
	ChecksumTransmission string `pulumi:"checksumTransmission"`
	Diffservcode         string `pulumi:"diffservcode"`
	DscpCopying          string `pulumi:"dscpCopying"`
	// The provider-assigned unique ID for this managed resource.
	Id                         string  `pulumi:"id"`
	Interface                  string  `pulumi:"interface"`
	IpVersion                  string  `pulumi:"ipVersion"`
	KeepaliveFailtimes         int     `pulumi:"keepaliveFailtimes"`
	KeepaliveInterval          int     `pulumi:"keepaliveInterval"`
	KeyInbound                 int     `pulumi:"keyInbound"`
	KeyOutbound                int     `pulumi:"keyOutbound"`
	LocalGw                    string  `pulumi:"localGw"`
	LocalGw6                   string  `pulumi:"localGw6"`
	Name                       string  `pulumi:"name"`
	RemoteGw                   string  `pulumi:"remoteGw"`
	RemoteGw6                  string  `pulumi:"remoteGw6"`
	SequenceNumberReception    string  `pulumi:"sequenceNumberReception"`
	SequenceNumberTransmission string  `pulumi:"sequenceNumberTransmission"`
	UseSdwan                   string  `pulumi:"useSdwan"`
	Vdomparam                  *string `pulumi:"vdomparam"`
}

func LookupSystemGreTunnelOutput(ctx *pulumi.Context, args LookupSystemGreTunnelOutputArgs, opts ...pulumi.InvokeOption) LookupSystemGreTunnelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemGreTunnelResult, error) {
			args := v.(LookupSystemGreTunnelArgs)
			r, err := LookupSystemGreTunnel(ctx, &args, opts...)
			var s LookupSystemGreTunnelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemGreTunnelResultOutput)
}

// A collection of arguments for invoking GetSystemGreTunnel.
type LookupSystemGreTunnelOutputArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemGreTunnelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemGreTunnelArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemGreTunnel.
type LookupSystemGreTunnelResultOutput struct{ *pulumi.OutputState }

func (LookupSystemGreTunnelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemGreTunnelResult)(nil)).Elem()
}

func (o LookupSystemGreTunnelResultOutput) ToLookupSystemGreTunnelResultOutput() LookupSystemGreTunnelResultOutput {
	return o
}

func (o LookupSystemGreTunnelResultOutput) ToLookupSystemGreTunnelResultOutputWithContext(ctx context.Context) LookupSystemGreTunnelResultOutput {
	return o
}

func (o LookupSystemGreTunnelResultOutput) ChecksumReception() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) string { return v.ChecksumReception }).(pulumi.StringOutput)
}

func (o LookupSystemGreTunnelResultOutput) ChecksumTransmission() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) string { return v.ChecksumTransmission }).(pulumi.StringOutput)
}

func (o LookupSystemGreTunnelResultOutput) Diffservcode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) string { return v.Diffservcode }).(pulumi.StringOutput)
}

func (o LookupSystemGreTunnelResultOutput) DscpCopying() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) string { return v.DscpCopying }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemGreTunnelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemGreTunnelResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) string { return v.Interface }).(pulumi.StringOutput)
}

func (o LookupSystemGreTunnelResultOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) string { return v.IpVersion }).(pulumi.StringOutput)
}

func (o LookupSystemGreTunnelResultOutput) KeepaliveFailtimes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) int { return v.KeepaliveFailtimes }).(pulumi.IntOutput)
}

func (o LookupSystemGreTunnelResultOutput) KeepaliveInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) int { return v.KeepaliveInterval }).(pulumi.IntOutput)
}

func (o LookupSystemGreTunnelResultOutput) KeyInbound() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) int { return v.KeyInbound }).(pulumi.IntOutput)
}

func (o LookupSystemGreTunnelResultOutput) KeyOutbound() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) int { return v.KeyOutbound }).(pulumi.IntOutput)
}

func (o LookupSystemGreTunnelResultOutput) LocalGw() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) string { return v.LocalGw }).(pulumi.StringOutput)
}

func (o LookupSystemGreTunnelResultOutput) LocalGw6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) string { return v.LocalGw6 }).(pulumi.StringOutput)
}

func (o LookupSystemGreTunnelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSystemGreTunnelResultOutput) RemoteGw() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) string { return v.RemoteGw }).(pulumi.StringOutput)
}

func (o LookupSystemGreTunnelResultOutput) RemoteGw6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) string { return v.RemoteGw6 }).(pulumi.StringOutput)
}

func (o LookupSystemGreTunnelResultOutput) SequenceNumberReception() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) string { return v.SequenceNumberReception }).(pulumi.StringOutput)
}

func (o LookupSystemGreTunnelResultOutput) SequenceNumberTransmission() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) string { return v.SequenceNumberTransmission }).(pulumi.StringOutput)
}

func (o LookupSystemGreTunnelResultOutput) UseSdwan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) string { return v.UseSdwan }).(pulumi.StringOutput)
}

func (o LookupSystemGreTunnelResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemGreTunnelResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemGreTunnelResultOutput{})
}
