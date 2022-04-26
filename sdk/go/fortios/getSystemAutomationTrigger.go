// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system automationtrigger
func LookupSystemAutomationTrigger(ctx *pulumi.Context, args *LookupSystemAutomationTriggerArgs, opts ...pulumi.InvokeOption) (*LookupSystemAutomationTriggerResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemAutomationTriggerResult
	err := ctx.Invoke("fortios:index/getSystemAutomationTrigger:GetSystemAutomationTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemAutomationTrigger.
type LookupSystemAutomationTriggerArgs struct {
	// Specify the name of the desired system automationtrigger.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemAutomationTrigger.
type LookupSystemAutomationTriggerResult struct {
	// Description.
	Description string `pulumi:"description"`
	// Event type.
	EventType string `pulumi:"eventType"`
	// Fabric connector event handler name.
	FabricEventName string `pulumi:"fabricEventName"`
	// Fabric connector event severity.
	FabricEventSeverity string `pulumi:"fabricEventSeverity"`
	// FortiAnalyzer event handler name.
	FazEventName string `pulumi:"fazEventName"`
	// FortiAnalyzer event severity.
	FazEventSeverity string `pulumi:"fazEventSeverity"`
	// FortiAnalyzer event tags.
	FazEventTags string `pulumi:"fazEventTags"`
	// Customized trigger field settings. The structure of `fields` block is documented below.
	Fields []GetSystemAutomationTriggerField `pulumi:"fields"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IOC threat level.
	IocLevel string `pulumi:"iocLevel"`
	// License type.
	LicenseType string `pulumi:"licenseType"`
	// Log ID to trigger event.
	Logid int `pulumi:"logid"`
	// Log ID to trigger event. Only applies on FortiOS v7.0.0+. The structure of `logidBlock` block is documented below.
	LogidBlocks []GetSystemAutomationTriggerLogidBlock `pulumi:"logidBlocks"`
	// Name.
	Name string `pulumi:"name"`
	// Security Rating report.
	ReportType string `pulumi:"reportType"`
	// Fabric connector serial number.
	Serial string `pulumi:"serial"`
	// Day within a month to trigger.
	TriggerDay int `pulumi:"triggerDay"`
	// Scheduled trigger frequency (default = daily).
	TriggerFrequency string `pulumi:"triggerFrequency"`
	// Hour of the day on which to trigger (0 - 23, default = 1).
	TriggerHour int `pulumi:"triggerHour"`
	// Minute of the hour on which to trigger (0 - 59, 60 to randomize).
	TriggerMinute int `pulumi:"triggerMinute"`
	// Trigger type.
	TriggerType string `pulumi:"triggerType"`
	// Day of week for trigger.
	TriggerWeekday string  `pulumi:"triggerWeekday"`
	Vdomparam      *string `pulumi:"vdomparam"`
}

func LookupSystemAutomationTriggerOutput(ctx *pulumi.Context, args LookupSystemAutomationTriggerOutputArgs, opts ...pulumi.InvokeOption) LookupSystemAutomationTriggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemAutomationTriggerResult, error) {
			args := v.(LookupSystemAutomationTriggerArgs)
			r, err := LookupSystemAutomationTrigger(ctx, &args, opts...)
			return *r, err
		}).(LookupSystemAutomationTriggerResultOutput)
}

// A collection of arguments for invoking GetSystemAutomationTrigger.
type LookupSystemAutomationTriggerOutputArgs struct {
	// Specify the name of the desired system automationtrigger.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemAutomationTriggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAutomationTriggerArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemAutomationTrigger.
type LookupSystemAutomationTriggerResultOutput struct{ *pulumi.OutputState }

func (LookupSystemAutomationTriggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAutomationTriggerResult)(nil)).Elem()
}

func (o LookupSystemAutomationTriggerResultOutput) ToLookupSystemAutomationTriggerResultOutput() LookupSystemAutomationTriggerResultOutput {
	return o
}

func (o LookupSystemAutomationTriggerResultOutput) ToLookupSystemAutomationTriggerResultOutputWithContext(ctx context.Context) LookupSystemAutomationTriggerResultOutput {
	return o
}

// Description.
func (o LookupSystemAutomationTriggerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.Description }).(pulumi.StringOutput)
}

// Event type.
func (o LookupSystemAutomationTriggerResultOutput) EventType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.EventType }).(pulumi.StringOutput)
}

// Fabric connector event handler name.
func (o LookupSystemAutomationTriggerResultOutput) FabricEventName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.FabricEventName }).(pulumi.StringOutput)
}

// Fabric connector event severity.
func (o LookupSystemAutomationTriggerResultOutput) FabricEventSeverity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.FabricEventSeverity }).(pulumi.StringOutput)
}

// FortiAnalyzer event handler name.
func (o LookupSystemAutomationTriggerResultOutput) FazEventName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.FazEventName }).(pulumi.StringOutput)
}

// FortiAnalyzer event severity.
func (o LookupSystemAutomationTriggerResultOutput) FazEventSeverity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.FazEventSeverity }).(pulumi.StringOutput)
}

// FortiAnalyzer event tags.
func (o LookupSystemAutomationTriggerResultOutput) FazEventTags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.FazEventTags }).(pulumi.StringOutput)
}

// Customized trigger field settings. The structure of `fields` block is documented below.
func (o LookupSystemAutomationTriggerResultOutput) Fields() GetSystemAutomationTriggerFieldArrayOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) []GetSystemAutomationTriggerField { return v.Fields }).(GetSystemAutomationTriggerFieldArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemAutomationTriggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.Id }).(pulumi.StringOutput)
}

// IOC threat level.
func (o LookupSystemAutomationTriggerResultOutput) IocLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.IocLevel }).(pulumi.StringOutput)
}

// License type.
func (o LookupSystemAutomationTriggerResultOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.LicenseType }).(pulumi.StringOutput)
}

// Log ID to trigger event.
func (o LookupSystemAutomationTriggerResultOutput) Logid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) int { return v.Logid }).(pulumi.IntOutput)
}

// Log ID to trigger event. Only applies on FortiOS v7.0.0+. The structure of `logidBlock` block is documented below.
func (o LookupSystemAutomationTriggerResultOutput) LogidBlocks() GetSystemAutomationTriggerLogidBlockArrayOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) []GetSystemAutomationTriggerLogidBlock {
		return v.LogidBlocks
	}).(GetSystemAutomationTriggerLogidBlockArrayOutput)
}

// Name.
func (o LookupSystemAutomationTriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Security Rating report.
func (o LookupSystemAutomationTriggerResultOutput) ReportType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.ReportType }).(pulumi.StringOutput)
}

// Fabric connector serial number.
func (o LookupSystemAutomationTriggerResultOutput) Serial() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.Serial }).(pulumi.StringOutput)
}

// Day within a month to trigger.
func (o LookupSystemAutomationTriggerResultOutput) TriggerDay() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) int { return v.TriggerDay }).(pulumi.IntOutput)
}

// Scheduled trigger frequency (default = daily).
func (o LookupSystemAutomationTriggerResultOutput) TriggerFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.TriggerFrequency }).(pulumi.StringOutput)
}

// Hour of the day on which to trigger (0 - 23, default = 1).
func (o LookupSystemAutomationTriggerResultOutput) TriggerHour() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) int { return v.TriggerHour }).(pulumi.IntOutput)
}

// Minute of the hour on which to trigger (0 - 59, 60 to randomize).
func (o LookupSystemAutomationTriggerResultOutput) TriggerMinute() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) int { return v.TriggerMinute }).(pulumi.IntOutput)
}

// Trigger type.
func (o LookupSystemAutomationTriggerResultOutput) TriggerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.TriggerType }).(pulumi.StringOutput)
}

// Day of week for trigger.
func (o LookupSystemAutomationTriggerResultOutput) TriggerWeekday() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.TriggerWeekday }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemAutomationTriggerResultOutput{})
}
