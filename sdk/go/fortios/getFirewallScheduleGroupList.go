// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetFirewallScheduleGroupList(ctx *pulumi.Context, args *GetFirewallScheduleGroupListArgs, opts ...pulumi.InvokeOption) (*GetFirewallScheduleGroupListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallScheduleGroupListResult
	err := ctx.Invoke("fortios:index/getFirewallScheduleGroupList:GetFirewallScheduleGroupList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallScheduleGroupList.
type GetFirewallScheduleGroupListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallScheduleGroupList.
type GetFirewallScheduleGroupListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallScheduleGroupListOutput(ctx *pulumi.Context, args GetFirewallScheduleGroupListOutputArgs, opts ...pulumi.InvokeOption) GetFirewallScheduleGroupListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallScheduleGroupListResult, error) {
			args := v.(GetFirewallScheduleGroupListArgs)
			r, err := GetFirewallScheduleGroupList(ctx, &args, opts...)
			var s GetFirewallScheduleGroupListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallScheduleGroupListResultOutput)
}

// A collection of arguments for invoking GetFirewallScheduleGroupList.
type GetFirewallScheduleGroupListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallScheduleGroupListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallScheduleGroupListArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallScheduleGroupList.
type GetFirewallScheduleGroupListResultOutput struct{ *pulumi.OutputState }

func (GetFirewallScheduleGroupListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallScheduleGroupListResult)(nil)).Elem()
}

func (o GetFirewallScheduleGroupListResultOutput) ToGetFirewallScheduleGroupListResultOutput() GetFirewallScheduleGroupListResultOutput {
	return o
}

func (o GetFirewallScheduleGroupListResultOutput) ToGetFirewallScheduleGroupListResultOutputWithContext(ctx context.Context) GetFirewallScheduleGroupListResultOutput {
	return o
}

func (o GetFirewallScheduleGroupListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallScheduleGroupListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallScheduleGroupListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallScheduleGroupListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFirewallScheduleGroupListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallScheduleGroupListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallScheduleGroupListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallScheduleGroupListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallScheduleGroupListResultOutput{})
}
