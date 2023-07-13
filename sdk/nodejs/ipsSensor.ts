// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class IpsSensor extends pulumi.CustomResource {
    /**
     * Get an existing IpsSensor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpsSensorState, opts?: pulumi.CustomResourceOptions): IpsSensor {
        return new IpsSensor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/ipsSensor:IpsSensor';

    /**
     * Returns true if the given object is an instance of IpsSensor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpsSensor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpsSensor.__pulumiType;
    }

    public readonly blockMaliciousUrl!: pulumi.Output<string>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly entries!: pulumi.Output<outputs.IpsSensorEntry[] | undefined>;
    public readonly extendedLog!: pulumi.Output<string>;
    public readonly filters!: pulumi.Output<outputs.IpsSensorFilter[] | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly overrides!: pulumi.Output<outputs.IpsSensorOverride[] | undefined>;
    public readonly replacemsgGroup!: pulumi.Output<string>;
    public readonly scanBotnetConnections!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a IpsSensor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IpsSensorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpsSensorArgs | IpsSensorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpsSensorState | undefined;
            resourceInputs["blockMaliciousUrl"] = state ? state.blockMaliciousUrl : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = state ? state.entries : undefined;
            resourceInputs["extendedLog"] = state ? state.extendedLog : undefined;
            resourceInputs["filters"] = state ? state.filters : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["overrides"] = state ? state.overrides : undefined;
            resourceInputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            resourceInputs["scanBotnetConnections"] = state ? state.scanBotnetConnections : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as IpsSensorArgs | undefined;
            resourceInputs["blockMaliciousUrl"] = args ? args.blockMaliciousUrl : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["entries"] = args ? args.entries : undefined;
            resourceInputs["extendedLog"] = args ? args.extendedLog : undefined;
            resourceInputs["filters"] = args ? args.filters : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["overrides"] = args ? args.overrides : undefined;
            resourceInputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            resourceInputs["scanBotnetConnections"] = args ? args.scanBotnetConnections : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpsSensor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpsSensor resources.
 */
export interface IpsSensorState {
    blockMaliciousUrl?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    entries?: pulumi.Input<pulumi.Input<inputs.IpsSensorEntry>[]>;
    extendedLog?: pulumi.Input<string>;
    filters?: pulumi.Input<pulumi.Input<inputs.IpsSensorFilter>[]>;
    getAllTables?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    overrides?: pulumi.Input<pulumi.Input<inputs.IpsSensorOverride>[]>;
    replacemsgGroup?: pulumi.Input<string>;
    scanBotnetConnections?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpsSensor resource.
 */
export interface IpsSensorArgs {
    blockMaliciousUrl?: pulumi.Input<string>;
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    entries?: pulumi.Input<pulumi.Input<inputs.IpsSensorEntry>[]>;
    extendedLog?: pulumi.Input<string>;
    filters?: pulumi.Input<pulumi.Input<inputs.IpsSensorFilter>[]>;
    getAllTables?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    overrides?: pulumi.Input<pulumi.Input<inputs.IpsSensorOverride>[]>;
    replacemsgGroup?: pulumi.Input<string>;
    scanBotnetConnections?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
