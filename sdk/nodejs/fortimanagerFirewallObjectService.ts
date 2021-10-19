// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource supports Create/Read/Update/Delete firewall object service for FortiManager.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const test1 = new fortios.FortimanagerFirewallObjectService("test1", {
 *     category: "Email",
 *     comment: "test obj service",
 *     iprange: "1.1.1.1",
 *     protocol: "TCP/UDP/SCTP",
 *     sctpPortranges: ["100-200:150-250"],
 *     tcpPortranges: ["100-200:150-250"],
 *     udpPortranges: ["100-200:150-250"],
 * });
 * const test2 = new fortios.FortimanagerFirewallObjectService("test2", {
 *     category: "Web Access",
 *     comment: "test obj service",
 *     icmpCode: 3,
 *     icmpType: 2,
 *     protocol: "ICMP",
 * });
 * const test3 = new fortios.FortimanagerFirewallObjectService("test3", {
 *     category: "File Access",
 *     comment: "test obj service",
 *     protocol: "IP",
 *     protocolNumber: 4,
 * });
 * ```
 */
export class FortimanagerFirewallObjectService extends pulumi.CustomResource {
    /**
     * Get an existing FortimanagerFirewallObjectService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FortimanagerFirewallObjectServiceState, opts?: pulumi.CustomResourceOptions): FortimanagerFirewallObjectService {
        return new FortimanagerFirewallObjectService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/fortimanagerFirewallObjectService:FortimanagerFirewallObjectService';

    /**
     * Returns true if the given object is an instance of FortimanagerFirewallObjectService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FortimanagerFirewallObjectService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FortimanagerFirewallObjectService.__pulumiType;
    }

    /**
     * ADOM name. default is 'root'.
     */
    public readonly adom!: pulumi.Output<string | undefined>;
    /**
     * Service category. Enum: ["", "File Access", "Authentication", "Email", "General", "Network Services", "Remote Access", "Tunneling", "VoIP, Messaging & Other Applications", "Web Access", "Web Proxy"]
     */
    public readonly category!: pulumi.Output<string | undefined>;
    /**
     * Comments.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Fully qualified domain name.
     */
    public readonly fqdn!: pulumi.Output<string | undefined>;
    /**
     * ICMP Code.
     */
    public readonly icmpCode!: pulumi.Output<number | undefined>;
    /**
     * ICMP Type.
     */
    public readonly icmpType!: pulumi.Output<number | undefined>;
    /**
     * Start and end of the IP range associated with service. Ip or Ip range(eg: 1.1.1.1-1.1.1.100).
     */
    public readonly iprange!: pulumi.Output<string | undefined>;
    /**
     * Custom service name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Protocol type. Enum: ["TCP/UDP/SCTP", "ICMP", "ICMP6", "IP"]
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    /**
     * IP protocol number.
     */
    public readonly protocolNumber!: pulumi.Output<number | undefined>;
    public readonly proxy!: pulumi.Output<string | undefined>;
    /**
     * SCTP port ranges. e.g: ["dst-port-range:src-port-range"]
     */
    public readonly sctpPortranges!: pulumi.Output<string[] | undefined>;
    /**
     * TCP port ranges. e.g: ["dst-port-range:src-port-range"]
     */
    public readonly tcpPortranges!: pulumi.Output<string[] | undefined>;
    /**
     * UDP port ranges. e.g: ["dst-port-range:src-port-range"]
     */
    public readonly udpPortranges!: pulumi.Output<string[] | undefined>;

    /**
     * Create a FortimanagerFirewallObjectService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FortimanagerFirewallObjectServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FortimanagerFirewallObjectServiceArgs | FortimanagerFirewallObjectServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FortimanagerFirewallObjectServiceState | undefined;
            inputs["adom"] = state ? state.adom : undefined;
            inputs["category"] = state ? state.category : undefined;
            inputs["comment"] = state ? state.comment : undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["icmpCode"] = state ? state.icmpCode : undefined;
            inputs["icmpType"] = state ? state.icmpType : undefined;
            inputs["iprange"] = state ? state.iprange : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["protocolNumber"] = state ? state.protocolNumber : undefined;
            inputs["proxy"] = state ? state.proxy : undefined;
            inputs["sctpPortranges"] = state ? state.sctpPortranges : undefined;
            inputs["tcpPortranges"] = state ? state.tcpPortranges : undefined;
            inputs["udpPortranges"] = state ? state.udpPortranges : undefined;
        } else {
            const args = argsOrState as FortimanagerFirewallObjectServiceArgs | undefined;
            inputs["adom"] = args ? args.adom : undefined;
            inputs["category"] = args ? args.category : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["fqdn"] = args ? args.fqdn : undefined;
            inputs["icmpCode"] = args ? args.icmpCode : undefined;
            inputs["icmpType"] = args ? args.icmpType : undefined;
            inputs["iprange"] = args ? args.iprange : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["protocolNumber"] = args ? args.protocolNumber : undefined;
            inputs["proxy"] = args ? args.proxy : undefined;
            inputs["sctpPortranges"] = args ? args.sctpPortranges : undefined;
            inputs["tcpPortranges"] = args ? args.tcpPortranges : undefined;
            inputs["udpPortranges"] = args ? args.udpPortranges : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FortimanagerFirewallObjectService.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FortimanagerFirewallObjectService resources.
 */
export interface FortimanagerFirewallObjectServiceState {
    /**
     * ADOM name. default is 'root'.
     */
    adom?: pulumi.Input<string>;
    /**
     * Service category. Enum: ["", "File Access", "Authentication", "Email", "General", "Network Services", "Remote Access", "Tunneling", "VoIP, Messaging & Other Applications", "Web Access", "Web Proxy"]
     */
    category?: pulumi.Input<string>;
    /**
     * Comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Fully qualified domain name.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * ICMP Code.
     */
    icmpCode?: pulumi.Input<number>;
    /**
     * ICMP Type.
     */
    icmpType?: pulumi.Input<number>;
    /**
     * Start and end of the IP range associated with service. Ip or Ip range(eg: 1.1.1.1-1.1.1.100).
     */
    iprange?: pulumi.Input<string>;
    /**
     * Custom service name.
     */
    name?: pulumi.Input<string>;
    /**
     * Protocol type. Enum: ["TCP/UDP/SCTP", "ICMP", "ICMP6", "IP"]
     */
    protocol?: pulumi.Input<string>;
    /**
     * IP protocol number.
     */
    protocolNumber?: pulumi.Input<number>;
    proxy?: pulumi.Input<string>;
    /**
     * SCTP port ranges. e.g: ["dst-port-range:src-port-range"]
     */
    sctpPortranges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * TCP port ranges. e.g: ["dst-port-range:src-port-range"]
     */
    tcpPortranges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * UDP port ranges. e.g: ["dst-port-range:src-port-range"]
     */
    udpPortranges?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a FortimanagerFirewallObjectService resource.
 */
export interface FortimanagerFirewallObjectServiceArgs {
    /**
     * ADOM name. default is 'root'.
     */
    adom?: pulumi.Input<string>;
    /**
     * Service category. Enum: ["", "File Access", "Authentication", "Email", "General", "Network Services", "Remote Access", "Tunneling", "VoIP, Messaging & Other Applications", "Web Access", "Web Proxy"]
     */
    category?: pulumi.Input<string>;
    /**
     * Comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Fully qualified domain name.
     */
    fqdn?: pulumi.Input<string>;
    /**
     * ICMP Code.
     */
    icmpCode?: pulumi.Input<number>;
    /**
     * ICMP Type.
     */
    icmpType?: pulumi.Input<number>;
    /**
     * Start and end of the IP range associated with service. Ip or Ip range(eg: 1.1.1.1-1.1.1.100).
     */
    iprange?: pulumi.Input<string>;
    /**
     * Custom service name.
     */
    name?: pulumi.Input<string>;
    /**
     * Protocol type. Enum: ["TCP/UDP/SCTP", "ICMP", "ICMP6", "IP"]
     */
    protocol?: pulumi.Input<string>;
    /**
     * IP protocol number.
     */
    protocolNumber?: pulumi.Input<number>;
    proxy?: pulumi.Input<string>;
    /**
     * SCTP port ranges. e.g: ["dst-port-range:src-port-range"]
     */
    sctpPortranges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * TCP port ranges. e.g: ["dst-port-range:src-port-range"]
     */
    tcpPortranges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * UDP port ranges. e.g: ["dst-port-range:src-port-range"]
     */
    udpPortranges?: pulumi.Input<pulumi.Input<string>[]>;
}
