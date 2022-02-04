// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv4 virtual IP groups.
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
// 		trname1, err := fortios.NewFirewallVip(ctx, "trname1", &fortios.FirewallVipArgs{
// 			Extintf: pulumi.String("any"),
// 			Extport: pulumi.String("0-65535"),
// 			Extip:   pulumi.String("2.0.0.1-2.0.0.4"),
// 			Mappedips: FirewallVipMappedipArray{
// 				&FirewallVipMappedipArgs{
// 					Range: pulumi.String("3.0.0.0-3.0.0.3"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = fortios.NewFirewallVipgrp(ctx, "trname", &fortios.FirewallVipgrpArgs{
// 			Color:     pulumi.Int(0),
// 			Interface: pulumi.String("any"),
// 			Members: FirewallVipgrpMemberArray{
// 				&FirewallVipgrpMemberArgs{
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
// Firewall Vipgrp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallVipgrp:FirewallVipgrp labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallVipgrp struct {
	pulumi.CustomResourceState

	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color pulumi.IntOutput `pulumi:"color"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// interface
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members FirewallVipgrpMemberArrayOutput `pulumi:"members"`
	// VIP name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallVipgrp registers a new resource with the given unique name, arguments, and options.
func NewFirewallVipgrp(ctx *pulumi.Context,
	name string, args *FirewallVipgrpArgs, opts ...pulumi.ResourceOption) (*FirewallVipgrp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallVipgrp
	err := ctx.RegisterResource("fortios:index/firewallVipgrp:FirewallVipgrp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallVipgrp gets an existing FirewallVipgrp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallVipgrp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallVipgrpState, opts ...pulumi.ResourceOption) (*FirewallVipgrp, error) {
	var resource FirewallVipgrp
	err := ctx.ReadResource("fortios:index/firewallVipgrp:FirewallVipgrp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallVipgrp resources.
type firewallVipgrpState struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color *int `pulumi:"color"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// interface
	Interface *string `pulumi:"interface"`
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members []FirewallVipgrpMember `pulumi:"members"`
	// VIP name.
	Name *string `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallVipgrpState struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// interface
	Interface pulumi.StringPtrInput
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members FirewallVipgrpMemberArrayInput
	// VIP name.
	Name pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallVipgrpState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallVipgrpState)(nil)).Elem()
}

type firewallVipgrpArgs struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color *int `pulumi:"color"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// interface
	Interface string `pulumi:"interface"`
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members []FirewallVipgrpMember `pulumi:"members"`
	// VIP name.
	Name *string `pulumi:"name"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallVipgrp resource.
type FirewallVipgrpArgs struct {
	// Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
	Color pulumi.IntPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// interface
	Interface pulumi.StringInput
	// Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
	Members FirewallVipgrpMemberArrayInput
	// VIP name.
	Name pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallVipgrpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallVipgrpArgs)(nil)).Elem()
}

type FirewallVipgrpInput interface {
	pulumi.Input

	ToFirewallVipgrpOutput() FirewallVipgrpOutput
	ToFirewallVipgrpOutputWithContext(ctx context.Context) FirewallVipgrpOutput
}

func (*FirewallVipgrp) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallVipgrp)(nil)).Elem()
}

func (i *FirewallVipgrp) ToFirewallVipgrpOutput() FirewallVipgrpOutput {
	return i.ToFirewallVipgrpOutputWithContext(context.Background())
}

func (i *FirewallVipgrp) ToFirewallVipgrpOutputWithContext(ctx context.Context) FirewallVipgrpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVipgrpOutput)
}

// FirewallVipgrpArrayInput is an input type that accepts FirewallVipgrpArray and FirewallVipgrpArrayOutput values.
// You can construct a concrete instance of `FirewallVipgrpArrayInput` via:
//
//          FirewallVipgrpArray{ FirewallVipgrpArgs{...} }
type FirewallVipgrpArrayInput interface {
	pulumi.Input

	ToFirewallVipgrpArrayOutput() FirewallVipgrpArrayOutput
	ToFirewallVipgrpArrayOutputWithContext(context.Context) FirewallVipgrpArrayOutput
}

type FirewallVipgrpArray []FirewallVipgrpInput

func (FirewallVipgrpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallVipgrp)(nil)).Elem()
}

func (i FirewallVipgrpArray) ToFirewallVipgrpArrayOutput() FirewallVipgrpArrayOutput {
	return i.ToFirewallVipgrpArrayOutputWithContext(context.Background())
}

func (i FirewallVipgrpArray) ToFirewallVipgrpArrayOutputWithContext(ctx context.Context) FirewallVipgrpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVipgrpArrayOutput)
}

// FirewallVipgrpMapInput is an input type that accepts FirewallVipgrpMap and FirewallVipgrpMapOutput values.
// You can construct a concrete instance of `FirewallVipgrpMapInput` via:
//
//          FirewallVipgrpMap{ "key": FirewallVipgrpArgs{...} }
type FirewallVipgrpMapInput interface {
	pulumi.Input

	ToFirewallVipgrpMapOutput() FirewallVipgrpMapOutput
	ToFirewallVipgrpMapOutputWithContext(context.Context) FirewallVipgrpMapOutput
}

type FirewallVipgrpMap map[string]FirewallVipgrpInput

func (FirewallVipgrpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallVipgrp)(nil)).Elem()
}

func (i FirewallVipgrpMap) ToFirewallVipgrpMapOutput() FirewallVipgrpMapOutput {
	return i.ToFirewallVipgrpMapOutputWithContext(context.Background())
}

func (i FirewallVipgrpMap) ToFirewallVipgrpMapOutputWithContext(ctx context.Context) FirewallVipgrpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVipgrpMapOutput)
}

type FirewallVipgrpOutput struct{ *pulumi.OutputState }

func (FirewallVipgrpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallVipgrp)(nil)).Elem()
}

func (o FirewallVipgrpOutput) ToFirewallVipgrpOutput() FirewallVipgrpOutput {
	return o
}

func (o FirewallVipgrpOutput) ToFirewallVipgrpOutputWithContext(ctx context.Context) FirewallVipgrpOutput {
	return o
}

type FirewallVipgrpArrayOutput struct{ *pulumi.OutputState }

func (FirewallVipgrpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallVipgrp)(nil)).Elem()
}

func (o FirewallVipgrpArrayOutput) ToFirewallVipgrpArrayOutput() FirewallVipgrpArrayOutput {
	return o
}

func (o FirewallVipgrpArrayOutput) ToFirewallVipgrpArrayOutputWithContext(ctx context.Context) FirewallVipgrpArrayOutput {
	return o
}

func (o FirewallVipgrpArrayOutput) Index(i pulumi.IntInput) FirewallVipgrpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallVipgrp {
		return vs[0].([]*FirewallVipgrp)[vs[1].(int)]
	}).(FirewallVipgrpOutput)
}

type FirewallVipgrpMapOutput struct{ *pulumi.OutputState }

func (FirewallVipgrpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallVipgrp)(nil)).Elem()
}

func (o FirewallVipgrpMapOutput) ToFirewallVipgrpMapOutput() FirewallVipgrpMapOutput {
	return o
}

func (o FirewallVipgrpMapOutput) ToFirewallVipgrpMapOutputWithContext(ctx context.Context) FirewallVipgrpMapOutput {
	return o
}

func (o FirewallVipgrpMapOutput) MapIndex(k pulumi.StringInput) FirewallVipgrpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallVipgrp {
		return vs[0].(map[string]*FirewallVipgrp)[vs[1].(string)]
	}).(FirewallVipgrpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVipgrpInput)(nil)).Elem(), &FirewallVipgrp{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVipgrpArrayInput)(nil)).Elem(), FirewallVipgrpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVipgrpMapInput)(nil)).Elem(), FirewallVipgrpMap{})
	pulumi.RegisterOutputType(FirewallVipgrpOutput{})
	pulumi.RegisterOutputType(FirewallVipgrpArrayOutput{})
	pulumi.RegisterOutputType(FirewallVipgrpMapOutput{})
}
