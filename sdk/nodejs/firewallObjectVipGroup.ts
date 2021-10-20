// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a resource to configure virtual IP groups of FortiOS.
 *
 * !> **Warning:** The resource will be deprecated and replaced by new resource `fortios.FirewallVipgrp`, we recommend that you use the new resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const v11 = new fortios.FirewallObjectVipGroup("v11", {
 *     comments: "comments",
 *     interface: "port3",
 *     members: [
 *         "vip1",
 *         "vip3",
 *     ],
 * });
 * ```
 */
export class FirewallObjectVipGroup extends pulumi.CustomResource {
    /**
     * Get an existing FirewallObjectVipGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallObjectVipGroupState, opts?: pulumi.CustomResourceOptions): FirewallObjectVipGroup {
        return new FirewallObjectVipGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallObjectVipGroup:FirewallObjectVipGroup';

    /**
     * Returns true if the given object is an instance of FirewallObjectVipGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallObjectVipGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallObjectVipGroup.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Interface name.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Member VIP objects of the group.
     */
    public readonly members!: pulumi.Output<string[]>;
    /**
     * VIP group name.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a FirewallObjectVipGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallObjectVipGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallObjectVipGroupArgs | FirewallObjectVipGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallObjectVipGroupState | undefined;
            inputs["comments"] = state ? state.comments : undefined;
            inputs["interface"] = state ? state.interface : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as FirewallObjectVipGroupArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            inputs["comments"] = args ? args.comments : undefined;
            inputs["interface"] = args ? args.interface : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FirewallObjectVipGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallObjectVipGroup resources.
 */
export interface FirewallObjectVipGroupState {
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Interface name.
     */
    interface?: pulumi.Input<string>;
    /**
     * Member VIP objects of the group.
     */
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * VIP group name.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallObjectVipGroup resource.
 */
export interface FirewallObjectVipGroupArgs {
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Interface name.
     */
    interface?: pulumi.Input<string>;
    /**
     * Member VIP objects of the group.
     */
    members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * VIP group name.
     */
    name?: pulumi.Input<string>;
}