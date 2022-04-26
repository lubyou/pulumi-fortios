// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Show Internet Service botnet. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// Firewall InternetServiceBotnet can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/firewallInternetServiceBotnet:FirewallInternetServiceBotnet labelname {{fosid}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/firewallInternetServiceBotnet:FirewallInternetServiceBotnet labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallInternetServiceBotnet struct {
	pulumi.CustomResourceState

	// Internet Service Botnet ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Internet Service Botnet name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallInternetServiceBotnet registers a new resource with the given unique name, arguments, and options.
func NewFirewallInternetServiceBotnet(ctx *pulumi.Context,
	name string, args *FirewallInternetServiceBotnetArgs, opts ...pulumi.ResourceOption) (*FirewallInternetServiceBotnet, error) {
	if args == nil {
		args = &FirewallInternetServiceBotnetArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallInternetServiceBotnet
	err := ctx.RegisterResource("fortios:index/firewallInternetServiceBotnet:FirewallInternetServiceBotnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallInternetServiceBotnet gets an existing FirewallInternetServiceBotnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallInternetServiceBotnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallInternetServiceBotnetState, opts ...pulumi.ResourceOption) (*FirewallInternetServiceBotnet, error) {
	var resource FirewallInternetServiceBotnet
	err := ctx.ReadResource("fortios:index/firewallInternetServiceBotnet:FirewallInternetServiceBotnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallInternetServiceBotnet resources.
type firewallInternetServiceBotnetState struct {
	// Internet Service Botnet ID.
	Fosid *int `pulumi:"fosid"`
	// Internet Service Botnet name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallInternetServiceBotnetState struct {
	// Internet Service Botnet ID.
	Fosid pulumi.IntPtrInput
	// Internet Service Botnet name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetServiceBotnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceBotnetState)(nil)).Elem()
}

type firewallInternetServiceBotnetArgs struct {
	// Internet Service Botnet ID.
	Fosid *int `pulumi:"fosid"`
	// Internet Service Botnet name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallInternetServiceBotnet resource.
type FirewallInternetServiceBotnetArgs struct {
	// Internet Service Botnet ID.
	Fosid pulumi.IntPtrInput
	// Internet Service Botnet name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetServiceBotnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceBotnetArgs)(nil)).Elem()
}

type FirewallInternetServiceBotnetInput interface {
	pulumi.Input

	ToFirewallInternetServiceBotnetOutput() FirewallInternetServiceBotnetOutput
	ToFirewallInternetServiceBotnetOutputWithContext(ctx context.Context) FirewallInternetServiceBotnetOutput
}

func (*FirewallInternetServiceBotnet) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceBotnet)(nil)).Elem()
}

func (i *FirewallInternetServiceBotnet) ToFirewallInternetServiceBotnetOutput() FirewallInternetServiceBotnetOutput {
	return i.ToFirewallInternetServiceBotnetOutputWithContext(context.Background())
}

func (i *FirewallInternetServiceBotnet) ToFirewallInternetServiceBotnetOutputWithContext(ctx context.Context) FirewallInternetServiceBotnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceBotnetOutput)
}

// FirewallInternetServiceBotnetArrayInput is an input type that accepts FirewallInternetServiceBotnetArray and FirewallInternetServiceBotnetArrayOutput values.
// You can construct a concrete instance of `FirewallInternetServiceBotnetArrayInput` via:
//
//          FirewallInternetServiceBotnetArray{ FirewallInternetServiceBotnetArgs{...} }
type FirewallInternetServiceBotnetArrayInput interface {
	pulumi.Input

	ToFirewallInternetServiceBotnetArrayOutput() FirewallInternetServiceBotnetArrayOutput
	ToFirewallInternetServiceBotnetArrayOutputWithContext(context.Context) FirewallInternetServiceBotnetArrayOutput
}

type FirewallInternetServiceBotnetArray []FirewallInternetServiceBotnetInput

func (FirewallInternetServiceBotnetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceBotnet)(nil)).Elem()
}

func (i FirewallInternetServiceBotnetArray) ToFirewallInternetServiceBotnetArrayOutput() FirewallInternetServiceBotnetArrayOutput {
	return i.ToFirewallInternetServiceBotnetArrayOutputWithContext(context.Background())
}

func (i FirewallInternetServiceBotnetArray) ToFirewallInternetServiceBotnetArrayOutputWithContext(ctx context.Context) FirewallInternetServiceBotnetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceBotnetArrayOutput)
}

// FirewallInternetServiceBotnetMapInput is an input type that accepts FirewallInternetServiceBotnetMap and FirewallInternetServiceBotnetMapOutput values.
// You can construct a concrete instance of `FirewallInternetServiceBotnetMapInput` via:
//
//          FirewallInternetServiceBotnetMap{ "key": FirewallInternetServiceBotnetArgs{...} }
type FirewallInternetServiceBotnetMapInput interface {
	pulumi.Input

	ToFirewallInternetServiceBotnetMapOutput() FirewallInternetServiceBotnetMapOutput
	ToFirewallInternetServiceBotnetMapOutputWithContext(context.Context) FirewallInternetServiceBotnetMapOutput
}

type FirewallInternetServiceBotnetMap map[string]FirewallInternetServiceBotnetInput

func (FirewallInternetServiceBotnetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceBotnet)(nil)).Elem()
}

func (i FirewallInternetServiceBotnetMap) ToFirewallInternetServiceBotnetMapOutput() FirewallInternetServiceBotnetMapOutput {
	return i.ToFirewallInternetServiceBotnetMapOutputWithContext(context.Background())
}

func (i FirewallInternetServiceBotnetMap) ToFirewallInternetServiceBotnetMapOutputWithContext(ctx context.Context) FirewallInternetServiceBotnetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceBotnetMapOutput)
}

type FirewallInternetServiceBotnetOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceBotnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceBotnet)(nil)).Elem()
}

func (o FirewallInternetServiceBotnetOutput) ToFirewallInternetServiceBotnetOutput() FirewallInternetServiceBotnetOutput {
	return o
}

func (o FirewallInternetServiceBotnetOutput) ToFirewallInternetServiceBotnetOutputWithContext(ctx context.Context) FirewallInternetServiceBotnetOutput {
	return o
}

type FirewallInternetServiceBotnetArrayOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceBotnetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceBotnet)(nil)).Elem()
}

func (o FirewallInternetServiceBotnetArrayOutput) ToFirewallInternetServiceBotnetArrayOutput() FirewallInternetServiceBotnetArrayOutput {
	return o
}

func (o FirewallInternetServiceBotnetArrayOutput) ToFirewallInternetServiceBotnetArrayOutputWithContext(ctx context.Context) FirewallInternetServiceBotnetArrayOutput {
	return o
}

func (o FirewallInternetServiceBotnetArrayOutput) Index(i pulumi.IntInput) FirewallInternetServiceBotnetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallInternetServiceBotnet {
		return vs[0].([]*FirewallInternetServiceBotnet)[vs[1].(int)]
	}).(FirewallInternetServiceBotnetOutput)
}

type FirewallInternetServiceBotnetMapOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceBotnetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceBotnet)(nil)).Elem()
}

func (o FirewallInternetServiceBotnetMapOutput) ToFirewallInternetServiceBotnetMapOutput() FirewallInternetServiceBotnetMapOutput {
	return o
}

func (o FirewallInternetServiceBotnetMapOutput) ToFirewallInternetServiceBotnetMapOutputWithContext(ctx context.Context) FirewallInternetServiceBotnetMapOutput {
	return o
}

func (o FirewallInternetServiceBotnetMapOutput) MapIndex(k pulumi.StringInput) FirewallInternetServiceBotnetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallInternetServiceBotnet {
		return vs[0].(map[string]*FirewallInternetServiceBotnet)[vs[1].(string)]
	}).(FirewallInternetServiceBotnetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceBotnetInput)(nil)).Elem(), &FirewallInternetServiceBotnet{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceBotnetArrayInput)(nil)).Elem(), FirewallInternetServiceBotnetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceBotnetMapInput)(nil)).Elem(), FirewallInternetServiceBotnetMap{})
	pulumi.RegisterOutputType(FirewallInternetServiceBotnetOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceBotnetArrayOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceBotnetMapOutput{})
}
