// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure NetFlow per VDOM.
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
// 		_, err := fortios.NewSystemVdomNetflow(ctx, "trname", &fortios.SystemVdomNetflowArgs{
// 			CollectorIp:   pulumi.String("0.0.0.0"),
// 			CollectorPort: pulumi.Int(2055),
// 			SourceIp:      pulumi.String("0.0.0.0"),
// 			VdomNetflow:   pulumi.String("disable"),
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
// System VdomNetflow can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemVdomNetflow:SystemVdomNetflow labelname SystemVdomNetflow
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemVdomNetflow struct {
	pulumi.CustomResourceState

	// NetFlow collector IP address.
	CollectorIp pulumi.StringOutput `pulumi:"collectorIp"`
	// NetFlow collector port number.
	CollectorPort pulumi.IntOutput `pulumi:"collectorPort"`
	// Source IP address for communication with the NetFlow agent.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Enable/disable NetFlow per VDOM. Valid values: `enable`, `disable`.
	VdomNetflow pulumi.StringOutput `pulumi:"vdomNetflow"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemVdomNetflow registers a new resource with the given unique name, arguments, and options.
func NewSystemVdomNetflow(ctx *pulumi.Context,
	name string, args *SystemVdomNetflowArgs, opts ...pulumi.ResourceOption) (*SystemVdomNetflow, error) {
	if args == nil {
		args = &SystemVdomNetflowArgs{}
	}

	var resource SystemVdomNetflow
	err := ctx.RegisterResource("fortios:index/systemVdomNetflow:SystemVdomNetflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemVdomNetflow gets an existing SystemVdomNetflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemVdomNetflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemVdomNetflowState, opts ...pulumi.ResourceOption) (*SystemVdomNetflow, error) {
	var resource SystemVdomNetflow
	err := ctx.ReadResource("fortios:index/systemVdomNetflow:SystemVdomNetflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemVdomNetflow resources.
type systemVdomNetflowState struct {
	// NetFlow collector IP address.
	CollectorIp *string `pulumi:"collectorIp"`
	// NetFlow collector port number.
	CollectorPort *int `pulumi:"collectorPort"`
	// Source IP address for communication with the NetFlow agent.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable NetFlow per VDOM. Valid values: `enable`, `disable`.
	VdomNetflow *string `pulumi:"vdomNetflow"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemVdomNetflowState struct {
	// NetFlow collector IP address.
	CollectorIp pulumi.StringPtrInput
	// NetFlow collector port number.
	CollectorPort pulumi.IntPtrInput
	// Source IP address for communication with the NetFlow agent.
	SourceIp pulumi.StringPtrInput
	// Enable/disable NetFlow per VDOM. Valid values: `enable`, `disable`.
	VdomNetflow pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemVdomNetflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomNetflowState)(nil)).Elem()
}

type systemVdomNetflowArgs struct {
	// NetFlow collector IP address.
	CollectorIp *string `pulumi:"collectorIp"`
	// NetFlow collector port number.
	CollectorPort *int `pulumi:"collectorPort"`
	// Source IP address for communication with the NetFlow agent.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable NetFlow per VDOM. Valid values: `enable`, `disable`.
	VdomNetflow *string `pulumi:"vdomNetflow"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemVdomNetflow resource.
type SystemVdomNetflowArgs struct {
	// NetFlow collector IP address.
	CollectorIp pulumi.StringPtrInput
	// NetFlow collector port number.
	CollectorPort pulumi.IntPtrInput
	// Source IP address for communication with the NetFlow agent.
	SourceIp pulumi.StringPtrInput
	// Enable/disable NetFlow per VDOM. Valid values: `enable`, `disable`.
	VdomNetflow pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemVdomNetflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomNetflowArgs)(nil)).Elem()
}

type SystemVdomNetflowInput interface {
	pulumi.Input

	ToSystemVdomNetflowOutput() SystemVdomNetflowOutput
	ToSystemVdomNetflowOutputWithContext(ctx context.Context) SystemVdomNetflowOutput
}

func (*SystemVdomNetflow) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemVdomNetflow)(nil))
}

func (i *SystemVdomNetflow) ToSystemVdomNetflowOutput() SystemVdomNetflowOutput {
	return i.ToSystemVdomNetflowOutputWithContext(context.Background())
}

func (i *SystemVdomNetflow) ToSystemVdomNetflowOutputWithContext(ctx context.Context) SystemVdomNetflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomNetflowOutput)
}

func (i *SystemVdomNetflow) ToSystemVdomNetflowPtrOutput() SystemVdomNetflowPtrOutput {
	return i.ToSystemVdomNetflowPtrOutputWithContext(context.Background())
}

func (i *SystemVdomNetflow) ToSystemVdomNetflowPtrOutputWithContext(ctx context.Context) SystemVdomNetflowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomNetflowPtrOutput)
}

type SystemVdomNetflowPtrInput interface {
	pulumi.Input

	ToSystemVdomNetflowPtrOutput() SystemVdomNetflowPtrOutput
	ToSystemVdomNetflowPtrOutputWithContext(ctx context.Context) SystemVdomNetflowPtrOutput
}

type systemVdomNetflowPtrType SystemVdomNetflowArgs

func (*systemVdomNetflowPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomNetflow)(nil))
}

func (i *systemVdomNetflowPtrType) ToSystemVdomNetflowPtrOutput() SystemVdomNetflowPtrOutput {
	return i.ToSystemVdomNetflowPtrOutputWithContext(context.Background())
}

func (i *systemVdomNetflowPtrType) ToSystemVdomNetflowPtrOutputWithContext(ctx context.Context) SystemVdomNetflowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomNetflowPtrOutput)
}

// SystemVdomNetflowArrayInput is an input type that accepts SystemVdomNetflowArray and SystemVdomNetflowArrayOutput values.
// You can construct a concrete instance of `SystemVdomNetflowArrayInput` via:
//
//          SystemVdomNetflowArray{ SystemVdomNetflowArgs{...} }
type SystemVdomNetflowArrayInput interface {
	pulumi.Input

	ToSystemVdomNetflowArrayOutput() SystemVdomNetflowArrayOutput
	ToSystemVdomNetflowArrayOutputWithContext(context.Context) SystemVdomNetflowArrayOutput
}

type SystemVdomNetflowArray []SystemVdomNetflowInput

func (SystemVdomNetflowArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SystemVdomNetflow)(nil))
}

func (i SystemVdomNetflowArray) ToSystemVdomNetflowArrayOutput() SystemVdomNetflowArrayOutput {
	return i.ToSystemVdomNetflowArrayOutputWithContext(context.Background())
}

func (i SystemVdomNetflowArray) ToSystemVdomNetflowArrayOutputWithContext(ctx context.Context) SystemVdomNetflowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomNetflowArrayOutput)
}

// SystemVdomNetflowMapInput is an input type that accepts SystemVdomNetflowMap and SystemVdomNetflowMapOutput values.
// You can construct a concrete instance of `SystemVdomNetflowMapInput` via:
//
//          SystemVdomNetflowMap{ "key": SystemVdomNetflowArgs{...} }
type SystemVdomNetflowMapInput interface {
	pulumi.Input

	ToSystemVdomNetflowMapOutput() SystemVdomNetflowMapOutput
	ToSystemVdomNetflowMapOutputWithContext(context.Context) SystemVdomNetflowMapOutput
}

type SystemVdomNetflowMap map[string]SystemVdomNetflowInput

func (SystemVdomNetflowMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SystemVdomNetflow)(nil))
}

func (i SystemVdomNetflowMap) ToSystemVdomNetflowMapOutput() SystemVdomNetflowMapOutput {
	return i.ToSystemVdomNetflowMapOutputWithContext(context.Background())
}

func (i SystemVdomNetflowMap) ToSystemVdomNetflowMapOutputWithContext(ctx context.Context) SystemVdomNetflowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomNetflowMapOutput)
}

type SystemVdomNetflowOutput struct {
	*pulumi.OutputState
}

func (SystemVdomNetflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemVdomNetflow)(nil))
}

func (o SystemVdomNetflowOutput) ToSystemVdomNetflowOutput() SystemVdomNetflowOutput {
	return o
}

func (o SystemVdomNetflowOutput) ToSystemVdomNetflowOutputWithContext(ctx context.Context) SystemVdomNetflowOutput {
	return o
}

func (o SystemVdomNetflowOutput) ToSystemVdomNetflowPtrOutput() SystemVdomNetflowPtrOutput {
	return o.ToSystemVdomNetflowPtrOutputWithContext(context.Background())
}

func (o SystemVdomNetflowOutput) ToSystemVdomNetflowPtrOutputWithContext(ctx context.Context) SystemVdomNetflowPtrOutput {
	return o.ApplyT(func(v SystemVdomNetflow) *SystemVdomNetflow {
		return &v
	}).(SystemVdomNetflowPtrOutput)
}

type SystemVdomNetflowPtrOutput struct {
	*pulumi.OutputState
}

func (SystemVdomNetflowPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomNetflow)(nil))
}

func (o SystemVdomNetflowPtrOutput) ToSystemVdomNetflowPtrOutput() SystemVdomNetflowPtrOutput {
	return o
}

func (o SystemVdomNetflowPtrOutput) ToSystemVdomNetflowPtrOutputWithContext(ctx context.Context) SystemVdomNetflowPtrOutput {
	return o
}

type SystemVdomNetflowArrayOutput struct{ *pulumi.OutputState }

func (SystemVdomNetflowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemVdomNetflow)(nil))
}

func (o SystemVdomNetflowArrayOutput) ToSystemVdomNetflowArrayOutput() SystemVdomNetflowArrayOutput {
	return o
}

func (o SystemVdomNetflowArrayOutput) ToSystemVdomNetflowArrayOutputWithContext(ctx context.Context) SystemVdomNetflowArrayOutput {
	return o
}

func (o SystemVdomNetflowArrayOutput) Index(i pulumi.IntInput) SystemVdomNetflowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemVdomNetflow {
		return vs[0].([]SystemVdomNetflow)[vs[1].(int)]
	}).(SystemVdomNetflowOutput)
}

type SystemVdomNetflowMapOutput struct{ *pulumi.OutputState }

func (SystemVdomNetflowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SystemVdomNetflow)(nil))
}

func (o SystemVdomNetflowMapOutput) ToSystemVdomNetflowMapOutput() SystemVdomNetflowMapOutput {
	return o
}

func (o SystemVdomNetflowMapOutput) ToSystemVdomNetflowMapOutputWithContext(ctx context.Context) SystemVdomNetflowMapOutput {
	return o
}

func (o SystemVdomNetflowMapOutput) MapIndex(k pulumi.StringInput) SystemVdomNetflowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SystemVdomNetflow {
		return vs[0].(map[string]SystemVdomNetflow)[vs[1].(string)]
	}).(SystemVdomNetflowOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemVdomNetflowOutput{})
	pulumi.RegisterOutputType(SystemVdomNetflowPtrOutput{})
	pulumi.RegisterOutputType(SystemVdomNetflowArrayOutput{})
	pulumi.RegisterOutputType(SystemVdomNetflowMapOutput{})
}