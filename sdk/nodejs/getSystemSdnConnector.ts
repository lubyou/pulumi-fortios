// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios system sdnconnector
 */
export function getSystemSdnConnector(args: GetSystemSdnConnectorArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemSdnConnectorResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("fortios:index/getSystemSdnConnector:GetSystemSdnConnector", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking GetSystemSdnConnector.
 */
export interface GetSystemSdnConnectorArgs {
    /**
     * Specify the name of the desired system sdnconnector.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by GetSystemSdnConnector.
 */
export interface GetSystemSdnConnectorResult {
    /**
     * AWS access key ID.
     */
    readonly accessKey: string;
    /**
     * IBM cloud API key or service ID API key.
     */
    readonly apiKey: string;
    /**
     * Azure server region.
     */
    readonly azureRegion: string;
    /**
     * Azure client ID (application ID).
     */
    readonly clientId: string;
    /**
     * Azure client secret (application key).
     */
    readonly clientSecret: string;
    /**
     * Compartment ID.
     */
    readonly compartmentId: string;
    /**
     * Compute generation for IBM cloud infrastructure.
     */
    readonly computeGeneration: number;
    /**
     * Domain name.
     */
    readonly domain: string;
    /**
     * Configure AWS external account list. The structure of `externalAccountList` block is documented below.
     */
    readonly externalAccountLists: outputs.GetSystemSdnConnectorExternalAccountList[];
    /**
     * Configure GCP external IP. The structure of `externalIp` block is documented below.
     */
    readonly externalIps: outputs.GetSystemSdnConnectorExternalIp[];
    /**
     * Configure GCP forwarding rule. The structure of `forwardingRule` block is documented below.
     */
    readonly forwardingRules: outputs.GetSystemSdnConnectorForwardingRule[];
    /**
     * GCP project name.
     */
    readonly gcpProject: string;
    /**
     * Configure GCP project list. The structure of `gcpProjectList` block is documented below.
     */
    readonly gcpProjectLists: outputs.GetSystemSdnConnectorGcpProjectList[];
    /**
     * Group name of computers.
     */
    readonly groupName: string;
    /**
     * Enable/disable use for FortiGate HA service.
     */
    readonly haStatus: string;
    /**
     * IBM cloud region name.
     */
    readonly ibmRegion: string;
    /**
     * IBM cloud compute generation 1 region name.
     */
    readonly ibmRegionGen1: string;
    /**
     * IBM cloud compute generation 2 region name.
     */
    readonly ibmRegionGen2: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Private key password.
     */
    readonly keyPasswd: string;
    /**
     * Azure Stack login endpoint.
     */
    readonly loginEndpoint: string;
    /**
     * GCP zone name.
     */
    readonly name: string;
    /**
     * Configure Azure network interface. The structure of `nic` block is documented below.
     */
    readonly nics: outputs.GetSystemSdnConnectorNic[];
    /**
     * OCI certificate.
     */
    readonly ociCert: string;
    /**
     * OCI pubkey fingerprint.
     */
    readonly ociFingerprint: string;
    /**
     * OCI server region.
     */
    readonly ociRegion: string;
    /**
     * OCI region type.
     */
    readonly ociRegionType: string;
    /**
     * Password of the remote SDN connector as login credentials.
     */
    readonly password: string;
    /**
     * Private key of GCP service account.
     */
    readonly privateKey: string;
    /**
     * AWS region name.
     */
    readonly region: string;
    /**
     * Resource group of Azure route table.
     */
    readonly resourceGroup: string;
    /**
     * Azure Stack resource URL.
     */
    readonly resourceUrl: string;
    /**
     * Configure Azure route table. The structure of `routeTable` block is documented below.
     */
    readonly routeTables: outputs.GetSystemSdnConnectorRouteTable[];
    /**
     * Configure Azure route. The structure of `route` block is documented below.
     */
    readonly routes: outputs.GetSystemSdnConnectorRoute[];
    /**
     * AWS secret access key.
     */
    readonly secretKey: string;
    /**
     * Secret token of Kubernetes service account.
     */
    readonly secretToken: string;
    /**
     * Server address of the remote SDN connector.
     */
    readonly server: string;
    /**
     * Server address list of the remote SDN connector. The structure of `serverList` block is documented below.
     */
    readonly serverLists: outputs.GetSystemSdnConnectorServerList[];
    /**
     * Port number of the remote SDN connector.
     */
    readonly serverPort: number;
    /**
     * GCP service account email.
     */
    readonly serviceAccount: string;
    /**
     * Enable/disable connection to the remote SDN connector.
     */
    readonly status: string;
    /**
     * Subscription ID of Azure route table.
     */
    readonly subscriptionId: string;
    /**
     * Tenant ID (directory ID).
     */
    readonly tenantId: string;
    /**
     * Type of SDN connector.
     */
    readonly type: string;
    /**
     * Dynamic object update interval (0 - 3600 sec, 0 means disabled, default = 60).
     */
    readonly updateInterval: number;
    /**
     * Enable/disable using IAM role from metadata to call API.
     */
    readonly useMetadataIam: string;
    /**
     * User ID.
     */
    readonly userId: string;
    /**
     * Username of the remote SDN connector as login credentials.
     */
    readonly username: string;
    /**
     * vCenter server password for NSX quarantine.
     */
    readonly vcenterPassword: string;
    /**
     * vCenter server address for NSX quarantine.
     */
    readonly vcenterServer: string;
    /**
     * vCenter server username for NSX quarantine.
     */
    readonly vcenterUsername: string;
    readonly vdomparam?: string;
    /**
     * Enable/disable server certificate verification.
     */
    readonly verifyCertificate: string;
    /**
     * AWS VPC ID.
     */
    readonly vpcId: string;
}

export function getSystemSdnConnectorOutput(args: GetSystemSdnConnectorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemSdnConnectorResult> {
    return pulumi.output(args).apply(a => getSystemSdnConnector(a, opts))
}

/**
 * A collection of arguments for invoking GetSystemSdnConnector.
 */
export interface GetSystemSdnConnectorOutputArgs {
    /**
     * Specify the name of the desired system sdnconnector.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
