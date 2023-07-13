// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class FirewallProfileProtocolOptions extends pulumi.CustomResource {
    /**
     * Get an existing FirewallProfileProtocolOptions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallProfileProtocolOptionsState, opts?: pulumi.CustomResourceOptions): FirewallProfileProtocolOptions {
        return new FirewallProfileProtocolOptions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallProfileProtocolOptions:FirewallProfileProtocolOptions';

    /**
     * Returns true if the given object is an instance of FirewallProfileProtocolOptions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallProfileProtocolOptions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallProfileProtocolOptions.__pulumiType;
    }

    public readonly cifs!: pulumi.Output<outputs.FirewallProfileProtocolOptionsCifs>;
    public readonly comment!: pulumi.Output<string | undefined>;
    public readonly dns!: pulumi.Output<outputs.FirewallProfileProtocolOptionsDns>;
    public readonly featureSet!: pulumi.Output<string>;
    public readonly ftp!: pulumi.Output<outputs.FirewallProfileProtocolOptionsFtp>;
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    public readonly http!: pulumi.Output<outputs.FirewallProfileProtocolOptionsHttp>;
    public readonly imap!: pulumi.Output<outputs.FirewallProfileProtocolOptionsImap>;
    public readonly mailSignature!: pulumi.Output<outputs.FirewallProfileProtocolOptionsMailSignature>;
    public readonly mapi!: pulumi.Output<outputs.FirewallProfileProtocolOptionsMapi>;
    public readonly name!: pulumi.Output<string>;
    public readonly nntp!: pulumi.Output<outputs.FirewallProfileProtocolOptionsNntp>;
    public readonly oversizeLog!: pulumi.Output<string>;
    public readonly pop3!: pulumi.Output<outputs.FirewallProfileProtocolOptionsPop3>;
    public readonly replacemsgGroup!: pulumi.Output<string>;
    public readonly rpcOverHttp!: pulumi.Output<string>;
    public readonly smtp!: pulumi.Output<outputs.FirewallProfileProtocolOptionsSmtp>;
    public readonly ssh!: pulumi.Output<outputs.FirewallProfileProtocolOptionsSsh>;
    public readonly switchingProtocolsLog!: pulumi.Output<string>;
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallProfileProtocolOptions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallProfileProtocolOptionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallProfileProtocolOptionsArgs | FirewallProfileProtocolOptionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallProfileProtocolOptionsState | undefined;
            resourceInputs["cifs"] = state ? state.cifs : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dns"] = state ? state.dns : undefined;
            resourceInputs["featureSet"] = state ? state.featureSet : undefined;
            resourceInputs["ftp"] = state ? state.ftp : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["http"] = state ? state.http : undefined;
            resourceInputs["imap"] = state ? state.imap : undefined;
            resourceInputs["mailSignature"] = state ? state.mailSignature : undefined;
            resourceInputs["mapi"] = state ? state.mapi : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nntp"] = state ? state.nntp : undefined;
            resourceInputs["oversizeLog"] = state ? state.oversizeLog : undefined;
            resourceInputs["pop3"] = state ? state.pop3 : undefined;
            resourceInputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            resourceInputs["rpcOverHttp"] = state ? state.rpcOverHttp : undefined;
            resourceInputs["smtp"] = state ? state.smtp : undefined;
            resourceInputs["ssh"] = state ? state.ssh : undefined;
            resourceInputs["switchingProtocolsLog"] = state ? state.switchingProtocolsLog : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallProfileProtocolOptionsArgs | undefined;
            resourceInputs["cifs"] = args ? args.cifs : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dns"] = args ? args.dns : undefined;
            resourceInputs["featureSet"] = args ? args.featureSet : undefined;
            resourceInputs["ftp"] = args ? args.ftp : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["http"] = args ? args.http : undefined;
            resourceInputs["imap"] = args ? args.imap : undefined;
            resourceInputs["mailSignature"] = args ? args.mailSignature : undefined;
            resourceInputs["mapi"] = args ? args.mapi : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nntp"] = args ? args.nntp : undefined;
            resourceInputs["oversizeLog"] = args ? args.oversizeLog : undefined;
            resourceInputs["pop3"] = args ? args.pop3 : undefined;
            resourceInputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            resourceInputs["rpcOverHttp"] = args ? args.rpcOverHttp : undefined;
            resourceInputs["smtp"] = args ? args.smtp : undefined;
            resourceInputs["ssh"] = args ? args.ssh : undefined;
            resourceInputs["switchingProtocolsLog"] = args ? args.switchingProtocolsLog : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallProfileProtocolOptions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallProfileProtocolOptions resources.
 */
export interface FirewallProfileProtocolOptionsState {
    cifs?: pulumi.Input<inputs.FirewallProfileProtocolOptionsCifs>;
    comment?: pulumi.Input<string>;
    dns?: pulumi.Input<inputs.FirewallProfileProtocolOptionsDns>;
    featureSet?: pulumi.Input<string>;
    ftp?: pulumi.Input<inputs.FirewallProfileProtocolOptionsFtp>;
    getAllTables?: pulumi.Input<string>;
    http?: pulumi.Input<inputs.FirewallProfileProtocolOptionsHttp>;
    imap?: pulumi.Input<inputs.FirewallProfileProtocolOptionsImap>;
    mailSignature?: pulumi.Input<inputs.FirewallProfileProtocolOptionsMailSignature>;
    mapi?: pulumi.Input<inputs.FirewallProfileProtocolOptionsMapi>;
    name?: pulumi.Input<string>;
    nntp?: pulumi.Input<inputs.FirewallProfileProtocolOptionsNntp>;
    oversizeLog?: pulumi.Input<string>;
    pop3?: pulumi.Input<inputs.FirewallProfileProtocolOptionsPop3>;
    replacemsgGroup?: pulumi.Input<string>;
    rpcOverHttp?: pulumi.Input<string>;
    smtp?: pulumi.Input<inputs.FirewallProfileProtocolOptionsSmtp>;
    ssh?: pulumi.Input<inputs.FirewallProfileProtocolOptionsSsh>;
    switchingProtocolsLog?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallProfileProtocolOptions resource.
 */
export interface FirewallProfileProtocolOptionsArgs {
    cifs?: pulumi.Input<inputs.FirewallProfileProtocolOptionsCifs>;
    comment?: pulumi.Input<string>;
    dns?: pulumi.Input<inputs.FirewallProfileProtocolOptionsDns>;
    featureSet?: pulumi.Input<string>;
    ftp?: pulumi.Input<inputs.FirewallProfileProtocolOptionsFtp>;
    getAllTables?: pulumi.Input<string>;
    http?: pulumi.Input<inputs.FirewallProfileProtocolOptionsHttp>;
    imap?: pulumi.Input<inputs.FirewallProfileProtocolOptionsImap>;
    mailSignature?: pulumi.Input<inputs.FirewallProfileProtocolOptionsMailSignature>;
    mapi?: pulumi.Input<inputs.FirewallProfileProtocolOptionsMapi>;
    name?: pulumi.Input<string>;
    nntp?: pulumi.Input<inputs.FirewallProfileProtocolOptionsNntp>;
    oversizeLog?: pulumi.Input<string>;
    pop3?: pulumi.Input<inputs.FirewallProfileProtocolOptionsPop3>;
    replacemsgGroup?: pulumi.Input<string>;
    rpcOverHttp?: pulumi.Input<string>;
    smtp?: pulumi.Input<inputs.FirewallProfileProtocolOptionsSmtp>;
    ssh?: pulumi.Input<inputs.FirewallProfileProtocolOptionsSsh>;
    switchingProtocolsLog?: pulumi.Input<string>;
    vdomparam?: pulumi.Input<string>;
}
