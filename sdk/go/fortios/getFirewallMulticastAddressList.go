// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallMulticastAddress`.
func GetFirewallMulticastAddressList(ctx *pulumi.Context, args *GetFirewallMulticastAddressListArgs, opts ...pulumi.InvokeOption) (*GetFirewallMulticastAddressListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallMulticastAddressListResult
	err := ctx.Invoke("fortios:index/getFirewallMulticastAddressList:GetFirewallMulticastAddressList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallMulticastAddressList.
type GetFirewallMulticastAddressListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallMulticastAddressList.
type GetFirewallMulticastAddressListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallMulticastAddress`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallMulticastAddressListOutput(ctx *pulumi.Context, args GetFirewallMulticastAddressListOutputArgs, opts ...pulumi.InvokeOption) GetFirewallMulticastAddressListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallMulticastAddressListResult, error) {
			args := v.(GetFirewallMulticastAddressListArgs)
			r, err := GetFirewallMulticastAddressList(ctx, &args, opts...)
			var s GetFirewallMulticastAddressListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallMulticastAddressListResultOutput)
}

// A collection of arguments for invoking GetFirewallMulticastAddressList.
type GetFirewallMulticastAddressListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallMulticastAddressListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallMulticastAddressListArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallMulticastAddressList.
type GetFirewallMulticastAddressListResultOutput struct{ *pulumi.OutputState }

func (GetFirewallMulticastAddressListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallMulticastAddressListResult)(nil)).Elem()
}

func (o GetFirewallMulticastAddressListResultOutput) ToGetFirewallMulticastAddressListResultOutput() GetFirewallMulticastAddressListResultOutput {
	return o
}

func (o GetFirewallMulticastAddressListResultOutput) ToGetFirewallMulticastAddressListResultOutputWithContext(ctx context.Context) GetFirewallMulticastAddressListResultOutput {
	return o
}

func (o GetFirewallMulticastAddressListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallMulticastAddressListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallMulticastAddressListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallMulticastAddressListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `FirewallMulticastAddress`.
func (o GetFirewallMulticastAddressListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallMulticastAddressListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallMulticastAddressListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallMulticastAddressListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallMulticastAddressListResultOutput{})
}
