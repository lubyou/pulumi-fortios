// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Web application firewall configuration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.WafProfile("trname", {
 *     extendedLog: "disable",
 *     external: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * Waf Profile can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/wafProfile:WafProfile labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/wafProfile:WafProfile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WafProfile extends pulumi.CustomResource {
    /**
     * Get an existing WafProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WafProfileState, opts?: pulumi.CustomResourceOptions): WafProfile {
        return new WafProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wafProfile:WafProfile';

    /**
     * Returns true if the given object is an instance of WafProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WafProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WafProfile.__pulumiType;
    }

    /**
     * Black address list and white address list. The structure of `addressList` block is documented below.
     */
    public readonly addressList!: pulumi.Output<outputs.WafProfileAddressList | undefined>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * WAF HTTP protocol restrictions. The structure of `constraint` block is documented below.
     */
    public readonly constraint!: pulumi.Output<outputs.WafProfileConstraint | undefined>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable extended logging. Valid values: `enable`, `disable`.
     */
    public readonly extendedLog!: pulumi.Output<string>;
    /**
     * Disable/Enable external HTTP Inspection. Valid values: `disable`, `enable`.
     */
    public readonly external!: pulumi.Output<string>;
    /**
     * Enable/disable HTTP method check. Valid values: `enable`, `disable`.
     */
    public readonly method!: pulumi.Output<outputs.WafProfileMethod | undefined>;
    /**
     * Address name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * WAF signatures. The structure of `signature` block is documented below.
     */
    public readonly signature!: pulumi.Output<outputs.WafProfileSignature | undefined>;
    /**
     * URL access list The structure of `urlAccess` block is documented below.
     */
    public readonly urlAccesses!: pulumi.Output<outputs.WafProfileUrlAccess[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WafProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WafProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WafProfileArgs | WafProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WafProfileState | undefined;
            resourceInputs["addressList"] = state ? state.addressList : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["constraint"] = state ? state.constraint : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["extendedLog"] = state ? state.extendedLog : undefined;
            resourceInputs["external"] = state ? state.external : undefined;
            resourceInputs["method"] = state ? state.method : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["signature"] = state ? state.signature : undefined;
            resourceInputs["urlAccesses"] = state ? state.urlAccesses : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WafProfileArgs | undefined;
            resourceInputs["addressList"] = args ? args.addressList : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["constraint"] = args ? args.constraint : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["extendedLog"] = args ? args.extendedLog : undefined;
            resourceInputs["external"] = args ? args.external : undefined;
            resourceInputs["method"] = args ? args.method : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["signature"] = args ? args.signature : undefined;
            resourceInputs["urlAccesses"] = args ? args.urlAccesses : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WafProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WafProfile resources.
 */
export interface WafProfileState {
    /**
     * Black address list and white address list. The structure of `addressList` block is documented below.
     */
    addressList?: pulumi.Input<inputs.WafProfileAddressList>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * WAF HTTP protocol restrictions. The structure of `constraint` block is documented below.
     */
    constraint?: pulumi.Input<inputs.WafProfileConstraint>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable extended logging. Valid values: `enable`, `disable`.
     */
    extendedLog?: pulumi.Input<string>;
    /**
     * Disable/Enable external HTTP Inspection. Valid values: `disable`, `enable`.
     */
    external?: pulumi.Input<string>;
    /**
     * Enable/disable HTTP method check. Valid values: `enable`, `disable`.
     */
    method?: pulumi.Input<inputs.WafProfileMethod>;
    /**
     * Address name.
     */
    name?: pulumi.Input<string>;
    /**
     * WAF signatures. The structure of `signature` block is documented below.
     */
    signature?: pulumi.Input<inputs.WafProfileSignature>;
    /**
     * URL access list The structure of `urlAccess` block is documented below.
     */
    urlAccesses?: pulumi.Input<pulumi.Input<inputs.WafProfileUrlAccess>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WafProfile resource.
 */
export interface WafProfileArgs {
    /**
     * Black address list and white address list. The structure of `addressList` block is documented below.
     */
    addressList?: pulumi.Input<inputs.WafProfileAddressList>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * WAF HTTP protocol restrictions. The structure of `constraint` block is documented below.
     */
    constraint?: pulumi.Input<inputs.WafProfileConstraint>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable extended logging. Valid values: `enable`, `disable`.
     */
    extendedLog?: pulumi.Input<string>;
    /**
     * Disable/Enable external HTTP Inspection. Valid values: `disable`, `enable`.
     */
    external?: pulumi.Input<string>;
    /**
     * Enable/disable HTTP method check. Valid values: `enable`, `disable`.
     */
    method?: pulumi.Input<inputs.WafProfileMethod>;
    /**
     * Address name.
     */
    name?: pulumi.Input<string>;
    /**
     * WAF signatures. The structure of `signature` block is documented below.
     */
    signature?: pulumi.Input<inputs.WafProfileSignature>;
    /**
     * URL access list The structure of `urlAccess` block is documented below.
     */
    urlAccesses?: pulumi.Input<pulumi.Input<inputs.WafProfileUrlAccess>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
