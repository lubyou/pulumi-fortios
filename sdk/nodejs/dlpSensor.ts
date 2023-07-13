// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class DlpSensor extends pulumi.CustomResource {
    /**
     * Get an existing DlpSensor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DlpSensorState, opts?: pulumi.CustomResourceOptions): DlpSensor {
        return new DlpSensor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/dlpSensor:DlpSensor';

    /**
     * Returns true if the given object is an instance of DlpSensor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DlpSensor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DlpSensor.__pulumiType;
    }

    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly dlpLog!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly entries!: pulumi.Output<outputs.DlpSensorEntry[] | undefined>;
    public readonly eval!: pulumi.Output<string>;
    public readonly extendedLog!: pulumi.Output<string>;
    public readonly featureSet!: pulumi.Output<string>;
    public readonly filters!: pulumi.Output<outputs.DlpSensorFilter[] | undefined>;
    public readonly flowBased!: pulumi.Output<string>;
    public readonly fullArchiveProto!: pulumi.Output<string>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly matchType!: pulumi.Output<string>;
    public readonly nacQuarLog!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly options!: pulumi.Output<string>;
    public readonly replacemsgGroup!: pulumi.Output<string>;
    public readonly summaryProto!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a DlpSensor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DlpSensorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DlpSensorArgs | DlpSensorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DlpSensorState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dlpLog"] = state ? state.dlpLog : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = state ? state.entries : undefined;
            resourceInputs["eval"] = state ? state.eval : undefined;
            resourceInputs["extendedLog"] = state ? state.extendedLog : undefined;
            resourceInputs["featureSet"] = state ? state.featureSet : undefined;
            resourceInputs["filters"] = state ? state.filters : undefined;
            resourceInputs["flowBased"] = state ? state.flowBased : undefined;
            resourceInputs["fullArchiveProto"] = state ? state.fullArchiveProto : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["matchType"] = state ? state.matchType : undefined;
            resourceInputs["nacQuarLog"] = state ? state.nacQuarLog : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            resourceInputs["summaryProto"] = state ? state.summaryProto : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as DlpSensorArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dlpLog"] = args ? args.dlpLog : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = args ? args.entries : undefined;
            resourceInputs["eval"] = args ? args.eval : undefined;
            resourceInputs["extendedLog"] = args ? args.extendedLog : undefined;
            resourceInputs["featureSet"] = args ? args.featureSet : undefined;
            resourceInputs["filters"] = args ? args.filters : undefined;
            resourceInputs["flowBased"] = args ? args.flowBased : undefined;
            resourceInputs["fullArchiveProto"] = args ? args.fullArchiveProto : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["matchType"] = args ? args.matchType : undefined;
            resourceInputs["nacQuarLog"] = args ? args.nacQuarLog : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            resourceInputs["summaryProto"] = args ? args.summaryProto : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DlpSensor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DlpSensor resources.
 */
export interface DlpSensorState {
    comment?: pulumi.Input<string>;
    dlpLog?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    entries?: pulumi.Input<pulumi.Input<inputs.DlpSensorEntry>[]>;
    eval?: pulumi.Input<string>;
    extendedLog?: pulumi.Input<string>;
    featureSet?: pulumi.Input<string>;
    filters?: pulumi.Input<pulumi.Input<inputs.DlpSensorFilter>[]>;
    flowBased?: pulumi.Input<string>;
    fullArchiveProto?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    matchType?: pulumi.Input<string>;
    nacQuarLog?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    options?: pulumi.Input<string>;
    replacemsgGroup?: pulumi.Input<string>;
    summaryProto?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DlpSensor resource.
 */
export interface DlpSensorArgs {
    comment?: pulumi.Input<string>;
    dlpLog?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    entries?: pulumi.Input<pulumi.Input<inputs.DlpSensorEntry>[]>;
    eval?: pulumi.Input<string>;
    extendedLog?: pulumi.Input<string>;
    featureSet?: pulumi.Input<string>;
    filters?: pulumi.Input<pulumi.Input<inputs.DlpSensorFilter>[]>;
    flowBased?: pulumi.Input<string>;
    fullArchiveProto?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    matchType?: pulumi.Input<string>;
    nacQuarLog?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    options?: pulumi.Input<string>;
    replacemsgGroup?: pulumi.Input<string>;
    summaryProto?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
