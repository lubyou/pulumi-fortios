// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure SMS server for sending SMS messages to support user authentication.
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
// 		_, err := fortios.NewSystemSmsServer(ctx, "trname", &fortios.SystemSmsServerArgs{
// 			MailServer: pulumi.String("1.1.1.2"),
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
// System SmsServer can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemSmsServer:SystemSmsServer labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemSmsServer struct {
	pulumi.CustomResourceState

	// Email-to-SMS server domain name.
	MailServer pulumi.StringOutput `pulumi:"mailServer"`
	// Name of SMS server.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemSmsServer registers a new resource with the given unique name, arguments, and options.
func NewSystemSmsServer(ctx *pulumi.Context,
	name string, args *SystemSmsServerArgs, opts ...pulumi.ResourceOption) (*SystemSmsServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MailServer == nil {
		return nil, errors.New("invalid value for required argument 'MailServer'")
	}
	var resource SystemSmsServer
	err := ctx.RegisterResource("fortios:index/systemSmsServer:SystemSmsServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSmsServer gets an existing SystemSmsServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSmsServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSmsServerState, opts ...pulumi.ResourceOption) (*SystemSmsServer, error) {
	var resource SystemSmsServer
	err := ctx.ReadResource("fortios:index/systemSmsServer:SystemSmsServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSmsServer resources.
type systemSmsServerState struct {
	// Email-to-SMS server domain name.
	MailServer *string `pulumi:"mailServer"`
	// Name of SMS server.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemSmsServerState struct {
	// Email-to-SMS server domain name.
	MailServer pulumi.StringPtrInput
	// Name of SMS server.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSmsServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSmsServerState)(nil)).Elem()
}

type systemSmsServerArgs struct {
	// Email-to-SMS server domain name.
	MailServer string `pulumi:"mailServer"`
	// Name of SMS server.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemSmsServer resource.
type SystemSmsServerArgs struct {
	// Email-to-SMS server domain name.
	MailServer pulumi.StringInput
	// Name of SMS server.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemSmsServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSmsServerArgs)(nil)).Elem()
}

type SystemSmsServerInput interface {
	pulumi.Input

	ToSystemSmsServerOutput() SystemSmsServerOutput
	ToSystemSmsServerOutputWithContext(ctx context.Context) SystemSmsServerOutput
}

func (*SystemSmsServer) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemSmsServer)(nil))
}

func (i *SystemSmsServer) ToSystemSmsServerOutput() SystemSmsServerOutput {
	return i.ToSystemSmsServerOutputWithContext(context.Background())
}

func (i *SystemSmsServer) ToSystemSmsServerOutputWithContext(ctx context.Context) SystemSmsServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSmsServerOutput)
}

func (i *SystemSmsServer) ToSystemSmsServerPtrOutput() SystemSmsServerPtrOutput {
	return i.ToSystemSmsServerPtrOutputWithContext(context.Background())
}

func (i *SystemSmsServer) ToSystemSmsServerPtrOutputWithContext(ctx context.Context) SystemSmsServerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSmsServerPtrOutput)
}

type SystemSmsServerPtrInput interface {
	pulumi.Input

	ToSystemSmsServerPtrOutput() SystemSmsServerPtrOutput
	ToSystemSmsServerPtrOutputWithContext(ctx context.Context) SystemSmsServerPtrOutput
}

type systemSmsServerPtrType SystemSmsServerArgs

func (*systemSmsServerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSmsServer)(nil))
}

func (i *systemSmsServerPtrType) ToSystemSmsServerPtrOutput() SystemSmsServerPtrOutput {
	return i.ToSystemSmsServerPtrOutputWithContext(context.Background())
}

func (i *systemSmsServerPtrType) ToSystemSmsServerPtrOutputWithContext(ctx context.Context) SystemSmsServerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSmsServerPtrOutput)
}

// SystemSmsServerArrayInput is an input type that accepts SystemSmsServerArray and SystemSmsServerArrayOutput values.
// You can construct a concrete instance of `SystemSmsServerArrayInput` via:
//
//          SystemSmsServerArray{ SystemSmsServerArgs{...} }
type SystemSmsServerArrayInput interface {
	pulumi.Input

	ToSystemSmsServerArrayOutput() SystemSmsServerArrayOutput
	ToSystemSmsServerArrayOutputWithContext(context.Context) SystemSmsServerArrayOutput
}

type SystemSmsServerArray []SystemSmsServerInput

func (SystemSmsServerArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SystemSmsServer)(nil))
}

func (i SystemSmsServerArray) ToSystemSmsServerArrayOutput() SystemSmsServerArrayOutput {
	return i.ToSystemSmsServerArrayOutputWithContext(context.Background())
}

func (i SystemSmsServerArray) ToSystemSmsServerArrayOutputWithContext(ctx context.Context) SystemSmsServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSmsServerArrayOutput)
}

// SystemSmsServerMapInput is an input type that accepts SystemSmsServerMap and SystemSmsServerMapOutput values.
// You can construct a concrete instance of `SystemSmsServerMapInput` via:
//
//          SystemSmsServerMap{ "key": SystemSmsServerArgs{...} }
type SystemSmsServerMapInput interface {
	pulumi.Input

	ToSystemSmsServerMapOutput() SystemSmsServerMapOutput
	ToSystemSmsServerMapOutputWithContext(context.Context) SystemSmsServerMapOutput
}

type SystemSmsServerMap map[string]SystemSmsServerInput

func (SystemSmsServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SystemSmsServer)(nil))
}

func (i SystemSmsServerMap) ToSystemSmsServerMapOutput() SystemSmsServerMapOutput {
	return i.ToSystemSmsServerMapOutputWithContext(context.Background())
}

func (i SystemSmsServerMap) ToSystemSmsServerMapOutputWithContext(ctx context.Context) SystemSmsServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSmsServerMapOutput)
}

type SystemSmsServerOutput struct {
	*pulumi.OutputState
}

func (SystemSmsServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemSmsServer)(nil))
}

func (o SystemSmsServerOutput) ToSystemSmsServerOutput() SystemSmsServerOutput {
	return o
}

func (o SystemSmsServerOutput) ToSystemSmsServerOutputWithContext(ctx context.Context) SystemSmsServerOutput {
	return o
}

func (o SystemSmsServerOutput) ToSystemSmsServerPtrOutput() SystemSmsServerPtrOutput {
	return o.ToSystemSmsServerPtrOutputWithContext(context.Background())
}

func (o SystemSmsServerOutput) ToSystemSmsServerPtrOutputWithContext(ctx context.Context) SystemSmsServerPtrOutput {
	return o.ApplyT(func(v SystemSmsServer) *SystemSmsServer {
		return &v
	}).(SystemSmsServerPtrOutput)
}

type SystemSmsServerPtrOutput struct {
	*pulumi.OutputState
}

func (SystemSmsServerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSmsServer)(nil))
}

func (o SystemSmsServerPtrOutput) ToSystemSmsServerPtrOutput() SystemSmsServerPtrOutput {
	return o
}

func (o SystemSmsServerPtrOutput) ToSystemSmsServerPtrOutputWithContext(ctx context.Context) SystemSmsServerPtrOutput {
	return o
}

type SystemSmsServerArrayOutput struct{ *pulumi.OutputState }

func (SystemSmsServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemSmsServer)(nil))
}

func (o SystemSmsServerArrayOutput) ToSystemSmsServerArrayOutput() SystemSmsServerArrayOutput {
	return o
}

func (o SystemSmsServerArrayOutput) ToSystemSmsServerArrayOutputWithContext(ctx context.Context) SystemSmsServerArrayOutput {
	return o
}

func (o SystemSmsServerArrayOutput) Index(i pulumi.IntInput) SystemSmsServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemSmsServer {
		return vs[0].([]SystemSmsServer)[vs[1].(int)]
	}).(SystemSmsServerOutput)
}

type SystemSmsServerMapOutput struct{ *pulumi.OutputState }

func (SystemSmsServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SystemSmsServer)(nil))
}

func (o SystemSmsServerMapOutput) ToSystemSmsServerMapOutput() SystemSmsServerMapOutput {
	return o
}

func (o SystemSmsServerMapOutput) ToSystemSmsServerMapOutputWithContext(ctx context.Context) SystemSmsServerMapOutput {
	return o
}

func (o SystemSmsServerMapOutput) MapIndex(k pulumi.StringInput) SystemSmsServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SystemSmsServer {
		return vs[0].(map[string]SystemSmsServer)[vs[1].(string)]
	}).(SystemSmsServerOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemSmsServerOutput{})
	pulumi.RegisterOutputType(SystemSmsServerPtrOutput{})
	pulumi.RegisterOutputType(SystemSmsServerArrayOutput{})
	pulumi.RegisterOutputType(SystemSmsServerMapOutput{})
}