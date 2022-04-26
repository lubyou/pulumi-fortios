// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure DNS domain filter profiles.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewDnsfilterProfile(ctx, "trname", &fortios.DnsfilterProfileArgs{
// 			BlockAction: pulumi.String("redirect"),
// 			BlockBotnet: pulumi.String("disable"),
// 			DomainFilter: &DnsfilterProfileDomainFilterArgs{
// 				DomainFilterTable: pulumi.Int(0),
// 			},
// 			FtgdDns: &DnsfilterProfileFtgdDnsArgs{
// 				Filters: DnsfilterProfileFtgdDnsFilterArray{
// 					&DnsfilterProfileFtgdDnsFilterArgs{
// 						Action:   pulumi.String("block"),
// 						Category: pulumi.Int(26),
// 						Id:       pulumi.Int(1),
// 						Log:      pulumi.String("enable"),
// 					},
// 					&DnsfilterProfileFtgdDnsFilterArgs{
// 						Action:   pulumi.String("block"),
// 						Category: pulumi.Int(61),
// 						Id:       pulumi.Int(2),
// 						Log:      pulumi.String("enable"),
// 					},
// 					&DnsfilterProfileFtgdDnsFilterArgs{
// 						Action:   pulumi.String("block"),
// 						Category: pulumi.Int(86),
// 						Id:       pulumi.Int(3),
// 						Log:      pulumi.String("enable"),
// 					},
// 					&DnsfilterProfileFtgdDnsFilterArgs{
// 						Action:   pulumi.String("block"),
// 						Category: pulumi.Int(88),
// 						Id:       pulumi.Int(4),
// 						Log:      pulumi.String("enable"),
// 					},
// 				},
// 			},
// 			LogAllDomain:    pulumi.String("disable"),
// 			RedirectPortal:  pulumi.String("0.0.0.0"),
// 			SafeSearch:      pulumi.String("disable"),
// 			SdnsDomainLog:   pulumi.String("enable"),
// 			SdnsFtgdErrLog:  pulumi.String("enable"),
// 			YoutubeRestrict: pulumi.String("strict"),
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
// Dnsfilter Profile can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/dnsfilterProfile:DnsfilterProfile labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/dnsfilterProfile:DnsfilterProfile labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type DnsfilterProfile struct {
	pulumi.CustomResourceState

	// Action to take for blocked domains.
	BlockAction pulumi.StringOutput `pulumi:"blockAction"`
	// Enable/disable blocking botnet C&C DNS lookups. Valid values: `disable`, `enable`.
	BlockBotnet pulumi.StringOutput `pulumi:"blockBotnet"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// DNS translation settings. The structure of `dnsTranslation` block is documented below.
	DnsTranslations DnsfilterProfileDnsTranslationArrayOutput `pulumi:"dnsTranslations"`
	// Domain filter settings. The structure of `domainFilter` block is documented below.
	DomainFilter DnsfilterProfileDomainFilterPtrOutput `pulumi:"domainFilter"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// One or more external IP block lists. The structure of `externalIpBlocklist` block is documented below.
	ExternalIpBlocklists DnsfilterProfileExternalIpBlocklistArrayOutput `pulumi:"externalIpBlocklists"`
	// FortiGuard DNS Filter settings. The structure of `ftgdDns` block is documented below.
	FtgdDns DnsfilterProfileFtgdDnsPtrOutput `pulumi:"ftgdDns"`
	// Enable/disable logging of all domains visited (detailed DNS logging). Valid values: `enable`, `disable`.
	LogAllDomain pulumi.StringOutput `pulumi:"logAllDomain"`
	// External domain block list name.
	Name pulumi.StringOutput `pulumi:"name"`
	// IP address of the SDNS redirect portal.
	RedirectPortal pulumi.StringOutput `pulumi:"redirectPortal"`
	// IPv6 address of the SDNS redirect portal.
	RedirectPortal6 pulumi.StringOutput `pulumi:"redirectPortal6"`
	// Enable/disable Google, Bing, and YouTube safe search. Valid values: `disable`, `enable`.
	SafeSearch pulumi.StringOutput `pulumi:"safeSearch"`
	// Enable/disable domain filtering and botnet domain logging. Valid values: `enable`, `disable`.
	SdnsDomainLog pulumi.StringOutput `pulumi:"sdnsDomainLog"`
	// Enable/disable FortiGuard SDNS rating error logging. Valid values: `enable`, `disable`.
	SdnsFtgdErrLog pulumi.StringOutput `pulumi:"sdnsFtgdErrLog"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Set safe search for YouTube restriction level. Valid values: `strict`, `moderate`.
	YoutubeRestrict pulumi.StringOutput `pulumi:"youtubeRestrict"`
}

// NewDnsfilterProfile registers a new resource with the given unique name, arguments, and options.
func NewDnsfilterProfile(ctx *pulumi.Context,
	name string, args *DnsfilterProfileArgs, opts ...pulumi.ResourceOption) (*DnsfilterProfile, error) {
	if args == nil {
		args = &DnsfilterProfileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource DnsfilterProfile
	err := ctx.RegisterResource("fortios:index/dnsfilterProfile:DnsfilterProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsfilterProfile gets an existing DnsfilterProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsfilterProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsfilterProfileState, opts ...pulumi.ResourceOption) (*DnsfilterProfile, error) {
	var resource DnsfilterProfile
	err := ctx.ReadResource("fortios:index/dnsfilterProfile:DnsfilterProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsfilterProfile resources.
type dnsfilterProfileState struct {
	// Action to take for blocked domains.
	BlockAction *string `pulumi:"blockAction"`
	// Enable/disable blocking botnet C&C DNS lookups. Valid values: `disable`, `enable`.
	BlockBotnet *string `pulumi:"blockBotnet"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// DNS translation settings. The structure of `dnsTranslation` block is documented below.
	DnsTranslations []DnsfilterProfileDnsTranslation `pulumi:"dnsTranslations"`
	// Domain filter settings. The structure of `domainFilter` block is documented below.
	DomainFilter *DnsfilterProfileDomainFilter `pulumi:"domainFilter"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// One or more external IP block lists. The structure of `externalIpBlocklist` block is documented below.
	ExternalIpBlocklists []DnsfilterProfileExternalIpBlocklist `pulumi:"externalIpBlocklists"`
	// FortiGuard DNS Filter settings. The structure of `ftgdDns` block is documented below.
	FtgdDns *DnsfilterProfileFtgdDns `pulumi:"ftgdDns"`
	// Enable/disable logging of all domains visited (detailed DNS logging). Valid values: `enable`, `disable`.
	LogAllDomain *string `pulumi:"logAllDomain"`
	// External domain block list name.
	Name *string `pulumi:"name"`
	// IP address of the SDNS redirect portal.
	RedirectPortal *string `pulumi:"redirectPortal"`
	// IPv6 address of the SDNS redirect portal.
	RedirectPortal6 *string `pulumi:"redirectPortal6"`
	// Enable/disable Google, Bing, and YouTube safe search. Valid values: `disable`, `enable`.
	SafeSearch *string `pulumi:"safeSearch"`
	// Enable/disable domain filtering and botnet domain logging. Valid values: `enable`, `disable`.
	SdnsDomainLog *string `pulumi:"sdnsDomainLog"`
	// Enable/disable FortiGuard SDNS rating error logging. Valid values: `enable`, `disable`.
	SdnsFtgdErrLog *string `pulumi:"sdnsFtgdErrLog"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Set safe search for YouTube restriction level. Valid values: `strict`, `moderate`.
	YoutubeRestrict *string `pulumi:"youtubeRestrict"`
}

type DnsfilterProfileState struct {
	// Action to take for blocked domains.
	BlockAction pulumi.StringPtrInput
	// Enable/disable blocking botnet C&C DNS lookups. Valid values: `disable`, `enable`.
	BlockBotnet pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// DNS translation settings. The structure of `dnsTranslation` block is documented below.
	DnsTranslations DnsfilterProfileDnsTranslationArrayInput
	// Domain filter settings. The structure of `domainFilter` block is documented below.
	DomainFilter DnsfilterProfileDomainFilterPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// One or more external IP block lists. The structure of `externalIpBlocklist` block is documented below.
	ExternalIpBlocklists DnsfilterProfileExternalIpBlocklistArrayInput
	// FortiGuard DNS Filter settings. The structure of `ftgdDns` block is documented below.
	FtgdDns DnsfilterProfileFtgdDnsPtrInput
	// Enable/disable logging of all domains visited (detailed DNS logging). Valid values: `enable`, `disable`.
	LogAllDomain pulumi.StringPtrInput
	// External domain block list name.
	Name pulumi.StringPtrInput
	// IP address of the SDNS redirect portal.
	RedirectPortal pulumi.StringPtrInput
	// IPv6 address of the SDNS redirect portal.
	RedirectPortal6 pulumi.StringPtrInput
	// Enable/disable Google, Bing, and YouTube safe search. Valid values: `disable`, `enable`.
	SafeSearch pulumi.StringPtrInput
	// Enable/disable domain filtering and botnet domain logging. Valid values: `enable`, `disable`.
	SdnsDomainLog pulumi.StringPtrInput
	// Enable/disable FortiGuard SDNS rating error logging. Valid values: `enable`, `disable`.
	SdnsFtgdErrLog pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Set safe search for YouTube restriction level. Valid values: `strict`, `moderate`.
	YoutubeRestrict pulumi.StringPtrInput
}

func (DnsfilterProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsfilterProfileState)(nil)).Elem()
}

type dnsfilterProfileArgs struct {
	// Action to take for blocked domains.
	BlockAction *string `pulumi:"blockAction"`
	// Enable/disable blocking botnet C&C DNS lookups. Valid values: `disable`, `enable`.
	BlockBotnet *string `pulumi:"blockBotnet"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// DNS translation settings. The structure of `dnsTranslation` block is documented below.
	DnsTranslations []DnsfilterProfileDnsTranslation `pulumi:"dnsTranslations"`
	// Domain filter settings. The structure of `domainFilter` block is documented below.
	DomainFilter *DnsfilterProfileDomainFilter `pulumi:"domainFilter"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// One or more external IP block lists. The structure of `externalIpBlocklist` block is documented below.
	ExternalIpBlocklists []DnsfilterProfileExternalIpBlocklist `pulumi:"externalIpBlocklists"`
	// FortiGuard DNS Filter settings. The structure of `ftgdDns` block is documented below.
	FtgdDns *DnsfilterProfileFtgdDns `pulumi:"ftgdDns"`
	// Enable/disable logging of all domains visited (detailed DNS logging). Valid values: `enable`, `disable`.
	LogAllDomain *string `pulumi:"logAllDomain"`
	// External domain block list name.
	Name *string `pulumi:"name"`
	// IP address of the SDNS redirect portal.
	RedirectPortal *string `pulumi:"redirectPortal"`
	// IPv6 address of the SDNS redirect portal.
	RedirectPortal6 *string `pulumi:"redirectPortal6"`
	// Enable/disable Google, Bing, and YouTube safe search. Valid values: `disable`, `enable`.
	SafeSearch *string `pulumi:"safeSearch"`
	// Enable/disable domain filtering and botnet domain logging. Valid values: `enable`, `disable`.
	SdnsDomainLog *string `pulumi:"sdnsDomainLog"`
	// Enable/disable FortiGuard SDNS rating error logging. Valid values: `enable`, `disable`.
	SdnsFtgdErrLog *string `pulumi:"sdnsFtgdErrLog"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Set safe search for YouTube restriction level. Valid values: `strict`, `moderate`.
	YoutubeRestrict *string `pulumi:"youtubeRestrict"`
}

// The set of arguments for constructing a DnsfilterProfile resource.
type DnsfilterProfileArgs struct {
	// Action to take for blocked domains.
	BlockAction pulumi.StringPtrInput
	// Enable/disable blocking botnet C&C DNS lookups. Valid values: `disable`, `enable`.
	BlockBotnet pulumi.StringPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// DNS translation settings. The structure of `dnsTranslation` block is documented below.
	DnsTranslations DnsfilterProfileDnsTranslationArrayInput
	// Domain filter settings. The structure of `domainFilter` block is documented below.
	DomainFilter DnsfilterProfileDomainFilterPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// One or more external IP block lists. The structure of `externalIpBlocklist` block is documented below.
	ExternalIpBlocklists DnsfilterProfileExternalIpBlocklistArrayInput
	// FortiGuard DNS Filter settings. The structure of `ftgdDns` block is documented below.
	FtgdDns DnsfilterProfileFtgdDnsPtrInput
	// Enable/disable logging of all domains visited (detailed DNS logging). Valid values: `enable`, `disable`.
	LogAllDomain pulumi.StringPtrInput
	// External domain block list name.
	Name pulumi.StringPtrInput
	// IP address of the SDNS redirect portal.
	RedirectPortal pulumi.StringPtrInput
	// IPv6 address of the SDNS redirect portal.
	RedirectPortal6 pulumi.StringPtrInput
	// Enable/disable Google, Bing, and YouTube safe search. Valid values: `disable`, `enable`.
	SafeSearch pulumi.StringPtrInput
	// Enable/disable domain filtering and botnet domain logging. Valid values: `enable`, `disable`.
	SdnsDomainLog pulumi.StringPtrInput
	// Enable/disable FortiGuard SDNS rating error logging. Valid values: `enable`, `disable`.
	SdnsFtgdErrLog pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Set safe search for YouTube restriction level. Valid values: `strict`, `moderate`.
	YoutubeRestrict pulumi.StringPtrInput
}

func (DnsfilterProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsfilterProfileArgs)(nil)).Elem()
}

type DnsfilterProfileInput interface {
	pulumi.Input

	ToDnsfilterProfileOutput() DnsfilterProfileOutput
	ToDnsfilterProfileOutputWithContext(ctx context.Context) DnsfilterProfileOutput
}

func (*DnsfilterProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsfilterProfile)(nil)).Elem()
}

func (i *DnsfilterProfile) ToDnsfilterProfileOutput() DnsfilterProfileOutput {
	return i.ToDnsfilterProfileOutputWithContext(context.Background())
}

func (i *DnsfilterProfile) ToDnsfilterProfileOutputWithContext(ctx context.Context) DnsfilterProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsfilterProfileOutput)
}

// DnsfilterProfileArrayInput is an input type that accepts DnsfilterProfileArray and DnsfilterProfileArrayOutput values.
// You can construct a concrete instance of `DnsfilterProfileArrayInput` via:
//
//          DnsfilterProfileArray{ DnsfilterProfileArgs{...} }
type DnsfilterProfileArrayInput interface {
	pulumi.Input

	ToDnsfilterProfileArrayOutput() DnsfilterProfileArrayOutput
	ToDnsfilterProfileArrayOutputWithContext(context.Context) DnsfilterProfileArrayOutput
}

type DnsfilterProfileArray []DnsfilterProfileInput

func (DnsfilterProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsfilterProfile)(nil)).Elem()
}

func (i DnsfilterProfileArray) ToDnsfilterProfileArrayOutput() DnsfilterProfileArrayOutput {
	return i.ToDnsfilterProfileArrayOutputWithContext(context.Background())
}

func (i DnsfilterProfileArray) ToDnsfilterProfileArrayOutputWithContext(ctx context.Context) DnsfilterProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsfilterProfileArrayOutput)
}

// DnsfilterProfileMapInput is an input type that accepts DnsfilterProfileMap and DnsfilterProfileMapOutput values.
// You can construct a concrete instance of `DnsfilterProfileMapInput` via:
//
//          DnsfilterProfileMap{ "key": DnsfilterProfileArgs{...} }
type DnsfilterProfileMapInput interface {
	pulumi.Input

	ToDnsfilterProfileMapOutput() DnsfilterProfileMapOutput
	ToDnsfilterProfileMapOutputWithContext(context.Context) DnsfilterProfileMapOutput
}

type DnsfilterProfileMap map[string]DnsfilterProfileInput

func (DnsfilterProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsfilterProfile)(nil)).Elem()
}

func (i DnsfilterProfileMap) ToDnsfilterProfileMapOutput() DnsfilterProfileMapOutput {
	return i.ToDnsfilterProfileMapOutputWithContext(context.Background())
}

func (i DnsfilterProfileMap) ToDnsfilterProfileMapOutputWithContext(ctx context.Context) DnsfilterProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsfilterProfileMapOutput)
}

type DnsfilterProfileOutput struct{ *pulumi.OutputState }

func (DnsfilterProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsfilterProfile)(nil)).Elem()
}

func (o DnsfilterProfileOutput) ToDnsfilterProfileOutput() DnsfilterProfileOutput {
	return o
}

func (o DnsfilterProfileOutput) ToDnsfilterProfileOutputWithContext(ctx context.Context) DnsfilterProfileOutput {
	return o
}

type DnsfilterProfileArrayOutput struct{ *pulumi.OutputState }

func (DnsfilterProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsfilterProfile)(nil)).Elem()
}

func (o DnsfilterProfileArrayOutput) ToDnsfilterProfileArrayOutput() DnsfilterProfileArrayOutput {
	return o
}

func (o DnsfilterProfileArrayOutput) ToDnsfilterProfileArrayOutputWithContext(ctx context.Context) DnsfilterProfileArrayOutput {
	return o
}

func (o DnsfilterProfileArrayOutput) Index(i pulumi.IntInput) DnsfilterProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsfilterProfile {
		return vs[0].([]*DnsfilterProfile)[vs[1].(int)]
	}).(DnsfilterProfileOutput)
}

type DnsfilterProfileMapOutput struct{ *pulumi.OutputState }

func (DnsfilterProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsfilterProfile)(nil)).Elem()
}

func (o DnsfilterProfileMapOutput) ToDnsfilterProfileMapOutput() DnsfilterProfileMapOutput {
	return o
}

func (o DnsfilterProfileMapOutput) ToDnsfilterProfileMapOutputWithContext(ctx context.Context) DnsfilterProfileMapOutput {
	return o
}

func (o DnsfilterProfileMapOutput) MapIndex(k pulumi.StringInput) DnsfilterProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsfilterProfile {
		return vs[0].(map[string]*DnsfilterProfile)[vs[1].(string)]
	}).(DnsfilterProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsfilterProfileInput)(nil)).Elem(), &DnsfilterProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsfilterProfileArrayInput)(nil)).Elem(), DnsfilterProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsfilterProfileMapInput)(nil)).Elem(), DnsfilterProfileMap{})
	pulumi.RegisterOutputType(DnsfilterProfileOutput{})
	pulumi.RegisterOutputType(DnsfilterProfileArrayOutput{})
	pulumi.RegisterOutputType(DnsfilterProfileMapOutput{})
}
