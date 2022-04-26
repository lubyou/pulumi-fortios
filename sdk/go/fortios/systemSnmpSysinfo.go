// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SNMP system info configuration.
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
// 		_, err := fortios.NewSystemSnmpSysinfo(ctx, "trname", &fortios.SystemSnmpSysinfoArgs{
// 			Status:                 pulumi.String("disable"),
// 			TrapHighCpuThreshold:   pulumi.Int(80),
// 			TrapLogFullThreshold:   pulumi.Int(90),
// 			TrapLowMemoryThreshold: pulumi.Int(80),
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
// SystemSnmp Sysinfo can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemSnmpSysinfo:SystemSnmpSysinfo labelname SystemSnmpSysinfo
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemSnmpSysinfo:SystemSnmpSysinfo labelname SystemSnmpSysinfo
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemSnmpSysinfo struct {
	pulumi.CustomResourceState

	// Contact information.
	ContactInfo pulumi.StringPtrOutput `pulumi:"contactInfo"`
	// System description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Local SNMP engineID string (maximum 24 characters).
	EngineId pulumi.StringOutput `pulumi:"engineId"`
	// Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.
	EngineIdType pulumi.StringOutput `pulumi:"engineIdType"`
	// System location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Enable/disable SNMP. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// CPU usage when trap is sent.
	TrapHighCpuThreshold pulumi.IntOutput `pulumi:"trapHighCpuThreshold"`
	// Log disk usage when trap is sent.
	TrapLogFullThreshold pulumi.IntOutput `pulumi:"trapLogFullThreshold"`
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold pulumi.IntOutput `pulumi:"trapLowMemoryThreshold"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemSnmpSysinfo registers a new resource with the given unique name, arguments, and options.
func NewSystemSnmpSysinfo(ctx *pulumi.Context,
	name string, args *SystemSnmpSysinfoArgs, opts ...pulumi.ResourceOption) (*SystemSnmpSysinfo, error) {
	if args == nil {
		args = &SystemSnmpSysinfoArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemSnmpSysinfo
	err := ctx.RegisterResource("fortios:index/systemSnmpSysinfo:SystemSnmpSysinfo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSnmpSysinfo gets an existing SystemSnmpSysinfo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSnmpSysinfo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSnmpSysinfoState, opts ...pulumi.ResourceOption) (*SystemSnmpSysinfo, error) {
	var resource SystemSnmpSysinfo
	err := ctx.ReadResource("fortios:index/systemSnmpSysinfo:SystemSnmpSysinfo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSnmpSysinfo resources.
type systemSnmpSysinfoState struct {
	// Contact information.
	ContactInfo *string `pulumi:"contactInfo"`
	// System description.
	Description *string `pulumi:"description"`
	// Local SNMP engineID string (maximum 24 characters).
	EngineId *string `pulumi:"engineId"`
	// Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.
	EngineIdType *string `pulumi:"engineIdType"`
	// System location.
	Location *string `pulumi:"location"`
	// Enable/disable SNMP. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// CPU usage when trap is sent.
	TrapHighCpuThreshold *int `pulumi:"trapHighCpuThreshold"`
	// Log disk usage when trap is sent.
	TrapLogFullThreshold *int `pulumi:"trapLogFullThreshold"`
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold *int `pulumi:"trapLowMemoryThreshold"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemSnmpSysinfoState struct {
	// Contact information.
	ContactInfo pulumi.StringPtrInput
	// System description.
	Description pulumi.StringPtrInput
	// Local SNMP engineID string (maximum 24 characters).
	EngineId pulumi.StringPtrInput
	// Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.
	EngineIdType pulumi.StringPtrInput
	// System location.
	Location pulumi.StringPtrInput
	// Enable/disable SNMP. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// CPU usage when trap is sent.
	TrapHighCpuThreshold pulumi.IntPtrInput
	// Log disk usage when trap is sent.
	TrapLogFullThreshold pulumi.IntPtrInput
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSnmpSysinfoState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSnmpSysinfoState)(nil)).Elem()
}

type systemSnmpSysinfoArgs struct {
	// Contact information.
	ContactInfo *string `pulumi:"contactInfo"`
	// System description.
	Description *string `pulumi:"description"`
	// Local SNMP engineID string (maximum 24 characters).
	EngineId *string `pulumi:"engineId"`
	// Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.
	EngineIdType *string `pulumi:"engineIdType"`
	// System location.
	Location *string `pulumi:"location"`
	// Enable/disable SNMP. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// CPU usage when trap is sent.
	TrapHighCpuThreshold *int `pulumi:"trapHighCpuThreshold"`
	// Log disk usage when trap is sent.
	TrapLogFullThreshold *int `pulumi:"trapLogFullThreshold"`
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold *int `pulumi:"trapLowMemoryThreshold"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemSnmpSysinfo resource.
type SystemSnmpSysinfoArgs struct {
	// Contact information.
	ContactInfo pulumi.StringPtrInput
	// System description.
	Description pulumi.StringPtrInput
	// Local SNMP engineID string (maximum 24 characters).
	EngineId pulumi.StringPtrInput
	// Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.
	EngineIdType pulumi.StringPtrInput
	// System location.
	Location pulumi.StringPtrInput
	// Enable/disable SNMP. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// CPU usage when trap is sent.
	TrapHighCpuThreshold pulumi.IntPtrInput
	// Log disk usage when trap is sent.
	TrapLogFullThreshold pulumi.IntPtrInput
	// Memory usage when trap is sent.
	TrapLowMemoryThreshold pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSnmpSysinfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSnmpSysinfoArgs)(nil)).Elem()
}

type SystemSnmpSysinfoInput interface {
	pulumi.Input

	ToSystemSnmpSysinfoOutput() SystemSnmpSysinfoOutput
	ToSystemSnmpSysinfoOutputWithContext(ctx context.Context) SystemSnmpSysinfoOutput
}

func (*SystemSnmpSysinfo) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSnmpSysinfo)(nil)).Elem()
}

func (i *SystemSnmpSysinfo) ToSystemSnmpSysinfoOutput() SystemSnmpSysinfoOutput {
	return i.ToSystemSnmpSysinfoOutputWithContext(context.Background())
}

func (i *SystemSnmpSysinfo) ToSystemSnmpSysinfoOutputWithContext(ctx context.Context) SystemSnmpSysinfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSnmpSysinfoOutput)
}

// SystemSnmpSysinfoArrayInput is an input type that accepts SystemSnmpSysinfoArray and SystemSnmpSysinfoArrayOutput values.
// You can construct a concrete instance of `SystemSnmpSysinfoArrayInput` via:
//
//          SystemSnmpSysinfoArray{ SystemSnmpSysinfoArgs{...} }
type SystemSnmpSysinfoArrayInput interface {
	pulumi.Input

	ToSystemSnmpSysinfoArrayOutput() SystemSnmpSysinfoArrayOutput
	ToSystemSnmpSysinfoArrayOutputWithContext(context.Context) SystemSnmpSysinfoArrayOutput
}

type SystemSnmpSysinfoArray []SystemSnmpSysinfoInput

func (SystemSnmpSysinfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSnmpSysinfo)(nil)).Elem()
}

func (i SystemSnmpSysinfoArray) ToSystemSnmpSysinfoArrayOutput() SystemSnmpSysinfoArrayOutput {
	return i.ToSystemSnmpSysinfoArrayOutputWithContext(context.Background())
}

func (i SystemSnmpSysinfoArray) ToSystemSnmpSysinfoArrayOutputWithContext(ctx context.Context) SystemSnmpSysinfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSnmpSysinfoArrayOutput)
}

// SystemSnmpSysinfoMapInput is an input type that accepts SystemSnmpSysinfoMap and SystemSnmpSysinfoMapOutput values.
// You can construct a concrete instance of `SystemSnmpSysinfoMapInput` via:
//
//          SystemSnmpSysinfoMap{ "key": SystemSnmpSysinfoArgs{...} }
type SystemSnmpSysinfoMapInput interface {
	pulumi.Input

	ToSystemSnmpSysinfoMapOutput() SystemSnmpSysinfoMapOutput
	ToSystemSnmpSysinfoMapOutputWithContext(context.Context) SystemSnmpSysinfoMapOutput
}

type SystemSnmpSysinfoMap map[string]SystemSnmpSysinfoInput

func (SystemSnmpSysinfoMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSnmpSysinfo)(nil)).Elem()
}

func (i SystemSnmpSysinfoMap) ToSystemSnmpSysinfoMapOutput() SystemSnmpSysinfoMapOutput {
	return i.ToSystemSnmpSysinfoMapOutputWithContext(context.Background())
}

func (i SystemSnmpSysinfoMap) ToSystemSnmpSysinfoMapOutputWithContext(ctx context.Context) SystemSnmpSysinfoMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSnmpSysinfoMapOutput)
}

type SystemSnmpSysinfoOutput struct{ *pulumi.OutputState }

func (SystemSnmpSysinfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSnmpSysinfo)(nil)).Elem()
}

func (o SystemSnmpSysinfoOutput) ToSystemSnmpSysinfoOutput() SystemSnmpSysinfoOutput {
	return o
}

func (o SystemSnmpSysinfoOutput) ToSystemSnmpSysinfoOutputWithContext(ctx context.Context) SystemSnmpSysinfoOutput {
	return o
}

type SystemSnmpSysinfoArrayOutput struct{ *pulumi.OutputState }

func (SystemSnmpSysinfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSnmpSysinfo)(nil)).Elem()
}

func (o SystemSnmpSysinfoArrayOutput) ToSystemSnmpSysinfoArrayOutput() SystemSnmpSysinfoArrayOutput {
	return o
}

func (o SystemSnmpSysinfoArrayOutput) ToSystemSnmpSysinfoArrayOutputWithContext(ctx context.Context) SystemSnmpSysinfoArrayOutput {
	return o
}

func (o SystemSnmpSysinfoArrayOutput) Index(i pulumi.IntInput) SystemSnmpSysinfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSnmpSysinfo {
		return vs[0].([]*SystemSnmpSysinfo)[vs[1].(int)]
	}).(SystemSnmpSysinfoOutput)
}

type SystemSnmpSysinfoMapOutput struct{ *pulumi.OutputState }

func (SystemSnmpSysinfoMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSnmpSysinfo)(nil)).Elem()
}

func (o SystemSnmpSysinfoMapOutput) ToSystemSnmpSysinfoMapOutput() SystemSnmpSysinfoMapOutput {
	return o
}

func (o SystemSnmpSysinfoMapOutput) ToSystemSnmpSysinfoMapOutputWithContext(ctx context.Context) SystemSnmpSysinfoMapOutput {
	return o
}

func (o SystemSnmpSysinfoMapOutput) MapIndex(k pulumi.StringInput) SystemSnmpSysinfoOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSnmpSysinfo {
		return vs[0].(map[string]*SystemSnmpSysinfo)[vs[1].(string)]
	}).(SystemSnmpSysinfoOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSnmpSysinfoInput)(nil)).Elem(), &SystemSnmpSysinfo{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSnmpSysinfoArrayInput)(nil)).Elem(), SystemSnmpSysinfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSnmpSysinfoMapInput)(nil)).Elem(), SystemSnmpSysinfoMap{})
	pulumi.RegisterOutputType(SystemSnmpSysinfoOutput{})
	pulumi.RegisterOutputType(SystemSnmpSysinfoArrayOutput{})
	pulumi.RegisterOutputType(SystemSnmpSysinfoMapOutput{})
}
