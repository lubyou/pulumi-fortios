// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure USB LTE/WIMAX devices. Applies to FortiOS Version `>= 7.0.4`.
 *
 * ## Import
 *
 * System LteModem can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemLteModem:SystemLteModem labelname SystemLteModem
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemLteModem extends pulumi.CustomResource {
    /**
     * Get an existing SystemLteModem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemLteModemState, opts?: pulumi.CustomResourceOptions): SystemLteModem {
        return new SystemLteModem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemLteModem:SystemLteModem';

    /**
     * Returns true if the given object is an instance of SystemLteModem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemLteModem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemLteModem.__pulumiType;
    }

    /**
     * Login APN string for PDP-IP packet data calls.
     */
    public readonly apn!: pulumi.Output<string>;
    /**
     * Authentication type for PDP-IP packet data calls. Valid values: `none`, `pap`, `chap`.
     */
    public readonly authtype!: pulumi.Output<string>;
    /**
     * Extra initialization string for USB LTE/WIMAX devices.
     */
    public readonly extraInit!: pulumi.Output<string>;
    /**
     * Hold down timer (10 - 60 sec).
     */
    public readonly holddownTimer!: pulumi.Output<number>;
    /**
     * The interface that the modem is acting as a redundant interface for.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Modem operation mode. Valid values: `standalone`, `redundant`.
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * Modem port index (0 - 20).
     */
    public readonly modemPort!: pulumi.Output<number>;
    /**
     * Authentication password for PDP-IP packet data calls.
     */
    public readonly passwd!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable USB LTE/WIMAX device. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Authentication username for PDP-IP packet data calls.
     */
    public readonly username!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemLteModem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemLteModemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemLteModemArgs | SystemLteModemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemLteModemState | undefined;
            resourceInputs["apn"] = state ? state.apn : undefined;
            resourceInputs["authtype"] = state ? state.authtype : undefined;
            resourceInputs["extraInit"] = state ? state.extraInit : undefined;
            resourceInputs["holddownTimer"] = state ? state.holddownTimer : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["modemPort"] = state ? state.modemPort : undefined;
            resourceInputs["passwd"] = state ? state.passwd : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemLteModemArgs | undefined;
            resourceInputs["apn"] = args ? args.apn : undefined;
            resourceInputs["authtype"] = args ? args.authtype : undefined;
            resourceInputs["extraInit"] = args ? args.extraInit : undefined;
            resourceInputs["holddownTimer"] = args ? args.holddownTimer : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["modemPort"] = args ? args.modemPort : undefined;
            resourceInputs["passwd"] = args ? args.passwd : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemLteModem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemLteModem resources.
 */
export interface SystemLteModemState {
    /**
     * Login APN string for PDP-IP packet data calls.
     */
    apn?: pulumi.Input<string>;
    /**
     * Authentication type for PDP-IP packet data calls. Valid values: `none`, `pap`, `chap`.
     */
    authtype?: pulumi.Input<string>;
    /**
     * Extra initialization string for USB LTE/WIMAX devices.
     */
    extraInit?: pulumi.Input<string>;
    /**
     * Hold down timer (10 - 60 sec).
     */
    holddownTimer?: pulumi.Input<number>;
    /**
     * The interface that the modem is acting as a redundant interface for.
     */
    interface?: pulumi.Input<string>;
    /**
     * Modem operation mode. Valid values: `standalone`, `redundant`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Modem port index (0 - 20).
     */
    modemPort?: pulumi.Input<number>;
    /**
     * Authentication password for PDP-IP packet data calls.
     */
    passwd?: pulumi.Input<string>;
    /**
     * Enable/disable USB LTE/WIMAX device. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Authentication username for PDP-IP packet data calls.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemLteModem resource.
 */
export interface SystemLteModemArgs {
    /**
     * Login APN string for PDP-IP packet data calls.
     */
    apn?: pulumi.Input<string>;
    /**
     * Authentication type for PDP-IP packet data calls. Valid values: `none`, `pap`, `chap`.
     */
    authtype?: pulumi.Input<string>;
    /**
     * Extra initialization string for USB LTE/WIMAX devices.
     */
    extraInit?: pulumi.Input<string>;
    /**
     * Hold down timer (10 - 60 sec).
     */
    holddownTimer?: pulumi.Input<number>;
    /**
     * The interface that the modem is acting as a redundant interface for.
     */
    interface?: pulumi.Input<string>;
    /**
     * Modem operation mode. Valid values: `standalone`, `redundant`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Modem port index (0 - 20).
     */
    modemPort?: pulumi.Input<number>;
    /**
     * Authentication password for PDP-IP packet data calls.
     */
    passwd?: pulumi.Input<string>;
    /**
     * Enable/disable USB LTE/WIMAX device. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Authentication username for PDP-IP packet data calls.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
