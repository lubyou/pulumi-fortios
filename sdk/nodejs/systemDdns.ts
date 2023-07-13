// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SystemDdns extends pulumi.CustomResource {
    /**
     * Get an existing SystemDdns resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemDdnsState, opts?: pulumi.CustomResourceOptions): SystemDdns {
        return new SystemDdns(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemDdns:SystemDdns';

    /**
     * Returns true if the given object is an instance of SystemDdns.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemDdns {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemDdns.__pulumiType;
    }

    public readonly addrType!: pulumi.Output<string>;
    public readonly boundIp!: pulumi.Output<string>;
    public readonly clearText!: pulumi.Output<string>;
    public readonly ddnsAuth!: pulumi.Output<string>;
    public readonly ddnsDomain!: pulumi.Output<string>;
    public readonly ddnsKey!: pulumi.Output<string>;
    public readonly ddnsKeyname!: pulumi.Output<string>;
    public readonly ddnsPassword!: pulumi.Output<string | undefined>;
    public readonly ddnsServer!: pulumi.Output<string>;
    public readonly ddnsServerAddrs!: pulumi.Output<outputs.SystemDdnsDdnsServerAddr[] | undefined>;
    public readonly ddnsServerIp!: pulumi.Output<string>;
    public readonly ddnsSn!: pulumi.Output<string>;
    public readonly ddnsTtl!: pulumi.Output<number>;
    public readonly ddnsUsername!: pulumi.Output<string>;
    public readonly ddnsZone!: pulumi.Output<string>;
    public readonly ddnsid!: pulumi.Output<number>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly monitorInterfaces!: pulumi.Output<outputs.SystemDdnsMonitorInterface[]>;
    public readonly serverType!: pulumi.Output<string>;
    public readonly sslCertificate!: pulumi.Output<string>;
    public readonly updateInterval!: pulumi.Output<number>;
    public readonly usePublicIp!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemDdns resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemDdnsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemDdnsArgs | SystemDdnsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemDdnsState | undefined;
            resourceInputs["addrType"] = state ? state.addrType : undefined;
            resourceInputs["boundIp"] = state ? state.boundIp : undefined;
            resourceInputs["clearText"] = state ? state.clearText : undefined;
            resourceInputs["ddnsAuth"] = state ? state.ddnsAuth : undefined;
            resourceInputs["ddnsDomain"] = state ? state.ddnsDomain : undefined;
            resourceInputs["ddnsKey"] = state ? state.ddnsKey : undefined;
            resourceInputs["ddnsKeyname"] = state ? state.ddnsKeyname : undefined;
            resourceInputs["ddnsPassword"] = state ? state.ddnsPassword : undefined;
            resourceInputs["ddnsServer"] = state ? state.ddnsServer : undefined;
            resourceInputs["ddnsServerAddrs"] = state ? state.ddnsServerAddrs : undefined;
            resourceInputs["ddnsServerIp"] = state ? state.ddnsServerIp : undefined;
            resourceInputs["ddnsSn"] = state ? state.ddnsSn : undefined;
            resourceInputs["ddnsTtl"] = state ? state.ddnsTtl : undefined;
            resourceInputs["ddnsUsername"] = state ? state.ddnsUsername : undefined;
            resourceInputs["ddnsZone"] = state ? state.ddnsZone : undefined;
            resourceInputs["ddnsid"] = state ? state.ddnsid : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["monitorInterfaces"] = state ? state.monitorInterfaces : undefined;
            resourceInputs["serverType"] = state ? state.serverType : undefined;
            resourceInputs["sslCertificate"] = state ? state.sslCertificate : undefined;
            resourceInputs["updateInterval"] = state ? state.updateInterval : undefined;
            resourceInputs["usePublicIp"] = state ? state.usePublicIp : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemDdnsArgs | undefined;
            if ((!args || args.ddnsServer === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ddnsServer'");
            }
            if ((!args || args.monitorInterfaces === undefined) && !opts.urn) {
                throw new Error("Missing required property 'monitorInterfaces'");
            }
            resourceInputs["addrType"] = args ? args.addrType : undefined;
            resourceInputs["boundIp"] = args ? args.boundIp : undefined;
            resourceInputs["clearText"] = args ? args.clearText : undefined;
            resourceInputs["ddnsAuth"] = args ? args.ddnsAuth : undefined;
            resourceInputs["ddnsDomain"] = args ? args.ddnsDomain : undefined;
            resourceInputs["ddnsKey"] = args?.ddnsKey ? pulumi.secret(args.ddnsKey) : undefined;
            resourceInputs["ddnsKeyname"] = args ? args.ddnsKeyname : undefined;
            resourceInputs["ddnsPassword"] = args?.ddnsPassword ? pulumi.secret(args.ddnsPassword) : undefined;
            resourceInputs["ddnsServer"] = args ? args.ddnsServer : undefined;
            resourceInputs["ddnsServerAddrs"] = args ? args.ddnsServerAddrs : undefined;
            resourceInputs["ddnsServerIp"] = args ? args.ddnsServerIp : undefined;
            resourceInputs["ddnsSn"] = args ? args.ddnsSn : undefined;
            resourceInputs["ddnsTtl"] = args ? args.ddnsTtl : undefined;
            resourceInputs["ddnsUsername"] = args ? args.ddnsUsername : undefined;
            resourceInputs["ddnsZone"] = args ? args.ddnsZone : undefined;
            resourceInputs["ddnsid"] = args ? args.ddnsid : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["monitorInterfaces"] = args ? args.monitorInterfaces : undefined;
            resourceInputs["serverType"] = args ? args.serverType : undefined;
            resourceInputs["sslCertificate"] = args ? args.sslCertificate : undefined;
            resourceInputs["updateInterval"] = args ? args.updateInterval : undefined;
            resourceInputs["usePublicIp"] = args ? args.usePublicIp : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["ddnsKey", "ddnsPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SystemDdns.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemDdns resources.
 */
export interface SystemDdnsState {
    addrType?: pulumi.Input<string>;
    boundIp?: pulumi.Input<string>;
    clearText?: pulumi.Input<string>;
    ddnsAuth?: pulumi.Input<string>;
    ddnsDomain?: pulumi.Input<string>;
    ddnsKey?: pulumi.Input<string>;
    ddnsKeyname?: pulumi.Input<string>;
    ddnsPassword?: pulumi.Input<string>;
    ddnsServer?: pulumi.Input<string>;
    ddnsServerAddrs?: pulumi.Input<pulumi.Input<inputs.SystemDdnsDdnsServerAddr>[]>;
    ddnsServerIp?: pulumi.Input<string>;
    ddnsSn?: pulumi.Input<string>;
    ddnsTtl?: pulumi.Input<number>;
    ddnsUsername?: pulumi.Input<string>;
    ddnsZone?: pulumi.Input<string>;
    ddnsid?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    monitorInterfaces?: pulumi.Input<pulumi.Input<inputs.SystemDdnsMonitorInterface>[]>;
    serverType?: pulumi.Input<string>;
    sslCertificate?: pulumi.Input<string>;
    updateInterval?: pulumi.Input<number>;
    usePublicIp?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemDdns resource.
 */
export interface SystemDdnsArgs {
    addrType?: pulumi.Input<string>;
    boundIp?: pulumi.Input<string>;
    clearText?: pulumi.Input<string>;
    ddnsAuth?: pulumi.Input<string>;
    ddnsDomain?: pulumi.Input<string>;
    ddnsKey?: pulumi.Input<string>;
    ddnsKeyname?: pulumi.Input<string>;
    ddnsPassword?: pulumi.Input<string>;
    ddnsServer: pulumi.Input<string>;
    ddnsServerAddrs?: pulumi.Input<pulumi.Input<inputs.SystemDdnsDdnsServerAddr>[]>;
    ddnsServerIp?: pulumi.Input<string>;
    ddnsSn?: pulumi.Input<string>;
    ddnsTtl?: pulumi.Input<number>;
    ddnsUsername?: pulumi.Input<string>;
    ddnsZone?: pulumi.Input<string>;
    ddnsid?: pulumi.Input<number>;
    dynamicSortSubtable?: pulumi.Input<string>;
    getAllTables?: pulumi.Input<string>;
    monitorInterfaces: pulumi.Input<pulumi.Input<inputs.SystemDdnsMonitorInterface>[]>;
    serverType?: pulumi.Input<string>;
    sslCertificate?: pulumi.Input<string>;
    updateInterval?: pulumi.Input<number>;
    usePublicIp?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
