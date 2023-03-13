// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReportLayout struct {
	pulumi.CustomResourceState

	BodyItems           ReportLayoutBodyItemArrayOutput `pulumi:"bodyItems"`
	CutoffOption        pulumi.StringOutput             `pulumi:"cutoffOption"`
	CutoffTime          pulumi.StringOutput             `pulumi:"cutoffTime"`
	Day                 pulumi.StringOutput             `pulumi:"day"`
	Description         pulumi.StringOutput             `pulumi:"description"`
	DynamicSortSubtable pulumi.StringPtrOutput          `pulumi:"dynamicSortSubtable"`
	EmailRecipients     pulumi.StringOutput             `pulumi:"emailRecipients"`
	EmailSend           pulumi.StringOutput             `pulumi:"emailSend"`
	Format              pulumi.StringOutput             `pulumi:"format"`
	MaxPdfReport        pulumi.IntOutput                `pulumi:"maxPdfReport"`
	Name                pulumi.StringOutput             `pulumi:"name"`
	Options             pulumi.StringOutput             `pulumi:"options"`
	Page                ReportLayoutPageOutput          `pulumi:"page"`
	ScheduleType        pulumi.StringOutput             `pulumi:"scheduleType"`
	StyleTheme          pulumi.StringOutput             `pulumi:"styleTheme"`
	Subtitle            pulumi.StringOutput             `pulumi:"subtitle"`
	Time                pulumi.StringOutput             `pulumi:"time"`
	Title               pulumi.StringOutput             `pulumi:"title"`
	Vdomparam           pulumi.StringPtrOutput          `pulumi:"vdomparam"`
}

// NewReportLayout registers a new resource with the given unique name, arguments, and options.
func NewReportLayout(ctx *pulumi.Context,
	name string, args *ReportLayoutArgs, opts ...pulumi.ResourceOption) (*ReportLayout, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StyleTheme == nil {
		return nil, errors.New("invalid value for required argument 'StyleTheme'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ReportLayout
	err := ctx.RegisterResource("fortios:index/reportLayout:ReportLayout", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReportLayout gets an existing ReportLayout resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReportLayout(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportLayoutState, opts ...pulumi.ResourceOption) (*ReportLayout, error) {
	var resource ReportLayout
	err := ctx.ReadResource("fortios:index/reportLayout:ReportLayout", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReportLayout resources.
type reportLayoutState struct {
	BodyItems           []ReportLayoutBodyItem `pulumi:"bodyItems"`
	CutoffOption        *string                `pulumi:"cutoffOption"`
	CutoffTime          *string                `pulumi:"cutoffTime"`
	Day                 *string                `pulumi:"day"`
	Description         *string                `pulumi:"description"`
	DynamicSortSubtable *string                `pulumi:"dynamicSortSubtable"`
	EmailRecipients     *string                `pulumi:"emailRecipients"`
	EmailSend           *string                `pulumi:"emailSend"`
	Format              *string                `pulumi:"format"`
	MaxPdfReport        *int                   `pulumi:"maxPdfReport"`
	Name                *string                `pulumi:"name"`
	Options             *string                `pulumi:"options"`
	Page                *ReportLayoutPage      `pulumi:"page"`
	ScheduleType        *string                `pulumi:"scheduleType"`
	StyleTheme          *string                `pulumi:"styleTheme"`
	Subtitle            *string                `pulumi:"subtitle"`
	Time                *string                `pulumi:"time"`
	Title               *string                `pulumi:"title"`
	Vdomparam           *string                `pulumi:"vdomparam"`
}

type ReportLayoutState struct {
	BodyItems           ReportLayoutBodyItemArrayInput
	CutoffOption        pulumi.StringPtrInput
	CutoffTime          pulumi.StringPtrInput
	Day                 pulumi.StringPtrInput
	Description         pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	EmailRecipients     pulumi.StringPtrInput
	EmailSend           pulumi.StringPtrInput
	Format              pulumi.StringPtrInput
	MaxPdfReport        pulumi.IntPtrInput
	Name                pulumi.StringPtrInput
	Options             pulumi.StringPtrInput
	Page                ReportLayoutPagePtrInput
	ScheduleType        pulumi.StringPtrInput
	StyleTheme          pulumi.StringPtrInput
	Subtitle            pulumi.StringPtrInput
	Time                pulumi.StringPtrInput
	Title               pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (ReportLayoutState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportLayoutState)(nil)).Elem()
}

type reportLayoutArgs struct {
	BodyItems           []ReportLayoutBodyItem `pulumi:"bodyItems"`
	CutoffOption        *string                `pulumi:"cutoffOption"`
	CutoffTime          *string                `pulumi:"cutoffTime"`
	Day                 *string                `pulumi:"day"`
	Description         *string                `pulumi:"description"`
	DynamicSortSubtable *string                `pulumi:"dynamicSortSubtable"`
	EmailRecipients     *string                `pulumi:"emailRecipients"`
	EmailSend           *string                `pulumi:"emailSend"`
	Format              *string                `pulumi:"format"`
	MaxPdfReport        *int                   `pulumi:"maxPdfReport"`
	Name                *string                `pulumi:"name"`
	Options             *string                `pulumi:"options"`
	Page                *ReportLayoutPage      `pulumi:"page"`
	ScheduleType        *string                `pulumi:"scheduleType"`
	StyleTheme          string                 `pulumi:"styleTheme"`
	Subtitle            *string                `pulumi:"subtitle"`
	Time                *string                `pulumi:"time"`
	Title               *string                `pulumi:"title"`
	Vdomparam           *string                `pulumi:"vdomparam"`
}

// The set of arguments for constructing a ReportLayout resource.
type ReportLayoutArgs struct {
	BodyItems           ReportLayoutBodyItemArrayInput
	CutoffOption        pulumi.StringPtrInput
	CutoffTime          pulumi.StringPtrInput
	Day                 pulumi.StringPtrInput
	Description         pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	EmailRecipients     pulumi.StringPtrInput
	EmailSend           pulumi.StringPtrInput
	Format              pulumi.StringPtrInput
	MaxPdfReport        pulumi.IntPtrInput
	Name                pulumi.StringPtrInput
	Options             pulumi.StringPtrInput
	Page                ReportLayoutPagePtrInput
	ScheduleType        pulumi.StringPtrInput
	StyleTheme          pulumi.StringInput
	Subtitle            pulumi.StringPtrInput
	Time                pulumi.StringPtrInput
	Title               pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (ReportLayoutArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportLayoutArgs)(nil)).Elem()
}

type ReportLayoutInput interface {
	pulumi.Input

	ToReportLayoutOutput() ReportLayoutOutput
	ToReportLayoutOutputWithContext(ctx context.Context) ReportLayoutOutput
}

func (*ReportLayout) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportLayout)(nil)).Elem()
}

func (i *ReportLayout) ToReportLayoutOutput() ReportLayoutOutput {
	return i.ToReportLayoutOutputWithContext(context.Background())
}

func (i *ReportLayout) ToReportLayoutOutputWithContext(ctx context.Context) ReportLayoutOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportLayoutOutput)
}

// ReportLayoutArrayInput is an input type that accepts ReportLayoutArray and ReportLayoutArrayOutput values.
// You can construct a concrete instance of `ReportLayoutArrayInput` via:
//
//	ReportLayoutArray{ ReportLayoutArgs{...} }
type ReportLayoutArrayInput interface {
	pulumi.Input

	ToReportLayoutArrayOutput() ReportLayoutArrayOutput
	ToReportLayoutArrayOutputWithContext(context.Context) ReportLayoutArrayOutput
}

type ReportLayoutArray []ReportLayoutInput

func (ReportLayoutArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReportLayout)(nil)).Elem()
}

func (i ReportLayoutArray) ToReportLayoutArrayOutput() ReportLayoutArrayOutput {
	return i.ToReportLayoutArrayOutputWithContext(context.Background())
}

func (i ReportLayoutArray) ToReportLayoutArrayOutputWithContext(ctx context.Context) ReportLayoutArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportLayoutArrayOutput)
}

// ReportLayoutMapInput is an input type that accepts ReportLayoutMap and ReportLayoutMapOutput values.
// You can construct a concrete instance of `ReportLayoutMapInput` via:
//
//	ReportLayoutMap{ "key": ReportLayoutArgs{...} }
type ReportLayoutMapInput interface {
	pulumi.Input

	ToReportLayoutMapOutput() ReportLayoutMapOutput
	ToReportLayoutMapOutputWithContext(context.Context) ReportLayoutMapOutput
}

type ReportLayoutMap map[string]ReportLayoutInput

func (ReportLayoutMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReportLayout)(nil)).Elem()
}

func (i ReportLayoutMap) ToReportLayoutMapOutput() ReportLayoutMapOutput {
	return i.ToReportLayoutMapOutputWithContext(context.Background())
}

func (i ReportLayoutMap) ToReportLayoutMapOutputWithContext(ctx context.Context) ReportLayoutMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportLayoutMapOutput)
}

type ReportLayoutOutput struct{ *pulumi.OutputState }

func (ReportLayoutOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportLayout)(nil)).Elem()
}

func (o ReportLayoutOutput) ToReportLayoutOutput() ReportLayoutOutput {
	return o
}

func (o ReportLayoutOutput) ToReportLayoutOutputWithContext(ctx context.Context) ReportLayoutOutput {
	return o
}

func (o ReportLayoutOutput) BodyItems() ReportLayoutBodyItemArrayOutput {
	return o.ApplyT(func(v *ReportLayout) ReportLayoutBodyItemArrayOutput { return v.BodyItems }).(ReportLayoutBodyItemArrayOutput)
}

func (o ReportLayoutOutput) CutoffOption() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringOutput { return v.CutoffOption }).(pulumi.StringOutput)
}

func (o ReportLayoutOutput) CutoffTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringOutput { return v.CutoffTime }).(pulumi.StringOutput)
}

func (o ReportLayoutOutput) Day() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringOutput { return v.Day }).(pulumi.StringOutput)
}

func (o ReportLayoutOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o ReportLayoutOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o ReportLayoutOutput) EmailRecipients() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringOutput { return v.EmailRecipients }).(pulumi.StringOutput)
}

func (o ReportLayoutOutput) EmailSend() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringOutput { return v.EmailSend }).(pulumi.StringOutput)
}

func (o ReportLayoutOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

func (o ReportLayoutOutput) MaxPdfReport() pulumi.IntOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.IntOutput { return v.MaxPdfReport }).(pulumi.IntOutput)
}

func (o ReportLayoutOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReportLayoutOutput) Options() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringOutput { return v.Options }).(pulumi.StringOutput)
}

func (o ReportLayoutOutput) Page() ReportLayoutPageOutput {
	return o.ApplyT(func(v *ReportLayout) ReportLayoutPageOutput { return v.Page }).(ReportLayoutPageOutput)
}

func (o ReportLayoutOutput) ScheduleType() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringOutput { return v.ScheduleType }).(pulumi.StringOutput)
}

func (o ReportLayoutOutput) StyleTheme() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringOutput { return v.StyleTheme }).(pulumi.StringOutput)
}

func (o ReportLayoutOutput) Subtitle() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringOutput { return v.Subtitle }).(pulumi.StringOutput)
}

func (o ReportLayoutOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringOutput { return v.Time }).(pulumi.StringOutput)
}

func (o ReportLayoutOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

func (o ReportLayoutOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportLayout) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ReportLayoutArrayOutput struct{ *pulumi.OutputState }

func (ReportLayoutArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReportLayout)(nil)).Elem()
}

func (o ReportLayoutArrayOutput) ToReportLayoutArrayOutput() ReportLayoutArrayOutput {
	return o
}

func (o ReportLayoutArrayOutput) ToReportLayoutArrayOutputWithContext(ctx context.Context) ReportLayoutArrayOutput {
	return o
}

func (o ReportLayoutArrayOutput) Index(i pulumi.IntInput) ReportLayoutOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReportLayout {
		return vs[0].([]*ReportLayout)[vs[1].(int)]
	}).(ReportLayoutOutput)
}

type ReportLayoutMapOutput struct{ *pulumi.OutputState }

func (ReportLayoutMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReportLayout)(nil)).Elem()
}

func (o ReportLayoutMapOutput) ToReportLayoutMapOutput() ReportLayoutMapOutput {
	return o
}

func (o ReportLayoutMapOutput) ToReportLayoutMapOutputWithContext(ctx context.Context) ReportLayoutMapOutput {
	return o
}

func (o ReportLayoutMapOutput) MapIndex(k pulumi.StringInput) ReportLayoutOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReportLayout {
		return vs[0].(map[string]*ReportLayout)[vs[1].(string)]
	}).(ReportLayoutOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReportLayoutInput)(nil)).Elem(), &ReportLayout{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReportLayoutArrayInput)(nil)).Elem(), ReportLayoutArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReportLayoutMapInput)(nil)).Elem(), ReportLayoutMap{})
	pulumi.RegisterOutputType(ReportLayoutOutput{})
	pulumi.RegisterOutputType(ReportLayoutArrayOutput{})
	pulumi.RegisterOutputType(ReportLayoutMapOutput{})
}
