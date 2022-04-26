// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure custom services.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.FirewallServiceCustom("trname", {
 *     appServiceType: "disable",
 *     category: "General",
 *     checkResetRange: "default",
 *     color: 0,
 *     helper: "auto",
 *     iprange: "0.0.0.0",
 *     protocol: "TCP/UDP/SCTP",
 *     protocolNumber: 6,
 *     proxy: "disable",
 *     tcpHalfcloseTimer: 0,
 *     tcpHalfopenTimer: 0,
 *     tcpPortrange: "223-332",
 *     tcpTimewaitTimer: 0,
 *     udpIdleTimer: 0,
 *     visibility: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * FirewallService Custom can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallServiceCustom:FirewallServiceCustom labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallServiceCustom:FirewallServiceCustom labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallServiceCustom extends pulumi.CustomResource {
    /**
     * Get an existing FirewallServiceCustom resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallServiceCustomState, opts?: pulumi.CustomResourceOptions): FirewallServiceCustom {
        return new FirewallServiceCustom(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallServiceCustom:FirewallServiceCustom';

    /**
     * Returns true if the given object is an instance of FirewallServiceCustom.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallServiceCustom {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallServiceCustom.__pulumiType;
    }

    /**
     * Application category ID. The structure of `appCategory` block is documented below.
     */
    public readonly appCategories!: pulumi.Output<outputs.FirewallServiceCustomAppCategory[] | undefined>;
    /**
     * Application service type. Valid values: `disable`, `app-id`, `app-category`.
     */
    public readonly appServiceType!: pulumi.Output<string>;
    /**
     * Application ID. The structure of `application` block is documented below.
     */
    public readonly applications!: pulumi.Output<outputs.FirewallServiceCustomApplication[] | undefined>;
    /**
     * Service category.
     */
    public readonly category!: pulumi.Output<string>;
    /**
     * Configure the type of ICMP error message verification. Valid values: `disable`, `strict`, `default`.
     */
    public readonly checkResetRange!: pulumi.Output<string>;
    /**
     * Color of icon on the GUI.
     */
    public readonly color!: pulumi.Output<number>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Security Fabric global object setting. Valid values: `enable`, `disable`.
     */
    public readonly fabricObject!: pulumi.Output<string>;
    /**
     * Fully qualified domain name.
     */
    public readonly fqdn!: pulumi.Output<string>;
    /**
     * Helper name.
     */
    public readonly helper!: pulumi.Output<string>;
    /**
     * ICMP code.
     */
    public readonly icmpcode!: pulumi.Output<number>;
    /**
     * ICMP type.
     */
    public readonly icmptype!: pulumi.Output<number>;
    /**
     * Start and end of the IP range associated with service.
     */
    public readonly iprange!: pulumi.Output<string>;
    /**
     * Custom service name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Protocol type based on IANA numbers. Valid values: `TCP/UDP/SCTP`, `ICMP`, `ICMP6`, `IP`, `HTTP`, `FTP`, `CONNECT`, `SOCKS-TCP`, `SOCKS-UDP`, `ALL`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * IP protocol number.
     */
    public readonly protocolNumber!: pulumi.Output<number>;
    /**
     * Enable/disable web proxy service. Valid values: `enable`, `disable`.
     */
    public readonly proxy!: pulumi.Output<string>;
    /**
     * Multiple SCTP port ranges.
     */
    public readonly sctpPortrange!: pulumi.Output<string>;
    /**
     * Session TTL (300 - 604800, 0 = default).
     */
    public readonly sessionTtl!: pulumi.Output<number>;
    /**
     * Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
     */
    public readonly tcpHalfcloseTimer!: pulumi.Output<number>;
    /**
     * Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
     */
    public readonly tcpHalfopenTimer!: pulumi.Output<number>;
    /**
     * Multiple TCP port ranges.
     */
    public readonly tcpPortrange!: pulumi.Output<string>;
    /**
     * Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).
     */
    public readonly tcpRstTimer!: pulumi.Output<number>;
    /**
     * Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
     */
    public readonly tcpTimewaitTimer!: pulumi.Output<number>;
    /**
     * UDP half close timeout (0 - 86400 sec, 0 = default).
     */
    public readonly udpIdleTimer!: pulumi.Output<number>;
    /**
     * Multiple UDP port ranges.
     */
    public readonly udpPortrange!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable the visibility of the service on the GUI. Valid values: `enable`, `disable`.
     */
    public readonly visibility!: pulumi.Output<string>;

    /**
     * Create a FirewallServiceCustom resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallServiceCustomArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallServiceCustomArgs | FirewallServiceCustomState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallServiceCustomState | undefined;
            resourceInputs["appCategories"] = state ? state.appCategories : undefined;
            resourceInputs["appServiceType"] = state ? state.appServiceType : undefined;
            resourceInputs["applications"] = state ? state.applications : undefined;
            resourceInputs["category"] = state ? state.category : undefined;
            resourceInputs["checkResetRange"] = state ? state.checkResetRange : undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fabricObject"] = state ? state.fabricObject : undefined;
            resourceInputs["fqdn"] = state ? state.fqdn : undefined;
            resourceInputs["helper"] = state ? state.helper : undefined;
            resourceInputs["icmpcode"] = state ? state.icmpcode : undefined;
            resourceInputs["icmptype"] = state ? state.icmptype : undefined;
            resourceInputs["iprange"] = state ? state.iprange : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["protocolNumber"] = state ? state.protocolNumber : undefined;
            resourceInputs["proxy"] = state ? state.proxy : undefined;
            resourceInputs["sctpPortrange"] = state ? state.sctpPortrange : undefined;
            resourceInputs["sessionTtl"] = state ? state.sessionTtl : undefined;
            resourceInputs["tcpHalfcloseTimer"] = state ? state.tcpHalfcloseTimer : undefined;
            resourceInputs["tcpHalfopenTimer"] = state ? state.tcpHalfopenTimer : undefined;
            resourceInputs["tcpPortrange"] = state ? state.tcpPortrange : undefined;
            resourceInputs["tcpRstTimer"] = state ? state.tcpRstTimer : undefined;
            resourceInputs["tcpTimewaitTimer"] = state ? state.tcpTimewaitTimer : undefined;
            resourceInputs["udpIdleTimer"] = state ? state.udpIdleTimer : undefined;
            resourceInputs["udpPortrange"] = state ? state.udpPortrange : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as FirewallServiceCustomArgs | undefined;
            resourceInputs["appCategories"] = args ? args.appCategories : undefined;
            resourceInputs["appServiceType"] = args ? args.appServiceType : undefined;
            resourceInputs["applications"] = args ? args.applications : undefined;
            resourceInputs["category"] = args ? args.category : undefined;
            resourceInputs["checkResetRange"] = args ? args.checkResetRange : undefined;
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fabricObject"] = args ? args.fabricObject : undefined;
            resourceInputs["fqdn"] = args ? args.fqdn : undefined;
            resourceInputs["helper"] = args ? args.helper : undefined;
            resourceInputs["icmpcode"] = args ? args.icmpcode : undefined;
            resourceInputs["icmptype"] = args ? args.icmptype : undefined;
            resourceInputs["iprange"] = args ? args.iprange : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["protocolNumber"] = args ? args.protocolNumber : undefined;
            resourceInputs["proxy"] = args ? args.proxy : undefined;
            resourceInputs["sctpPortrange"] = args ? args.sctpPortrange : undefined;
            resourceInputs["sessionTtl"] = args ? args.sessionTtl : undefined;
            resourceInputs["tcpHalfcloseTimer"] = args ? args.tcpHalfcloseTimer : undefined;
            resourceInputs["tcpHalfopenTimer"] = args ? args.tcpHalfopenTimer : undefined;
            resourceInputs["tcpPortrange"] = args ? args.tcpPortrange : undefined;
            resourceInputs["tcpRstTimer"] = args ? args.tcpRstTimer : undefined;
            resourceInputs["tcpTimewaitTimer"] = args ? args.tcpTimewaitTimer : undefined;
            resourceInputs["udpIdleTimer"] = args ? args.udpIdleTimer : undefined;
            resourceInputs["udpPortrange"] = args ? args.udpPortrange : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallServiceCustom.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallServiceCustom resources.
 */
export interface FirewallServiceCustomState {
    /**
     * Application category ID. The structure of `appCategory` block is documented below.
     */
    appCategories?: pulumi.Input<pulumi.Input<inputs.FirewallServiceCustomAppCategory>[]>;
    /**
     * Application service type. Valid values: `disable`, `app-id`, `app-category`.
     */
    appServiceType?: pulumi.Input<string>;
    /**
     * Application ID. The structure of `application` block is documented below.
     */
    applications?: pulumi.Input<pulumi.Input<inputs.FirewallServiceCustomApplication>[]>;
    /**
     * Service category.
     */
    category?: pulumi.Input<string>;
    /**
     * Configure the type of ICMP error message verification. Valid values: `disable`, `strict`, `default`.
     */
    checkResetRange?: pulumi.Input<string>;
    /**
     * Color of icon on the GUI.
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Security Fabric global object setting. Valid values: `enable`, `disable`.
     */
    fabricObject?: pulumi.Input<string>;
    /**
     * Fully qualified domain name.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * Helper name.
     */
    helper?: pulumi.Input<string>;
    /**
     * ICMP code.
     */
    icmpcode?: pulumi.Input<number>;
    /**
     * ICMP type.
     */
    icmptype?: pulumi.Input<number>;
    /**
     * Start and end of the IP range associated with service.
     */
    iprange?: pulumi.Input<string>;
    /**
     * Custom service name.
     */
    name?: pulumi.Input<string>;
    /**
     * Protocol type based on IANA numbers. Valid values: `TCP/UDP/SCTP`, `ICMP`, `ICMP6`, `IP`, `HTTP`, `FTP`, `CONNECT`, `SOCKS-TCP`, `SOCKS-UDP`, `ALL`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * IP protocol number.
     */
    protocolNumber?: pulumi.Input<number>;
    /**
     * Enable/disable web proxy service. Valid values: `enable`, `disable`.
     */
    proxy?: pulumi.Input<string>;
    /**
     * Multiple SCTP port ranges.
     */
    sctpPortrange?: pulumi.Input<string>;
    /**
     * Session TTL (300 - 604800, 0 = default).
     */
    sessionTtl?: pulumi.Input<number>;
    /**
     * Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
     */
    tcpHalfcloseTimer?: pulumi.Input<number>;
    /**
     * Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
     */
    tcpHalfopenTimer?: pulumi.Input<number>;
    /**
     * Multiple TCP port ranges.
     */
    tcpPortrange?: pulumi.Input<string>;
    /**
     * Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).
     */
    tcpRstTimer?: pulumi.Input<number>;
    /**
     * Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
     */
    tcpTimewaitTimer?: pulumi.Input<number>;
    /**
     * UDP half close timeout (0 - 86400 sec, 0 = default).
     */
    udpIdleTimer?: pulumi.Input<number>;
    /**
     * Multiple UDP port ranges.
     */
    udpPortrange?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable the visibility of the service on the GUI. Valid values: `enable`, `disable`.
     */
    visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallServiceCustom resource.
 */
export interface FirewallServiceCustomArgs {
    /**
     * Application category ID. The structure of `appCategory` block is documented below.
     */
    appCategories?: pulumi.Input<pulumi.Input<inputs.FirewallServiceCustomAppCategory>[]>;
    /**
     * Application service type. Valid values: `disable`, `app-id`, `app-category`.
     */
    appServiceType?: pulumi.Input<string>;
    /**
     * Application ID. The structure of `application` block is documented below.
     */
    applications?: pulumi.Input<pulumi.Input<inputs.FirewallServiceCustomApplication>[]>;
    /**
     * Service category.
     */
    category?: pulumi.Input<string>;
    /**
     * Configure the type of ICMP error message verification. Valid values: `disable`, `strict`, `default`.
     */
    checkResetRange?: pulumi.Input<string>;
    /**
     * Color of icon on the GUI.
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Security Fabric global object setting. Valid values: `enable`, `disable`.
     */
    fabricObject?: pulumi.Input<string>;
    /**
     * Fully qualified domain name.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * Helper name.
     */
    helper?: pulumi.Input<string>;
    /**
     * ICMP code.
     */
    icmpcode?: pulumi.Input<number>;
    /**
     * ICMP type.
     */
    icmptype?: pulumi.Input<number>;
    /**
     * Start and end of the IP range associated with service.
     */
    iprange?: pulumi.Input<string>;
    /**
     * Custom service name.
     */
    name?: pulumi.Input<string>;
    /**
     * Protocol type based on IANA numbers. Valid values: `TCP/UDP/SCTP`, `ICMP`, `ICMP6`, `IP`, `HTTP`, `FTP`, `CONNECT`, `SOCKS-TCP`, `SOCKS-UDP`, `ALL`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * IP protocol number.
     */
    protocolNumber?: pulumi.Input<number>;
    /**
     * Enable/disable web proxy service. Valid values: `enable`, `disable`.
     */
    proxy?: pulumi.Input<string>;
    /**
     * Multiple SCTP port ranges.
     */
    sctpPortrange?: pulumi.Input<string>;
    /**
     * Session TTL (300 - 604800, 0 = default).
     */
    sessionTtl?: pulumi.Input<number>;
    /**
     * Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
     */
    tcpHalfcloseTimer?: pulumi.Input<number>;
    /**
     * Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
     */
    tcpHalfopenTimer?: pulumi.Input<number>;
    /**
     * Multiple TCP port ranges.
     */
    tcpPortrange?: pulumi.Input<string>;
    /**
     * Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).
     */
    tcpRstTimer?: pulumi.Input<number>;
    /**
     * Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
     */
    tcpTimewaitTimer?: pulumi.Input<number>;
    /**
     * UDP half close timeout (0 - 86400 sec, 0 = default).
     */
    udpIdleTimer?: pulumi.Input<number>;
    /**
     * Multiple UDP port ranges.
     */
    udpPortrange?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable the visibility of the service on the GUI. Valid values: `enable`, `disable`.
     */
    visibility?: pulumi.Input<string>;
}
