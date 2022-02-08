// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure group of Internet Service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewFirewallInternetServiceGroup(ctx, "trname", &fortios.FirewallInternetServiceGroupArgs{
// 			Direction: pulumi.String("both"),
// 			Members: FirewallInternetServiceGroupMemberArray{
// 				&FirewallInternetServiceGroupMemberArgs{
// 					Id: pulumi.Int(65641),
// 				},
// 				&FirewallInternetServiceGroupMemberArgs{
// 					Id: pulumi.Int(65646),
// 				},
// 				&FirewallInternetServiceGroupMemberArgs{
// 					Id: pulumi.Int(196747),
// 				},
// 				&FirewallInternetServiceGroupMemberArgs{
// 					Id: pulumi.Int(327781),
// 				},
// 				&FirewallInternetServiceGroupMemberArgs{
// 					Id: pulumi.Int(327786),
// 				},
// 				&FirewallInternetServiceGroupMemberArgs{
// 					Id: pulumi.Int(327791),
// 				},
// 				&FirewallInternetServiceGroupMemberArgs{
// 					Id: pulumi.Int(327839),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Firewall InternetServiceGroup can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallInternetServiceGroup:FirewallInternetServiceGroup labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallInternetServiceGroup struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// How this service may be used (source, destination or both). Valid values: `source`, `destination`, `both`.
	Direction pulumi.StringOutput `pulumi:"direction"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Internet Service group member. The structure of `member` block is documented below.
	Members FirewallInternetServiceGroupMemberArrayOutput `pulumi:"members"`
	// Internet Service name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallInternetServiceGroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallInternetServiceGroup(ctx *pulumi.Context,
	name string, args *FirewallInternetServiceGroupArgs, opts ...pulumi.ResourceOption) (*FirewallInternetServiceGroup, error) {
	if args == nil {
		args = &FirewallInternetServiceGroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallInternetServiceGroup
	err := ctx.RegisterResource("fortios:index/firewallInternetServiceGroup:FirewallInternetServiceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallInternetServiceGroup gets an existing FirewallInternetServiceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallInternetServiceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallInternetServiceGroupState, opts ...pulumi.ResourceOption) (*FirewallInternetServiceGroup, error) {
	var resource FirewallInternetServiceGroup
	err := ctx.ReadResource("fortios:index/firewallInternetServiceGroup:FirewallInternetServiceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallInternetServiceGroup resources.
type firewallInternetServiceGroupState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// How this service may be used (source, destination or both). Valid values: `source`, `destination`, `both`.
	Direction *string `pulumi:"direction"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Internet Service group member. The structure of `member` block is documented below.
	Members []FirewallInternetServiceGroupMember `pulumi:"members"`
	// Internet Service name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallInternetServiceGroupState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// How this service may be used (source, destination or both). Valid values: `source`, `destination`, `both`.
	Direction pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Internet Service group member. The structure of `member` block is documented below.
	Members FirewallInternetServiceGroupMemberArrayInput
	// Internet Service name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetServiceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceGroupState)(nil)).Elem()
}

type firewallInternetServiceGroupArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// How this service may be used (source, destination or both). Valid values: `source`, `destination`, `both`.
	Direction *string `pulumi:"direction"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Internet Service group member. The structure of `member` block is documented below.
	Members []FirewallInternetServiceGroupMember `pulumi:"members"`
	// Internet Service name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallInternetServiceGroup resource.
type FirewallInternetServiceGroupArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// How this service may be used (source, destination or both). Valid values: `source`, `destination`, `both`.
	Direction pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Internet Service group member. The structure of `member` block is documented below.
	Members FirewallInternetServiceGroupMemberArrayInput
	// Internet Service name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetServiceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceGroupArgs)(nil)).Elem()
}

type FirewallInternetServiceGroupInput interface {
	pulumi.Input

	ToFirewallInternetServiceGroupOutput() FirewallInternetServiceGroupOutput
	ToFirewallInternetServiceGroupOutputWithContext(ctx context.Context) FirewallInternetServiceGroupOutput
}

func (*FirewallInternetServiceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceGroup)(nil)).Elem()
}

func (i *FirewallInternetServiceGroup) ToFirewallInternetServiceGroupOutput() FirewallInternetServiceGroupOutput {
	return i.ToFirewallInternetServiceGroupOutputWithContext(context.Background())
}

func (i *FirewallInternetServiceGroup) ToFirewallInternetServiceGroupOutputWithContext(ctx context.Context) FirewallInternetServiceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceGroupOutput)
}

// FirewallInternetServiceGroupArrayInput is an input type that accepts FirewallInternetServiceGroupArray and FirewallInternetServiceGroupArrayOutput values.
// You can construct a concrete instance of `FirewallInternetServiceGroupArrayInput` via:
//
//          FirewallInternetServiceGroupArray{ FirewallInternetServiceGroupArgs{...} }
type FirewallInternetServiceGroupArrayInput interface {
	pulumi.Input

	ToFirewallInternetServiceGroupArrayOutput() FirewallInternetServiceGroupArrayOutput
	ToFirewallInternetServiceGroupArrayOutputWithContext(context.Context) FirewallInternetServiceGroupArrayOutput
}

type FirewallInternetServiceGroupArray []FirewallInternetServiceGroupInput

func (FirewallInternetServiceGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceGroup)(nil)).Elem()
}

func (i FirewallInternetServiceGroupArray) ToFirewallInternetServiceGroupArrayOutput() FirewallInternetServiceGroupArrayOutput {
	return i.ToFirewallInternetServiceGroupArrayOutputWithContext(context.Background())
}

func (i FirewallInternetServiceGroupArray) ToFirewallInternetServiceGroupArrayOutputWithContext(ctx context.Context) FirewallInternetServiceGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceGroupArrayOutput)
}

// FirewallInternetServiceGroupMapInput is an input type that accepts FirewallInternetServiceGroupMap and FirewallInternetServiceGroupMapOutput values.
// You can construct a concrete instance of `FirewallInternetServiceGroupMapInput` via:
//
//          FirewallInternetServiceGroupMap{ "key": FirewallInternetServiceGroupArgs{...} }
type FirewallInternetServiceGroupMapInput interface {
	pulumi.Input

	ToFirewallInternetServiceGroupMapOutput() FirewallInternetServiceGroupMapOutput
	ToFirewallInternetServiceGroupMapOutputWithContext(context.Context) FirewallInternetServiceGroupMapOutput
}

type FirewallInternetServiceGroupMap map[string]FirewallInternetServiceGroupInput

func (FirewallInternetServiceGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceGroup)(nil)).Elem()
}

func (i FirewallInternetServiceGroupMap) ToFirewallInternetServiceGroupMapOutput() FirewallInternetServiceGroupMapOutput {
	return i.ToFirewallInternetServiceGroupMapOutputWithContext(context.Background())
}

func (i FirewallInternetServiceGroupMap) ToFirewallInternetServiceGroupMapOutputWithContext(ctx context.Context) FirewallInternetServiceGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceGroupMapOutput)
}

type FirewallInternetServiceGroupOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceGroup)(nil)).Elem()
}

func (o FirewallInternetServiceGroupOutput) ToFirewallInternetServiceGroupOutput() FirewallInternetServiceGroupOutput {
	return o
}

func (o FirewallInternetServiceGroupOutput) ToFirewallInternetServiceGroupOutputWithContext(ctx context.Context) FirewallInternetServiceGroupOutput {
	return o
}

type FirewallInternetServiceGroupArrayOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceGroup)(nil)).Elem()
}

func (o FirewallInternetServiceGroupArrayOutput) ToFirewallInternetServiceGroupArrayOutput() FirewallInternetServiceGroupArrayOutput {
	return o
}

func (o FirewallInternetServiceGroupArrayOutput) ToFirewallInternetServiceGroupArrayOutputWithContext(ctx context.Context) FirewallInternetServiceGroupArrayOutput {
	return o
}

func (o FirewallInternetServiceGroupArrayOutput) Index(i pulumi.IntInput) FirewallInternetServiceGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallInternetServiceGroup {
		return vs[0].([]*FirewallInternetServiceGroup)[vs[1].(int)]
	}).(FirewallInternetServiceGroupOutput)
}

type FirewallInternetServiceGroupMapOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceGroup)(nil)).Elem()
}

func (o FirewallInternetServiceGroupMapOutput) ToFirewallInternetServiceGroupMapOutput() FirewallInternetServiceGroupMapOutput {
	return o
}

func (o FirewallInternetServiceGroupMapOutput) ToFirewallInternetServiceGroupMapOutputWithContext(ctx context.Context) FirewallInternetServiceGroupMapOutput {
	return o
}

func (o FirewallInternetServiceGroupMapOutput) MapIndex(k pulumi.StringInput) FirewallInternetServiceGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallInternetServiceGroup {
		return vs[0].(map[string]*FirewallInternetServiceGroup)[vs[1].(string)]
	}).(FirewallInternetServiceGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceGroupInput)(nil)).Elem(), &FirewallInternetServiceGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceGroupArrayInput)(nil)).Elem(), FirewallInternetServiceGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceGroupMapInput)(nil)).Elem(), FirewallInternetServiceGroupMap{})
	pulumi.RegisterOutputType(FirewallInternetServiceGroupOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceGroupArrayOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceGroupMapOutput{})
}
