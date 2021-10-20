// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure VLAN policy to be applied on the managed FortiSwitch ports through port-policy. Applies to FortiOS Version `>= 6.4.0`.
 *
 * ## Import
 *
 * SwitchController VlanPolicy can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerVlanPolicy:SwitchControllerVlanPolicy labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
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

    /**
     * Allowed VLANs to be applied when using this VLAN policy. The structure of `allowedVlans` block is documented below.
     */
    public readonly allowedVlans!: pulumi.Output<outputs.SwitchControllerVlanPolicyAllowedVlan[] | undefined>;
    /**
     * Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
     */
    public readonly allowedVlansAll!: pulumi.Output<string>;
    /**
     * Description for the VLAN policy.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
     */
    public readonly discardMode!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * FortiLink interface for which this VLAN policy belongs to.
     */
    public readonly fortilink!: pulumi.Output<string>;
    /**
     * VLAN policy name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Untagged VLANs to be applied when using this VLAN policy. The structure of `untaggedVlans` block is documented below.
     */
    public readonly untaggedVlans!: pulumi.Output<outputs.SwitchControllerVlanPolicyUntaggedVlan[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Native VLAN to be applied when using this VLAN policy.
     */
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerVlanPolicyState | undefined;
            inputs["allowedVlans"] = state ? state.allowedVlans : undefined;
            inputs["allowedVlansAll"] = state ? state.allowedVlansAll : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["discardMode"] = state ? state.discardMode : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["fortilink"] = state ? state.fortilink : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["untaggedVlans"] = state ? state.untaggedVlans : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["vlan"] = state ? state.vlan : undefined;
        } else {
            const args = argsOrState as SwitchControllerVlanPolicyArgs | undefined;
            inputs["allowedVlans"] = args ? args.allowedVlans : undefined;
            inputs["allowedVlansAll"] = args ? args.allowedVlansAll : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["discardMode"] = args ? args.discardMode : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["fortilink"] = args ? args.fortilink : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["untaggedVlans"] = args ? args.untaggedVlans : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["vlan"] = args ? args.vlan : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SwitchControllerVlanPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerVlanPolicy resources.
 */
export interface SwitchControllerVlanPolicyState {
    /**
     * Allowed VLANs to be applied when using this VLAN policy. The structure of `allowedVlans` block is documented below.
     */
    allowedVlans?: pulumi.Input<pulumi.Input<inputs.SwitchControllerVlanPolicyAllowedVlan>[]>;
    /**
     * Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
     */
    allowedVlansAll?: pulumi.Input<string>;
    /**
     * Description for the VLAN policy.
     */
    description?: pulumi.Input<string>;
    /**
     * Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
     */
    discardMode?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * FortiLink interface for which this VLAN policy belongs to.
     */
    fortilink?: pulumi.Input<string>;
    /**
     * VLAN policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Untagged VLANs to be applied when using this VLAN policy. The structure of `untaggedVlans` block is documented below.
     */
    untaggedVlans?: pulumi.Input<pulumi.Input<inputs.SwitchControllerVlanPolicyUntaggedVlan>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Native VLAN to be applied when using this VLAN policy.
     */
    vlan?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerVlanPolicy resource.
 */
export interface SwitchControllerVlanPolicyArgs {
    /**
     * Allowed VLANs to be applied when using this VLAN policy. The structure of `allowedVlans` block is documented below.
     */
    allowedVlans?: pulumi.Input<pulumi.Input<inputs.SwitchControllerVlanPolicyAllowedVlan>[]>;
    /**
     * Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
     */
    allowedVlansAll?: pulumi.Input<string>;
    /**
     * Description for the VLAN policy.
     */
    description?: pulumi.Input<string>;
    /**
     * Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
     */
    discardMode?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * FortiLink interface for which this VLAN policy belongs to.
     */
    fortilink?: pulumi.Input<string>;
    /**
     * VLAN policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Untagged VLANs to be applied when using this VLAN policy. The structure of `untaggedVlans` block is documented below.
     */
    untaggedVlans?: pulumi.Input<pulumi.Input<inputs.SwitchControllerVlanPolicyUntaggedVlan>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Native VLAN to be applied when using this VLAN policy.
     */
    vlan?: pulumi.Input<string>;
}