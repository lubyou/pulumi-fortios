// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure CPUs enabled to run engines in each DPDK stage.
//
// ## Import
//
// Dpdk Cpus can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/dpdkCpus:DpdkCpus labelname DpdkCpus
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type DpdkCpus struct {
	pulumi.CustomResourceState

	// CPUs enabled to run DPDK IPS engines.
	IpsCpus pulumi.StringOutput `pulumi:"ipsCpus"`
	// CPUs enabled to run DPDK RX engines.
	RxCpus pulumi.StringOutput `pulumi:"rxCpus"`
	// CPUs enabled to run DPDK TX engines.
	TxCpus pulumi.StringOutput `pulumi:"txCpus"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// CPUs enabled to run DPDK VNP engines.
	VnpCpus pulumi.StringOutput `pulumi:"vnpCpus"`
}

// NewDpdkCpus registers a new resource with the given unique name, arguments, and options.
func NewDpdkCpus(ctx *pulumi.Context,
	name string, args *DpdkCpusArgs, opts ...pulumi.ResourceOption) (*DpdkCpus, error) {
	if args == nil {
		args = &DpdkCpusArgs{}
	}

	var resource DpdkCpus
	err := ctx.RegisterResource("fortios:index/dpdkCpus:DpdkCpus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDpdkCpus gets an existing DpdkCpus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDpdkCpus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DpdkCpusState, opts ...pulumi.ResourceOption) (*DpdkCpus, error) {
	var resource DpdkCpus
	err := ctx.ReadResource("fortios:index/dpdkCpus:DpdkCpus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DpdkCpus resources.
type dpdkCpusState struct {
	// CPUs enabled to run DPDK IPS engines.
	IpsCpus *string `pulumi:"ipsCpus"`
	// CPUs enabled to run DPDK RX engines.
	RxCpus *string `pulumi:"rxCpus"`
	// CPUs enabled to run DPDK TX engines.
	TxCpus *string `pulumi:"txCpus"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// CPUs enabled to run DPDK VNP engines.
	VnpCpus *string `pulumi:"vnpCpus"`
}

type DpdkCpusState struct {
	// CPUs enabled to run DPDK IPS engines.
	IpsCpus pulumi.StringPtrInput
	// CPUs enabled to run DPDK RX engines.
	RxCpus pulumi.StringPtrInput
	// CPUs enabled to run DPDK TX engines.
	TxCpus pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// CPUs enabled to run DPDK VNP engines.
	VnpCpus pulumi.StringPtrInput
}

func (DpdkCpusState) ElementType() reflect.Type {
	return reflect.TypeOf((*dpdkCpusState)(nil)).Elem()
}

type dpdkCpusArgs struct {
	// CPUs enabled to run DPDK IPS engines.
	IpsCpus *string `pulumi:"ipsCpus"`
	// CPUs enabled to run DPDK RX engines.
	RxCpus *string `pulumi:"rxCpus"`
	// CPUs enabled to run DPDK TX engines.
	TxCpus *string `pulumi:"txCpus"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// CPUs enabled to run DPDK VNP engines.
	VnpCpus *string `pulumi:"vnpCpus"`
}

// The set of arguments for constructing a DpdkCpus resource.
type DpdkCpusArgs struct {
	// CPUs enabled to run DPDK IPS engines.
	IpsCpus pulumi.StringPtrInput
	// CPUs enabled to run DPDK RX engines.
	RxCpus pulumi.StringPtrInput
	// CPUs enabled to run DPDK TX engines.
	TxCpus pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// CPUs enabled to run DPDK VNP engines.
	VnpCpus pulumi.StringPtrInput
}

func (DpdkCpusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dpdkCpusArgs)(nil)).Elem()
}

type DpdkCpusInput interface {
	pulumi.Input

	ToDpdkCpusOutput() DpdkCpusOutput
	ToDpdkCpusOutputWithContext(ctx context.Context) DpdkCpusOutput
}

func (*DpdkCpus) ElementType() reflect.Type {
	return reflect.TypeOf((*DpdkCpus)(nil))
}

func (i *DpdkCpus) ToDpdkCpusOutput() DpdkCpusOutput {
	return i.ToDpdkCpusOutputWithContext(context.Background())
}

func (i *DpdkCpus) ToDpdkCpusOutputWithContext(ctx context.Context) DpdkCpusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DpdkCpusOutput)
}

func (i *DpdkCpus) ToDpdkCpusPtrOutput() DpdkCpusPtrOutput {
	return i.ToDpdkCpusPtrOutputWithContext(context.Background())
}

func (i *DpdkCpus) ToDpdkCpusPtrOutputWithContext(ctx context.Context) DpdkCpusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DpdkCpusPtrOutput)
}

type DpdkCpusPtrInput interface {
	pulumi.Input

	ToDpdkCpusPtrOutput() DpdkCpusPtrOutput
	ToDpdkCpusPtrOutputWithContext(ctx context.Context) DpdkCpusPtrOutput
}

type dpdkCpusPtrType DpdkCpusArgs

func (*dpdkCpusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DpdkCpus)(nil))
}

func (i *dpdkCpusPtrType) ToDpdkCpusPtrOutput() DpdkCpusPtrOutput {
	return i.ToDpdkCpusPtrOutputWithContext(context.Background())
}

func (i *dpdkCpusPtrType) ToDpdkCpusPtrOutputWithContext(ctx context.Context) DpdkCpusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DpdkCpusPtrOutput)
}

// DpdkCpusArrayInput is an input type that accepts DpdkCpusArray and DpdkCpusArrayOutput values.
// You can construct a concrete instance of `DpdkCpusArrayInput` via:
//
//          DpdkCpusArray{ DpdkCpusArgs{...} }
type DpdkCpusArrayInput interface {
	pulumi.Input

	ToDpdkCpusArrayOutput() DpdkCpusArrayOutput
	ToDpdkCpusArrayOutputWithContext(context.Context) DpdkCpusArrayOutput
}

type DpdkCpusArray []DpdkCpusInput

func (DpdkCpusArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DpdkCpus)(nil))
}

func (i DpdkCpusArray) ToDpdkCpusArrayOutput() DpdkCpusArrayOutput {
	return i.ToDpdkCpusArrayOutputWithContext(context.Background())
}

func (i DpdkCpusArray) ToDpdkCpusArrayOutputWithContext(ctx context.Context) DpdkCpusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DpdkCpusArrayOutput)
}

// DpdkCpusMapInput is an input type that accepts DpdkCpusMap and DpdkCpusMapOutput values.
// You can construct a concrete instance of `DpdkCpusMapInput` via:
//
//          DpdkCpusMap{ "key": DpdkCpusArgs{...} }
type DpdkCpusMapInput interface {
	pulumi.Input

	ToDpdkCpusMapOutput() DpdkCpusMapOutput
	ToDpdkCpusMapOutputWithContext(context.Context) DpdkCpusMapOutput
}

type DpdkCpusMap map[string]DpdkCpusInput

func (DpdkCpusMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DpdkCpus)(nil))
}

func (i DpdkCpusMap) ToDpdkCpusMapOutput() DpdkCpusMapOutput {
	return i.ToDpdkCpusMapOutputWithContext(context.Background())
}

func (i DpdkCpusMap) ToDpdkCpusMapOutputWithContext(ctx context.Context) DpdkCpusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DpdkCpusMapOutput)
}

type DpdkCpusOutput struct {
	*pulumi.OutputState
}

func (DpdkCpusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DpdkCpus)(nil))
}

func (o DpdkCpusOutput) ToDpdkCpusOutput() DpdkCpusOutput {
	return o
}

func (o DpdkCpusOutput) ToDpdkCpusOutputWithContext(ctx context.Context) DpdkCpusOutput {
	return o
}

func (o DpdkCpusOutput) ToDpdkCpusPtrOutput() DpdkCpusPtrOutput {
	return o.ToDpdkCpusPtrOutputWithContext(context.Background())
}

func (o DpdkCpusOutput) ToDpdkCpusPtrOutputWithContext(ctx context.Context) DpdkCpusPtrOutput {
	return o.ApplyT(func(v DpdkCpus) *DpdkCpus {
		return &v
	}).(DpdkCpusPtrOutput)
}

type DpdkCpusPtrOutput struct {
	*pulumi.OutputState
}

func (DpdkCpusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DpdkCpus)(nil))
}

func (o DpdkCpusPtrOutput) ToDpdkCpusPtrOutput() DpdkCpusPtrOutput {
	return o
}

func (o DpdkCpusPtrOutput) ToDpdkCpusPtrOutputWithContext(ctx context.Context) DpdkCpusPtrOutput {
	return o
}

type DpdkCpusArrayOutput struct{ *pulumi.OutputState }

func (DpdkCpusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DpdkCpus)(nil))
}

func (o DpdkCpusArrayOutput) ToDpdkCpusArrayOutput() DpdkCpusArrayOutput {
	return o
}

func (o DpdkCpusArrayOutput) ToDpdkCpusArrayOutputWithContext(ctx context.Context) DpdkCpusArrayOutput {
	return o
}

func (o DpdkCpusArrayOutput) Index(i pulumi.IntInput) DpdkCpusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DpdkCpus {
		return vs[0].([]DpdkCpus)[vs[1].(int)]
	}).(DpdkCpusOutput)
}

type DpdkCpusMapOutput struct{ *pulumi.OutputState }

func (DpdkCpusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DpdkCpus)(nil))
}

func (o DpdkCpusMapOutput) ToDpdkCpusMapOutput() DpdkCpusMapOutput {
	return o
}

func (o DpdkCpusMapOutput) ToDpdkCpusMapOutputWithContext(ctx context.Context) DpdkCpusMapOutput {
	return o
}

func (o DpdkCpusMapOutput) MapIndex(k pulumi.StringInput) DpdkCpusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DpdkCpus {
		return vs[0].(map[string]DpdkCpus)[vs[1].(string)]
	}).(DpdkCpusOutput)
}

func init() {
	pulumi.RegisterOutputType(DpdkCpusOutput{})
	pulumi.RegisterOutputType(DpdkCpusPtrOutput{})
	pulumi.RegisterOutputType(DpdkCpusArrayOutput{})
	pulumi.RegisterOutputType(DpdkCpusMapOutput{})
}
