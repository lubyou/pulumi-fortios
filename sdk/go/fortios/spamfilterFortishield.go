// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiGuard - AntiSpam. Applies to FortiOS Version `<= 6.2.0`.
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
// 		_, err := fortios.NewSpamfilterFortishield(ctx, "trname", &fortios.SpamfilterFortishieldArgs{
// 			SpamSubmitForce:   pulumi.String("enable"),
// 			SpamSubmitSrv:     pulumi.String("www.nospammer.net"),
// 			SpamSubmitTxt2htm: pulumi.String("enable"),
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
// Spamfilter Fortishield can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/spamfilterFortishield:SpamfilterFortishield labelname SpamfilterFortishield
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SpamfilterFortishield struct {
	pulumi.CustomResourceState

	// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
	SpamSubmitForce pulumi.StringOutput `pulumi:"spamSubmitForce"`
	// Hostname of the spam submission server.
	SpamSubmitSrv pulumi.StringOutput `pulumi:"spamSubmitSrv"`
	// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
	SpamSubmitTxt2htm pulumi.StringOutput `pulumi:"spamSubmitTxt2htm"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSpamfilterFortishield registers a new resource with the given unique name, arguments, and options.
func NewSpamfilterFortishield(ctx *pulumi.Context,
	name string, args *SpamfilterFortishieldArgs, opts ...pulumi.ResourceOption) (*SpamfilterFortishield, error) {
	if args == nil {
		args = &SpamfilterFortishieldArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SpamfilterFortishield
	err := ctx.RegisterResource("fortios:index/spamfilterFortishield:SpamfilterFortishield", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpamfilterFortishield gets an existing SpamfilterFortishield resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpamfilterFortishield(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpamfilterFortishieldState, opts ...pulumi.ResourceOption) (*SpamfilterFortishield, error) {
	var resource SpamfilterFortishield
	err := ctx.ReadResource("fortios:index/spamfilterFortishield:SpamfilterFortishield", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpamfilterFortishield resources.
type spamfilterFortishieldState struct {
	// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
	SpamSubmitForce *string `pulumi:"spamSubmitForce"`
	// Hostname of the spam submission server.
	SpamSubmitSrv *string `pulumi:"spamSubmitSrv"`
	// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
	SpamSubmitTxt2htm *string `pulumi:"spamSubmitTxt2htm"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SpamfilterFortishieldState struct {
	// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
	SpamSubmitForce pulumi.StringPtrInput
	// Hostname of the spam submission server.
	SpamSubmitSrv pulumi.StringPtrInput
	// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
	SpamSubmitTxt2htm pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpamfilterFortishieldState) ElementType() reflect.Type {
	return reflect.TypeOf((*spamfilterFortishieldState)(nil)).Elem()
}

type spamfilterFortishieldArgs struct {
	// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
	SpamSubmitForce *string `pulumi:"spamSubmitForce"`
	// Hostname of the spam submission server.
	SpamSubmitSrv *string `pulumi:"spamSubmitSrv"`
	// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
	SpamSubmitTxt2htm *string `pulumi:"spamSubmitTxt2htm"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SpamfilterFortishield resource.
type SpamfilterFortishieldArgs struct {
	// Enable/disable force insertion of a new mime entity for the submission text. Valid values: `enable`, `disable`.
	SpamSubmitForce pulumi.StringPtrInput
	// Hostname of the spam submission server.
	SpamSubmitSrv pulumi.StringPtrInput
	// Enable/disable conversion of text email to HTML email. Valid values: `enable`, `disable`.
	SpamSubmitTxt2htm pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SpamfilterFortishieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spamfilterFortishieldArgs)(nil)).Elem()
}

type SpamfilterFortishieldInput interface {
	pulumi.Input

	ToSpamfilterFortishieldOutput() SpamfilterFortishieldOutput
	ToSpamfilterFortishieldOutputWithContext(ctx context.Context) SpamfilterFortishieldOutput
}

func (*SpamfilterFortishield) ElementType() reflect.Type {
	return reflect.TypeOf((**SpamfilterFortishield)(nil)).Elem()
}

func (i *SpamfilterFortishield) ToSpamfilterFortishieldOutput() SpamfilterFortishieldOutput {
	return i.ToSpamfilterFortishieldOutputWithContext(context.Background())
}

func (i *SpamfilterFortishield) ToSpamfilterFortishieldOutputWithContext(ctx context.Context) SpamfilterFortishieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterFortishieldOutput)
}

// SpamfilterFortishieldArrayInput is an input type that accepts SpamfilterFortishieldArray and SpamfilterFortishieldArrayOutput values.
// You can construct a concrete instance of `SpamfilterFortishieldArrayInput` via:
//
//          SpamfilterFortishieldArray{ SpamfilterFortishieldArgs{...} }
type SpamfilterFortishieldArrayInput interface {
	pulumi.Input

	ToSpamfilterFortishieldArrayOutput() SpamfilterFortishieldArrayOutput
	ToSpamfilterFortishieldArrayOutputWithContext(context.Context) SpamfilterFortishieldArrayOutput
}

type SpamfilterFortishieldArray []SpamfilterFortishieldInput

func (SpamfilterFortishieldArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpamfilterFortishield)(nil)).Elem()
}

func (i SpamfilterFortishieldArray) ToSpamfilterFortishieldArrayOutput() SpamfilterFortishieldArrayOutput {
	return i.ToSpamfilterFortishieldArrayOutputWithContext(context.Background())
}

func (i SpamfilterFortishieldArray) ToSpamfilterFortishieldArrayOutputWithContext(ctx context.Context) SpamfilterFortishieldArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterFortishieldArrayOutput)
}

// SpamfilterFortishieldMapInput is an input type that accepts SpamfilterFortishieldMap and SpamfilterFortishieldMapOutput values.
// You can construct a concrete instance of `SpamfilterFortishieldMapInput` via:
//
//          SpamfilterFortishieldMap{ "key": SpamfilterFortishieldArgs{...} }
type SpamfilterFortishieldMapInput interface {
	pulumi.Input

	ToSpamfilterFortishieldMapOutput() SpamfilterFortishieldMapOutput
	ToSpamfilterFortishieldMapOutputWithContext(context.Context) SpamfilterFortishieldMapOutput
}

type SpamfilterFortishieldMap map[string]SpamfilterFortishieldInput

func (SpamfilterFortishieldMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpamfilterFortishield)(nil)).Elem()
}

func (i SpamfilterFortishieldMap) ToSpamfilterFortishieldMapOutput() SpamfilterFortishieldMapOutput {
	return i.ToSpamfilterFortishieldMapOutputWithContext(context.Background())
}

func (i SpamfilterFortishieldMap) ToSpamfilterFortishieldMapOutputWithContext(ctx context.Context) SpamfilterFortishieldMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpamfilterFortishieldMapOutput)
}

type SpamfilterFortishieldOutput struct{ *pulumi.OutputState }

func (SpamfilterFortishieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpamfilterFortishield)(nil)).Elem()
}

func (o SpamfilterFortishieldOutput) ToSpamfilterFortishieldOutput() SpamfilterFortishieldOutput {
	return o
}

func (o SpamfilterFortishieldOutput) ToSpamfilterFortishieldOutputWithContext(ctx context.Context) SpamfilterFortishieldOutput {
	return o
}

type SpamfilterFortishieldArrayOutput struct{ *pulumi.OutputState }

func (SpamfilterFortishieldArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpamfilterFortishield)(nil)).Elem()
}

func (o SpamfilterFortishieldArrayOutput) ToSpamfilterFortishieldArrayOutput() SpamfilterFortishieldArrayOutput {
	return o
}

func (o SpamfilterFortishieldArrayOutput) ToSpamfilterFortishieldArrayOutputWithContext(ctx context.Context) SpamfilterFortishieldArrayOutput {
	return o
}

func (o SpamfilterFortishieldArrayOutput) Index(i pulumi.IntInput) SpamfilterFortishieldOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SpamfilterFortishield {
		return vs[0].([]*SpamfilterFortishield)[vs[1].(int)]
	}).(SpamfilterFortishieldOutput)
}

type SpamfilterFortishieldMapOutput struct{ *pulumi.OutputState }

func (SpamfilterFortishieldMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpamfilterFortishield)(nil)).Elem()
}

func (o SpamfilterFortishieldMapOutput) ToSpamfilterFortishieldMapOutput() SpamfilterFortishieldMapOutput {
	return o
}

func (o SpamfilterFortishieldMapOutput) ToSpamfilterFortishieldMapOutputWithContext(ctx context.Context) SpamfilterFortishieldMapOutput {
	return o
}

func (o SpamfilterFortishieldMapOutput) MapIndex(k pulumi.StringInput) SpamfilterFortishieldOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SpamfilterFortishield {
		return vs[0].(map[string]*SpamfilterFortishield)[vs[1].(string)]
	}).(SpamfilterFortishieldOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterFortishieldInput)(nil)).Elem(), &SpamfilterFortishield{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterFortishieldArrayInput)(nil)).Elem(), SpamfilterFortishieldArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpamfilterFortishieldMapInput)(nil)).Elem(), SpamfilterFortishieldMap{})
	pulumi.RegisterOutputType(SpamfilterFortishieldOutput{})
	pulumi.RegisterOutputType(SpamfilterFortishieldArrayOutput{})
	pulumi.RegisterOutputType(SpamfilterFortishieldMapOutput{})
}
