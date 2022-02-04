// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPS custom signature.
//
// ## Import
//
// Ips Custom can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/ipsCustom:IpsCustom labelname {{tag}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type IpsCustom struct {
	pulumi.CustomResourceState

	// Default action (pass or block) for this signature. Valid values: `pass`, `block`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Applications to be protected. Blank for all applications.
	Application pulumi.StringOutput `pulumi:"application"`
	// Comment.
	Comment pulumi.StringOutput `pulumi:"comment"`
	// Protect client or server traffic.
	Location pulumi.StringOutput `pulumi:"location"`
	// Enable/disable logging. Valid values: `disable`, `enable`.
	Log pulumi.StringOutput `pulumi:"log"`
	// Enable/disable packet logging. Valid values: `disable`, `enable`.
	LogPacket pulumi.StringOutput `pulumi:"logPacket"`
	// Operating system(s) that the signature protects. Blank for all operating systems.
	Os pulumi.StringOutput `pulumi:"os"`
	// Protocol(s) that the signature scans. Blank for all protocols.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Signature ID.
	RuleId pulumi.IntOutput `pulumi:"ruleId"`
	// Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// Signature name.
	SigName pulumi.StringOutput `pulumi:"sigName"`
	// Custom signature enclosed in single quotes.
	Signature pulumi.StringOutput `pulumi:"signature"`
	// Enable/disable this signature. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Signature tag.
	Tag pulumi.StringOutput `pulumi:"tag"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIpsCustom registers a new resource with the given unique name, arguments, and options.
func NewIpsCustom(ctx *pulumi.Context,
	name string, args *IpsCustomArgs, opts ...pulumi.ResourceOption) (*IpsCustom, error) {
	if args == nil {
		args = &IpsCustomArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource IpsCustom
	err := ctx.RegisterResource("fortios:index/ipsCustom:IpsCustom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpsCustom gets an existing IpsCustom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpsCustom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpsCustomState, opts ...pulumi.ResourceOption) (*IpsCustom, error) {
	var resource IpsCustom
	err := ctx.ReadResource("fortios:index/ipsCustom:IpsCustom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpsCustom resources.
type ipsCustomState struct {
	// Default action (pass or block) for this signature. Valid values: `pass`, `block`.
	Action *string `pulumi:"action"`
	// Applications to be protected. Blank for all applications.
	Application *string `pulumi:"application"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Protect client or server traffic.
	Location *string `pulumi:"location"`
	// Enable/disable logging. Valid values: `disable`, `enable`.
	Log *string `pulumi:"log"`
	// Enable/disable packet logging. Valid values: `disable`, `enable`.
	LogPacket *string `pulumi:"logPacket"`
	// Operating system(s) that the signature protects. Blank for all operating systems.
	Os *string `pulumi:"os"`
	// Protocol(s) that the signature scans. Blank for all protocols.
	Protocol *string `pulumi:"protocol"`
	// Signature ID.
	RuleId *int `pulumi:"ruleId"`
	// Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
	Severity *string `pulumi:"severity"`
	// Signature name.
	SigName *string `pulumi:"sigName"`
	// Custom signature enclosed in single quotes.
	Signature *string `pulumi:"signature"`
	// Enable/disable this signature. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Signature tag.
	Tag *string `pulumi:"tag"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type IpsCustomState struct {
	// Default action (pass or block) for this signature. Valid values: `pass`, `block`.
	Action pulumi.StringPtrInput
	// Applications to be protected. Blank for all applications.
	Application pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Protect client or server traffic.
	Location pulumi.StringPtrInput
	// Enable/disable logging. Valid values: `disable`, `enable`.
	Log pulumi.StringPtrInput
	// Enable/disable packet logging. Valid values: `disable`, `enable`.
	LogPacket pulumi.StringPtrInput
	// Operating system(s) that the signature protects. Blank for all operating systems.
	Os pulumi.StringPtrInput
	// Protocol(s) that the signature scans. Blank for all protocols.
	Protocol pulumi.StringPtrInput
	// Signature ID.
	RuleId pulumi.IntPtrInput
	// Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
	Severity pulumi.StringPtrInput
	// Signature name.
	SigName pulumi.StringPtrInput
	// Custom signature enclosed in single quotes.
	Signature pulumi.StringPtrInput
	// Enable/disable this signature. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Signature tag.
	Tag pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsCustomState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsCustomState)(nil)).Elem()
}

type ipsCustomArgs struct {
	// Default action (pass or block) for this signature. Valid values: `pass`, `block`.
	Action *string `pulumi:"action"`
	// Applications to be protected. Blank for all applications.
	Application *string `pulumi:"application"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// Protect client or server traffic.
	Location *string `pulumi:"location"`
	// Enable/disable logging. Valid values: `disable`, `enable`.
	Log *string `pulumi:"log"`
	// Enable/disable packet logging. Valid values: `disable`, `enable`.
	LogPacket *string `pulumi:"logPacket"`
	// Operating system(s) that the signature protects. Blank for all operating systems.
	Os *string `pulumi:"os"`
	// Protocol(s) that the signature scans. Blank for all protocols.
	Protocol *string `pulumi:"protocol"`
	// Signature ID.
	RuleId *int `pulumi:"ruleId"`
	// Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
	Severity *string `pulumi:"severity"`
	// Signature name.
	SigName *string `pulumi:"sigName"`
	// Custom signature enclosed in single quotes.
	Signature *string `pulumi:"signature"`
	// Enable/disable this signature. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Signature tag.
	Tag *string `pulumi:"tag"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a IpsCustom resource.
type IpsCustomArgs struct {
	// Default action (pass or block) for this signature. Valid values: `pass`, `block`.
	Action pulumi.StringPtrInput
	// Applications to be protected. Blank for all applications.
	Application pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// Protect client or server traffic.
	Location pulumi.StringPtrInput
	// Enable/disable logging. Valid values: `disable`, `enable`.
	Log pulumi.StringPtrInput
	// Enable/disable packet logging. Valid values: `disable`, `enable`.
	LogPacket pulumi.StringPtrInput
	// Operating system(s) that the signature protects. Blank for all operating systems.
	Os pulumi.StringPtrInput
	// Protocol(s) that the signature scans. Blank for all protocols.
	Protocol pulumi.StringPtrInput
	// Signature ID.
	RuleId pulumi.IntPtrInput
	// Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.
	Severity pulumi.StringPtrInput
	// Signature name.
	SigName pulumi.StringPtrInput
	// Custom signature enclosed in single quotes.
	Signature pulumi.StringPtrInput
	// Enable/disable this signature. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Signature tag.
	Tag pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsCustomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsCustomArgs)(nil)).Elem()
}

type IpsCustomInput interface {
	pulumi.Input

	ToIpsCustomOutput() IpsCustomOutput
	ToIpsCustomOutputWithContext(ctx context.Context) IpsCustomOutput
}

func (*IpsCustom) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsCustom)(nil)).Elem()
}

func (i *IpsCustom) ToIpsCustomOutput() IpsCustomOutput {
	return i.ToIpsCustomOutputWithContext(context.Background())
}

func (i *IpsCustom) ToIpsCustomOutputWithContext(ctx context.Context) IpsCustomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsCustomOutput)
}

// IpsCustomArrayInput is an input type that accepts IpsCustomArray and IpsCustomArrayOutput values.
// You can construct a concrete instance of `IpsCustomArrayInput` via:
//
//          IpsCustomArray{ IpsCustomArgs{...} }
type IpsCustomArrayInput interface {
	pulumi.Input

	ToIpsCustomArrayOutput() IpsCustomArrayOutput
	ToIpsCustomArrayOutputWithContext(context.Context) IpsCustomArrayOutput
}

type IpsCustomArray []IpsCustomInput

func (IpsCustomArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsCustom)(nil)).Elem()
}

func (i IpsCustomArray) ToIpsCustomArrayOutput() IpsCustomArrayOutput {
	return i.ToIpsCustomArrayOutputWithContext(context.Background())
}

func (i IpsCustomArray) ToIpsCustomArrayOutputWithContext(ctx context.Context) IpsCustomArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsCustomArrayOutput)
}

// IpsCustomMapInput is an input type that accepts IpsCustomMap and IpsCustomMapOutput values.
// You can construct a concrete instance of `IpsCustomMapInput` via:
//
//          IpsCustomMap{ "key": IpsCustomArgs{...} }
type IpsCustomMapInput interface {
	pulumi.Input

	ToIpsCustomMapOutput() IpsCustomMapOutput
	ToIpsCustomMapOutputWithContext(context.Context) IpsCustomMapOutput
}

type IpsCustomMap map[string]IpsCustomInput

func (IpsCustomMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsCustom)(nil)).Elem()
}

func (i IpsCustomMap) ToIpsCustomMapOutput() IpsCustomMapOutput {
	return i.ToIpsCustomMapOutputWithContext(context.Background())
}

func (i IpsCustomMap) ToIpsCustomMapOutputWithContext(ctx context.Context) IpsCustomMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsCustomMapOutput)
}

type IpsCustomOutput struct{ *pulumi.OutputState }

func (IpsCustomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsCustom)(nil)).Elem()
}

func (o IpsCustomOutput) ToIpsCustomOutput() IpsCustomOutput {
	return o
}

func (o IpsCustomOutput) ToIpsCustomOutputWithContext(ctx context.Context) IpsCustomOutput {
	return o
}

type IpsCustomArrayOutput struct{ *pulumi.OutputState }

func (IpsCustomArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsCustom)(nil)).Elem()
}

func (o IpsCustomArrayOutput) ToIpsCustomArrayOutput() IpsCustomArrayOutput {
	return o
}

func (o IpsCustomArrayOutput) ToIpsCustomArrayOutputWithContext(ctx context.Context) IpsCustomArrayOutput {
	return o
}

func (o IpsCustomArrayOutput) Index(i pulumi.IntInput) IpsCustomOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpsCustom {
		return vs[0].([]*IpsCustom)[vs[1].(int)]
	}).(IpsCustomOutput)
}

type IpsCustomMapOutput struct{ *pulumi.OutputState }

func (IpsCustomMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsCustom)(nil)).Elem()
}

func (o IpsCustomMapOutput) ToIpsCustomMapOutput() IpsCustomMapOutput {
	return o
}

func (o IpsCustomMapOutput) ToIpsCustomMapOutputWithContext(ctx context.Context) IpsCustomMapOutput {
	return o
}

func (o IpsCustomMapOutput) MapIndex(k pulumi.StringInput) IpsCustomOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpsCustom {
		return vs[0].(map[string]*IpsCustom)[vs[1].(string)]
	}).(IpsCustomOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpsCustomInput)(nil)).Elem(), &IpsCustom{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsCustomArrayInput)(nil)).Elem(), IpsCustomArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsCustomMapInput)(nil)).Elem(), IpsCustomMap{})
	pulumi.RegisterOutputType(IpsCustomOutput{})
	pulumi.RegisterOutputType(IpsCustomArrayOutput{})
	pulumi.RegisterOutputType(IpsCustomMapOutput{})
}
