// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemNpu struct {
	pulumi.CustomResourceState

	CapwapOffload               pulumi.StringOutput             `pulumi:"capwapOffload"`
	DedicatedManagementAffinity pulumi.StringOutput             `pulumi:"dedicatedManagementAffinity"`
	DedicatedManagementCpu      pulumi.StringOutput             `pulumi:"dedicatedManagementCpu"`
	Fastpath                    pulumi.StringOutput             `pulumi:"fastpath"`
	GetAllTables                pulumi.StringPtrOutput          `pulumi:"getAllTables"`
	IpsecDecSubengineMask       pulumi.StringOutput             `pulumi:"ipsecDecSubengineMask"`
	IpsecEncSubengineMask       pulumi.StringOutput             `pulumi:"ipsecEncSubengineMask"`
	IpsecInboundCache           pulumi.StringOutput             `pulumi:"ipsecInboundCache"`
	IpsecMtuOverride            pulumi.StringOutput             `pulumi:"ipsecMtuOverride"`
	IpsecOverVlink              pulumi.StringOutput             `pulumi:"ipsecOverVlink"`
	McastSessionAccounting      pulumi.StringOutput             `pulumi:"mcastSessionAccounting"`
	Np6CpsOptimizationMode      pulumi.StringOutput             `pulumi:"np6CpsOptimizationMode"`
	PriorityProtocol            SystemNpuPriorityProtocolOutput `pulumi:"priorityProtocol"`
	RdpOffload                  pulumi.StringOutput             `pulumi:"rdpOffload"`
	SessionDeniedOffload        pulumi.StringOutput             `pulumi:"sessionDeniedOffload"`
	SseBackpressure             pulumi.StringOutput             `pulumi:"sseBackpressure"`
	StripClearTextPadding       pulumi.StringOutput             `pulumi:"stripClearTextPadding"`
	StripEspPadding             pulumi.StringOutput             `pulumi:"stripEspPadding"`
	SwNpBandwidth               pulumi.StringOutput             `pulumi:"swNpBandwidth"`
	UespOffload                 pulumi.StringOutput             `pulumi:"uespOffload"`
	Vdomparam                   pulumi.StringPtrOutput          `pulumi:"vdomparam"`
}

// NewSystemNpu registers a new resource with the given unique name, arguments, and options.
func NewSystemNpu(ctx *pulumi.Context,
	name string, args *SystemNpuArgs, opts ...pulumi.ResourceOption) (*SystemNpu, error) {
	if args == nil {
		args = &SystemNpuArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemNpu
	err := ctx.RegisterResource("fortios:index/systemNpu:SystemNpu", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemNpu gets an existing SystemNpu resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemNpu(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemNpuState, opts ...pulumi.ResourceOption) (*SystemNpu, error) {
	var resource SystemNpu
	err := ctx.ReadResource("fortios:index/systemNpu:SystemNpu", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemNpu resources.
type systemNpuState struct {
	CapwapOffload               *string                    `pulumi:"capwapOffload"`
	DedicatedManagementAffinity *string                    `pulumi:"dedicatedManagementAffinity"`
	DedicatedManagementCpu      *string                    `pulumi:"dedicatedManagementCpu"`
	Fastpath                    *string                    `pulumi:"fastpath"`
	GetAllTables                *string                    `pulumi:"getAllTables"`
	IpsecDecSubengineMask       *string                    `pulumi:"ipsecDecSubengineMask"`
	IpsecEncSubengineMask       *string                    `pulumi:"ipsecEncSubengineMask"`
	IpsecInboundCache           *string                    `pulumi:"ipsecInboundCache"`
	IpsecMtuOverride            *string                    `pulumi:"ipsecMtuOverride"`
	IpsecOverVlink              *string                    `pulumi:"ipsecOverVlink"`
	McastSessionAccounting      *string                    `pulumi:"mcastSessionAccounting"`
	Np6CpsOptimizationMode      *string                    `pulumi:"np6CpsOptimizationMode"`
	PriorityProtocol            *SystemNpuPriorityProtocol `pulumi:"priorityProtocol"`
	RdpOffload                  *string                    `pulumi:"rdpOffload"`
	SessionDeniedOffload        *string                    `pulumi:"sessionDeniedOffload"`
	SseBackpressure             *string                    `pulumi:"sseBackpressure"`
	StripClearTextPadding       *string                    `pulumi:"stripClearTextPadding"`
	StripEspPadding             *string                    `pulumi:"stripEspPadding"`
	SwNpBandwidth               *string                    `pulumi:"swNpBandwidth"`
	UespOffload                 *string                    `pulumi:"uespOffload"`
	Vdomparam                   *string                    `pulumi:"vdomparam"`
}

type SystemNpuState struct {
	CapwapOffload               pulumi.StringPtrInput
	DedicatedManagementAffinity pulumi.StringPtrInput
	DedicatedManagementCpu      pulumi.StringPtrInput
	Fastpath                    pulumi.StringPtrInput
	GetAllTables                pulumi.StringPtrInput
	IpsecDecSubengineMask       pulumi.StringPtrInput
	IpsecEncSubengineMask       pulumi.StringPtrInput
	IpsecInboundCache           pulumi.StringPtrInput
	IpsecMtuOverride            pulumi.StringPtrInput
	IpsecOverVlink              pulumi.StringPtrInput
	McastSessionAccounting      pulumi.StringPtrInput
	Np6CpsOptimizationMode      pulumi.StringPtrInput
	PriorityProtocol            SystemNpuPriorityProtocolPtrInput
	RdpOffload                  pulumi.StringPtrInput
	SessionDeniedOffload        pulumi.StringPtrInput
	SseBackpressure             pulumi.StringPtrInput
	StripClearTextPadding       pulumi.StringPtrInput
	StripEspPadding             pulumi.StringPtrInput
	SwNpBandwidth               pulumi.StringPtrInput
	UespOffload                 pulumi.StringPtrInput
	Vdomparam                   pulumi.StringPtrInput
}

func (SystemNpuState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemNpuState)(nil)).Elem()
}

type systemNpuArgs struct {
	CapwapOffload               *string                    `pulumi:"capwapOffload"`
	DedicatedManagementAffinity *string                    `pulumi:"dedicatedManagementAffinity"`
	DedicatedManagementCpu      *string                    `pulumi:"dedicatedManagementCpu"`
	Fastpath                    *string                    `pulumi:"fastpath"`
	GetAllTables                *string                    `pulumi:"getAllTables"`
	IpsecDecSubengineMask       *string                    `pulumi:"ipsecDecSubengineMask"`
	IpsecEncSubengineMask       *string                    `pulumi:"ipsecEncSubengineMask"`
	IpsecInboundCache           *string                    `pulumi:"ipsecInboundCache"`
	IpsecMtuOverride            *string                    `pulumi:"ipsecMtuOverride"`
	IpsecOverVlink              *string                    `pulumi:"ipsecOverVlink"`
	McastSessionAccounting      *string                    `pulumi:"mcastSessionAccounting"`
	Np6CpsOptimizationMode      *string                    `pulumi:"np6CpsOptimizationMode"`
	PriorityProtocol            *SystemNpuPriorityProtocol `pulumi:"priorityProtocol"`
	RdpOffload                  *string                    `pulumi:"rdpOffload"`
	SessionDeniedOffload        *string                    `pulumi:"sessionDeniedOffload"`
	SseBackpressure             *string                    `pulumi:"sseBackpressure"`
	StripClearTextPadding       *string                    `pulumi:"stripClearTextPadding"`
	StripEspPadding             *string                    `pulumi:"stripEspPadding"`
	SwNpBandwidth               *string                    `pulumi:"swNpBandwidth"`
	UespOffload                 *string                    `pulumi:"uespOffload"`
	Vdomparam                   *string                    `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemNpu resource.
type SystemNpuArgs struct {
	CapwapOffload               pulumi.StringPtrInput
	DedicatedManagementAffinity pulumi.StringPtrInput
	DedicatedManagementCpu      pulumi.StringPtrInput
	Fastpath                    pulumi.StringPtrInput
	GetAllTables                pulumi.StringPtrInput
	IpsecDecSubengineMask       pulumi.StringPtrInput
	IpsecEncSubengineMask       pulumi.StringPtrInput
	IpsecInboundCache           pulumi.StringPtrInput
	IpsecMtuOverride            pulumi.StringPtrInput
	IpsecOverVlink              pulumi.StringPtrInput
	McastSessionAccounting      pulumi.StringPtrInput
	Np6CpsOptimizationMode      pulumi.StringPtrInput
	PriorityProtocol            SystemNpuPriorityProtocolPtrInput
	RdpOffload                  pulumi.StringPtrInput
	SessionDeniedOffload        pulumi.StringPtrInput
	SseBackpressure             pulumi.StringPtrInput
	StripClearTextPadding       pulumi.StringPtrInput
	StripEspPadding             pulumi.StringPtrInput
	SwNpBandwidth               pulumi.StringPtrInput
	UespOffload                 pulumi.StringPtrInput
	Vdomparam                   pulumi.StringPtrInput
}

func (SystemNpuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemNpuArgs)(nil)).Elem()
}

type SystemNpuInput interface {
	pulumi.Input

	ToSystemNpuOutput() SystemNpuOutput
	ToSystemNpuOutputWithContext(ctx context.Context) SystemNpuOutput
}

func (*SystemNpu) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemNpu)(nil)).Elem()
}

func (i *SystemNpu) ToSystemNpuOutput() SystemNpuOutput {
	return i.ToSystemNpuOutputWithContext(context.Background())
}

func (i *SystemNpu) ToSystemNpuOutputWithContext(ctx context.Context) SystemNpuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNpuOutput)
}

// SystemNpuArrayInput is an input type that accepts SystemNpuArray and SystemNpuArrayOutput values.
// You can construct a concrete instance of `SystemNpuArrayInput` via:
//
//	SystemNpuArray{ SystemNpuArgs{...} }
type SystemNpuArrayInput interface {
	pulumi.Input

	ToSystemNpuArrayOutput() SystemNpuArrayOutput
	ToSystemNpuArrayOutputWithContext(context.Context) SystemNpuArrayOutput
}

type SystemNpuArray []SystemNpuInput

func (SystemNpuArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemNpu)(nil)).Elem()
}

func (i SystemNpuArray) ToSystemNpuArrayOutput() SystemNpuArrayOutput {
	return i.ToSystemNpuArrayOutputWithContext(context.Background())
}

func (i SystemNpuArray) ToSystemNpuArrayOutputWithContext(ctx context.Context) SystemNpuArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNpuArrayOutput)
}

// SystemNpuMapInput is an input type that accepts SystemNpuMap and SystemNpuMapOutput values.
// You can construct a concrete instance of `SystemNpuMapInput` via:
//
//	SystemNpuMap{ "key": SystemNpuArgs{...} }
type SystemNpuMapInput interface {
	pulumi.Input

	ToSystemNpuMapOutput() SystemNpuMapOutput
	ToSystemNpuMapOutputWithContext(context.Context) SystemNpuMapOutput
}

type SystemNpuMap map[string]SystemNpuInput

func (SystemNpuMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemNpu)(nil)).Elem()
}

func (i SystemNpuMap) ToSystemNpuMapOutput() SystemNpuMapOutput {
	return i.ToSystemNpuMapOutputWithContext(context.Background())
}

func (i SystemNpuMap) ToSystemNpuMapOutputWithContext(ctx context.Context) SystemNpuMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNpuMapOutput)
}

type SystemNpuOutput struct{ *pulumi.OutputState }

func (SystemNpuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemNpu)(nil)).Elem()
}

func (o SystemNpuOutput) ToSystemNpuOutput() SystemNpuOutput {
	return o
}

func (o SystemNpuOutput) ToSystemNpuOutputWithContext(ctx context.Context) SystemNpuOutput {
	return o
}

func (o SystemNpuOutput) CapwapOffload() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.CapwapOffload }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) DedicatedManagementAffinity() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.DedicatedManagementAffinity }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) DedicatedManagementCpu() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.DedicatedManagementCpu }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) Fastpath() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.Fastpath }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o SystemNpuOutput) IpsecDecSubengineMask() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.IpsecDecSubengineMask }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) IpsecEncSubengineMask() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.IpsecEncSubengineMask }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) IpsecInboundCache() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.IpsecInboundCache }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) IpsecMtuOverride() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.IpsecMtuOverride }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) IpsecOverVlink() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.IpsecOverVlink }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) McastSessionAccounting() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.McastSessionAccounting }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) Np6CpsOptimizationMode() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.Np6CpsOptimizationMode }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) PriorityProtocol() SystemNpuPriorityProtocolOutput {
	return o.ApplyT(func(v *SystemNpu) SystemNpuPriorityProtocolOutput { return v.PriorityProtocol }).(SystemNpuPriorityProtocolOutput)
}

func (o SystemNpuOutput) RdpOffload() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.RdpOffload }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) SessionDeniedOffload() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.SessionDeniedOffload }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) SseBackpressure() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.SseBackpressure }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) StripClearTextPadding() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.StripClearTextPadding }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) StripEspPadding() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.StripEspPadding }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) SwNpBandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.SwNpBandwidth }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) UespOffload() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringOutput { return v.UespOffload }).(pulumi.StringOutput)
}

func (o SystemNpuOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemNpu) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemNpuArrayOutput struct{ *pulumi.OutputState }

func (SystemNpuArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemNpu)(nil)).Elem()
}

func (o SystemNpuArrayOutput) ToSystemNpuArrayOutput() SystemNpuArrayOutput {
	return o
}

func (o SystemNpuArrayOutput) ToSystemNpuArrayOutputWithContext(ctx context.Context) SystemNpuArrayOutput {
	return o
}

func (o SystemNpuArrayOutput) Index(i pulumi.IntInput) SystemNpuOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemNpu {
		return vs[0].([]*SystemNpu)[vs[1].(int)]
	}).(SystemNpuOutput)
}

type SystemNpuMapOutput struct{ *pulumi.OutputState }

func (SystemNpuMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemNpu)(nil)).Elem()
}

func (o SystemNpuMapOutput) ToSystemNpuMapOutput() SystemNpuMapOutput {
	return o
}

func (o SystemNpuMapOutput) ToSystemNpuMapOutputWithContext(ctx context.Context) SystemNpuMapOutput {
	return o
}

func (o SystemNpuMapOutput) MapIndex(k pulumi.StringInput) SystemNpuOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemNpu {
		return vs[0].(map[string]*SystemNpu)[vs[1].(string)]
	}).(SystemNpuOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemNpuInput)(nil)).Elem(), &SystemNpu{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemNpuArrayInput)(nil)).Elem(), SystemNpuArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemNpuMapInput)(nil)).Elem(), SystemNpuMap{})
	pulumi.RegisterOutputType(SystemNpuOutput{})
	pulumi.RegisterOutputType(SystemNpuArrayOutput{})
	pulumi.RegisterOutputType(SystemNpuMapOutput{})
}
