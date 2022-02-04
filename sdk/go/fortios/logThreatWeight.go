// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure threat weight settings.
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
// 		_, err := fortios.NewLogThreatWeight(ctx, "trname", &fortios.LogThreatWeightArgs{
// 			Applications: LogThreatWeightApplicationArray{
// 				&LogThreatWeightApplicationArgs{
// 					Category: pulumi.Int(2),
// 					Id:       pulumi.Int(1),
// 					Level:    pulumi.String("low"),
// 				},
// 				&LogThreatWeightApplicationArgs{
// 					Category: pulumi.Int(6),
// 					Id:       pulumi.Int(2),
// 					Level:    pulumi.String("medium"),
// 				},
// 			},
// 			BlockedConnection: pulumi.String("high"),
// 			FailedConnection:  pulumi.String("low"),
// 			Ips: &LogThreatWeightIpsArgs{
// 				CriticalSeverity: pulumi.String("critical"),
// 				HighSeverity:     pulumi.String("high"),
// 				InfoSeverity:     pulumi.String("disable"),
// 				LowSeverity:      pulumi.String("low"),
// 				MediumSeverity:   pulumi.String("medium"),
// 			},
// 			Level: &LogThreatWeightLevelArgs{
// 				Critical: pulumi.Int(50),
// 				High:     pulumi.Int(30),
// 				Low:      pulumi.Int(5),
// 				Medium:   pulumi.Int(10),
// 			},
// 			Malware: &LogThreatWeightMalwareArgs{
// 				BotnetConnection:        pulumi.String("critical"),
// 				CommandBlocked:          pulumi.String("disable"),
// 				ContentDisarm:           pulumi.String("medium"),
// 				FileBlocked:             pulumi.String("low"),
// 				Mimefragmented:          pulumi.String("disable"),
// 				Oversized:               pulumi.String("disable"),
// 				SwitchProto:             pulumi.String("disable"),
// 				VirusFileTypeExecutable: pulumi.String("medium"),
// 				VirusInfected:           pulumi.String("critical"),
// 				VirusOutbreakPrevention: pulumi.String("critical"),
// 				VirusScanError:          pulumi.String("high"),
// 			},
// 			Status:           pulumi.String("enable"),
// 			UrlBlockDetected: pulumi.String("high"),
// 			Webs: LogThreatWeightWebArray{
// 				&LogThreatWeightWebArgs{
// 					Category: pulumi.Int(26),
// 					Id:       pulumi.Int(1),
// 					Level:    pulumi.String("high"),
// 				},
// 				&LogThreatWeightWebArgs{
// 					Category: pulumi.Int(61),
// 					Id:       pulumi.Int(2),
// 					Level:    pulumi.String("high"),
// 				},
// 				&LogThreatWeightWebArgs{
// 					Category: pulumi.Int(86),
// 					Id:       pulumi.Int(3),
// 					Level:    pulumi.String("high"),
// 				},
// 				&LogThreatWeightWebArgs{
// 					Category: pulumi.Int(1),
// 					Id:       pulumi.Int(4),
// 					Level:    pulumi.String("medium"),
// 				},
// 				&LogThreatWeightWebArgs{
// 					Category: pulumi.Int(3),
// 					Id:       pulumi.Int(5),
// 					Level:    pulumi.String("medium"),
// 				},
// 				&LogThreatWeightWebArgs{
// 					Category: pulumi.Int(4),
// 					Id:       pulumi.Int(6),
// 					Level:    pulumi.String("medium"),
// 				},
// 				&LogThreatWeightWebArgs{
// 					Category: pulumi.Int(5),
// 					Id:       pulumi.Int(7),
// 					Level:    pulumi.String("medium"),
// 				},
// 				&LogThreatWeightWebArgs{
// 					Category: pulumi.Int(6),
// 					Id:       pulumi.Int(8),
// 					Level:    pulumi.String("medium"),
// 				},
// 				&LogThreatWeightWebArgs{
// 					Category: pulumi.Int(12),
// 					Id:       pulumi.Int(9),
// 					Level:    pulumi.String("medium"),
// 				},
// 				&LogThreatWeightWebArgs{
// 					Category: pulumi.Int(59),
// 					Id:       pulumi.Int(10),
// 					Level:    pulumi.String("medium"),
// 				},
// 				&LogThreatWeightWebArgs{
// 					Category: pulumi.Int(62),
// 					Id:       pulumi.Int(11),
// 					Level:    pulumi.String("medium"),
// 				},
// 				&LogThreatWeightWebArgs{
// 					Category: pulumi.Int(83),
// 					Id:       pulumi.Int(12),
// 					Level:    pulumi.String("medium"),
// 				},
// 				&LogThreatWeightWebArgs{
// 					Category: pulumi.Int(72),
// 					Id:       pulumi.Int(13),
// 					Level:    pulumi.String("low"),
// 				},
// 				&LogThreatWeightWebArgs{
// 					Category: pulumi.Int(14),
// 					Id:       pulumi.Int(14),
// 					Level:    pulumi.String("low"),
// 				},
// 			},
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
// Log ThreatWeight can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/logThreatWeight:LogThreatWeight labelname LogThreatWeight
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type LogThreatWeight struct {
	pulumi.CustomResourceState

	// Application-control threat weight settings. The structure of `application` block is documented below.
	Applications LogThreatWeightApplicationArrayOutput `pulumi:"applications"`
	// Threat weight score for blocked connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	BlockedConnection pulumi.StringOutput `pulumi:"blockedConnection"`
	// Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	BotnetConnectionDetected pulumi.StringOutput `pulumi:"botnetConnectionDetected"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Threat weight score for failed connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	FailedConnection pulumi.StringOutput `pulumi:"failedConnection"`
	// Geolocation-based threat weight settings. The structure of `geolocation` block is documented below.
	Geolocations LogThreatWeightGeolocationArrayOutput `pulumi:"geolocations"`
	// IPS threat weight settings. The structure of `ips` block is documented below.
	Ips LogThreatWeightIpsPtrOutput `pulumi:"ips"`
	// Threat weight score for Application events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	Level LogThreatWeightLevelPtrOutput `pulumi:"level"`
	// Anti-virus malware threat weight settings. The structure of `malware` block is documented below.
	Malware LogThreatWeightMalwarePtrOutput `pulumi:"malware"`
	// Enable/disable the threat weight feature. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Threat weight score for URL blocking. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	UrlBlockDetected pulumi.StringOutput `pulumi:"urlBlockDetected"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Web filtering threat weight settings. The structure of `web` block is documented below.
	Webs LogThreatWeightWebArrayOutput `pulumi:"webs"`
}

// NewLogThreatWeight registers a new resource with the given unique name, arguments, and options.
func NewLogThreatWeight(ctx *pulumi.Context,
	name string, args *LogThreatWeightArgs, opts ...pulumi.ResourceOption) (*LogThreatWeight, error) {
	if args == nil {
		args = &LogThreatWeightArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogThreatWeight
	err := ctx.RegisterResource("fortios:index/logThreatWeight:LogThreatWeight", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogThreatWeight gets an existing LogThreatWeight resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogThreatWeight(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogThreatWeightState, opts ...pulumi.ResourceOption) (*LogThreatWeight, error) {
	var resource LogThreatWeight
	err := ctx.ReadResource("fortios:index/logThreatWeight:LogThreatWeight", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogThreatWeight resources.
type logThreatWeightState struct {
	// Application-control threat weight settings. The structure of `application` block is documented below.
	Applications []LogThreatWeightApplication `pulumi:"applications"`
	// Threat weight score for blocked connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	BlockedConnection *string `pulumi:"blockedConnection"`
	// Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	BotnetConnectionDetected *string `pulumi:"botnetConnectionDetected"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Threat weight score for failed connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	FailedConnection *string `pulumi:"failedConnection"`
	// Geolocation-based threat weight settings. The structure of `geolocation` block is documented below.
	Geolocations []LogThreatWeightGeolocation `pulumi:"geolocations"`
	// IPS threat weight settings. The structure of `ips` block is documented below.
	Ips *LogThreatWeightIps `pulumi:"ips"`
	// Threat weight score for Application events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	Level *LogThreatWeightLevel `pulumi:"level"`
	// Anti-virus malware threat weight settings. The structure of `malware` block is documented below.
	Malware *LogThreatWeightMalware `pulumi:"malware"`
	// Enable/disable the threat weight feature. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Threat weight score for URL blocking. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	UrlBlockDetected *string `pulumi:"urlBlockDetected"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Web filtering threat weight settings. The structure of `web` block is documented below.
	Webs []LogThreatWeightWeb `pulumi:"webs"`
}

type LogThreatWeightState struct {
	// Application-control threat weight settings. The structure of `application` block is documented below.
	Applications LogThreatWeightApplicationArrayInput
	// Threat weight score for blocked connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	BlockedConnection pulumi.StringPtrInput
	// Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	BotnetConnectionDetected pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Threat weight score for failed connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	FailedConnection pulumi.StringPtrInput
	// Geolocation-based threat weight settings. The structure of `geolocation` block is documented below.
	Geolocations LogThreatWeightGeolocationArrayInput
	// IPS threat weight settings. The structure of `ips` block is documented below.
	Ips LogThreatWeightIpsPtrInput
	// Threat weight score for Application events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	Level LogThreatWeightLevelPtrInput
	// Anti-virus malware threat weight settings. The structure of `malware` block is documented below.
	Malware LogThreatWeightMalwarePtrInput
	// Enable/disable the threat weight feature. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Threat weight score for URL blocking. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	UrlBlockDetected pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Web filtering threat weight settings. The structure of `web` block is documented below.
	Webs LogThreatWeightWebArrayInput
}

func (LogThreatWeightState) ElementType() reflect.Type {
	return reflect.TypeOf((*logThreatWeightState)(nil)).Elem()
}

type logThreatWeightArgs struct {
	// Application-control threat weight settings. The structure of `application` block is documented below.
	Applications []LogThreatWeightApplication `pulumi:"applications"`
	// Threat weight score for blocked connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	BlockedConnection *string `pulumi:"blockedConnection"`
	// Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	BotnetConnectionDetected *string `pulumi:"botnetConnectionDetected"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Threat weight score for failed connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	FailedConnection *string `pulumi:"failedConnection"`
	// Geolocation-based threat weight settings. The structure of `geolocation` block is documented below.
	Geolocations []LogThreatWeightGeolocation `pulumi:"geolocations"`
	// IPS threat weight settings. The structure of `ips` block is documented below.
	Ips *LogThreatWeightIps `pulumi:"ips"`
	// Threat weight score for Application events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	Level *LogThreatWeightLevel `pulumi:"level"`
	// Anti-virus malware threat weight settings. The structure of `malware` block is documented below.
	Malware *LogThreatWeightMalware `pulumi:"malware"`
	// Enable/disable the threat weight feature. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Threat weight score for URL blocking. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	UrlBlockDetected *string `pulumi:"urlBlockDetected"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Web filtering threat weight settings. The structure of `web` block is documented below.
	Webs []LogThreatWeightWeb `pulumi:"webs"`
}

// The set of arguments for constructing a LogThreatWeight resource.
type LogThreatWeightArgs struct {
	// Application-control threat weight settings. The structure of `application` block is documented below.
	Applications LogThreatWeightApplicationArrayInput
	// Threat weight score for blocked connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	BlockedConnection pulumi.StringPtrInput
	// Threat weight score for detected botnet connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	BotnetConnectionDetected pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Threat weight score for failed connections. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	FailedConnection pulumi.StringPtrInput
	// Geolocation-based threat weight settings. The structure of `geolocation` block is documented below.
	Geolocations LogThreatWeightGeolocationArrayInput
	// IPS threat weight settings. The structure of `ips` block is documented below.
	Ips LogThreatWeightIpsPtrInput
	// Threat weight score for Application events. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	Level LogThreatWeightLevelPtrInput
	// Anti-virus malware threat weight settings. The structure of `malware` block is documented below.
	Malware LogThreatWeightMalwarePtrInput
	// Enable/disable the threat weight feature. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Threat weight score for URL blocking. Valid values: `disable`, `low`, `medium`, `high`, `critical`.
	UrlBlockDetected pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Web filtering threat weight settings. The structure of `web` block is documented below.
	Webs LogThreatWeightWebArrayInput
}

func (LogThreatWeightArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logThreatWeightArgs)(nil)).Elem()
}

type LogThreatWeightInput interface {
	pulumi.Input

	ToLogThreatWeightOutput() LogThreatWeightOutput
	ToLogThreatWeightOutputWithContext(ctx context.Context) LogThreatWeightOutput
}

func (*LogThreatWeight) ElementType() reflect.Type {
	return reflect.TypeOf((**LogThreatWeight)(nil)).Elem()
}

func (i *LogThreatWeight) ToLogThreatWeightOutput() LogThreatWeightOutput {
	return i.ToLogThreatWeightOutputWithContext(context.Background())
}

func (i *LogThreatWeight) ToLogThreatWeightOutputWithContext(ctx context.Context) LogThreatWeightOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogThreatWeightOutput)
}

// LogThreatWeightArrayInput is an input type that accepts LogThreatWeightArray and LogThreatWeightArrayOutput values.
// You can construct a concrete instance of `LogThreatWeightArrayInput` via:
//
//          LogThreatWeightArray{ LogThreatWeightArgs{...} }
type LogThreatWeightArrayInput interface {
	pulumi.Input

	ToLogThreatWeightArrayOutput() LogThreatWeightArrayOutput
	ToLogThreatWeightArrayOutputWithContext(context.Context) LogThreatWeightArrayOutput
}

type LogThreatWeightArray []LogThreatWeightInput

func (LogThreatWeightArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogThreatWeight)(nil)).Elem()
}

func (i LogThreatWeightArray) ToLogThreatWeightArrayOutput() LogThreatWeightArrayOutput {
	return i.ToLogThreatWeightArrayOutputWithContext(context.Background())
}

func (i LogThreatWeightArray) ToLogThreatWeightArrayOutputWithContext(ctx context.Context) LogThreatWeightArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogThreatWeightArrayOutput)
}

// LogThreatWeightMapInput is an input type that accepts LogThreatWeightMap and LogThreatWeightMapOutput values.
// You can construct a concrete instance of `LogThreatWeightMapInput` via:
//
//          LogThreatWeightMap{ "key": LogThreatWeightArgs{...} }
type LogThreatWeightMapInput interface {
	pulumi.Input

	ToLogThreatWeightMapOutput() LogThreatWeightMapOutput
	ToLogThreatWeightMapOutputWithContext(context.Context) LogThreatWeightMapOutput
}

type LogThreatWeightMap map[string]LogThreatWeightInput

func (LogThreatWeightMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogThreatWeight)(nil)).Elem()
}

func (i LogThreatWeightMap) ToLogThreatWeightMapOutput() LogThreatWeightMapOutput {
	return i.ToLogThreatWeightMapOutputWithContext(context.Background())
}

func (i LogThreatWeightMap) ToLogThreatWeightMapOutputWithContext(ctx context.Context) LogThreatWeightMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogThreatWeightMapOutput)
}

type LogThreatWeightOutput struct{ *pulumi.OutputState }

func (LogThreatWeightOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogThreatWeight)(nil)).Elem()
}

func (o LogThreatWeightOutput) ToLogThreatWeightOutput() LogThreatWeightOutput {
	return o
}

func (o LogThreatWeightOutput) ToLogThreatWeightOutputWithContext(ctx context.Context) LogThreatWeightOutput {
	return o
}

type LogThreatWeightArrayOutput struct{ *pulumi.OutputState }

func (LogThreatWeightArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogThreatWeight)(nil)).Elem()
}

func (o LogThreatWeightArrayOutput) ToLogThreatWeightArrayOutput() LogThreatWeightArrayOutput {
	return o
}

func (o LogThreatWeightArrayOutput) ToLogThreatWeightArrayOutputWithContext(ctx context.Context) LogThreatWeightArrayOutput {
	return o
}

func (o LogThreatWeightArrayOutput) Index(i pulumi.IntInput) LogThreatWeightOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogThreatWeight {
		return vs[0].([]*LogThreatWeight)[vs[1].(int)]
	}).(LogThreatWeightOutput)
}

type LogThreatWeightMapOutput struct{ *pulumi.OutputState }

func (LogThreatWeightMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogThreatWeight)(nil)).Elem()
}

func (o LogThreatWeightMapOutput) ToLogThreatWeightMapOutput() LogThreatWeightMapOutput {
	return o
}

func (o LogThreatWeightMapOutput) ToLogThreatWeightMapOutputWithContext(ctx context.Context) LogThreatWeightMapOutput {
	return o
}

func (o LogThreatWeightMapOutput) MapIndex(k pulumi.StringInput) LogThreatWeightOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogThreatWeight {
		return vs[0].(map[string]*LogThreatWeight)[vs[1].(string)]
	}).(LogThreatWeightOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogThreatWeightInput)(nil)).Elem(), &LogThreatWeight{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogThreatWeightArrayInput)(nil)).Elem(), LogThreatWeightArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogThreatWeightMapInput)(nil)).Elem(), LogThreatWeightMap{})
	pulumi.RegisterOutputType(LogThreatWeightOutput{})
	pulumi.RegisterOutputType(LogThreatWeightArrayOutput{})
	pulumi.RegisterOutputType(LogThreatWeightMapOutput{})
}
