// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallPolicy46`.
func GetFirewallPolicy46List(ctx *pulumi.Context, args *GetFirewallPolicy46ListArgs, opts ...pulumi.InvokeOption) (*GetFirewallPolicy46ListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallPolicy46ListResult
	err := ctx.Invoke("fortios:index/getFirewallPolicy46List:GetFirewallPolicy46List", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallPolicy46List.
type GetFirewallPolicy46ListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallPolicy46List.
type GetFirewallPolicy46ListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallPolicy46`.
	Policyidlists []int   `pulumi:"policyidlists"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

func GetFirewallPolicy46ListOutput(ctx *pulumi.Context, args GetFirewallPolicy46ListOutputArgs, opts ...pulumi.InvokeOption) GetFirewallPolicy46ListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallPolicy46ListResult, error) {
			args := v.(GetFirewallPolicy46ListArgs)
			r, err := GetFirewallPolicy46List(ctx, &args, opts...)
			var s GetFirewallPolicy46ListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallPolicy46ListResultOutput)
}

// A collection of arguments for invoking GetFirewallPolicy46List.
type GetFirewallPolicy46ListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallPolicy46ListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallPolicy46ListArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallPolicy46List.
type GetFirewallPolicy46ListResultOutput struct{ *pulumi.OutputState }

func (GetFirewallPolicy46ListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallPolicy46ListResult)(nil)).Elem()
}

func (o GetFirewallPolicy46ListResultOutput) ToGetFirewallPolicy46ListResultOutput() GetFirewallPolicy46ListResultOutput {
	return o
}

func (o GetFirewallPolicy46ListResultOutput) ToGetFirewallPolicy46ListResultOutputWithContext(ctx context.Context) GetFirewallPolicy46ListResultOutput {
	return o
}

func (o GetFirewallPolicy46ListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallPolicy46ListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallPolicy46ListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallPolicy46ListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `FirewallPolicy46`.
func (o GetFirewallPolicy46ListResultOutput) Policyidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetFirewallPolicy46ListResult) []int { return v.Policyidlists }).(pulumi.IntArrayOutput)
}

func (o GetFirewallPolicy46ListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallPolicy46ListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallPolicy46ListResultOutput{})
}
