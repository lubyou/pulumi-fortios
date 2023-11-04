// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type WirelessControllerLog struct {
	pulumi.CustomResourceState

	AddrgrpLog    pulumi.StringOutput    `pulumi:"addrgrpLog"`
	BleLog        pulumi.StringOutput    `pulumi:"bleLog"`
	ClbLog        pulumi.StringOutput    `pulumi:"clbLog"`
	DhcpStarvLog  pulumi.StringOutput    `pulumi:"dhcpStarvLog"`
	LedSchedLog   pulumi.StringOutput    `pulumi:"ledSchedLog"`
	RadioEventLog pulumi.StringOutput    `pulumi:"radioEventLog"`
	RogueEventLog pulumi.StringOutput    `pulumi:"rogueEventLog"`
	StaEventLog   pulumi.StringOutput    `pulumi:"staEventLog"`
	StaLocateLog  pulumi.StringOutput    `pulumi:"staLocateLog"`
	Status        pulumi.StringOutput    `pulumi:"status"`
	Vdomparam     pulumi.StringPtrOutput `pulumi:"vdomparam"`
	WidsLog       pulumi.StringOutput    `pulumi:"widsLog"`
	WtpEventLog   pulumi.StringOutput    `pulumi:"wtpEventLog"`
}

// NewWirelessControllerLog registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerLog(ctx *pulumi.Context,
	name string, args *WirelessControllerLogArgs, opts ...pulumi.ResourceOption) (*WirelessControllerLog, error) {
	if args == nil {
		args = &WirelessControllerLogArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WirelessControllerLog
	err := ctx.RegisterResource("fortios:index/wirelessControllerLog:WirelessControllerLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerLog gets an existing WirelessControllerLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerLogState, opts ...pulumi.ResourceOption) (*WirelessControllerLog, error) {
	var resource WirelessControllerLog
	err := ctx.ReadResource("fortios:index/wirelessControllerLog:WirelessControllerLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerLog resources.
type wirelessControllerLogState struct {
	AddrgrpLog    *string `pulumi:"addrgrpLog"`
	BleLog        *string `pulumi:"bleLog"`
	ClbLog        *string `pulumi:"clbLog"`
	DhcpStarvLog  *string `pulumi:"dhcpStarvLog"`
	LedSchedLog   *string `pulumi:"ledSchedLog"`
	RadioEventLog *string `pulumi:"radioEventLog"`
	RogueEventLog *string `pulumi:"rogueEventLog"`
	StaEventLog   *string `pulumi:"staEventLog"`
	StaLocateLog  *string `pulumi:"staLocateLog"`
	Status        *string `pulumi:"status"`
	Vdomparam     *string `pulumi:"vdomparam"`
	WidsLog       *string `pulumi:"widsLog"`
	WtpEventLog   *string `pulumi:"wtpEventLog"`
}

type WirelessControllerLogState struct {
	AddrgrpLog    pulumi.StringPtrInput
	BleLog        pulumi.StringPtrInput
	ClbLog        pulumi.StringPtrInput
	DhcpStarvLog  pulumi.StringPtrInput
	LedSchedLog   pulumi.StringPtrInput
	RadioEventLog pulumi.StringPtrInput
	RogueEventLog pulumi.StringPtrInput
	StaEventLog   pulumi.StringPtrInput
	StaLocateLog  pulumi.StringPtrInput
	Status        pulumi.StringPtrInput
	Vdomparam     pulumi.StringPtrInput
	WidsLog       pulumi.StringPtrInput
	WtpEventLog   pulumi.StringPtrInput
}

func (WirelessControllerLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerLogState)(nil)).Elem()
}

type wirelessControllerLogArgs struct {
	AddrgrpLog    *string `pulumi:"addrgrpLog"`
	BleLog        *string `pulumi:"bleLog"`
	ClbLog        *string `pulumi:"clbLog"`
	DhcpStarvLog  *string `pulumi:"dhcpStarvLog"`
	LedSchedLog   *string `pulumi:"ledSchedLog"`
	RadioEventLog *string `pulumi:"radioEventLog"`
	RogueEventLog *string `pulumi:"rogueEventLog"`
	StaEventLog   *string `pulumi:"staEventLog"`
	StaLocateLog  *string `pulumi:"staLocateLog"`
	Status        *string `pulumi:"status"`
	Vdomparam     *string `pulumi:"vdomparam"`
	WidsLog       *string `pulumi:"widsLog"`
	WtpEventLog   *string `pulumi:"wtpEventLog"`
}

// The set of arguments for constructing a WirelessControllerLog resource.
type WirelessControllerLogArgs struct {
	AddrgrpLog    pulumi.StringPtrInput
	BleLog        pulumi.StringPtrInput
	ClbLog        pulumi.StringPtrInput
	DhcpStarvLog  pulumi.StringPtrInput
	LedSchedLog   pulumi.StringPtrInput
	RadioEventLog pulumi.StringPtrInput
	RogueEventLog pulumi.StringPtrInput
	StaEventLog   pulumi.StringPtrInput
	StaLocateLog  pulumi.StringPtrInput
	Status        pulumi.StringPtrInput
	Vdomparam     pulumi.StringPtrInput
	WidsLog       pulumi.StringPtrInput
	WtpEventLog   pulumi.StringPtrInput
}

func (WirelessControllerLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerLogArgs)(nil)).Elem()
}

type WirelessControllerLogInput interface {
	pulumi.Input

	ToWirelessControllerLogOutput() WirelessControllerLogOutput
	ToWirelessControllerLogOutputWithContext(ctx context.Context) WirelessControllerLogOutput
}

func (*WirelessControllerLog) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerLog)(nil)).Elem()
}

func (i *WirelessControllerLog) ToWirelessControllerLogOutput() WirelessControllerLogOutput {
	return i.ToWirelessControllerLogOutputWithContext(context.Background())
}

func (i *WirelessControllerLog) ToWirelessControllerLogOutputWithContext(ctx context.Context) WirelessControllerLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerLogOutput)
}

func (i *WirelessControllerLog) ToOutput(ctx context.Context) pulumix.Output[*WirelessControllerLog] {
	return pulumix.Output[*WirelessControllerLog]{
		OutputState: i.ToWirelessControllerLogOutputWithContext(ctx).OutputState,
	}
}

// WirelessControllerLogArrayInput is an input type that accepts WirelessControllerLogArray and WirelessControllerLogArrayOutput values.
// You can construct a concrete instance of `WirelessControllerLogArrayInput` via:
//
//	WirelessControllerLogArray{ WirelessControllerLogArgs{...} }
type WirelessControllerLogArrayInput interface {
	pulumi.Input

	ToWirelessControllerLogArrayOutput() WirelessControllerLogArrayOutput
	ToWirelessControllerLogArrayOutputWithContext(context.Context) WirelessControllerLogArrayOutput
}

type WirelessControllerLogArray []WirelessControllerLogInput

func (WirelessControllerLogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerLog)(nil)).Elem()
}

func (i WirelessControllerLogArray) ToWirelessControllerLogArrayOutput() WirelessControllerLogArrayOutput {
	return i.ToWirelessControllerLogArrayOutputWithContext(context.Background())
}

func (i WirelessControllerLogArray) ToWirelessControllerLogArrayOutputWithContext(ctx context.Context) WirelessControllerLogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerLogArrayOutput)
}

func (i WirelessControllerLogArray) ToOutput(ctx context.Context) pulumix.Output[[]*WirelessControllerLog] {
	return pulumix.Output[[]*WirelessControllerLog]{
		OutputState: i.ToWirelessControllerLogArrayOutputWithContext(ctx).OutputState,
	}
}

// WirelessControllerLogMapInput is an input type that accepts WirelessControllerLogMap and WirelessControllerLogMapOutput values.
// You can construct a concrete instance of `WirelessControllerLogMapInput` via:
//
//	WirelessControllerLogMap{ "key": WirelessControllerLogArgs{...} }
type WirelessControllerLogMapInput interface {
	pulumi.Input

	ToWirelessControllerLogMapOutput() WirelessControllerLogMapOutput
	ToWirelessControllerLogMapOutputWithContext(context.Context) WirelessControllerLogMapOutput
}

type WirelessControllerLogMap map[string]WirelessControllerLogInput

func (WirelessControllerLogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerLog)(nil)).Elem()
}

func (i WirelessControllerLogMap) ToWirelessControllerLogMapOutput() WirelessControllerLogMapOutput {
	return i.ToWirelessControllerLogMapOutputWithContext(context.Background())
}

func (i WirelessControllerLogMap) ToWirelessControllerLogMapOutputWithContext(ctx context.Context) WirelessControllerLogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerLogMapOutput)
}

func (i WirelessControllerLogMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*WirelessControllerLog] {
	return pulumix.Output[map[string]*WirelessControllerLog]{
		OutputState: i.ToWirelessControllerLogMapOutputWithContext(ctx).OutputState,
	}
}

type WirelessControllerLogOutput struct{ *pulumi.OutputState }

func (WirelessControllerLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerLog)(nil)).Elem()
}

func (o WirelessControllerLogOutput) ToWirelessControllerLogOutput() WirelessControllerLogOutput {
	return o
}

func (o WirelessControllerLogOutput) ToWirelessControllerLogOutputWithContext(ctx context.Context) WirelessControllerLogOutput {
	return o
}

func (o WirelessControllerLogOutput) ToOutput(ctx context.Context) pulumix.Output[*WirelessControllerLog] {
	return pulumix.Output[*WirelessControllerLog]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerLogOutput) AddrgrpLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerLog) pulumi.StringOutput { return v.AddrgrpLog }).(pulumi.StringOutput)
}

func (o WirelessControllerLogOutput) BleLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerLog) pulumi.StringOutput { return v.BleLog }).(pulumi.StringOutput)
}

func (o WirelessControllerLogOutput) ClbLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerLog) pulumi.StringOutput { return v.ClbLog }).(pulumi.StringOutput)
}

func (o WirelessControllerLogOutput) DhcpStarvLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerLog) pulumi.StringOutput { return v.DhcpStarvLog }).(pulumi.StringOutput)
}

func (o WirelessControllerLogOutput) LedSchedLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerLog) pulumi.StringOutput { return v.LedSchedLog }).(pulumi.StringOutput)
}

func (o WirelessControllerLogOutput) RadioEventLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerLog) pulumi.StringOutput { return v.RadioEventLog }).(pulumi.StringOutput)
}

func (o WirelessControllerLogOutput) RogueEventLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerLog) pulumi.StringOutput { return v.RogueEventLog }).(pulumi.StringOutput)
}

func (o WirelessControllerLogOutput) StaEventLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerLog) pulumi.StringOutput { return v.StaEventLog }).(pulumi.StringOutput)
}

func (o WirelessControllerLogOutput) StaLocateLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerLog) pulumi.StringOutput { return v.StaLocateLog }).(pulumi.StringOutput)
}

func (o WirelessControllerLogOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerLog) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o WirelessControllerLogOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerLog) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerLogOutput) WidsLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerLog) pulumi.StringOutput { return v.WidsLog }).(pulumi.StringOutput)
}

func (o WirelessControllerLogOutput) WtpEventLog() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerLog) pulumi.StringOutput { return v.WtpEventLog }).(pulumi.StringOutput)
}

type WirelessControllerLogArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerLogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerLog)(nil)).Elem()
}

func (o WirelessControllerLogArrayOutput) ToWirelessControllerLogArrayOutput() WirelessControllerLogArrayOutput {
	return o
}

func (o WirelessControllerLogArrayOutput) ToWirelessControllerLogArrayOutputWithContext(ctx context.Context) WirelessControllerLogArrayOutput {
	return o
}

func (o WirelessControllerLogArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*WirelessControllerLog] {
	return pulumix.Output[[]*WirelessControllerLog]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerLogArrayOutput) Index(i pulumi.IntInput) WirelessControllerLogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerLog {
		return vs[0].([]*WirelessControllerLog)[vs[1].(int)]
	}).(WirelessControllerLogOutput)
}

type WirelessControllerLogMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerLogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerLog)(nil)).Elem()
}

func (o WirelessControllerLogMapOutput) ToWirelessControllerLogMapOutput() WirelessControllerLogMapOutput {
	return o
}

func (o WirelessControllerLogMapOutput) ToWirelessControllerLogMapOutputWithContext(ctx context.Context) WirelessControllerLogMapOutput {
	return o
}

func (o WirelessControllerLogMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*WirelessControllerLog] {
	return pulumix.Output[map[string]*WirelessControllerLog]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerLogMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerLogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerLog {
		return vs[0].(map[string]*WirelessControllerLog)[vs[1].(string)]
	}).(WirelessControllerLogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerLogInput)(nil)).Elem(), &WirelessControllerLog{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerLogArrayInput)(nil)).Elem(), WirelessControllerLogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerLogMapInput)(nil)).Elem(), WirelessControllerLogMap{})
	pulumi.RegisterOutputType(WirelessControllerLogOutput{})
	pulumi.RegisterOutputType(WirelessControllerLogArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerLogMapOutput{})
}
