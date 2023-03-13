// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SwitchControllerVlanPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerVlanPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerVlanPolicyState, opts?: pulumi.CustomResourceOptions): SwitchControllerVlanPolicy {
        return new SwitchControllerVlanPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerVlanPolicy:SwitchControllerVlanPolicy';

    /**
     * Returns true if the given object is an instance of SwitchControllerVlanPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerVlanPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerVlanPolicy.__pulumiType;
    }

    public readonly allowedVlans!: pulumi.Output<outputs.SwitchControllerVlanPolicyAllowedVlan[] | undefined>;
    public readonly allowedVlansAll!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string>;
    public readonly discardMode!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly fortilink!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly untaggedVlans!: pulumi.Output<outputs.SwitchControllerVlanPolicyUntaggedVlan[] | undefined>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly vlan!: pulumi.Output<string>;

    /**
     * Create a SwitchControllerVlanPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerVlanPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerVlanPolicyArgs | SwitchControllerVlanPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerVlanPolicyState | undefined;
            resourceInputs["allowedVlans"] = state ? state.allowedVlans : undefined;
            resourceInputs["allowedVlansAll"] = state ? state.allowedVlansAll : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["discardMode"] = state ? state.discardMode : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fortilink"] = state ? state.fortilink : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["untaggedVlans"] = state ? state.untaggedVlans : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vlan"] = state ? state.vlan : undefined;
        } else {
            const args = argsOrState as SwitchControllerVlanPolicyArgs | undefined;
            resourceInputs["allowedVlans"] = args ? args.allowedVlans : undefined;
            resourceInputs["allowedVlansAll"] = args ? args.allowedVlansAll : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["discardMode"] = args ? args.discardMode : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fortilink"] = args ? args.fortilink : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["untaggedVlans"] = args ? args.untaggedVlans : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vlan"] = args ? args.vlan : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwitchControllerVlanPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerVlanPolicy resources.
 */
export interface SwitchControllerVlanPolicyState {
    allowedVlans?: pulumi.Input<pulumi.Input<inputs.SwitchControllerVlanPolicyAllowedVlan>[]>;
    allowedVlansAll?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    discardMode?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    fortilink?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    untaggedVlans?: pulumi.Input<pulumi.Input<inputs.SwitchControllerVlanPolicyUntaggedVlan>[]>;
    vdomparam?: pulumi.Input<string>;
    vlan?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerVlanPolicy resource.
 */
export interface SwitchControllerVlanPolicyArgs {
    allowedVlans?: pulumi.Input<pulumi.Input<inputs.SwitchControllerVlanPolicyAllowedVlan>[]>;
    allowedVlansAll?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    discardMode?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    fortilink?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    untaggedVlans?: pulumi.Input<pulumi.Input<inputs.SwitchControllerVlanPolicyUntaggedVlan>[]>;
    vdomparam?: pulumi.Input<string>;
    vlan?: pulumi.Input<string>;
}
