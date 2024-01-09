// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemDnsDatabase(ctx *pulumi.Context, args *LookupSystemDnsDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupSystemDnsDatabaseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSystemDnsDatabaseResult
	err := ctx.Invoke("fortios:index/getSystemDnsDatabase:GetSystemDnsDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemDnsDatabase.
type LookupSystemDnsDatabaseArgs struct {
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemDnsDatabase.
type LookupSystemDnsDatabaseResult struct {
	AllowTransfer string                         `pulumi:"allowTransfer"`
	Authoritative string                         `pulumi:"authoritative"`
	Contact       string                         `pulumi:"contact"`
	DnsEntries    []GetSystemDnsDatabaseDnsEntry `pulumi:"dnsEntries"`
	Domain        string                         `pulumi:"domain"`
	Forwarder     string                         `pulumi:"forwarder"`
	Forwarder6    string                         `pulumi:"forwarder6"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	IpMaster    string  `pulumi:"ipMaster"`
	IpPrimary   string  `pulumi:"ipPrimary"`
	Name        string  `pulumi:"name"`
	PrimaryName string  `pulumi:"primaryName"`
	RrMax       int     `pulumi:"rrMax"`
	SourceIp    string  `pulumi:"sourceIp"`
	SourceIp6   string  `pulumi:"sourceIp6"`
	Status      string  `pulumi:"status"`
	Ttl         int     `pulumi:"ttl"`
	Type        string  `pulumi:"type"`
	Vdomparam   *string `pulumi:"vdomparam"`
	View        string  `pulumi:"view"`
}

func LookupSystemDnsDatabaseOutput(ctx *pulumi.Context, args LookupSystemDnsDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupSystemDnsDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemDnsDatabaseResult, error) {
			args := v.(LookupSystemDnsDatabaseArgs)
			r, err := LookupSystemDnsDatabase(ctx, &args, opts...)
			var s LookupSystemDnsDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemDnsDatabaseResultOutput)
}

// A collection of arguments for invoking GetSystemDnsDatabase.
type LookupSystemDnsDatabaseOutputArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemDnsDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemDnsDatabaseArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemDnsDatabase.
type LookupSystemDnsDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupSystemDnsDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemDnsDatabaseResult)(nil)).Elem()
}

func (o LookupSystemDnsDatabaseResultOutput) ToLookupSystemDnsDatabaseResultOutput() LookupSystemDnsDatabaseResultOutput {
	return o
}

func (o LookupSystemDnsDatabaseResultOutput) ToLookupSystemDnsDatabaseResultOutputWithContext(ctx context.Context) LookupSystemDnsDatabaseResultOutput {
	return o
}

func (o LookupSystemDnsDatabaseResultOutput) AllowTransfer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.AllowTransfer }).(pulumi.StringOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) Authoritative() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.Authoritative }).(pulumi.StringOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) Contact() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.Contact }).(pulumi.StringOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) DnsEntries() GetSystemDnsDatabaseDnsEntryArrayOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) []GetSystemDnsDatabaseDnsEntry { return v.DnsEntries }).(GetSystemDnsDatabaseDnsEntryArrayOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.Domain }).(pulumi.StringOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) Forwarder() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.Forwarder }).(pulumi.StringOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) Forwarder6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.Forwarder6 }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemDnsDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) IpMaster() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.IpMaster }).(pulumi.StringOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) IpPrimary() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.IpPrimary }).(pulumi.StringOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) PrimaryName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.PrimaryName }).(pulumi.StringOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) RrMax() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) int { return v.RrMax }).(pulumi.IntOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) SourceIp6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.SourceIp6 }).(pulumi.StringOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) int { return v.Ttl }).(pulumi.IntOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LookupSystemDnsDatabaseResultOutput) View() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsDatabaseResult) string { return v.View }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemDnsDatabaseResultOutput{})
}
