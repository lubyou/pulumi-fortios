// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure NPU attributes. Applies to FortiOS Version `>= 7.0.4`.
 *
 * ## Import
 *
 * System Npu can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemNpu:SystemNpu labelname SystemNpu
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemNpu extends pulumi.CustomResource {
    /**
     * Get an existing SystemNpu resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemNpuState, opts?: pulumi.CustomResourceOptions): SystemNpu {
        return new SystemNpu(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemNpu:SystemNpu';

    /**
     * Returns true if the given object is an instance of SystemNpu.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemNpu {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemNpu.__pulumiType;
    }

    /**
     * Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `enable`, `disable`.
     */
    public readonly capwapOffload!: pulumi.Output<string>;
    /**
     * Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
     */
    public readonly dedicatedManagementAffinity!: pulumi.Output<string>;
    /**
     * Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `enable`, `disable`.
     */
    public readonly dedicatedManagementCpu!: pulumi.Output<string>;
    /**
     * Enable/disable NP6 offloading (also called fast path). Valid values: `disable`, `enable`.
     */
    public readonly fastpath!: pulumi.Output<string>;
    /**
     * IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
     */
    public readonly ipsecDecSubengineMask!: pulumi.Output<string>;
    /**
     * IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
     */
    public readonly ipsecEncSubengineMask!: pulumi.Output<string>;
    /**
     * Enable/disable IPsec inbound cache for anti-replay. Valid values: `enable`, `disable`.
     */
    public readonly ipsecInboundCache!: pulumi.Output<string>;
    /**
     * Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.
     */
    public readonly ipsecMtuOverride!: pulumi.Output<string>;
    /**
     * Enable/disable IPSEC over vlink. Valid values: `enable`, `disable`.
     */
    public readonly ipsecOverVlink!: pulumi.Output<string>;
    /**
     * Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `tpe-based`, `session-based`, `disable`.
     */
    public readonly mcastSessionAccounting!: pulumi.Output<string>;
    /**
     * Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `enable`, `disable`.
     */
    public readonly np6CpsOptimizationMode!: pulumi.Output<string>;
    /**
     * Configure NPU priority protocol. The structure of `priorityProtocol` block is documented below.
     */
    public readonly priorityProtocol!: pulumi.Output<outputs.SystemNpuPriorityProtocol | undefined>;
    /**
     * Enable/disable rdp offload. Valid values: `enable`, `disable`.
     */
    public readonly rdpOffload!: pulumi.Output<string>;
    /**
     * Enable/disable offloading of denied sessions. Requires ses-denied-traffic to be set. Valid values: `disable`, `enable`.
     */
    public readonly sessionDeniedOffload!: pulumi.Output<string>;
    /**
     * Enable/disable sse backpressure. Valid values: `enable`, `disable`.
     */
    public readonly sseBackpressure!: pulumi.Output<string>;
    /**
     * Enable/disable stripping clear text padding. Valid values: `enable`, `disable`.
     */
    public readonly stripClearTextPadding!: pulumi.Output<string>;
    /**
     * Enable/disable stripping ESP padding. Valid values: `enable`, `disable`.
     */
    public readonly stripEspPadding!: pulumi.Output<string>;
    /**
     * Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`.
     */
    public readonly swNpBandwidth!: pulumi.Output<string>;
    /**
     * Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly uespOffload!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemNpu resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemNpuArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemNpuArgs | SystemNpuState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemNpuState | undefined;
            resourceInputs["capwapOffload"] = state ? state.capwapOffload : undefined;
            resourceInputs["dedicatedManagementAffinity"] = state ? state.dedicatedManagementAffinity : undefined;
            resourceInputs["dedicatedManagementCpu"] = state ? state.dedicatedManagementCpu : undefined;
            resourceInputs["fastpath"] = state ? state.fastpath : undefined;
            resourceInputs["ipsecDecSubengineMask"] = state ? state.ipsecDecSubengineMask : undefined;
            resourceInputs["ipsecEncSubengineMask"] = state ? state.ipsecEncSubengineMask : undefined;
            resourceInputs["ipsecInboundCache"] = state ? state.ipsecInboundCache : undefined;
            resourceInputs["ipsecMtuOverride"] = state ? state.ipsecMtuOverride : undefined;
            resourceInputs["ipsecOverVlink"] = state ? state.ipsecOverVlink : undefined;
            resourceInputs["mcastSessionAccounting"] = state ? state.mcastSessionAccounting : undefined;
            resourceInputs["np6CpsOptimizationMode"] = state ? state.np6CpsOptimizationMode : undefined;
            resourceInputs["priorityProtocol"] = state ? state.priorityProtocol : undefined;
            resourceInputs["rdpOffload"] = state ? state.rdpOffload : undefined;
            resourceInputs["sessionDeniedOffload"] = state ? state.sessionDeniedOffload : undefined;
            resourceInputs["sseBackpressure"] = state ? state.sseBackpressure : undefined;
            resourceInputs["stripClearTextPadding"] = state ? state.stripClearTextPadding : undefined;
            resourceInputs["stripEspPadding"] = state ? state.stripEspPadding : undefined;
            resourceInputs["swNpBandwidth"] = state ? state.swNpBandwidth : undefined;
            resourceInputs["uespOffload"] = state ? state.uespOffload : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemNpuArgs | undefined;
            resourceInputs["capwapOffload"] = args ? args.capwapOffload : undefined;
            resourceInputs["dedicatedManagementAffinity"] = args ? args.dedicatedManagementAffinity : undefined;
            resourceInputs["dedicatedManagementCpu"] = args ? args.dedicatedManagementCpu : undefined;
            resourceInputs["fastpath"] = args ? args.fastpath : undefined;
            resourceInputs["ipsecDecSubengineMask"] = args ? args.ipsecDecSubengineMask : undefined;
            resourceInputs["ipsecEncSubengineMask"] = args ? args.ipsecEncSubengineMask : undefined;
            resourceInputs["ipsecInboundCache"] = args ? args.ipsecInboundCache : undefined;
            resourceInputs["ipsecMtuOverride"] = args ? args.ipsecMtuOverride : undefined;
            resourceInputs["ipsecOverVlink"] = args ? args.ipsecOverVlink : undefined;
            resourceInputs["mcastSessionAccounting"] = args ? args.mcastSessionAccounting : undefined;
            resourceInputs["np6CpsOptimizationMode"] = args ? args.np6CpsOptimizationMode : undefined;
            resourceInputs["priorityProtocol"] = args ? args.priorityProtocol : undefined;
            resourceInputs["rdpOffload"] = args ? args.rdpOffload : undefined;
            resourceInputs["sessionDeniedOffload"] = args ? args.sessionDeniedOffload : undefined;
            resourceInputs["sseBackpressure"] = args ? args.sseBackpressure : undefined;
            resourceInputs["stripClearTextPadding"] = args ? args.stripClearTextPadding : undefined;
            resourceInputs["stripEspPadding"] = args ? args.stripEspPadding : undefined;
            resourceInputs["swNpBandwidth"] = args ? args.swNpBandwidth : undefined;
            resourceInputs["uespOffload"] = args ? args.uespOffload : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemNpu.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemNpu resources.
 */
export interface SystemNpuState {
    /**
     * Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `enable`, `disable`.
     */
    capwapOffload?: pulumi.Input<string>;
    /**
     * Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
     */
    dedicatedManagementAffinity?: pulumi.Input<string>;
    /**
     * Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `enable`, `disable`.
     */
    dedicatedManagementCpu?: pulumi.Input<string>;
    /**
     * Enable/disable NP6 offloading (also called fast path). Valid values: `disable`, `enable`.
     */
    fastpath?: pulumi.Input<string>;
    /**
     * IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
     */
    ipsecDecSubengineMask?: pulumi.Input<string>;
    /**
     * IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
     */
    ipsecEncSubengineMask?: pulumi.Input<string>;
    /**
     * Enable/disable IPsec inbound cache for anti-replay. Valid values: `enable`, `disable`.
     */
    ipsecInboundCache?: pulumi.Input<string>;
    /**
     * Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.
     */
    ipsecMtuOverride?: pulumi.Input<string>;
    /**
     * Enable/disable IPSEC over vlink. Valid values: `enable`, `disable`.
     */
    ipsecOverVlink?: pulumi.Input<string>;
    /**
     * Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `tpe-based`, `session-based`, `disable`.
     */
    mcastSessionAccounting?: pulumi.Input<string>;
    /**
     * Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `enable`, `disable`.
     */
    np6CpsOptimizationMode?: pulumi.Input<string>;
    /**
     * Configure NPU priority protocol. The structure of `priorityProtocol` block is documented below.
     */
    priorityProtocol?: pulumi.Input<inputs.SystemNpuPriorityProtocol>;
    /**
     * Enable/disable rdp offload. Valid values: `enable`, `disable`.
     */
    rdpOffload?: pulumi.Input<string>;
    /**
     * Enable/disable offloading of denied sessions. Requires ses-denied-traffic to be set. Valid values: `disable`, `enable`.
     */
    sessionDeniedOffload?: pulumi.Input<string>;
    /**
     * Enable/disable sse backpressure. Valid values: `enable`, `disable`.
     */
    sseBackpressure?: pulumi.Input<string>;
    /**
     * Enable/disable stripping clear text padding. Valid values: `enable`, `disable`.
     */
    stripClearTextPadding?: pulumi.Input<string>;
    /**
     * Enable/disable stripping ESP padding. Valid values: `enable`, `disable`.
     */
    stripEspPadding?: pulumi.Input<string>;
    /**
     * Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`.
     */
    swNpBandwidth?: pulumi.Input<string>;
    /**
     * Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `enable`, `disable`.
     */
    uespOffload?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemNpu resource.
 */
export interface SystemNpuArgs {
    /**
     * Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `enable`, `disable`.
     */
    capwapOffload?: pulumi.Input<string>;
    /**
     * Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
     */
    dedicatedManagementAffinity?: pulumi.Input<string>;
    /**
     * Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `enable`, `disable`.
     */
    dedicatedManagementCpu?: pulumi.Input<string>;
    /**
     * Enable/disable NP6 offloading (also called fast path). Valid values: `disable`, `enable`.
     */
    fastpath?: pulumi.Input<string>;
    /**
     * IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
     */
    ipsecDecSubengineMask?: pulumi.Input<string>;
    /**
     * IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
     */
    ipsecEncSubengineMask?: pulumi.Input<string>;
    /**
     * Enable/disable IPsec inbound cache for anti-replay. Valid values: `enable`, `disable`.
     */
    ipsecInboundCache?: pulumi.Input<string>;
    /**
     * Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.
     */
    ipsecMtuOverride?: pulumi.Input<string>;
    /**
     * Enable/disable IPSEC over vlink. Valid values: `enable`, `disable`.
     */
    ipsecOverVlink?: pulumi.Input<string>;
    /**
     * Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `tpe-based`, `session-based`, `disable`.
     */
    mcastSessionAccounting?: pulumi.Input<string>;
    /**
     * Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `enable`, `disable`.
     */
    np6CpsOptimizationMode?: pulumi.Input<string>;
    /**
     * Configure NPU priority protocol. The structure of `priorityProtocol` block is documented below.
     */
    priorityProtocol?: pulumi.Input<inputs.SystemNpuPriorityProtocol>;
    /**
     * Enable/disable rdp offload. Valid values: `enable`, `disable`.
     */
    rdpOffload?: pulumi.Input<string>;
    /**
     * Enable/disable offloading of denied sessions. Requires ses-denied-traffic to be set. Valid values: `disable`, `enable`.
     */
    sessionDeniedOffload?: pulumi.Input<string>;
    /**
     * Enable/disable sse backpressure. Valid values: `enable`, `disable`.
     */
    sseBackpressure?: pulumi.Input<string>;
    /**
     * Enable/disable stripping clear text padding. Valid values: `enable`, `disable`.
     */
    stripClearTextPadding?: pulumi.Input<string>;
    /**
     * Enable/disable stripping ESP padding. Valid values: `enable`, `disable`.
     */
    stripEspPadding?: pulumi.Input<string>;
    /**
     * Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`.
     */
    swNpBandwidth?: pulumi.Input<string>;
    /**
     * Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `enable`, `disable`.
     */
    uespOffload?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
