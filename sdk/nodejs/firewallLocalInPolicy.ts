// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class FirewallLocalInPolicy extends pulumi.CustomResource {
    /**
     * Get an existing FirewallLocalInPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallLocalInPolicyState, opts?: pulumi.CustomResourceOptions): FirewallLocalInPolicy {
        return new FirewallLocalInPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallLocalInPolicy:FirewallLocalInPolicy';

    /**
     * Returns true if the given object is an instance of FirewallLocalInPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallLocalInPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallLocalInPolicy.__pulumiType;
    }

    public readonly action!: pulumi.Output<string>;
    public readonly comments!: pulumi.Output<string | undefined>;
    public readonly dstaddrNegate!: pulumi.Output<string>;
    public readonly dstaddrs!: pulumi.Output<outputs.FirewallLocalInPolicyDstaddr[]>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly haMgmtIntfOnly!: pulumi.Output<string>;
    public readonly intf!: pulumi.Output<string>;
    public readonly policyid!: pulumi.Output<number>;
    public readonly schedule!: pulumi.Output<string>;
    public readonly serviceNegate!: pulumi.Output<string>;
    public readonly services!: pulumi.Output<outputs.FirewallLocalInPolicyService[] | undefined>;
    public readonly srcaddrNegate!: pulumi.Output<string>;
    public readonly srcaddrs!: pulumi.Output<outputs.FirewallLocalInPolicySrcaddr[]>;
    public readonly status!: pulumi.Output<string>;
    public readonly uuid!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly virtualPatch!: pulumi.Output<string>;

    /**
     * Create a FirewallLocalInPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallLocalInPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallLocalInPolicyArgs | FirewallLocalInPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallLocalInPolicyState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["dstaddrNegate"] = state ? state.dstaddrNegate : undefined;
            resourceInputs["dstaddrs"] = state ? state.dstaddrs : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["haMgmtIntfOnly"] = state ? state.haMgmtIntfOnly : undefined;
            resourceInputs["intf"] = state ? state.intf : undefined;
            resourceInputs["policyid"] = state ? state.policyid : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
            resourceInputs["serviceNegate"] = state ? state.serviceNegate : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["srcaddrNegate"] = state ? state.srcaddrNegate : undefined;
            resourceInputs["srcaddrs"] = state ? state.srcaddrs : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["virtualPatch"] = state ? state.virtualPatch : undefined;
        } else {
            const args = argsOrState as FirewallLocalInPolicyArgs | undefined;
            if ((!args || args.dstaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstaddrs'");
            }
            if ((!args || args.schedule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schedule'");
            }
            if ((!args || args.srcaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcaddrs'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["dstaddrNegate"] = args ? args.dstaddrNegate : undefined;
            resourceInputs["dstaddrs"] = args ? args.dstaddrs : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["haMgmtIntfOnly"] = args ? args.haMgmtIntfOnly : undefined;
            resourceInputs["intf"] = args ? args.intf : undefined;
            resourceInputs["policyid"] = args ? args.policyid : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["serviceNegate"] = args ? args.serviceNegate : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["srcaddrNegate"] = args ? args.srcaddrNegate : undefined;
            resourceInputs["srcaddrs"] = args ? args.srcaddrs : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["virtualPatch"] = args ? args.virtualPatch : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallLocalInPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallLocalInPolicy resources.
 */
export interface FirewallLocalInPolicyState {
    action?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    dstaddrNegate?: pulumi.Input<string>;
    dstaddrs?: pulumi.Input<pulumi.Input<inputs.FirewallLocalInPolicyDstaddr>[]>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    haMgmtIntfOnly?: pulumi.Input<string>;
    intf?: pulumi.Input<string>;
    policyid?: pulumi.Input<number>;
    schedule?: pulumi.Input<string>;
    serviceNegate?: pulumi.Input<string>;
    services?: pulumi.Input<pulumi.Input<inputs.FirewallLocalInPolicyService>[]>;
    srcaddrNegate?: pulumi.Input<string>;
    srcaddrs?: pulumi.Input<pulumi.Input<inputs.FirewallLocalInPolicySrcaddr>[]>;
    status?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    virtualPatch?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallLocalInPolicy resource.
 */
export interface FirewallLocalInPolicyArgs {
    action?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    dstaddrNegate?: pulumi.Input<string>;
    dstaddrs: pulumi.Input<pulumi.Input<inputs.FirewallLocalInPolicyDstaddr>[]>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    haMgmtIntfOnly?: pulumi.Input<string>;
    intf?: pulumi.Input<string>;
    policyid?: pulumi.Input<number>;
    schedule: pulumi.Input<string>;
    serviceNegate?: pulumi.Input<string>;
    services?: pulumi.Input<pulumi.Input<inputs.FirewallLocalInPolicyService>[]>;
    srcaddrNegate?: pulumi.Input<string>;
    srcaddrs: pulumi.Input<pulumi.Input<inputs.FirewallLocalInPolicySrcaddr>[]>;
    status?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    virtualPatch?: pulumi.Input<string>;
}
