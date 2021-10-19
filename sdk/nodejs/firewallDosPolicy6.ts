// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure IPv6 DoS policies.
 *
 * ## Import
 *
 * Firewall DosPolicy6 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallDosPolicy6:FirewallDosPolicy6 labelname {{policyid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallDosPolicy6 extends pulumi.CustomResource {
    /**
     * Get an existing FirewallDosPolicy6 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallDosPolicy6State, opts?: pulumi.CustomResourceOptions): FirewallDosPolicy6 {
        return new FirewallDosPolicy6(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallDosPolicy6:FirewallDosPolicy6';

    /**
     * Returns true if the given object is an instance of FirewallDosPolicy6.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallDosPolicy6 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallDosPolicy6.__pulumiType;
    }

    /**
     * Anomaly name. The structure of `anomaly` block is documented below.
     */
    public readonly anomalies!: pulumi.Output<outputs.FirewallDosPolicy6Anomaly[] | undefined>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Destination address name from available addresses. The structure of `dstaddr` block is documented below.
     */
    public readonly dstaddrs!: pulumi.Output<outputs.FirewallDosPolicy6Dstaddr[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Incoming interface name from available interfaces.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Anomaly name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Policy ID.
     */
    public readonly policyid!: pulumi.Output<number>;
    /**
     * Service object from available options. The structure of `service` block is documented below.
     */
    public readonly services!: pulumi.Output<outputs.FirewallDosPolicy6Service[] | undefined>;
    /**
     * Source address name from available addresses. The structure of `srcaddr` block is documented below.
     */
    public readonly srcaddrs!: pulumi.Output<outputs.FirewallDosPolicy6Srcaddr[]>;
    /**
     * Enable/disable this anomaly. Valid values: `disable`, `enable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallDosPolicy6 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallDosPolicy6Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallDosPolicy6Args | FirewallDosPolicy6State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallDosPolicy6State | undefined;
            inputs["anomalies"] = state ? state.anomalies : undefined;
            inputs["comments"] = state ? state.comments : undefined;
            inputs["dstaddrs"] = state ? state.dstaddrs : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["interface"] = state ? state.interface : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["policyid"] = state ? state.policyid : undefined;
            inputs["services"] = state ? state.services : undefined;
            inputs["srcaddrs"] = state ? state.srcaddrs : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallDosPolicy6Args | undefined;
            if ((!args || args.dstaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstaddrs'");
            }
            if ((!args || args.interface === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interface'");
            }
            if ((!args || args.srcaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcaddrs'");
            }
            inputs["anomalies"] = args ? args.anomalies : undefined;
            inputs["comments"] = args ? args.comments : undefined;
            inputs["dstaddrs"] = args ? args.dstaddrs : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["interface"] = args ? args.interface : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policyid"] = args ? args.policyid : undefined;
            inputs["services"] = args ? args.services : undefined;
            inputs["srcaddrs"] = args ? args.srcaddrs : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FirewallDosPolicy6.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallDosPolicy6 resources.
 */
export interface FirewallDosPolicy6State {
    /**
     * Anomaly name. The structure of `anomaly` block is documented below.
     */
    anomalies?: pulumi.Input<pulumi.Input<inputs.FirewallDosPolicy6Anomaly>[]>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Destination address name from available addresses. The structure of `dstaddr` block is documented below.
     */
    dstaddrs?: pulumi.Input<pulumi.Input<inputs.FirewallDosPolicy6Dstaddr>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Incoming interface name from available interfaces.
     */
    interface?: pulumi.Input<string>;
    /**
     * Anomaly name.
     */
    name?: pulumi.Input<string>;
    /**
     * Policy ID.
     */
    policyid?: pulumi.Input<number>;
    /**
     * Service object from available options. The structure of `service` block is documented below.
     */
    services?: pulumi.Input<pulumi.Input<inputs.FirewallDosPolicy6Service>[]>;
    /**
     * Source address name from available addresses. The structure of `srcaddr` block is documented below.
     */
    srcaddrs?: pulumi.Input<pulumi.Input<inputs.FirewallDosPolicy6Srcaddr>[]>;
    /**
     * Enable/disable this anomaly. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallDosPolicy6 resource.
 */
export interface FirewallDosPolicy6Args {
    /**
     * Anomaly name. The structure of `anomaly` block is documented below.
     */
    anomalies?: pulumi.Input<pulumi.Input<inputs.FirewallDosPolicy6Anomaly>[]>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Destination address name from available addresses. The structure of `dstaddr` block is documented below.
     */
    dstaddrs: pulumi.Input<pulumi.Input<inputs.FirewallDosPolicy6Dstaddr>[]>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Incoming interface name from available interfaces.
     */
    interface: pulumi.Input<string>;
    /**
     * Anomaly name.
     */
    name?: pulumi.Input<string>;
    /**
     * Policy ID.
     */
    policyid?: pulumi.Input<number>;
    /**
     * Service object from available options. The structure of `service` block is documented below.
     */
    services?: pulumi.Input<pulumi.Input<inputs.FirewallDosPolicy6Service>[]>;
    /**
     * Source address name from available addresses. The structure of `srcaddr` block is documented below.
     */
    srcaddrs: pulumi.Input<pulumi.Input<inputs.FirewallDosPolicy6Srcaddr>[]>;
    /**
     * Enable/disable this anomaly. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
