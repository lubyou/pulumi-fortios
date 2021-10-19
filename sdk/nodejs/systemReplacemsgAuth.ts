// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Replacement messages.
 *
 * ## Import
 *
 * SystemReplacemsg Auth can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemReplacemsgAuth:SystemReplacemsgAuth labelname {{msg_type}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemReplacemsgAuth extends pulumi.CustomResource {
    /**
     * Get an existing SystemReplacemsgAuth resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemReplacemsgAuthState, opts?: pulumi.CustomResourceOptions): SystemReplacemsgAuth {
        return new SystemReplacemsgAuth(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemReplacemsgAuth:SystemReplacemsgAuth';

    /**
     * Returns true if the given object is an instance of SystemReplacemsgAuth.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemReplacemsgAuth {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemReplacemsgAuth.__pulumiType;
    }

    /**
     * Message string.
     */
    public readonly buffer!: pulumi.Output<string | undefined>;
    /**
     * Format flag.
     */
    public readonly format!: pulumi.Output<string>;
    /**
     * Header flag. Valid values: `none`, `http`, `8bit`.
     */
    public readonly header!: pulumi.Output<string>;
    /**
     * Message type.
     */
    public readonly msgType!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemReplacemsgAuth resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemReplacemsgAuthArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemReplacemsgAuthArgs | SystemReplacemsgAuthState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemReplacemsgAuthState | undefined;
            inputs["buffer"] = state ? state.buffer : undefined;
            inputs["format"] = state ? state.format : undefined;
            inputs["header"] = state ? state.header : undefined;
            inputs["msgType"] = state ? state.msgType : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemReplacemsgAuthArgs | undefined;
            if ((!args || args.msgType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'msgType'");
            }
            inputs["buffer"] = args ? args.buffer : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["header"] = args ? args.header : undefined;
            inputs["msgType"] = args ? args.msgType : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SystemReplacemsgAuth.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemReplacemsgAuth resources.
 */
export interface SystemReplacemsgAuthState {
    /**
     * Message string.
     */
    buffer?: pulumi.Input<string>;
    /**
     * Format flag.
     */
    format?: pulumi.Input<string>;
    /**
     * Header flag. Valid values: `none`, `http`, `8bit`.
     */
    header?: pulumi.Input<string>;
    /**
     * Message type.
     */
    msgType?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemReplacemsgAuth resource.
 */
export interface SystemReplacemsgAuthArgs {
    /**
     * Message string.
     */
    buffer?: pulumi.Input<string>;
    /**
     * Format flag.
     */
    format?: pulumi.Input<string>;
    /**
     * Header flag. Valid values: `none`, `http`, `8bit`.
     */
    header?: pulumi.Input<string>;
    /**
     * Message type.
     */
    msgType: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
