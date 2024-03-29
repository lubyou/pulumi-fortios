// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class VpnIpsecConcentrator extends pulumi.CustomResource {
    /**
     * Get an existing VpnIpsecConcentrator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpnIpsecConcentratorState, opts?: pulumi.CustomResourceOptions): VpnIpsecConcentrator {
        return new VpnIpsecConcentrator(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/vpnIpsecConcentrator:VpnIpsecConcentrator';

    /**
     * Returns true if the given object is an instance of VpnIpsecConcentrator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnIpsecConcentrator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnIpsecConcentrator.__pulumiType;
    }

    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly fosid!: pulumi.Output<number>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly members!: pulumi.Output<outputs.VpnIpsecConcentratorMember[] | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly srcCheck!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a VpnIpsecConcentrator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VpnIpsecConcentratorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnIpsecConcentratorArgs | VpnIpsecConcentratorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpnIpsecConcentratorState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["srcCheck"] = state ? state.srcCheck : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as VpnIpsecConcentratorArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["srcCheck"] = args ? args.srcCheck : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpnIpsecConcentrator.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpnIpsecConcentrator resources.
 */
export interface VpnIpsecConcentratorState {
    dynamicSortSubtable?: pulumi.Input<string>;
    fosid?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    members?: pulumi.Input<pulumi.Input<inputs.VpnIpsecConcentratorMember>[]>;
    name?: pulumi.Input<string>;
    srcCheck?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpnIpsecConcentrator resource.
 */
export interface VpnIpsecConcentratorArgs {
    dynamicSortSubtable?: pulumi.Input<string>;
    fosid?: pulumi.Input<number>;
    getAllTables?: pulumi.Input<string>;
    members?: pulumi.Input<pulumi.Input<inputs.VpnIpsecConcentratorMember>[]>;
    name?: pulumi.Input<string>;
    srcCheck?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
