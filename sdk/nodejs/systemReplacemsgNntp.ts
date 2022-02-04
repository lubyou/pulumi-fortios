// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Replacement messages. Applies to FortiOS Version `<= 6.4.0`.
 *
 * ## Import
 *
 * SystemReplacemsg Nntp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemReplacemsgNntp:SystemReplacemsgNntp labelname {{msg_type}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemReplacemsgNntp extends pulumi.CustomResource {
    /**
     * Get an existing SystemReplacemsgNntp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemReplacemsgNntpState, opts?: pulumi.CustomResourceOptions): SystemReplacemsgNntp {
        return new SystemReplacemsgNntp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemReplacemsgNntp:SystemReplacemsgNntp';

    /**
     * Returns true if the given object is an instance of SystemReplacemsgNntp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemReplacemsgNntp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemReplacemsgNntp.__pulumiType;
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
     * Create a SystemReplacemsgNntp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemReplacemsgNntpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemReplacemsgNntpArgs | SystemReplacemsgNntpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemReplacemsgNntpState | undefined;
            resourceInputs["buffer"] = state ? state.buffer : undefined;
            resourceInputs["format"] = state ? state.format : undefined;
            resourceInputs["header"] = state ? state.header : undefined;
            resourceInputs["msgType"] = state ? state.msgType : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemReplacemsgNntpArgs | undefined;
            if ((!args || args.msgType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'msgType'");
            }
            resourceInputs["buffer"] = args ? args.buffer : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["header"] = args ? args.header : undefined;
            resourceInputs["msgType"] = args ? args.msgType : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemReplacemsgNntp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemReplacemsgNntp resources.
 */
export interface SystemReplacemsgNntpState {
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
 * The set of arguments for constructing a SystemReplacemsgNntp resource.
 */
export interface SystemReplacemsgNntpArgs {
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
