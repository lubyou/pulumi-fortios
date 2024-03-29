// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class FirewallAuthPortal extends pulumi.CustomResource {
    /**
     * Get an existing FirewallAuthPortal resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallAuthPortalState, opts?: pulumi.CustomResourceOptions): FirewallAuthPortal {
        return new FirewallAuthPortal(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallAuthPortal:FirewallAuthPortal';

    /**
     * Returns true if the given object is an instance of FirewallAuthPortal.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallAuthPortal {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallAuthPortal.__pulumiType;
    }

    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly groups!: pulumi.Output<outputs.FirewallAuthPortalGroup[] | undefined>;
    public readonly identityBasedRoute!: pulumi.Output<string>;
    public readonly portalAddr!: pulumi.Output<string>;
    public readonly portalAddr6!: pulumi.Output<string>;
    public readonly proxyAuth!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallAuthPortal resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallAuthPortalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallAuthPortalArgs | FirewallAuthPortalState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallAuthPortalState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["identityBasedRoute"] = state ? state.identityBasedRoute : undefined;
            resourceInputs["portalAddr"] = state ? state.portalAddr : undefined;
            resourceInputs["portalAddr6"] = state ? state.portalAddr6 : undefined;
            resourceInputs["proxyAuth"] = state ? state.proxyAuth : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallAuthPortalArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["identityBasedRoute"] = args ? args.identityBasedRoute : undefined;
            resourceInputs["portalAddr"] = args ? args.portalAddr : undefined;
            resourceInputs["portalAddr6"] = args ? args.portalAddr6 : undefined;
            resourceInputs["proxyAuth"] = args ? args.proxyAuth : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallAuthPortal.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallAuthPortal resources.
 */
export interface FirewallAuthPortalState {
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    groups?: pulumi.Input<pulumi.Input<inputs.FirewallAuthPortalGroup>[]>;
    identityBasedRoute?: pulumi.Input<string>;
    portalAddr?: pulumi.Input<string>;
    portalAddr6?: pulumi.Input<string>;
    proxyAuth?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallAuthPortal resource.
 */
export interface FirewallAuthPortalArgs {
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    groups?: pulumi.Input<pulumi.Input<inputs.FirewallAuthPortalGroup>[]>;
    identityBasedRoute?: pulumi.Input<string>;
    portalAddr?: pulumi.Input<string>;
    portalAddr6?: pulumi.Input<string>;
    proxyAuth?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
