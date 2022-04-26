// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv4 address groups.
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
// 		trname1, err := fortios.NewFirewallAddress(ctx, "trname1", &fortios.FirewallAddressArgs{
// 			AllowRouting: pulumi.String("disable"),
// 			CacheTtl:     pulumi.Int(0),
// 			Color:        pulumi.Int(0),
// 			EndIp:        pulumi.String("255.0.0.0"),
// 			StartIp:      pulumi.String("12.0.0.0"),
// 			Subnet:       pulumi.String("12.0.0.0 255.0.0.0"),
// 			Type:         pulumi.String("ipmask"),
// 			Visibility:   pulumi.String("enable"),
// 			Wildcard:     pulumi.String("12.0.0.0 255.0.0.0"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = fortios.NewFirewallAddrgrp(ctx, "trname", &fortios.FirewallAddrgrpArgs{
// 			AllowRouting: pulumi.String("disable"),
// 			Color:        pulumi.Int(0),
// 			Exclude:      pulumi.String("disable"),
// 			Visibility:   pulumi.String("enable"),
// 			Members: FirewallAddrgrpMemberArray{
// 				&FirewallAddrgrpMemberArgs{
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
// Firewall Addrgrp can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/firewallAddrgrp:FirewallAddrgrp labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/firewallAddrgrp:FirewallAddrgrp labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallAddrgrp struct {
	pulumi.CustomResourceState

	// Enable/disable use of this group in the static route configuration. Valid values: `enable`, `disable`.
	AllowRouting pulumi.StringOutput `pulumi:"allowRouting"`
	// Tag category.
	Category pulumi.StringOutput `pulumi:"category"`
	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable address exclusion. Valid values: `enable`, `disable`.
	Exclude pulumi.StringOutput `pulumi:"exclude"`
	// Address exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers FirewallAddrgrpExcludeMemberArrayOutput `pulumi:"excludeMembers"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringOutput `pulumi:"fabricObject"`
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members FirewallAddrgrpMemberArrayOutput `pulumi:"members"`
	// Tag name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallAddrgrpTaggingArrayOutput `pulumi:"taggings"`
	// Address group type. Valid values: `default`, `folder`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewFirewallAddrgrp registers a new resource with the given unique name, arguments, and options.
func NewFirewallAddrgrp(ctx *pulumi.Context,
	name string, args *FirewallAddrgrpArgs, opts ...pulumi.ResourceOption) (*FirewallAddrgrp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallAddrgrp
	err := ctx.RegisterResource("fortios:index/firewallAddrgrp:FirewallAddrgrp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallAddrgrp gets an existing FirewallAddrgrp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallAddrgrp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallAddrgrpState, opts ...pulumi.ResourceOption) (*FirewallAddrgrp, error) {
	var resource FirewallAddrgrp
	err := ctx.ReadResource("fortios:index/firewallAddrgrp:FirewallAddrgrp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallAddrgrp resources.
type firewallAddrgrpState struct {
	// Enable/disable use of this group in the static route configuration. Valid values: `enable`, `disable`.
	AllowRouting *string `pulumi:"allowRouting"`
	// Tag category.
	Category *string `pulumi:"category"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable address exclusion. Valid values: `enable`, `disable`.
	Exclude *string `pulumi:"exclude"`
	// Address exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers []FirewallAddrgrpExcludeMember `pulumi:"excludeMembers"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members []FirewallAddrgrpMember `pulumi:"members"`
	// Tag name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []FirewallAddrgrpTagging `pulumi:"taggings"`
	// Address group type. Valid values: `default`, `folder`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

type FirewallAddrgrpState struct {
	// Enable/disable use of this group in the static route configuration. Valid values: `enable`, `disable`.
	AllowRouting pulumi.StringPtrInput
	// Tag category.
	Category pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable address exclusion. Valid values: `enable`, `disable`.
	Exclude pulumi.StringPtrInput
	// Address exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers FirewallAddrgrpExcludeMemberArrayInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members FirewallAddrgrpMemberArrayInput
	// Tag name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallAddrgrpTaggingArrayInput
	// Address group type. Valid values: `default`, `folder`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (FirewallAddrgrpState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallAddrgrpState)(nil)).Elem()
}

type firewallAddrgrpArgs struct {
	// Enable/disable use of this group in the static route configuration. Valid values: `enable`, `disable`.
	AllowRouting *string `pulumi:"allowRouting"`
	// Tag category.
	Category *string `pulumi:"category"`
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable address exclusion. Valid values: `enable`, `disable`.
	Exclude *string `pulumi:"exclude"`
	// Address exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers []FirewallAddrgrpExcludeMember `pulumi:"excludeMembers"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members []FirewallAddrgrpMember `pulumi:"members"`
	// Tag name.
	Name *string `pulumi:"name"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []FirewallAddrgrpTagging `pulumi:"taggings"`
	// Address group type. Valid values: `default`, `folder`.
	Type *string `pulumi:"type"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a FirewallAddrgrp resource.
type FirewallAddrgrpArgs struct {
	// Enable/disable use of this group in the static route configuration. Valid values: `enable`, `disable`.
	AllowRouting pulumi.StringPtrInput
	// Tag category.
	Category pulumi.StringPtrInput
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable address exclusion. Valid values: `enable`, `disable`.
	Exclude pulumi.StringPtrInput
	// Address exclusion member. The structure of `excludeMember` block is documented below.
	ExcludeMembers FirewallAddrgrpExcludeMemberArrayInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Address objects contained within the group. The structure of `member` block is documented below.
	Members FirewallAddrgrpMemberArrayInput
	// Tag name.
	Name pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallAddrgrpTaggingArrayInput
	// Address group type. Valid values: `default`, `folder`.
	Type pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (FirewallAddrgrpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallAddrgrpArgs)(nil)).Elem()
}

type FirewallAddrgrpInput interface {
	pulumi.Input

	ToFirewallAddrgrpOutput() FirewallAddrgrpOutput
	ToFirewallAddrgrpOutputWithContext(ctx context.Context) FirewallAddrgrpOutput
}

func (*FirewallAddrgrp) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallAddrgrp)(nil)).Elem()
}

func (i *FirewallAddrgrp) ToFirewallAddrgrpOutput() FirewallAddrgrpOutput {
	return i.ToFirewallAddrgrpOutputWithContext(context.Background())
}

func (i *FirewallAddrgrp) ToFirewallAddrgrpOutputWithContext(ctx context.Context) FirewallAddrgrpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAddrgrpOutput)
}

// FirewallAddrgrpArrayInput is an input type that accepts FirewallAddrgrpArray and FirewallAddrgrpArrayOutput values.
// You can construct a concrete instance of `FirewallAddrgrpArrayInput` via:
//
//          FirewallAddrgrpArray{ FirewallAddrgrpArgs{...} }
type FirewallAddrgrpArrayInput interface {
	pulumi.Input

	ToFirewallAddrgrpArrayOutput() FirewallAddrgrpArrayOutput
	ToFirewallAddrgrpArrayOutputWithContext(context.Context) FirewallAddrgrpArrayOutput
}

type FirewallAddrgrpArray []FirewallAddrgrpInput

func (FirewallAddrgrpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallAddrgrp)(nil)).Elem()
}

func (i FirewallAddrgrpArray) ToFirewallAddrgrpArrayOutput() FirewallAddrgrpArrayOutput {
	return i.ToFirewallAddrgrpArrayOutputWithContext(context.Background())
}

func (i FirewallAddrgrpArray) ToFirewallAddrgrpArrayOutputWithContext(ctx context.Context) FirewallAddrgrpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAddrgrpArrayOutput)
}

// FirewallAddrgrpMapInput is an input type that accepts FirewallAddrgrpMap and FirewallAddrgrpMapOutput values.
// You can construct a concrete instance of `FirewallAddrgrpMapInput` via:
//
//          FirewallAddrgrpMap{ "key": FirewallAddrgrpArgs{...} }
type FirewallAddrgrpMapInput interface {
	pulumi.Input

	ToFirewallAddrgrpMapOutput() FirewallAddrgrpMapOutput
	ToFirewallAddrgrpMapOutputWithContext(context.Context) FirewallAddrgrpMapOutput
}

type FirewallAddrgrpMap map[string]FirewallAddrgrpInput

func (FirewallAddrgrpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallAddrgrp)(nil)).Elem()
}

func (i FirewallAddrgrpMap) ToFirewallAddrgrpMapOutput() FirewallAddrgrpMapOutput {
	return i.ToFirewallAddrgrpMapOutputWithContext(context.Background())
}

func (i FirewallAddrgrpMap) ToFirewallAddrgrpMapOutputWithContext(ctx context.Context) FirewallAddrgrpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAddrgrpMapOutput)
}

type FirewallAddrgrpOutput struct{ *pulumi.OutputState }

func (FirewallAddrgrpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallAddrgrp)(nil)).Elem()
}

func (o FirewallAddrgrpOutput) ToFirewallAddrgrpOutput() FirewallAddrgrpOutput {
	return o
}

func (o FirewallAddrgrpOutput) ToFirewallAddrgrpOutputWithContext(ctx context.Context) FirewallAddrgrpOutput {
	return o
}

type FirewallAddrgrpArrayOutput struct{ *pulumi.OutputState }

func (FirewallAddrgrpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallAddrgrp)(nil)).Elem()
}

func (o FirewallAddrgrpArrayOutput) ToFirewallAddrgrpArrayOutput() FirewallAddrgrpArrayOutput {
	return o
}

func (o FirewallAddrgrpArrayOutput) ToFirewallAddrgrpArrayOutputWithContext(ctx context.Context) FirewallAddrgrpArrayOutput {
	return o
}

func (o FirewallAddrgrpArrayOutput) Index(i pulumi.IntInput) FirewallAddrgrpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallAddrgrp {
		return vs[0].([]*FirewallAddrgrp)[vs[1].(int)]
	}).(FirewallAddrgrpOutput)
}

type FirewallAddrgrpMapOutput struct{ *pulumi.OutputState }

func (FirewallAddrgrpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallAddrgrp)(nil)).Elem()
}

func (o FirewallAddrgrpMapOutput) ToFirewallAddrgrpMapOutput() FirewallAddrgrpMapOutput {
	return o
}

func (o FirewallAddrgrpMapOutput) ToFirewallAddrgrpMapOutputWithContext(ctx context.Context) FirewallAddrgrpMapOutput {
	return o
}

func (o FirewallAddrgrpMapOutput) MapIndex(k pulumi.StringInput) FirewallAddrgrpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallAddrgrp {
		return vs[0].(map[string]*FirewallAddrgrp)[vs[1].(string)]
	}).(FirewallAddrgrpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallAddrgrpInput)(nil)).Elem(), &FirewallAddrgrp{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallAddrgrpArrayInput)(nil)).Elem(), FirewallAddrgrpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallAddrgrpMapInput)(nil)).Elem(), FirewallAddrgrpMap{})
	pulumi.RegisterOutputType(FirewallAddrgrpOutput{})
	pulumi.RegisterOutputType(FirewallAddrgrpArrayOutput{})
	pulumi.RegisterOutputType(FirewallAddrgrpMapOutput{})
}
