// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure FortiSwitch RSPAN/ERSPAN traffic sniffing parameters. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * SwitchController TrafficSniffer can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerTrafficSniffer:SwitchControllerTrafficSniffer labelname SwitchControllerTrafficSniffer
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/switchControllerTrafficSniffer:SwitchControllerTrafficSniffer labelname SwitchControllerTrafficSniffer
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SwitchControllerTrafficSniffer extends pulumi.CustomResource {
    /**
     * Get an existing SwitchControllerTrafficSniffer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SwitchControllerTrafficSnifferState, opts?: pulumi.CustomResourceOptions): SwitchControllerTrafficSniffer {
        return new SwitchControllerTrafficSniffer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/switchControllerTrafficSniffer:SwitchControllerTrafficSniffer';

    /**
     * Returns true if the given object is an instance of SwitchControllerTrafficSniffer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SwitchControllerTrafficSniffer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SwitchControllerTrafficSniffer.__pulumiType;
    }

    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Configure ERSPAN collector IP address.
     */
    public readonly erspanIp!: pulumi.Output<string>;
    /**
     * Configure traffic sniffer mode. Valid values: `erspan-auto`, `rspan`, `none`.
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * Sniffer IPs to filter. The structure of `targetIp` block is documented below.
     */
    public readonly targetIps!: pulumi.Output<outputs.SwitchControllerTrafficSnifferTargetIp[] | undefined>;
    /**
     * Sniffer MACs to filter. The structure of `targetMac` block is documented below.
     */
    public readonly targetMacs!: pulumi.Output<outputs.SwitchControllerTrafficSnifferTargetMac[] | undefined>;
    /**
     * Sniffer ports to filter. The structure of `targetPort` block is documented below.
     */
    public readonly targetPorts!: pulumi.Output<outputs.SwitchControllerTrafficSnifferTargetPort[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SwitchControllerTrafficSniffer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SwitchControllerTrafficSnifferArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SwitchControllerTrafficSnifferArgs | SwitchControllerTrafficSnifferState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SwitchControllerTrafficSnifferState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["erspanIp"] = state ? state.erspanIp : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["targetIps"] = state ? state.targetIps : undefined;
            resourceInputs["targetMacs"] = state ? state.targetMacs : undefined;
            resourceInputs["targetPorts"] = state ? state.targetPorts : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SwitchControllerTrafficSnifferArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["erspanIp"] = args ? args.erspanIp : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["targetIps"] = args ? args.targetIps : undefined;
            resourceInputs["targetMacs"] = args ? args.targetMacs : undefined;
            resourceInputs["targetPorts"] = args ? args.targetPorts : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SwitchControllerTrafficSniffer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SwitchControllerTrafficSniffer resources.
 */
export interface SwitchControllerTrafficSnifferState {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Configure ERSPAN collector IP address.
     */
    erspanIp?: pulumi.Input<string>;
    /**
     * Configure traffic sniffer mode. Valid values: `erspan-auto`, `rspan`, `none`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Sniffer IPs to filter. The structure of `targetIp` block is documented below.
     */
    targetIps?: pulumi.Input<pulumi.Input<inputs.SwitchControllerTrafficSnifferTargetIp>[]>;
    /**
     * Sniffer MACs to filter. The structure of `targetMac` block is documented below.
     */
    targetMacs?: pulumi.Input<pulumi.Input<inputs.SwitchControllerTrafficSnifferTargetMac>[]>;
    /**
     * Sniffer ports to filter. The structure of `targetPort` block is documented below.
     */
    targetPorts?: pulumi.Input<pulumi.Input<inputs.SwitchControllerTrafficSnifferTargetPort>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SwitchControllerTrafficSniffer resource.
 */
export interface SwitchControllerTrafficSnifferArgs {
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Configure ERSPAN collector IP address.
     */
    erspanIp?: pulumi.Input<string>;
    /**
     * Configure traffic sniffer mode. Valid values: `erspan-auto`, `rspan`, `none`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Sniffer IPs to filter. The structure of `targetIp` block is documented below.
     */
    targetIps?: pulumi.Input<pulumi.Input<inputs.SwitchControllerTrafficSnifferTargetIp>[]>;
    /**
     * Sniffer MACs to filter. The structure of `targetMac` block is documented below.
     */
    targetMacs?: pulumi.Input<pulumi.Input<inputs.SwitchControllerTrafficSnifferTargetMac>[]>;
    /**
     * Sniffer ports to filter. The structure of `targetPort` block is documented below.
     */
    targetPorts?: pulumi.Input<pulumi.Input<inputs.SwitchControllerTrafficSnifferTargetPort>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
