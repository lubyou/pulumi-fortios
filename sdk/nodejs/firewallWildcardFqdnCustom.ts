// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class FirewallWildcardFqdnCustom extends pulumi.CustomResource {
    /**
     * Get an existing FirewallWildcardFqdnCustom resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallWildcardFqdnCustomState, opts?: pulumi.CustomResourceOptions): FirewallWildcardFqdnCustom {
        return new FirewallWildcardFqdnCustom(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallWildcardFqdnCustom:FirewallWildcardFqdnCustom';

    /**
     * Returns true if the given object is an instance of FirewallWildcardFqdnCustom.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallWildcardFqdnCustom {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallWildcardFqdnCustom.__pulumiType;
    }

    public readonly color!: pulumi.Output<number>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly uuid!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    public readonly visibility!: pulumi.Output<string>;
    public readonly wildcardFqdn!: pulumi.Output<string>;

    /**
     * Create a FirewallWildcardFqdnCustom resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallWildcardFqdnCustomArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallWildcardFqdnCustomArgs | FirewallWildcardFqdnCustomState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallWildcardFqdnCustomState | undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
            resourceInputs["wildcardFqdn"] = state ? state.wildcardFqdn : undefined;
        } else {
            const args = argsOrState as FirewallWildcardFqdnCustomArgs | undefined;
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
            resourceInputs["wildcardFqdn"] = args ? args.wildcardFqdn : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallWildcardFqdnCustom.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallWildcardFqdnCustom resources.
 */
export interface FirewallWildcardFqdnCustomState {
    color?: pulumi.Input<number>;
    comment?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    visibility?: pulumi.Input<string>;
    wildcardFqdn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallWildcardFqdnCustom resource.
 */
export interface FirewallWildcardFqdnCustomArgs {
    color?: pulumi.Input<number>;
    comment?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    uuid?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
    visibility?: pulumi.Input<string>;
    wildcardFqdn?: pulumi.Input<string>;
}
