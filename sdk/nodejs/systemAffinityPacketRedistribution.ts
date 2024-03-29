// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SystemAffinityPacketRedistribution extends pulumi.CustomResource {
    /**
     * Get an existing SystemAffinityPacketRedistribution resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemAffinityPacketRedistributionState, opts?: pulumi.CustomResourceOptions): SystemAffinityPacketRedistribution {
        return new SystemAffinityPacketRedistribution(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemAffinityPacketRedistribution:SystemAffinityPacketRedistribution';

    /**
     * Returns true if the given object is an instance of SystemAffinityPacketRedistribution.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemAffinityPacketRedistribution {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemAffinityPacketRedistribution.__pulumiType;
    }

    public readonly affinityCpumask!: pulumi.Output<string>;
    public readonly fosid!: pulumi.Output<number>;
    public readonly interface!: pulumi.Output<string>;
    public readonly roundRobin!: pulumi.Output<string>;
    public readonly rxqid!: pulumi.Output<number>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemAffinityPacketRedistribution resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemAffinityPacketRedistributionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemAffinityPacketRedistributionArgs | SystemAffinityPacketRedistributionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemAffinityPacketRedistributionState | undefined;
            resourceInputs["affinityCpumask"] = state ? state.affinityCpumask : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["roundRobin"] = state ? state.roundRobin : undefined;
            resourceInputs["rxqid"] = state ? state.rxqid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemAffinityPacketRedistributionArgs | undefined;
            if ((!args || args.affinityCpumask === undefined) && !opts.urn) {
                throw new Error("Missing required property 'affinityCpumask'");
            }
            if ((!args || args.fosid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fosid'");
            }
            if ((!args || args.interface === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interface'");
            }
            if ((!args || args.rxqid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rxqid'");
            }
            resourceInputs["affinityCpumask"] = args ? args.affinityCpumask : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["roundRobin"] = args ? args.roundRobin : undefined;
            resourceInputs["rxqid"] = args ? args.rxqid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemAffinityPacketRedistribution.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemAffinityPacketRedistribution resources.
 */
export interface SystemAffinityPacketRedistributionState {
    affinityCpumask?: pulumi.Input<string>;
    fosid?: pulumi.Input<number>;
    interface?: pulumi.Input<string>;
    roundRobin?: pulumi.Input<string>;
    rxqid?: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemAffinityPacketRedistribution resource.
 */
export interface SystemAffinityPacketRedistributionArgs {
    affinityCpumask: pulumi.Input<string>;
    fosid: pulumi.Input<number>;
    interface: pulumi.Input<string>;
    roundRobin?: pulumi.Input<string>;
    rxqid: pulumi.Input<number>;
    vdomparam?: pulumi.Input<string>;
}
