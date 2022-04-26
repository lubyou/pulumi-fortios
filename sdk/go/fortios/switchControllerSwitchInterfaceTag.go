// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure switch object tags.
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
// 		_, err := fortios.NewSwitchControllerSwitchInterfaceTag(ctx, "trname", nil)
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
// SwitchController SwitchInterfaceTag can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/switchControllerSwitchInterfaceTag:SwitchControllerSwitchInterfaceTag labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/switchControllerSwitchInterfaceTag:SwitchControllerSwitchInterfaceTag labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SwitchControllerSwitchInterfaceTag struct {
	pulumi.CustomResourceState

	// Tag name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchControllerSwitchInterfaceTag registers a new resource with the given unique name, arguments, and options.
func NewSwitchControllerSwitchInterfaceTag(ctx *pulumi.Context,
	name string, args *SwitchControllerSwitchInterfaceTagArgs, opts ...pulumi.ResourceOption) (*SwitchControllerSwitchInterfaceTag, error) {
	if args == nil {
		args = &SwitchControllerSwitchInterfaceTagArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchControllerSwitchInterfaceTag
	err := ctx.RegisterResource("fortios:index/switchControllerSwitchInterfaceTag:SwitchControllerSwitchInterfaceTag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchControllerSwitchInterfaceTag gets an existing SwitchControllerSwitchInterfaceTag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchControllerSwitchInterfaceTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchControllerSwitchInterfaceTagState, opts ...pulumi.ResourceOption) (*SwitchControllerSwitchInterfaceTag, error) {
	var resource SwitchControllerSwitchInterfaceTag
	err := ctx.ReadResource("fortios:index/switchControllerSwitchInterfaceTag:SwitchControllerSwitchInterfaceTag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchControllerSwitchInterfaceTag resources.
type switchControllerSwitchInterfaceTagState struct {
	// Tag name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchControllerSwitchInterfaceTagState struct {
	// Tag name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerSwitchInterfaceTagState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSwitchInterfaceTagState)(nil)).Elem()
}

type switchControllerSwitchInterfaceTagArgs struct {
	// Tag name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchControllerSwitchInterfaceTag resource.
type SwitchControllerSwitchInterfaceTagArgs struct {
	// Tag name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchControllerSwitchInterfaceTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchControllerSwitchInterfaceTagArgs)(nil)).Elem()
}

type SwitchControllerSwitchInterfaceTagInput interface {
	pulumi.Input

	ToSwitchControllerSwitchInterfaceTagOutput() SwitchControllerSwitchInterfaceTagOutput
	ToSwitchControllerSwitchInterfaceTagOutputWithContext(ctx context.Context) SwitchControllerSwitchInterfaceTagOutput
}

func (*SwitchControllerSwitchInterfaceTag) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSwitchInterfaceTag)(nil)).Elem()
}

func (i *SwitchControllerSwitchInterfaceTag) ToSwitchControllerSwitchInterfaceTagOutput() SwitchControllerSwitchInterfaceTagOutput {
	return i.ToSwitchControllerSwitchInterfaceTagOutputWithContext(context.Background())
}

func (i *SwitchControllerSwitchInterfaceTag) ToSwitchControllerSwitchInterfaceTagOutputWithContext(ctx context.Context) SwitchControllerSwitchInterfaceTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSwitchInterfaceTagOutput)
}

// SwitchControllerSwitchInterfaceTagArrayInput is an input type that accepts SwitchControllerSwitchInterfaceTagArray and SwitchControllerSwitchInterfaceTagArrayOutput values.
// You can construct a concrete instance of `SwitchControllerSwitchInterfaceTagArrayInput` via:
//
//          SwitchControllerSwitchInterfaceTagArray{ SwitchControllerSwitchInterfaceTagArgs{...} }
type SwitchControllerSwitchInterfaceTagArrayInput interface {
	pulumi.Input

	ToSwitchControllerSwitchInterfaceTagArrayOutput() SwitchControllerSwitchInterfaceTagArrayOutput
	ToSwitchControllerSwitchInterfaceTagArrayOutputWithContext(context.Context) SwitchControllerSwitchInterfaceTagArrayOutput
}

type SwitchControllerSwitchInterfaceTagArray []SwitchControllerSwitchInterfaceTagInput

func (SwitchControllerSwitchInterfaceTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSwitchInterfaceTag)(nil)).Elem()
}

func (i SwitchControllerSwitchInterfaceTagArray) ToSwitchControllerSwitchInterfaceTagArrayOutput() SwitchControllerSwitchInterfaceTagArrayOutput {
	return i.ToSwitchControllerSwitchInterfaceTagArrayOutputWithContext(context.Background())
}

func (i SwitchControllerSwitchInterfaceTagArray) ToSwitchControllerSwitchInterfaceTagArrayOutputWithContext(ctx context.Context) SwitchControllerSwitchInterfaceTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSwitchInterfaceTagArrayOutput)
}

// SwitchControllerSwitchInterfaceTagMapInput is an input type that accepts SwitchControllerSwitchInterfaceTagMap and SwitchControllerSwitchInterfaceTagMapOutput values.
// You can construct a concrete instance of `SwitchControllerSwitchInterfaceTagMapInput` via:
//
//          SwitchControllerSwitchInterfaceTagMap{ "key": SwitchControllerSwitchInterfaceTagArgs{...} }
type SwitchControllerSwitchInterfaceTagMapInput interface {
	pulumi.Input

	ToSwitchControllerSwitchInterfaceTagMapOutput() SwitchControllerSwitchInterfaceTagMapOutput
	ToSwitchControllerSwitchInterfaceTagMapOutputWithContext(context.Context) SwitchControllerSwitchInterfaceTagMapOutput
}

type SwitchControllerSwitchInterfaceTagMap map[string]SwitchControllerSwitchInterfaceTagInput

func (SwitchControllerSwitchInterfaceTagMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSwitchInterfaceTag)(nil)).Elem()
}

func (i SwitchControllerSwitchInterfaceTagMap) ToSwitchControllerSwitchInterfaceTagMapOutput() SwitchControllerSwitchInterfaceTagMapOutput {
	return i.ToSwitchControllerSwitchInterfaceTagMapOutputWithContext(context.Background())
}

func (i SwitchControllerSwitchInterfaceTagMap) ToSwitchControllerSwitchInterfaceTagMapOutputWithContext(ctx context.Context) SwitchControllerSwitchInterfaceTagMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchControllerSwitchInterfaceTagMapOutput)
}

type SwitchControllerSwitchInterfaceTagOutput struct{ *pulumi.OutputState }

func (SwitchControllerSwitchInterfaceTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchControllerSwitchInterfaceTag)(nil)).Elem()
}

func (o SwitchControllerSwitchInterfaceTagOutput) ToSwitchControllerSwitchInterfaceTagOutput() SwitchControllerSwitchInterfaceTagOutput {
	return o
}

func (o SwitchControllerSwitchInterfaceTagOutput) ToSwitchControllerSwitchInterfaceTagOutputWithContext(ctx context.Context) SwitchControllerSwitchInterfaceTagOutput {
	return o
}

type SwitchControllerSwitchInterfaceTagArrayOutput struct{ *pulumi.OutputState }

func (SwitchControllerSwitchInterfaceTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchControllerSwitchInterfaceTag)(nil)).Elem()
}

func (o SwitchControllerSwitchInterfaceTagArrayOutput) ToSwitchControllerSwitchInterfaceTagArrayOutput() SwitchControllerSwitchInterfaceTagArrayOutput {
	return o
}

func (o SwitchControllerSwitchInterfaceTagArrayOutput) ToSwitchControllerSwitchInterfaceTagArrayOutputWithContext(ctx context.Context) SwitchControllerSwitchInterfaceTagArrayOutput {
	return o
}

func (o SwitchControllerSwitchInterfaceTagArrayOutput) Index(i pulumi.IntInput) SwitchControllerSwitchInterfaceTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchControllerSwitchInterfaceTag {
		return vs[0].([]*SwitchControllerSwitchInterfaceTag)[vs[1].(int)]
	}).(SwitchControllerSwitchInterfaceTagOutput)
}

type SwitchControllerSwitchInterfaceTagMapOutput struct{ *pulumi.OutputState }

func (SwitchControllerSwitchInterfaceTagMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchControllerSwitchInterfaceTag)(nil)).Elem()
}

func (o SwitchControllerSwitchInterfaceTagMapOutput) ToSwitchControllerSwitchInterfaceTagMapOutput() SwitchControllerSwitchInterfaceTagMapOutput {
	return o
}

func (o SwitchControllerSwitchInterfaceTagMapOutput) ToSwitchControllerSwitchInterfaceTagMapOutputWithContext(ctx context.Context) SwitchControllerSwitchInterfaceTagMapOutput {
	return o
}

func (o SwitchControllerSwitchInterfaceTagMapOutput) MapIndex(k pulumi.StringInput) SwitchControllerSwitchInterfaceTagOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchControllerSwitchInterfaceTag {
		return vs[0].(map[string]*SwitchControllerSwitchInterfaceTag)[vs[1].(string)]
	}).(SwitchControllerSwitchInterfaceTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSwitchInterfaceTagInput)(nil)).Elem(), &SwitchControllerSwitchInterfaceTag{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSwitchInterfaceTagArrayInput)(nil)).Elem(), SwitchControllerSwitchInterfaceTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchControllerSwitchInterfaceTagMapInput)(nil)).Elem(), SwitchControllerSwitchInterfaceTagMap{})
	pulumi.RegisterOutputType(SwitchControllerSwitchInterfaceTagOutput{})
	pulumi.RegisterOutputType(SwitchControllerSwitchInterfaceTagArrayOutput{})
	pulumi.RegisterOutputType(SwitchControllerSwitchInterfaceTagMapOutput{})
}
