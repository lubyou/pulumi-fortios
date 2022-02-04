// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure global DPDK options. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// Dpdk Global can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/dpdkGlobal:DpdkGlobal labelname DpdkGlobal
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type DpdkGlobal struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
	Elasticbuffer pulumi.StringOutput `pulumi:"elasticbuffer"`
	// Percentage of main memory allocated to hugepages, which are available for DPDK operation.
	HugepagePercentage pulumi.IntOutput `pulumi:"hugepagePercentage"`
	// Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
	Interfaces DpdkGlobalInterfaceArrayOutput `pulumi:"interfaces"`
	// Percentage of main memory allocated to DPDK packet buffer.
	MbufpoolPercentage pulumi.IntOutput `pulumi:"mbufpoolPercentage"`
	// Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
	Multiqueue pulumi.StringOutput `pulumi:"multiqueue"`
	// Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
	PerSessionAccounting pulumi.StringOutput `pulumi:"perSessionAccounting"`
	// Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
	SleepOnIdle pulumi.StringOutput `pulumi:"sleepOnIdle"`
	// Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDpdkGlobal registers a new resource with the given unique name, arguments, and options.
func NewDpdkGlobal(ctx *pulumi.Context,
	name string, args *DpdkGlobalArgs, opts ...pulumi.ResourceOption) (*DpdkGlobal, error) {
	if args == nil {
		args = &DpdkGlobalArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource DpdkGlobal
	err := ctx.RegisterResource("fortios:index/dpdkGlobal:DpdkGlobal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDpdkGlobal gets an existing DpdkGlobal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDpdkGlobal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DpdkGlobalState, opts ...pulumi.ResourceOption) (*DpdkGlobal, error) {
	var resource DpdkGlobal
	err := ctx.ReadResource("fortios:index/dpdkGlobal:DpdkGlobal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DpdkGlobal resources.
type dpdkGlobalState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
	Elasticbuffer *string `pulumi:"elasticbuffer"`
	// Percentage of main memory allocated to hugepages, which are available for DPDK operation.
	HugepagePercentage *int `pulumi:"hugepagePercentage"`
	// Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
	Interfaces []DpdkGlobalInterface `pulumi:"interfaces"`
	// Percentage of main memory allocated to DPDK packet buffer.
	MbufpoolPercentage *int `pulumi:"mbufpoolPercentage"`
	// Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
	Multiqueue *string `pulumi:"multiqueue"`
	// Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
	PerSessionAccounting *string `pulumi:"perSessionAccounting"`
	// Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
	SleepOnIdle *string `pulumi:"sleepOnIdle"`
	// Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DpdkGlobalState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
	Elasticbuffer pulumi.StringPtrInput
	// Percentage of main memory allocated to hugepages, which are available for DPDK operation.
	HugepagePercentage pulumi.IntPtrInput
	// Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
	Interfaces DpdkGlobalInterfaceArrayInput
	// Percentage of main memory allocated to DPDK packet buffer.
	MbufpoolPercentage pulumi.IntPtrInput
	// Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
	Multiqueue pulumi.StringPtrInput
	// Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
	PerSessionAccounting pulumi.StringPtrInput
	// Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
	SleepOnIdle pulumi.StringPtrInput
	// Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DpdkGlobalState) ElementType() reflect.Type {
	return reflect.TypeOf((*dpdkGlobalState)(nil)).Elem()
}

type dpdkGlobalArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
	Elasticbuffer *string `pulumi:"elasticbuffer"`
	// Percentage of main memory allocated to hugepages, which are available for DPDK operation.
	HugepagePercentage *int `pulumi:"hugepagePercentage"`
	// Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
	Interfaces []DpdkGlobalInterface `pulumi:"interfaces"`
	// Percentage of main memory allocated to DPDK packet buffer.
	MbufpoolPercentage *int `pulumi:"mbufpoolPercentage"`
	// Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
	Multiqueue *string `pulumi:"multiqueue"`
	// Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
	PerSessionAccounting *string `pulumi:"perSessionAccounting"`
	// Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
	SleepOnIdle *string `pulumi:"sleepOnIdle"`
	// Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a DpdkGlobal resource.
type DpdkGlobalArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
	Elasticbuffer pulumi.StringPtrInput
	// Percentage of main memory allocated to hugepages, which are available for DPDK operation.
	HugepagePercentage pulumi.IntPtrInput
	// Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
	Interfaces DpdkGlobalInterfaceArrayInput
	// Percentage of main memory allocated to DPDK packet buffer.
	MbufpoolPercentage pulumi.IntPtrInput
	// Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
	Multiqueue pulumi.StringPtrInput
	// Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
	PerSessionAccounting pulumi.StringPtrInput
	// Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
	SleepOnIdle pulumi.StringPtrInput
	// Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DpdkGlobalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dpdkGlobalArgs)(nil)).Elem()
}

type DpdkGlobalInput interface {
	pulumi.Input

	ToDpdkGlobalOutput() DpdkGlobalOutput
	ToDpdkGlobalOutputWithContext(ctx context.Context) DpdkGlobalOutput
}

func (*DpdkGlobal) ElementType() reflect.Type {
	return reflect.TypeOf((**DpdkGlobal)(nil)).Elem()
}

func (i *DpdkGlobal) ToDpdkGlobalOutput() DpdkGlobalOutput {
	return i.ToDpdkGlobalOutputWithContext(context.Background())
}

func (i *DpdkGlobal) ToDpdkGlobalOutputWithContext(ctx context.Context) DpdkGlobalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DpdkGlobalOutput)
}

// DpdkGlobalArrayInput is an input type that accepts DpdkGlobalArray and DpdkGlobalArrayOutput values.
// You can construct a concrete instance of `DpdkGlobalArrayInput` via:
//
//          DpdkGlobalArray{ DpdkGlobalArgs{...} }
type DpdkGlobalArrayInput interface {
	pulumi.Input

	ToDpdkGlobalArrayOutput() DpdkGlobalArrayOutput
	ToDpdkGlobalArrayOutputWithContext(context.Context) DpdkGlobalArrayOutput
}

type DpdkGlobalArray []DpdkGlobalInput

func (DpdkGlobalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DpdkGlobal)(nil)).Elem()
}

func (i DpdkGlobalArray) ToDpdkGlobalArrayOutput() DpdkGlobalArrayOutput {
	return i.ToDpdkGlobalArrayOutputWithContext(context.Background())
}

func (i DpdkGlobalArray) ToDpdkGlobalArrayOutputWithContext(ctx context.Context) DpdkGlobalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DpdkGlobalArrayOutput)
}

// DpdkGlobalMapInput is an input type that accepts DpdkGlobalMap and DpdkGlobalMapOutput values.
// You can construct a concrete instance of `DpdkGlobalMapInput` via:
//
//          DpdkGlobalMap{ "key": DpdkGlobalArgs{...} }
type DpdkGlobalMapInput interface {
	pulumi.Input

	ToDpdkGlobalMapOutput() DpdkGlobalMapOutput
	ToDpdkGlobalMapOutputWithContext(context.Context) DpdkGlobalMapOutput
}

type DpdkGlobalMap map[string]DpdkGlobalInput

func (DpdkGlobalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DpdkGlobal)(nil)).Elem()
}

func (i DpdkGlobalMap) ToDpdkGlobalMapOutput() DpdkGlobalMapOutput {
	return i.ToDpdkGlobalMapOutputWithContext(context.Background())
}

func (i DpdkGlobalMap) ToDpdkGlobalMapOutputWithContext(ctx context.Context) DpdkGlobalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DpdkGlobalMapOutput)
}

type DpdkGlobalOutput struct{ *pulumi.OutputState }

func (DpdkGlobalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DpdkGlobal)(nil)).Elem()
}

func (o DpdkGlobalOutput) ToDpdkGlobalOutput() DpdkGlobalOutput {
	return o
}

func (o DpdkGlobalOutput) ToDpdkGlobalOutputWithContext(ctx context.Context) DpdkGlobalOutput {
	return o
}

type DpdkGlobalArrayOutput struct{ *pulumi.OutputState }

func (DpdkGlobalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DpdkGlobal)(nil)).Elem()
}

func (o DpdkGlobalArrayOutput) ToDpdkGlobalArrayOutput() DpdkGlobalArrayOutput {
	return o
}

func (o DpdkGlobalArrayOutput) ToDpdkGlobalArrayOutputWithContext(ctx context.Context) DpdkGlobalArrayOutput {
	return o
}

func (o DpdkGlobalArrayOutput) Index(i pulumi.IntInput) DpdkGlobalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DpdkGlobal {
		return vs[0].([]*DpdkGlobal)[vs[1].(int)]
	}).(DpdkGlobalOutput)
}

type DpdkGlobalMapOutput struct{ *pulumi.OutputState }

func (DpdkGlobalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DpdkGlobal)(nil)).Elem()
}

func (o DpdkGlobalMapOutput) ToDpdkGlobalMapOutput() DpdkGlobalMapOutput {
	return o
}

func (o DpdkGlobalMapOutput) ToDpdkGlobalMapOutputWithContext(ctx context.Context) DpdkGlobalMapOutput {
	return o
}

func (o DpdkGlobalMapOutput) MapIndex(k pulumi.StringInput) DpdkGlobalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DpdkGlobal {
		return vs[0].(map[string]*DpdkGlobal)[vs[1].(string)]
	}).(DpdkGlobalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DpdkGlobalInput)(nil)).Elem(), &DpdkGlobal{})
	pulumi.RegisterInputType(reflect.TypeOf((*DpdkGlobalArrayInput)(nil)).Elem(), DpdkGlobalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DpdkGlobalMapInput)(nil)).Elem(), DpdkGlobalMap{})
	pulumi.RegisterOutputType(DpdkGlobalOutput{})
	pulumi.RegisterOutputType(DpdkGlobalArrayOutput{})
	pulumi.RegisterOutputType(DpdkGlobalMapOutput{})
}
