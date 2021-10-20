// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource supports Create/Read/Update/Delete devicemanager script for FortiManager.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const test1 = new fortios.FortimanagerDVMScript("test1", {
 *     content: `config system interface 
 *  edit port3 
 * 	 set vdom "root"
 * 	 set ip 10.7.0.200 255.255.0.0 
 * 	 set allowaccess ping http https
 * 	 next 
 *  end`,
 *     description: "description",
 *     target: "remote_device",
 * });
 * ```
 */
export class FortimanagerDVMScript extends pulumi.CustomResource {
    /**
     * Get an existing FortimanagerDVMScript resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FortimanagerDVMScriptState, opts?: pulumi.CustomResourceOptions): FortimanagerDVMScript {
        return new FortimanagerDVMScript(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/fortimanagerDVMScript:FortimanagerDVMScript';

    /**
     * Returns true if the given object is an instance of FortimanagerDVMScript.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FortimanagerDVMScript {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FortimanagerDVMScript.__pulumiType;
    }

    /**
     * ADOM name. default is 'root'.
     */
    public readonly adom!: pulumi.Output<string | undefined>;
    /**
     * Script content, only cli script is supported now
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Script name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Script target, Enum: ["deviceDatabase", "remoteDevice", "adomDatabase"]
     */
    public readonly target!: pulumi.Output<string | undefined>;

    /**
     * Create a FortimanagerDVMScript resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FortimanagerDVMScriptArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FortimanagerDVMScriptArgs | FortimanagerDVMScriptState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FortimanagerDVMScriptState | undefined;
            inputs["adom"] = state ? state.adom : undefined;
            inputs["content"] = state ? state.content : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["target"] = state ? state.target : undefined;
        } else {
            const args = argsOrState as FortimanagerDVMScriptArgs | undefined;
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            inputs["adom"] = args ? args.adom : undefined;
            inputs["content"] = args ? args.content : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["target"] = args ? args.target : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FortimanagerDVMScript.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FortimanagerDVMScript resources.
 */
export interface FortimanagerDVMScriptState {
    /**
     * ADOM name. default is 'root'.
     */
    adom?: pulumi.Input<string>;
    /**
     * Script content, only cli script is supported now
     */
    content?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Script name.
     */
    name?: pulumi.Input<string>;
    /**
     * Script target, Enum: ["deviceDatabase", "remoteDevice", "adomDatabase"]
     */
    target?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FortimanagerDVMScript resource.
 */
export interface FortimanagerDVMScriptArgs {
    /**
     * ADOM name. default is 'root'.
     */
    adom?: pulumi.Input<string>;
    /**
     * Script content, only cli script is supported now
     */
    content: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Script name.
     */
    name?: pulumi.Input<string>;
    /**
     * Script target, Enum: ["deviceDatabase", "remoteDevice", "adomDatabase"]
     */
    target?: pulumi.Input<string>;
}