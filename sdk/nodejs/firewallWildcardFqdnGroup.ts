// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Config global Wildcard FQDN address groups.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi_fortios from "@lubyou/pulumi-fortios";
 *
 * const trname1 = new fortios.FirewallWildcardFqdnCustom("trname1", {
 *     color: 0,
 *     visibility: "enable",
 *     wildcardFqdn: "*.ms.com",
 * });
 * const trname = new fortios.FirewallWildcardFqdnGroup("trname", {
 *     color: 0,
 *     visibility: "enable",
 *     members: [{
 *         name: trname1.name,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * FirewallWildcardFqdn Group can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallWildcardFqdnGroup:FirewallWildcardFqdnGroup labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallWildcardFqdnGroup:FirewallWildcardFqdnGroup labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallWildcardFqdnGroup extends pulumi.CustomResource {
    /**
     * Get an existing FirewallWildcardFqdnGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallWildcardFqdnGroupState, opts?: pulumi.CustomResourceOptions): FirewallWildcardFqdnGroup {
        return new FirewallWildcardFqdnGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallWildcardFqdnGroup:FirewallWildcardFqdnGroup';

    /**
     * Returns true if the given object is an instance of FirewallWildcardFqdnGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallWildcardFqdnGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallWildcardFqdnGroup.__pulumiType;
    }

    /**
     * GUI icon color.
     */
    public readonly color!: pulumi.Output<number>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Address group members. The structure of `member` block is documented below.
     */
    public readonly members!: pulumi.Output<outputs.FirewallWildcardFqdnGroupMember[]>;
    /**
     * Address name.
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
     * Enable/disable address visibility. Valid values: `enable`, `disable`.
     */
    public readonly visibility!: pulumi.Output<string>;

    /**
     * Create a FirewallWildcardFqdnGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallWildcardFqdnGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallWildcardFqdnGroupArgs | FirewallWildcardFqdnGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallWildcardFqdnGroupState | undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as FirewallWildcardFqdnGroupArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallWildcardFqdnGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallWildcardFqdnGroup resources.
 */
export interface FirewallWildcardFqdnGroupState {
    /**
     * GUI icon color.
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Address group members. The structure of `member` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.FirewallWildcardFqdnGroupMember>[]>;
    /**
     * Address name.
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
    /**
     * Enable/disable address visibility. Valid values: `enable`, `disable`.
     */
    visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallWildcardFqdnGroup resource.
 */
export interface FirewallWildcardFqdnGroupArgs {
    /**
     * GUI icon color.
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Address group members. The structure of `member` block is documented below.
     */
    members: pulumi.Input<pulumi.Input<inputs.FirewallWildcardFqdnGroupMember>[]>;
    /**
     * Address name.
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
    /**
     * Enable/disable address visibility. Valid values: `enable`, `disable`.
     */
    visibility?: pulumi.Input<string>;
}
