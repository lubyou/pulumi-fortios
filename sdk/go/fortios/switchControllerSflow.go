// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch sFlow.
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
// 		_, err := fortios.NewSwitchControllerSflow(ctx, "trname", &fortios.SwitchControllerSflowArgs{
// 			CollectorIp:   pulumi.String("0.0.0.0"),
// 			CollectorPort: pulumi.Int(6343),
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
// SwitchController Sflow can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/switchControllerSflow:SwitchControllerSflow labelname SwitchControllerSflow
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SwitchControllerSflow struct {
	pulumi.CustomResourceState

	// Collector IP.
	CollectorIp pulumi.StringOutput `pulumi:"collectorIp"`
	// SFlow collector port (0 - 65535).
	CollectorPort pulumi.IntOutput `pulumi:"collectorPort"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerSflow registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerSflow(ctx *pulumi.Context,
	name string, args *SwitchControllerSflowArgs, opts ...pulumi.ResourceOption) (*SwitchControllerSflow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CollectorIp == nil {
		return nil, errors.New("invalid value for required argument 'CollectorIp'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchControllerSflow
	err := ctx.RegisterResource("fortios:index/switchControllerSflow:SwitchControllerSflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerSflow gets an existing SwitchControllerSflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerSflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerSflowState, opts ...pulumi.ResourceOption) (*SwitchControllerSflow, error) {
	var resource SwitchControllerSflow
	err := ctx.ReadResource("fortios:index/switchControllerSflow:SwitchControllerSflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerSflow resources.
type switchControllerSflowState struct {
	// Collector IP.
	CollectorIp *string `pulumi:"collectorIp"`
	// SFlow collector port (0 - 65535).
	CollectorPort *int `pulumi:"collectorPort"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchControllerSflowState struct {
	// Collector IP.
	CollectorIp pulumi.StringPtrInput
	// SFlow collector port (0 - 65535).
	CollectorPort pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerSflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSflowState)(nil)).Elem()
}

type switchControllerSflowArgs struct {
	// Collector IP.
	CollectorIp string `pulumi:"collectorIp"`
	// SFlow collector port (0 - 65535).
	CollectorPort *int `pulumi:"collectorPort"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerSflow resource.
type SwitchControllerSflowArgs struct {
	// Collector IP.
	CollectorIp pulumi.StringInput
	// SFlow collector port (0 - 65535).
	CollectorPort pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerSflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSflowArgs)(nil)).Elem()
}

type SwitchControllerSflowInput interface {
	pulumi.Input

	ToSwitchControllerSflowOutput() SwitchControllerSflowOutput
	ToSwitchControllerSflowOutputWithContext(ctx context.Context) SwitchControllerSflowOutput
}

func (*SwitchControllerSflow) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSflow)(nil)).Elem()
}

func (i *SwitchControllerSflow) ToSwitchControllerSflowOutput() SwitchControllerSflowOutput {
	return i.ToSwitchControllerSflowOutputWithContext(context.Background())
}

func (i *SwitchControllerSflow) ToSwitchControllerSflowOutputWithContext(ctx context.Context) SwitchControllerSflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSflowOutput)
}

// SwitchControllerSflowArrayInput is an input type that accepts SwitchControllerSflowArray and SwitchControllerSflowArrayOutput values.
// You can construct a concrete instance of `SwitchControllerSflowArrayInput` via:
//
//          SwitchControllerSflowArray{ SwitchControllerSflowArgs{...} }
type SwitchControllerSflowArrayInput interface {
	pulumi.Input

	ToSwitchControllerSflowArrayOutput() SwitchControllerSflowArrayOutput
	ToSwitchControllerSflowArrayOutputWithContext(context.Context) SwitchControllerSflowArrayOutput
}

type SwitchControllerSflowArray []SwitchControllerSflowInput

func (SwitchControllerSflowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSflow)(nil)).Elem()
}

func (i SwitchControllerSflowArray) ToSwitchControllerSflowArrayOutput() SwitchControllerSflowArrayOutput {
	return i.ToSwitchControllerSflowArrayOutputWithContext(context.Background())
}

func (i SwitchControllerSflowArray) ToSwitchControllerSflowArrayOutputWithContext(ctx context.Context) SwitchControllerSflowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSflowArrayOutput)
}

// SwitchControllerSflowMapInput is an input type that accepts SwitchControllerSflowMap and SwitchControllerSflowMapOutput values.
// You can construct a concrete instance of `SwitchControllerSflowMapInput` via:
//
//          SwitchControllerSflowMap{ "key": SwitchControllerSflowArgs{...} }
type SwitchControllerSflowMapInput interface {
	pulumi.Input

	ToSwitchControllerSflowMapOutput() SwitchControllerSflowMapOutput
	ToSwitchControllerSflowMapOutputWithContext(context.Context) SwitchControllerSflowMapOutput
}

type SwitchControllerSflowMap map[string]SwitchControllerSflowInput

func (SwitchControllerSflowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSflow)(nil)).Elem()
}

func (i SwitchControllerSflowMap) ToSwitchControllerSflowMapOutput() SwitchControllerSflowMapOutput {
	return i.ToSwitchControllerSflowMapOutputWithContext(context.Background())
}

func (i SwitchControllerSflowMap) ToSwitchControllerSflowMapOutputWithContext(ctx context.Context) SwitchControllerSflowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSflowMapOutput)
}

type SwitchControllerSflowOutput struct{ *pulumi.OutputState }

func (SwitchControllerSflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSflow)(nil)).Elem()
}

func (o SwitchControllerSflowOutput) ToSwitchControllerSflowOutput() SwitchControllerSflowOutput {
	return o
}

func (o SwitchControllerSflowOutput) ToSwitchControllerSflowOutputWithContext(ctx context.Context) SwitchControllerSflowOutput {
	return o
}

type SwitchControllerSflowArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerSflowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSflow)(nil)).Elem()
}

func (o SwitchControllerSflowArrayOutput) ToSwitchControllerSflowArrayOutput() SwitchControllerSflowArrayOutput {
	return o
}

func (o SwitchControllerSflowArrayOutput) ToSwitchControllerSflowArrayOutputWithContext(ctx context.Context) SwitchControllerSflowArrayOutput {
	return o
}

func (o SwitchControllerSflowArrayOutput) Index(i pulumi.IntInput) SwitchControllerSflowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerSflow {
		return vs[0].([]*SwitchControllerSflow)[vs[1].(int)]
	}).(SwitchControllerSflowOutput)
}

type SwitchControllerSflowMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerSflowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSflow)(nil)).Elem()
}

func (o SwitchControllerSflowMapOutput) ToSwitchControllerSflowMapOutput() SwitchControllerSflowMapOutput {
	return o
}

func (o SwitchControllerSflowMapOutput) ToSwitchControllerSflowMapOutputWithContext(ctx context.Context) SwitchControllerSflowMapOutput {
	return o
}

func (o SwitchControllerSflowMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerSflowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerSflow {
		return vs[0].(map[string]*SwitchControllerSflow)[vs[1].(string)]
	}).(SwitchControllerSflowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSflowInput)(nil)).Elem(), &SwitchControllerSflow{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSflowArrayInput)(nil)).Elem(), SwitchControllerSflowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSflowMapInput)(nil)).Elem(), SwitchControllerSflowMap{})
	pulumi.RegisterOutputType(SwitchControllerSflowOutput{})
	pulumi.RegisterOutputType(SwitchControllerSflowArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerSflowMapOutput{})
}
