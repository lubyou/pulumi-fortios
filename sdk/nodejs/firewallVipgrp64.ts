// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure IPv6 to IPv4 virtual IP groups. Applies to FortiOS Version `<= 7.0.0`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi_fortios from "@lubyou/pulumi-fortios";
 *
 * const trname1 = new fortios.FirewallVip64("trname1", {
 *     arpReply: "enable",
 *     color: 0,
 *     extip: "2001:db8:99:503::22",
 *     extport: "0-65535",
 *     fosid: 0,
 *     ldbMethod: "static",
 *     mappedip: "1.1.3.1",
 *     mappedport: "0-65535",
 *     portforward: "disable",
 *     protocol: "tcp",
 *     type: "static-nat",
 * });
 * const trname = new fortios.FirewallVipgrp64("trname", {
 *     color: 0,
 *     members: [{
 *         name: trname1.name,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Firewall Vipgrp64 can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallVipgrp64:FirewallVipgrp64 labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallVipgrp64:FirewallVipgrp64 labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallVipgrp64 extends pulumi.CustomResource {
    /**
     * Get an existing FirewallVipgrp64 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallVipgrp64State, opts?: pulumi.CustomResourceOptions): FirewallVipgrp64 {
        return new FirewallVipgrp64(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallVipgrp64:FirewallVipgrp64';

    /**
     * Returns true if the given object is an instance of FirewallVipgrp64.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallVipgrp64 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallVipgrp64.__pulumiType;
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
     * Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
     */
    public readonly members!: pulumi.Output<outputs.FirewallVipgrp64Member[]>;
    /**
     * VIP64 name.
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
     * Create a FirewallVipgrp64 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallVipgrp64Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallVipgrp64Args | FirewallVipgrp64State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallVipgrp64State | undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallVipgrp64Args | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallVipgrp64.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallVipgrp64 resources.
 */
export interface FirewallVipgrp64State {
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
     * Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.FirewallVipgrp64Member>[]>;
    /**
     * VIP64 name.
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
 * The set of arguments for constructing a FirewallVipgrp64 resource.
 */
export interface FirewallVipgrp64Args {
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
     * Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.
     */
    members: pulumi.Input<pulumi.Input<inputs.FirewallVipgrp64Member>[]>;
    /**
     * VIP64 name.
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
