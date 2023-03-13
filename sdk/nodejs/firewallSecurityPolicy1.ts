// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class FirewallSecurityPolicy1 extends pulumi.CustomResource {
    /**
     * Get an existing FirewallSecurityPolicy1 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallSecurityPolicy1State, opts?: pulumi.CustomResourceOptions): FirewallSecurityPolicy1 {
        return new FirewallSecurityPolicy1(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallSecurityPolicy1:FirewallSecurityPolicy1';

    /**
     * Returns true if the given object is an instance of FirewallSecurityPolicy1.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallSecurityPolicy1 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallSecurityPolicy1.__pulumiType;
    }

    public readonly action!: pulumi.Output<string>;
    public readonly applicationList!: pulumi.Output<string>;
    public readonly avProfile!: pulumi.Output<string>;
    public readonly capturePacket!: pulumi.Output<string>;
    public readonly comments!: pulumi.Output<string | undefined>;
    public readonly devices!: pulumi.Output<string[] | undefined>;
    public readonly dnsfilterProfile!: pulumi.Output<string>;
    public readonly dstaddrs!: pulumi.Output<string[]>;
    public readonly dstintfs!: pulumi.Output<string[]>;
    public readonly groups!: pulumi.Output<string[] | undefined>;
    public readonly internetService!: pulumi.Output<string>;
    public readonly internetServiceIds!: pulumi.Output<number[] | undefined>;
    public readonly internetServiceSrc!: pulumi.Output<string>;
    public readonly internetServiceSrcIds!: pulumi.Output<number[] | undefined>;
    public readonly ippool!: pulumi.Output<string>;
    public readonly ipsSensor!: pulumi.Output<string>;
    public readonly logtraffic!: pulumi.Output<string>;
    public readonly logtrafficStart!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly nat!: pulumi.Output<string>;
    public readonly poolnames!: pulumi.Output<string[] | undefined>;
    public readonly profileProtocolOptions!: pulumi.Output<string>;
    public readonly schedule!: pulumi.Output<string>;
    public readonly services!: pulumi.Output<string[]>;
    public readonly srcaddrs!: pulumi.Output<string[]>;
    public readonly srcintfs!: pulumi.Output<string[]>;
    public readonly sslSshProfile!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string>;
    public readonly users!: pulumi.Output<string[] | undefined>;
    public readonly utmStatus!: pulumi.Output<string>;
    public readonly webfilterProfile!: pulumi.Output<string>;

    /**
     * Create a FirewallSecurityPolicy1 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallSecurityPolicy1Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallSecurityPolicy1Args | FirewallSecurityPolicy1State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallSecurityPolicy1State | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["applicationList"] = state ? state.applicationList : undefined;
            resourceInputs["avProfile"] = state ? state.avProfile : undefined;
            resourceInputs["capturePacket"] = state ? state.capturePacket : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["devices"] = state ? state.devices : undefined;
            resourceInputs["dnsfilterProfile"] = state ? state.dnsfilterProfile : undefined;
            resourceInputs["dstaddrs"] = state ? state.dstaddrs : undefined;
            resourceInputs["dstintfs"] = state ? state.dstintfs : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["internetService"] = state ? state.internetService : undefined;
            resourceInputs["internetServiceIds"] = state ? state.internetServiceIds : undefined;
            resourceInputs["internetServiceSrc"] = state ? state.internetServiceSrc : undefined;
            resourceInputs["internetServiceSrcIds"] = state ? state.internetServiceSrcIds : undefined;
            resourceInputs["ippool"] = state ? state.ippool : undefined;
            resourceInputs["ipsSensor"] = state ? state.ipsSensor : undefined;
            resourceInputs["logtraffic"] = state ? state.logtraffic : undefined;
            resourceInputs["logtrafficStart"] = state ? state.logtrafficStart : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nat"] = state ? state.nat : undefined;
            resourceInputs["poolnames"] = state ? state.poolnames : undefined;
            resourceInputs["profileProtocolOptions"] = state ? state.profileProtocolOptions : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["srcaddrs"] = state ? state.srcaddrs : undefined;
            resourceInputs["srcintfs"] = state ? state.srcintfs : undefined;
            resourceInputs["sslSshProfile"] = state ? state.sslSshProfile : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["users"] = state ? state.users : undefined;
            resourceInputs["utmStatus"] = state ? state.utmStatus : undefined;
            resourceInputs["webfilterProfile"] = state ? state.webfilterProfile : undefined;
        } else {
            const args = argsOrState as FirewallSecurityPolicy1Args | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.dstaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstaddrs'");
            }
            if ((!args || args.dstintfs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dstintfs'");
            }
            if ((!args || args.schedule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schedule'");
            }
            if ((!args || args.services === undefined) && !opts.urn) {
                throw new Error("Missing required property 'services'");
            }
            if ((!args || args.srcaddrs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcaddrs'");
            }
            if ((!args || args.srcintfs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcintfs'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["applicationList"] = args ? args.applicationList : undefined;
            resourceInputs["avProfile"] = args ? args.avProfile : undefined;
            resourceInputs["capturePacket"] = args ? args.capturePacket : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["devices"] = args ? args.devices : undefined;
            resourceInputs["dnsfilterProfile"] = args ? args.dnsfilterProfile : undefined;
            resourceInputs["dstaddrs"] = args ? args.dstaddrs : undefined;
            resourceInputs["dstintfs"] = args ? args.dstintfs : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["internetService"] = args ? args.internetService : undefined;
            resourceInputs["internetServiceIds"] = args ? args.internetServiceIds : undefined;
            resourceInputs["internetServiceSrc"] = args ? args.internetServiceSrc : undefined;
            resourceInputs["internetServiceSrcIds"] = args ? args.internetServiceSrcIds : undefined;
            resourceInputs["ippool"] = args ? args.ippool : undefined;
            resourceInputs["ipsSensor"] = args ? args.ipsSensor : undefined;
            resourceInputs["logtraffic"] = args ? args.logtraffic : undefined;
            resourceInputs["logtrafficStart"] = args ? args.logtrafficStart : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nat"] = args ? args.nat : undefined;
            resourceInputs["poolnames"] = args ? args.poolnames : undefined;
            resourceInputs["profileProtocolOptions"] = args ? args.profileProtocolOptions : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["srcaddrs"] = args ? args.srcaddrs : undefined;
            resourceInputs["srcintfs"] = args ? args.srcintfs : undefined;
            resourceInputs["sslSshProfile"] = args ? args.sslSshProfile : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["users"] = args ? args.users : undefined;
            resourceInputs["utmStatus"] = args ? args.utmStatus : undefined;
            resourceInputs["webfilterProfile"] = args ? args.webfilterProfile : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallSecurityPolicy1.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallSecurityPolicy1 resources.
 */
export interface FirewallSecurityPolicy1State {
    action?: pulumi.Input<string>;
    applicationList?: pulumi.Input<string>;
    avProfile?: pulumi.Input<string>;
    capturePacket?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    devices?: pulumi.Input<pulumi.Input<string>[]>;
    dnsfilterProfile?: pulumi.Input<string>;
    dstaddrs?: pulumi.Input<pulumi.Input<string>[]>;
    dstintfs?: pulumi.Input<pulumi.Input<string>[]>;
    groups?: pulumi.Input<pulumi.Input<string>[]>;
    internetService?: pulumi.Input<string>;
    internetServiceIds?: pulumi.Input<pulumi.Input<number>[]>;
    internetServiceSrc?: pulumi.Input<string>;
    internetServiceSrcIds?: pulumi.Input<pulumi.Input<number>[]>;
    ippool?: pulumi.Input<string>;
    ipsSensor?: pulumi.Input<string>;
    logtraffic?: pulumi.Input<string>;
    logtrafficStart?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    nat?: pulumi.Input<string>;
    poolnames?: pulumi.Input<pulumi.Input<string>[]>;
    profileProtocolOptions?: pulumi.Input<string>;
    schedule?: pulumi.Input<string>;
    services?: pulumi.Input<pulumi.Input<string>[]>;
    srcaddrs?: pulumi.Input<pulumi.Input<string>[]>;
    srcintfs?: pulumi.Input<pulumi.Input<string>[]>;
    sslSshProfile?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    users?: pulumi.Input<pulumi.Input<string>[]>;
    utmStatus?: pulumi.Input<string>;
    webfilterProfile?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallSecurityPolicy1 resource.
 */
export interface FirewallSecurityPolicy1Args {
    action: pulumi.Input<string>;
    applicationList?: pulumi.Input<string>;
    avProfile?: pulumi.Input<string>;
    capturePacket?: pulumi.Input<string>;
    comments?: pulumi.Input<string>;
    devices?: pulumi.Input<pulumi.Input<string>[]>;
    dnsfilterProfile?: pulumi.Input<string>;
    dstaddrs: pulumi.Input<pulumi.Input<string>[]>;
    dstintfs: pulumi.Input<pulumi.Input<string>[]>;
    groups?: pulumi.Input<pulumi.Input<string>[]>;
    internetService?: pulumi.Input<string>;
    internetServiceIds?: pulumi.Input<pulumi.Input<number>[]>;
    internetServiceSrc?: pulumi.Input<string>;
    internetServiceSrcIds?: pulumi.Input<pulumi.Input<number>[]>;
    ippool?: pulumi.Input<string>;
    ipsSensor?: pulumi.Input<string>;
    logtraffic?: pulumi.Input<string>;
    logtrafficStart?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    nat?: pulumi.Input<string>;
    poolnames?: pulumi.Input<pulumi.Input<string>[]>;
    profileProtocolOptions?: pulumi.Input<string>;
    schedule: pulumi.Input<string>;
    services: pulumi.Input<pulumi.Input<string>[]>;
    srcaddrs: pulumi.Input<pulumi.Input<string>[]>;
    srcintfs: pulumi.Input<pulumi.Input<string>[]>;
    sslSshProfile?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    users?: pulumi.Input<pulumi.Input<string>[]>;
    utmStatus?: pulumi.Input<string>;
    webfilterProfile?: pulumi.Input<string>;
}
