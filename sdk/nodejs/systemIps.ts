// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure IPS system settings. Applies to FortiOS Version `>= 6.4.2`.
 *
 * ## Import
 *
 * System Ips can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemIps:SystemIps labelname SystemIps
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemIps extends pulumi.CustomResource {
    /**
     * Get an existing SystemIps resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemIpsState, opts?: pulumi.CustomResourceOptions): SystemIps {
        return new SystemIps(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemIps:SystemIps';

    /**
     * Returns true if the given object is an instance of SystemIps.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemIps {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemIps.__pulumiType;
    }

    /**
     * Enable/disable override of hold of triggering signatures that are specified by IDs regardless of hold. Valid values: `enable`, `disable`.
     */
    public readonly overrideSignatureHoldById!: pulumi.Output<string>;
    /**
     * Time to hold and monitor IPS signatures. Format <#d##h> (day range: 0 - 7, hour range: 0 - 23, max hold time: 7d0h, default hold time: 0d0h).
     */
    public readonly signatureHoldTime!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemIps resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemIpsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemIpsArgs | SystemIpsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemIpsState | undefined;
            inputs["overrideSignatureHoldById"] = state ? state.overrideSignatureHoldById : undefined;
            inputs["signatureHoldTime"] = state ? state.signatureHoldTime : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemIpsArgs | undefined;
            inputs["overrideSignatureHoldById"] = args ? args.overrideSignatureHoldById : undefined;
            inputs["signatureHoldTime"] = args ? args.signatureHoldTime : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SystemIps.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemIps resources.
 */
export interface SystemIpsState {
    /**
     * Enable/disable override of hold of triggering signatures that are specified by IDs regardless of hold. Valid values: `enable`, `disable`.
     */
    overrideSignatureHoldById?: pulumi.Input<string>;
    /**
     * Time to hold and monitor IPS signatures. Format <#d##h> (day range: 0 - 7, hour range: 0 - 23, max hold time: 7d0h, default hold time: 0d0h).
     */
    signatureHoldTime?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemIps resource.
 */
export interface SystemIpsArgs {
    /**
     * Enable/disable override of hold of triggering signatures that are specified by IDs regardless of hold. Valid values: `enable`, `disable`.
     */
    overrideSignatureHoldById?: pulumi.Input<string>;
    /**
     * Time to hold and monitor IPS signatures. Format <#d##h> (day range: 0 - 7, hour range: 0 - 23, max hold time: 7d0h, default hold time: 0d0h).
     */
    signatureHoldTime?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
