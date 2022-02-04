// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure AntiSpam options. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * Emailfilter Options can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/emailfilterOptions:EmailfilterOptions labelname EmailfilterOptions
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class EmailfilterOptions extends pulumi.CustomResource {
    /**
     * Get an existing EmailfilterOptions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EmailfilterOptionsState, opts?: pulumi.CustomResourceOptions): EmailfilterOptions {
        return new EmailfilterOptions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/emailfilterOptions:EmailfilterOptions';

    /**
     * Returns true if the given object is an instance of EmailfilterOptions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EmailfilterOptions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EmailfilterOptions.__pulumiType;
    }

    /**
     * DNS query time out (1 - 30 sec).
     */
    public readonly dnsTimeout!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a EmailfilterOptions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EmailfilterOptionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EmailfilterOptionsArgs | EmailfilterOptionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EmailfilterOptionsState | undefined;
            resourceInputs["dnsTimeout"] = state ? state.dnsTimeout : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as EmailfilterOptionsArgs | undefined;
            resourceInputs["dnsTimeout"] = args ? args.dnsTimeout : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EmailfilterOptions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EmailfilterOptions resources.
 */
export interface EmailfilterOptionsState {
    /**
     * DNS query time out (1 - 30 sec).
     */
    dnsTimeout?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EmailfilterOptions resource.
 */
export interface EmailfilterOptionsArgs {
    /**
     * DNS query time out (1 - 30 sec).
     */
    dnsTimeout?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
