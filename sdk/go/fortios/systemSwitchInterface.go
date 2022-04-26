// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure software switch interfaces by grouping physical and WiFi interfaces.
//
// ## Import
//
// System SwitchInterface can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemSwitchInterface:SystemSwitchInterface labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemSwitchInterface:SystemSwitchInterface labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemSwitchInterface struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
	IntraSwitchPolicy pulumi.StringOutput `pulumi:"intraSwitchPolicy"`
	// Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
	MacTtl pulumi.IntOutput `pulumi:"macTtl"`
	// Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
	Members SystemSwitchInterfaceMemberArrayOutput `pulumi:"members"`
	// Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
	Span pulumi.StringOutput `pulumi:"span"`
	// SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
	SpanDestPort pulumi.StringOutput `pulumi:"spanDestPort"`
	// The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
	SpanDirection pulumi.StringOutput `pulumi:"spanDirection"`
	// Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `spanSourcePort` block is documented below.
	SpanSourcePorts SystemSwitchInterfaceSpanSourcePortArrayOutput `pulumi:"spanSourcePorts"`
	// Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
	Type pulumi.StringOutput `pulumi:"type"`
	// VDOM that the software switch belongs to.
	Vdom pulumi.StringOutput `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemSwitchInterface registers a new resource with the given unique name, arguments, and options.
func NewSystemSwitchInterface(ctx *pulumi.Context,
	name string, args *SystemSwitchInterfaceArgs, opts ...pulumi.ResourceOption) (*SystemSwitchInterface, error) {
	if args == nil {
		args = &SystemSwitchInterfaceArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemSwitchInterface
	err := ctx.RegisterResource("fortios:index/systemSwitchInterface:SystemSwitchInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSwitchInterface gets an existing SystemSwitchInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSwitchInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSwitchInterfaceState, opts ...pulumi.ResourceOption) (*SystemSwitchInterface, error) {
	var resource SystemSwitchInterface
	err := ctx.ReadResource("fortios:index/systemSwitchInterface:SystemSwitchInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSwitchInterface resources.
type systemSwitchInterfaceState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
	IntraSwitchPolicy *string `pulumi:"intraSwitchPolicy"`
	// Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
	MacTtl *int `pulumi:"macTtl"`
	// Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
	Members []SystemSwitchInterfaceMember `pulumi:"members"`
	// Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
	Name *string `pulumi:"name"`
	// Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
	Span *string `pulumi:"span"`
	// SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
	SpanDestPort *string `pulumi:"spanDestPort"`
	// The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
	SpanDirection *string `pulumi:"spanDirection"`
	// Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `spanSourcePort` block is documented below.
	SpanSourcePorts []SystemSwitchInterfaceSpanSourcePort `pulumi:"spanSourcePorts"`
	// Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
	Type *string `pulumi:"type"`
	// VDOM that the software switch belongs to.
	Vdom *string `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemSwitchInterfaceState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
	IntraSwitchPolicy pulumi.StringPtrInput
	// Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
	MacTtl pulumi.IntPtrInput
	// Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
	Members SystemSwitchInterfaceMemberArrayInput
	// Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
	Name pulumi.StringPtrInput
	// Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
	Span pulumi.StringPtrInput
	// SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
	SpanDestPort pulumi.StringPtrInput
	// The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
	SpanDirection pulumi.StringPtrInput
	// Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `spanSourcePort` block is documented below.
	SpanSourcePorts SystemSwitchInterfaceSpanSourcePortArrayInput
	// Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
	Type pulumi.StringPtrInput
	// VDOM that the software switch belongs to.
	Vdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSwitchInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSwitchInterfaceState)(nil)).Elem()
}

type systemSwitchInterfaceArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
	IntraSwitchPolicy *string `pulumi:"intraSwitchPolicy"`
	// Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
	MacTtl *int `pulumi:"macTtl"`
	// Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
	Members []SystemSwitchInterfaceMember `pulumi:"members"`
	// Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
	Name *string `pulumi:"name"`
	// Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
	Span *string `pulumi:"span"`
	// SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
	SpanDestPort *string `pulumi:"spanDestPort"`
	// The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
	SpanDirection *string `pulumi:"spanDirection"`
	// Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `spanSourcePort` block is documented below.
	SpanSourcePorts []SystemSwitchInterfaceSpanSourcePort `pulumi:"spanSourcePorts"`
	// Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
	Type *string `pulumi:"type"`
	// VDOM that the software switch belongs to.
	Vdom *string `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemSwitchInterface resource.
type SystemSwitchInterfaceArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
	IntraSwitchPolicy pulumi.StringPtrInput
	// Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
	MacTtl pulumi.IntPtrInput
	// Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
	Members SystemSwitchInterfaceMemberArrayInput
	// Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
	Name pulumi.StringPtrInput
	// Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
	Span pulumi.StringPtrInput
	// SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
	SpanDestPort pulumi.StringPtrInput
	// The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
	SpanDirection pulumi.StringPtrInput
	// Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `spanSourcePort` block is documented below.
	SpanSourcePorts SystemSwitchInterfaceSpanSourcePortArrayInput
	// Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
	Type pulumi.StringPtrInput
	// VDOM that the software switch belongs to.
	Vdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSwitchInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSwitchInterfaceArgs)(nil)).Elem()
}

type SystemSwitchInterfaceInput interface {
	pulumi.Input

	ToSystemSwitchInterfaceOutput() SystemSwitchInterfaceOutput
	ToSystemSwitchInterfaceOutputWithContext(ctx context.Context) SystemSwitchInterfaceOutput
}

func (*SystemSwitchInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSwitchInterface)(nil)).Elem()
}

func (i *SystemSwitchInterface) ToSystemSwitchInterfaceOutput() SystemSwitchInterfaceOutput {
	return i.ToSystemSwitchInterfaceOutputWithContext(context.Background())
}

func (i *SystemSwitchInterface) ToSystemSwitchInterfaceOutputWithContext(ctx context.Context) SystemSwitchInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSwitchInterfaceOutput)
}

// SystemSwitchInterfaceArrayInput is an input type that accepts SystemSwitchInterfaceArray and SystemSwitchInterfaceArrayOutput values.
// You can construct a concrete instance of `SystemSwitchInterfaceArrayInput` via:
//
//          SystemSwitchInterfaceArray{ SystemSwitchInterfaceArgs{...} }
type SystemSwitchInterfaceArrayInput interface {
	pulumi.Input

	ToSystemSwitchInterfaceArrayOutput() SystemSwitchInterfaceArrayOutput
	ToSystemSwitchInterfaceArrayOutputWithContext(context.Context) SystemSwitchInterfaceArrayOutput
}

type SystemSwitchInterfaceArray []SystemSwitchInterfaceInput

func (SystemSwitchInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSwitchInterface)(nil)).Elem()
}

func (i SystemSwitchInterfaceArray) ToSystemSwitchInterfaceArrayOutput() SystemSwitchInterfaceArrayOutput {
	return i.ToSystemSwitchInterfaceArrayOutputWithContext(context.Background())
}

func (i SystemSwitchInterfaceArray) ToSystemSwitchInterfaceArrayOutputWithContext(ctx context.Context) SystemSwitchInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSwitchInterfaceArrayOutput)
}

// SystemSwitchInterfaceMapInput is an input type that accepts SystemSwitchInterfaceMap and SystemSwitchInterfaceMapOutput values.
// You can construct a concrete instance of `SystemSwitchInterfaceMapInput` via:
//
//          SystemSwitchInterfaceMap{ "key": SystemSwitchInterfaceArgs{...} }
type SystemSwitchInterfaceMapInput interface {
	pulumi.Input

	ToSystemSwitchInterfaceMapOutput() SystemSwitchInterfaceMapOutput
	ToSystemSwitchInterfaceMapOutputWithContext(context.Context) SystemSwitchInterfaceMapOutput
}

type SystemSwitchInterfaceMap map[string]SystemSwitchInterfaceInput

func (SystemSwitchInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSwitchInterface)(nil)).Elem()
}

func (i SystemSwitchInterfaceMap) ToSystemSwitchInterfaceMapOutput() SystemSwitchInterfaceMapOutput {
	return i.ToSystemSwitchInterfaceMapOutputWithContext(context.Background())
}

func (i SystemSwitchInterfaceMap) ToSystemSwitchInterfaceMapOutputWithContext(ctx context.Context) SystemSwitchInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSwitchInterfaceMapOutput)
}

type SystemSwitchInterfaceOutput struct{ *pulumi.OutputState }

func (SystemSwitchInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSwitchInterface)(nil)).Elem()
}

func (o SystemSwitchInterfaceOutput) ToSystemSwitchInterfaceOutput() SystemSwitchInterfaceOutput {
	return o
}

func (o SystemSwitchInterfaceOutput) ToSystemSwitchInterfaceOutputWithContext(ctx context.Context) SystemSwitchInterfaceOutput {
	return o
}

type SystemSwitchInterfaceArrayOutput struct{ *pulumi.OutputState }

func (SystemSwitchInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSwitchInterface)(nil)).Elem()
}

func (o SystemSwitchInterfaceArrayOutput) ToSystemSwitchInterfaceArrayOutput() SystemSwitchInterfaceArrayOutput {
	return o
}

func (o SystemSwitchInterfaceArrayOutput) ToSystemSwitchInterfaceArrayOutputWithContext(ctx context.Context) SystemSwitchInterfaceArrayOutput {
	return o
}

func (o SystemSwitchInterfaceArrayOutput) Index(i pulumi.IntInput) SystemSwitchInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSwitchInterface {
		return vs[0].([]*SystemSwitchInterface)[vs[1].(int)]
	}).(SystemSwitchInterfaceOutput)
}

type SystemSwitchInterfaceMapOutput struct{ *pulumi.OutputState }

func (SystemSwitchInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSwitchInterface)(nil)).Elem()
}

func (o SystemSwitchInterfaceMapOutput) ToSystemSwitchInterfaceMapOutput() SystemSwitchInterfaceMapOutput {
	return o
}

func (o SystemSwitchInterfaceMapOutput) ToSystemSwitchInterfaceMapOutputWithContext(ctx context.Context) SystemSwitchInterfaceMapOutput {
	return o
}

func (o SystemSwitchInterfaceMapOutput) MapIndex(k pulumi.StringInput) SystemSwitchInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSwitchInterface {
		return vs[0].(map[string]*SystemSwitchInterface)[vs[1].(string)]
	}).(SystemSwitchInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSwitchInterfaceInput)(nil)).Elem(), &SystemSwitchInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSwitchInterfaceArrayInput)(nil)).Elem(), SystemSwitchInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSwitchInterfaceMapInput)(nil)).Elem(), SystemSwitchInterfaceMap{})
	pulumi.RegisterOutputType(SystemSwitchInterfaceOutput{})
	pulumi.RegisterOutputType(SystemSwitchInterfaceArrayOutput{})
	pulumi.RegisterOutputType(SystemSwitchInterfaceMapOutput{})
}
