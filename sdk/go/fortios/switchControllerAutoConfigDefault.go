// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure default auto-config QoS policy for FortiSwitch.
//
// ## Import
//
// SwitchControllerAutoConfig Default can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/switchControllerAutoConfigDefault:SwitchControllerAutoConfigDefault labelname SwitchControllerAutoConfigDefault
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/switchControllerAutoConfigDefault:SwitchControllerAutoConfigDefault labelname SwitchControllerAutoConfigDefault
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SwitchControllerAutoConfigDefault struct {
	pulumi.CustomResourceState

	// Default FortiLink auto-config policy.
	FgtPolicy pulumi.StringOutput `pulumi:"fgtPolicy"`
	// Default ICL auto-config policy.
	IclPolicy pulumi.StringOutput `pulumi:"iclPolicy"`
	// Default ISL auto-config policy.
	IslPolicy pulumi.StringOutput `pulumi:"islPolicy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerAutoConfigDefault registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerAutoConfigDefault(ctx *pulumi.Context,
	name string, args *SwitchControllerAutoConfigDefaultArgs, opts ...pulumi.ResourceOption) (*SwitchControllerAutoConfigDefault, error) {
	if args == nil {
		args = &SwitchControllerAutoConfigDefaultArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchControllerAutoConfigDefault
	err := ctx.RegisterResource("fortios:index/switchControllerAutoConfigDefault:SwitchControllerAutoConfigDefault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerAutoConfigDefault gets an existing SwitchControllerAutoConfigDefault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerAutoConfigDefault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerAutoConfigDefaultState, opts ...pulumi.ResourceOption) (*SwitchControllerAutoConfigDefault, error) {
	var resource SwitchControllerAutoConfigDefault
	err := ctx.ReadResource("fortios:index/switchControllerAutoConfigDefault:SwitchControllerAutoConfigDefault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerAutoConfigDefault resources.
type switchControllerAutoConfigDefaultState struct {
	// Default FortiLink auto-config policy.
	FgtPolicy *string `pulumi:"fgtPolicy"`
	// Default ICL auto-config policy.
	IclPolicy *string `pulumi:"iclPolicy"`
	// Default ISL auto-config policy.
	IslPolicy *string `pulumi:"islPolicy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchControllerAutoConfigDefaultState struct {
	// Default FortiLink auto-config policy.
	FgtPolicy pulumi.StringPtrInput
	// Default ICL auto-config policy.
	IclPolicy pulumi.StringPtrInput
	// Default ISL auto-config policy.
	IslPolicy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerAutoConfigDefaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerAutoConfigDefaultState)(nil)).Elem()
}

type switchControllerAutoConfigDefaultArgs struct {
	// Default FortiLink auto-config policy.
	FgtPolicy *string `pulumi:"fgtPolicy"`
	// Default ICL auto-config policy.
	IclPolicy *string `pulumi:"iclPolicy"`
	// Default ISL auto-config policy.
	IslPolicy *string `pulumi:"islPolicy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerAutoConfigDefault resource.
type SwitchControllerAutoConfigDefaultArgs struct {
	// Default FortiLink auto-config policy.
	FgtPolicy pulumi.StringPtrInput
	// Default ICL auto-config policy.
	IclPolicy pulumi.StringPtrInput
	// Default ISL auto-config policy.
	IslPolicy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerAutoConfigDefaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerAutoConfigDefaultArgs)(nil)).Elem()
}

type SwitchControllerAutoConfigDefaultInput interface {
	pulumi.Input

	ToSwitchControllerAutoConfigDefaultOutput() SwitchControllerAutoConfigDefaultOutput
	ToSwitchControllerAutoConfigDefaultOutputWithContext(ctx context.Context) SwitchControllerAutoConfigDefaultOutput
}

func (*SwitchControllerAutoConfigDefault) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerAutoConfigDefault)(nil)).Elem()
}

func (i *SwitchControllerAutoConfigDefault) ToSwitchControllerAutoConfigDefaultOutput() SwitchControllerAutoConfigDefaultOutput {
	return i.ToSwitchControllerAutoConfigDefaultOutputWithContext(context.Background())
}

func (i *SwitchControllerAutoConfigDefault) ToSwitchControllerAutoConfigDefaultOutputWithContext(ctx context.Context) SwitchControllerAutoConfigDefaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerAutoConfigDefaultOutput)
}

// SwitchControllerAutoConfigDefaultArrayInput is an input type that accepts SwitchControllerAutoConfigDefaultArray and SwitchControllerAutoConfigDefaultArrayOutput values.
// You can construct a concrete instance of `SwitchControllerAutoConfigDefaultArrayInput` via:
//
//          SwitchControllerAutoConfigDefaultArray{ SwitchControllerAutoConfigDefaultArgs{...} }
type SwitchControllerAutoConfigDefaultArrayInput interface {
	pulumi.Input

	ToSwitchControllerAutoConfigDefaultArrayOutput() SwitchControllerAutoConfigDefaultArrayOutput
	ToSwitchControllerAutoConfigDefaultArrayOutputWithContext(context.Context) SwitchControllerAutoConfigDefaultArrayOutput
}

type SwitchControllerAutoConfigDefaultArray []SwitchControllerAutoConfigDefaultInput

func (SwitchControllerAutoConfigDefaultArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerAutoConfigDefault)(nil)).Elem()
}

func (i SwitchControllerAutoConfigDefaultArray) ToSwitchControllerAutoConfigDefaultArrayOutput() SwitchControllerAutoConfigDefaultArrayOutput {
	return i.ToSwitchControllerAutoConfigDefaultArrayOutputWithContext(context.Background())
}

func (i SwitchControllerAutoConfigDefaultArray) ToSwitchControllerAutoConfigDefaultArrayOutputWithContext(ctx context.Context) SwitchControllerAutoConfigDefaultArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerAutoConfigDefaultArrayOutput)
}

// SwitchControllerAutoConfigDefaultMapInput is an input type that accepts SwitchControllerAutoConfigDefaultMap and SwitchControllerAutoConfigDefaultMapOutput values.
// You can construct a concrete instance of `SwitchControllerAutoConfigDefaultMapInput` via:
//
//          SwitchControllerAutoConfigDefaultMap{ "key": SwitchControllerAutoConfigDefaultArgs{...} }
type SwitchControllerAutoConfigDefaultMapInput interface {
	pulumi.Input

	ToSwitchControllerAutoConfigDefaultMapOutput() SwitchControllerAutoConfigDefaultMapOutput
	ToSwitchControllerAutoConfigDefaultMapOutputWithContext(context.Context) SwitchControllerAutoConfigDefaultMapOutput
}

type SwitchControllerAutoConfigDefaultMap map[string]SwitchControllerAutoConfigDefaultInput

func (SwitchControllerAutoConfigDefaultMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerAutoConfigDefault)(nil)).Elem()
}

func (i SwitchControllerAutoConfigDefaultMap) ToSwitchControllerAutoConfigDefaultMapOutput() SwitchControllerAutoConfigDefaultMapOutput {
	return i.ToSwitchControllerAutoConfigDefaultMapOutputWithContext(context.Background())
}

func (i SwitchControllerAutoConfigDefaultMap) ToSwitchControllerAutoConfigDefaultMapOutputWithContext(ctx context.Context) SwitchControllerAutoConfigDefaultMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerAutoConfigDefaultMapOutput)
}

type SwitchControllerAutoConfigDefaultOutput struct{ *pulumi.OutputState }

func (SwitchControllerAutoConfigDefaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerAutoConfigDefault)(nil)).Elem()
}

func (o SwitchControllerAutoConfigDefaultOutput) ToSwitchControllerAutoConfigDefaultOutput() SwitchControllerAutoConfigDefaultOutput {
	return o
}

func (o SwitchControllerAutoConfigDefaultOutput) ToSwitchControllerAutoConfigDefaultOutputWithContext(ctx context.Context) SwitchControllerAutoConfigDefaultOutput {
	return o
}

type SwitchControllerAutoConfigDefaultArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerAutoConfigDefaultArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerAutoConfigDefault)(nil)).Elem()
}

func (o SwitchControllerAutoConfigDefaultArrayOutput) ToSwitchControllerAutoConfigDefaultArrayOutput() SwitchControllerAutoConfigDefaultArrayOutput {
	return o
}

func (o SwitchControllerAutoConfigDefaultArrayOutput) ToSwitchControllerAutoConfigDefaultArrayOutputWithContext(ctx context.Context) SwitchControllerAutoConfigDefaultArrayOutput {
	return o
}

func (o SwitchControllerAutoConfigDefaultArrayOutput) Index(i pulumi.IntInput) SwitchControllerAutoConfigDefaultOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerAutoConfigDefault {
		return vs[0].([]*SwitchControllerAutoConfigDefault)[vs[1].(int)]
	}).(SwitchControllerAutoConfigDefaultOutput)
}

type SwitchControllerAutoConfigDefaultMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerAutoConfigDefaultMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerAutoConfigDefault)(nil)).Elem()
}

func (o SwitchControllerAutoConfigDefaultMapOutput) ToSwitchControllerAutoConfigDefaultMapOutput() SwitchControllerAutoConfigDefaultMapOutput {
	return o
}

func (o SwitchControllerAutoConfigDefaultMapOutput) ToSwitchControllerAutoConfigDefaultMapOutputWithContext(ctx context.Context) SwitchControllerAutoConfigDefaultMapOutput {
	return o
}

func (o SwitchControllerAutoConfigDefaultMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerAutoConfigDefaultOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerAutoConfigDefault {
		return vs[0].(map[string]*SwitchControllerAutoConfigDefault)[vs[1].(string)]
	}).(SwitchControllerAutoConfigDefaultOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerAutoConfigDefaultInput)(nil)).Elem(), &SwitchControllerAutoConfigDefault{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerAutoConfigDefaultArrayInput)(nil)).Elem(), SwitchControllerAutoConfigDefaultArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerAutoConfigDefaultMapInput)(nil)).Elem(), SwitchControllerAutoConfigDefaultMap{})
	pulumi.RegisterOutputType(SwitchControllerAutoConfigDefaultOutput{})
	pulumi.RegisterOutputType(SwitchControllerAutoConfigDefaultArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerAutoConfigDefaultMapOutput{})
}
