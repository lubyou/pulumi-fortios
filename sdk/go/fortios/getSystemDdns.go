// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemDdns(ctx *pulumi.Context, args *LookupSystemDdnsArgs, opts ...pulumi.InvokeOption) (*LookupSystemDdnsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSystemDdnsResult
	err := ctx.Invoke("fortios:index/getSystemDdns:GetSystemDdns", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemDdns.
type LookupSystemDdnsArgs struct {
	Ddnsid    int     `pulumi:"ddnsid"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemDdns.
type LookupSystemDdnsResult struct {
	AddrType        string                        `pulumi:"addrType"`
	BoundIp         string                        `pulumi:"boundIp"`
	ClearText       string                        `pulumi:"clearText"`
	DdnsAuth        string                        `pulumi:"ddnsAuth"`
	DdnsDomain      string                        `pulumi:"ddnsDomain"`
	DdnsKey         string                        `pulumi:"ddnsKey"`
	DdnsKeyname     string                        `pulumi:"ddnsKeyname"`
	DdnsPassword    string                        `pulumi:"ddnsPassword"`
	DdnsServer      string                        `pulumi:"ddnsServer"`
	DdnsServerAddrs []GetSystemDdnsDdnsServerAddr `pulumi:"ddnsServerAddrs"`
	DdnsServerIp    string                        `pulumi:"ddnsServerIp"`
	DdnsSn          string                        `pulumi:"ddnsSn"`
	DdnsTtl         int                           `pulumi:"ddnsTtl"`
	DdnsUsername    string                        `pulumi:"ddnsUsername"`
	DdnsZone        string                        `pulumi:"ddnsZone"`
	Ddnsid          int                           `pulumi:"ddnsid"`
	// The provider-assigned unique ID for this managed resource.
	Id                string                          `pulumi:"id"`
	MonitorInterfaces []GetSystemDdnsMonitorInterface `pulumi:"monitorInterfaces"`
	ServerType        string                          `pulumi:"serverType"`
	SslCertificate    string                          `pulumi:"sslCertificate"`
	UpdateInterval    int                             `pulumi:"updateInterval"`
	UsePublicIp       string                          `pulumi:"usePublicIp"`
	Vdomparam         *string                         `pulumi:"vdomparam"`
}

func LookupSystemDdnsOutput(ctx *pulumi.Context, args LookupSystemDdnsOutputArgs, opts ...pulumi.InvokeOption) LookupSystemDdnsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemDdnsResult, error) {
			args := v.(LookupSystemDdnsArgs)
			r, err := LookupSystemDdns(ctx, &args, opts...)
			var s LookupSystemDdnsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemDdnsResultOutput)
}

// A collection of arguments for invoking GetSystemDdns.
type LookupSystemDdnsOutputArgs struct {
	Ddnsid    pulumi.IntInput       `pulumi:"ddnsid"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemDdnsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemDdnsArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemDdns.
type LookupSystemDdnsResultOutput struct{ *pulumi.OutputState }

func (LookupSystemDdnsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemDdnsResult)(nil)).Elem()
}

func (o LookupSystemDdnsResultOutput) ToLookupSystemDdnsResultOutput() LookupSystemDdnsResultOutput {
	return o
}

func (o LookupSystemDdnsResultOutput) ToLookupSystemDdnsResultOutputWithContext(ctx context.Context) LookupSystemDdnsResultOutput {
	return o
}

func (o LookupSystemDdnsResultOutput) AddrType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.AddrType }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) BoundIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.BoundIp }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) ClearText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.ClearText }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) DdnsAuth() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.DdnsAuth }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) DdnsDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.DdnsDomain }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) DdnsKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.DdnsKey }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) DdnsKeyname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.DdnsKeyname }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) DdnsPassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.DdnsPassword }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) DdnsServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.DdnsServer }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) DdnsServerAddrs() GetSystemDdnsDdnsServerAddrArrayOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) []GetSystemDdnsDdnsServerAddr { return v.DdnsServerAddrs }).(GetSystemDdnsDdnsServerAddrArrayOutput)
}

func (o LookupSystemDdnsResultOutput) DdnsServerIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.DdnsServerIp }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) DdnsSn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.DdnsSn }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) DdnsTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) int { return v.DdnsTtl }).(pulumi.IntOutput)
}

func (o LookupSystemDdnsResultOutput) DdnsUsername() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.DdnsUsername }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) DdnsZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.DdnsZone }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) Ddnsid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) int { return v.Ddnsid }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemDdnsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) MonitorInterfaces() GetSystemDdnsMonitorInterfaceArrayOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) []GetSystemDdnsMonitorInterface { return v.MonitorInterfaces }).(GetSystemDdnsMonitorInterfaceArrayOutput)
}

func (o LookupSystemDdnsResultOutput) ServerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.ServerType }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) SslCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.SslCertificate }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) UpdateInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) int { return v.UpdateInterval }).(pulumi.IntOutput)
}

func (o LookupSystemDdnsResultOutput) UsePublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) string { return v.UsePublicIp }).(pulumi.StringOutput)
}

func (o LookupSystemDdnsResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemDdnsResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemDdnsResultOutput{})
}
