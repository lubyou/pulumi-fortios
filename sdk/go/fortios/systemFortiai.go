// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiAI. Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// System Fortiai can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemFortiai:SystemFortiai labelname SystemFortiai
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemFortiai struct {
	pulumi.CustomResourceState

	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// Source IP address for communications to FortiAI.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Enable/disable FortiAI. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemFortiai registers a new resource with the given unique name, arguments, and options.
func NewSystemFortiai(ctx *pulumi.Context,
	name string, args *SystemFortiaiArgs, opts ...pulumi.ResourceOption) (*SystemFortiai, error) {
	if args == nil {
		args = &SystemFortiaiArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemFortiai
	err := ctx.RegisterResource("fortios:index/systemFortiai:SystemFortiai", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemFortiai gets an existing SystemFortiai resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemFortiai(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemFortiaiState, opts ...pulumi.ResourceOption) (*SystemFortiai, error) {
	var resource SystemFortiai
	err := ctx.ReadResource("fortios:index/systemFortiai:SystemFortiai", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemFortiai resources.
type systemFortiaiState struct {
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Source IP address for communications to FortiAI.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable FortiAI. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemFortiaiState struct {
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Source IP address for communications to FortiAI.
	SourceIp pulumi.StringPtrInput
	// Enable/disable FortiAI. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemFortiaiState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFortiaiState)(nil)).Elem()
}

type systemFortiaiArgs struct {
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Source IP address for communications to FortiAI.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable FortiAI. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemFortiai resource.
type SystemFortiaiArgs struct {
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Source IP address for communications to FortiAI.
	SourceIp pulumi.StringPtrInput
	// Enable/disable FortiAI. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemFortiaiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFortiaiArgs)(nil)).Elem()
}

type SystemFortiaiInput interface {
	pulumi.Input

	ToSystemFortiaiOutput() SystemFortiaiOutput
	ToSystemFortiaiOutputWithContext(ctx context.Context) SystemFortiaiOutput
}

func (*SystemFortiai) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFortiai)(nil)).Elem()
}

func (i *SystemFortiai) ToSystemFortiaiOutput() SystemFortiaiOutput {
	return i.ToSystemFortiaiOutputWithContext(context.Background())
}

func (i *SystemFortiai) ToSystemFortiaiOutputWithContext(ctx context.Context) SystemFortiaiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFortiaiOutput)
}

// SystemFortiaiArrayInput is an input type that accepts SystemFortiaiArray and SystemFortiaiArrayOutput values.
// You can construct a concrete instance of `SystemFortiaiArrayInput` via:
//
//          SystemFortiaiArray{ SystemFortiaiArgs{...} }
type SystemFortiaiArrayInput interface {
	pulumi.Input

	ToSystemFortiaiArrayOutput() SystemFortiaiArrayOutput
	ToSystemFortiaiArrayOutputWithContext(context.Context) SystemFortiaiArrayOutput
}

type SystemFortiaiArray []SystemFortiaiInput

func (SystemFortiaiArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFortiai)(nil)).Elem()
}

func (i SystemFortiaiArray) ToSystemFortiaiArrayOutput() SystemFortiaiArrayOutput {
	return i.ToSystemFortiaiArrayOutputWithContext(context.Background())
}

func (i SystemFortiaiArray) ToSystemFortiaiArrayOutputWithContext(ctx context.Context) SystemFortiaiArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFortiaiArrayOutput)
}

// SystemFortiaiMapInput is an input type that accepts SystemFortiaiMap and SystemFortiaiMapOutput values.
// You can construct a concrete instance of `SystemFortiaiMapInput` via:
//
//          SystemFortiaiMap{ "key": SystemFortiaiArgs{...} }
type SystemFortiaiMapInput interface {
	pulumi.Input

	ToSystemFortiaiMapOutput() SystemFortiaiMapOutput
	ToSystemFortiaiMapOutputWithContext(context.Context) SystemFortiaiMapOutput
}

type SystemFortiaiMap map[string]SystemFortiaiInput

func (SystemFortiaiMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFortiai)(nil)).Elem()
}

func (i SystemFortiaiMap) ToSystemFortiaiMapOutput() SystemFortiaiMapOutput {
	return i.ToSystemFortiaiMapOutputWithContext(context.Background())
}

func (i SystemFortiaiMap) ToSystemFortiaiMapOutputWithContext(ctx context.Context) SystemFortiaiMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFortiaiMapOutput)
}

type SystemFortiaiOutput struct{ *pulumi.OutputState }

func (SystemFortiaiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFortiai)(nil)).Elem()
}

func (o SystemFortiaiOutput) ToSystemFortiaiOutput() SystemFortiaiOutput {
	return o
}

func (o SystemFortiaiOutput) ToSystemFortiaiOutputWithContext(ctx context.Context) SystemFortiaiOutput {
	return o
}

type SystemFortiaiArrayOutput struct{ *pulumi.OutputState }

func (SystemFortiaiArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFortiai)(nil)).Elem()
}

func (o SystemFortiaiArrayOutput) ToSystemFortiaiArrayOutput() SystemFortiaiArrayOutput {
	return o
}

func (o SystemFortiaiArrayOutput) ToSystemFortiaiArrayOutputWithContext(ctx context.Context) SystemFortiaiArrayOutput {
	return o
}

func (o SystemFortiaiArrayOutput) Index(i pulumi.IntInput) SystemFortiaiOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemFortiai {
		return vs[0].([]*SystemFortiai)[vs[1].(int)]
	}).(SystemFortiaiOutput)
}

type SystemFortiaiMapOutput struct{ *pulumi.OutputState }

func (SystemFortiaiMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFortiai)(nil)).Elem()
}

func (o SystemFortiaiMapOutput) ToSystemFortiaiMapOutput() SystemFortiaiMapOutput {
	return o
}

func (o SystemFortiaiMapOutput) ToSystemFortiaiMapOutputWithContext(ctx context.Context) SystemFortiaiMapOutput {
	return o
}

func (o SystemFortiaiMapOutput) MapIndex(k pulumi.StringInput) SystemFortiaiOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemFortiai {
		return vs[0].(map[string]*SystemFortiai)[vs[1].(string)]
	}).(SystemFortiaiOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFortiaiInput)(nil)).Elem(), &SystemFortiai{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFortiaiArrayInput)(nil)).Elem(), SystemFortiaiArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFortiaiMapInput)(nil)).Elem(), SystemFortiaiMap{})
	pulumi.RegisterOutputType(SystemFortiaiOutput{})
	pulumi.RegisterOutputType(SystemFortiaiArrayOutput{})
	pulumi.RegisterOutputType(SystemFortiaiMapOutput{})
}
