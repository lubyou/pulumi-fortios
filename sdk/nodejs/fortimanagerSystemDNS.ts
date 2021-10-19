// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource supports modifying system dns setting for FortiManager.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const test1 = new fortios.FortimanagerSystemDNS("test1", {
 *     primary: "208.91.112.52",
 *     secondary: "208.91.112.54",
 * });
 * ```
 */
export class FortimanagerSystemDNS extends pulumi.CustomResource {
    /**
     * Get an existing FortimanagerSystemDNS resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FortimanagerSystemDNSState, opts?: pulumi.CustomResourceOptions): FortimanagerSystemDNS {
        return new FortimanagerSystemDNS(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/fortimanagerSystemDNS:FortimanagerSystemDNS';

    /**
     * Returns true if the given object is an instance of FortimanagerSystemDNS.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FortimanagerSystemDNS {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FortimanagerSystemDNS.__pulumiType;
    }

    /**
     * Primary DNS IP.
     */
    public readonly primary!: pulumi.Output<string | undefined>;
    /**
     * Secondary DNS IP.
     */
    public readonly secondary!: pulumi.Output<string | undefined>;

    /**
     * Create a FortimanagerSystemDNS resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FortimanagerSystemDNSArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FortimanagerSystemDNSArgs | FortimanagerSystemDNSState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FortimanagerSystemDNSState | undefined;
            inputs["primary"] = state ? state.primary : undefined;
            inputs["secondary"] = state ? state.secondary : undefined;
        } else {
            const args = argsOrState as FortimanagerSystemDNSArgs | undefined;
            inputs["primary"] = args ? args.primary : undefined;
            inputs["secondary"] = args ? args.secondary : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FortimanagerSystemDNS.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FortimanagerSystemDNS resources.
 */
export interface FortimanagerSystemDNSState {
    /**
     * Primary DNS IP.
     */
    primary?: pulumi.Input<string>;
    /**
     * Secondary DNS IP.
     */
    secondary?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FortimanagerSystemDNS resource.
 */
export interface FortimanagerSystemDNSArgs {
    /**
     * Primary DNS IP.
     */
    primary?: pulumi.Input<string>;
    /**
     * Secondary DNS IP.
     */
    secondary?: pulumi.Input<string>;
}
