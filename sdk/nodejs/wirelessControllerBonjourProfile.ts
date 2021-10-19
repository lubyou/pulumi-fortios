// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure Bonjour profiles. Bonjour is Apple's zero configuration networking protocol. Bonjour profiles allow APs and FortiAPs to connnect to networks using Bonjour.
 *
 * ## Import
 *
 * WirelessController BonjourProfile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelessControllerBonjourProfile:WirelessControllerBonjourProfile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WirelessControllerBonjourProfile extends pulumi.CustomResource {
    /**
     * Get an existing WirelessControllerBonjourProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WirelessControllerBonjourProfileState, opts?: pulumi.CustomResourceOptions): WirelessControllerBonjourProfile {
        return new WirelessControllerBonjourProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelessControllerBonjourProfile:WirelessControllerBonjourProfile';

    /**
     * Returns true if the given object is an instance of WirelessControllerBonjourProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WirelessControllerBonjourProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WirelessControllerBonjourProfile.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Bonjour profile name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Bonjour policy list. The structure of `policyList` block is documented below.
     */
    public readonly policyLists!: pulumi.Output<outputs.WirelessControllerBonjourProfilePolicyList[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WirelessControllerBonjourProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WirelessControllerBonjourProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WirelessControllerBonjourProfileArgs | WirelessControllerBonjourProfileState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WirelessControllerBonjourProfileState | undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["policyLists"] = state ? state.policyLists : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WirelessControllerBonjourProfileArgs | undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policyLists"] = args ? args.policyLists : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(WirelessControllerBonjourProfile.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WirelessControllerBonjourProfile resources.
 */
export interface WirelessControllerBonjourProfileState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Bonjour profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Bonjour policy list. The structure of `policyList` block is documented below.
     */
    policyLists?: pulumi.Input<pulumi.Input<inputs.WirelessControllerBonjourProfilePolicyList>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WirelessControllerBonjourProfile resource.
 */
export interface WirelessControllerBonjourProfileArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Bonjour profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Bonjour policy list. The structure of `policyList` block is documented below.
     */
    policyLists?: pulumi.Input<pulumi.Input<inputs.WirelessControllerBonjourProfilePolicyList>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
