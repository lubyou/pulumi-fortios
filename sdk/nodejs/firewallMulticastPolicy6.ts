// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class FirewallMulticastPolicy6 extends pulumi.CustomResource {
    /**
     * Get an existing FirewallMulticastPolicy6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallMulticastPolicy6State, opts?: pulumi.CustomResourceOptions): FirewallMulticastPolicy6 {
        return new FirewallMulticastPolicy6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallMulticastPolicy6:FirewallMulticastPolicy6';

    /**
     * Returns true if the given object is an instance of FirewallMulticastPolicy6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallMulticastPolicy6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallMulticastPolicy6.__pulumiType;
    }

    public readonly action!: pulumi.Output<string>;
    public readonly autoAsicOffload!: pulumi.Output<string>;
    public readonly comments!: pulumi.Output<string | undefined>;
    public readonly dstaddrs!: pulumi.Output<outputs.FirewallMulticastPolicy6Dstaddr[]>;
    public readonly dstintf!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly endPort!: pulumi.Output<number>;
    public readonly fosid!: pulumi.Output<number>;
    public readonly logtraffic!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly protocol!: pulumi.Output<number>;
    public readonly srcaddrs!: pulumi.Output<outputs.FirewallMulticastPolicy6Srcaddr[]>;
    public readonly srcintf!: pulumi.Output<string>;
    public readonly startPort!: pulumi.Output<number>;
    public readonly status!: pulumi.Output<string>;
    public readonly uuid!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallMulticastPolicy6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallMulticastPolicy6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallMulticastPolicy6Args | FirewallMulticastPolicy6State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallMulticastPolicy6State | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["autoAsicOffload"] = state ? state.autoAsicOffload : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["dstaddrs"] = state ? state.dstaddrs : undefined;
            resourceInputs["dstintf"] = state ? state.dstintf : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["endPort"] = state ? state.endPort : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["logtraffic"] = state ? state.logtraffic : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["srcaddrs"] = state ? state.srcaddrs : undefined;
            resourceInputs["srcintf"] = state ? state.srcintf : undefined;
            resourceInputs["startPort"] = state ? state.startPort : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallMulticastPolicy6Args | undefined;
            if ((!args || args.dstaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstaddrs'");
            }
            if ((!args || args.dstintf === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstintf'");
            }
            if ((!args || args.srcaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcaddrs'");
            }
            if ((!args || args.srcintf === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcintf'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["autoAsicOffload"] = args ? args.autoAsicOffload : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["dstaddrs"] = args ? args.dstaddrs : undefined;
            resourceInputs["dstintf"] = args ? args.dstintf : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["endPort"] = args ? args.endPort : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["logtraffic"] = args ? args.logtraffic : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["srcaddrs"] = args ? args.srcaddrs : undefined;
            resourceInputs["srcintf"] = args ? args.srcintf : undefined;
            resourceInputs["startPort"] = args ? args.startPort : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallMulticastPolicy6.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallMulticastPolicy6 resources.
 */
export interface FirewallMulticastPolicy6State {
    action?: pulumi.Input<string>;
    autoAsicOffload?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    dstaddrs?: pulumi.Input<pulumi.Input<inputs.FirewallMulticastPolicy6Dstaddr>[]>;
    dstintf?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    endPort?: pulumi.Input<number>;
    fosid?: pulumi.Input<number>;
    logtraffic?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    protocol?: pulumi.Input<number>;
    srcaddrs?: pulumi.Input<pulumi.Input<inputs.FirewallMulticastPolicy6Srcaddr>[]>;
    srcintf?: pulumi.Input<string>;
    startPort?: pulumi.Input<number>;
    status?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallMulticastPolicy6 resource.
 */
export interface FirewallMulticastPolicy6Args {
    action?: pulumi.Input<string>;
    autoAsicOffload?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    dstaddrs: pulumi.Input<pulumi.Input<inputs.FirewallMulticastPolicy6Dstaddr>[]>;
    dstintf: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    endPort?: pulumi.Input<number>;
    fosid?: pulumi.Input<number>;
    logtraffic?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    protocol?: pulumi.Input<number>;
    srcaddrs: pulumi.Input<pulumi.Input<inputs.FirewallMulticastPolicy6Srcaddr>[]>;
    srcintf: pulumi.Input<string>;
    startPort?: pulumi.Input<number>;
    status?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
