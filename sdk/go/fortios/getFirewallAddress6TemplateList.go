// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetFirewallAddress6TemplateList(ctx *pulumi.Context, args *GetFirewallAddress6TemplateListArgs, opts ...pulumi.InvokeOption) (*GetFirewallAddress6TemplateListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFirewallAddress6TemplateListResult
	err := ctx.Invoke("fortios:index/getFirewallAddress6TemplateList:GetFirewallAddress6TemplateList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallAddress6TemplateList.
type GetFirewallAddress6TemplateListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallAddress6TemplateList.
type GetFirewallAddress6TemplateListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallAddress6TemplateListOutput(ctx *pulumi.Context, args GetFirewallAddress6TemplateListOutputArgs, opts ...pulumi.InvokeOption) GetFirewallAddress6TemplateListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallAddress6TemplateListResult, error) {
			args := v.(GetFirewallAddress6TemplateListArgs)
			r, err := GetFirewallAddress6TemplateList(ctx, &args, opts...)
			var s GetFirewallAddress6TemplateListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallAddress6TemplateListResultOutput)
}

// A collection of arguments for invoking GetFirewallAddress6TemplateList.
type GetFirewallAddress6TemplateListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallAddress6TemplateListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallAddress6TemplateListArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallAddress6TemplateList.
type GetFirewallAddress6TemplateListResultOutput struct{ *pulumi.OutputState }

func (GetFirewallAddress6TemplateListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallAddress6TemplateListResult)(nil)).Elem()
}

func (o GetFirewallAddress6TemplateListResultOutput) ToGetFirewallAddress6TemplateListResultOutput() GetFirewallAddress6TemplateListResultOutput {
	return o
}

func (o GetFirewallAddress6TemplateListResultOutput) ToGetFirewallAddress6TemplateListResultOutputWithContext(ctx context.Context) GetFirewallAddress6TemplateListResultOutput {
	return o
}

func (o GetFirewallAddress6TemplateListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallAddress6TemplateListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallAddress6TemplateListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallAddress6TemplateListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFirewallAddress6TemplateListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallAddress6TemplateListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallAddress6TemplateListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallAddress6TemplateListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallAddress6TemplateListResultOutput{})
}
