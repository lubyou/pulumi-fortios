// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch SNMP system information globally. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// SwitchController SnmpSysinfo can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/switchControllerSnmpSysinfo:SwitchControllerSnmpSysinfo labelname SwitchControllerSnmpSysinfo
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/switchControllerSnmpSysinfo:SwitchControllerSnmpSysinfo labelname SwitchControllerSnmpSysinfo
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SwitchControllerSnmpSysinfo struct {
	pulumi.CustomResourceState

	// Contact information.
	ContactInfo pulumi.StringOutput `pulumi:"contactInfo"`
	// System description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Local SNMP engine ID string (max 24 char).
	EngineId pulumi.StringOutput `pulumi:"engineId"`
	// System location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Enable/disable SNMP. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerSnmpSysinfo registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerSnmpSysinfo(ctx *pulumi.Context,
	name string, args *SwitchControllerSnmpSysinfoArgs, opts ...pulumi.ResourceOption) (*SwitchControllerSnmpSysinfo, error) {
	if args == nil {
		args = &SwitchControllerSnmpSysinfoArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchControllerSnmpSysinfo
	err := ctx.RegisterResource("fortios:index/switchControllerSnmpSysinfo:SwitchControllerSnmpSysinfo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerSnmpSysinfo gets an existing SwitchControllerSnmpSysinfo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerSnmpSysinfo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerSnmpSysinfoState, opts ...pulumi.ResourceOption) (*SwitchControllerSnmpSysinfo, error) {
	var resource SwitchControllerSnmpSysinfo
	err := ctx.ReadResource("fortios:index/switchControllerSnmpSysinfo:SwitchControllerSnmpSysinfo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerSnmpSysinfo resources.
type switchControllerSnmpSysinfoState struct {
	// Contact information.
	ContactInfo *string `pulumi:"contactInfo"`
	// System description.
	Description *string `pulumi:"description"`
	// Local SNMP engine ID string (max 24 char).
	EngineId *string `pulumi:"engineId"`
	// System location.
	Location *string `pulumi:"location"`
	// Enable/disable SNMP. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchControllerSnmpSysinfoState struct {
	// Contact information.
	ContactInfo pulumi.StringPtrInput
	// System description.
	Description pulumi.StringPtrInput
	// Local SNMP engine ID string (max 24 char).
	EngineId pulumi.StringPtrInput
	// System location.
	Location pulumi.StringPtrInput
	// Enable/disable SNMP. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerSnmpSysinfoState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSnmpSysinfoState)(nil)).Elem()
}

type switchControllerSnmpSysinfoArgs struct {
	// Contact information.
	ContactInfo *string `pulumi:"contactInfo"`
	// System description.
	Description *string `pulumi:"description"`
	// Local SNMP engine ID string (max 24 char).
	EngineId *string `pulumi:"engineId"`
	// System location.
	Location *string `pulumi:"location"`
	// Enable/disable SNMP. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerSnmpSysinfo resource.
type SwitchControllerSnmpSysinfoArgs struct {
	// Contact information.
	ContactInfo pulumi.StringPtrInput
	// System description.
	Description pulumi.StringPtrInput
	// Local SNMP engine ID string (max 24 char).
	EngineId pulumi.StringPtrInput
	// System location.
	Location pulumi.StringPtrInput
	// Enable/disable SNMP. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerSnmpSysinfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSnmpSysinfoArgs)(nil)).Elem()
}

type SwitchControllerSnmpSysinfoInput interface {
	pulumi.Input

	ToSwitchControllerSnmpSysinfoOutput() SwitchControllerSnmpSysinfoOutput
	ToSwitchControllerSnmpSysinfoOutputWithContext(ctx context.Context) SwitchControllerSnmpSysinfoOutput
}

func (*SwitchControllerSnmpSysinfo) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSnmpSysinfo)(nil)).Elem()
}

func (i *SwitchControllerSnmpSysinfo) ToSwitchControllerSnmpSysinfoOutput() SwitchControllerSnmpSysinfoOutput {
	return i.ToSwitchControllerSnmpSysinfoOutputWithContext(context.Background())
}

func (i *SwitchControllerSnmpSysinfo) ToSwitchControllerSnmpSysinfoOutputWithContext(ctx context.Context) SwitchControllerSnmpSysinfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSnmpSysinfoOutput)
}

// SwitchControllerSnmpSysinfoArrayInput is an input type that accepts SwitchControllerSnmpSysinfoArray and SwitchControllerSnmpSysinfoArrayOutput values.
// You can construct a concrete instance of `SwitchControllerSnmpSysinfoArrayInput` via:
//
//          SwitchControllerSnmpSysinfoArray{ SwitchControllerSnmpSysinfoArgs{...} }
type SwitchControllerSnmpSysinfoArrayInput interface {
	pulumi.Input

	ToSwitchControllerSnmpSysinfoArrayOutput() SwitchControllerSnmpSysinfoArrayOutput
	ToSwitchControllerSnmpSysinfoArrayOutputWithContext(context.Context) SwitchControllerSnmpSysinfoArrayOutput
}

type SwitchControllerSnmpSysinfoArray []SwitchControllerSnmpSysinfoInput

func (SwitchControllerSnmpSysinfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSnmpSysinfo)(nil)).Elem()
}

func (i SwitchControllerSnmpSysinfoArray) ToSwitchControllerSnmpSysinfoArrayOutput() SwitchControllerSnmpSysinfoArrayOutput {
	return i.ToSwitchControllerSnmpSysinfoArrayOutputWithContext(context.Background())
}

func (i SwitchControllerSnmpSysinfoArray) ToSwitchControllerSnmpSysinfoArrayOutputWithContext(ctx context.Context) SwitchControllerSnmpSysinfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSnmpSysinfoArrayOutput)
}

// SwitchControllerSnmpSysinfoMapInput is an input type that accepts SwitchControllerSnmpSysinfoMap and SwitchControllerSnmpSysinfoMapOutput values.
// You can construct a concrete instance of `SwitchControllerSnmpSysinfoMapInput` via:
//
//          SwitchControllerSnmpSysinfoMap{ "key": SwitchControllerSnmpSysinfoArgs{...} }
type SwitchControllerSnmpSysinfoMapInput interface {
	pulumi.Input

	ToSwitchControllerSnmpSysinfoMapOutput() SwitchControllerSnmpSysinfoMapOutput
	ToSwitchControllerSnmpSysinfoMapOutputWithContext(context.Context) SwitchControllerSnmpSysinfoMapOutput
}

type SwitchControllerSnmpSysinfoMap map[string]SwitchControllerSnmpSysinfoInput

func (SwitchControllerSnmpSysinfoMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSnmpSysinfo)(nil)).Elem()
}

func (i SwitchControllerSnmpSysinfoMap) ToSwitchControllerSnmpSysinfoMapOutput() SwitchControllerSnmpSysinfoMapOutput {
	return i.ToSwitchControllerSnmpSysinfoMapOutputWithContext(context.Background())
}

func (i SwitchControllerSnmpSysinfoMap) ToSwitchControllerSnmpSysinfoMapOutputWithContext(ctx context.Context) SwitchControllerSnmpSysinfoMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSnmpSysinfoMapOutput)
}

type SwitchControllerSnmpSysinfoOutput struct{ *pulumi.OutputState }

func (SwitchControllerSnmpSysinfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSnmpSysinfo)(nil)).Elem()
}

func (o SwitchControllerSnmpSysinfoOutput) ToSwitchControllerSnmpSysinfoOutput() SwitchControllerSnmpSysinfoOutput {
	return o
}

func (o SwitchControllerSnmpSysinfoOutput) ToSwitchControllerSnmpSysinfoOutputWithContext(ctx context.Context) SwitchControllerSnmpSysinfoOutput {
	return o
}

type SwitchControllerSnmpSysinfoArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerSnmpSysinfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSnmpSysinfo)(nil)).Elem()
}

func (o SwitchControllerSnmpSysinfoArrayOutput) ToSwitchControllerSnmpSysinfoArrayOutput() SwitchControllerSnmpSysinfoArrayOutput {
	return o
}

func (o SwitchControllerSnmpSysinfoArrayOutput) ToSwitchControllerSnmpSysinfoArrayOutputWithContext(ctx context.Context) SwitchControllerSnmpSysinfoArrayOutput {
	return o
}

func (o SwitchControllerSnmpSysinfoArrayOutput) Index(i pulumi.IntInput) SwitchControllerSnmpSysinfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerSnmpSysinfo {
		return vs[0].([]*SwitchControllerSnmpSysinfo)[vs[1].(int)]
	}).(SwitchControllerSnmpSysinfoOutput)
}

type SwitchControllerSnmpSysinfoMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerSnmpSysinfoMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSnmpSysinfo)(nil)).Elem()
}

func (o SwitchControllerSnmpSysinfoMapOutput) ToSwitchControllerSnmpSysinfoMapOutput() SwitchControllerSnmpSysinfoMapOutput {
	return o
}

func (o SwitchControllerSnmpSysinfoMapOutput) ToSwitchControllerSnmpSysinfoMapOutputWithContext(ctx context.Context) SwitchControllerSnmpSysinfoMapOutput {
	return o
}

func (o SwitchControllerSnmpSysinfoMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerSnmpSysinfoOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerSnmpSysinfo {
		return vs[0].(map[string]*SwitchControllerSnmpSysinfo)[vs[1].(string)]
	}).(SwitchControllerSnmpSysinfoOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSnmpSysinfoInput)(nil)).Elem(), &SwitchControllerSnmpSysinfo{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSnmpSysinfoArrayInput)(nil)).Elem(), SwitchControllerSnmpSysinfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSnmpSysinfoMapInput)(nil)).Elem(), SwitchControllerSnmpSysinfoMap{})
	pulumi.RegisterOutputType(SwitchControllerSnmpSysinfoOutput{})
	pulumi.RegisterOutputType(SwitchControllerSnmpSysinfoArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerSnmpSysinfoMapOutput{})
}
