// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource supports Create/Read/Update/Delete firewall security policypackage on FortiManager which could be installed to the FortiGate later
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const test1 = new fortios.FortimanagerFirewallSecurityPolicyPackage("test1", {
 *     target: "FGVM64-test",
 * });
 * ```
 */
export class FortimanagerFirewallSecurityPolicyPackage extends pulumi.CustomResource {
    /**
     * Get an existing FortimanagerFirewallSecurityPolicyPackage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FortimanagerFirewallSecurityPolicyPackageState, opts?: pulumi.CustomResourceOptions): FortimanagerFirewallSecurityPolicyPackage {
        return new FortimanagerFirewallSecurityPolicyPackage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/fortimanagerFirewallSecurityPolicyPackage:FortimanagerFirewallSecurityPolicyPackage';

    /**
     * Returns true if the given object is an instance of FortimanagerFirewallSecurityPolicyPackage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FortimanagerFirewallSecurityPolicyPackage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FortimanagerFirewallSecurityPolicyPackage.__pulumiType;
    }

    /**
     * Source ADOM name. default is 'root'
     */
    public readonly adom!: pulumi.Output<string | undefined>;
    /**
     * Inspection Mode. Enum:[flow, proxy]. default is 'flow'
     */
    public readonly inspectionMode!: pulumi.Output<string | undefined>;
    /**
     * Security policy package name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The installation target.
     */
    public readonly target!: pulumi.Output<string | undefined>;
    /**
     * Vdom of managed device. default is 'root'
     */
    public readonly vdom!: pulumi.Output<string | undefined>;

    /**
     * Create a FortimanagerFirewallSecurityPolicyPackage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FortimanagerFirewallSecurityPolicyPackageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FortimanagerFirewallSecurityPolicyPackageArgs | FortimanagerFirewallSecurityPolicyPackageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FortimanagerFirewallSecurityPolicyPackageState | undefined;
            inputs["adom"] = state ? state.adom : undefined;
            inputs["inspectionMode"] = state ? state.inspectionMode : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["target"] = state ? state.target : undefined;
            inputs["vdom"] = state ? state.vdom : undefined;
        } else {
            const args = argsOrState as FortimanagerFirewallSecurityPolicyPackageArgs | undefined;
            inputs["adom"] = args ? args.adom : undefined;
            inputs["inspectionMode"] = args ? args.inspectionMode : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["vdom"] = args ? args.vdom : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FortimanagerFirewallSecurityPolicyPackage.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FortimanagerFirewallSecurityPolicyPackage resources.
 */
export interface FortimanagerFirewallSecurityPolicyPackageState {
    /**
     * Source ADOM name. default is 'root'
     */
    adom?: pulumi.Input<string>;
    /**
     * Inspection Mode. Enum:[flow, proxy]. default is 'flow'
     */
    inspectionMode?: pulumi.Input<string>;
    /**
     * Security policy package name.
     */
    name?: pulumi.Input<string>;
    /**
     * The installation target.
     */
    target?: pulumi.Input<string>;
    /**
     * Vdom of managed device. default is 'root'
     */
    vdom?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FortimanagerFirewallSecurityPolicyPackage resource.
 */
export interface FortimanagerFirewallSecurityPolicyPackageArgs {
    /**
     * Source ADOM name. default is 'root'
     */
    adom?: pulumi.Input<string>;
    /**
     * Inspection Mode. Enum:[flow, proxy]. default is 'flow'
     */
    inspectionMode?: pulumi.Input<string>;
    /**
     * Security policy package name.
     */
    name?: pulumi.Input<string>;
    /**
     * The installation target.
     */
    target?: pulumi.Input<string>;
    /**
     * Vdom of managed device. default is 'root'
     */
    vdom?: pulumi.Input<string>;
}