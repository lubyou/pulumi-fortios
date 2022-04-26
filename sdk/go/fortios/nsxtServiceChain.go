// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure NSX-T service chain. Applies to FortiOS Version `>= 7.0.0`.
//
// ## Import
//
// Nsxt ServiceChain can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/nsxtServiceChain:NsxtServiceChain labelname {{fosid}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/nsxtServiceChain:NsxtServiceChain labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type NsxtServiceChain struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Chain ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Index name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configure service index. The structure of `serviceIndex` block is documented below.
	ServiceIndices NsxtServiceChainServiceIndexArrayOutput `pulumi:"serviceIndices"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewNsxtServiceChain registers a new resource with the given unique name, arguments, and options.
func NewNsxtServiceChain(ctx *pulumi.Context,
	name string, args *NsxtServiceChainArgs, opts ...pulumi.ResourceOption) (*NsxtServiceChain, error) {
	if args == nil {
		args = &NsxtServiceChainArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource NsxtServiceChain
	err := ctx.RegisterResource("fortios:index/nsxtServiceChain:NsxtServiceChain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsxtServiceChain gets an existing NsxtServiceChain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsxtServiceChain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsxtServiceChainState, opts ...pulumi.ResourceOption) (*NsxtServiceChain, error) {
	var resource NsxtServiceChain
	err := ctx.ReadResource("fortios:index/nsxtServiceChain:NsxtServiceChain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NsxtServiceChain resources.
type nsxtServiceChainState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Chain ID.
	Fosid *int `pulumi:"fosid"`
	// Index name.
	Name *string `pulumi:"name"`
	// Configure service index. The structure of `serviceIndex` block is documented below.
	ServiceIndices []NsxtServiceChainServiceIndex `pulumi:"serviceIndices"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type NsxtServiceChainState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Chain ID.
	Fosid pulumi.IntPtrInput
	// Index name.
	Name pulumi.StringPtrInput
	// Configure service index. The structure of `serviceIndex` block is documented below.
	ServiceIndices NsxtServiceChainServiceIndexArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NsxtServiceChainState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtServiceChainState)(nil)).Elem()
}

type nsxtServiceChainArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Chain ID.
	Fosid *int `pulumi:"fosid"`
	// Index name.
	Name *string `pulumi:"name"`
	// Configure service index. The structure of `serviceIndex` block is documented below.
	ServiceIndices []NsxtServiceChainServiceIndex `pulumi:"serviceIndices"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a NsxtServiceChain resource.
type NsxtServiceChainArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Chain ID.
	Fosid pulumi.IntPtrInput
	// Index name.
	Name pulumi.StringPtrInput
	// Configure service index. The structure of `serviceIndex` block is documented below.
	ServiceIndices NsxtServiceChainServiceIndexArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (NsxtServiceChainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtServiceChainArgs)(nil)).Elem()
}

type NsxtServiceChainInput interface {
	pulumi.Input

	ToNsxtServiceChainOutput() NsxtServiceChainOutput
	ToNsxtServiceChainOutputWithContext(ctx context.Context) NsxtServiceChainOutput
}

func (*NsxtServiceChain) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtServiceChain)(nil)).Elem()
}

func (i *NsxtServiceChain) ToNsxtServiceChainOutput() NsxtServiceChainOutput {
	return i.ToNsxtServiceChainOutputWithContext(context.Background())
}

func (i *NsxtServiceChain) ToNsxtServiceChainOutputWithContext(ctx context.Context) NsxtServiceChainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtServiceChainOutput)
}

// NsxtServiceChainArrayInput is an input type that accepts NsxtServiceChainArray and NsxtServiceChainArrayOutput values.
// You can construct a concrete instance of `NsxtServiceChainArrayInput` via:
//
//          NsxtServiceChainArray{ NsxtServiceChainArgs{...} }
type NsxtServiceChainArrayInput interface {
	pulumi.Input

	ToNsxtServiceChainArrayOutput() NsxtServiceChainArrayOutput
	ToNsxtServiceChainArrayOutputWithContext(context.Context) NsxtServiceChainArrayOutput
}

type NsxtServiceChainArray []NsxtServiceChainInput

func (NsxtServiceChainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtServiceChain)(nil)).Elem()
}

func (i NsxtServiceChainArray) ToNsxtServiceChainArrayOutput() NsxtServiceChainArrayOutput {
	return i.ToNsxtServiceChainArrayOutputWithContext(context.Background())
}

func (i NsxtServiceChainArray) ToNsxtServiceChainArrayOutputWithContext(ctx context.Context) NsxtServiceChainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtServiceChainArrayOutput)
}

// NsxtServiceChainMapInput is an input type that accepts NsxtServiceChainMap and NsxtServiceChainMapOutput values.
// You can construct a concrete instance of `NsxtServiceChainMapInput` via:
//
//          NsxtServiceChainMap{ "key": NsxtServiceChainArgs{...} }
type NsxtServiceChainMapInput interface {
	pulumi.Input

	ToNsxtServiceChainMapOutput() NsxtServiceChainMapOutput
	ToNsxtServiceChainMapOutputWithContext(context.Context) NsxtServiceChainMapOutput
}

type NsxtServiceChainMap map[string]NsxtServiceChainInput

func (NsxtServiceChainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtServiceChain)(nil)).Elem()
}

func (i NsxtServiceChainMap) ToNsxtServiceChainMapOutput() NsxtServiceChainMapOutput {
	return i.ToNsxtServiceChainMapOutputWithContext(context.Background())
}

func (i NsxtServiceChainMap) ToNsxtServiceChainMapOutputWithContext(ctx context.Context) NsxtServiceChainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtServiceChainMapOutput)
}

type NsxtServiceChainOutput struct{ *pulumi.OutputState }

func (NsxtServiceChainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtServiceChain)(nil)).Elem()
}

func (o NsxtServiceChainOutput) ToNsxtServiceChainOutput() NsxtServiceChainOutput {
	return o
}

func (o NsxtServiceChainOutput) ToNsxtServiceChainOutputWithContext(ctx context.Context) NsxtServiceChainOutput {
	return o
}

type NsxtServiceChainArrayOutput struct{ *pulumi.OutputState }

func (NsxtServiceChainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtServiceChain)(nil)).Elem()
}

func (o NsxtServiceChainArrayOutput) ToNsxtServiceChainArrayOutput() NsxtServiceChainArrayOutput {
	return o
}

func (o NsxtServiceChainArrayOutput) ToNsxtServiceChainArrayOutputWithContext(ctx context.Context) NsxtServiceChainArrayOutput {
	return o
}

func (o NsxtServiceChainArrayOutput) Index(i pulumi.IntInput) NsxtServiceChainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NsxtServiceChain {
		return vs[0].([]*NsxtServiceChain)[vs[1].(int)]
	}).(NsxtServiceChainOutput)
}

type NsxtServiceChainMapOutput struct{ *pulumi.OutputState }

func (NsxtServiceChainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtServiceChain)(nil)).Elem()
}

func (o NsxtServiceChainMapOutput) ToNsxtServiceChainMapOutput() NsxtServiceChainMapOutput {
	return o
}

func (o NsxtServiceChainMapOutput) ToNsxtServiceChainMapOutputWithContext(ctx context.Context) NsxtServiceChainMapOutput {
	return o
}

func (o NsxtServiceChainMapOutput) MapIndex(k pulumi.StringInput) NsxtServiceChainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NsxtServiceChain {
		return vs[0].(map[string]*NsxtServiceChain)[vs[1].(string)]
	}).(NsxtServiceChainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtServiceChainInput)(nil)).Elem(), &NsxtServiceChain{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtServiceChainArrayInput)(nil)).Elem(), NsxtServiceChainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtServiceChainMapInput)(nil)).Elem(), NsxtServiceChainMap{})
	pulumi.RegisterOutputType(NsxtServiceChainOutput{})
	pulumi.RegisterOutputType(NsxtServiceChainArrayOutput{})
	pulumi.RegisterOutputType(NsxtServiceChainMapOutput{})
}
