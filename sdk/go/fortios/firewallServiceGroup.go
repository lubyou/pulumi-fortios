// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure service groups.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		trname1, err := fortios.NewFirewallServiceCustom(ctx, "trname1", &fortios.FirewallServiceCustomArgs{
// 			AppServiceType:    pulumi.String("disable"),
// 			Category:          pulumi.String("General"),
// 			CheckResetRange:   pulumi.String("default"),
// 			Color:             pulumi.Int(0),
// 			Helper:            pulumi.String("auto"),
// 			Iprange:           pulumi.String("0.0.0.0"),
// 			Protocol:          pulumi.String("TCP/UDP/SCTP"),
// 			ProtocolNumber:    pulumi.Int(6),
// 			Proxy:             pulumi.String("disable"),
// 			TcpHalfcloseTimer: pulumi.Int(0),
// 			TcpHalfopenTimer:  pulumi.Int(0),
// 			TcpPortrange:      pulumi.String("223-332"),
// 			TcpTimewaitTimer:  pulumi.Int(0),
// 			UdpIdleTimer:      pulumi.Int(0),
// 			Visibility:        pulumi.String("enable"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = fortios.NewFirewallServiceGroup(ctx, "trname", &fortios.FirewallServiceGroupArgs{
// 			Color: pulumi.Int(0),
// 			Proxy: pulumi.String("disable"),
// 			Members: FirewallServiceGroupMemberArray{
// 				&FirewallServiceGroupMemberArgs{
// 					Name: trname1.Name,
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
// FirewallService Group can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallServiceGroup:FirewallServiceGroup labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallServiceGroup struct {
	pulumi.CustomResourceState

	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringOutput `pulumi:"fabricObject"`
	// Service objects contained within the group. The structure of `member` block is documented below.
	Members FirewallServiceGroupMemberArrayOutput `pulumi:"members"`
	// Address name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable web proxy service group. Valid values: `enable`, `disable`.
	Proxy pulumi.StringOutput `pulumi:"proxy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallServiceGroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallServiceGroup(ctx *pulumi.Context,
	name string, args *FirewallServiceGroupArgs, opts ...pulumi.ResourceOption) (*FirewallServiceGroup, error) {
	if args == nil {
		args = &FirewallServiceGroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallServiceGroup
	err := ctx.RegisterResource("fortios:index/firewallServiceGroup:FirewallServiceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallServiceGroup gets an existing FirewallServiceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallServiceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallServiceGroupState, opts ...pulumi.ResourceOption) (*FirewallServiceGroup, error) {
	var resource FirewallServiceGroup
	err := ctx.ReadResource("fortios:index/firewallServiceGroup:FirewallServiceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallServiceGroup resources.
type firewallServiceGroupState struct {
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Service objects contained within the group. The structure of `member` block is documented below.
	Members []FirewallServiceGroupMember `pulumi:"members"`
	// Address name.
	Name *string `pulumi:"name"`
	// Enable/disable web proxy service group. Valid values: `enable`, `disable`.
	Proxy *string `pulumi:"proxy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallServiceGroupState struct {
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Service objects contained within the group. The structure of `member` block is documented below.
	Members FirewallServiceGroupMemberArrayInput
	// Address name.
	Name pulumi.StringPtrInput
	// Enable/disable web proxy service group. Valid values: `enable`, `disable`.
	Proxy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallServiceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallServiceGroupState)(nil)).Elem()
}

type firewallServiceGroupArgs struct {
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Service objects contained within the group. The structure of `member` block is documented below.
	Members []FirewallServiceGroupMember `pulumi:"members"`
	// Address name.
	Name *string `pulumi:"name"`
	// Enable/disable web proxy service group. Valid values: `enable`, `disable`.
	Proxy *string `pulumi:"proxy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallServiceGroup resource.
type FirewallServiceGroupArgs struct {
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Service objects contained within the group. The structure of `member` block is documented below.
	Members FirewallServiceGroupMemberArrayInput
	// Address name.
	Name pulumi.StringPtrInput
	// Enable/disable web proxy service group. Valid values: `enable`, `disable`.
	Proxy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallServiceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallServiceGroupArgs)(nil)).Elem()
}

type FirewallServiceGroupInput interface {
	pulumi.Input

	ToFirewallServiceGroupOutput() FirewallServiceGroupOutput
	ToFirewallServiceGroupOutputWithContext(ctx context.Context) FirewallServiceGroupOutput
}

func (*FirewallServiceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallServiceGroup)(nil)).Elem()
}

func (i *FirewallServiceGroup) ToFirewallServiceGroupOutput() FirewallServiceGroupOutput {
	return i.ToFirewallServiceGroupOutputWithContext(context.Background())
}

func (i *FirewallServiceGroup) ToFirewallServiceGroupOutputWithContext(ctx context.Context) FirewallServiceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallServiceGroupOutput)
}

// FirewallServiceGroupArrayInput is an input type that accepts FirewallServiceGroupArray and FirewallServiceGroupArrayOutput values.
// You can construct a concrete instance of `FirewallServiceGroupArrayInput` via:
//
//          FirewallServiceGroupArray{ FirewallServiceGroupArgs{...} }
type FirewallServiceGroupArrayInput interface {
	pulumi.Input

	ToFirewallServiceGroupArrayOutput() FirewallServiceGroupArrayOutput
	ToFirewallServiceGroupArrayOutputWithContext(context.Context) FirewallServiceGroupArrayOutput
}

type FirewallServiceGroupArray []FirewallServiceGroupInput

func (FirewallServiceGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallServiceGroup)(nil)).Elem()
}

func (i FirewallServiceGroupArray) ToFirewallServiceGroupArrayOutput() FirewallServiceGroupArrayOutput {
	return i.ToFirewallServiceGroupArrayOutputWithContext(context.Background())
}

func (i FirewallServiceGroupArray) ToFirewallServiceGroupArrayOutputWithContext(ctx context.Context) FirewallServiceGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallServiceGroupArrayOutput)
}

// FirewallServiceGroupMapInput is an input type that accepts FirewallServiceGroupMap and FirewallServiceGroupMapOutput values.
// You can construct a concrete instance of `FirewallServiceGroupMapInput` via:
//
//          FirewallServiceGroupMap{ "key": FirewallServiceGroupArgs{...} }
type FirewallServiceGroupMapInput interface {
	pulumi.Input

	ToFirewallServiceGroupMapOutput() FirewallServiceGroupMapOutput
	ToFirewallServiceGroupMapOutputWithContext(context.Context) FirewallServiceGroupMapOutput
}

type FirewallServiceGroupMap map[string]FirewallServiceGroupInput

func (FirewallServiceGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallServiceGroup)(nil)).Elem()
}

func (i FirewallServiceGroupMap) ToFirewallServiceGroupMapOutput() FirewallServiceGroupMapOutput {
	return i.ToFirewallServiceGroupMapOutputWithContext(context.Background())
}

func (i FirewallServiceGroupMap) ToFirewallServiceGroupMapOutputWithContext(ctx context.Context) FirewallServiceGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallServiceGroupMapOutput)
}

type FirewallServiceGroupOutput struct{ *pulumi.OutputState }

func (FirewallServiceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallServiceGroup)(nil)).Elem()
}

func (o FirewallServiceGroupOutput) ToFirewallServiceGroupOutput() FirewallServiceGroupOutput {
	return o
}

func (o FirewallServiceGroupOutput) ToFirewallServiceGroupOutputWithContext(ctx context.Context) FirewallServiceGroupOutput {
	return o
}

type FirewallServiceGroupArrayOutput struct{ *pulumi.OutputState }

func (FirewallServiceGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallServiceGroup)(nil)).Elem()
}

func (o FirewallServiceGroupArrayOutput) ToFirewallServiceGroupArrayOutput() FirewallServiceGroupArrayOutput {
	return o
}

func (o FirewallServiceGroupArrayOutput) ToFirewallServiceGroupArrayOutputWithContext(ctx context.Context) FirewallServiceGroupArrayOutput {
	return o
}

func (o FirewallServiceGroupArrayOutput) Index(i pulumi.IntInput) FirewallServiceGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallServiceGroup {
		return vs[0].([]*FirewallServiceGroup)[vs[1].(int)]
	}).(FirewallServiceGroupOutput)
}

type FirewallServiceGroupMapOutput struct{ *pulumi.OutputState }

func (FirewallServiceGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallServiceGroup)(nil)).Elem()
}

func (o FirewallServiceGroupMapOutput) ToFirewallServiceGroupMapOutput() FirewallServiceGroupMapOutput {
	return o
}

func (o FirewallServiceGroupMapOutput) ToFirewallServiceGroupMapOutputWithContext(ctx context.Context) FirewallServiceGroupMapOutput {
	return o
}

func (o FirewallServiceGroupMapOutput) MapIndex(k pulumi.StringInput) FirewallServiceGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallServiceGroup {
		return vs[0].(map[string]*FirewallServiceGroup)[vs[1].(string)]
	}).(FirewallServiceGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallServiceGroupInput)(nil)).Elem(), &FirewallServiceGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallServiceGroupArrayInput)(nil)).Elem(), FirewallServiceGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallServiceGroupMapInput)(nil)).Elem(), FirewallServiceGroupMap{})
	pulumi.RegisterOutputType(FirewallServiceGroupOutput{})
	pulumi.RegisterOutputType(FirewallServiceGroupArrayOutput{})
	pulumi.RegisterOutputType(FirewallServiceGroupMapOutput{})
}
