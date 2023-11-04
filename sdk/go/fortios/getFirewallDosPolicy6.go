// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func LookupFirewallDosPolicy6(ctx *pulumi.Context, args *LookupFirewallDosPolicy6Args, opts ...pulumi.InvokeOption) (*LookupFirewallDosPolicy6Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFirewallDosPolicy6Result
	err := ctx.Invoke("fortios:index/getFirewallDosPolicy6:GetFirewallDosPolicy6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallDosPolicy6.
type LookupFirewallDosPolicy6Args struct {
	Policyid  int     `pulumi:"policyid"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallDosPolicy6.
type LookupFirewallDosPolicy6Result struct {
	Anomalies []GetFirewallDosPolicy6Anomaly `pulumi:"anomalies"`
	Comments  string                         `pulumi:"comments"`
	Dstaddrs  []GetFirewallDosPolicy6Dstaddr `pulumi:"dstaddrs"`
	// The provider-assigned unique ID for this managed resource.
	Id        string                         `pulumi:"id"`
	Interface string                         `pulumi:"interface"`
	Name      string                         `pulumi:"name"`
	Policyid  int                            `pulumi:"policyid"`
	Services  []GetFirewallDosPolicy6Service `pulumi:"services"`
	Srcaddrs  []GetFirewallDosPolicy6Srcaddr `pulumi:"srcaddrs"`
	Status    string                         `pulumi:"status"`
	Vdomparam *string                        `pulumi:"vdomparam"`
}

func LookupFirewallDosPolicy6Output(ctx *pulumi.Context, args LookupFirewallDosPolicy6OutputArgs, opts ...pulumi.InvokeOption) LookupFirewallDosPolicy6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallDosPolicy6Result, error) {
			args := v.(LookupFirewallDosPolicy6Args)
			r, err := LookupFirewallDosPolicy6(ctx, &args, opts...)
			var s LookupFirewallDosPolicy6Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallDosPolicy6ResultOutput)
}

// A collection of arguments for invoking GetFirewallDosPolicy6.
type LookupFirewallDosPolicy6OutputArgs struct {
	Policyid  pulumi.IntInput       `pulumi:"policyid"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallDosPolicy6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallDosPolicy6Args)(nil)).Elem()
}

// A collection of values returned by GetFirewallDosPolicy6.
type LookupFirewallDosPolicy6ResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallDosPolicy6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallDosPolicy6Result)(nil)).Elem()
}

func (o LookupFirewallDosPolicy6ResultOutput) ToLookupFirewallDosPolicy6ResultOutput() LookupFirewallDosPolicy6ResultOutput {
	return o
}

func (o LookupFirewallDosPolicy6ResultOutput) ToLookupFirewallDosPolicy6ResultOutputWithContext(ctx context.Context) LookupFirewallDosPolicy6ResultOutput {
	return o
}

func (o LookupFirewallDosPolicy6ResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFirewallDosPolicy6Result] {
	return pulumix.Output[LookupFirewallDosPolicy6Result]{
		OutputState: o.OutputState,
	}
}

func (o LookupFirewallDosPolicy6ResultOutput) Anomalies() GetFirewallDosPolicy6AnomalyArrayOutput {
	return o.ApplyT(func(v LookupFirewallDosPolicy6Result) []GetFirewallDosPolicy6Anomaly { return v.Anomalies }).(GetFirewallDosPolicy6AnomalyArrayOutput)
}

func (o LookupFirewallDosPolicy6ResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallDosPolicy6Result) string { return v.Comments }).(pulumi.StringOutput)
}

func (o LookupFirewallDosPolicy6ResultOutput) Dstaddrs() GetFirewallDosPolicy6DstaddrArrayOutput {
	return o.ApplyT(func(v LookupFirewallDosPolicy6Result) []GetFirewallDosPolicy6Dstaddr { return v.Dstaddrs }).(GetFirewallDosPolicy6DstaddrArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallDosPolicy6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallDosPolicy6Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFirewallDosPolicy6ResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallDosPolicy6Result) string { return v.Interface }).(pulumi.StringOutput)
}

func (o LookupFirewallDosPolicy6ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallDosPolicy6Result) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFirewallDosPolicy6ResultOutput) Policyid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallDosPolicy6Result) int { return v.Policyid }).(pulumi.IntOutput)
}

func (o LookupFirewallDosPolicy6ResultOutput) Services() GetFirewallDosPolicy6ServiceArrayOutput {
	return o.ApplyT(func(v LookupFirewallDosPolicy6Result) []GetFirewallDosPolicy6Service { return v.Services }).(GetFirewallDosPolicy6ServiceArrayOutput)
}

func (o LookupFirewallDosPolicy6ResultOutput) Srcaddrs() GetFirewallDosPolicy6SrcaddrArrayOutput {
	return o.ApplyT(func(v LookupFirewallDosPolicy6Result) []GetFirewallDosPolicy6Srcaddr { return v.Srcaddrs }).(GetFirewallDosPolicy6SrcaddrArrayOutput)
}

func (o LookupFirewallDosPolicy6ResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallDosPolicy6Result) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupFirewallDosPolicy6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallDosPolicy6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallDosPolicy6ResultOutput{})
}
