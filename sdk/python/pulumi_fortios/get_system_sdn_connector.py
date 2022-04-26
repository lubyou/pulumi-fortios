# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetSystemSdnConnectorResult',
    'AwaitableGetSystemSdnConnectorResult',
    'get_system_sdn_connector',
    'get_system_sdn_connector_output',
]

@pulumi.output_type
class GetSystemSdnConnectorResult:
    """
    A collection of values returned by GetSystemSdnConnector.
    """
    def __init__(__self__, access_key=None, api_key=None, azure_region=None, client_id=None, client_secret=None, compartment_id=None, compute_generation=None, domain=None, external_account_lists=None, external_ips=None, forwarding_rules=None, gcp_project=None, gcp_project_lists=None, group_name=None, ha_status=None, ibm_region=None, ibm_region_gen1=None, ibm_region_gen2=None, id=None, key_passwd=None, login_endpoint=None, name=None, nics=None, oci_cert=None, oci_fingerprint=None, oci_region=None, oci_region_type=None, password=None, private_key=None, region=None, resource_group=None, resource_url=None, route_tables=None, routes=None, secret_key=None, secret_token=None, server=None, server_lists=None, server_port=None, service_account=None, status=None, subscription_id=None, tenant_id=None, type=None, update_interval=None, use_metadata_iam=None, user_id=None, username=None, vcenter_password=None, vcenter_server=None, vcenter_username=None, vdomparam=None, verify_certificate=None, vpc_id=None):
        if access_key and not isinstance(access_key, str):
            raise TypeError("Expected argument 'access_key' to be a str")
        pulumi.set(__self__, "access_key", access_key)
        if api_key and not isinstance(api_key, str):
            raise TypeError("Expected argument 'api_key' to be a str")
        pulumi.set(__self__, "api_key", api_key)
        if azure_region and not isinstance(azure_region, str):
            raise TypeError("Expected argument 'azure_region' to be a str")
        pulumi.set(__self__, "azure_region", azure_region)
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if client_secret and not isinstance(client_secret, str):
            raise TypeError("Expected argument 'client_secret' to be a str")
        pulumi.set(__self__, "client_secret", client_secret)
        if compartment_id and not isinstance(compartment_id, str):
            raise TypeError("Expected argument 'compartment_id' to be a str")
        pulumi.set(__self__, "compartment_id", compartment_id)
        if compute_generation and not isinstance(compute_generation, int):
            raise TypeError("Expected argument 'compute_generation' to be a int")
        pulumi.set(__self__, "compute_generation", compute_generation)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if external_account_lists and not isinstance(external_account_lists, list):
            raise TypeError("Expected argument 'external_account_lists' to be a list")
        pulumi.set(__self__, "external_account_lists", external_account_lists)
        if external_ips and not isinstance(external_ips, list):
            raise TypeError("Expected argument 'external_ips' to be a list")
        pulumi.set(__self__, "external_ips", external_ips)
        if forwarding_rules and not isinstance(forwarding_rules, list):
            raise TypeError("Expected argument 'forwarding_rules' to be a list")
        pulumi.set(__self__, "forwarding_rules", forwarding_rules)
        if gcp_project and not isinstance(gcp_project, str):
            raise TypeError("Expected argument 'gcp_project' to be a str")
        pulumi.set(__self__, "gcp_project", gcp_project)
        if gcp_project_lists and not isinstance(gcp_project_lists, list):
            raise TypeError("Expected argument 'gcp_project_lists' to be a list")
        pulumi.set(__self__, "gcp_project_lists", gcp_project_lists)
        if group_name and not isinstance(group_name, str):
            raise TypeError("Expected argument 'group_name' to be a str")
        pulumi.set(__self__, "group_name", group_name)
        if ha_status and not isinstance(ha_status, str):
            raise TypeError("Expected argument 'ha_status' to be a str")
        pulumi.set(__self__, "ha_status", ha_status)
        if ibm_region and not isinstance(ibm_region, str):
            raise TypeError("Expected argument 'ibm_region' to be a str")
        pulumi.set(__self__, "ibm_region", ibm_region)
        if ibm_region_gen1 and not isinstance(ibm_region_gen1, str):
            raise TypeError("Expected argument 'ibm_region_gen1' to be a str")
        pulumi.set(__self__, "ibm_region_gen1", ibm_region_gen1)
        if ibm_region_gen2 and not isinstance(ibm_region_gen2, str):
            raise TypeError("Expected argument 'ibm_region_gen2' to be a str")
        pulumi.set(__self__, "ibm_region_gen2", ibm_region_gen2)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key_passwd and not isinstance(key_passwd, str):
            raise TypeError("Expected argument 'key_passwd' to be a str")
        pulumi.set(__self__, "key_passwd", key_passwd)
        if login_endpoint and not isinstance(login_endpoint, str):
            raise TypeError("Expected argument 'login_endpoint' to be a str")
        pulumi.set(__self__, "login_endpoint", login_endpoint)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nics and not isinstance(nics, list):
            raise TypeError("Expected argument 'nics' to be a list")
        pulumi.set(__self__, "nics", nics)
        if oci_cert and not isinstance(oci_cert, str):
            raise TypeError("Expected argument 'oci_cert' to be a str")
        pulumi.set(__self__, "oci_cert", oci_cert)
        if oci_fingerprint and not isinstance(oci_fingerprint, str):
            raise TypeError("Expected argument 'oci_fingerprint' to be a str")
        pulumi.set(__self__, "oci_fingerprint", oci_fingerprint)
        if oci_region and not isinstance(oci_region, str):
            raise TypeError("Expected argument 'oci_region' to be a str")
        pulumi.set(__self__, "oci_region", oci_region)
        if oci_region_type and not isinstance(oci_region_type, str):
            raise TypeError("Expected argument 'oci_region_type' to be a str")
        pulumi.set(__self__, "oci_region_type", oci_region_type)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if private_key and not isinstance(private_key, str):
            raise TypeError("Expected argument 'private_key' to be a str")
        pulumi.set(__self__, "private_key", private_key)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if resource_group and not isinstance(resource_group, str):
            raise TypeError("Expected argument 'resource_group' to be a str")
        pulumi.set(__self__, "resource_group", resource_group)
        if resource_url and not isinstance(resource_url, str):
            raise TypeError("Expected argument 'resource_url' to be a str")
        pulumi.set(__self__, "resource_url", resource_url)
        if route_tables and not isinstance(route_tables, list):
            raise TypeError("Expected argument 'route_tables' to be a list")
        pulumi.set(__self__, "route_tables", route_tables)
        if routes and not isinstance(routes, list):
            raise TypeError("Expected argument 'routes' to be a list")
        pulumi.set(__self__, "routes", routes)
        if secret_key and not isinstance(secret_key, str):
            raise TypeError("Expected argument 'secret_key' to be a str")
        pulumi.set(__self__, "secret_key", secret_key)
        if secret_token and not isinstance(secret_token, str):
            raise TypeError("Expected argument 'secret_token' to be a str")
        pulumi.set(__self__, "secret_token", secret_token)
        if server and not isinstance(server, str):
            raise TypeError("Expected argument 'server' to be a str")
        pulumi.set(__self__, "server", server)
        if server_lists and not isinstance(server_lists, list):
            raise TypeError("Expected argument 'server_lists' to be a list")
        pulumi.set(__self__, "server_lists", server_lists)
        if server_port and not isinstance(server_port, int):
            raise TypeError("Expected argument 'server_port' to be a int")
        pulumi.set(__self__, "server_port", server_port)
        if service_account and not isinstance(service_account, str):
            raise TypeError("Expected argument 'service_account' to be a str")
        pulumi.set(__self__, "service_account", service_account)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if subscription_id and not isinstance(subscription_id, str):
            raise TypeError("Expected argument 'subscription_id' to be a str")
        pulumi.set(__self__, "subscription_id", subscription_id)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if update_interval and not isinstance(update_interval, int):
            raise TypeError("Expected argument 'update_interval' to be a int")
        pulumi.set(__self__, "update_interval", update_interval)
        if use_metadata_iam and not isinstance(use_metadata_iam, str):
            raise TypeError("Expected argument 'use_metadata_iam' to be a str")
        pulumi.set(__self__, "use_metadata_iam", use_metadata_iam)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)
        if vcenter_password and not isinstance(vcenter_password, str):
            raise TypeError("Expected argument 'vcenter_password' to be a str")
        pulumi.set(__self__, "vcenter_password", vcenter_password)
        if vcenter_server and not isinstance(vcenter_server, str):
            raise TypeError("Expected argument 'vcenter_server' to be a str")
        pulumi.set(__self__, "vcenter_server", vcenter_server)
        if vcenter_username and not isinstance(vcenter_username, str):
            raise TypeError("Expected argument 'vcenter_username' to be a str")
        pulumi.set(__self__, "vcenter_username", vcenter_username)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if verify_certificate and not isinstance(verify_certificate, str):
            raise TypeError("Expected argument 'verify_certificate' to be a str")
        pulumi.set(__self__, "verify_certificate", verify_certificate)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> str:
        """
        AWS access key ID.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> str:
        """
        IBM cloud API key or service ID API key.
        """
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter(name="azureRegion")
    def azure_region(self) -> str:
        """
        Azure server region.
        """
        return pulumi.get(self, "azure_region")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        Azure client ID (application ID).
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> str:
        """
        Azure client secret (application key).
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="compartmentId")
    def compartment_id(self) -> str:
        """
        Compartment ID.
        """
        return pulumi.get(self, "compartment_id")

    @property
    @pulumi.getter(name="computeGeneration")
    def compute_generation(self) -> int:
        """
        Compute generation for IBM cloud infrastructure.
        """
        return pulumi.get(self, "compute_generation")

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        Domain name.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="externalAccountLists")
    def external_account_lists(self) -> Sequence['outputs.GetSystemSdnConnectorExternalAccountListResult']:
        """
        Configure AWS external account list. The structure of `external_account_list` block is documented below.
        """
        return pulumi.get(self, "external_account_lists")

    @property
    @pulumi.getter(name="externalIps")
    def external_ips(self) -> Sequence['outputs.GetSystemSdnConnectorExternalIpResult']:
        """
        Configure GCP external IP. The structure of `external_ip` block is documented below.
        """
        return pulumi.get(self, "external_ips")

    @property
    @pulumi.getter(name="forwardingRules")
    def forwarding_rules(self) -> Sequence['outputs.GetSystemSdnConnectorForwardingRuleResult']:
        """
        Configure GCP forwarding rule. The structure of `forwarding_rule` block is documented below.
        """
        return pulumi.get(self, "forwarding_rules")

    @property
    @pulumi.getter(name="gcpProject")
    def gcp_project(self) -> str:
        """
        GCP project name.
        """
        return pulumi.get(self, "gcp_project")

    @property
    @pulumi.getter(name="gcpProjectLists")
    def gcp_project_lists(self) -> Sequence['outputs.GetSystemSdnConnectorGcpProjectListResult']:
        """
        Configure GCP project list. The structure of `gcp_project_list` block is documented below.
        """
        return pulumi.get(self, "gcp_project_lists")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> str:
        """
        Group name of computers.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter(name="haStatus")
    def ha_status(self) -> str:
        """
        Enable/disable use for FortiGate HA service.
        """
        return pulumi.get(self, "ha_status")

    @property
    @pulumi.getter(name="ibmRegion")
    def ibm_region(self) -> str:
        """
        IBM cloud region name.
        """
        return pulumi.get(self, "ibm_region")

    @property
    @pulumi.getter(name="ibmRegionGen1")
    def ibm_region_gen1(self) -> str:
        """
        IBM cloud compute generation 1 region name.
        """
        return pulumi.get(self, "ibm_region_gen1")

    @property
    @pulumi.getter(name="ibmRegionGen2")
    def ibm_region_gen2(self) -> str:
        """
        IBM cloud compute generation 2 region name.
        """
        return pulumi.get(self, "ibm_region_gen2")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keyPasswd")
    def key_passwd(self) -> str:
        """
        Private key password.
        """
        return pulumi.get(self, "key_passwd")

    @property
    @pulumi.getter(name="loginEndpoint")
    def login_endpoint(self) -> str:
        """
        Azure Stack login endpoint.
        """
        return pulumi.get(self, "login_endpoint")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        GCP zone name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def nics(self) -> Sequence['outputs.GetSystemSdnConnectorNicResult']:
        """
        Configure Azure network interface. The structure of `nic` block is documented below.
        """
        return pulumi.get(self, "nics")

    @property
    @pulumi.getter(name="ociCert")
    def oci_cert(self) -> str:
        """
        OCI certificate.
        """
        return pulumi.get(self, "oci_cert")

    @property
    @pulumi.getter(name="ociFingerprint")
    def oci_fingerprint(self) -> str:
        """
        OCI pubkey fingerprint.
        """
        return pulumi.get(self, "oci_fingerprint")

    @property
    @pulumi.getter(name="ociRegion")
    def oci_region(self) -> str:
        """
        OCI server region.
        """
        return pulumi.get(self, "oci_region")

    @property
    @pulumi.getter(name="ociRegionType")
    def oci_region_type(self) -> str:
        """
        OCI region type.
        """
        return pulumi.get(self, "oci_region_type")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        Password of the remote SDN connector as login credentials.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> str:
        """
        Private key of GCP service account.
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        AWS region name.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> str:
        """
        Resource group of Azure route table.
        """
        return pulumi.get(self, "resource_group")

    @property
    @pulumi.getter(name="resourceUrl")
    def resource_url(self) -> str:
        """
        Azure Stack resource URL.
        """
        return pulumi.get(self, "resource_url")

    @property
    @pulumi.getter(name="routeTables")
    def route_tables(self) -> Sequence['outputs.GetSystemSdnConnectorRouteTableResult']:
        """
        Configure Azure route table. The structure of `route_table` block is documented below.
        """
        return pulumi.get(self, "route_tables")

    @property
    @pulumi.getter
    def routes(self) -> Sequence['outputs.GetSystemSdnConnectorRouteResult']:
        """
        Configure Azure route. The structure of `route` block is documented below.
        """
        return pulumi.get(self, "routes")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> str:
        """
        AWS secret access key.
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="secretToken")
    def secret_token(self) -> str:
        """
        Secret token of Kubernetes service account.
        """
        return pulumi.get(self, "secret_token")

    @property
    @pulumi.getter
    def server(self) -> str:
        """
        Server address of the remote SDN connector.
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter(name="serverLists")
    def server_lists(self) -> Sequence['outputs.GetSystemSdnConnectorServerListResult']:
        """
        Server address list of the remote SDN connector. The structure of `server_list` block is documented below.
        """
        return pulumi.get(self, "server_lists")

    @property
    @pulumi.getter(name="serverPort")
    def server_port(self) -> int:
        """
        Port number of the remote SDN connector.
        """
        return pulumi.get(self, "server_port")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> str:
        """
        GCP service account email.
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable connection to the remote SDN connector.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> str:
        """
        Subscription ID of Azure route table.
        """
        return pulumi.get(self, "subscription_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        Tenant ID (directory ID).
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of SDN connector.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updateInterval")
    def update_interval(self) -> int:
        """
        Dynamic object update interval (0 - 3600 sec, 0 means disabled, default = 60).
        """
        return pulumi.get(self, "update_interval")

    @property
    @pulumi.getter(name="useMetadataIam")
    def use_metadata_iam(self) -> str:
        """
        Enable/disable using IAM role from metadata to call API.
        """
        return pulumi.get(self, "use_metadata_iam")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        """
        User ID.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        Username of the remote SDN connector as login credentials.
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="vcenterPassword")
    def vcenter_password(self) -> str:
        """
        vCenter server password for NSX quarantine.
        """
        return pulumi.get(self, "vcenter_password")

    @property
    @pulumi.getter(name="vcenterServer")
    def vcenter_server(self) -> str:
        """
        vCenter server address for NSX quarantine.
        """
        return pulumi.get(self, "vcenter_server")

    @property
    @pulumi.getter(name="vcenterUsername")
    def vcenter_username(self) -> str:
        """
        vCenter server username for NSX quarantine.
        """
        return pulumi.get(self, "vcenter_username")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="verifyCertificate")
    def verify_certificate(self) -> str:
        """
        Enable/disable server certificate verification.
        """
        return pulumi.get(self, "verify_certificate")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        AWS VPC ID.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableGetSystemSdnConnectorResult(GetSystemSdnConnectorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemSdnConnectorResult(
            access_key=self.access_key,
            api_key=self.api_key,
            azure_region=self.azure_region,
            client_id=self.client_id,
            client_secret=self.client_secret,
            compartment_id=self.compartment_id,
            compute_generation=self.compute_generation,
            domain=self.domain,
            external_account_lists=self.external_account_lists,
            external_ips=self.external_ips,
            forwarding_rules=self.forwarding_rules,
            gcp_project=self.gcp_project,
            gcp_project_lists=self.gcp_project_lists,
            group_name=self.group_name,
            ha_status=self.ha_status,
            ibm_region=self.ibm_region,
            ibm_region_gen1=self.ibm_region_gen1,
            ibm_region_gen2=self.ibm_region_gen2,
            id=self.id,
            key_passwd=self.key_passwd,
            login_endpoint=self.login_endpoint,
            name=self.name,
            nics=self.nics,
            oci_cert=self.oci_cert,
            oci_fingerprint=self.oci_fingerprint,
            oci_region=self.oci_region,
            oci_region_type=self.oci_region_type,
            password=self.password,
            private_key=self.private_key,
            region=self.region,
            resource_group=self.resource_group,
            resource_url=self.resource_url,
            route_tables=self.route_tables,
            routes=self.routes,
            secret_key=self.secret_key,
            secret_token=self.secret_token,
            server=self.server,
            server_lists=self.server_lists,
            server_port=self.server_port,
            service_account=self.service_account,
            status=self.status,
            subscription_id=self.subscription_id,
            tenant_id=self.tenant_id,
            type=self.type,
            update_interval=self.update_interval,
            use_metadata_iam=self.use_metadata_iam,
            user_id=self.user_id,
            username=self.username,
            vcenter_password=self.vcenter_password,
            vcenter_server=self.vcenter_server,
            vcenter_username=self.vcenter_username,
            vdomparam=self.vdomparam,
            verify_certificate=self.verify_certificate,
            vpc_id=self.vpc_id)


def get_system_sdn_connector(name: Optional[str] = None,
                             vdomparam: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemSdnConnectorResult:
    """
    Use this data source to get information on an fortios system sdnconnector


    :param str name: Specify the name of the desired system sdnconnector.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemSdnConnector:GetSystemSdnConnector', __args__, opts=opts, typ=GetSystemSdnConnectorResult).value

    return AwaitableGetSystemSdnConnectorResult(
        access_key=__ret__.access_key,
        api_key=__ret__.api_key,
        azure_region=__ret__.azure_region,
        client_id=__ret__.client_id,
        client_secret=__ret__.client_secret,
        compartment_id=__ret__.compartment_id,
        compute_generation=__ret__.compute_generation,
        domain=__ret__.domain,
        external_account_lists=__ret__.external_account_lists,
        external_ips=__ret__.external_ips,
        forwarding_rules=__ret__.forwarding_rules,
        gcp_project=__ret__.gcp_project,
        gcp_project_lists=__ret__.gcp_project_lists,
        group_name=__ret__.group_name,
        ha_status=__ret__.ha_status,
        ibm_region=__ret__.ibm_region,
        ibm_region_gen1=__ret__.ibm_region_gen1,
        ibm_region_gen2=__ret__.ibm_region_gen2,
        id=__ret__.id,
        key_passwd=__ret__.key_passwd,
        login_endpoint=__ret__.login_endpoint,
        name=__ret__.name,
        nics=__ret__.nics,
        oci_cert=__ret__.oci_cert,
        oci_fingerprint=__ret__.oci_fingerprint,
        oci_region=__ret__.oci_region,
        oci_region_type=__ret__.oci_region_type,
        password=__ret__.password,
        private_key=__ret__.private_key,
        region=__ret__.region,
        resource_group=__ret__.resource_group,
        resource_url=__ret__.resource_url,
        route_tables=__ret__.route_tables,
        routes=__ret__.routes,
        secret_key=__ret__.secret_key,
        secret_token=__ret__.secret_token,
        server=__ret__.server,
        server_lists=__ret__.server_lists,
        server_port=__ret__.server_port,
        service_account=__ret__.service_account,
        status=__ret__.status,
        subscription_id=__ret__.subscription_id,
        tenant_id=__ret__.tenant_id,
        type=__ret__.type,
        update_interval=__ret__.update_interval,
        use_metadata_iam=__ret__.use_metadata_iam,
        user_id=__ret__.user_id,
        username=__ret__.username,
        vcenter_password=__ret__.vcenter_password,
        vcenter_server=__ret__.vcenter_server,
        vcenter_username=__ret__.vcenter_username,
        vdomparam=__ret__.vdomparam,
        verify_certificate=__ret__.verify_certificate,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(get_system_sdn_connector)
def get_system_sdn_connector_output(name: Optional[pulumi.Input[str]] = None,
                                    vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemSdnConnectorResult]:
    """
    Use this data source to get information on an fortios system sdnconnector


    :param str name: Specify the name of the desired system sdnconnector.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
