// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FIPS-CC mode.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewSystemFipsCc(ctx, "trname", &fortios.SystemFipsCcArgs{
// 			EntropyToken:          pulumi.String("enable"),
// 			KeyGenerationSelfTest: pulumi.String("disable"),
// 			SelfTestPeriod:        pulumi.Int(1440),
// 			Status:                pulumi.String("disable"),
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
// System FipsCc can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemFipsCc:SystemFipsCc labelname SystemFipsCc
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemFipsCc:SystemFipsCc labelname SystemFipsCc
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemFipsCc struct {
	pulumi.CustomResourceState

	// Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
	EntropyToken pulumi.StringOutput `pulumi:"entropyToken"`
	// Enable/disable self tests after key generation. Valid values: `enable`, `disable`.
	KeyGenerationSelfTest pulumi.StringOutput `pulumi:"keyGenerationSelfTest"`
	// Self test period.
	SelfTestPeriod pulumi.IntOutput `pulumi:"selfTestPeriod"`
	// Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemFipsCc registers a new resource with the given unique name, arguments, and options.
func NewSystemFipsCc(ctx *pulumi.Context,
	name string, args *SystemFipsCcArgs, opts ...pulumi.ResourceOption) (*SystemFipsCc, error) {
	if args == nil {
		args = &SystemFipsCcArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemFipsCc
	err := ctx.RegisterResource("fortios:index/systemFipsCc:SystemFipsCc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemFipsCc gets an existing SystemFipsCc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemFipsCc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemFipsCcState, opts ...pulumi.ResourceOption) (*SystemFipsCc, error) {
	var resource SystemFipsCc
	err := ctx.ReadResource("fortios:index/systemFipsCc:SystemFipsCc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemFipsCc resources.
type systemFipsCcState struct {
	// Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
	EntropyToken *string `pulumi:"entropyToken"`
	// Enable/disable self tests after key generation. Valid values: `enable`, `disable`.
	KeyGenerationSelfTest *string `pulumi:"keyGenerationSelfTest"`
	// Self test period.
	SelfTestPeriod *int `pulumi:"selfTestPeriod"`
	// Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemFipsCcState struct {
	// Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
	EntropyToken pulumi.StringPtrInput
	// Enable/disable self tests after key generation. Valid values: `enable`, `disable`.
	KeyGenerationSelfTest pulumi.StringPtrInput
	// Self test period.
	SelfTestPeriod pulumi.IntPtrInput
	// Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemFipsCcState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFipsCcState)(nil)).Elem()
}

type systemFipsCcArgs struct {
	// Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
	EntropyToken *string `pulumi:"entropyToken"`
	// Enable/disable self tests after key generation. Valid values: `enable`, `disable`.
	KeyGenerationSelfTest *string `pulumi:"keyGenerationSelfTest"`
	// Self test period.
	SelfTestPeriod *int `pulumi:"selfTestPeriod"`
	// Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemFipsCc resource.
type SystemFipsCcArgs struct {
	// Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
	EntropyToken pulumi.StringPtrInput
	// Enable/disable self tests after key generation. Valid values: `enable`, `disable`.
	KeyGenerationSelfTest pulumi.StringPtrInput
	// Self test period.
	SelfTestPeriod pulumi.IntPtrInput
	// Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemFipsCcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemFipsCcArgs)(nil)).Elem()
}

type SystemFipsCcInput interface {
	pulumi.Input

	ToSystemFipsCcOutput() SystemFipsCcOutput
	ToSystemFipsCcOutputWithContext(ctx context.Context) SystemFipsCcOutput
}

func (*SystemFipsCc) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFipsCc)(nil)).Elem()
}

func (i *SystemFipsCc) ToSystemFipsCcOutput() SystemFipsCcOutput {
	return i.ToSystemFipsCcOutputWithContext(context.Background())
}

func (i *SystemFipsCc) ToSystemFipsCcOutputWithContext(ctx context.Context) SystemFipsCcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFipsCcOutput)
}

// SystemFipsCcArrayInput is an input type that accepts SystemFipsCcArray and SystemFipsCcArrayOutput values.
// You can construct a concrete instance of `SystemFipsCcArrayInput` via:
//
//          SystemFipsCcArray{ SystemFipsCcArgs{...} }
type SystemFipsCcArrayInput interface {
	pulumi.Input

	ToSystemFipsCcArrayOutput() SystemFipsCcArrayOutput
	ToSystemFipsCcArrayOutputWithContext(context.Context) SystemFipsCcArrayOutput
}

type SystemFipsCcArray []SystemFipsCcInput

func (SystemFipsCcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFipsCc)(nil)).Elem()
}

func (i SystemFipsCcArray) ToSystemFipsCcArrayOutput() SystemFipsCcArrayOutput {
	return i.ToSystemFipsCcArrayOutputWithContext(context.Background())
}

func (i SystemFipsCcArray) ToSystemFipsCcArrayOutputWithContext(ctx context.Context) SystemFipsCcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFipsCcArrayOutput)
}

// SystemFipsCcMapInput is an input type that accepts SystemFipsCcMap and SystemFipsCcMapOutput values.
// You can construct a concrete instance of `SystemFipsCcMapInput` via:
//
//          SystemFipsCcMap{ "key": SystemFipsCcArgs{...} }
type SystemFipsCcMapInput interface {
	pulumi.Input

	ToSystemFipsCcMapOutput() SystemFipsCcMapOutput
	ToSystemFipsCcMapOutputWithContext(context.Context) SystemFipsCcMapOutput
}

type SystemFipsCcMap map[string]SystemFipsCcInput

func (SystemFipsCcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFipsCc)(nil)).Elem()
}

func (i SystemFipsCcMap) ToSystemFipsCcMapOutput() SystemFipsCcMapOutput {
	return i.ToSystemFipsCcMapOutputWithContext(context.Background())
}

func (i SystemFipsCcMap) ToSystemFipsCcMapOutputWithContext(ctx context.Context) SystemFipsCcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemFipsCcMapOutput)
}

type SystemFipsCcOutput struct{ *pulumi.OutputState }

func (SystemFipsCcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemFipsCc)(nil)).Elem()
}

func (o SystemFipsCcOutput) ToSystemFipsCcOutput() SystemFipsCcOutput {
	return o
}

func (o SystemFipsCcOutput) ToSystemFipsCcOutputWithContext(ctx context.Context) SystemFipsCcOutput {
	return o
}

type SystemFipsCcArrayOutput struct{ *pulumi.OutputState }

func (SystemFipsCcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemFipsCc)(nil)).Elem()
}

func (o SystemFipsCcArrayOutput) ToSystemFipsCcArrayOutput() SystemFipsCcArrayOutput {
	return o
}

func (o SystemFipsCcArrayOutput) ToSystemFipsCcArrayOutputWithContext(ctx context.Context) SystemFipsCcArrayOutput {
	return o
}

func (o SystemFipsCcArrayOutput) Index(i pulumi.IntInput) SystemFipsCcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemFipsCc {
		return vs[0].([]*SystemFipsCc)[vs[1].(int)]
	}).(SystemFipsCcOutput)
}

type SystemFipsCcMapOutput struct{ *pulumi.OutputState }

func (SystemFipsCcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemFipsCc)(nil)).Elem()
}

func (o SystemFipsCcMapOutput) ToSystemFipsCcMapOutput() SystemFipsCcMapOutput {
	return o
}

func (o SystemFipsCcMapOutput) ToSystemFipsCcMapOutputWithContext(ctx context.Context) SystemFipsCcMapOutput {
	return o
}

func (o SystemFipsCcMapOutput) MapIndex(k pulumi.StringInput) SystemFipsCcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemFipsCc {
		return vs[0].(map[string]*SystemFipsCc)[vs[1].(string)]
	}).(SystemFipsCcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFipsCcInput)(nil)).Elem(), &SystemFipsCc{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFipsCcArrayInput)(nil)).Elem(), SystemFipsCcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemFipsCcMapInput)(nil)).Elem(), SystemFipsCcMap{})
	pulumi.RegisterOutputType(SystemFipsCcOutput{})
	pulumi.RegisterOutputType(SystemFipsCcArrayOutput{})
	pulumi.RegisterOutputType(SystemFipsCcMapOutput{})
}
