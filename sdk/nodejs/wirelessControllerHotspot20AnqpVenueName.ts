// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class WirelessControllerHotspot20AnqpVenueName extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerHotspot20AnqpVenueName resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerHotspot20AnqpVenueNameState, opts?: pulumi.CustomResourceOptions): WirelessControllerHotspot20AnqpVenueName {
        return new WirelessControllerHotspot20AnqpVenueName(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerHotspot20AnqpVenueName:WirelessControllerHotspot20AnqpVenueName';

    /**
     * Returns true if the given object is an instance of WirelessControllerHotspot20AnqpVenueName.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerHotspot20AnqpVenueName {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerHotspot20AnqpVenueName.__pulumiType;
    }

    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly valueLists!: pulumi.Output<outputs.WirelessControllerHotspot20AnqpVenueNameValueList[] | undefined>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessControllerHotspot20AnqpVenueName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerHotspot20AnqpVenueNameArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerHotspot20AnqpVenueNameArgs | WirelessControllerHotspot20AnqpVenueNameState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerHotspot20AnqpVenueNameState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["valueLists"] = state ? state.valueLists : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelessControllerHotspot20AnqpVenueNameArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["valueLists"] = args ? args.valueLists : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerHotspot20AnqpVenueName.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerHotspot20AnqpVenueName resources.
 */
export interface WirelessControllerHotspot20AnqpVenueNameState {
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    valueLists?: pulumi.Input<pulumi.Input<inputs.WirelessControllerHotspot20AnqpVenueNameValueList>[]>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerHotspot20AnqpVenueName resource.
 */
export interface WirelessControllerHotspot20AnqpVenueNameArgs {
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    valueLists?: pulumi.Input<pulumi.Input<inputs.WirelessControllerHotspot20AnqpVenueNameValueList>[]>;
    vdomparam?: pulumi.Input<string>;
}
