// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure sniffer.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.FirewallSniffer("trname", {
 *     applicationListStatus: "disable",
 *     avProfileStatus: "disable",
 *     dlpSensorStatus: "disable",
 *     dsri: "disable",
 *     fosid: 1,
 *     interface: "port4",
 *     ipsDosStatus: "disable",
 *     ipsSensorStatus: "disable",
 *     ipv6: "disable",
 *     logtraffic: "utm",
 *     maxPacketCount: 4000,
 *     nonIp: "enable",
 *     scanBotnetConnections: "disable",
 *     spamfilterProfileStatus: "disable",
 *     status: "enable",
 *     webfilterProfileStatus: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall Sniffer can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallSniffer:FirewallSniffer labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallSniffer extends pulumi.CustomResource {
    /**
     * Get an existing FirewallSniffer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallSnifferState, opts?: pulumi.CustomResourceOptions): FirewallSniffer {
        return new FirewallSniffer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallSniffer:FirewallSniffer';

    /**
     * Returns true if the given object is an instance of FirewallSniffer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallSniffer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallSniffer.__pulumiType;
    }

    /**
     * Configuration method to edit Denial of Service (DoS) anomaly settings. The structure of `anomaly` block is documented below.
     */
    public readonly anomalies!: pulumi.Output<outputs.FirewallSnifferAnomaly[] | undefined>;
    /**
     * Name of an existing application list.
     */
    public readonly applicationList!: pulumi.Output<string>;
    /**
     * Enable/disable application control profile. Valid values: `enable`, `disable`.
     */
    public readonly applicationListStatus!: pulumi.Output<string>;
    /**
     * Name of an existing antivirus profile.
     */
    public readonly avProfile!: pulumi.Output<string>;
    /**
     * Enable/disable antivirus profile. Valid values: `enable`, `disable`.
     */
    public readonly avProfileStatus!: pulumi.Output<string>;
    /**
     * Name of an existing DLP sensor.
     */
    public readonly dlpSensor!: pulumi.Output<string>;
    /**
     * Enable/disable DLP sensor. Valid values: `enable`, `disable`.
     */
    public readonly dlpSensorStatus!: pulumi.Output<string>;
    /**
     * Enable/disable DSRI. Valid values: `enable`, `disable`.
     */
    public readonly dsri!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Name of an existing email filter profile.
     */
    public readonly emailfilterProfile!: pulumi.Output<string>;
    /**
     * Enable/disable emailfilter. Valid values: `enable`, `disable`.
     */
    public readonly emailfilterProfileStatus!: pulumi.Output<string>;
    /**
     * Sniffer ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * Interface name that traffic sniffing will take place on.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Enable/disable IP threat feed. Valid values: `enable`, `disable`.
     */
    public readonly ipThreatfeedStatus!: pulumi.Output<string>;
    /**
     * Name of an existing IP threat feed. The structure of `ipThreatfeed` block is documented below.
     */
    public readonly ipThreatfeeds!: pulumi.Output<outputs.FirewallSnifferIpThreatfeed[] | undefined>;
    /**
     * Enable/disable IPS DoS anomaly detection. Valid values: `enable`, `disable`.
     */
    public readonly ipsDosStatus!: pulumi.Output<string>;
    /**
     * Name of an existing IPS sensor.
     */
    public readonly ipsSensor!: pulumi.Output<string>;
    /**
     * Enable/disable IPS sensor. Valid values: `enable`, `disable`.
     */
    public readonly ipsSensorStatus!: pulumi.Output<string>;
    /**
     * Enable/disable sniffing IPv6 packets. Valid values: `enable`, `disable`.
     */
    public readonly ipv6!: pulumi.Output<string>;
    /**
     * Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy. Valid values: `all`, `utm`, `disable`.
     */
    public readonly logtraffic!: pulumi.Output<string>;
    /**
     * Maximum packet count (1 - 1000000, default = 10000).
     */
    public readonly maxPacketCount!: pulumi.Output<number>;
    /**
     * Enable/disable sniffing non-IP packets. Valid values: `enable`, `disable`.
     */
    public readonly nonIp!: pulumi.Output<string>;
    /**
     * Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).
     */
    public readonly port!: pulumi.Output<string>;
    /**
     * Integer value for the protocol type as defined by IANA (0 - 255).
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Enable/disable scanning of connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
     */
    public readonly scanBotnetConnections!: pulumi.Output<string>;
    /**
     * Name of an existing spam filter profile.
     */
    public readonly spamfilterProfile!: pulumi.Output<string>;
    /**
     * Enable/disable spam filter. Valid values: `enable`, `disable`.
     */
    public readonly spamfilterProfileStatus!: pulumi.Output<string>;
    /**
     * Enable/disable this anomaly. Valid values: `disable`, `enable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * List of VLANs to sniff.
     */
    public readonly vlan!: pulumi.Output<string>;
    /**
     * Name of an existing web filter profile.
     */
    public readonly webfilterProfile!: pulumi.Output<string>;
    /**
     * Enable/disable web filter profile. Valid values: `enable`, `disable`.
     */
    public readonly webfilterProfileStatus!: pulumi.Output<string>;

    /**
     * Create a FirewallSniffer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallSnifferArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallSnifferArgs | FirewallSnifferState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallSnifferState | undefined;
            inputs["anomalies"] = state ? state.anomalies : undefined;
            inputs["applicationList"] = state ? state.applicationList : undefined;
            inputs["applicationListStatus"] = state ? state.applicationListStatus : undefined;
            inputs["avProfile"] = state ? state.avProfile : undefined;
            inputs["avProfileStatus"] = state ? state.avProfileStatus : undefined;
            inputs["dlpSensor"] = state ? state.dlpSensor : undefined;
            inputs["dlpSensorStatus"] = state ? state.dlpSensorStatus : undefined;
            inputs["dsri"] = state ? state.dsri : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["emailfilterProfile"] = state ? state.emailfilterProfile : undefined;
            inputs["emailfilterProfileStatus"] = state ? state.emailfilterProfileStatus : undefined;
            inputs["fosid"] = state ? state.fosid : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["interface"] = state ? state.interface : undefined;
            inputs["ipThreatfeedStatus"] = state ? state.ipThreatfeedStatus : undefined;
            inputs["ipThreatfeeds"] = state ? state.ipThreatfeeds : undefined;
            inputs["ipsDosStatus"] = state ? state.ipsDosStatus : undefined;
            inputs["ipsSensor"] = state ? state.ipsSensor : undefined;
            inputs["ipsSensorStatus"] = state ? state.ipsSensorStatus : undefined;
            inputs["ipv6"] = state ? state.ipv6 : undefined;
            inputs["logtraffic"] = state ? state.logtraffic : undefined;
            inputs["maxPacketCount"] = state ? state.maxPacketCount : undefined;
            inputs["nonIp"] = state ? state.nonIp : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["scanBotnetConnections"] = state ? state.scanBotnetConnections : undefined;
            inputs["spamfilterProfile"] = state ? state.spamfilterProfile : undefined;
            inputs["spamfilterProfileStatus"] = state ? state.spamfilterProfileStatus : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["vlan"] = state ? state.vlan : undefined;
            inputs["webfilterProfile"] = state ? state.webfilterProfile : undefined;
            inputs["webfilterProfileStatus"] = state ? state.webfilterProfileStatus : undefined;
        } else {
            const args = argsOrState as FirewallSnifferArgs | undefined;
            if ((!args || args.interface === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interface'");
            }
            inputs["anomalies"] = args ? args.anomalies : undefined;
            inputs["applicationList"] = args ? args.applicationList : undefined;
            inputs["applicationListStatus"] = args ? args.applicationListStatus : undefined;
            inputs["avProfile"] = args ? args.avProfile : undefined;
            inputs["avProfileStatus"] = args ? args.avProfileStatus : undefined;
            inputs["dlpSensor"] = args ? args.dlpSensor : undefined;
            inputs["dlpSensorStatus"] = args ? args.dlpSensorStatus : undefined;
            inputs["dsri"] = args ? args.dsri : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["emailfilterProfile"] = args ? args.emailfilterProfile : undefined;
            inputs["emailfilterProfileStatus"] = args ? args.emailfilterProfileStatus : undefined;
            inputs["fosid"] = args ? args.fosid : undefined;
            inputs["host"] = args ? args.host : undefined;
            inputs["interface"] = args ? args.interface : undefined;
            inputs["ipThreatfeedStatus"] = args ? args.ipThreatfeedStatus : undefined;
            inputs["ipThreatfeeds"] = args ? args.ipThreatfeeds : undefined;
            inputs["ipsDosStatus"] = args ? args.ipsDosStatus : undefined;
            inputs["ipsSensor"] = args ? args.ipsSensor : undefined;
            inputs["ipsSensorStatus"] = args ? args.ipsSensorStatus : undefined;
            inputs["ipv6"] = args ? args.ipv6 : undefined;
            inputs["logtraffic"] = args ? args.logtraffic : undefined;
            inputs["maxPacketCount"] = args ? args.maxPacketCount : undefined;
            inputs["nonIp"] = args ? args.nonIp : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["scanBotnetConnections"] = args ? args.scanBotnetConnections : undefined;
            inputs["spamfilterProfile"] = args ? args.spamfilterProfile : undefined;
            inputs["spamfilterProfileStatus"] = args ? args.spamfilterProfileStatus : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["vlan"] = args ? args.vlan : undefined;
            inputs["webfilterProfile"] = args ? args.webfilterProfile : undefined;
            inputs["webfilterProfileStatus"] = args ? args.webfilterProfileStatus : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FirewallSniffer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallSniffer resources.
 */
export interface FirewallSnifferState {
    /**
     * Configuration method to edit Denial of Service (DoS) anomaly settings. The structure of `anomaly` block is documented below.
     */
    anomalies?: pulumi.Input<pulumi.Input<inputs.FirewallSnifferAnomaly>[]>;
    /**
     * Name of an existing application list.
     */
    applicationList?: pulumi.Input<string>;
    /**
     * Enable/disable application control profile. Valid values: `enable`, `disable`.
     */
    applicationListStatus?: pulumi.Input<string>;
    /**
     * Name of an existing antivirus profile.
     */
    avProfile?: pulumi.Input<string>;
    /**
     * Enable/disable antivirus profile. Valid values: `enable`, `disable`.
     */
    avProfileStatus?: pulumi.Input<string>;
    /**
     * Name of an existing DLP sensor.
     */
    dlpSensor?: pulumi.Input<string>;
    /**
     * Enable/disable DLP sensor. Valid values: `enable`, `disable`.
     */
    dlpSensorStatus?: pulumi.Input<string>;
    /**
     * Enable/disable DSRI. Valid values: `enable`, `disable`.
     */
    dsri?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Name of an existing email filter profile.
     */
    emailfilterProfile?: pulumi.Input<string>;
    /**
     * Enable/disable emailfilter. Valid values: `enable`, `disable`.
     */
    emailfilterProfileStatus?: pulumi.Input<string>;
    /**
     * Sniffer ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).
     */
    host?: pulumi.Input<string>;
    /**
     * Interface name that traffic sniffing will take place on.
     */
    interface?: pulumi.Input<string>;
    /**
     * Enable/disable IP threat feed. Valid values: `enable`, `disable`.
     */
    ipThreatfeedStatus?: pulumi.Input<string>;
    /**
     * Name of an existing IP threat feed. The structure of `ipThreatfeed` block is documented below.
     */
    ipThreatfeeds?: pulumi.Input<pulumi.Input<inputs.FirewallSnifferIpThreatfeed>[]>;
    /**
     * Enable/disable IPS DoS anomaly detection. Valid values: `enable`, `disable`.
     */
    ipsDosStatus?: pulumi.Input<string>;
    /**
     * Name of an existing IPS sensor.
     */
    ipsSensor?: pulumi.Input<string>;
    /**
     * Enable/disable IPS sensor. Valid values: `enable`, `disable`.
     */
    ipsSensorStatus?: pulumi.Input<string>;
    /**
     * Enable/disable sniffing IPv6 packets. Valid values: `enable`, `disable`.
     */
    ipv6?: pulumi.Input<string>;
    /**
     * Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy. Valid values: `all`, `utm`, `disable`.
     */
    logtraffic?: pulumi.Input<string>;
    /**
     * Maximum packet count (1 - 1000000, default = 10000).
     */
    maxPacketCount?: pulumi.Input<number>;
    /**
     * Enable/disable sniffing non-IP packets. Valid values: `enable`, `disable`.
     */
    nonIp?: pulumi.Input<string>;
    /**
     * Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).
     */
    port?: pulumi.Input<string>;
    /**
     * Integer value for the protocol type as defined by IANA (0 - 255).
     */
    protocol?: pulumi.Input<string>;
    /**
     * Enable/disable scanning of connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
     */
    scanBotnetConnections?: pulumi.Input<string>;
    /**
     * Name of an existing spam filter profile.
     */
    spamfilterProfile?: pulumi.Input<string>;
    /**
     * Enable/disable spam filter. Valid values: `enable`, `disable`.
     */
    spamfilterProfileStatus?: pulumi.Input<string>;
    /**
     * Enable/disable this anomaly. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * List of VLANs to sniff.
     */
    vlan?: pulumi.Input<string>;
    /**
     * Name of an existing web filter profile.
     */
    webfilterProfile?: pulumi.Input<string>;
    /**
     * Enable/disable web filter profile. Valid values: `enable`, `disable`.
     */
    webfilterProfileStatus?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallSniffer resource.
 */
export interface FirewallSnifferArgs {
    /**
     * Configuration method to edit Denial of Service (DoS) anomaly settings. The structure of `anomaly` block is documented below.
     */
    anomalies?: pulumi.Input<pulumi.Input<inputs.FirewallSnifferAnomaly>[]>;
    /**
     * Name of an existing application list.
     */
    applicationList?: pulumi.Input<string>;
    /**
     * Enable/disable application control profile. Valid values: `enable`, `disable`.
     */
    applicationListStatus?: pulumi.Input<string>;
    /**
     * Name of an existing antivirus profile.
     */
    avProfile?: pulumi.Input<string>;
    /**
     * Enable/disable antivirus profile. Valid values: `enable`, `disable`.
     */
    avProfileStatus?: pulumi.Input<string>;
    /**
     * Name of an existing DLP sensor.
     */
    dlpSensor?: pulumi.Input<string>;
    /**
     * Enable/disable DLP sensor. Valid values: `enable`, `disable`.
     */
    dlpSensorStatus?: pulumi.Input<string>;
    /**
     * Enable/disable DSRI. Valid values: `enable`, `disable`.
     */
    dsri?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Name of an existing email filter profile.
     */
    emailfilterProfile?: pulumi.Input<string>;
    /**
     * Enable/disable emailfilter. Valid values: `enable`, `disable`.
     */
    emailfilterProfileStatus?: pulumi.Input<string>;
    /**
     * Sniffer ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).
     */
    host?: pulumi.Input<string>;
    /**
     * Interface name that traffic sniffing will take place on.
     */
    interface: pulumi.Input<string>;
    /**
     * Enable/disable IP threat feed. Valid values: `enable`, `disable`.
     */
    ipThreatfeedStatus?: pulumi.Input<string>;
    /**
     * Name of an existing IP threat feed. The structure of `ipThreatfeed` block is documented below.
     */
    ipThreatfeeds?: pulumi.Input<pulumi.Input<inputs.FirewallSnifferIpThreatfeed>[]>;
    /**
     * Enable/disable IPS DoS anomaly detection. Valid values: `enable`, `disable`.
     */
    ipsDosStatus?: pulumi.Input<string>;
    /**
     * Name of an existing IPS sensor.
     */
    ipsSensor?: pulumi.Input<string>;
    /**
     * Enable/disable IPS sensor. Valid values: `enable`, `disable`.
     */
    ipsSensorStatus?: pulumi.Input<string>;
    /**
     * Enable/disable sniffing IPv6 packets. Valid values: `enable`, `disable`.
     */
    ipv6?: pulumi.Input<string>;
    /**
     * Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy. Valid values: `all`, `utm`, `disable`.
     */
    logtraffic?: pulumi.Input<string>;
    /**
     * Maximum packet count (1 - 1000000, default = 10000).
     */
    maxPacketCount?: pulumi.Input<number>;
    /**
     * Enable/disable sniffing non-IP packets. Valid values: `enable`, `disable`.
     */
    nonIp?: pulumi.Input<string>;
    /**
     * Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).
     */
    port?: pulumi.Input<string>;
    /**
     * Integer value for the protocol type as defined by IANA (0 - 255).
     */
    protocol?: pulumi.Input<string>;
    /**
     * Enable/disable scanning of connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
     */
    scanBotnetConnections?: pulumi.Input<string>;
    /**
     * Name of an existing spam filter profile.
     */
    spamfilterProfile?: pulumi.Input<string>;
    /**
     * Enable/disable spam filter. Valid values: `enable`, `disable`.
     */
    spamfilterProfileStatus?: pulumi.Input<string>;
    /**
     * Enable/disable this anomaly. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * List of VLANs to sniff.
     */
    vlan?: pulumi.Input<string>;
    /**
     * Name of an existing web filter profile.
     */
    webfilterProfile?: pulumi.Input<string>;
    /**
     * Enable/disable web filter profile. Valid values: `enable`, `disable`.
     */
    webfilterProfileStatus?: pulumi.Input<string>;
}