// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class FirewallAddrgrp extends pulumi.CustomResource {
    /**
     * Get an existing FirewallAddrgrp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallAddrgrpState, opts?: pulumi.CustomResourceOptions): FirewallAddrgrp {
        return new FirewallAddrgrp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallAddrgrp:FirewallAddrgrp';

    /**
     * Returns true if the given object is an instance of FirewallAddrgrp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallAddrgrp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallAddrgrp.__pulumiType;
    }

    public readonly allowRouting!: pulumi.Output<string>;
    public readonly category!: pulumi.Output<string>;
    public readonly color!: pulumi.Output<number>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly exclude!: pulumi.Output<string>;
    public readonly excludeMembers!: pulumi.Output<outputs.FirewallAddrgrpExcludeMember[] | undefined>;
    public readonly fabricObject!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<outputs.FirewallAddrgrpMember[]>;
    public readonly name!: pulumi.Output<string>;
    public readonly taggings!: pulumi.Output<outputs.FirewallAddrgrpTagging[] | undefined>;
    public readonly type!: pulumi.Output<string>;
    public readonly uuid!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly visibility!: pulumi.Output<string>;

    /**
     * Create a FirewallAddrgrp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallAddrgrpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallAddrgrpArgs | FirewallAddrgrpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallAddrgrpState | undefined;
            resourceInputs["allowRouting"] = state ? state.allowRouting : undefined;
            resourceInputs["category"] = state ? state.category : undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["exclude"] = state ? state.exclude : undefined;
            resourceInputs["excludeMembers"] = state ? state.excludeMembers : undefined;
            resourceInputs["fabricObject"] = state ? state.fabricObject : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["taggings"] = state ? state.taggings : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as FirewallAddrgrpArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            resourceInputs["allowRouting"] = args ? args.allowRouting : undefined;
            resourceInputs["category"] = args ? args.category : undefined;
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["exclude"] = args ? args.exclude : undefined;
            resourceInputs["excludeMembers"] = args ? args.excludeMembers : undefined;
            resourceInputs["fabricObject"] = args ? args.fabricObject : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["taggings"] = args ? args.taggings : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallAddrgrp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallAddrgrp resources.
 */
export interface FirewallAddrgrpState {
    allowRouting?: pulumi.Input<string>;
    category?: pulumi.Input<string>;
    color?: pulumi.Input<number>;
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    exclude?: pulumi.Input<string>;
    excludeMembers?: pulumi.Input<pulumi.Input<inputs.FirewallAddrgrpExcludeMember>[]>;
    fabricObject?: pulumi.Input<string>;
    members?: pulumi.Input<pulumi.Input<inputs.FirewallAddrgrpMember>[]>;
    name?: pulumi.Input<string>;
    taggings?: pulumi.Input<pulumi.Input<inputs.FirewallAddrgrpTagging>[]>;
    type?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallAddrgrp resource.
 */
export interface FirewallAddrgrpArgs {
    allowRouting?: pulumi.Input<string>;
    category?: pulumi.Input<string>;
    color?: pulumi.Input<number>;
    comment?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    exclude?: pulumi.Input<string>;
    excludeMembers?: pulumi.Input<pulumi.Input<inputs.FirewallAddrgrpExcludeMember>[]>;
    fabricObject?: pulumi.Input<string>;
    members: pulumi.Input<pulumi.Input<inputs.FirewallAddrgrpMember>[]>;
    name?: pulumi.Input<string>;
    taggings?: pulumi.Input<pulumi.Input<inputs.FirewallAddrgrpTagging>[]>;
    type?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    visibility?: pulumi.Input<string>;
}
