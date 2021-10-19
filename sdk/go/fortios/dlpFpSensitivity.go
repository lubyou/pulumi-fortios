// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create self-explanatory DLP sensitivity levels to be used when setting sensitivity under config fp-doc-source. Applies to FortiOS Version `<= 6.2.0`.
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
// 		_, err := fortios.NewDlpFpSensitivity(ctx, "trname", nil)
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
// Dlp FpSensitivity can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/dlpFpSensitivity:DlpFpSensitivity labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type DlpFpSensitivity struct {
	pulumi.CustomResourceState

	// DLP Sensitivity Levels.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewDlpFpSensitivity registers a new resource with the given unique name, arguments, and options.
func NewDlpFpSensitivity(ctx *pulumi.Context,
	name string, args *DlpFpSensitivityArgs, opts ...pulumi.ResourceOption) (*DlpFpSensitivity, error) {
	if args == nil {
		args = &DlpFpSensitivityArgs{}
	}

	var resource DlpFpSensitivity
	err := ctx.RegisterResource("fortios:index/dlpFpSensitivity:DlpFpSensitivity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDlpFpSensitivity gets an existing DlpFpSensitivity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDlpFpSensitivity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DlpFpSensitivityState, opts ...pulumi.ResourceOption) (*DlpFpSensitivity, error) {
	var resource DlpFpSensitivity
	err := ctx.ReadResource("fortios:index/dlpFpSensitivity:DlpFpSensitivity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DlpFpSensitivity resources.
type dlpFpSensitivityState struct {
	// DLP Sensitivity Levels.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type DlpFpSensitivityState struct {
	// DLP Sensitivity Levels.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DlpFpSensitivityState) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpFpSensitivityState)(nil)).Elem()
}

type dlpFpSensitivityArgs struct {
	// DLP Sensitivity Levels.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a DlpFpSensitivity resource.
type DlpFpSensitivityArgs struct {
	// DLP Sensitivity Levels.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (DlpFpSensitivityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dlpFpSensitivityArgs)(nil)).Elem()
}

type DlpFpSensitivityInput interface {
	pulumi.Input

	ToDlpFpSensitivityOutput() DlpFpSensitivityOutput
	ToDlpFpSensitivityOutputWithContext(ctx context.Context) DlpFpSensitivityOutput
}

func (*DlpFpSensitivity) ElementType() reflect.Type {
	return reflect.TypeOf((*DlpFpSensitivity)(nil))
}

func (i *DlpFpSensitivity) ToDlpFpSensitivityOutput() DlpFpSensitivityOutput {
	return i.ToDlpFpSensitivityOutputWithContext(context.Background())
}

func (i *DlpFpSensitivity) ToDlpFpSensitivityOutputWithContext(ctx context.Context) DlpFpSensitivityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpFpSensitivityOutput)
}

func (i *DlpFpSensitivity) ToDlpFpSensitivityPtrOutput() DlpFpSensitivityPtrOutput {
	return i.ToDlpFpSensitivityPtrOutputWithContext(context.Background())
}

func (i *DlpFpSensitivity) ToDlpFpSensitivityPtrOutputWithContext(ctx context.Context) DlpFpSensitivityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpFpSensitivityPtrOutput)
}

type DlpFpSensitivityPtrInput interface {
	pulumi.Input

	ToDlpFpSensitivityPtrOutput() DlpFpSensitivityPtrOutput
	ToDlpFpSensitivityPtrOutputWithContext(ctx context.Context) DlpFpSensitivityPtrOutput
}

type dlpFpSensitivityPtrType DlpFpSensitivityArgs

func (*dlpFpSensitivityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpFpSensitivity)(nil))
}

func (i *dlpFpSensitivityPtrType) ToDlpFpSensitivityPtrOutput() DlpFpSensitivityPtrOutput {
	return i.ToDlpFpSensitivityPtrOutputWithContext(context.Background())
}

func (i *dlpFpSensitivityPtrType) ToDlpFpSensitivityPtrOutputWithContext(ctx context.Context) DlpFpSensitivityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpFpSensitivityPtrOutput)
}

// DlpFpSensitivityArrayInput is an input type that accepts DlpFpSensitivityArray and DlpFpSensitivityArrayOutput values.
// You can construct a concrete instance of `DlpFpSensitivityArrayInput` via:
//
//          DlpFpSensitivityArray{ DlpFpSensitivityArgs{...} }
type DlpFpSensitivityArrayInput interface {
	pulumi.Input

	ToDlpFpSensitivityArrayOutput() DlpFpSensitivityArrayOutput
	ToDlpFpSensitivityArrayOutputWithContext(context.Context) DlpFpSensitivityArrayOutput
}

type DlpFpSensitivityArray []DlpFpSensitivityInput

func (DlpFpSensitivityArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DlpFpSensitivity)(nil))
}

func (i DlpFpSensitivityArray) ToDlpFpSensitivityArrayOutput() DlpFpSensitivityArrayOutput {
	return i.ToDlpFpSensitivityArrayOutputWithContext(context.Background())
}

func (i DlpFpSensitivityArray) ToDlpFpSensitivityArrayOutputWithContext(ctx context.Context) DlpFpSensitivityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpFpSensitivityArrayOutput)
}

// DlpFpSensitivityMapInput is an input type that accepts DlpFpSensitivityMap and DlpFpSensitivityMapOutput values.
// You can construct a concrete instance of `DlpFpSensitivityMapInput` via:
//
//          DlpFpSensitivityMap{ "key": DlpFpSensitivityArgs{...} }
type DlpFpSensitivityMapInput interface {
	pulumi.Input

	ToDlpFpSensitivityMapOutput() DlpFpSensitivityMapOutput
	ToDlpFpSensitivityMapOutputWithContext(context.Context) DlpFpSensitivityMapOutput
}

type DlpFpSensitivityMap map[string]DlpFpSensitivityInput

func (DlpFpSensitivityMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DlpFpSensitivity)(nil))
}

func (i DlpFpSensitivityMap) ToDlpFpSensitivityMapOutput() DlpFpSensitivityMapOutput {
	return i.ToDlpFpSensitivityMapOutputWithContext(context.Background())
}

func (i DlpFpSensitivityMap) ToDlpFpSensitivityMapOutputWithContext(ctx context.Context) DlpFpSensitivityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DlpFpSensitivityMapOutput)
}

type DlpFpSensitivityOutput struct {
	*pulumi.OutputState
}

func (DlpFpSensitivityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DlpFpSensitivity)(nil))
}

func (o DlpFpSensitivityOutput) ToDlpFpSensitivityOutput() DlpFpSensitivityOutput {
	return o
}

func (o DlpFpSensitivityOutput) ToDlpFpSensitivityOutputWithContext(ctx context.Context) DlpFpSensitivityOutput {
	return o
}

func (o DlpFpSensitivityOutput) ToDlpFpSensitivityPtrOutput() DlpFpSensitivityPtrOutput {
	return o.ToDlpFpSensitivityPtrOutputWithContext(context.Background())
}

func (o DlpFpSensitivityOutput) ToDlpFpSensitivityPtrOutputWithContext(ctx context.Context) DlpFpSensitivityPtrOutput {
	return o.ApplyT(func(v DlpFpSensitivity) *DlpFpSensitivity {
		return &v
	}).(DlpFpSensitivityPtrOutput)
}

type DlpFpSensitivityPtrOutput struct {
	*pulumi.OutputState
}

func (DlpFpSensitivityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DlpFpSensitivity)(nil))
}

func (o DlpFpSensitivityPtrOutput) ToDlpFpSensitivityPtrOutput() DlpFpSensitivityPtrOutput {
	return o
}

func (o DlpFpSensitivityPtrOutput) ToDlpFpSensitivityPtrOutputWithContext(ctx context.Context) DlpFpSensitivityPtrOutput {
	return o
}

type DlpFpSensitivityArrayOutput struct{ *pulumi.OutputState }

func (DlpFpSensitivityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DlpFpSensitivity)(nil))
}

func (o DlpFpSensitivityArrayOutput) ToDlpFpSensitivityArrayOutput() DlpFpSensitivityArrayOutput {
	return o
}

func (o DlpFpSensitivityArrayOutput) ToDlpFpSensitivityArrayOutputWithContext(ctx context.Context) DlpFpSensitivityArrayOutput {
	return o
}

func (o DlpFpSensitivityArrayOutput) Index(i pulumi.IntInput) DlpFpSensitivityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DlpFpSensitivity {
		return vs[0].([]DlpFpSensitivity)[vs[1].(int)]
	}).(DlpFpSensitivityOutput)
}

type DlpFpSensitivityMapOutput struct{ *pulumi.OutputState }

func (DlpFpSensitivityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DlpFpSensitivity)(nil))
}

func (o DlpFpSensitivityMapOutput) ToDlpFpSensitivityMapOutput() DlpFpSensitivityMapOutput {
	return o
}

func (o DlpFpSensitivityMapOutput) ToDlpFpSensitivityMapOutputWithContext(ctx context.Context) DlpFpSensitivityMapOutput {
	return o
}

func (o DlpFpSensitivityMapOutput) MapIndex(k pulumi.StringInput) DlpFpSensitivityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DlpFpSensitivity {
		return vs[0].(map[string]DlpFpSensitivity)[vs[1].(string)]
	}).(DlpFpSensitivityOutput)
}

func init() {
	pulumi.RegisterOutputType(DlpFpSensitivityOutput{})
	pulumi.RegisterOutputType(DlpFpSensitivityPtrOutput{})
	pulumi.RegisterOutputType(DlpFpSensitivityArrayOutput{})
	pulumi.RegisterOutputType(DlpFpSensitivityMapOutput{})
}
