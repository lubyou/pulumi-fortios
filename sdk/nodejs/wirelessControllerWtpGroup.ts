// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class WirelessControllerWtpGroup extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerWtpGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerWtpGroupState, opts?: pulumi.CustomResourceOptions): WirelessControllerWtpGroup {
        return new WirelessControllerWtpGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerWtpGroup:WirelessControllerWtpGroup';

    /**
     * Returns true if the given object is an instance of WirelessControllerWtpGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerWtpGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerWtpGroup.__pulumiType;
    }

    public readonly bleMajorId!: pulumi.Output<number>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly platformType!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly wtps!: pulumi.Output<outputs.WirelessControllerWtpGroupWtp[] | undefined>;

    /**
     * Create a WirelessControllerWtpGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerWtpGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerWtpGroupArgs | WirelessControllerWtpGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerWtpGroupState | undefined;
            resourceInputs["bleMajorId"] = state ? state.bleMajorId : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["platformType"] = state ? state.platformType : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["wtps"] = state ? state.wtps : undefined;
        } else {
            const args = argsOrState as WirelessControllerWtpGroupArgs | undefined;
            resourceInputs["bleMajorId"] = args ? args.bleMajorId : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["platformType"] = args ? args.platformType : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["wtps"] = args ? args.wtps : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WirelessControllerWtpGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerWtpGroup resources.
 */
export interface WirelessControllerWtpGroupState {
    bleMajorId?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    platformType?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    wtps?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWtpGroupWtp>[]>;
}

/**
 * The set of arguments for constructing a WirelessControllerWtpGroup resource.
 */
export interface WirelessControllerWtpGroupArgs {
    bleMajorId?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    platformType?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    wtps?: pulumi.Input<pulumi.Input<inputs.WirelessControllerWtpGroupWtp>[]>;
}
