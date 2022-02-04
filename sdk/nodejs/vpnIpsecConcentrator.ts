// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Concentrator configuration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.VpnIpsecConcentrator("trname", {
 *     srcCheck: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * VpnIpsec Concentrator can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/vpnIpsecConcentrator:VpnIpsecConcentrator labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class VpnIpsecConcentrator extends pulumi.CustomResource {
    /**
     * Get an existing VpnIpsecConcentrator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpnIpsecConcentratorState, opts?: pulumi.CustomResourceOptions): VpnIpsecConcentrator {
        return new VpnIpsecConcentrator(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/vpnIpsecConcentrator:VpnIpsecConcentrator';

    /**
     * Returns true if the given object is an instance of VpnIpsecConcentrator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnIpsecConcentrator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnIpsecConcentrator.__pulumiType;
    }

    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Concentrator ID. (1-65535)
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Names of up to 3 VPN tunnels to add to the concentrator. The structure of `member` block is documented below.
     */
    public readonly members!: pulumi.Output<outputs.VpnIpsecConcentratorMember[] | undefined>;
    /**
     * Member name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable to check source address of phase 2 selector. Disable to check only the destination selector. Valid values: `disable`, `enable`.
     */
    public readonly srcCheck!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a VpnIpsecConcentrator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VpnIpsecConcentratorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnIpsecConcentratorArgs | VpnIpsecConcentratorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpnIpsecConcentratorState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["srcCheck"] = state ? state.srcCheck : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as VpnIpsecConcentratorArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["srcCheck"] = args ? args.srcCheck : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpnIpsecConcentrator.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpnIpsecConcentrator resources.
 */
export interface VpnIpsecConcentratorState {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Concentrator ID. (1-65535)
     */
    fosid?: pulumi.Input<number>;
    /**
     * Names of up to 3 VPN tunnels to add to the concentrator. The structure of `member` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.VpnIpsecConcentratorMember>[]>;
    /**
     * Member name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable to check source address of phase 2 selector. Disable to check only the destination selector. Valid values: `disable`, `enable`.
     */
    srcCheck?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpnIpsecConcentrator resource.
 */
export interface VpnIpsecConcentratorArgs {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Concentrator ID. (1-65535)
     */
    fosid?: pulumi.Input<number>;
    /**
     * Names of up to 3 VPN tunnels to add to the concentrator. The structure of `member` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.VpnIpsecConcentratorMember>[]>;
    /**
     * Member name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable to check source address of phase 2 selector. Disable to check only the destination selector. Valid values: `disable`, `enable`.
     */
    srcCheck?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
