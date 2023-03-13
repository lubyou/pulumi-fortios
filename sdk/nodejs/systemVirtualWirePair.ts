// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SystemVirtualWirePair extends pulumi.CustomResource {
    /**
     * Get an existing SystemVirtualWirePair resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemVirtualWirePairState, opts?: pulumi.CustomResourceOptions): SystemVirtualWirePair {
        return new SystemVirtualWirePair(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemVirtualWirePair:SystemVirtualWirePair';

    /**
     * Returns true if the given object is an instance of SystemVirtualWirePair.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemVirtualWirePair {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemVirtualWirePair.__pulumiType;
    }

    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly members!: pulumi.Output<outputs.SystemVirtualWirePairMember[]>;
    public readonly name!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly vlanFilter!: pulumi.Output<string>;
    public readonly wildcardVlan!: pulumi.Output<string>;

    /**
     * Create a SystemVirtualWirePair resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemVirtualWirePairArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemVirtualWirePairArgs | SystemVirtualWirePairState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemVirtualWirePairState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vlanFilter"] = state ? state.vlanFilter : undefined;
            resourceInputs["wildcardVlan"] = state ? state.wildcardVlan : undefined;
        } else {
            const args = argsOrState as SystemVirtualWirePairArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vlanFilter"] = args ? args.vlanFilter : undefined;
            resourceInputs["wildcardVlan"] = args ? args.wildcardVlan : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemVirtualWirePair.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemVirtualWirePair resources.
 */
export interface SystemVirtualWirePairState {
    dynamicSortSubtable?: pulumi.Input<string>;
    members?: pulumi.Input<pulumi.Input<inputs.SystemVirtualWirePairMember>[]>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    vlanFilter?: pulumi.Input<string>;
    wildcardVlan?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemVirtualWirePair resource.
 */
export interface SystemVirtualWirePairArgs {
    dynamicSortSubtable?: pulumi.Input<string>;
    members: pulumi.Input<pulumi.Input<inputs.SystemVirtualWirePairMember>[]>;
    name?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    vlanFilter?: pulumi.Input<string>;
    wildcardVlan?: pulumi.Input<string>;
}
