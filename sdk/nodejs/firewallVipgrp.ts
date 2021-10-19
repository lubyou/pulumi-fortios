// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure IPv4 virtual IP groups.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname1 = new fortios.FirewallVip("trname1", {
 *     extintf: "any",
 *     extport: "0-65535",
 *     extip: "2.0.0.1-2.0.0.4",
 *     mappedips: [{
 *         range: "3.0.0.0-3.0.0.3",
 *     }],
 * });
 * const trname = new fortios.FirewallVipgrp("trname", {
 *     color: 0,
 *     "interface": "any",
 *     members: [{
 *         name: trname1.name,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Firewall Vipgrp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallVipgrp:FirewallVipgrp labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallVipgrp extends pulumi.CustomResource {
    /**
     * Get an existing FirewallVipgrp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallVipgrpState, opts?: pulumi.CustomResourceOptions): FirewallVipgrp {
        return new FirewallVipgrp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallVipgrp:FirewallVipgrp';

    /**
     * Returns true if the given object is an instance of FirewallVipgrp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallVipgrp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallVipgrp.__pulumiType;
    }

    /**
     * Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
     */
    public readonly color!: pulumi.Output<number>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * interface
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
     */
    public readonly members!: pulumi.Output<outputs.FirewallVipgrpMember[]>;
    /**
     * VIP name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    public readonly uuid!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallVipgrp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallVipgrpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallVipgrpArgs | FirewallVipgrpState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallVipgrpState | undefined;
            inputs["color"] = state ? state.color : undefined;
            inputs["comments"] = state ? state.comments : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["interface"] = state ? state.interface : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["uuid"] = state ? state.uuid : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallVipgrpArgs | undefined;
            if ((!args || args.interface === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interface'");
            }
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            inputs["color"] = args ? args.color : undefined;
            inputs["comments"] = args ? args.comments : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["interface"] = args ? args.interface : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["uuid"] = args ? args.uuid : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FirewallVipgrp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallVipgrp resources.
 */
export interface FirewallVipgrpState {
    /**
     * Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * interface
     */
    interface?: pulumi.Input<string>;
    /**
     * Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.FirewallVipgrpMember>[]>;
    /**
     * VIP name.
     */
    name?: pulumi.Input<string>;
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
 * The set of arguments for constructing a FirewallVipgrp resource.
 */
export interface FirewallVipgrpArgs {
    /**
     * Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * interface
     */
    interface: pulumi.Input<string>;
    /**
     * Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
     */
    members: pulumi.Input<pulumi.Input<inputs.FirewallVipgrpMember>[]>;
    /**
     * VIP name.
     */
    name?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
