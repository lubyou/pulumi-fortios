// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure template for auto-generated VLANs. Applies to FortiOS Version `>= 6.4.0`.
 *
 * ## Import
 *
 * SwitchControllerInitialConfig Template can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerInitialConfigTemplate:SwitchControllerInitialConfigTemplate labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchControllerInitialConfigTemplate extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerInitialConfigTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerInitialConfigTemplateState, opts?: pulumi.CustomResourceOptions): SwitchControllerInitialConfigTemplate {
        return new SwitchControllerInitialConfigTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerInitialConfigTemplate:SwitchControllerInitialConfigTemplate';

    /**
     * Returns true if the given object is an instance of SwitchControllerInitialConfigTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerInitialConfigTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerInitialConfigTemplate.__pulumiType;
    }

    /**
     * Permitted types of management access to this interface. Valid values: `ping`, `https`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `radius-acct`, `probe-response`, `fabric`, `ftm`.
     */
    public readonly allowaccess!: pulumi.Output<string>;
    /**
     * Automatically allocate interface address and subnet block. Valid values: `enable`, `disable`.
     */
    public readonly autoIp!: pulumi.Output<string>;
    /**
     * Enable/disable a DHCP server on this interface. Valid values: `enable`, `disable`.
     */
    public readonly dhcpServer!: pulumi.Output<string>;
    /**
     * Interface IPv4 address and subnet mask.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * Initial config template name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Unique VLAN ID.
     */
    public readonly vlanid!: pulumi.Output<number>;

    /**
     * Create a SwitchControllerInitialConfigTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerInitialConfigTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerInitialConfigTemplateArgs | SwitchControllerInitialConfigTemplateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerInitialConfigTemplateState | undefined;
            inputs["allowaccess"] = state ? state.allowaccess : undefined;
            inputs["autoIp"] = state ? state.autoIp : undefined;
            inputs["dhcpServer"] = state ? state.dhcpServer : undefined;
            inputs["ip"] = state ? state.ip : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["vlanid"] = state ? state.vlanid : undefined;
        } else {
            const args = argsOrState as SwitchControllerInitialConfigTemplateArgs | undefined;
            inputs["allowaccess"] = args ? args.allowaccess : undefined;
            inputs["autoIp"] = args ? args.autoIp : undefined;
            inputs["dhcpServer"] = args ? args.dhcpServer : undefined;
            inputs["ip"] = args ? args.ip : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["vlanid"] = args ? args.vlanid : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SwitchControllerInitialConfigTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerInitialConfigTemplate resources.
 */
export interface SwitchControllerInitialConfigTemplateState {
    /**
     * Permitted types of management access to this interface. Valid values: `ping`, `https`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `radius-acct`, `probe-response`, `fabric`, `ftm`.
     */
    allowaccess?: pulumi.Input<string>;
    /**
     * Automatically allocate interface address and subnet block. Valid values: `enable`, `disable`.
     */
    autoIp?: pulumi.Input<string>;
    /**
     * Enable/disable a DHCP server on this interface. Valid values: `enable`, `disable`.
     */
    dhcpServer?: pulumi.Input<string>;
    /**
     * Interface IPv4 address and subnet mask.
     */
    ip?: pulumi.Input<string>;
    /**
     * Initial config template name
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Unique VLAN ID.
     */
    vlanid?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SwitchControllerInitialConfigTemplate resource.
 */
export interface SwitchControllerInitialConfigTemplateArgs {
    /**
     * Permitted types of management access to this interface. Valid values: `ping`, `https`, `ssh`, `snmp`, `http`, `telnet`, `fgfm`, `radius-acct`, `probe-response`, `fabric`, `ftm`.
     */
    allowaccess?: pulumi.Input<string>;
    /**
     * Automatically allocate interface address and subnet block. Valid values: `enable`, `disable`.
     */
    autoIp?: pulumi.Input<string>;
    /**
     * Enable/disable a DHCP server on this interface. Valid values: `enable`, `disable`.
     */
    dhcpServer?: pulumi.Input<string>;
    /**
     * Interface IPv4 address and subnet mask.
     */
    ip?: pulumi.Input<string>;
    /**
     * Initial config template name
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Unique VLAN ID.
     */
    vlanid?: pulumi.Input<number>;
}