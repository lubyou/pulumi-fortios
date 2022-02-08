// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IP to MAC binding settings.
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
// 		_, err := fortios.NewFirewallIpmacbindingSetting(ctx, "trname", &fortios.FirewallIpmacbindingSettingArgs{
// 			Bindthroughfw: pulumi.String("disable"),
// 			Bindtofw:      pulumi.String("disable"),
// 			Undefinedhost: pulumi.String("block"),
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
// FirewallIpmacbinding Setting can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallIpmacbindingSetting:FirewallIpmacbindingSetting labelname FirewallIpmacbindingSetting
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallIpmacbindingSetting struct {
	pulumi.CustomResourceState

	// Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
	Bindthroughfw pulumi.StringOutput `pulumi:"bindthroughfw"`
	// Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
	Bindtofw pulumi.StringOutput `pulumi:"bindtofw"`
	// Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
	Undefinedhost pulumi.StringOutput `pulumi:"undefinedhost"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallIpmacbindingSetting registers a new resource with the given unique name, arguments, and options.
func NewFirewallIpmacbindingSetting(ctx *pulumi.Context,
	name string, args *FirewallIpmacbindingSettingArgs, opts ...pulumi.ResourceOption) (*FirewallIpmacbindingSetting, error) {
	if args == nil {
		args = &FirewallIpmacbindingSettingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallIpmacbindingSetting
	err := ctx.RegisterResource("fortios:index/firewallIpmacbindingSetting:FirewallIpmacbindingSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallIpmacbindingSetting gets an existing FirewallIpmacbindingSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallIpmacbindingSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallIpmacbindingSettingState, opts ...pulumi.ResourceOption) (*FirewallIpmacbindingSetting, error) {
	var resource FirewallIpmacbindingSetting
	err := ctx.ReadResource("fortios:index/firewallIpmacbindingSetting:FirewallIpmacbindingSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallIpmacbindingSetting resources.
type firewallIpmacbindingSettingState struct {
	// Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
	Bindthroughfw *string `pulumi:"bindthroughfw"`
	// Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
	Bindtofw *string `pulumi:"bindtofw"`
	// Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
	Undefinedhost *string `pulumi:"undefinedhost"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallIpmacbindingSettingState struct {
	// Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
	Bindthroughfw pulumi.StringPtrInput
	// Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
	Bindtofw pulumi.StringPtrInput
	// Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
	Undefinedhost pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallIpmacbindingSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallIpmacbindingSettingState)(nil)).Elem()
}

type firewallIpmacbindingSettingArgs struct {
	// Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
	Bindthroughfw *string `pulumi:"bindthroughfw"`
	// Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
	Bindtofw *string `pulumi:"bindtofw"`
	// Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
	Undefinedhost *string `pulumi:"undefinedhost"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallIpmacbindingSetting resource.
type FirewallIpmacbindingSettingArgs struct {
	// Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
	Bindthroughfw pulumi.StringPtrInput
	// Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
	Bindtofw pulumi.StringPtrInput
	// Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.
	Undefinedhost pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallIpmacbindingSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallIpmacbindingSettingArgs)(nil)).Elem()
}

type FirewallIpmacbindingSettingInput interface {
	pulumi.Input

	ToFirewallIpmacbindingSettingOutput() FirewallIpmacbindingSettingOutput
	ToFirewallIpmacbindingSettingOutputWithContext(ctx context.Context) FirewallIpmacbindingSettingOutput
}

func (*FirewallIpmacbindingSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallIpmacbindingSetting)(nil)).Elem()
}

func (i *FirewallIpmacbindingSetting) ToFirewallIpmacbindingSettingOutput() FirewallIpmacbindingSettingOutput {
	return i.ToFirewallIpmacbindingSettingOutputWithContext(context.Background())
}

func (i *FirewallIpmacbindingSetting) ToFirewallIpmacbindingSettingOutputWithContext(ctx context.Context) FirewallIpmacbindingSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIpmacbindingSettingOutput)
}

// FirewallIpmacbindingSettingArrayInput is an input type that accepts FirewallIpmacbindingSettingArray and FirewallIpmacbindingSettingArrayOutput values.
// You can construct a concrete instance of `FirewallIpmacbindingSettingArrayInput` via:
//
//          FirewallIpmacbindingSettingArray{ FirewallIpmacbindingSettingArgs{...} }
type FirewallIpmacbindingSettingArrayInput interface {
	pulumi.Input

	ToFirewallIpmacbindingSettingArrayOutput() FirewallIpmacbindingSettingArrayOutput
	ToFirewallIpmacbindingSettingArrayOutputWithContext(context.Context) FirewallIpmacbindingSettingArrayOutput
}

type FirewallIpmacbindingSettingArray []FirewallIpmacbindingSettingInput

func (FirewallIpmacbindingSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallIpmacbindingSetting)(nil)).Elem()
}

func (i FirewallIpmacbindingSettingArray) ToFirewallIpmacbindingSettingArrayOutput() FirewallIpmacbindingSettingArrayOutput {
	return i.ToFirewallIpmacbindingSettingArrayOutputWithContext(context.Background())
}

func (i FirewallIpmacbindingSettingArray) ToFirewallIpmacbindingSettingArrayOutputWithContext(ctx context.Context) FirewallIpmacbindingSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIpmacbindingSettingArrayOutput)
}

// FirewallIpmacbindingSettingMapInput is an input type that accepts FirewallIpmacbindingSettingMap and FirewallIpmacbindingSettingMapOutput values.
// You can construct a concrete instance of `FirewallIpmacbindingSettingMapInput` via:
//
//          FirewallIpmacbindingSettingMap{ "key": FirewallIpmacbindingSettingArgs{...} }
type FirewallIpmacbindingSettingMapInput interface {
	pulumi.Input

	ToFirewallIpmacbindingSettingMapOutput() FirewallIpmacbindingSettingMapOutput
	ToFirewallIpmacbindingSettingMapOutputWithContext(context.Context) FirewallIpmacbindingSettingMapOutput
}

type FirewallIpmacbindingSettingMap map[string]FirewallIpmacbindingSettingInput

func (FirewallIpmacbindingSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallIpmacbindingSetting)(nil)).Elem()
}

func (i FirewallIpmacbindingSettingMap) ToFirewallIpmacbindingSettingMapOutput() FirewallIpmacbindingSettingMapOutput {
	return i.ToFirewallIpmacbindingSettingMapOutputWithContext(context.Background())
}

func (i FirewallIpmacbindingSettingMap) ToFirewallIpmacbindingSettingMapOutputWithContext(ctx context.Context) FirewallIpmacbindingSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallIpmacbindingSettingMapOutput)
}

type FirewallIpmacbindingSettingOutput struct{ *pulumi.OutputState }

func (FirewallIpmacbindingSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallIpmacbindingSetting)(nil)).Elem()
}

func (o FirewallIpmacbindingSettingOutput) ToFirewallIpmacbindingSettingOutput() FirewallIpmacbindingSettingOutput {
	return o
}

func (o FirewallIpmacbindingSettingOutput) ToFirewallIpmacbindingSettingOutputWithContext(ctx context.Context) FirewallIpmacbindingSettingOutput {
	return o
}

type FirewallIpmacbindingSettingArrayOutput struct{ *pulumi.OutputState }

func (FirewallIpmacbindingSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallIpmacbindingSetting)(nil)).Elem()
}

func (o FirewallIpmacbindingSettingArrayOutput) ToFirewallIpmacbindingSettingArrayOutput() FirewallIpmacbindingSettingArrayOutput {
	return o
}

func (o FirewallIpmacbindingSettingArrayOutput) ToFirewallIpmacbindingSettingArrayOutputWithContext(ctx context.Context) FirewallIpmacbindingSettingArrayOutput {
	return o
}

func (o FirewallIpmacbindingSettingArrayOutput) Index(i pulumi.IntInput) FirewallIpmacbindingSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallIpmacbindingSetting {
		return vs[0].([]*FirewallIpmacbindingSetting)[vs[1].(int)]
	}).(FirewallIpmacbindingSettingOutput)
}

type FirewallIpmacbindingSettingMapOutput struct{ *pulumi.OutputState }

func (FirewallIpmacbindingSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallIpmacbindingSetting)(nil)).Elem()
}

func (o FirewallIpmacbindingSettingMapOutput) ToFirewallIpmacbindingSettingMapOutput() FirewallIpmacbindingSettingMapOutput {
	return o
}

func (o FirewallIpmacbindingSettingMapOutput) ToFirewallIpmacbindingSettingMapOutputWithContext(ctx context.Context) FirewallIpmacbindingSettingMapOutput {
	return o
}

func (o FirewallIpmacbindingSettingMapOutput) MapIndex(k pulumi.StringInput) FirewallIpmacbindingSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallIpmacbindingSetting {
		return vs[0].(map[string]*FirewallIpmacbindingSetting)[vs[1].(string)]
	}).(FirewallIpmacbindingSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallIpmacbindingSettingInput)(nil)).Elem(), &FirewallIpmacbindingSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallIpmacbindingSettingArrayInput)(nil)).Elem(), FirewallIpmacbindingSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallIpmacbindingSettingMapInput)(nil)).Elem(), FirewallIpmacbindingSettingMap{})
	pulumi.RegisterOutputType(FirewallIpmacbindingSettingOutput{})
	pulumi.RegisterOutputType(FirewallIpmacbindingSettingArrayOutput{})
	pulumi.RegisterOutputType(FirewallIpmacbindingSettingMapOutput{})
}
