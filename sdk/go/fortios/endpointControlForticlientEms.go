// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiClient Enterprise Management Server (EMS) entries. Applies to FortiOS Version `<= 6.2.0`.
//
// ## Import
//
// EndpointControl ForticlientEms can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/endpointControlForticlientEms:EndpointControlForticlientEms labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type EndpointControlForticlientEms struct {
	pulumi.CustomResourceState

	// Firewall address name.
	Address pulumi.StringOutput `pulumi:"address"`
	// FortiClient EMS admin password.
	AdminPassword pulumi.StringPtrOutput `pulumi:"adminPassword"`
	// FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
	AdminType pulumi.StringOutput `pulumi:"adminType"`
	// FortiClient EMS admin username.
	AdminUsername pulumi.StringOutput `pulumi:"adminUsername"`
	// FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
	HttpsPort pulumi.IntOutput `pulumi:"httpsPort"`
	// FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
	ListenPort pulumi.IntOutput `pulumi:"listenPort"`
	// FortiClient Enterprise Management Server (EMS) name.
	Name pulumi.StringOutput `pulumi:"name"`
	// FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
	RestApiAuth pulumi.StringOutput `pulumi:"restApiAuth"`
	// FortiClient EMS Serial Number.
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
	UploadPort pulumi.IntOutput `pulumi:"uploadPort"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewEndpointControlForticlientEms registers a new resource with the given unique name, arguments, and options.
func NewEndpointControlForticlientEms(ctx *pulumi.Context,
	name string, args *EndpointControlForticlientEmsArgs, opts ...pulumi.ResourceOption) (*EndpointControlForticlientEms, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.AdminUsername == nil {
		return nil, errors.New("invalid value for required argument 'AdminUsername'")
	}
	if args.SerialNumber == nil {
		return nil, errors.New("invalid value for required argument 'SerialNumber'")
	}
	var resource EndpointControlForticlientEms
	err := ctx.RegisterResource("fortios:index/endpointControlForticlientEms:EndpointControlForticlientEms", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointControlForticlientEms gets an existing EndpointControlForticlientEms resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointControlForticlientEms(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointControlForticlientEmsState, opts ...pulumi.ResourceOption) (*EndpointControlForticlientEms, error) {
	var resource EndpointControlForticlientEms
	err := ctx.ReadResource("fortios:index/endpointControlForticlientEms:EndpointControlForticlientEms", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointControlForticlientEms resources.
type endpointControlForticlientEmsState struct {
	// Firewall address name.
	Address *string `pulumi:"address"`
	// FortiClient EMS admin password.
	AdminPassword *string `pulumi:"adminPassword"`
	// FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
	AdminType *string `pulumi:"adminType"`
	// FortiClient EMS admin username.
	AdminUsername *string `pulumi:"adminUsername"`
	// FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
	HttpsPort *int `pulumi:"httpsPort"`
	// FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
	ListenPort *int `pulumi:"listenPort"`
	// FortiClient Enterprise Management Server (EMS) name.
	Name *string `pulumi:"name"`
	// FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
	RestApiAuth *string `pulumi:"restApiAuth"`
	// FortiClient EMS Serial Number.
	SerialNumber *string `pulumi:"serialNumber"`
	// FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
	UploadPort *int `pulumi:"uploadPort"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type EndpointControlForticlientEmsState struct {
	// Firewall address name.
	Address pulumi.StringPtrInput
	// FortiClient EMS admin password.
	AdminPassword pulumi.StringPtrInput
	// FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
	AdminType pulumi.StringPtrInput
	// FortiClient EMS admin username.
	AdminUsername pulumi.StringPtrInput
	// FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
	HttpsPort pulumi.IntPtrInput
	// FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
	ListenPort pulumi.IntPtrInput
	// FortiClient Enterprise Management Server (EMS) name.
	Name pulumi.StringPtrInput
	// FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
	RestApiAuth pulumi.StringPtrInput
	// FortiClient EMS Serial Number.
	SerialNumber pulumi.StringPtrInput
	// FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
	UploadPort pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EndpointControlForticlientEmsState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointControlForticlientEmsState)(nil)).Elem()
}

type endpointControlForticlientEmsArgs struct {
	// Firewall address name.
	Address string `pulumi:"address"`
	// FortiClient EMS admin password.
	AdminPassword *string `pulumi:"adminPassword"`
	// FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
	AdminType *string `pulumi:"adminType"`
	// FortiClient EMS admin username.
	AdminUsername string `pulumi:"adminUsername"`
	// FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
	HttpsPort *int `pulumi:"httpsPort"`
	// FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
	ListenPort *int `pulumi:"listenPort"`
	// FortiClient Enterprise Management Server (EMS) name.
	Name *string `pulumi:"name"`
	// FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
	RestApiAuth *string `pulumi:"restApiAuth"`
	// FortiClient EMS Serial Number.
	SerialNumber string `pulumi:"serialNumber"`
	// FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
	UploadPort *int `pulumi:"uploadPort"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a EndpointControlForticlientEms resource.
type EndpointControlForticlientEmsArgs struct {
	// Firewall address name.
	Address pulumi.StringInput
	// FortiClient EMS admin password.
	AdminPassword pulumi.StringPtrInput
	// FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
	AdminType pulumi.StringPtrInput
	// FortiClient EMS admin username.
	AdminUsername pulumi.StringInput
	// FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
	HttpsPort pulumi.IntPtrInput
	// FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
	ListenPort pulumi.IntPtrInput
	// FortiClient Enterprise Management Server (EMS) name.
	Name pulumi.StringPtrInput
	// FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
	RestApiAuth pulumi.StringPtrInput
	// FortiClient EMS Serial Number.
	SerialNumber pulumi.StringInput
	// FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
	UploadPort pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EndpointControlForticlientEmsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointControlForticlientEmsArgs)(nil)).Elem()
}

type EndpointControlForticlientEmsInput interface {
	pulumi.Input

	ToEndpointControlForticlientEmsOutput() EndpointControlForticlientEmsOutput
	ToEndpointControlForticlientEmsOutputWithContext(ctx context.Context) EndpointControlForticlientEmsOutput
}

func (*EndpointControlForticlientEms) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointControlForticlientEms)(nil))
}

func (i *EndpointControlForticlientEms) ToEndpointControlForticlientEmsOutput() EndpointControlForticlientEmsOutput {
	return i.ToEndpointControlForticlientEmsOutputWithContext(context.Background())
}

func (i *EndpointControlForticlientEms) ToEndpointControlForticlientEmsOutputWithContext(ctx context.Context) EndpointControlForticlientEmsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlForticlientEmsOutput)
}

func (i *EndpointControlForticlientEms) ToEndpointControlForticlientEmsPtrOutput() EndpointControlForticlientEmsPtrOutput {
	return i.ToEndpointControlForticlientEmsPtrOutputWithContext(context.Background())
}

func (i *EndpointControlForticlientEms) ToEndpointControlForticlientEmsPtrOutputWithContext(ctx context.Context) EndpointControlForticlientEmsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlForticlientEmsPtrOutput)
}

type EndpointControlForticlientEmsPtrInput interface {
	pulumi.Input

	ToEndpointControlForticlientEmsPtrOutput() EndpointControlForticlientEmsPtrOutput
	ToEndpointControlForticlientEmsPtrOutputWithContext(ctx context.Context) EndpointControlForticlientEmsPtrOutput
}

type endpointControlForticlientEmsPtrType EndpointControlForticlientEmsArgs

func (*endpointControlForticlientEmsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointControlForticlientEms)(nil))
}

func (i *endpointControlForticlientEmsPtrType) ToEndpointControlForticlientEmsPtrOutput() EndpointControlForticlientEmsPtrOutput {
	return i.ToEndpointControlForticlientEmsPtrOutputWithContext(context.Background())
}

func (i *endpointControlForticlientEmsPtrType) ToEndpointControlForticlientEmsPtrOutputWithContext(ctx context.Context) EndpointControlForticlientEmsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlForticlientEmsPtrOutput)
}

// EndpointControlForticlientEmsArrayInput is an input type that accepts EndpointControlForticlientEmsArray and EndpointControlForticlientEmsArrayOutput values.
// You can construct a concrete instance of `EndpointControlForticlientEmsArrayInput` via:
//
//          EndpointControlForticlientEmsArray{ EndpointControlForticlientEmsArgs{...} }
type EndpointControlForticlientEmsArrayInput interface {
	pulumi.Input

	ToEndpointControlForticlientEmsArrayOutput() EndpointControlForticlientEmsArrayOutput
	ToEndpointControlForticlientEmsArrayOutputWithContext(context.Context) EndpointControlForticlientEmsArrayOutput
}

type EndpointControlForticlientEmsArray []EndpointControlForticlientEmsInput

func (EndpointControlForticlientEmsArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EndpointControlForticlientEms)(nil))
}

func (i EndpointControlForticlientEmsArray) ToEndpointControlForticlientEmsArrayOutput() EndpointControlForticlientEmsArrayOutput {
	return i.ToEndpointControlForticlientEmsArrayOutputWithContext(context.Background())
}

func (i EndpointControlForticlientEmsArray) ToEndpointControlForticlientEmsArrayOutputWithContext(ctx context.Context) EndpointControlForticlientEmsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlForticlientEmsArrayOutput)
}

// EndpointControlForticlientEmsMapInput is an input type that accepts EndpointControlForticlientEmsMap and EndpointControlForticlientEmsMapOutput values.
// You can construct a concrete instance of `EndpointControlForticlientEmsMapInput` via:
//
//          EndpointControlForticlientEmsMap{ "key": EndpointControlForticlientEmsArgs{...} }
type EndpointControlForticlientEmsMapInput interface {
	pulumi.Input

	ToEndpointControlForticlientEmsMapOutput() EndpointControlForticlientEmsMapOutput
	ToEndpointControlForticlientEmsMapOutputWithContext(context.Context) EndpointControlForticlientEmsMapOutput
}

type EndpointControlForticlientEmsMap map[string]EndpointControlForticlientEmsInput

func (EndpointControlForticlientEmsMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EndpointControlForticlientEms)(nil))
}

func (i EndpointControlForticlientEmsMap) ToEndpointControlForticlientEmsMapOutput() EndpointControlForticlientEmsMapOutput {
	return i.ToEndpointControlForticlientEmsMapOutputWithContext(context.Background())
}

func (i EndpointControlForticlientEmsMap) ToEndpointControlForticlientEmsMapOutputWithContext(ctx context.Context) EndpointControlForticlientEmsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointControlForticlientEmsMapOutput)
}

type EndpointControlForticlientEmsOutput struct {
	*pulumi.OutputState
}

func (EndpointControlForticlientEmsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointControlForticlientEms)(nil))
}

func (o EndpointControlForticlientEmsOutput) ToEndpointControlForticlientEmsOutput() EndpointControlForticlientEmsOutput {
	return o
}

func (o EndpointControlForticlientEmsOutput) ToEndpointControlForticlientEmsOutputWithContext(ctx context.Context) EndpointControlForticlientEmsOutput {
	return o
}

func (o EndpointControlForticlientEmsOutput) ToEndpointControlForticlientEmsPtrOutput() EndpointControlForticlientEmsPtrOutput {
	return o.ToEndpointControlForticlientEmsPtrOutputWithContext(context.Background())
}

func (o EndpointControlForticlientEmsOutput) ToEndpointControlForticlientEmsPtrOutputWithContext(ctx context.Context) EndpointControlForticlientEmsPtrOutput {
	return o.ApplyT(func(v EndpointControlForticlientEms) *EndpointControlForticlientEms {
		return &v
	}).(EndpointControlForticlientEmsPtrOutput)
}

type EndpointControlForticlientEmsPtrOutput struct {
	*pulumi.OutputState
}

func (EndpointControlForticlientEmsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointControlForticlientEms)(nil))
}

func (o EndpointControlForticlientEmsPtrOutput) ToEndpointControlForticlientEmsPtrOutput() EndpointControlForticlientEmsPtrOutput {
	return o
}

func (o EndpointControlForticlientEmsPtrOutput) ToEndpointControlForticlientEmsPtrOutputWithContext(ctx context.Context) EndpointControlForticlientEmsPtrOutput {
	return o
}

type EndpointControlForticlientEmsArrayOutput struct{ *pulumi.OutputState }

func (EndpointControlForticlientEmsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointControlForticlientEms)(nil))
}

func (o EndpointControlForticlientEmsArrayOutput) ToEndpointControlForticlientEmsArrayOutput() EndpointControlForticlientEmsArrayOutput {
	return o
}

func (o EndpointControlForticlientEmsArrayOutput) ToEndpointControlForticlientEmsArrayOutputWithContext(ctx context.Context) EndpointControlForticlientEmsArrayOutput {
	return o
}

func (o EndpointControlForticlientEmsArrayOutput) Index(i pulumi.IntInput) EndpointControlForticlientEmsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointControlForticlientEms {
		return vs[0].([]EndpointControlForticlientEms)[vs[1].(int)]
	}).(EndpointControlForticlientEmsOutput)
}

type EndpointControlForticlientEmsMapOutput struct{ *pulumi.OutputState }

func (EndpointControlForticlientEmsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EndpointControlForticlientEms)(nil))
}

func (o EndpointControlForticlientEmsMapOutput) ToEndpointControlForticlientEmsMapOutput() EndpointControlForticlientEmsMapOutput {
	return o
}

func (o EndpointControlForticlientEmsMapOutput) ToEndpointControlForticlientEmsMapOutputWithContext(ctx context.Context) EndpointControlForticlientEmsMapOutput {
	return o
}

func (o EndpointControlForticlientEmsMapOutput) MapIndex(k pulumi.StringInput) EndpointControlForticlientEmsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EndpointControlForticlientEms {
		return vs[0].(map[string]EndpointControlForticlientEms)[vs[1].(string)]
	}).(EndpointControlForticlientEmsOutput)
}

func init() {
	pulumi.RegisterOutputType(EndpointControlForticlientEmsOutput{})
	pulumi.RegisterOutputType(EndpointControlForticlientEmsPtrOutput{})
	pulumi.RegisterOutputType(EndpointControlForticlientEmsArrayOutput{})
	pulumi.RegisterOutputType(EndpointControlForticlientEmsMapOutput{})
}
