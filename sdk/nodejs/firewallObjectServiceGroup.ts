// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a resource to configure firewall service group of FortiOS.
 *
 * !> **Warning:** The resource will be deprecated and replaced by new resource `fortios.FirewallServiceGroup`, we recommend that you use the new resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const v11 = new fortios.FirewallObjectServiceGroup("v11", {
 *     comment: "fdsafdsa",
 *     members: [
 *         "DCE-RPC",
 *         "DNS",
 *         "HTTPS",
 *     ],
 * });
 * ```
 */
export class FirewallObjectServiceGroup extends pulumi.CustomResource {
    /**
     * Get an existing FirewallObjectServiceGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallObjectServiceGroupState, opts?: pulumi.CustomResourceOptions): FirewallObjectServiceGroup {
        return new FirewallObjectServiceGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallObjectServiceGroup:FirewallObjectServiceGroup';

    /**
     * Returns true if the given object is an instance of FirewallObjectServiceGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallObjectServiceGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallObjectServiceGroup.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Service objects contained within the group.
     */
    public readonly members!: pulumi.Output<string[]>;
    /**
     * Service group name.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a FirewallObjectServiceGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallObjectServiceGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallObjectServiceGroupArgs | FirewallObjectServiceGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallObjectServiceGroupState | undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as FirewallObjectServiceGroupArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            inputs["comment"] = args ? args.comment : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FirewallObjectServiceGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallObjectServiceGroup resources.
 */
export interface FirewallObjectServiceGroupState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Service objects contained within the group.
     */
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Service group name.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallObjectServiceGroup resource.
 */
export interface FirewallObjectServiceGroupArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Service objects contained within the group.
     */
    members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Service group name.
     */
    name?: pulumi.Input<string>;
}