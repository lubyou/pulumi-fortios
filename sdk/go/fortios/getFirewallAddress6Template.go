// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewallAddress6Template(ctx *pulumi.Context, args *LookupFirewallAddress6TemplateArgs, opts ...pulumi.InvokeOption) (*LookupFirewallAddress6TemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFirewallAddress6TemplateResult
	err := ctx.Invoke("fortios:index/getFirewallAddress6Template:GetFirewallAddress6Template", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallAddress6Template.
type LookupFirewallAddress6TemplateArgs struct {
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallAddress6Template.
type LookupFirewallAddress6TemplateResult struct {
	FabricObject string `pulumi:"fabricObject"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string                                     `pulumi:"id"`
	Ip6                string                                     `pulumi:"ip6"`
	Name               string                                     `pulumi:"name"`
	SubnetSegmentCount int                                        `pulumi:"subnetSegmentCount"`
	SubnetSegments     []GetFirewallAddress6TemplateSubnetSegment `pulumi:"subnetSegments"`
	Vdomparam          *string                                    `pulumi:"vdomparam"`
}

func LookupFirewallAddress6TemplateOutput(ctx *pulumi.Context, args LookupFirewallAddress6TemplateOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallAddress6TemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallAddress6TemplateResult, error) {
			args := v.(LookupFirewallAddress6TemplateArgs)
			r, err := LookupFirewallAddress6Template(ctx, &args, opts...)
			var s LookupFirewallAddress6TemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallAddress6TemplateResultOutput)
}

// A collection of arguments for invoking GetFirewallAddress6Template.
type LookupFirewallAddress6TemplateOutputArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallAddress6TemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallAddress6TemplateArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallAddress6Template.
type LookupFirewallAddress6TemplateResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallAddress6TemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallAddress6TemplateResult)(nil)).Elem()
}

func (o LookupFirewallAddress6TemplateResultOutput) ToLookupFirewallAddress6TemplateResultOutput() LookupFirewallAddress6TemplateResultOutput {
	return o
}

func (o LookupFirewallAddress6TemplateResultOutput) ToLookupFirewallAddress6TemplateResultOutputWithContext(ctx context.Context) LookupFirewallAddress6TemplateResultOutput {
	return o
}

func (o LookupFirewallAddress6TemplateResultOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddress6TemplateResult) string { return v.FabricObject }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallAddress6TemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddress6TemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFirewallAddress6TemplateResultOutput) Ip6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddress6TemplateResult) string { return v.Ip6 }).(pulumi.StringOutput)
}

func (o LookupFirewallAddress6TemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddress6TemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFirewallAddress6TemplateResultOutput) SubnetSegmentCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallAddress6TemplateResult) int { return v.SubnetSegmentCount }).(pulumi.IntOutput)
}

func (o LookupFirewallAddress6TemplateResultOutput) SubnetSegments() GetFirewallAddress6TemplateSubnetSegmentArrayOutput {
	return o.ApplyT(func(v LookupFirewallAddress6TemplateResult) []GetFirewallAddress6TemplateSubnetSegment {
		return v.SubnetSegments
	}).(GetFirewallAddress6TemplateSubnetSegmentArrayOutput)
}

func (o LookupFirewallAddress6TemplateResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallAddress6TemplateResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallAddress6TemplateResultOutput{})
}
