// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure Web filter profiles.
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
// 		_, err := fortios.NewWebfilterProfile(ctx, "trname", &fortios.WebfilterProfileArgs{
// 			ExtendedLog: pulumi.String("disable"),
// 			FtgdWf: &WebfilterProfileFtgdWfArgs{
// 				ExemptQuota: pulumi.String("17"),
// 				Filters: WebfilterProfileFtgdWfFilterArray{
// 					&WebfilterProfileFtgdWfFilterArgs{
// 						Action:              pulumi.String("warning"),
// 						Category:            pulumi.Int(2),
// 						Id:                  pulumi.Int(1),
// 						Log:                 pulumi.String("enable"),
// 						WarnDuration:        pulumi.String("5m"),
// 						WarningDurationType: pulumi.String("timeout"),
// 						WarningPrompt:       pulumi.String("per-category"),
// 					},
// 					&WebfilterProfileFtgdWfFilterArgs{
// 						Action:              pulumi.String("warning"),
// 						Category:            pulumi.Int(7),
// 						Id:                  pulumi.Int(2),
// 						Log:                 pulumi.String("enable"),
// 						WarnDuration:        pulumi.String("5m"),
// 						WarningDurationType: pulumi.String("timeout"),
// 						WarningPrompt:       pulumi.String("per-category"),
// 					},
// 				},
// 				MaxQuotaTimeout:    pulumi.Int(300),
// 				RateCrlUrls:        pulumi.String("enable"),
// 				RateCssUrls:        pulumi.String("enable"),
// 				RateImageUrls:      pulumi.String("enable"),
// 				RateJavascriptUrls: pulumi.String("enable"),
// 			},
// 			HttpsReplacemsg: pulumi.String("enable"),
// 			InspectionMode:  pulumi.String("flow-based"),
// 			LogAllUrl:       pulumi.String("disable"),
// 			Override: &WebfilterProfileOverrideArgs{
// 				OvrdCookie:       pulumi.String("deny"),
// 				OvrdDur:          pulumi.String("15m"),
// 				OvrdDurMode:      pulumi.String("constant"),
// 				OvrdScope:        pulumi.String("user"),
// 				ProfileAttribute: pulumi.String("Login-LAT-Service"),
// 				ProfileType:      pulumi.String("list"),
// 			},
// 			PostAction: pulumi.String("normal"),
// 			Web: &WebfilterProfileWebArgs{
// 				Blacklist:         pulumi.String("disable"),
// 				BwordTable:        pulumi.Int(0),
// 				BwordThreshold:    pulumi.Int(10),
// 				ContentHeaderList: pulumi.Int(0),
// 				LogSearch:         pulumi.String("disable"),
// 				UrlfilterTable:    pulumi.Int(0),
// 				YoutubeRestrict:   pulumi.String("none"),
// 			},
// 			WebContentLog:             pulumi.String("enable"),
// 			WebExtendedAllActionLog:   pulumi.String("disable"),
// 			WebFilterActivexLog:       pulumi.String("enable"),
// 			WebFilterAppletLog:        pulumi.String("enable"),
// 			WebFilterCommandBlockLog:  pulumi.String("enable"),
// 			WebFilterCookieLog:        pulumi.String("enable"),
// 			WebFilterCookieRemovalLog: pulumi.String("enable"),
// 			WebFilterJsLog:            pulumi.String("enable"),
// 			WebFilterJscriptLog:       pulumi.String("enable"),
// 			WebFilterRefererLog:       pulumi.String("enable"),
// 			WebFilterUnknownLog:       pulumi.String("enable"),
// 			WebFilterVbsLog:           pulumi.String("enable"),
// 			WebFtgdErrLog:             pulumi.String("enable"),
// 			WebFtgdQuotaUsage:         pulumi.String("enable"),
// 			WebInvalidDomainLog:       pulumi.String("enable"),
// 			WebUrlLog:                 pulumi.String("enable"),
// 			Wisp:                      pulumi.String("disable"),
// 			WispAlgorithm:             pulumi.String("auto-learning"),
// 			YoutubeChannelStatus:      pulumi.String("disable"),
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
// Webfilter Profile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/webfilterProfile:WebfilterProfile labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WebfilterProfile struct {
	pulumi.CustomResourceState

	// AntiPhishing profile. The structure of `antiphish` block is documented below.
	Antiphish WebfilterProfileAntiphishPtrOutput `pulumi:"antiphish"`
	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable extended logging for web filtering. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringOutput `pulumi:"extendedLog"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringOutput `pulumi:"featureSet"`
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter WebfilterProfileFileFilterPtrOutput `pulumi:"fileFilter"`
	// FortiGuard Web Filter settings. The structure of `ftgdWf` block is documented below.
	FtgdWf WebfilterProfileFtgdWfPtrOutput `pulumi:"ftgdWf"`
	// Enable replacement messages for HTTPS. Valid values: `enable`, `disable`.
	HttpsReplacemsg pulumi.StringOutput `pulumi:"httpsReplacemsg"`
	// Web filtering inspection mode. Valid values: `proxy`, `flow-based`.
	InspectionMode pulumi.StringOutput `pulumi:"inspectionMode"`
	// Enable/disable logging all URLs visited. Valid values: `enable`, `disable`.
	LogAllUrl pulumi.StringOutput `pulumi:"logAllUrl"`
	// Server name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Options for FortiGuard Web Filter. Valid values: `error-allow`, `rate-server-ip`, `connect-request-bypass`, `ftgd-disable`.
	Options pulumi.StringOutput `pulumi:"options"`
	// Web Filter override settings. The structure of `override` block is documented below.
	Override WebfilterProfileOverridePtrOutput `pulumi:"override"`
	// Permitted override types. Valid values: `bannedword-override`, `urlfilter-override`, `fortiguard-wf-override`, `contenttype-check-override`.
	OvrdPerm pulumi.StringOutput `pulumi:"ovrdPerm"`
	// Action taken for HTTP POST traffic. Valid values: `normal`, `block`.
	PostAction pulumi.StringOutput `pulumi:"postAction"`
	// Replacement message group.
	ReplacemsgGroup pulumi.StringOutput `pulumi:"replacemsgGroup"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Web content filtering settings. The structure of `web` block is documented below.
	Web WebfilterProfileWebPtrOutput `pulumi:"web"`
	// Enable/disable logging of AntiPhishing checks. Valid values: `enable`, `disable`.
	WebAntiphishingLog pulumi.StringOutput `pulumi:"webAntiphishingLog"`
	// Enable/disable logging logging blocked web content. Valid values: `enable`, `disable`.
	WebContentLog pulumi.StringOutput `pulumi:"webContentLog"`
	// Enable/disable extended any filter action logging for web filtering. Valid values: `enable`, `disable`.
	WebExtendedAllActionLog pulumi.StringOutput `pulumi:"webExtendedAllActionLog"`
	// Enable/disable logging ActiveX. Valid values: `enable`, `disable`.
	WebFilterActivexLog pulumi.StringOutput `pulumi:"webFilterActivexLog"`
	// Enable/disable logging Java applets. Valid values: `enable`, `disable`.
	WebFilterAppletLog pulumi.StringOutput `pulumi:"webFilterAppletLog"`
	// Enable/disable logging blocked commands. Valid values: `enable`, `disable`.
	WebFilterCommandBlockLog pulumi.StringOutput `pulumi:"webFilterCommandBlockLog"`
	// Enable/disable logging cookie filtering. Valid values: `enable`, `disable`.
	WebFilterCookieLog pulumi.StringOutput `pulumi:"webFilterCookieLog"`
	// Enable/disable logging blocked cookies. Valid values: `enable`, `disable`.
	WebFilterCookieRemovalLog pulumi.StringOutput `pulumi:"webFilterCookieRemovalLog"`
	// Enable/disable logging Java scripts. Valid values: `enable`, `disable`.
	WebFilterJsLog pulumi.StringOutput `pulumi:"webFilterJsLog"`
	// Enable/disable logging JScripts. Valid values: `enable`, `disable`.
	WebFilterJscriptLog pulumi.StringOutput `pulumi:"webFilterJscriptLog"`
	// Enable/disable logging referrers. Valid values: `enable`, `disable`.
	WebFilterRefererLog pulumi.StringOutput `pulumi:"webFilterRefererLog"`
	// Enable/disable logging unknown scripts. Valid values: `enable`, `disable`.
	WebFilterUnknownLog pulumi.StringOutput `pulumi:"webFilterUnknownLog"`
	// Enable/disable logging VBS scripts. Valid values: `enable`, `disable`.
	WebFilterVbsLog pulumi.StringOutput `pulumi:"webFilterVbsLog"`
	// Enable/disable logging rating errors. Valid values: `enable`, `disable`.
	WebFtgdErrLog pulumi.StringOutput `pulumi:"webFtgdErrLog"`
	// Enable/disable logging daily quota usage. Valid values: `enable`, `disable`.
	WebFtgdQuotaUsage pulumi.StringOutput `pulumi:"webFtgdQuotaUsage"`
	// Enable/disable logging invalid domain names. Valid values: `enable`, `disable`.
	WebInvalidDomainLog pulumi.StringOutput `pulumi:"webInvalidDomainLog"`
	// Enable/disable logging URL filtering. Valid values: `enable`, `disable`.
	WebUrlLog pulumi.StringOutput `pulumi:"webUrlLog"`
	// Enable/disable web proxy WISP. Valid values: `enable`, `disable`.
	Wisp pulumi.StringOutput `pulumi:"wisp"`
	// WISP server selection algorithm. Valid values: `primary-secondary`, `round-robin`, `auto-learning`.
	WispAlgorithm pulumi.StringOutput `pulumi:"wispAlgorithm"`
	// WISP servers. The structure of `wispServers` block is documented below.
	WispServers WebfilterProfileWispServerArrayOutput `pulumi:"wispServers"`
	// YouTube channel filter. The structure of `youtubeChannelFilter` block is documented below.
	YoutubeChannelFilters WebfilterProfileYoutubeChannelFilterArrayOutput `pulumi:"youtubeChannelFilters"`
	// YouTube channel filter status.
	YoutubeChannelStatus pulumi.StringOutput `pulumi:"youtubeChannelStatus"`
}

// NewWebfilterProfile registers a new resource with the given unique name, arguments, and options.
func NewWebfilterProfile(ctx *pulumi.Context,
	name string, args *WebfilterProfileArgs, opts ...pulumi.ResourceOption) (*WebfilterProfile, error) {
	if args == nil {
		args = &WebfilterProfileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WebfilterProfile
	err := ctx.RegisterResource("fortios:index/webfilterProfile:WebfilterProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebfilterProfile gets an existing WebfilterProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebfilterProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebfilterProfileState, opts ...pulumi.ResourceOption) (*WebfilterProfile, error) {
	var resource WebfilterProfile
	err := ctx.ReadResource("fortios:index/webfilterProfile:WebfilterProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebfilterProfile resources.
type webfilterProfileState struct {
	// AntiPhishing profile. The structure of `antiphish` block is documented below.
	Antiphish *WebfilterProfileAntiphish `pulumi:"antiphish"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable extended logging for web filtering. Valid values: `enable`, `disable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet *string `pulumi:"featureSet"`
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter *WebfilterProfileFileFilter `pulumi:"fileFilter"`
	// FortiGuard Web Filter settings. The structure of `ftgdWf` block is documented below.
	FtgdWf *WebfilterProfileFtgdWf `pulumi:"ftgdWf"`
	// Enable replacement messages for HTTPS. Valid values: `enable`, `disable`.
	HttpsReplacemsg *string `pulumi:"httpsReplacemsg"`
	// Web filtering inspection mode. Valid values: `proxy`, `flow-based`.
	InspectionMode *string `pulumi:"inspectionMode"`
	// Enable/disable logging all URLs visited. Valid values: `enable`, `disable`.
	LogAllUrl *string `pulumi:"logAllUrl"`
	// Server name.
	Name *string `pulumi:"name"`
	// Options for FortiGuard Web Filter. Valid values: `error-allow`, `rate-server-ip`, `connect-request-bypass`, `ftgd-disable`.
	Options *string `pulumi:"options"`
	// Web Filter override settings. The structure of `override` block is documented below.
	Override *WebfilterProfileOverride `pulumi:"override"`
	// Permitted override types. Valid values: `bannedword-override`, `urlfilter-override`, `fortiguard-wf-override`, `contenttype-check-override`.
	OvrdPerm *string `pulumi:"ovrdPerm"`
	// Action taken for HTTP POST traffic. Valid values: `normal`, `block`.
	PostAction *string `pulumi:"postAction"`
	// Replacement message group.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Web content filtering settings. The structure of `web` block is documented below.
	Web *WebfilterProfileWeb `pulumi:"web"`
	// Enable/disable logging of AntiPhishing checks. Valid values: `enable`, `disable`.
	WebAntiphishingLog *string `pulumi:"webAntiphishingLog"`
	// Enable/disable logging logging blocked web content. Valid values: `enable`, `disable`.
	WebContentLog *string `pulumi:"webContentLog"`
	// Enable/disable extended any filter action logging for web filtering. Valid values: `enable`, `disable`.
	WebExtendedAllActionLog *string `pulumi:"webExtendedAllActionLog"`
	// Enable/disable logging ActiveX. Valid values: `enable`, `disable`.
	WebFilterActivexLog *string `pulumi:"webFilterActivexLog"`
	// Enable/disable logging Java applets. Valid values: `enable`, `disable`.
	WebFilterAppletLog *string `pulumi:"webFilterAppletLog"`
	// Enable/disable logging blocked commands. Valid values: `enable`, `disable`.
	WebFilterCommandBlockLog *string `pulumi:"webFilterCommandBlockLog"`
	// Enable/disable logging cookie filtering. Valid values: `enable`, `disable`.
	WebFilterCookieLog *string `pulumi:"webFilterCookieLog"`
	// Enable/disable logging blocked cookies. Valid values: `enable`, `disable`.
	WebFilterCookieRemovalLog *string `pulumi:"webFilterCookieRemovalLog"`
	// Enable/disable logging Java scripts. Valid values: `enable`, `disable`.
	WebFilterJsLog *string `pulumi:"webFilterJsLog"`
	// Enable/disable logging JScripts. Valid values: `enable`, `disable`.
	WebFilterJscriptLog *string `pulumi:"webFilterJscriptLog"`
	// Enable/disable logging referrers. Valid values: `enable`, `disable`.
	WebFilterRefererLog *string `pulumi:"webFilterRefererLog"`
	// Enable/disable logging unknown scripts. Valid values: `enable`, `disable`.
	WebFilterUnknownLog *string `pulumi:"webFilterUnknownLog"`
	// Enable/disable logging VBS scripts. Valid values: `enable`, `disable`.
	WebFilterVbsLog *string `pulumi:"webFilterVbsLog"`
	// Enable/disable logging rating errors. Valid values: `enable`, `disable`.
	WebFtgdErrLog *string `pulumi:"webFtgdErrLog"`
	// Enable/disable logging daily quota usage. Valid values: `enable`, `disable`.
	WebFtgdQuotaUsage *string `pulumi:"webFtgdQuotaUsage"`
	// Enable/disable logging invalid domain names. Valid values: `enable`, `disable`.
	WebInvalidDomainLog *string `pulumi:"webInvalidDomainLog"`
	// Enable/disable logging URL filtering. Valid values: `enable`, `disable`.
	WebUrlLog *string `pulumi:"webUrlLog"`
	// Enable/disable web proxy WISP. Valid values: `enable`, `disable`.
	Wisp *string `pulumi:"wisp"`
	// WISP server selection algorithm. Valid values: `primary-secondary`, `round-robin`, `auto-learning`.
	WispAlgorithm *string `pulumi:"wispAlgorithm"`
	// WISP servers. The structure of `wispServers` block is documented below.
	WispServers []WebfilterProfileWispServer `pulumi:"wispServers"`
	// YouTube channel filter. The structure of `youtubeChannelFilter` block is documented below.
	YoutubeChannelFilters []WebfilterProfileYoutubeChannelFilter `pulumi:"youtubeChannelFilters"`
	// YouTube channel filter status.
	YoutubeChannelStatus *string `pulumi:"youtubeChannelStatus"`
}

type WebfilterProfileState struct {
	// AntiPhishing profile. The structure of `antiphish` block is documented below.
	Antiphish WebfilterProfileAntiphishPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable extended logging for web filtering. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringPtrInput
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringPtrInput
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter WebfilterProfileFileFilterPtrInput
	// FortiGuard Web Filter settings. The structure of `ftgdWf` block is documented below.
	FtgdWf WebfilterProfileFtgdWfPtrInput
	// Enable replacement messages for HTTPS. Valid values: `enable`, `disable`.
	HttpsReplacemsg pulumi.StringPtrInput
	// Web filtering inspection mode. Valid values: `proxy`, `flow-based`.
	InspectionMode pulumi.StringPtrInput
	// Enable/disable logging all URLs visited. Valid values: `enable`, `disable`.
	LogAllUrl pulumi.StringPtrInput
	// Server name.
	Name pulumi.StringPtrInput
	// Options for FortiGuard Web Filter. Valid values: `error-allow`, `rate-server-ip`, `connect-request-bypass`, `ftgd-disable`.
	Options pulumi.StringPtrInput
	// Web Filter override settings. The structure of `override` block is documented below.
	Override WebfilterProfileOverridePtrInput
	// Permitted override types. Valid values: `bannedword-override`, `urlfilter-override`, `fortiguard-wf-override`, `contenttype-check-override`.
	OvrdPerm pulumi.StringPtrInput
	// Action taken for HTTP POST traffic. Valid values: `normal`, `block`.
	PostAction pulumi.StringPtrInput
	// Replacement message group.
	ReplacemsgGroup pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Web content filtering settings. The structure of `web` block is documented below.
	Web WebfilterProfileWebPtrInput
	// Enable/disable logging of AntiPhishing checks. Valid values: `enable`, `disable`.
	WebAntiphishingLog pulumi.StringPtrInput
	// Enable/disable logging logging blocked web content. Valid values: `enable`, `disable`.
	WebContentLog pulumi.StringPtrInput
	// Enable/disable extended any filter action logging for web filtering. Valid values: `enable`, `disable`.
	WebExtendedAllActionLog pulumi.StringPtrInput
	// Enable/disable logging ActiveX. Valid values: `enable`, `disable`.
	WebFilterActivexLog pulumi.StringPtrInput
	// Enable/disable logging Java applets. Valid values: `enable`, `disable`.
	WebFilterAppletLog pulumi.StringPtrInput
	// Enable/disable logging blocked commands. Valid values: `enable`, `disable`.
	WebFilterCommandBlockLog pulumi.StringPtrInput
	// Enable/disable logging cookie filtering. Valid values: `enable`, `disable`.
	WebFilterCookieLog pulumi.StringPtrInput
	// Enable/disable logging blocked cookies. Valid values: `enable`, `disable`.
	WebFilterCookieRemovalLog pulumi.StringPtrInput
	// Enable/disable logging Java scripts. Valid values: `enable`, `disable`.
	WebFilterJsLog pulumi.StringPtrInput
	// Enable/disable logging JScripts. Valid values: `enable`, `disable`.
	WebFilterJscriptLog pulumi.StringPtrInput
	// Enable/disable logging referrers. Valid values: `enable`, `disable`.
	WebFilterRefererLog pulumi.StringPtrInput
	// Enable/disable logging unknown scripts. Valid values: `enable`, `disable`.
	WebFilterUnknownLog pulumi.StringPtrInput
	// Enable/disable logging VBS scripts. Valid values: `enable`, `disable`.
	WebFilterVbsLog pulumi.StringPtrInput
	// Enable/disable logging rating errors. Valid values: `enable`, `disable`.
	WebFtgdErrLog pulumi.StringPtrInput
	// Enable/disable logging daily quota usage. Valid values: `enable`, `disable`.
	WebFtgdQuotaUsage pulumi.StringPtrInput
	// Enable/disable logging invalid domain names. Valid values: `enable`, `disable`.
	WebInvalidDomainLog pulumi.StringPtrInput
	// Enable/disable logging URL filtering. Valid values: `enable`, `disable`.
	WebUrlLog pulumi.StringPtrInput
	// Enable/disable web proxy WISP. Valid values: `enable`, `disable`.
	Wisp pulumi.StringPtrInput
	// WISP server selection algorithm. Valid values: `primary-secondary`, `round-robin`, `auto-learning`.
	WispAlgorithm pulumi.StringPtrInput
	// WISP servers. The structure of `wispServers` block is documented below.
	WispServers WebfilterProfileWispServerArrayInput
	// YouTube channel filter. The structure of `youtubeChannelFilter` block is documented below.
	YoutubeChannelFilters WebfilterProfileYoutubeChannelFilterArrayInput
	// YouTube channel filter status.
	YoutubeChannelStatus pulumi.StringPtrInput
}

func (WebfilterProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterProfileState)(nil)).Elem()
}

type webfilterProfileArgs struct {
	// AntiPhishing profile. The structure of `antiphish` block is documented below.
	Antiphish *WebfilterProfileAntiphish `pulumi:"antiphish"`
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable extended logging for web filtering. Valid values: `enable`, `disable`.
	ExtendedLog *string `pulumi:"extendedLog"`
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet *string `pulumi:"featureSet"`
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter *WebfilterProfileFileFilter `pulumi:"fileFilter"`
	// FortiGuard Web Filter settings. The structure of `ftgdWf` block is documented below.
	FtgdWf *WebfilterProfileFtgdWf `pulumi:"ftgdWf"`
	// Enable replacement messages for HTTPS. Valid values: `enable`, `disable`.
	HttpsReplacemsg *string `pulumi:"httpsReplacemsg"`
	// Web filtering inspection mode. Valid values: `proxy`, `flow-based`.
	InspectionMode *string `pulumi:"inspectionMode"`
	// Enable/disable logging all URLs visited. Valid values: `enable`, `disable`.
	LogAllUrl *string `pulumi:"logAllUrl"`
	// Server name.
	Name *string `pulumi:"name"`
	// Options for FortiGuard Web Filter. Valid values: `error-allow`, `rate-server-ip`, `connect-request-bypass`, `ftgd-disable`.
	Options *string `pulumi:"options"`
	// Web Filter override settings. The structure of `override` block is documented below.
	Override *WebfilterProfileOverride `pulumi:"override"`
	// Permitted override types. Valid values: `bannedword-override`, `urlfilter-override`, `fortiguard-wf-override`, `contenttype-check-override`.
	OvrdPerm *string `pulumi:"ovrdPerm"`
	// Action taken for HTTP POST traffic. Valid values: `normal`, `block`.
	PostAction *string `pulumi:"postAction"`
	// Replacement message group.
	ReplacemsgGroup *string `pulumi:"replacemsgGroup"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Web content filtering settings. The structure of `web` block is documented below.
	Web *WebfilterProfileWeb `pulumi:"web"`
	// Enable/disable logging of AntiPhishing checks. Valid values: `enable`, `disable`.
	WebAntiphishingLog *string `pulumi:"webAntiphishingLog"`
	// Enable/disable logging logging blocked web content. Valid values: `enable`, `disable`.
	WebContentLog *string `pulumi:"webContentLog"`
	// Enable/disable extended any filter action logging for web filtering. Valid values: `enable`, `disable`.
	WebExtendedAllActionLog *string `pulumi:"webExtendedAllActionLog"`
	// Enable/disable logging ActiveX. Valid values: `enable`, `disable`.
	WebFilterActivexLog *string `pulumi:"webFilterActivexLog"`
	// Enable/disable logging Java applets. Valid values: `enable`, `disable`.
	WebFilterAppletLog *string `pulumi:"webFilterAppletLog"`
	// Enable/disable logging blocked commands. Valid values: `enable`, `disable`.
	WebFilterCommandBlockLog *string `pulumi:"webFilterCommandBlockLog"`
	// Enable/disable logging cookie filtering. Valid values: `enable`, `disable`.
	WebFilterCookieLog *string `pulumi:"webFilterCookieLog"`
	// Enable/disable logging blocked cookies. Valid values: `enable`, `disable`.
	WebFilterCookieRemovalLog *string `pulumi:"webFilterCookieRemovalLog"`
	// Enable/disable logging Java scripts. Valid values: `enable`, `disable`.
	WebFilterJsLog *string `pulumi:"webFilterJsLog"`
	// Enable/disable logging JScripts. Valid values: `enable`, `disable`.
	WebFilterJscriptLog *string `pulumi:"webFilterJscriptLog"`
	// Enable/disable logging referrers. Valid values: `enable`, `disable`.
	WebFilterRefererLog *string `pulumi:"webFilterRefererLog"`
	// Enable/disable logging unknown scripts. Valid values: `enable`, `disable`.
	WebFilterUnknownLog *string `pulumi:"webFilterUnknownLog"`
	// Enable/disable logging VBS scripts. Valid values: `enable`, `disable`.
	WebFilterVbsLog *string `pulumi:"webFilterVbsLog"`
	// Enable/disable logging rating errors. Valid values: `enable`, `disable`.
	WebFtgdErrLog *string `pulumi:"webFtgdErrLog"`
	// Enable/disable logging daily quota usage. Valid values: `enable`, `disable`.
	WebFtgdQuotaUsage *string `pulumi:"webFtgdQuotaUsage"`
	// Enable/disable logging invalid domain names. Valid values: `enable`, `disable`.
	WebInvalidDomainLog *string `pulumi:"webInvalidDomainLog"`
	// Enable/disable logging URL filtering. Valid values: `enable`, `disable`.
	WebUrlLog *string `pulumi:"webUrlLog"`
	// Enable/disable web proxy WISP. Valid values: `enable`, `disable`.
	Wisp *string `pulumi:"wisp"`
	// WISP server selection algorithm. Valid values: `primary-secondary`, `round-robin`, `auto-learning`.
	WispAlgorithm *string `pulumi:"wispAlgorithm"`
	// WISP servers. The structure of `wispServers` block is documented below.
	WispServers []WebfilterProfileWispServer `pulumi:"wispServers"`
	// YouTube channel filter. The structure of `youtubeChannelFilter` block is documented below.
	YoutubeChannelFilters []WebfilterProfileYoutubeChannelFilter `pulumi:"youtubeChannelFilters"`
	// YouTube channel filter status.
	YoutubeChannelStatus *string `pulumi:"youtubeChannelStatus"`
}

// The set of arguments for constructing a WebfilterProfile resource.
type WebfilterProfileArgs struct {
	// AntiPhishing profile. The structure of `antiphish` block is documented below.
	Antiphish WebfilterProfileAntiphishPtrInput
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable extended logging for web filtering. Valid values: `enable`, `disable`.
	ExtendedLog pulumi.StringPtrInput
	// Flow/proxy feature set. Valid values: `flow`, `proxy`.
	FeatureSet pulumi.StringPtrInput
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter WebfilterProfileFileFilterPtrInput
	// FortiGuard Web Filter settings. The structure of `ftgdWf` block is documented below.
	FtgdWf WebfilterProfileFtgdWfPtrInput
	// Enable replacement messages for HTTPS. Valid values: `enable`, `disable`.
	HttpsReplacemsg pulumi.StringPtrInput
	// Web filtering inspection mode. Valid values: `proxy`, `flow-based`.
	InspectionMode pulumi.StringPtrInput
	// Enable/disable logging all URLs visited. Valid values: `enable`, `disable`.
	LogAllUrl pulumi.StringPtrInput
	// Server name.
	Name pulumi.StringPtrInput
	// Options for FortiGuard Web Filter. Valid values: `error-allow`, `rate-server-ip`, `connect-request-bypass`, `ftgd-disable`.
	Options pulumi.StringPtrInput
	// Web Filter override settings. The structure of `override` block is documented below.
	Override WebfilterProfileOverridePtrInput
	// Permitted override types. Valid values: `bannedword-override`, `urlfilter-override`, `fortiguard-wf-override`, `contenttype-check-override`.
	OvrdPerm pulumi.StringPtrInput
	// Action taken for HTTP POST traffic. Valid values: `normal`, `block`.
	PostAction pulumi.StringPtrInput
	// Replacement message group.
	ReplacemsgGroup pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Web content filtering settings. The structure of `web` block is documented below.
	Web WebfilterProfileWebPtrInput
	// Enable/disable logging of AntiPhishing checks. Valid values: `enable`, `disable`.
	WebAntiphishingLog pulumi.StringPtrInput
	// Enable/disable logging logging blocked web content. Valid values: `enable`, `disable`.
	WebContentLog pulumi.StringPtrInput
	// Enable/disable extended any filter action logging for web filtering. Valid values: `enable`, `disable`.
	WebExtendedAllActionLog pulumi.StringPtrInput
	// Enable/disable logging ActiveX. Valid values: `enable`, `disable`.
	WebFilterActivexLog pulumi.StringPtrInput
	// Enable/disable logging Java applets. Valid values: `enable`, `disable`.
	WebFilterAppletLog pulumi.StringPtrInput
	// Enable/disable logging blocked commands. Valid values: `enable`, `disable`.
	WebFilterCommandBlockLog pulumi.StringPtrInput
	// Enable/disable logging cookie filtering. Valid values: `enable`, `disable`.
	WebFilterCookieLog pulumi.StringPtrInput
	// Enable/disable logging blocked cookies. Valid values: `enable`, `disable`.
	WebFilterCookieRemovalLog pulumi.StringPtrInput
	// Enable/disable logging Java scripts. Valid values: `enable`, `disable`.
	WebFilterJsLog pulumi.StringPtrInput
	// Enable/disable logging JScripts. Valid values: `enable`, `disable`.
	WebFilterJscriptLog pulumi.StringPtrInput
	// Enable/disable logging referrers. Valid values: `enable`, `disable`.
	WebFilterRefererLog pulumi.StringPtrInput
	// Enable/disable logging unknown scripts. Valid values: `enable`, `disable`.
	WebFilterUnknownLog pulumi.StringPtrInput
	// Enable/disable logging VBS scripts. Valid values: `enable`, `disable`.
	WebFilterVbsLog pulumi.StringPtrInput
	// Enable/disable logging rating errors. Valid values: `enable`, `disable`.
	WebFtgdErrLog pulumi.StringPtrInput
	// Enable/disable logging daily quota usage. Valid values: `enable`, `disable`.
	WebFtgdQuotaUsage pulumi.StringPtrInput
	// Enable/disable logging invalid domain names. Valid values: `enable`, `disable`.
	WebInvalidDomainLog pulumi.StringPtrInput
	// Enable/disable logging URL filtering. Valid values: `enable`, `disable`.
	WebUrlLog pulumi.StringPtrInput
	// Enable/disable web proxy WISP. Valid values: `enable`, `disable`.
	Wisp pulumi.StringPtrInput
	// WISP server selection algorithm. Valid values: `primary-secondary`, `round-robin`, `auto-learning`.
	WispAlgorithm pulumi.StringPtrInput
	// WISP servers. The structure of `wispServers` block is documented below.
	WispServers WebfilterProfileWispServerArrayInput
	// YouTube channel filter. The structure of `youtubeChannelFilter` block is documented below.
	YoutubeChannelFilters WebfilterProfileYoutubeChannelFilterArrayInput
	// YouTube channel filter status.
	YoutubeChannelStatus pulumi.StringPtrInput
}

func (WebfilterProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webfilterProfileArgs)(nil)).Elem()
}

type WebfilterProfileInput interface {
	pulumi.Input

	ToWebfilterProfileOutput() WebfilterProfileOutput
	ToWebfilterProfileOutputWithContext(ctx context.Context) WebfilterProfileOutput
}

func (*WebfilterProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterProfile)(nil)).Elem()
}

func (i *WebfilterProfile) ToWebfilterProfileOutput() WebfilterProfileOutput {
	return i.ToWebfilterProfileOutputWithContext(context.Background())
}

func (i *WebfilterProfile) ToWebfilterProfileOutputWithContext(ctx context.Context) WebfilterProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterProfileOutput)
}

// WebfilterProfileArrayInput is an input type that accepts WebfilterProfileArray and WebfilterProfileArrayOutput values.
// You can construct a concrete instance of `WebfilterProfileArrayInput` via:
//
//          WebfilterProfileArray{ WebfilterProfileArgs{...} }
type WebfilterProfileArrayInput interface {
	pulumi.Input

	ToWebfilterProfileArrayOutput() WebfilterProfileArrayOutput
	ToWebfilterProfileArrayOutputWithContext(context.Context) WebfilterProfileArrayOutput
}

type WebfilterProfileArray []WebfilterProfileInput

func (WebfilterProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebfilterProfile)(nil)).Elem()
}

func (i WebfilterProfileArray) ToWebfilterProfileArrayOutput() WebfilterProfileArrayOutput {
	return i.ToWebfilterProfileArrayOutputWithContext(context.Background())
}

func (i WebfilterProfileArray) ToWebfilterProfileArrayOutputWithContext(ctx context.Context) WebfilterProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterProfileArrayOutput)
}

// WebfilterProfileMapInput is an input type that accepts WebfilterProfileMap and WebfilterProfileMapOutput values.
// You can construct a concrete instance of `WebfilterProfileMapInput` via:
//
//          WebfilterProfileMap{ "key": WebfilterProfileArgs{...} }
type WebfilterProfileMapInput interface {
	pulumi.Input

	ToWebfilterProfileMapOutput() WebfilterProfileMapOutput
	ToWebfilterProfileMapOutputWithContext(context.Context) WebfilterProfileMapOutput
}

type WebfilterProfileMap map[string]WebfilterProfileInput

func (WebfilterProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebfilterProfile)(nil)).Elem()
}

func (i WebfilterProfileMap) ToWebfilterProfileMapOutput() WebfilterProfileMapOutput {
	return i.ToWebfilterProfileMapOutputWithContext(context.Background())
}

func (i WebfilterProfileMap) ToWebfilterProfileMapOutputWithContext(ctx context.Context) WebfilterProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebfilterProfileMapOutput)
}

type WebfilterProfileOutput struct{ *pulumi.OutputState }

func (WebfilterProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebfilterProfile)(nil)).Elem()
}

func (o WebfilterProfileOutput) ToWebfilterProfileOutput() WebfilterProfileOutput {
	return o
}

func (o WebfilterProfileOutput) ToWebfilterProfileOutputWithContext(ctx context.Context) WebfilterProfileOutput {
	return o
}

type WebfilterProfileArrayOutput struct{ *pulumi.OutputState }

func (WebfilterProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebfilterProfile)(nil)).Elem()
}

func (o WebfilterProfileArrayOutput) ToWebfilterProfileArrayOutput() WebfilterProfileArrayOutput {
	return o
}

func (o WebfilterProfileArrayOutput) ToWebfilterProfileArrayOutputWithContext(ctx context.Context) WebfilterProfileArrayOutput {
	return o
}

func (o WebfilterProfileArrayOutput) Index(i pulumi.IntInput) WebfilterProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebfilterProfile {
		return vs[0].([]*WebfilterProfile)[vs[1].(int)]
	}).(WebfilterProfileOutput)
}

type WebfilterProfileMapOutput struct{ *pulumi.OutputState }

func (WebfilterProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebfilterProfile)(nil)).Elem()
}

func (o WebfilterProfileMapOutput) ToWebfilterProfileMapOutput() WebfilterProfileMapOutput {
	return o
}

func (o WebfilterProfileMapOutput) ToWebfilterProfileMapOutputWithContext(ctx context.Context) WebfilterProfileMapOutput {
	return o
}

func (o WebfilterProfileMapOutput) MapIndex(k pulumi.StringInput) WebfilterProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebfilterProfile {
		return vs[0].(map[string]*WebfilterProfile)[vs[1].(string)]
	}).(WebfilterProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterProfileInput)(nil)).Elem(), &WebfilterProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterProfileArrayInput)(nil)).Elem(), WebfilterProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebfilterProfileMapInput)(nil)).Elem(), WebfilterProfileMap{})
	pulumi.RegisterOutputType(WebfilterProfileOutput{})
	pulumi.RegisterOutputType(WebfilterProfileArrayOutput{})
	pulumi.RegisterOutputType(WebfilterProfileMapOutput{})
}
