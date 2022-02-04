// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define known domain controller servers. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0`.
//
// ## Import
//
// Cifs DomainController can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/cifsDomainController:CifsDomainController labelname {{server_name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type CifsDomainController struct {
	pulumi.CustomResourceState

	// Fully qualified domain name (FQDN). E.g. 'EXAMPLE.COM'.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// IPv4 server address.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// IPv6 server address.
	Ip6 pulumi.StringOutput `pulumi:"ip6"`
	// Password for specified username.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Port number of service. Port number 0 indicates automatic discovery.
	Port pulumi.IntOutput `pulumi:"port"`
	// Name of the server to connect to.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// User name to sign in with. Must have proper permissions for service.
	Username pulumi.StringOutput `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewCifsDomainController registers a new resource with the given unique name, arguments, and options.
func NewCifsDomainController(ctx *pulumi.Context,
	name string, args *CifsDomainControllerArgs, opts ...pulumi.ResourceOption) (*CifsDomainController, error) {
	if args == nil {
		args = &CifsDomainControllerArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource CifsDomainController
	err := ctx.RegisterResource("fortios:index/cifsDomainController:CifsDomainController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCifsDomainController gets an existing CifsDomainController resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCifsDomainController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CifsDomainControllerState, opts ...pulumi.ResourceOption) (*CifsDomainController, error) {
	var resource CifsDomainController
	err := ctx.ReadResource("fortios:index/cifsDomainController:CifsDomainController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CifsDomainController resources.
type cifsDomainControllerState struct {
	// Fully qualified domain name (FQDN). E.g. 'EXAMPLE.COM'.
	DomainName *string `pulumi:"domainName"`
	// IPv4 server address.
	Ip *string `pulumi:"ip"`
	// IPv6 server address.
	Ip6 *string `pulumi:"ip6"`
	// Password for specified username.
	Password *string `pulumi:"password"`
	// Port number of service. Port number 0 indicates automatic discovery.
	Port *int `pulumi:"port"`
	// Name of the server to connect to.
	ServerName *string `pulumi:"serverName"`
	// User name to sign in with. Must have proper permissions for service.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type CifsDomainControllerState struct {
	// Fully qualified domain name (FQDN). E.g. 'EXAMPLE.COM'.
	DomainName pulumi.StringPtrInput
	// IPv4 server address.
	Ip pulumi.StringPtrInput
	// IPv6 server address.
	Ip6 pulumi.StringPtrInput
	// Password for specified username.
	Password pulumi.StringPtrInput
	// Port number of service. Port number 0 indicates automatic discovery.
	Port pulumi.IntPtrInput
	// Name of the server to connect to.
	ServerName pulumi.StringPtrInput
	// User name to sign in with. Must have proper permissions for service.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CifsDomainControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*cifsDomainControllerState)(nil)).Elem()
}

type cifsDomainControllerArgs struct {
	// Fully qualified domain name (FQDN). E.g. 'EXAMPLE.COM'.
	DomainName *string `pulumi:"domainName"`
	// IPv4 server address.
	Ip *string `pulumi:"ip"`
	// IPv6 server address.
	Ip6 *string `pulumi:"ip6"`
	// Password for specified username.
	Password *string `pulumi:"password"`
	// Port number of service. Port number 0 indicates automatic discovery.
	Port *int `pulumi:"port"`
	// Name of the server to connect to.
	ServerName *string `pulumi:"serverName"`
	// User name to sign in with. Must have proper permissions for service.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a CifsDomainController resource.
type CifsDomainControllerArgs struct {
	// Fully qualified domain name (FQDN). E.g. 'EXAMPLE.COM'.
	DomainName pulumi.StringPtrInput
	// IPv4 server address.
	Ip pulumi.StringPtrInput
	// IPv6 server address.
	Ip6 pulumi.StringPtrInput
	// Password for specified username.
	Password pulumi.StringPtrInput
	// Port number of service. Port number 0 indicates automatic discovery.
	Port pulumi.IntPtrInput
	// Name of the server to connect to.
	ServerName pulumi.StringPtrInput
	// User name to sign in with. Must have proper permissions for service.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CifsDomainControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cifsDomainControllerArgs)(nil)).Elem()
}

type CifsDomainControllerInput interface {
	pulumi.Input

	ToCifsDomainControllerOutput() CifsDomainControllerOutput
	ToCifsDomainControllerOutputWithContext(ctx context.Context) CifsDomainControllerOutput
}

func (*CifsDomainController) ElementType() reflect.Type {
	return reflect.TypeOf((**CifsDomainController)(nil)).Elem()
}

func (i *CifsDomainController) ToCifsDomainControllerOutput() CifsDomainControllerOutput {
	return i.ToCifsDomainControllerOutputWithContext(context.Background())
}

func (i *CifsDomainController) ToCifsDomainControllerOutputWithContext(ctx context.Context) CifsDomainControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CifsDomainControllerOutput)
}

// CifsDomainControllerArrayInput is an input type that accepts CifsDomainControllerArray and CifsDomainControllerArrayOutput values.
// You can construct a concrete instance of `CifsDomainControllerArrayInput` via:
//
//          CifsDomainControllerArray{ CifsDomainControllerArgs{...} }
type CifsDomainControllerArrayInput interface {
	pulumi.Input

	ToCifsDomainControllerArrayOutput() CifsDomainControllerArrayOutput
	ToCifsDomainControllerArrayOutputWithContext(context.Context) CifsDomainControllerArrayOutput
}

type CifsDomainControllerArray []CifsDomainControllerInput

func (CifsDomainControllerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CifsDomainController)(nil)).Elem()
}

func (i CifsDomainControllerArray) ToCifsDomainControllerArrayOutput() CifsDomainControllerArrayOutput {
	return i.ToCifsDomainControllerArrayOutputWithContext(context.Background())
}

func (i CifsDomainControllerArray) ToCifsDomainControllerArrayOutputWithContext(ctx context.Context) CifsDomainControllerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CifsDomainControllerArrayOutput)
}

// CifsDomainControllerMapInput is an input type that accepts CifsDomainControllerMap and CifsDomainControllerMapOutput values.
// You can construct a concrete instance of `CifsDomainControllerMapInput` via:
//
//          CifsDomainControllerMap{ "key": CifsDomainControllerArgs{...} }
type CifsDomainControllerMapInput interface {
	pulumi.Input

	ToCifsDomainControllerMapOutput() CifsDomainControllerMapOutput
	ToCifsDomainControllerMapOutputWithContext(context.Context) CifsDomainControllerMapOutput
}

type CifsDomainControllerMap map[string]CifsDomainControllerInput

func (CifsDomainControllerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CifsDomainController)(nil)).Elem()
}

func (i CifsDomainControllerMap) ToCifsDomainControllerMapOutput() CifsDomainControllerMapOutput {
	return i.ToCifsDomainControllerMapOutputWithContext(context.Background())
}

func (i CifsDomainControllerMap) ToCifsDomainControllerMapOutputWithContext(ctx context.Context) CifsDomainControllerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CifsDomainControllerMapOutput)
}

type CifsDomainControllerOutput struct{ *pulumi.OutputState }

func (CifsDomainControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CifsDomainController)(nil)).Elem()
}

func (o CifsDomainControllerOutput) ToCifsDomainControllerOutput() CifsDomainControllerOutput {
	return o
}

func (o CifsDomainControllerOutput) ToCifsDomainControllerOutputWithContext(ctx context.Context) CifsDomainControllerOutput {
	return o
}

type CifsDomainControllerArrayOutput struct{ *pulumi.OutputState }

func (CifsDomainControllerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CifsDomainController)(nil)).Elem()
}

func (o CifsDomainControllerArrayOutput) ToCifsDomainControllerArrayOutput() CifsDomainControllerArrayOutput {
	return o
}

func (o CifsDomainControllerArrayOutput) ToCifsDomainControllerArrayOutputWithContext(ctx context.Context) CifsDomainControllerArrayOutput {
	return o
}

func (o CifsDomainControllerArrayOutput) Index(i pulumi.IntInput) CifsDomainControllerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CifsDomainController {
		return vs[0].([]*CifsDomainController)[vs[1].(int)]
	}).(CifsDomainControllerOutput)
}

type CifsDomainControllerMapOutput struct{ *pulumi.OutputState }

func (CifsDomainControllerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CifsDomainController)(nil)).Elem()
}

func (o CifsDomainControllerMapOutput) ToCifsDomainControllerMapOutput() CifsDomainControllerMapOutput {
	return o
}

func (o CifsDomainControllerMapOutput) ToCifsDomainControllerMapOutputWithContext(ctx context.Context) CifsDomainControllerMapOutput {
	return o
}

func (o CifsDomainControllerMapOutput) MapIndex(k pulumi.StringInput) CifsDomainControllerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CifsDomainController {
		return vs[0].(map[string]*CifsDomainController)[vs[1].(string)]
	}).(CifsDomainControllerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CifsDomainControllerInput)(nil)).Elem(), &CifsDomainController{})
	pulumi.RegisterInputType(reflect.TypeOf((*CifsDomainControllerArrayInput)(nil)).Elem(), CifsDomainControllerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CifsDomainControllerMapInput)(nil)).Elem(), CifsDomainControllerMap{})
	pulumi.RegisterOutputType(CifsDomainControllerOutput{})
	pulumi.RegisterOutputType(CifsDomainControllerArrayOutput{})
	pulumi.RegisterOutputType(CifsDomainControllerMapOutput{})
}
