// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure user defined IPv4 local-in policies.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.FirewallLocalInPolicy("trname", {
 *     action: "accept",
 *     dstaddrs: [{
 *         name: "all",
 *     }],
 *     haMgmtIntfOnly: "disable",
 *     intf: "port4",
 *     policyid: 1,
 *     schedule: "always",
 *     services: [{
 *         name: "ALL",
 *     }],
 *     srcaddrs: [{
 *         name: "all",
 *     }],
 *     status: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall LocalInPolicy can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallLocalInPolicy:FirewallLocalInPolicy labelname {{policyid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
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

    /**
     * Action performed on traffic matching the policy (default = deny). Valid values: `accept`, `deny`.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Destination address object from available options. The structure of `dstaddr` block is documented below.
     */
    public readonly dstaddrs!: pulumi.Output<outputs.FirewallLocalInPolicyDstaddr[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable dedicating the HA management interface only for local-in policy. Valid values: `enable`, `disable`.
     */
    public readonly haMgmtIntfOnly!: pulumi.Output<string>;
    /**
     * Incoming interface name from available options.
     */
    public readonly intf!: pulumi.Output<string>;
    /**
     * User defined local in policy ID.
     */
    public readonly policyid!: pulumi.Output<number>;
    /**
     * Schedule object from available options.
     */
    public readonly schedule!: pulumi.Output<string>;
    /**
     * Service object from available options. The structure of `service` block is documented below.
     */
    public readonly services!: pulumi.Output<outputs.FirewallLocalInPolicyService[] | undefined>;
    /**
     * Source address object from available options. The structure of `srcaddr` block is documented below.
     */
    public readonly srcaddrs!: pulumi.Output<outputs.FirewallLocalInPolicySrcaddr[]>;
    /**
     * Enable/disable this local-in policy. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    public readonly uuid!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallLocalInPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallLocalInPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallLocalInPolicyArgs | FirewallLocalInPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallLocalInPolicyState | undefined;
            inputs["action"] = state ? state.action : undefined;
            inputs["comments"] = state ? state.comments : undefined;
            inputs["dstaddrs"] = state ? state.dstaddrs : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["haMgmtIntfOnly"] = state ? state.haMgmtIntfOnly : undefined;
            inputs["intf"] = state ? state.intf : undefined;
            inputs["policyid"] = state ? state.policyid : undefined;
            inputs["schedule"] = state ? state.schedule : undefined;
            inputs["services"] = state ? state.services : undefined;
            inputs["srcaddrs"] = state ? state.srcaddrs : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["uuid"] = state ? state.uuid : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
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
            inputs["action"] = args ? args.action : undefined;
            inputs["comments"] = args ? args.comments : undefined;
            inputs["dstaddrs"] = args ? args.dstaddrs : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["haMgmtIntfOnly"] = args ? args.haMgmtIntfOnly : undefined;
            inputs["intf"] = args ? args.intf : undefined;
            inputs["policyid"] = args ? args.policyid : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["services"] = args ? args.services : undefined;
            inputs["srcaddrs"] = args ? args.srcaddrs : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["uuid"] = args ? args.uuid : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FirewallLocalInPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallLocalInPolicy resources.
 */
export interface FirewallLocalInPolicyState {
    /**
     * Action performed on traffic matching the policy (default = deny). Valid values: `accept`, `deny`.
     */
    action?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Destination address object from available options. The structure of `dstaddr` block is documented below.
     */
    dstaddrs?: pulumi.Input<pulumi.Input<inputs.FirewallLocalInPolicyDstaddr>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable dedicating the HA management interface only for local-in policy. Valid values: `enable`, `disable`.
     */
    haMgmtIntfOnly?: pulumi.Input<string>;
    /**
     * Incoming interface name from available options.
     */
    intf?: pulumi.Input<string>;
    /**
     * User defined local in policy ID.
     */
    policyid?: pulumi.Input<number>;
    /**
     * Schedule object from available options.
     */
    schedule?: pulumi.Input<string>;
    /**
     * Service object from available options. The structure of `service` block is documented below.
     */
    services?: pulumi.Input<pulumi.Input<inputs.FirewallLocalInPolicyService>[]>;
    /**
     * Source address object from available options. The structure of `srcaddr` block is documented below.
     */
    srcaddrs?: pulumi.Input<pulumi.Input<inputs.FirewallLocalInPolicySrcaddr>[]>;
    /**
     * Enable/disable this local-in policy. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallLocalInPolicy resource.
 */
export interface FirewallLocalInPolicyArgs {
    /**
     * Action performed on traffic matching the policy (default = deny). Valid values: `accept`, `deny`.
     */
    action?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Destination address object from available options. The structure of `dstaddr` block is documented below.
     */
    dstaddrs: pulumi.Input<pulumi.Input<inputs.FirewallLocalInPolicyDstaddr>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable dedicating the HA management interface only for local-in policy. Valid values: `enable`, `disable`.
     */
    haMgmtIntfOnly?: pulumi.Input<string>;
    /**
     * Incoming interface name from available options.
     */
    intf?: pulumi.Input<string>;
    /**
     * User defined local in policy ID.
     */
    policyid?: pulumi.Input<number>;
    /**
     * Schedule object from available options.
     */
    schedule: pulumi.Input<string>;
    /**
     * Service object from available options. The structure of `service` block is documented below.
     */
    services?: pulumi.Input<pulumi.Input<inputs.FirewallLocalInPolicyService>[]>;
    /**
     * Source address object from available options. The structure of `srcaddr` block is documented below.
     */
    srcaddrs: pulumi.Input<pulumi.Input<inputs.FirewallLocalInPolicySrcaddr>[]>;
    /**
     * Enable/disable this local-in policy. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}