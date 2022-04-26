// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPS URL filter IPv6 DNS servers.
//
// ## Import
//
// System IpsUrlfilterDns6 can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemIpsUrlfilterDns6:SystemIpsUrlfilterDns6 labelname {{address6}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemIpsUrlfilterDns6:SystemIpsUrlfilterDns6 labelname {{address6}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemIpsUrlfilterDns6 struct {
	pulumi.CustomResourceState

	// IPv6 address of DNS server.
	Address6 pulumi.StringOutput `pulumi:"address6"`
	// Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemIpsUrlfilterDns6 registers a new resource with the given unique name, arguments, and options.
func NewSystemIpsUrlfilterDns6(ctx *pulumi.Context,
	name string, args *SystemIpsUrlfilterDns6Args, opts ...pulumi.ResourceOption) (*SystemIpsUrlfilterDns6, error) {
	if args == nil {
		args = &SystemIpsUrlfilterDns6Args{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemIpsUrlfilterDns6
	err := ctx.RegisterResource("fortios:index/systemIpsUrlfilterDns6:SystemIpsUrlfilterDns6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemIpsUrlfilterDns6 gets an existing SystemIpsUrlfilterDns6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemIpsUrlfilterDns6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemIpsUrlfilterDns6State, opts ...pulumi.ResourceOption) (*SystemIpsUrlfilterDns6, error) {
	var resource SystemIpsUrlfilterDns6
	err := ctx.ReadResource("fortios:index/systemIpsUrlfilterDns6:SystemIpsUrlfilterDns6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemIpsUrlfilterDns6 resources.
type systemIpsUrlfilterDns6State struct {
	// IPv6 address of DNS server.
	Address6 *string `pulumi:"address6"`
	// Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemIpsUrlfilterDns6State struct {
	// IPv6 address of DNS server.
	Address6 pulumi.StringPtrInput
	// Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemIpsUrlfilterDns6State) ElementType() reflect.Type {
	return reflect.TypeOf((*systemIpsUrlfilterDns6State)(nil)).Elem()
}

type systemIpsUrlfilterDns6Args struct {
	// IPv6 address of DNS server.
	Address6 *string `pulumi:"address6"`
	// Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemIpsUrlfilterDns6 resource.
type SystemIpsUrlfilterDns6Args struct {
	// IPv6 address of DNS server.
	Address6 pulumi.StringPtrInput
	// Enable/disable this server for IPv6 DNS queries. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemIpsUrlfilterDns6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*systemIpsUrlfilterDns6Args)(nil)).Elem()
}

type SystemIpsUrlfilterDns6Input interface {
	pulumi.Input

	ToSystemIpsUrlfilterDns6Output() SystemIpsUrlfilterDns6Output
	ToSystemIpsUrlfilterDns6OutputWithContext(ctx context.Context) SystemIpsUrlfilterDns6Output
}

func (*SystemIpsUrlfilterDns6) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemIpsUrlfilterDns6)(nil)).Elem()
}

func (i *SystemIpsUrlfilterDns6) ToSystemIpsUrlfilterDns6Output() SystemIpsUrlfilterDns6Output {
	return i.ToSystemIpsUrlfilterDns6OutputWithContext(context.Background())
}

func (i *SystemIpsUrlfilterDns6) ToSystemIpsUrlfilterDns6OutputWithContext(ctx context.Context) SystemIpsUrlfilterDns6Output {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIpsUrlfilterDns6Output)
}

// SystemIpsUrlfilterDns6ArrayInput is an input type that accepts SystemIpsUrlfilterDns6Array and SystemIpsUrlfilterDns6ArrayOutput values.
// You can construct a concrete instance of `SystemIpsUrlfilterDns6ArrayInput` via:
//
//          SystemIpsUrlfilterDns6Array{ SystemIpsUrlfilterDns6Args{...} }
type SystemIpsUrlfilterDns6ArrayInput interface {
	pulumi.Input

	ToSystemIpsUrlfilterDns6ArrayOutput() SystemIpsUrlfilterDns6ArrayOutput
	ToSystemIpsUrlfilterDns6ArrayOutputWithContext(context.Context) SystemIpsUrlfilterDns6ArrayOutput
}

type SystemIpsUrlfilterDns6Array []SystemIpsUrlfilterDns6Input

func (SystemIpsUrlfilterDns6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemIpsUrlfilterDns6)(nil)).Elem()
}

func (i SystemIpsUrlfilterDns6Array) ToSystemIpsUrlfilterDns6ArrayOutput() SystemIpsUrlfilterDns6ArrayOutput {
	return i.ToSystemIpsUrlfilterDns6ArrayOutputWithContext(context.Background())
}

func (i SystemIpsUrlfilterDns6Array) ToSystemIpsUrlfilterDns6ArrayOutputWithContext(ctx context.Context) SystemIpsUrlfilterDns6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIpsUrlfilterDns6ArrayOutput)
}

// SystemIpsUrlfilterDns6MapInput is an input type that accepts SystemIpsUrlfilterDns6Map and SystemIpsUrlfilterDns6MapOutput values.
// You can construct a concrete instance of `SystemIpsUrlfilterDns6MapInput` via:
//
//          SystemIpsUrlfilterDns6Map{ "key": SystemIpsUrlfilterDns6Args{...} }
type SystemIpsUrlfilterDns6MapInput interface {
	pulumi.Input

	ToSystemIpsUrlfilterDns6MapOutput() SystemIpsUrlfilterDns6MapOutput
	ToSystemIpsUrlfilterDns6MapOutputWithContext(context.Context) SystemIpsUrlfilterDns6MapOutput
}

type SystemIpsUrlfilterDns6Map map[string]SystemIpsUrlfilterDns6Input

func (SystemIpsUrlfilterDns6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemIpsUrlfilterDns6)(nil)).Elem()
}

func (i SystemIpsUrlfilterDns6Map) ToSystemIpsUrlfilterDns6MapOutput() SystemIpsUrlfilterDns6MapOutput {
	return i.ToSystemIpsUrlfilterDns6MapOutputWithContext(context.Background())
}

func (i SystemIpsUrlfilterDns6Map) ToSystemIpsUrlfilterDns6MapOutputWithContext(ctx context.Context) SystemIpsUrlfilterDns6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIpsUrlfilterDns6MapOutput)
}

type SystemIpsUrlfilterDns6Output struct{ *pulumi.OutputState }

func (SystemIpsUrlfilterDns6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemIpsUrlfilterDns6)(nil)).Elem()
}

func (o SystemIpsUrlfilterDns6Output) ToSystemIpsUrlfilterDns6Output() SystemIpsUrlfilterDns6Output {
	return o
}

func (o SystemIpsUrlfilterDns6Output) ToSystemIpsUrlfilterDns6OutputWithContext(ctx context.Context) SystemIpsUrlfilterDns6Output {
	return o
}

type SystemIpsUrlfilterDns6ArrayOutput struct{ *pulumi.OutputState }

func (SystemIpsUrlfilterDns6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemIpsUrlfilterDns6)(nil)).Elem()
}

func (o SystemIpsUrlfilterDns6ArrayOutput) ToSystemIpsUrlfilterDns6ArrayOutput() SystemIpsUrlfilterDns6ArrayOutput {
	return o
}

func (o SystemIpsUrlfilterDns6ArrayOutput) ToSystemIpsUrlfilterDns6ArrayOutputWithContext(ctx context.Context) SystemIpsUrlfilterDns6ArrayOutput {
	return o
}

func (o SystemIpsUrlfilterDns6ArrayOutput) Index(i pulumi.IntInput) SystemIpsUrlfilterDns6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemIpsUrlfilterDns6 {
		return vs[0].([]*SystemIpsUrlfilterDns6)[vs[1].(int)]
	}).(SystemIpsUrlfilterDns6Output)
}

type SystemIpsUrlfilterDns6MapOutput struct{ *pulumi.OutputState }

func (SystemIpsUrlfilterDns6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemIpsUrlfilterDns6)(nil)).Elem()
}

func (o SystemIpsUrlfilterDns6MapOutput) ToSystemIpsUrlfilterDns6MapOutput() SystemIpsUrlfilterDns6MapOutput {
	return o
}

func (o SystemIpsUrlfilterDns6MapOutput) ToSystemIpsUrlfilterDns6MapOutputWithContext(ctx context.Context) SystemIpsUrlfilterDns6MapOutput {
	return o
}

func (o SystemIpsUrlfilterDns6MapOutput) MapIndex(k pulumi.StringInput) SystemIpsUrlfilterDns6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemIpsUrlfilterDns6 {
		return vs[0].(map[string]*SystemIpsUrlfilterDns6)[vs[1].(string)]
	}).(SystemIpsUrlfilterDns6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIpsUrlfilterDns6Input)(nil)).Elem(), &SystemIpsUrlfilterDns6{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIpsUrlfilterDns6ArrayInput)(nil)).Elem(), SystemIpsUrlfilterDns6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIpsUrlfilterDns6MapInput)(nil)).Elem(), SystemIpsUrlfilterDns6Map{})
	pulumi.RegisterOutputType(SystemIpsUrlfilterDns6Output{})
	pulumi.RegisterOutputType(SystemIpsUrlfilterDns6ArrayOutput{})
	pulumi.RegisterOutputType(SystemIpsUrlfilterDns6MapOutput{})
}
