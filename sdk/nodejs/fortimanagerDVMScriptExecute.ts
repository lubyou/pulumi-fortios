// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource supports executing devicemanager script on Fortimanager.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const test3 = new fortios.FortimanagerDVMScriptExecute("test3", {
 *     scriptName: "config-intf3",
 *     targetDevname: "devname",
 *     timeout: 5,
 * });
 * ```
 */
export class FortimanagerDVMScriptExecute extends pulumi.CustomResource {
    /**
     * Get an existing FortimanagerDVMScriptExecute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FortimanagerDVMScriptExecuteState, opts?: pulumi.CustomResourceOptions): FortimanagerDVMScriptExecute {
        return new FortimanagerDVMScriptExecute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/fortimanagerDVMScriptExecute:FortimanagerDVMScriptExecute';

    /**
     * Returns true if the given object is an instance of FortimanagerDVMScriptExecute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FortimanagerDVMScriptExecute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FortimanagerDVMScriptExecute.__pulumiType;
    }

    /**
     * Source ADOM name. default is 'root'
     */
    public readonly adom!: pulumi.Output<string | undefined>;
    /**
     * Policy package.
     */
    public readonly package!: pulumi.Output<string | undefined>;
    /**
     * Script name.
     */
    public readonly scriptName!: pulumi.Output<string>;
    /**
     * Target device name, which the script will be installed.
     */
    public readonly targetDevname!: pulumi.Output<string | undefined>;
    /**
     * Timeout(minute) for executing the script, default is 3 minutes.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;
    /**
     * Vdom of managed device. default is 'root'
     */
    public readonly vdom!: pulumi.Output<string | undefined>;

    /**
     * Create a FortimanagerDVMScriptExecute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FortimanagerDVMScriptExecuteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FortimanagerDVMScriptExecuteArgs | FortimanagerDVMScriptExecuteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FortimanagerDVMScriptExecuteState | undefined;
            resourceInputs["adom"] = state ? state.adom : undefined;
            resourceInputs["package"] = state ? state.package : undefined;
            resourceInputs["scriptName"] = state ? state.scriptName : undefined;
            resourceInputs["targetDevname"] = state ? state.targetDevname : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["vdom"] = state ? state.vdom : undefined;
        } else {
            const args = argsOrState as FortimanagerDVMScriptExecuteArgs | undefined;
            if ((!args || args.scriptName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scriptName'");
            }
            resourceInputs["adom"] = args ? args.adom : undefined;
            resourceInputs["package"] = args ? args.package : undefined;
            resourceInputs["scriptName"] = args ? args.scriptName : undefined;
            resourceInputs["targetDevname"] = args ? args.targetDevname : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["vdom"] = args ? args.vdom : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FortimanagerDVMScriptExecute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FortimanagerDVMScriptExecute resources.
 */
export interface FortimanagerDVMScriptExecuteState {
    /**
     * Source ADOM name. default is 'root'
     */
    adom?: pulumi.Input<string>;
    /**
     * Policy package.
     */
    package?: pulumi.Input<string>;
    /**
     * Script name.
     */
    scriptName?: pulumi.Input<string>;
    /**
     * Target device name, which the script will be installed.
     */
    targetDevname?: pulumi.Input<string>;
    /**
     * Timeout(minute) for executing the script, default is 3 minutes.
     */
    timeout?: pulumi.Input<number>;
    /**
     * Vdom of managed device. default is 'root'
     */
    vdom?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FortimanagerDVMScriptExecute resource.
 */
export interface FortimanagerDVMScriptExecuteArgs {
    /**
     * Source ADOM name. default is 'root'
     */
    adom?: pulumi.Input<string>;
    /**
     * Policy package.
     */
    package?: pulumi.Input<string>;
    /**
     * Script name.
     */
    scriptName: pulumi.Input<string>;
    /**
     * Target device name, which the script will be installed.
     */
    targetDevname?: pulumi.Input<string>;
    /**
     * Timeout(minute) for executing the script, default is 3 minutes.
     */
    timeout?: pulumi.Input<number>;
    /**
     * Vdom of managed device. default is 'root'
     */
    vdom?: pulumi.Input<string>;
}
