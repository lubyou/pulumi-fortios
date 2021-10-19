// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiManager.
//
// Due to the limitations of the current system, the feature is temporarily unavailable. Please use the following resource configuration as an alternative.
//
// ## Example
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
// 		_, err := fortios.NewSystemCentralManagement(ctx, "trname", &fortios.SystemCentralManagementArgs{
// 			AllowMonitor:               pulumi.String("enable"),
// 			AllowPushConfiguration:     pulumi.String("enable"),
// 			AllowPushFirmware:          pulumi.String("enable"),
// 			AllowRemoteFirmwareUpgrade: pulumi.String("enable"),
// 			EncAlgorithm:               pulumi.String("high"),
// 			Fmg:                        pulumi.String("\"192.168.52.177\""),
// 			IncludeDefaultServers:      pulumi.String("enable"),
// 			Mode:                       pulumi.String("normal"),
// 			Type:                       pulumi.String("fortimanager"),
// 			Vdom:                       pulumi.String("root"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SystemFortimanager struct {
	pulumi.CustomResourceState

	CentralManagement                pulumi.StringOutput    `pulumi:"centralManagement"`
	CentralMgmtAutoBackup            pulumi.StringOutput    `pulumi:"centralMgmtAutoBackup"`
	CentralMgmtScheduleConfigRestore pulumi.StringOutput    `pulumi:"centralMgmtScheduleConfigRestore"`
	CentralMgmtScheduleScriptRestore pulumi.StringOutput    `pulumi:"centralMgmtScheduleScriptRestore"`
	Ip                               pulumi.StringOutput    `pulumi:"ip"`
	Ipsec                            pulumi.StringOutput    `pulumi:"ipsec"`
	Vdom                             pulumi.StringOutput    `pulumi:"vdom"`
	Vdomparam                        pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemFortimanager registers a new resource with the given unique name, arguments, and options.
func NewSystemFortimanager(ctx *pulumi.Context,
	name string, args *SystemFortimanagerArgs, opts ...pulumi.ResourceOption) (*SystemFortimanager, error) {
	if args == nil {
		args = &SystemFortimanagerArgs{}
	}

	var resource SystemFortimanager
	err := ctx.RegisterResource("fortios:index/systemFortimanager:SystemFortimanager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemFortimanager gets an existing SystemFortimanager resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemFortimanager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemFortimanagerState, opts ...pulumi.ResourceOption) (*SystemFortimanager, error) {
	var resource SystemFortimanager
	err := ctx.ReadResource("fortios:index/systemFortimanager:SystemFortimanager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemFortimanager resources.
type systemFortimanagerState struct {
	CentralManagement                *string `pulumi:"centralManagement"`
	CentralMgmtAutoBackup            *string `pulumi:"centralMgmtAutoBackup"`
	CentralMgmtScheduleConfigRestore *string `pulumi:"centralMgmtScheduleConfigRestore"`
	CentralMgmtScheduleScriptRestore *string `pulumi:"centralMgmtScheduleScriptRestore"`
	Ip                               *string `pulumi:"ip"`
	Ipsec                            *string `pulumi:"ipsec"`
	Vdom                             *string `pulumi:"vdom"`
	Vdomparam                        *string `pulumi:"vdomparam"`
}

type SystemFortimanagerState struct {
	CentralManagement                pulumi.StringPtrInput
	CentralMgmtAutoBackup            pulumi.StringPtrInput
	CentralMgmtScheduleConfigRestore pulumi.StringPtrInput
	CentralMgmtScheduleScriptRestore pulumi.StringPtrInput
	Ip                               pulumi.StringPtrInput
	Ipsec                            pulumi.StringPtrInput
	Vdom                             pulumi.StringPtrInput
	Vdomparam                        pulumi.StringPtrInput
}

func (SystemFortimanagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFortimanagerState)(nil)).Elem()
}

type systemFortimanagerArgs struct {
	CentralManagement                *string `pulumi:"centralManagement"`
	CentralMgmtAutoBackup            *string `pulumi:"centralMgmtAutoBackup"`
	CentralMgmtScheduleConfigRestore *string `pulumi:"centralMgmtScheduleConfigRestore"`
	CentralMgmtScheduleScriptRestore *string `pulumi:"centralMgmtScheduleScriptRestore"`
	Ip                               *string `pulumi:"ip"`
	Ipsec                            *string `pulumi:"ipsec"`
	Vdom                             *string `pulumi:"vdom"`
	Vdomparam                        *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemFortimanager resource.
type SystemFortimanagerArgs struct {
	CentralManagement                pulumi.StringPtrInput
	CentralMgmtAutoBackup            pulumi.StringPtrInput
	CentralMgmtScheduleConfigRestore pulumi.StringPtrInput
	CentralMgmtScheduleScriptRestore pulumi.StringPtrInput
	Ip                               pulumi.StringPtrInput
	Ipsec                            pulumi.StringPtrInput
	Vdom                             pulumi.StringPtrInput
	Vdomparam                        pulumi.StringPtrInput
}

func (SystemFortimanagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFortimanagerArgs)(nil)).Elem()
}

type SystemFortimanagerInput interface {
	pulumi.Input

	ToSystemFortimanagerOutput() SystemFortimanagerOutput
	ToSystemFortimanagerOutputWithContext(ctx context.Context) SystemFortimanagerOutput
}

func (*SystemFortimanager) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemFortimanager)(nil))
}

func (i *SystemFortimanager) ToSystemFortimanagerOutput() SystemFortimanagerOutput {
	return i.ToSystemFortimanagerOutputWithContext(context.Background())
}

func (i *SystemFortimanager) ToSystemFortimanagerOutputWithContext(ctx context.Context) SystemFortimanagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFortimanagerOutput)
}

func (i *SystemFortimanager) ToSystemFortimanagerPtrOutput() SystemFortimanagerPtrOutput {
	return i.ToSystemFortimanagerPtrOutputWithContext(context.Background())
}

func (i *SystemFortimanager) ToSystemFortimanagerPtrOutputWithContext(ctx context.Context) SystemFortimanagerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFortimanagerPtrOutput)
}

type SystemFortimanagerPtrInput interface {
	pulumi.Input

	ToSystemFortimanagerPtrOutput() SystemFortimanagerPtrOutput
	ToSystemFortimanagerPtrOutputWithContext(ctx context.Context) SystemFortimanagerPtrOutput
}

type systemFortimanagerPtrType SystemFortimanagerArgs

func (*systemFortimanagerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFortimanager)(nil))
}

func (i *systemFortimanagerPtrType) ToSystemFortimanagerPtrOutput() SystemFortimanagerPtrOutput {
	return i.ToSystemFortimanagerPtrOutputWithContext(context.Background())
}

func (i *systemFortimanagerPtrType) ToSystemFortimanagerPtrOutputWithContext(ctx context.Context) SystemFortimanagerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFortimanagerPtrOutput)
}

// SystemFortimanagerArrayInput is an input type that accepts SystemFortimanagerArray and SystemFortimanagerArrayOutput values.
// You can construct a concrete instance of `SystemFortimanagerArrayInput` via:
//
//          SystemFortimanagerArray{ SystemFortimanagerArgs{...} }
type SystemFortimanagerArrayInput interface {
	pulumi.Input

	ToSystemFortimanagerArrayOutput() SystemFortimanagerArrayOutput
	ToSystemFortimanagerArrayOutputWithContext(context.Context) SystemFortimanagerArrayOutput
}

type SystemFortimanagerArray []SystemFortimanagerInput

func (SystemFortimanagerArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SystemFortimanager)(nil))
}

func (i SystemFortimanagerArray) ToSystemFortimanagerArrayOutput() SystemFortimanagerArrayOutput {
	return i.ToSystemFortimanagerArrayOutputWithContext(context.Background())
}

func (i SystemFortimanagerArray) ToSystemFortimanagerArrayOutputWithContext(ctx context.Context) SystemFortimanagerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFortimanagerArrayOutput)
}

// SystemFortimanagerMapInput is an input type that accepts SystemFortimanagerMap and SystemFortimanagerMapOutput values.
// You can construct a concrete instance of `SystemFortimanagerMapInput` via:
//
//          SystemFortimanagerMap{ "key": SystemFortimanagerArgs{...} }
type SystemFortimanagerMapInput interface {
	pulumi.Input

	ToSystemFortimanagerMapOutput() SystemFortimanagerMapOutput
	ToSystemFortimanagerMapOutputWithContext(context.Context) SystemFortimanagerMapOutput
}

type SystemFortimanagerMap map[string]SystemFortimanagerInput

func (SystemFortimanagerMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SystemFortimanager)(nil))
}

func (i SystemFortimanagerMap) ToSystemFortimanagerMapOutput() SystemFortimanagerMapOutput {
	return i.ToSystemFortimanagerMapOutputWithContext(context.Background())
}

func (i SystemFortimanagerMap) ToSystemFortimanagerMapOutputWithContext(ctx context.Context) SystemFortimanagerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFortimanagerMapOutput)
}

type SystemFortimanagerOutput struct {
	*pulumi.OutputState
}

func (SystemFortimanagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemFortimanager)(nil))
}

func (o SystemFortimanagerOutput) ToSystemFortimanagerOutput() SystemFortimanagerOutput {
	return o
}

func (o SystemFortimanagerOutput) ToSystemFortimanagerOutputWithContext(ctx context.Context) SystemFortimanagerOutput {
	return o
}

func (o SystemFortimanagerOutput) ToSystemFortimanagerPtrOutput() SystemFortimanagerPtrOutput {
	return o.ToSystemFortimanagerPtrOutputWithContext(context.Background())
}

func (o SystemFortimanagerOutput) ToSystemFortimanagerPtrOutputWithContext(ctx context.Context) SystemFortimanagerPtrOutput {
	return o.ApplyT(func(v SystemFortimanager) *SystemFortimanager {
		return &v
	}).(SystemFortimanagerPtrOutput)
}

type SystemFortimanagerPtrOutput struct {
	*pulumi.OutputState
}

func (SystemFortimanagerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFortimanager)(nil))
}

func (o SystemFortimanagerPtrOutput) ToSystemFortimanagerPtrOutput() SystemFortimanagerPtrOutput {
	return o
}

func (o SystemFortimanagerPtrOutput) ToSystemFortimanagerPtrOutputWithContext(ctx context.Context) SystemFortimanagerPtrOutput {
	return o
}

type SystemFortimanagerArrayOutput struct{ *pulumi.OutputState }

func (SystemFortimanagerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemFortimanager)(nil))
}

func (o SystemFortimanagerArrayOutput) ToSystemFortimanagerArrayOutput() SystemFortimanagerArrayOutput {
	return o
}

func (o SystemFortimanagerArrayOutput) ToSystemFortimanagerArrayOutputWithContext(ctx context.Context) SystemFortimanagerArrayOutput {
	return o
}

func (o SystemFortimanagerArrayOutput) Index(i pulumi.IntInput) SystemFortimanagerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemFortimanager {
		return vs[0].([]SystemFortimanager)[vs[1].(int)]
	}).(SystemFortimanagerOutput)
}

type SystemFortimanagerMapOutput struct{ *pulumi.OutputState }

func (SystemFortimanagerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SystemFortimanager)(nil))
}

func (o SystemFortimanagerMapOutput) ToSystemFortimanagerMapOutput() SystemFortimanagerMapOutput {
	return o
}

func (o SystemFortimanagerMapOutput) ToSystemFortimanagerMapOutputWithContext(ctx context.Context) SystemFortimanagerMapOutput {
	return o
}

func (o SystemFortimanagerMapOutput) MapIndex(k pulumi.StringInput) SystemFortimanagerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SystemFortimanager {
		return vs[0].(map[string]SystemFortimanager)[vs[1].(string)]
	}).(SystemFortimanagerOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemFortimanagerOutput{})
	pulumi.RegisterOutputType(SystemFortimanagerPtrOutput{})
	pulumi.RegisterOutputType(SystemFortimanagerArrayOutput{})
	pulumi.RegisterOutputType(SystemFortimanagerMapOutput{})
}
