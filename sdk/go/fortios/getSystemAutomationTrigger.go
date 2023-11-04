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

func LookupSystemAutomationTrigger(ctx *pulumi.Context, args *LookupSystemAutomationTriggerArgs, opts ...pulumi.InvokeOption) (*LookupSystemAutomationTriggerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSystemAutomationTriggerResult
	err := ctx.Invoke("fortios:index/getSystemAutomationTrigger:GetSystemAutomationTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemAutomationTrigger.
type LookupSystemAutomationTriggerArgs struct {
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemAutomationTrigger.
type LookupSystemAutomationTriggerResult struct {
	Description         string                            `pulumi:"description"`
	EventType           string                            `pulumi:"eventType"`
	FabricEventName     string                            `pulumi:"fabricEventName"`
	FabricEventSeverity string                            `pulumi:"fabricEventSeverity"`
	FazEventName        string                            `pulumi:"fazEventName"`
	FazEventSeverity    string                            `pulumi:"fazEventSeverity"`
	FazEventTags        string                            `pulumi:"fazEventTags"`
	Fields              []GetSystemAutomationTriggerField `pulumi:"fields"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                                 `pulumi:"id"`
	IocLevel         string                                 `pulumi:"iocLevel"`
	LicenseType      string                                 `pulumi:"licenseType"`
	Logid            int                                    `pulumi:"logid"`
	LogidBlocks      []GetSystemAutomationTriggerLogidBlock `pulumi:"logidBlocks"`
	Name             string                                 `pulumi:"name"`
	ReportType       string                                 `pulumi:"reportType"`
	Serial           string                                 `pulumi:"serial"`
	TriggerDatetime  string                                 `pulumi:"triggerDatetime"`
	TriggerDay       int                                    `pulumi:"triggerDay"`
	TriggerFrequency string                                 `pulumi:"triggerFrequency"`
	TriggerHour      int                                    `pulumi:"triggerHour"`
	TriggerMinute    int                                    `pulumi:"triggerMinute"`
	TriggerType      string                                 `pulumi:"triggerType"`
	TriggerWeekday   string                                 `pulumi:"triggerWeekday"`
	Vdomparam        *string                                `pulumi:"vdomparam"`
	Vdoms            []GetSystemAutomationTriggerVdom       `pulumi:"vdoms"`
}

func LookupSystemAutomationTriggerOutput(ctx *pulumi.Context, args LookupSystemAutomationTriggerOutputArgs, opts ...pulumi.InvokeOption) LookupSystemAutomationTriggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemAutomationTriggerResult, error) {
			args := v.(LookupSystemAutomationTriggerArgs)
			r, err := LookupSystemAutomationTrigger(ctx, &args, opts...)
			var s LookupSystemAutomationTriggerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemAutomationTriggerResultOutput)
}

// A collection of arguments for invoking GetSystemAutomationTrigger.
type LookupSystemAutomationTriggerOutputArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
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

func (o LookupSystemAutomationTriggerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSystemAutomationTriggerResult] {
	return pulumix.Output[LookupSystemAutomationTriggerResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupSystemAutomationTriggerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) EventType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.EventType }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) FabricEventName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.FabricEventName }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) FabricEventSeverity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.FabricEventSeverity }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) FazEventName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.FazEventName }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) FazEventSeverity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.FazEventSeverity }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) FazEventTags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.FazEventTags }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) Fields() GetSystemAutomationTriggerFieldArrayOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) []GetSystemAutomationTriggerField { return v.Fields }).(GetSystemAutomationTriggerFieldArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemAutomationTriggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) IocLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.IocLevel }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.LicenseType }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) Logid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) int { return v.Logid }).(pulumi.IntOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) LogidBlocks() GetSystemAutomationTriggerLogidBlockArrayOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) []GetSystemAutomationTriggerLogidBlock {
		return v.LogidBlocks
	}).(GetSystemAutomationTriggerLogidBlockArrayOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) ReportType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.ReportType }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) Serial() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.Serial }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) TriggerDatetime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.TriggerDatetime }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) TriggerDay() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) int { return v.TriggerDay }).(pulumi.IntOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) TriggerFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.TriggerFrequency }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) TriggerHour() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) int { return v.TriggerHour }).(pulumi.IntOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) TriggerMinute() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) int { return v.TriggerMinute }).(pulumi.IntOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) TriggerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.TriggerType }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) TriggerWeekday() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) string { return v.TriggerWeekday }).(pulumi.StringOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LookupSystemAutomationTriggerResultOutput) Vdoms() GetSystemAutomationTriggerVdomArrayOutput {
	return o.ApplyT(func(v LookupSystemAutomationTriggerResult) []GetSystemAutomationTriggerVdom { return v.Vdoms }).(GetSystemAutomationTriggerVdomArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemAutomationTriggerResultOutput{})
}
