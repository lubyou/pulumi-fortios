// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class ReportChart extends pulumi.CustomResource {
    /**
     * Get an existing ReportChart resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReportChartState, opts?: pulumi.CustomResourceOptions): ReportChart {
        return new ReportChart(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/reportChart:ReportChart';

    /**
     * Returns true if the given object is an instance of ReportChart.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReportChart {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReportChart.__pulumiType;
    }

    public readonly background!: pulumi.Output<string>;
    public readonly category!: pulumi.Output<string>;
    public readonly categorySeries!: pulumi.Output<outputs.ReportChartCategorySeries>;
    public readonly colorPalette!: pulumi.Output<string>;
    public readonly columns!: pulumi.Output<outputs.ReportChartColumn[] | undefined>;
    public readonly comments!: pulumi.Output<string>;
    public readonly dataset!: pulumi.Output<string>;
    public readonly dimension!: pulumi.Output<string>;
    public readonly drillDownCharts!: pulumi.Output<outputs.ReportChartDrillDownChart[] | undefined>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly favorite!: pulumi.Output<string>;
    public readonly graphType!: pulumi.Output<string>;
    public readonly legend!: pulumi.Output<string>;
    public readonly legendFontSize!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public readonly period!: pulumi.Output<string>;
    public readonly policy!: pulumi.Output<number>;
    public readonly style!: pulumi.Output<string>;
    public readonly title!: pulumi.Output<string>;
    public readonly titleFontSize!: pulumi.Output<number>;
    public readonly type!: pulumi.Output<string>;
    public readonly valueSeries!: pulumi.Output<outputs.ReportChartValueSeries>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly xSeries!: pulumi.Output<outputs.ReportChartXSeries>;
    public readonly ySeries!: pulumi.Output<outputs.ReportChartYSeries>;

    /**
     * Create a ReportChart resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReportChartArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReportChartArgs | ReportChartState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReportChartState | undefined;
            resourceInputs["background"] = state ? state.background : undefined;
            resourceInputs["category"] = state ? state.category : undefined;
            resourceInputs["categorySeries"] = state ? state.categorySeries : undefined;
            resourceInputs["colorPalette"] = state ? state.colorPalette : undefined;
            resourceInputs["columns"] = state ? state.columns : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["dataset"] = state ? state.dataset : undefined;
            resourceInputs["dimension"] = state ? state.dimension : undefined;
            resourceInputs["drillDownCharts"] = state ? state.drillDownCharts : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["favorite"] = state ? state.favorite : undefined;
            resourceInputs["graphType"] = state ? state.graphType : undefined;
            resourceInputs["legend"] = state ? state.legend : undefined;
            resourceInputs["legendFontSize"] = state ? state.legendFontSize : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["style"] = state ? state.style : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["titleFontSize"] = state ? state.titleFontSize : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["valueSeries"] = state ? state.valueSeries : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["xSeries"] = state ? state.xSeries : undefined;
            resourceInputs["ySeries"] = state ? state.ySeries : undefined;
        } else {
            const args = argsOrState as ReportChartArgs | undefined;
            if ((!args || args.comments === undefined) && !opts.urn) {
                throw new Error("Missing required property 'comments'");
            }
            if ((!args || args.dataset === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataset'");
            }
            resourceInputs["background"] = args ? args.background : undefined;
            resourceInputs["category"] = args ? args.category : undefined;
            resourceInputs["categorySeries"] = args ? args.categorySeries : undefined;
            resourceInputs["colorPalette"] = args ? args.colorPalette : undefined;
            resourceInputs["columns"] = args ? args.columns : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["dataset"] = args ? args.dataset : undefined;
            resourceInputs["dimension"] = args ? args.dimension : undefined;
            resourceInputs["drillDownCharts"] = args ? args.drillDownCharts : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["favorite"] = args ? args.favorite : undefined;
            resourceInputs["graphType"] = args ? args.graphType : undefined;
            resourceInputs["legend"] = args ? args.legend : undefined;
            resourceInputs["legendFontSize"] = args ? args.legendFontSize : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["style"] = args ? args.style : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["titleFontSize"] = args ? args.titleFontSize : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["valueSeries"] = args ? args.valueSeries : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["xSeries"] = args ? args.xSeries : undefined;
            resourceInputs["ySeries"] = args ? args.ySeries : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReportChart.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReportChart resources.
 */
export interface ReportChartState {
    background?: pulumi.Input<string>;
    category?: pulumi.Input<string>;
    categorySeries?: pulumi.Input<inputs.ReportChartCategorySeries>;
    colorPalette?: pulumi.Input<string>;
    columns?: pulumi.Input<pulumi.Input<inputs.ReportChartColumn>[]>;
    comments?: pulumi.Input<string>;
    dataset?: pulumi.Input<string>;
    dimension?: pulumi.Input<string>;
    drillDownCharts?: pulumi.Input<pulumi.Input<inputs.ReportChartDrillDownChart>[]>;
    dynamicSortSubtable?: pulumi.Input<string>;
    favorite?: pulumi.Input<string>;
    graphType?: pulumi.Input<string>;
    legend?: pulumi.Input<string>;
    legendFontSize?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    period?: pulumi.Input<string>;
    policy?: pulumi.Input<number>;
    style?: pulumi.Input<string>;
    title?: pulumi.Input<string>;
    titleFontSize?: pulumi.Input<number>;
    type?: pulumi.Input<string>;
    valueSeries?: pulumi.Input<inputs.ReportChartValueSeries>;
    vdomparam?: pulumi.Input<string>;
    xSeries?: pulumi.Input<inputs.ReportChartXSeries>;
    ySeries?: pulumi.Input<inputs.ReportChartYSeries>;
}

/**
 * The set of arguments for constructing a ReportChart resource.
 */
export interface ReportChartArgs {
    background?: pulumi.Input<string>;
    category?: pulumi.Input<string>;
    categorySeries?: pulumi.Input<inputs.ReportChartCategorySeries>;
    colorPalette?: pulumi.Input<string>;
    columns?: pulumi.Input<pulumi.Input<inputs.ReportChartColumn>[]>;
    comments: pulumi.Input<string>;
    dataset: pulumi.Input<string>;
    dimension?: pulumi.Input<string>;
    drillDownCharts?: pulumi.Input<pulumi.Input<inputs.ReportChartDrillDownChart>[]>;
    dynamicSortSubtable?: pulumi.Input<string>;
    favorite?: pulumi.Input<string>;
    graphType?: pulumi.Input<string>;
    legend?: pulumi.Input<string>;
    legendFontSize?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    period?: pulumi.Input<string>;
    policy?: pulumi.Input<number>;
    style?: pulumi.Input<string>;
    title?: pulumi.Input<string>;
    titleFontSize?: pulumi.Input<number>;
    type?: pulumi.Input<string>;
    valueSeries?: pulumi.Input<inputs.ReportChartValueSeries>;
    vdomparam?: pulumi.Input<string>;
    xSeries?: pulumi.Input<inputs.ReportChartXSeries>;
    ySeries?: pulumi.Input<inputs.ReportChartYSeries>;
}
