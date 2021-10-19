// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy. Applies to FortiOS Version `>= 6.4.0`.
 *
 * ## Import
 *
 * SwitchController NacDevice can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerNacDevice:SwitchControllerNacDevice labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchControllerNacDevice extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerNacDevice resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerNacDeviceState, opts?: pulumi.CustomResourceOptions): SwitchControllerNacDevice {
        return new SwitchControllerNacDevice(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerNacDevice:SwitchControllerNacDevice';

    /**
     * Returns true if the given object is an instance of SwitchControllerNacDevice.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerNacDevice {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerNacDevice.__pulumiType;
    }

    /**
     * Description for the learned NAC device.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Device ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Managed FortiSwitch port where NAC device is last learned.
     */
    public readonly lastKnownPort!: pulumi.Output<string>;
    /**
     * Managed FortiSwitch where NAC device is last learned.
     */
    public readonly lastKnownSwitch!: pulumi.Output<string>;
    /**
     * Device last seen.
     */
    public readonly lastSeen!: pulumi.Output<number>;
    /**
     * MAC address of the learned NAC device.
     */
    public readonly mac!: pulumi.Output<string>;
    /**
     * MAC policy to be applied on this learned NAC device.
     */
    public readonly macPolicy!: pulumi.Output<string>;
    /**
     * Matched NAC policy for the learned NAC device.
     */
    public readonly matchedNacPolicy!: pulumi.Output<string>;
    /**
     * Port policy to be applied on this learned NAC device.
     */
    public readonly portPolicy!: pulumi.Output<string>;
    /**
     * Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchControllerNacDevice resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerNacDeviceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerNacDeviceArgs | SwitchControllerNacDeviceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerNacDeviceState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["fosid"] = state ? state.fosid : undefined;
            inputs["lastKnownPort"] = state ? state.lastKnownPort : undefined;
            inputs["lastKnownSwitch"] = state ? state.lastKnownSwitch : undefined;
            inputs["lastSeen"] = state ? state.lastSeen : undefined;
            inputs["mac"] = state ? state.mac : undefined;
            inputs["macPolicy"] = state ? state.macPolicy : undefined;
            inputs["matchedNacPolicy"] = state ? state.matchedNacPolicy : undefined;
            inputs["portPolicy"] = state ? state.portPolicy : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchControllerNacDeviceArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["fosid"] = args ? args.fosid : undefined;
            inputs["lastKnownPort"] = args ? args.lastKnownPort : undefined;
            inputs["lastKnownSwitch"] = args ? args.lastKnownSwitch : undefined;
            inputs["lastSeen"] = args ? args.lastSeen : undefined;
            inputs["mac"] = args ? args.mac : undefined;
            inputs["macPolicy"] = args ? args.macPolicy : undefined;
            inputs["matchedNacPolicy"] = args ? args.matchedNacPolicy : undefined;
            inputs["portPolicy"] = args ? args.portPolicy : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SwitchControllerNacDevice.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerNacDevice resources.
 */
export interface SwitchControllerNacDeviceState {
    /**
     * Description for the learned NAC device.
     */
    description?: pulumi.Input<string>;
    /**
     * Device ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Managed FortiSwitch port where NAC device is last learned.
     */
    lastKnownPort?: pulumi.Input<string>;
    /**
     * Managed FortiSwitch where NAC device is last learned.
     */
    lastKnownSwitch?: pulumi.Input<string>;
    /**
     * Device last seen.
     */
    lastSeen?: pulumi.Input<number>;
    /**
     * MAC address of the learned NAC device.
     */
    mac?: pulumi.Input<string>;
    /**
     * MAC policy to be applied on this learned NAC device.
     */
    macPolicy?: pulumi.Input<string>;
    /**
     * Matched NAC policy for the learned NAC device.
     */
    matchedNacPolicy?: pulumi.Input<string>;
    /**
     * Port policy to be applied on this learned NAC device.
     */
    portPolicy?: pulumi.Input<string>;
    /**
     * Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerNacDevice resource.
 */
export interface SwitchControllerNacDeviceArgs {
    /**
     * Description for the learned NAC device.
     */
    description?: pulumi.Input<string>;
    /**
     * Device ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Managed FortiSwitch port where NAC device is last learned.
     */
    lastKnownPort?: pulumi.Input<string>;
    /**
     * Managed FortiSwitch where NAC device is last learned.
     */
    lastKnownSwitch?: pulumi.Input<string>;
    /**
     * Device last seen.
     */
    lastSeen?: pulumi.Input<number>;
    /**
     * MAC address of the learned NAC device.
     */
    mac?: pulumi.Input<string>;
    /**
     * MAC policy to be applied on this learned NAC device.
     */
    macPolicy?: pulumi.Input<string>;
    /**
     * Matched NAC policy for the learned NAC device.
     */
    matchedNacPolicy?: pulumi.Input<string>;
    /**
     * Port policy to be applied on this learned NAC device.
     */
    portPolicy?: pulumi.Input<string>;
    /**
     * Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
