// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SwitchControllerQosQosPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerQosQosPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerQosQosPolicyState, opts?: pulumi.CustomResourceOptions): SwitchControllerQosQosPolicy {
        return new SwitchControllerQosQosPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerQosQosPolicy:SwitchControllerQosQosPolicy';

    /**
     * Returns true if the given object is an instance of SwitchControllerQosQosPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerQosQosPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerQosQosPolicy.__pulumiType;
    }

    public readonly defaultCos!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public readonly queuePolicy!: pulumi.Output<string>;
    public readonly trustDot1pMap!: pulumi.Output<string>;
    public readonly trustIpDscpMap!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchControllerQosQosPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SwitchControllerQosQosPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerQosQosPolicyArgs | SwitchControllerQosQosPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerQosQosPolicyState | undefined;
            resourceInputs["defaultCos"] = state ? state.defaultCos : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["queuePolicy"] = state ? state.queuePolicy : undefined;
            resourceInputs["trustDot1pMap"] = state ? state.trustDot1pMap : undefined;
            resourceInputs["trustIpDscpMap"] = state ? state.trustIpDscpMap : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchControllerQosQosPolicyArgs | undefined;
            if ((!args || args.defaultCos === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultCos'");
            }
            resourceInputs["defaultCos"] = args ? args.defaultCos : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["queuePolicy"] = args ? args.queuePolicy : undefined;
            resourceInputs["trustDot1pMap"] = args ? args.trustDot1pMap : undefined;
            resourceInputs["trustIpDscpMap"] = args ? args.trustIpDscpMap : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwitchControllerQosQosPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerQosQosPolicy resources.
 */
export interface SwitchControllerQosQosPolicyState {
    defaultCos?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    queuePolicy?: pulumi.Input<string>;
    trustDot1pMap?: pulumi.Input<string>;
    trustIpDscpMap?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerQosQosPolicy resource.
 */
export interface SwitchControllerQosQosPolicyArgs {
    defaultCos: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    queuePolicy?: pulumi.Input<string>;
    trustDot1pMap?: pulumi.Input<string>;
    trustIpDscpMap?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
