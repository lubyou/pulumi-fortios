// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure firewall authentication portals.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.FirewallAuthPortal("trname", {
 *     groups: [{
 *         name: "Guest-group",
 *     }],
 *     portalAddr: "1.1.1.1",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall AuthPortal can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallAuthPortal:FirewallAuthPortal labelname FirewallAuthPortal
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
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

    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Firewall user groups permitted to authenticate through this portal. Separate group names with spaces. The structure of `groups` block is documented below.
     */
    public readonly groups!: pulumi.Output<outputs.FirewallAuthPortalGroup[] | undefined>;
    /**
     * Name of the identity-based route that applies to this portal.
     */
    public readonly identityBasedRoute!: pulumi.Output<string>;
    /**
     * Address (or FQDN) of the authentication portal.
     */
    public readonly portalAddr!: pulumi.Output<string>;
    /**
     * IPv6 address (or FQDN) of authentication portal.
     */
    public readonly portalAddr6!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
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
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["identityBasedRoute"] = state ? state.identityBasedRoute : undefined;
            resourceInputs["portalAddr"] = state ? state.portalAddr : undefined;
            resourceInputs["portalAddr6"] = state ? state.portalAddr6 : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallAuthPortalArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["identityBasedRoute"] = args ? args.identityBasedRoute : undefined;
            resourceInputs["portalAddr"] = args ? args.portalAddr : undefined;
            resourceInputs["portalAddr6"] = args ? args.portalAddr6 : undefined;
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
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Firewall user groups permitted to authenticate through this portal. Separate group names with spaces. The structure of `groups` block is documented below.
     */
    groups?: pulumi.Input<pulumi.Input<inputs.FirewallAuthPortalGroup>[]>;
    /**
     * Name of the identity-based route that applies to this portal.
     */
    identityBasedRoute?: pulumi.Input<string>;
    /**
     * Address (or FQDN) of the authentication portal.
     */
    portalAddr?: pulumi.Input<string>;
    /**
     * IPv6 address (or FQDN) of authentication portal.
     */
    portalAddr6?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallAuthPortal resource.
 */
export interface FirewallAuthPortalArgs {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Firewall user groups permitted to authenticate through this portal. Separate group names with spaces. The structure of `groups` block is documented below.
     */
    groups?: pulumi.Input<pulumi.Input<inputs.FirewallAuthPortalGroup>[]>;
    /**
     * Name of the identity-based route that applies to this portal.
     */
    identityBasedRoute?: pulumi.Input<string>;
    /**
     * Address (or FQDN) of the authentication portal.
     */
    portalAddr?: pulumi.Input<string>;
    /**
     * IPv6 address (or FQDN) of authentication portal.
     */
    portalAddr6?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
