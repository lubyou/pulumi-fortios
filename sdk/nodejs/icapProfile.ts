// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Configure ICAP profiles.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumi/fortios";
 *
 * const trname = new fortios.IcapProfile("trname", {
 *     icapHeaders: [{
 *         base64Encoding: "disable",
 *         content: "$user",
 *         name: "X-Authenticated-User",
 *     }],
 *     methods: "delete get head options post put trace other",
 *     request: "disable",
 *     requestFailure: "error",
 *     response: "disable",
 *     responseFailure: "error",
 *     responseReqHdr: "disable",
 *     streamingContentBypass: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * Icap Profile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
 *
 * ```sh
 *  $ pulumi import fortios:index/icapProfile:IcapProfile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class IcapProfile extends pulumi.CustomResource {
    /**
     * Get an existing IcapProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IcapProfileState, opts?: pulumi.CustomResourceOptions): IcapProfile {
        return new IcapProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/icapProfile:IcapProfile';

    /**
     * Returns true if the given object is an instance of IcapProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IcapProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IcapProfile.__pulumiType;
    }

    /**
     * Enable/disable chunked encapsulation (default = disable). Valid values: `disable`, `enable`.
     */
    public readonly chunkEncap!: pulumi.Output<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable ICAP extension features. Valid values: `scan-progress`.
     */
    public readonly extensionFeature!: pulumi.Output<string>;
    /**
     * Enable/disable UTM log when infection found (default = disable). Valid values: `disable`, `enable`.
     */
    public readonly icapBlockLog!: pulumi.Output<string>;
    /**
     * Configure ICAP forwarded request headers. The structure of `icapHeaders` block is documented below.
     */
    public readonly icapHeaders!: pulumi.Output<outputs.IcapProfileIcapHeader[] | undefined>;
    /**
     * The allowed HTTP methods that will be sent to ICAP server for further processing. Valid values: `delete`, `get`, `head`, `options`, `post`, `put`, `trace`, `other`.
     */
    public readonly methods!: pulumi.Output<string>;
    /**
     * Address name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable preview of data to ICAP server. Valid values: `disable`, `enable`.
     */
    public readonly preview!: pulumi.Output<string>;
    /**
     * Preview data length to be sent to ICAP server.
     */
    public readonly previewDataLength!: pulumi.Output<number>;
    /**
     * Replacement message group.
     */
    public readonly replacemsgGroup!: pulumi.Output<string>;
    /**
     * Enable/disable whether an HTTP request is passed to an ICAP server. Valid values: `disable`, `enable`.
     */
    public readonly request!: pulumi.Output<string>;
    /**
     * Action to take if the ICAP server cannot be contacted when processing an HTTP request. Valid values: `error`, `bypass`.
     */
    public readonly requestFailure!: pulumi.Output<string>;
    /**
     * Path component of the ICAP URI that identifies the HTTP request processing service.
     */
    public readonly requestPath!: pulumi.Output<string>;
    /**
     * ICAP server to use for an HTTP request.
     */
    public readonly requestServer!: pulumi.Output<string>;
    /**
     * Default action to ICAP response modification (respmod) processing. Valid values: `forward`, `bypass`.
     */
    public readonly respmodDefaultAction!: pulumi.Output<string>;
    /**
     * ICAP response mode forward rules. The structure of `respmodForwardRules` block is documented below.
     */
    public readonly respmodForwardRules!: pulumi.Output<outputs.IcapProfileRespmodForwardRule[] | undefined>;
    /**
     * Enable/disable whether an HTTP response is passed to an ICAP server. Valid values: `disable`, `enable`.
     */
    public readonly response!: pulumi.Output<string>;
    /**
     * Action to take if the ICAP server cannot be contacted when processing an HTTP response. Valid values: `error`, `bypass`.
     */
    public readonly responseFailure!: pulumi.Output<string>;
    /**
     * Path component of the ICAP URI that identifies the HTTP response processing service.
     */
    public readonly responsePath!: pulumi.Output<string>;
    /**
     * Enable/disable addition of req-hdr for ICAP response modification (respmod) processing. Valid values: `disable`, `enable`.
     */
    public readonly responseReqHdr!: pulumi.Output<string>;
    /**
     * ICAP server to use for an HTTP response.
     */
    public readonly responseServer!: pulumi.Output<string>;
    /**
     * Scan progress interval value.
     */
    public readonly scanProgressInterval!: pulumi.Output<number>;
    /**
     * Enable/disable bypassing of ICAP server for streaming content. Valid values: `disable`, `enable`.
     */
    public readonly streamingContentBypass!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a IcapProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IcapProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IcapProfileArgs | IcapProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IcapProfileState | undefined;
            resourceInputs["chunkEncap"] = state ? state.chunkEncap : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["extensionFeature"] = state ? state.extensionFeature : undefined;
            resourceInputs["icapBlockLog"] = state ? state.icapBlockLog : undefined;
            resourceInputs["icapHeaders"] = state ? state.icapHeaders : undefined;
            resourceInputs["methods"] = state ? state.methods : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["preview"] = state ? state.preview : undefined;
            resourceInputs["previewDataLength"] = state ? state.previewDataLength : undefined;
            resourceInputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            resourceInputs["request"] = state ? state.request : undefined;
            resourceInputs["requestFailure"] = state ? state.requestFailure : undefined;
            resourceInputs["requestPath"] = state ? state.requestPath : undefined;
            resourceInputs["requestServer"] = state ? state.requestServer : undefined;
            resourceInputs["respmodDefaultAction"] = state ? state.respmodDefaultAction : undefined;
            resourceInputs["respmodForwardRules"] = state ? state.respmodForwardRules : undefined;
            resourceInputs["response"] = state ? state.response : undefined;
            resourceInputs["responseFailure"] = state ? state.responseFailure : undefined;
            resourceInputs["responsePath"] = state ? state.responsePath : undefined;
            resourceInputs["responseReqHdr"] = state ? state.responseReqHdr : undefined;
            resourceInputs["responseServer"] = state ? state.responseServer : undefined;
            resourceInputs["scanProgressInterval"] = state ? state.scanProgressInterval : undefined;
            resourceInputs["streamingContentBypass"] = state ? state.streamingContentBypass : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as IcapProfileArgs | undefined;
            resourceInputs["chunkEncap"] = args ? args.chunkEncap : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["extensionFeature"] = args ? args.extensionFeature : undefined;
            resourceInputs["icapBlockLog"] = args ? args.icapBlockLog : undefined;
            resourceInputs["icapHeaders"] = args ? args.icapHeaders : undefined;
            resourceInputs["methods"] = args ? args.methods : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["preview"] = args ? args.preview : undefined;
            resourceInputs["previewDataLength"] = args ? args.previewDataLength : undefined;
            resourceInputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            resourceInputs["request"] = args ? args.request : undefined;
            resourceInputs["requestFailure"] = args ? args.requestFailure : undefined;
            resourceInputs["requestPath"] = args ? args.requestPath : undefined;
            resourceInputs["requestServer"] = args ? args.requestServer : undefined;
            resourceInputs["respmodDefaultAction"] = args ? args.respmodDefaultAction : undefined;
            resourceInputs["respmodForwardRules"] = args ? args.respmodForwardRules : undefined;
            resourceInputs["response"] = args ? args.response : undefined;
            resourceInputs["responseFailure"] = args ? args.responseFailure : undefined;
            resourceInputs["responsePath"] = args ? args.responsePath : undefined;
            resourceInputs["responseReqHdr"] = args ? args.responseReqHdr : undefined;
            resourceInputs["responseServer"] = args ? args.responseServer : undefined;
            resourceInputs["scanProgressInterval"] = args ? args.scanProgressInterval : undefined;
            resourceInputs["streamingContentBypass"] = args ? args.streamingContentBypass : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IcapProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IcapProfile resources.
 */
export interface IcapProfileState {
    /**
     * Enable/disable chunked encapsulation (default = disable). Valid values: `disable`, `enable`.
     */
    chunkEncap?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable ICAP extension features. Valid values: `scan-progress`.
     */
    extensionFeature?: pulumi.Input<string>;
    /**
     * Enable/disable UTM log when infection found (default = disable). Valid values: `disable`, `enable`.
     */
    icapBlockLog?: pulumi.Input<string>;
    /**
     * Configure ICAP forwarded request headers. The structure of `icapHeaders` block is documented below.
     */
    icapHeaders?: pulumi.Input<pulumi.Input<inputs.IcapProfileIcapHeader>[]>;
    /**
     * The allowed HTTP methods that will be sent to ICAP server for further processing. Valid values: `delete`, `get`, `head`, `options`, `post`, `put`, `trace`, `other`.
     */
    methods?: pulumi.Input<string>;
    /**
     * Address name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable preview of data to ICAP server. Valid values: `disable`, `enable`.
     */
    preview?: pulumi.Input<string>;
    /**
     * Preview data length to be sent to ICAP server.
     */
    previewDataLength?: pulumi.Input<number>;
    /**
     * Replacement message group.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Enable/disable whether an HTTP request is passed to an ICAP server. Valid values: `disable`, `enable`.
     */
    request?: pulumi.Input<string>;
    /**
     * Action to take if the ICAP server cannot be contacted when processing an HTTP request. Valid values: `error`, `bypass`.
     */
    requestFailure?: pulumi.Input<string>;
    /**
     * Path component of the ICAP URI that identifies the HTTP request processing service.
     */
    requestPath?: pulumi.Input<string>;
    /**
     * ICAP server to use for an HTTP request.
     */
    requestServer?: pulumi.Input<string>;
    /**
     * Default action to ICAP response modification (respmod) processing. Valid values: `forward`, `bypass`.
     */
    respmodDefaultAction?: pulumi.Input<string>;
    /**
     * ICAP response mode forward rules. The structure of `respmodForwardRules` block is documented below.
     */
    respmodForwardRules?: pulumi.Input<pulumi.Input<inputs.IcapProfileRespmodForwardRule>[]>;
    /**
     * Enable/disable whether an HTTP response is passed to an ICAP server. Valid values: `disable`, `enable`.
     */
    response?: pulumi.Input<string>;
    /**
     * Action to take if the ICAP server cannot be contacted when processing an HTTP response. Valid values: `error`, `bypass`.
     */
    responseFailure?: pulumi.Input<string>;
    /**
     * Path component of the ICAP URI that identifies the HTTP response processing service.
     */
    responsePath?: pulumi.Input<string>;
    /**
     * Enable/disable addition of req-hdr for ICAP response modification (respmod) processing. Valid values: `disable`, `enable`.
     */
    responseReqHdr?: pulumi.Input<string>;
    /**
     * ICAP server to use for an HTTP response.
     */
    responseServer?: pulumi.Input<string>;
    /**
     * Scan progress interval value.
     */
    scanProgressInterval?: pulumi.Input<number>;
    /**
     * Enable/disable bypassing of ICAP server for streaming content. Valid values: `disable`, `enable`.
     */
    streamingContentBypass?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IcapProfile resource.
 */
export interface IcapProfileArgs {
    /**
     * Enable/disable chunked encapsulation (default = disable). Valid values: `disable`, `enable`.
     */
    chunkEncap?: pulumi.Input<string>;
    /**
     * true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable ICAP extension features. Valid values: `scan-progress`.
     */
    extensionFeature?: pulumi.Input<string>;
    /**
     * Enable/disable UTM log when infection found (default = disable). Valid values: `disable`, `enable`.
     */
    icapBlockLog?: pulumi.Input<string>;
    /**
     * Configure ICAP forwarded request headers. The structure of `icapHeaders` block is documented below.
     */
    icapHeaders?: pulumi.Input<pulumi.Input<inputs.IcapProfileIcapHeader>[]>;
    /**
     * The allowed HTTP methods that will be sent to ICAP server for further processing. Valid values: `delete`, `get`, `head`, `options`, `post`, `put`, `trace`, `other`.
     */
    methods?: pulumi.Input<string>;
    /**
     * Address name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable preview of data to ICAP server. Valid values: `disable`, `enable`.
     */
    preview?: pulumi.Input<string>;
    /**
     * Preview data length to be sent to ICAP server.
     */
    previewDataLength?: pulumi.Input<number>;
    /**
     * Replacement message group.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Enable/disable whether an HTTP request is passed to an ICAP server. Valid values: `disable`, `enable`.
     */
    request?: pulumi.Input<string>;
    /**
     * Action to take if the ICAP server cannot be contacted when processing an HTTP request. Valid values: `error`, `bypass`.
     */
    requestFailure?: pulumi.Input<string>;
    /**
     * Path component of the ICAP URI that identifies the HTTP request processing service.
     */
    requestPath?: pulumi.Input<string>;
    /**
     * ICAP server to use for an HTTP request.
     */
    requestServer?: pulumi.Input<string>;
    /**
     * Default action to ICAP response modification (respmod) processing. Valid values: `forward`, `bypass`.
     */
    respmodDefaultAction?: pulumi.Input<string>;
    /**
     * ICAP response mode forward rules. The structure of `respmodForwardRules` block is documented below.
     */
    respmodForwardRules?: pulumi.Input<pulumi.Input<inputs.IcapProfileRespmodForwardRule>[]>;
    /**
     * Enable/disable whether an HTTP response is passed to an ICAP server. Valid values: `disable`, `enable`.
     */
    response?: pulumi.Input<string>;
    /**
     * Action to take if the ICAP server cannot be contacted when processing an HTTP response. Valid values: `error`, `bypass`.
     */
    responseFailure?: pulumi.Input<string>;
    /**
     * Path component of the ICAP URI that identifies the HTTP response processing service.
     */
    responsePath?: pulumi.Input<string>;
    /**
     * Enable/disable addition of req-hdr for ICAP response modification (respmod) processing. Valid values: `disable`, `enable`.
     */
    responseReqHdr?: pulumi.Input<string>;
    /**
     * ICAP server to use for an HTTP response.
     */
    responseServer?: pulumi.Input<string>;
    /**
     * Scan progress interval value.
     */
    scanProgressInterval?: pulumi.Input<number>;
    /**
     * Enable/disable bypassing of ICAP server for streaming content. Valid values: `disable`, `enable`.
     */
    streamingContentBypass?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
