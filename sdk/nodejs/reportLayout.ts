// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Report layout configuration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.ReportLayout("trname", {
 *     cutoffOption: "run-time",
 *     cutoffTime: "00:00",
 *     day: "sunday",
 *     emailSend: "disable",
 *     format: "pdf",
 *     maxPdfReport: 31,
 *     options: "include-table-of-content view-chart-as-heading",
 *     scheduleType: "daily",
 *     styleTheme: "default-report",
 *     time: "00:00",
 *     title: "FortiGate System Analysis Report",
 * });
 * ```
 *
 * ## Import
 *
 * Report Layout can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/reportLayout:ReportLayout labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class ReportLayout extends pulumi.CustomResource {
    /**
     * Get an existing ReportLayout resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReportLayoutState, opts?: pulumi.CustomResourceOptions): ReportLayout {
        return new ReportLayout(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/reportLayout:ReportLayout';

    /**
     * Returns true if the given object is an instance of ReportLayout.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReportLayout {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReportLayout.__pulumiType;
    }

    /**
     * Configure report body item. The structure of `bodyItem` block is documented below.
     */
    public readonly bodyItems!: pulumi.Output<outputs.ReportLayoutBodyItem[] | undefined>;
    /**
     * Cutoff-option is either run-time or custom. Valid values: `run-time`, `custom`.
     */
    public readonly cutoffOption!: pulumi.Output<string>;
    /**
     * Custom cutoff time to generate report [hh:mm].
     */
    public readonly cutoffTime!: pulumi.Output<string>;
    /**
     * Schedule days of week to generate report. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
     */
    public readonly day!: pulumi.Output<string>;
    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Email recipients for generated reports.
     */
    public readonly emailRecipients!: pulumi.Output<string>;
    /**
     * Enable/disable sending emails after reports are generated. Valid values: `enable`, `disable`.
     */
    public readonly emailSend!: pulumi.Output<string>;
    /**
     * Report format. Valid values: `pdf`.
     */
    public readonly format!: pulumi.Output<string>;
    /**
     * Maximum number of PDF reports to keep at one time (oldest report is overwritten).
     */
    public readonly maxPdfReport!: pulumi.Output<number>;
    /**
     * Field name that match field of parameters defined in dataset.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Report page options. Valid values: `header-on-first-page`, `footer-on-first-page`.
     */
    public readonly options!: pulumi.Output<string>;
    /**
     * Configure report page. The structure of `page` block is documented below.
     */
    public readonly page!: pulumi.Output<outputs.ReportLayoutPage | undefined>;
    /**
     * Report schedule type. Valid values: `demand`, `daily`, `weekly`.
     */
    public readonly scheduleType!: pulumi.Output<string>;
    /**
     * Report style theme.
     */
    public readonly styleTheme!: pulumi.Output<string>;
    /**
     * Report subtitle.
     */
    public readonly subtitle!: pulumi.Output<string>;
    /**
     * Schedule time to generate report [hh:mm].
     */
    public readonly time!: pulumi.Output<string>;
    /**
     * Report section title.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a ReportLayout resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReportLayoutArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReportLayoutArgs | ReportLayoutState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReportLayoutState | undefined;
            inputs["bodyItems"] = state ? state.bodyItems : undefined;
            inputs["cutoffOption"] = state ? state.cutoffOption : undefined;
            inputs["cutoffTime"] = state ? state.cutoffTime : undefined;
            inputs["day"] = state ? state.day : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["emailRecipients"] = state ? state.emailRecipients : undefined;
            inputs["emailSend"] = state ? state.emailSend : undefined;
            inputs["format"] = state ? state.format : undefined;
            inputs["maxPdfReport"] = state ? state.maxPdfReport : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["options"] = state ? state.options : undefined;
            inputs["page"] = state ? state.page : undefined;
            inputs["scheduleType"] = state ? state.scheduleType : undefined;
            inputs["styleTheme"] = state ? state.styleTheme : undefined;
            inputs["subtitle"] = state ? state.subtitle : undefined;
            inputs["time"] = state ? state.time : undefined;
            inputs["title"] = state ? state.title : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ReportLayoutArgs | undefined;
            if ((!args || args.styleTheme === undefined) && !opts.urn) {
                throw new Error("Missing required property 'styleTheme'");
            }
            inputs["bodyItems"] = args ? args.bodyItems : undefined;
            inputs["cutoffOption"] = args ? args.cutoffOption : undefined;
            inputs["cutoffTime"] = args ? args.cutoffTime : undefined;
            inputs["day"] = args ? args.day : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["emailRecipients"] = args ? args.emailRecipients : undefined;
            inputs["emailSend"] = args ? args.emailSend : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["maxPdfReport"] = args ? args.maxPdfReport : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["page"] = args ? args.page : undefined;
            inputs["scheduleType"] = args ? args.scheduleType : undefined;
            inputs["styleTheme"] = args ? args.styleTheme : undefined;
            inputs["subtitle"] = args ? args.subtitle : undefined;
            inputs["time"] = args ? args.time : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ReportLayout.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReportLayout resources.
 */
export interface ReportLayoutState {
    /**
     * Configure report body item. The structure of `bodyItem` block is documented below.
     */
    bodyItems?: pulumi.Input<pulumi.Input<inputs.ReportLayoutBodyItem>[]>;
    /**
     * Cutoff-option is either run-time or custom. Valid values: `run-time`, `custom`.
     */
    cutoffOption?: pulumi.Input<string>;
    /**
     * Custom cutoff time to generate report [hh:mm].
     */
    cutoffTime?: pulumi.Input<string>;
    /**
     * Schedule days of week to generate report. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
     */
    day?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Email recipients for generated reports.
     */
    emailRecipients?: pulumi.Input<string>;
    /**
     * Enable/disable sending emails after reports are generated. Valid values: `enable`, `disable`.
     */
    emailSend?: pulumi.Input<string>;
    /**
     * Report format. Valid values: `pdf`.
     */
    format?: pulumi.Input<string>;
    /**
     * Maximum number of PDF reports to keep at one time (oldest report is overwritten).
     */
    maxPdfReport?: pulumi.Input<number>;
    /**
     * Field name that match field of parameters defined in dataset.
     */
    name?: pulumi.Input<string>;
    /**
     * Report page options. Valid values: `header-on-first-page`, `footer-on-first-page`.
     */
    options?: pulumi.Input<string>;
    /**
     * Configure report page. The structure of `page` block is documented below.
     */
    page?: pulumi.Input<inputs.ReportLayoutPage>;
    /**
     * Report schedule type. Valid values: `demand`, `daily`, `weekly`.
     */
    scheduleType?: pulumi.Input<string>;
    /**
     * Report style theme.
     */
    styleTheme?: pulumi.Input<string>;
    /**
     * Report subtitle.
     */
    subtitle?: pulumi.Input<string>;
    /**
     * Schedule time to generate report [hh:mm].
     */
    time?: pulumi.Input<string>;
    /**
     * Report section title.
     */
    title?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReportLayout resource.
 */
export interface ReportLayoutArgs {
    /**
     * Configure report body item. The structure of `bodyItem` block is documented below.
     */
    bodyItems?: pulumi.Input<pulumi.Input<inputs.ReportLayoutBodyItem>[]>;
    /**
     * Cutoff-option is either run-time or custom. Valid values: `run-time`, `custom`.
     */
    cutoffOption?: pulumi.Input<string>;
    /**
     * Custom cutoff time to generate report [hh:mm].
     */
    cutoffTime?: pulumi.Input<string>;
    /**
     * Schedule days of week to generate report. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
     */
    day?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Email recipients for generated reports.
     */
    emailRecipients?: pulumi.Input<string>;
    /**
     * Enable/disable sending emails after reports are generated. Valid values: `enable`, `disable`.
     */
    emailSend?: pulumi.Input<string>;
    /**
     * Report format. Valid values: `pdf`.
     */
    format?: pulumi.Input<string>;
    /**
     * Maximum number of PDF reports to keep at one time (oldest report is overwritten).
     */
    maxPdfReport?: pulumi.Input<number>;
    /**
     * Field name that match field of parameters defined in dataset.
     */
    name?: pulumi.Input<string>;
    /**
     * Report page options. Valid values: `header-on-first-page`, `footer-on-first-page`.
     */
    options?: pulumi.Input<string>;
    /**
     * Configure report page. The structure of `page` block is documented below.
     */
    page?: pulumi.Input<inputs.ReportLayoutPage>;
    /**
     * Report schedule type. Valid values: `demand`, `daily`, `weekly`.
     */
    scheduleType?: pulumi.Input<string>;
    /**
     * Report style theme.
     */
    styleTheme: pulumi.Input<string>;
    /**
     * Report subtitle.
     */
    subtitle?: pulumi.Input<string>;
    /**
     * Schedule time to generate report [hh:mm].
     */
    time?: pulumi.Input<string>;
    /**
     * Report section title.
     */
    title?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
