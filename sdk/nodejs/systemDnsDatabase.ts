// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure DNS databases.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.SystemDnsDatabase("trname", {
 *     authoritative: "enable",
 *     contact: "hostmaster",
 *     dnsEntries: [{
 *         hostname: "sghsgh.com",
 *         ttl: 3,
 *         type: "MX",
 *     }],
 *     domain: "s.com",
 *     forwarder: "\"9.9.9.9\" \"3.3.3.3\" ",
 *     ipMaster: "0.0.0.0",
 *     primaryName: "dns",
 *     sourceIp: "0.0.0.0",
 *     status: "enable",
 *     ttl: 86400,
 *     type: "master",
 *     view: "shadow",
 * });
 * ```
 *
 * ## Import
 *
 * System DnsDatabase can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemDnsDatabase:SystemDnsDatabase labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemDnsDatabase extends pulumi.CustomResource {
    /**
     * Get an existing SystemDnsDatabase resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemDnsDatabaseState, opts?: pulumi.CustomResourceOptions): SystemDnsDatabase {
        return new SystemDnsDatabase(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemDnsDatabase:SystemDnsDatabase';

    /**
     * Returns true if the given object is an instance of SystemDnsDatabase.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemDnsDatabase {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemDnsDatabase.__pulumiType;
    }

    /**
     * DNS zone transfer IP address list.
     */
    public readonly allowTransfer!: pulumi.Output<string>;
    /**
     * Enable/disable authoritative zone. Valid values: `enable`, `disable`.
     */
    public readonly authoritative!: pulumi.Output<string>;
    /**
     * Email address of the administrator for this zone.
     * You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com)
     * When using a simple username, the domain of the email will be this zone.
     */
    public readonly contact!: pulumi.Output<string>;
    /**
     * DNS entry. The structure of `dnsEntry` block is documented below.
     */
    public readonly dnsEntries!: pulumi.Output<outputs.SystemDnsDatabaseDnsEntry[] | undefined>;
    /**
     * Domain name.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * DNS zone forwarder IP address list.
     */
    public readonly forwarder!: pulumi.Output<string>;
    /**
     * IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
     */
    public readonly ipMaster!: pulumi.Output<string>;
    /**
     * IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
     */
    public readonly ipPrimary!: pulumi.Output<string>;
    /**
     * Zone name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Domain name of the default DNS server for this zone.
     */
    public readonly primaryName!: pulumi.Output<string>;
    /**
     * Maximum number of resource records (10 - 65536, 0 means infinite).
     */
    public readonly rrMax!: pulumi.Output<number>;
    /**
     * Source IP for forwarding to DNS server.
     */
    public readonly sourceIp!: pulumi.Output<string>;
    /**
     * Enable/disable resource record status. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Time-to-live for this entry (0 to 2147483647 sec, default = 0).
     */
    public readonly ttl!: pulumi.Output<number>;
    /**
     * Resource record type. Valid values: `A`, `NS`, `CNAME`, `MX`, `AAAA`, `PTR`, `PTR_V6`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Zone view (public to serve public clients, shadow to serve internal clients). Valid values: `shadow`, `public`.
     */
    public readonly view!: pulumi.Output<string>;

    /**
     * Create a SystemDnsDatabase resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemDnsDatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemDnsDatabaseArgs | SystemDnsDatabaseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemDnsDatabaseState | undefined;
            inputs["allowTransfer"] = state ? state.allowTransfer : undefined;
            inputs["authoritative"] = state ? state.authoritative : undefined;
            inputs["contact"] = state ? state.contact : undefined;
            inputs["dnsEntries"] = state ? state.dnsEntries : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            inputs["forwarder"] = state ? state.forwarder : undefined;
            inputs["ipMaster"] = state ? state.ipMaster : undefined;
            inputs["ipPrimary"] = state ? state.ipPrimary : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["primaryName"] = state ? state.primaryName : undefined;
            inputs["rrMax"] = state ? state.rrMax : undefined;
            inputs["sourceIp"] = state ? state.sourceIp : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["vdomparam"] = state ? state.vdomparam : undefined;
            inputs["view"] = state ? state.view : undefined;
        } else {
            const args = argsOrState as SystemDnsDatabaseArgs | undefined;
            if ((!args || args.authoritative === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authoritative'");
            }
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.ttl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ttl'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.view === undefined) && !opts.urn) {
                throw new Error("Missing required property 'view'");
            }
            inputs["allowTransfer"] = args ? args.allowTransfer : undefined;
            inputs["authoritative"] = args ? args.authoritative : undefined;
            inputs["contact"] = args ? args.contact : undefined;
            inputs["dnsEntries"] = args ? args.dnsEntries : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            inputs["forwarder"] = args ? args.forwarder : undefined;
            inputs["ipMaster"] = args ? args.ipMaster : undefined;
            inputs["ipPrimary"] = args ? args.ipPrimary : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["primaryName"] = args ? args.primaryName : undefined;
            inputs["rrMax"] = args ? args.rrMax : undefined;
            inputs["sourceIp"] = args ? args.sourceIp : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["vdomparam"] = args ? args.vdomparam : undefined;
            inputs["view"] = args ? args.view : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SystemDnsDatabase.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemDnsDatabase resources.
 */
export interface SystemDnsDatabaseState {
    /**
     * DNS zone transfer IP address list.
     */
    allowTransfer?: pulumi.Input<string>;
    /**
     * Enable/disable authoritative zone. Valid values: `enable`, `disable`.
     */
    authoritative?: pulumi.Input<string>;
    /**
     * Email address of the administrator for this zone.
     * You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com)
     * When using a simple username, the domain of the email will be this zone.
     */
    contact?: pulumi.Input<string>;
    /**
     * DNS entry. The structure of `dnsEntry` block is documented below.
     */
    dnsEntries?: pulumi.Input<pulumi.Input<inputs.SystemDnsDatabaseDnsEntry>[]>;
    /**
     * Domain name.
     */
    domain?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * DNS zone forwarder IP address list.
     */
    forwarder?: pulumi.Input<string>;
    /**
     * IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
     */
    ipMaster?: pulumi.Input<string>;
    /**
     * IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
     */
    ipPrimary?: pulumi.Input<string>;
    /**
     * Zone name.
     */
    name?: pulumi.Input<string>;
    /**
     * Domain name of the default DNS server for this zone.
     */
    primaryName?: pulumi.Input<string>;
    /**
     * Maximum number of resource records (10 - 65536, 0 means infinite).
     */
    rrMax?: pulumi.Input<number>;
    /**
     * Source IP for forwarding to DNS server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Enable/disable resource record status. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Time-to-live for this entry (0 to 2147483647 sec, default = 0).
     */
    ttl?: pulumi.Input<number>;
    /**
     * Resource record type. Valid values: `A`, `NS`, `CNAME`, `MX`, `AAAA`, `PTR`, `PTR_V6`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Zone view (public to serve public clients, shadow to serve internal clients). Valid values: `shadow`, `public`.
     */
    view?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemDnsDatabase resource.
 */
export interface SystemDnsDatabaseArgs {
    /**
     * DNS zone transfer IP address list.
     */
    allowTransfer?: pulumi.Input<string>;
    /**
     * Enable/disable authoritative zone. Valid values: `enable`, `disable`.
     */
    authoritative: pulumi.Input<string>;
    /**
     * Email address of the administrator for this zone.
     * You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com)
     * When using a simple username, the domain of the email will be this zone.
     */
    contact?: pulumi.Input<string>;
    /**
     * DNS entry. The structure of `dnsEntry` block is documented below.
     */
    dnsEntries?: pulumi.Input<pulumi.Input<inputs.SystemDnsDatabaseDnsEntry>[]>;
    /**
     * Domain name.
     */
    domain: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * DNS zone forwarder IP address list.
     */
    forwarder?: pulumi.Input<string>;
    /**
     * IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
     */
    ipMaster?: pulumi.Input<string>;
    /**
     * IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
     */
    ipPrimary?: pulumi.Input<string>;
    /**
     * Zone name.
     */
    name?: pulumi.Input<string>;
    /**
     * Domain name of the default DNS server for this zone.
     */
    primaryName?: pulumi.Input<string>;
    /**
     * Maximum number of resource records (10 - 65536, 0 means infinite).
     */
    rrMax?: pulumi.Input<number>;
    /**
     * Source IP for forwarding to DNS server.
     */
    sourceIp?: pulumi.Input<string>;
    /**
     * Enable/disable resource record status. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Time-to-live for this entry (0 to 2147483647 sec, default = 0).
     */
    ttl: pulumi.Input<number>;
    /**
     * Resource record type. Valid values: `A`, `NS`, `CNAME`, `MX`, `AAAA`, `PTR`, `PTR_V6`.
     */
    type: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Zone view (public to serve public clients, shadow to serve internal clients). Valid values: `shadow`, `public`.
     */
    view: pulumi.Input<string>;
}
