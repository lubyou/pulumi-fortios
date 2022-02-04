// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Replacement messages.
//
// ## Import
//
// SystemReplacemsg Alertmail can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemReplacemsgAlertmail:SystemReplacemsgAlertmail labelname {{msg_type}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemReplacemsgAlertmail struct {
	pulumi.CustomResourceState

	// Message string.
	Buffer pulumi.StringPtrOutput `pulumi:"buffer"`
	// Format flag.
	Format pulumi.StringOutput `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringOutput `pulumi:"header"`
	// Message type.
	MsgType pulumi.StringOutput `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemReplacemsgAlertmail registers a new resource with the given unique name, arguments, and options.
func NewSystemReplacemsgAlertmail(ctx *pulumi.Context,
	name string, args *SystemReplacemsgAlertmailArgs, opts ...pulumi.ResourceOption) (*SystemReplacemsgAlertmail, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemReplacemsgAlertmail
	err := ctx.RegisterResource("fortios:index/systemReplacemsgAlertmail:SystemReplacemsgAlertmail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemReplacemsgAlertmail gets an existing SystemReplacemsgAlertmail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemReplacemsgAlertmail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemReplacemsgAlertmailState, opts ...pulumi.ResourceOption) (*SystemReplacemsgAlertmail, error) {
	var resource SystemReplacemsgAlertmail
	err := ctx.ReadResource("fortios:index/systemReplacemsgAlertmail:SystemReplacemsgAlertmail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemReplacemsgAlertmail resources.
type systemReplacemsgAlertmailState struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType *string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemReplacemsgAlertmailState struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgAlertmailState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgAlertmailState)(nil)).Elem()
}

type systemReplacemsgAlertmailArgs struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemReplacemsgAlertmail resource.
type SystemReplacemsgAlertmailArgs struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemReplacemsgAlertmailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemReplacemsgAlertmailArgs)(nil)).Elem()
}

type SystemReplacemsgAlertmailInput interface {
	pulumi.Input

	ToSystemReplacemsgAlertmailOutput() SystemReplacemsgAlertmailOutput
	ToSystemReplacemsgAlertmailOutputWithContext(ctx context.Context) SystemReplacemsgAlertmailOutput
}

func (*SystemReplacemsgAlertmail) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgAlertmail)(nil)).Elem()
}

func (i *SystemReplacemsgAlertmail) ToSystemReplacemsgAlertmailOutput() SystemReplacemsgAlertmailOutput {
	return i.ToSystemReplacemsgAlertmailOutputWithContext(context.Background())
}

func (i *SystemReplacemsgAlertmail) ToSystemReplacemsgAlertmailOutputWithContext(ctx context.Context) SystemReplacemsgAlertmailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgAlertmailOutput)
}

// SystemReplacemsgAlertmailArrayInput is an input type that accepts SystemReplacemsgAlertmailArray and SystemReplacemsgAlertmailArrayOutput values.
// You can construct a concrete instance of `SystemReplacemsgAlertmailArrayInput` via:
//
//          SystemReplacemsgAlertmailArray{ SystemReplacemsgAlertmailArgs{...} }
type SystemReplacemsgAlertmailArrayInput interface {
	pulumi.Input

	ToSystemReplacemsgAlertmailArrayOutput() SystemReplacemsgAlertmailArrayOutput
	ToSystemReplacemsgAlertmailArrayOutputWithContext(context.Context) SystemReplacemsgAlertmailArrayOutput
}

type SystemReplacemsgAlertmailArray []SystemReplacemsgAlertmailInput

func (SystemReplacemsgAlertmailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgAlertmail)(nil)).Elem()
}

func (i SystemReplacemsgAlertmailArray) ToSystemReplacemsgAlertmailArrayOutput() SystemReplacemsgAlertmailArrayOutput {
	return i.ToSystemReplacemsgAlertmailArrayOutputWithContext(context.Background())
}

func (i SystemReplacemsgAlertmailArray) ToSystemReplacemsgAlertmailArrayOutputWithContext(ctx context.Context) SystemReplacemsgAlertmailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgAlertmailArrayOutput)
}

// SystemReplacemsgAlertmailMapInput is an input type that accepts SystemReplacemsgAlertmailMap and SystemReplacemsgAlertmailMapOutput values.
// You can construct a concrete instance of `SystemReplacemsgAlertmailMapInput` via:
//
//          SystemReplacemsgAlertmailMap{ "key": SystemReplacemsgAlertmailArgs{...} }
type SystemReplacemsgAlertmailMapInput interface {
	pulumi.Input

	ToSystemReplacemsgAlertmailMapOutput() SystemReplacemsgAlertmailMapOutput
	ToSystemReplacemsgAlertmailMapOutputWithContext(context.Context) SystemReplacemsgAlertmailMapOutput
}

type SystemReplacemsgAlertmailMap map[string]SystemReplacemsgAlertmailInput

func (SystemReplacemsgAlertmailMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgAlertmail)(nil)).Elem()
}

func (i SystemReplacemsgAlertmailMap) ToSystemReplacemsgAlertmailMapOutput() SystemReplacemsgAlertmailMapOutput {
	return i.ToSystemReplacemsgAlertmailMapOutputWithContext(context.Background())
}

func (i SystemReplacemsgAlertmailMap) ToSystemReplacemsgAlertmailMapOutputWithContext(ctx context.Context) SystemReplacemsgAlertmailMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemReplacemsgAlertmailMapOutput)
}

type SystemReplacemsgAlertmailOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgAlertmailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemReplacemsgAlertmail)(nil)).Elem()
}

func (o SystemReplacemsgAlertmailOutput) ToSystemReplacemsgAlertmailOutput() SystemReplacemsgAlertmailOutput {
	return o
}

func (o SystemReplacemsgAlertmailOutput) ToSystemReplacemsgAlertmailOutputWithContext(ctx context.Context) SystemReplacemsgAlertmailOutput {
	return o
}

type SystemReplacemsgAlertmailArrayOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgAlertmailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemReplacemsgAlertmail)(nil)).Elem()
}

func (o SystemReplacemsgAlertmailArrayOutput) ToSystemReplacemsgAlertmailArrayOutput() SystemReplacemsgAlertmailArrayOutput {
	return o
}

func (o SystemReplacemsgAlertmailArrayOutput) ToSystemReplacemsgAlertmailArrayOutputWithContext(ctx context.Context) SystemReplacemsgAlertmailArrayOutput {
	return o
}

func (o SystemReplacemsgAlertmailArrayOutput) Index(i pulumi.IntInput) SystemReplacemsgAlertmailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemReplacemsgAlertmail {
		return vs[0].([]*SystemReplacemsgAlertmail)[vs[1].(int)]
	}).(SystemReplacemsgAlertmailOutput)
}

type SystemReplacemsgAlertmailMapOutput struct{ *pulumi.OutputState }

func (SystemReplacemsgAlertmailMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemReplacemsgAlertmail)(nil)).Elem()
}

func (o SystemReplacemsgAlertmailMapOutput) ToSystemReplacemsgAlertmailMapOutput() SystemReplacemsgAlertmailMapOutput {
	return o
}

func (o SystemReplacemsgAlertmailMapOutput) ToSystemReplacemsgAlertmailMapOutputWithContext(ctx context.Context) SystemReplacemsgAlertmailMapOutput {
	return o
}

func (o SystemReplacemsgAlertmailMapOutput) MapIndex(k pulumi.StringInput) SystemReplacemsgAlertmailOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemReplacemsgAlertmail {
		return vs[0].(map[string]*SystemReplacemsgAlertmail)[vs[1].(string)]
	}).(SystemReplacemsgAlertmailOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgAlertmailInput)(nil)).Elem(), &SystemReplacemsgAlertmail{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgAlertmailArrayInput)(nil)).Elem(), SystemReplacemsgAlertmailArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemReplacemsgAlertmailMapInput)(nil)).Elem(), SystemReplacemsgAlertmailMap{})
	pulumi.RegisterOutputType(SystemReplacemsgAlertmailOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgAlertmailArrayOutput{})
	pulumi.RegisterOutputType(SystemReplacemsgAlertmailMapOutput{})
}
