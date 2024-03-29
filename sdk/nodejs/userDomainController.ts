// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class UserDomainController extends pulumi.CustomResource {
    /**
     * Get an existing UserDomainController resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserDomainControllerState, opts?: pulumi.CustomResourceOptions): UserDomainController {
        return new UserDomainController(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/userDomainController:UserDomainController';

    /**
     * Returns true if the given object is an instance of UserDomainController.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserDomainController {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserDomainController.__pulumiType;
    }

    public readonly adMode!: pulumi.Output<string>;
    public readonly adldsDn!: pulumi.Output<string>;
    public readonly adldsIp6!: pulumi.Output<string>;
    public readonly adldsIpAddress!: pulumi.Output<string>;
    public readonly adldsPort!: pulumi.Output<number>;
    public readonly changeDetection!: pulumi.Output<string>;
    public readonly changeDetectionPeriod!: pulumi.Output<number>;
    public readonly dnsSrvLookup!: pulumi.Output<string>;
    public readonly domainName!: pulumi.Output<string>;
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    public readonly extraServers!: pulumi.Output<outputs.UserDomainControllerExtraServer[] | undefined>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly hostname!: pulumi.Output<string>;
    public readonly interface!: pulumi.Output<string>;
    public readonly interfaceSelectMethod!: pulumi.Output<string>;
    public readonly ip6!: pulumi.Output<string>;
    public readonly ipAddress!: pulumi.Output<string>;
    public readonly ldapServer!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly port!: pulumi.Output<number>;
    public readonly replicationPort!: pulumi.Output<number>;
    public readonly sourceIp6!: pulumi.Output<string>;
    public readonly sourceIpAddress!: pulumi.Output<string>;
    public readonly sourcePort!: pulumi.Output<number>;
    public readonly username!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a UserDomainController resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserDomainControllerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserDomainControllerArgs | UserDomainControllerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserDomainControllerState | undefined;
            resourceInputs["adMode"] = state ? state.adMode : undefined;
            resourceInputs["adldsDn"] = state ? state.adldsDn : undefined;
            resourceInputs["adldsIp6"] = state ? state.adldsIp6 : undefined;
            resourceInputs["adldsIpAddress"] = state ? state.adldsIpAddress : undefined;
            resourceInputs["adldsPort"] = state ? state.adldsPort : undefined;
            resourceInputs["changeDetection"] = state ? state.changeDetection : undefined;
            resourceInputs["changeDetectionPeriod"] = state ? state.changeDetectionPeriod : undefined;
            resourceInputs["dnsSrvLookup"] = state ? state.dnsSrvLookup : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["extraServers"] = state ? state.extraServers : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = state ? state.interfaceSelectMethod : undefined;
            resourceInputs["ip6"] = state ? state.ip6 : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["ldapServer"] = state ? state.ldapServer : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["replicationPort"] = state ? state.replicationPort : undefined;
            resourceInputs["sourceIp6"] = state ? state.sourceIp6 : undefined;
            resourceInputs["sourceIpAddress"] = state ? state.sourceIpAddress : undefined;
            resourceInputs["sourcePort"] = state ? state.sourcePort : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as UserDomainControllerArgs | undefined;
            if ((!args || args.ipAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipAddress'");
            }
            if ((!args || args.ldapServer === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ldapServer'");
            }
            resourceInputs["adMode"] = args ? args.adMode : undefined;
            resourceInputs["adldsDn"] = args ? args.adldsDn : undefined;
            resourceInputs["adldsIp6"] = args ? args.adldsIp6 : undefined;
            resourceInputs["adldsIpAddress"] = args ? args.adldsIpAddress : undefined;
            resourceInputs["adldsPort"] = args ? args.adldsPort : undefined;
            resourceInputs["changeDetection"] = args ? args.changeDetection : undefined;
            resourceInputs["changeDetectionPeriod"] = args ? args.changeDetectionPeriod : undefined;
            resourceInputs["dnsSrvLookup"] = args ? args.dnsSrvLookup : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["extraServers"] = args ? args.extraServers : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["interfaceSelectMethod"] = args ? args.interfaceSelectMethod : undefined;
            resourceInputs["ip6"] = args ? args.ip6 : undefined;
            resourceInputs["ipAddress"] = args ? args.ipAddress : undefined;
            resourceInputs["ldapServer"] = args ? args.ldapServer : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["replicationPort"] = args ? args.replicationPort : undefined;
            resourceInputs["sourceIp6"] = args ? args.sourceIp6 : undefined;
            resourceInputs["sourceIpAddress"] = args ? args.sourceIpAddress : undefined;
            resourceInputs["sourcePort"] = args ? args.sourcePort : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserDomainController.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserDomainController resources.
 */
export interface UserDomainControllerState {
    adMode?: pulumi.Input<string>;
    adldsDn?: pulumi.Input<string>;
    adldsIp6?: pulumi.Input<string>;
    adldsIpAddress?: pulumi.Input<string>;
    adldsPort?: pulumi.Input<number>;
    changeDetection?: pulumi.Input<string>;
    changeDetectionPeriod?: pulumi.Input<number>;
    dnsSrvLookup?: pulumi.Input<string>;
    domainName?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    extraServers?: pulumi.Input<pulumi.Input<inputs.UserDomainControllerExtraServer>[]>;
    getAllTables?: pulumi.Input<string>;
    hostname?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    ip6?: pulumi.Input<string>;
    ipAddress?: pulumi.Input<string>;
    ldapServer?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    replicationPort?: pulumi.Input<number>;
    sourceIp6?: pulumi.Input<string>;
    sourceIpAddress?: pulumi.Input<string>;
    sourcePort?: pulumi.Input<number>;
    username?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserDomainController resource.
 */
export interface UserDomainControllerArgs {
    adMode?: pulumi.Input<string>;
    adldsDn?: pulumi.Input<string>;
    adldsIp6?: pulumi.Input<string>;
    adldsIpAddress?: pulumi.Input<string>;
    adldsPort?: pulumi.Input<number>;
    changeDetection?: pulumi.Input<string>;
    changeDetectionPeriod?: pulumi.Input<number>;
    dnsSrvLookup?: pulumi.Input<string>;
    domainName?: pulumi.Input<string>;
    dynamicSortSubtable?: pulumi.Input<string>;
    extraServers?: pulumi.Input<pulumi.Input<inputs.UserDomainControllerExtraServer>[]>;
    getAllTables?: pulumi.Input<string>;
    hostname?: pulumi.Input<string>;
    interface?: pulumi.Input<string>;
    interfaceSelectMethod?: pulumi.Input<string>;
    ip6?: pulumi.Input<string>;
    ipAddress: pulumi.Input<string>;
    ldapServer: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    replicationPort?: pulumi.Input<number>;
    sourceIp6?: pulumi.Input<string>;
    sourceIpAddress?: pulumi.Input<string>;
    sourcePort?: pulumi.Input<number>;
    username?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
