// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type FirewallAccessProxy struct {
	pulumi.CustomResourceState

	AddVhostDomainToDnsdb             pulumi.StringOutput                       `pulumi:"addVhostDomainToDnsdb"`
	ApiGateway6s                      FirewallAccessProxyApiGateway6ArrayOutput `pulumi:"apiGateway6s"`
	ApiGateways                       FirewallAccessProxyApiGatewayArrayOutput  `pulumi:"apiGateways"`
	AuthPortal                        pulumi.StringOutput                       `pulumi:"authPortal"`
	AuthVirtualHost                   pulumi.StringOutput                       `pulumi:"authVirtualHost"`
	ClientCert                        pulumi.StringOutput                       `pulumi:"clientCert"`
	DecryptedTrafficMirror            pulumi.StringOutput                       `pulumi:"decryptedTrafficMirror"`
	DynamicSortSubtable               pulumi.StringPtrOutput                    `pulumi:"dynamicSortSubtable"`
	EmptyCertAction                   pulumi.StringOutput                       `pulumi:"emptyCertAction"`
	GetAllTables                      pulumi.StringPtrOutput                    `pulumi:"getAllTables"`
	HttpSupportedMaxVersion           pulumi.StringOutput                       `pulumi:"httpSupportedMaxVersion"`
	LogBlockedTraffic                 pulumi.StringOutput                       `pulumi:"logBlockedTraffic"`
	Name                              pulumi.StringOutput                       `pulumi:"name"`
	SvrPoolMultiplex                  pulumi.StringOutput                       `pulumi:"svrPoolMultiplex"`
	SvrPoolServerMaxConcurrentRequest pulumi.IntOutput                          `pulumi:"svrPoolServerMaxConcurrentRequest"`
	SvrPoolServerMaxRequest           pulumi.IntOutput                          `pulumi:"svrPoolServerMaxRequest"`
	SvrPoolTtl                        pulumi.IntOutput                          `pulumi:"svrPoolTtl"`
	UserAgentDetect                   pulumi.StringOutput                       `pulumi:"userAgentDetect"`
	Vdomparam                         pulumi.StringPtrOutput                    `pulumi:"vdomparam"`
	Vip                               pulumi.StringOutput                       `pulumi:"vip"`
}

// NewFirewallAccessProxy registers a new resource with the given unique name, arguments, and options.
func NewFirewallAccessProxy(ctx *pulumi.Context,
	name string, args *FirewallAccessProxyArgs, opts ...pulumi.ResourceOption) (*FirewallAccessProxy, error) {
	if args == nil {
		args = &FirewallAccessProxyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallAccessProxy
	err := ctx.RegisterResource("fortios:index/firewallAccessProxy:FirewallAccessProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallAccessProxy gets an existing FirewallAccessProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallAccessProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallAccessProxyState, opts ...pulumi.ResourceOption) (*FirewallAccessProxy, error) {
	var resource FirewallAccessProxy
	err := ctx.ReadResource("fortios:index/firewallAccessProxy:FirewallAccessProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallAccessProxy resources.
type firewallAccessProxyState struct {
	AddVhostDomainToDnsdb             *string                          `pulumi:"addVhostDomainToDnsdb"`
	ApiGateway6s                      []FirewallAccessProxyApiGateway6 `pulumi:"apiGateway6s"`
	ApiGateways                       []FirewallAccessProxyApiGateway  `pulumi:"apiGateways"`
	AuthPortal                        *string                          `pulumi:"authPortal"`
	AuthVirtualHost                   *string                          `pulumi:"authVirtualHost"`
	ClientCert                        *string                          `pulumi:"clientCert"`
	DecryptedTrafficMirror            *string                          `pulumi:"decryptedTrafficMirror"`
	DynamicSortSubtable               *string                          `pulumi:"dynamicSortSubtable"`
	EmptyCertAction                   *string                          `pulumi:"emptyCertAction"`
	GetAllTables                      *string                          `pulumi:"getAllTables"`
	HttpSupportedMaxVersion           *string                          `pulumi:"httpSupportedMaxVersion"`
	LogBlockedTraffic                 *string                          `pulumi:"logBlockedTraffic"`
	Name                              *string                          `pulumi:"name"`
	SvrPoolMultiplex                  *string                          `pulumi:"svrPoolMultiplex"`
	SvrPoolServerMaxConcurrentRequest *int                             `pulumi:"svrPoolServerMaxConcurrentRequest"`
	SvrPoolServerMaxRequest           *int                             `pulumi:"svrPoolServerMaxRequest"`
	SvrPoolTtl                        *int                             `pulumi:"svrPoolTtl"`
	UserAgentDetect                   *string                          `pulumi:"userAgentDetect"`
	Vdomparam                         *string                          `pulumi:"vdomparam"`
	Vip                               *string                          `pulumi:"vip"`
}

type FirewallAccessProxyState struct {
	AddVhostDomainToDnsdb             pulumi.StringPtrInput
	ApiGateway6s                      FirewallAccessProxyApiGateway6ArrayInput
	ApiGateways                       FirewallAccessProxyApiGatewayArrayInput
	AuthPortal                        pulumi.StringPtrInput
	AuthVirtualHost                   pulumi.StringPtrInput
	ClientCert                        pulumi.StringPtrInput
	DecryptedTrafficMirror            pulumi.StringPtrInput
	DynamicSortSubtable               pulumi.StringPtrInput
	EmptyCertAction                   pulumi.StringPtrInput
	GetAllTables                      pulumi.StringPtrInput
	HttpSupportedMaxVersion           pulumi.StringPtrInput
	LogBlockedTraffic                 pulumi.StringPtrInput
	Name                              pulumi.StringPtrInput
	SvrPoolMultiplex                  pulumi.StringPtrInput
	SvrPoolServerMaxConcurrentRequest pulumi.IntPtrInput
	SvrPoolServerMaxRequest           pulumi.IntPtrInput
	SvrPoolTtl                        pulumi.IntPtrInput
	UserAgentDetect                   pulumi.StringPtrInput
	Vdomparam                         pulumi.StringPtrInput
	Vip                               pulumi.StringPtrInput
}

func (FirewallAccessProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallAccessProxyState)(nil)).Elem()
}

type firewallAccessProxyArgs struct {
	AddVhostDomainToDnsdb             *string                          `pulumi:"addVhostDomainToDnsdb"`
	ApiGateway6s                      []FirewallAccessProxyApiGateway6 `pulumi:"apiGateway6s"`
	ApiGateways                       []FirewallAccessProxyApiGateway  `pulumi:"apiGateways"`
	AuthPortal                        *string                          `pulumi:"authPortal"`
	AuthVirtualHost                   *string                          `pulumi:"authVirtualHost"`
	ClientCert                        *string                          `pulumi:"clientCert"`
	DecryptedTrafficMirror            *string                          `pulumi:"decryptedTrafficMirror"`
	DynamicSortSubtable               *string                          `pulumi:"dynamicSortSubtable"`
	EmptyCertAction                   *string                          `pulumi:"emptyCertAction"`
	GetAllTables                      *string                          `pulumi:"getAllTables"`
	HttpSupportedMaxVersion           *string                          `pulumi:"httpSupportedMaxVersion"`
	LogBlockedTraffic                 *string                          `pulumi:"logBlockedTraffic"`
	Name                              *string                          `pulumi:"name"`
	SvrPoolMultiplex                  *string                          `pulumi:"svrPoolMultiplex"`
	SvrPoolServerMaxConcurrentRequest *int                             `pulumi:"svrPoolServerMaxConcurrentRequest"`
	SvrPoolServerMaxRequest           *int                             `pulumi:"svrPoolServerMaxRequest"`
	SvrPoolTtl                        *int                             `pulumi:"svrPoolTtl"`
	UserAgentDetect                   *string                          `pulumi:"userAgentDetect"`
	Vdomparam                         *string                          `pulumi:"vdomparam"`
	Vip                               *string                          `pulumi:"vip"`
}

// The set of arguments for constructing a FirewallAccessProxy resource.
type FirewallAccessProxyArgs struct {
	AddVhostDomainToDnsdb             pulumi.StringPtrInput
	ApiGateway6s                      FirewallAccessProxyApiGateway6ArrayInput
	ApiGateways                       FirewallAccessProxyApiGatewayArrayInput
	AuthPortal                        pulumi.StringPtrInput
	AuthVirtualHost                   pulumi.StringPtrInput
	ClientCert                        pulumi.StringPtrInput
	DecryptedTrafficMirror            pulumi.StringPtrInput
	DynamicSortSubtable               pulumi.StringPtrInput
	EmptyCertAction                   pulumi.StringPtrInput
	GetAllTables                      pulumi.StringPtrInput
	HttpSupportedMaxVersion           pulumi.StringPtrInput
	LogBlockedTraffic                 pulumi.StringPtrInput
	Name                              pulumi.StringPtrInput
	SvrPoolMultiplex                  pulumi.StringPtrInput
	SvrPoolServerMaxConcurrentRequest pulumi.IntPtrInput
	SvrPoolServerMaxRequest           pulumi.IntPtrInput
	SvrPoolTtl                        pulumi.IntPtrInput
	UserAgentDetect                   pulumi.StringPtrInput
	Vdomparam                         pulumi.StringPtrInput
	Vip                               pulumi.StringPtrInput
}

func (FirewallAccessProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallAccessProxyArgs)(nil)).Elem()
}

type FirewallAccessProxyInput interface {
	pulumi.Input

	ToFirewallAccessProxyOutput() FirewallAccessProxyOutput
	ToFirewallAccessProxyOutputWithContext(ctx context.Context) FirewallAccessProxyOutput
}

func (*FirewallAccessProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallAccessProxy)(nil)).Elem()
}

func (i *FirewallAccessProxy) ToFirewallAccessProxyOutput() FirewallAccessProxyOutput {
	return i.ToFirewallAccessProxyOutputWithContext(context.Background())
}

func (i *FirewallAccessProxy) ToFirewallAccessProxyOutputWithContext(ctx context.Context) FirewallAccessProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAccessProxyOutput)
}

func (i *FirewallAccessProxy) ToOutput(ctx context.Context) pulumix.Output[*FirewallAccessProxy] {
	return pulumix.Output[*FirewallAccessProxy]{
		OutputState: i.ToFirewallAccessProxyOutputWithContext(ctx).OutputState,
	}
}

// FirewallAccessProxyArrayInput is an input type that accepts FirewallAccessProxyArray and FirewallAccessProxyArrayOutput values.
// You can construct a concrete instance of `FirewallAccessProxyArrayInput` via:
//
//	FirewallAccessProxyArray{ FirewallAccessProxyArgs{...} }
type FirewallAccessProxyArrayInput interface {
	pulumi.Input

	ToFirewallAccessProxyArrayOutput() FirewallAccessProxyArrayOutput
	ToFirewallAccessProxyArrayOutputWithContext(context.Context) FirewallAccessProxyArrayOutput
}

type FirewallAccessProxyArray []FirewallAccessProxyInput

func (FirewallAccessProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallAccessProxy)(nil)).Elem()
}

func (i FirewallAccessProxyArray) ToFirewallAccessProxyArrayOutput() FirewallAccessProxyArrayOutput {
	return i.ToFirewallAccessProxyArrayOutputWithContext(context.Background())
}

func (i FirewallAccessProxyArray) ToFirewallAccessProxyArrayOutputWithContext(ctx context.Context) FirewallAccessProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAccessProxyArrayOutput)
}

func (i FirewallAccessProxyArray) ToOutput(ctx context.Context) pulumix.Output[[]*FirewallAccessProxy] {
	return pulumix.Output[[]*FirewallAccessProxy]{
		OutputState: i.ToFirewallAccessProxyArrayOutputWithContext(ctx).OutputState,
	}
}

// FirewallAccessProxyMapInput is an input type that accepts FirewallAccessProxyMap and FirewallAccessProxyMapOutput values.
// You can construct a concrete instance of `FirewallAccessProxyMapInput` via:
//
//	FirewallAccessProxyMap{ "key": FirewallAccessProxyArgs{...} }
type FirewallAccessProxyMapInput interface {
	pulumi.Input

	ToFirewallAccessProxyMapOutput() FirewallAccessProxyMapOutput
	ToFirewallAccessProxyMapOutputWithContext(context.Context) FirewallAccessProxyMapOutput
}

type FirewallAccessProxyMap map[string]FirewallAccessProxyInput

func (FirewallAccessProxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallAccessProxy)(nil)).Elem()
}

func (i FirewallAccessProxyMap) ToFirewallAccessProxyMapOutput() FirewallAccessProxyMapOutput {
	return i.ToFirewallAccessProxyMapOutputWithContext(context.Background())
}

func (i FirewallAccessProxyMap) ToFirewallAccessProxyMapOutputWithContext(ctx context.Context) FirewallAccessProxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAccessProxyMapOutput)
}

func (i FirewallAccessProxyMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FirewallAccessProxy] {
	return pulumix.Output[map[string]*FirewallAccessProxy]{
		OutputState: i.ToFirewallAccessProxyMapOutputWithContext(ctx).OutputState,
	}
}

type FirewallAccessProxyOutput struct{ *pulumi.OutputState }

func (FirewallAccessProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallAccessProxy)(nil)).Elem()
}

func (o FirewallAccessProxyOutput) ToFirewallAccessProxyOutput() FirewallAccessProxyOutput {
	return o
}

func (o FirewallAccessProxyOutput) ToFirewallAccessProxyOutputWithContext(ctx context.Context) FirewallAccessProxyOutput {
	return o
}

func (o FirewallAccessProxyOutput) ToOutput(ctx context.Context) pulumix.Output[*FirewallAccessProxy] {
	return pulumix.Output[*FirewallAccessProxy]{
		OutputState: o.OutputState,
	}
}

func (o FirewallAccessProxyOutput) AddVhostDomainToDnsdb() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.StringOutput { return v.AddVhostDomainToDnsdb }).(pulumi.StringOutput)
}

func (o FirewallAccessProxyOutput) ApiGateway6s() FirewallAccessProxyApiGateway6ArrayOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) FirewallAccessProxyApiGateway6ArrayOutput { return v.ApiGateway6s }).(FirewallAccessProxyApiGateway6ArrayOutput)
}

func (o FirewallAccessProxyOutput) ApiGateways() FirewallAccessProxyApiGatewayArrayOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) FirewallAccessProxyApiGatewayArrayOutput { return v.ApiGateways }).(FirewallAccessProxyApiGatewayArrayOutput)
}

func (o FirewallAccessProxyOutput) AuthPortal() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.StringOutput { return v.AuthPortal }).(pulumi.StringOutput)
}

func (o FirewallAccessProxyOutput) AuthVirtualHost() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.StringOutput { return v.AuthVirtualHost }).(pulumi.StringOutput)
}

func (o FirewallAccessProxyOutput) ClientCert() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.StringOutput { return v.ClientCert }).(pulumi.StringOutput)
}

func (o FirewallAccessProxyOutput) DecryptedTrafficMirror() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.StringOutput { return v.DecryptedTrafficMirror }).(pulumi.StringOutput)
}

func (o FirewallAccessProxyOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o FirewallAccessProxyOutput) EmptyCertAction() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.StringOutput { return v.EmptyCertAction }).(pulumi.StringOutput)
}

func (o FirewallAccessProxyOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o FirewallAccessProxyOutput) HttpSupportedMaxVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.StringOutput { return v.HttpSupportedMaxVersion }).(pulumi.StringOutput)
}

func (o FirewallAccessProxyOutput) LogBlockedTraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.StringOutput { return v.LogBlockedTraffic }).(pulumi.StringOutput)
}

func (o FirewallAccessProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallAccessProxyOutput) SvrPoolMultiplex() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.StringOutput { return v.SvrPoolMultiplex }).(pulumi.StringOutput)
}

func (o FirewallAccessProxyOutput) SvrPoolServerMaxConcurrentRequest() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.IntOutput { return v.SvrPoolServerMaxConcurrentRequest }).(pulumi.IntOutput)
}

func (o FirewallAccessProxyOutput) SvrPoolServerMaxRequest() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.IntOutput { return v.SvrPoolServerMaxRequest }).(pulumi.IntOutput)
}

func (o FirewallAccessProxyOutput) SvrPoolTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.IntOutput { return v.SvrPoolTtl }).(pulumi.IntOutput)
}

func (o FirewallAccessProxyOutput) UserAgentDetect() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.StringOutput { return v.UserAgentDetect }).(pulumi.StringOutput)
}

func (o FirewallAccessProxyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o FirewallAccessProxyOutput) Vip() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAccessProxy) pulumi.StringOutput { return v.Vip }).(pulumi.StringOutput)
}

type FirewallAccessProxyArrayOutput struct{ *pulumi.OutputState }

func (FirewallAccessProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallAccessProxy)(nil)).Elem()
}

func (o FirewallAccessProxyArrayOutput) ToFirewallAccessProxyArrayOutput() FirewallAccessProxyArrayOutput {
	return o
}

func (o FirewallAccessProxyArrayOutput) ToFirewallAccessProxyArrayOutputWithContext(ctx context.Context) FirewallAccessProxyArrayOutput {
	return o
}

func (o FirewallAccessProxyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FirewallAccessProxy] {
	return pulumix.Output[[]*FirewallAccessProxy]{
		OutputState: o.OutputState,
	}
}

func (o FirewallAccessProxyArrayOutput) Index(i pulumi.IntInput) FirewallAccessProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallAccessProxy {
		return vs[0].([]*FirewallAccessProxy)[vs[1].(int)]
	}).(FirewallAccessProxyOutput)
}

type FirewallAccessProxyMapOutput struct{ *pulumi.OutputState }

func (FirewallAccessProxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallAccessProxy)(nil)).Elem()
}

func (o FirewallAccessProxyMapOutput) ToFirewallAccessProxyMapOutput() FirewallAccessProxyMapOutput {
	return o
}

func (o FirewallAccessProxyMapOutput) ToFirewallAccessProxyMapOutputWithContext(ctx context.Context) FirewallAccessProxyMapOutput {
	return o
}

func (o FirewallAccessProxyMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FirewallAccessProxy] {
	return pulumix.Output[map[string]*FirewallAccessProxy]{
		OutputState: o.OutputState,
	}
}

func (o FirewallAccessProxyMapOutput) MapIndex(k pulumi.StringInput) FirewallAccessProxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallAccessProxy {
		return vs[0].(map[string]*FirewallAccessProxy)[vs[1].(string)]
	}).(FirewallAccessProxyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallAccessProxyInput)(nil)).Elem(), &FirewallAccessProxy{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallAccessProxyArrayInput)(nil)).Elem(), FirewallAccessProxyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallAccessProxyMapInput)(nil)).Elem(), FirewallAccessProxyMap{})
	pulumi.RegisterOutputType(FirewallAccessProxyOutput{})
	pulumi.RegisterOutputType(FirewallAccessProxyArrayOutput{})
	pulumi.RegisterOutputType(FirewallAccessProxyMapOutput{})
}
