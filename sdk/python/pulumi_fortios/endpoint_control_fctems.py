# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['EndpointControlFctemsArgs', 'EndpointControlFctems']

@pulumi.input_type
class EndpointControlFctemsArgs:
    def __init__(__self__, *,
                 admin_password: Optional[pulumi.Input[str]] = None,
                 admin_username: Optional[pulumi.Input[str]] = None,
                 call_timeout: Optional[pulumi.Input[int]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 cloud_server_type: Optional[pulumi.Input[str]] = None,
                 fortinetone_cloud_authentication: Optional[pulumi.Input[str]] = None,
                 https_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pull_avatars: Optional[pulumi.Input[str]] = None,
                 pull_sysinfo: Optional[pulumi.Input[str]] = None,
                 pull_tags: Optional[pulumi.Input[str]] = None,
                 pull_vulnerabilities: Optional[pulumi.Input[str]] = None,
                 serial_number: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EndpointControlFctems resource.
        :param pulumi.Input[str] admin_password: FortiClient EMS admin password.
        :param pulumi.Input[str] admin_username: FortiClient EMS admin username.
        :param pulumi.Input[int] call_timeout: FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
        :param pulumi.Input[str] certificate: FortiClient EMS certificate.
        :param pulumi.Input[str] cloud_server_type: Cloud server type. Valid values: `production`, `alpha`, `beta`.
        :param pulumi.Input[str] fortinetone_cloud_authentication: Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] https_port: FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        :param pulumi.Input[str] name: FortiClient Enterprise Management Server (EMS) name.
        :param pulumi.Input[str] pull_avatars: Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pull_sysinfo: Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pull_tags: Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pull_vulnerabilities: Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] serial_number: FortiClient EMS Serial Number.
        :param pulumi.Input[str] server: FortiClient EMS FQDN or IPv4 address.
        :param pulumi.Input[str] source_ip: REST API call source IP.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if admin_password is not None:
            pulumi.set(__self__, "admin_password", admin_password)
        if admin_username is not None:
            pulumi.set(__self__, "admin_username", admin_username)
        if call_timeout is not None:
            pulumi.set(__self__, "call_timeout", call_timeout)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if cloud_server_type is not None:
            pulumi.set(__self__, "cloud_server_type", cloud_server_type)
        if fortinetone_cloud_authentication is not None:
            pulumi.set(__self__, "fortinetone_cloud_authentication", fortinetone_cloud_authentication)
        if https_port is not None:
            pulumi.set(__self__, "https_port", https_port)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pull_avatars is not None:
            pulumi.set(__self__, "pull_avatars", pull_avatars)
        if pull_sysinfo is not None:
            pulumi.set(__self__, "pull_sysinfo", pull_sysinfo)
        if pull_tags is not None:
            pulumi.set(__self__, "pull_tags", pull_tags)
        if pull_vulnerabilities is not None:
            pulumi.set(__self__, "pull_vulnerabilities", pull_vulnerabilities)
        if serial_number is not None:
            pulumi.set(__self__, "serial_number", serial_number)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="adminPassword")
    def admin_password(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS admin password.
        """
        return pulumi.get(self, "admin_password")

    @admin_password.setter
    def admin_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_password", value)

    @property
    @pulumi.getter(name="adminUsername")
    def admin_username(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS admin username.
        """
        return pulumi.get(self, "admin_username")

    @admin_username.setter
    def admin_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_username", value)

    @property
    @pulumi.getter(name="callTimeout")
    def call_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
        """
        return pulumi.get(self, "call_timeout")

    @call_timeout.setter
    def call_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "call_timeout", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS certificate.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="cloudServerType")
    def cloud_server_type(self) -> Optional[pulumi.Input[str]]:
        """
        Cloud server type. Valid values: `production`, `alpha`, `beta`.
        """
        return pulumi.get(self, "cloud_server_type")

    @cloud_server_type.setter
    def cloud_server_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_server_type", value)

    @property
    @pulumi.getter(name="fortinetoneCloudAuthentication")
    def fortinetone_cloud_authentication(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fortinetone_cloud_authentication")

    @fortinetone_cloud_authentication.setter
    def fortinetone_cloud_authentication(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortinetone_cloud_authentication", value)

    @property
    @pulumi.getter(name="httpsPort")
    def https_port(self) -> Optional[pulumi.Input[int]]:
        """
        FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        """
        return pulumi.get(self, "https_port")

    @https_port.setter
    def https_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "https_port", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient Enterprise Management Server (EMS) name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pullAvatars")
    def pull_avatars(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "pull_avatars")

    @pull_avatars.setter
    def pull_avatars(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pull_avatars", value)

    @property
    @pulumi.getter(name="pullSysinfo")
    def pull_sysinfo(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "pull_sysinfo")

    @pull_sysinfo.setter
    def pull_sysinfo(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pull_sysinfo", value)

    @property
    @pulumi.getter(name="pullTags")
    def pull_tags(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "pull_tags")

    @pull_tags.setter
    def pull_tags(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pull_tags", value)

    @property
    @pulumi.getter(name="pullVulnerabilities")
    def pull_vulnerabilities(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "pull_vulnerabilities")

    @pull_vulnerabilities.setter
    def pull_vulnerabilities(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pull_vulnerabilities", value)

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS Serial Number.
        """
        return pulumi.get(self, "serial_number")

    @serial_number.setter
    def serial_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "serial_number", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS FQDN or IPv4 address.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        REST API call source IP.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _EndpointControlFctemsState:
    def __init__(__self__, *,
                 admin_password: Optional[pulumi.Input[str]] = None,
                 admin_username: Optional[pulumi.Input[str]] = None,
                 call_timeout: Optional[pulumi.Input[int]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 cloud_server_type: Optional[pulumi.Input[str]] = None,
                 fortinetone_cloud_authentication: Optional[pulumi.Input[str]] = None,
                 https_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pull_avatars: Optional[pulumi.Input[str]] = None,
                 pull_sysinfo: Optional[pulumi.Input[str]] = None,
                 pull_tags: Optional[pulumi.Input[str]] = None,
                 pull_vulnerabilities: Optional[pulumi.Input[str]] = None,
                 serial_number: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EndpointControlFctems resources.
        :param pulumi.Input[str] admin_password: FortiClient EMS admin password.
        :param pulumi.Input[str] admin_username: FortiClient EMS admin username.
        :param pulumi.Input[int] call_timeout: FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
        :param pulumi.Input[str] certificate: FortiClient EMS certificate.
        :param pulumi.Input[str] cloud_server_type: Cloud server type. Valid values: `production`, `alpha`, `beta`.
        :param pulumi.Input[str] fortinetone_cloud_authentication: Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] https_port: FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        :param pulumi.Input[str] name: FortiClient Enterprise Management Server (EMS) name.
        :param pulumi.Input[str] pull_avatars: Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pull_sysinfo: Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pull_tags: Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pull_vulnerabilities: Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] serial_number: FortiClient EMS Serial Number.
        :param pulumi.Input[str] server: FortiClient EMS FQDN or IPv4 address.
        :param pulumi.Input[str] source_ip: REST API call source IP.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if admin_password is not None:
            pulumi.set(__self__, "admin_password", admin_password)
        if admin_username is not None:
            pulumi.set(__self__, "admin_username", admin_username)
        if call_timeout is not None:
            pulumi.set(__self__, "call_timeout", call_timeout)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if cloud_server_type is not None:
            pulumi.set(__self__, "cloud_server_type", cloud_server_type)
        if fortinetone_cloud_authentication is not None:
            pulumi.set(__self__, "fortinetone_cloud_authentication", fortinetone_cloud_authentication)
        if https_port is not None:
            pulumi.set(__self__, "https_port", https_port)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pull_avatars is not None:
            pulumi.set(__self__, "pull_avatars", pull_avatars)
        if pull_sysinfo is not None:
            pulumi.set(__self__, "pull_sysinfo", pull_sysinfo)
        if pull_tags is not None:
            pulumi.set(__self__, "pull_tags", pull_tags)
        if pull_vulnerabilities is not None:
            pulumi.set(__self__, "pull_vulnerabilities", pull_vulnerabilities)
        if serial_number is not None:
            pulumi.set(__self__, "serial_number", serial_number)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="adminPassword")
    def admin_password(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS admin password.
        """
        return pulumi.get(self, "admin_password")

    @admin_password.setter
    def admin_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_password", value)

    @property
    @pulumi.getter(name="adminUsername")
    def admin_username(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS admin username.
        """
        return pulumi.get(self, "admin_username")

    @admin_username.setter
    def admin_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_username", value)

    @property
    @pulumi.getter(name="callTimeout")
    def call_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
        """
        return pulumi.get(self, "call_timeout")

    @call_timeout.setter
    def call_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "call_timeout", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS certificate.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="cloudServerType")
    def cloud_server_type(self) -> Optional[pulumi.Input[str]]:
        """
        Cloud server type. Valid values: `production`, `alpha`, `beta`.
        """
        return pulumi.get(self, "cloud_server_type")

    @cloud_server_type.setter
    def cloud_server_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_server_type", value)

    @property
    @pulumi.getter(name="fortinetoneCloudAuthentication")
    def fortinetone_cloud_authentication(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fortinetone_cloud_authentication")

    @fortinetone_cloud_authentication.setter
    def fortinetone_cloud_authentication(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortinetone_cloud_authentication", value)

    @property
    @pulumi.getter(name="httpsPort")
    def https_port(self) -> Optional[pulumi.Input[int]]:
        """
        FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        """
        return pulumi.get(self, "https_port")

    @https_port.setter
    def https_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "https_port", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient Enterprise Management Server (EMS) name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pullAvatars")
    def pull_avatars(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "pull_avatars")

    @pull_avatars.setter
    def pull_avatars(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pull_avatars", value)

    @property
    @pulumi.getter(name="pullSysinfo")
    def pull_sysinfo(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "pull_sysinfo")

    @pull_sysinfo.setter
    def pull_sysinfo(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pull_sysinfo", value)

    @property
    @pulumi.getter(name="pullTags")
    def pull_tags(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "pull_tags")

    @pull_tags.setter
    def pull_tags(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pull_tags", value)

    @property
    @pulumi.getter(name="pullVulnerabilities")
    def pull_vulnerabilities(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "pull_vulnerabilities")

    @pull_vulnerabilities.setter
    def pull_vulnerabilities(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pull_vulnerabilities", value)

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS Serial Number.
        """
        return pulumi.get(self, "serial_number")

    @serial_number.setter
    def serial_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "serial_number", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        FortiClient EMS FQDN or IPv4 address.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        REST API call source IP.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class EndpointControlFctems(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_password: Optional[pulumi.Input[str]] = None,
                 admin_username: Optional[pulumi.Input[str]] = None,
                 call_timeout: Optional[pulumi.Input[int]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 cloud_server_type: Optional[pulumi.Input[str]] = None,
                 fortinetone_cloud_authentication: Optional[pulumi.Input[str]] = None,
                 https_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pull_avatars: Optional[pulumi.Input[str]] = None,
                 pull_sysinfo: Optional[pulumi.Input[str]] = None,
                 pull_tags: Optional[pulumi.Input[str]] = None,
                 pull_vulnerabilities: Optional[pulumi.Input[str]] = None,
                 serial_number: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiClient Enterprise Management Server (EMS) entries.

        ## Import

        EndpointControl Fctems can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/endpointControlFctems:EndpointControlFctems labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_password: FortiClient EMS admin password.
        :param pulumi.Input[str] admin_username: FortiClient EMS admin username.
        :param pulumi.Input[int] call_timeout: FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
        :param pulumi.Input[str] certificate: FortiClient EMS certificate.
        :param pulumi.Input[str] cloud_server_type: Cloud server type. Valid values: `production`, `alpha`, `beta`.
        :param pulumi.Input[str] fortinetone_cloud_authentication: Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] https_port: FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        :param pulumi.Input[str] name: FortiClient Enterprise Management Server (EMS) name.
        :param pulumi.Input[str] pull_avatars: Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pull_sysinfo: Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pull_tags: Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pull_vulnerabilities: Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] serial_number: FortiClient EMS Serial Number.
        :param pulumi.Input[str] server: FortiClient EMS FQDN or IPv4 address.
        :param pulumi.Input[str] source_ip: REST API call source IP.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[EndpointControlFctemsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiClient Enterprise Management Server (EMS) entries.

        ## Import

        EndpointControl Fctems can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/endpointControlFctems:EndpointControlFctems labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param EndpointControlFctemsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointControlFctemsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_password: Optional[pulumi.Input[str]] = None,
                 admin_username: Optional[pulumi.Input[str]] = None,
                 call_timeout: Optional[pulumi.Input[int]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 cloud_server_type: Optional[pulumi.Input[str]] = None,
                 fortinetone_cloud_authentication: Optional[pulumi.Input[str]] = None,
                 https_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pull_avatars: Optional[pulumi.Input[str]] = None,
                 pull_sysinfo: Optional[pulumi.Input[str]] = None,
                 pull_tags: Optional[pulumi.Input[str]] = None,
                 pull_vulnerabilities: Optional[pulumi.Input[str]] = None,
                 serial_number: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EndpointControlFctemsArgs.__new__(EndpointControlFctemsArgs)

            __props__.__dict__["admin_password"] = admin_password
            __props__.__dict__["admin_username"] = admin_username
            __props__.__dict__["call_timeout"] = call_timeout
            __props__.__dict__["certificate"] = certificate
            __props__.__dict__["cloud_server_type"] = cloud_server_type
            __props__.__dict__["fortinetone_cloud_authentication"] = fortinetone_cloud_authentication
            __props__.__dict__["https_port"] = https_port
            __props__.__dict__["name"] = name
            __props__.__dict__["pull_avatars"] = pull_avatars
            __props__.__dict__["pull_sysinfo"] = pull_sysinfo
            __props__.__dict__["pull_tags"] = pull_tags
            __props__.__dict__["pull_vulnerabilities"] = pull_vulnerabilities
            __props__.__dict__["serial_number"] = serial_number
            __props__.__dict__["server"] = server
            __props__.__dict__["source_ip"] = source_ip
            __props__.__dict__["vdomparam"] = vdomparam
        super(EndpointControlFctems, __self__).__init__(
            'fortios:index/endpointControlFctems:EndpointControlFctems',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_password: Optional[pulumi.Input[str]] = None,
            admin_username: Optional[pulumi.Input[str]] = None,
            call_timeout: Optional[pulumi.Input[int]] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            cloud_server_type: Optional[pulumi.Input[str]] = None,
            fortinetone_cloud_authentication: Optional[pulumi.Input[str]] = None,
            https_port: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pull_avatars: Optional[pulumi.Input[str]] = None,
            pull_sysinfo: Optional[pulumi.Input[str]] = None,
            pull_tags: Optional[pulumi.Input[str]] = None,
            pull_vulnerabilities: Optional[pulumi.Input[str]] = None,
            serial_number: Optional[pulumi.Input[str]] = None,
            server: Optional[pulumi.Input[str]] = None,
            source_ip: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'EndpointControlFctems':
        """
        Get an existing EndpointControlFctems resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_password: FortiClient EMS admin password.
        :param pulumi.Input[str] admin_username: FortiClient EMS admin username.
        :param pulumi.Input[int] call_timeout: FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
        :param pulumi.Input[str] certificate: FortiClient EMS certificate.
        :param pulumi.Input[str] cloud_server_type: Cloud server type. Valid values: `production`, `alpha`, `beta`.
        :param pulumi.Input[str] fortinetone_cloud_authentication: Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] https_port: FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        :param pulumi.Input[str] name: FortiClient Enterprise Management Server (EMS) name.
        :param pulumi.Input[str] pull_avatars: Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pull_sysinfo: Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pull_tags: Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] pull_vulnerabilities: Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] serial_number: FortiClient EMS Serial Number.
        :param pulumi.Input[str] server: FortiClient EMS FQDN or IPv4 address.
        :param pulumi.Input[str] source_ip: REST API call source IP.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EndpointControlFctemsState.__new__(_EndpointControlFctemsState)

        __props__.__dict__["admin_password"] = admin_password
        __props__.__dict__["admin_username"] = admin_username
        __props__.__dict__["call_timeout"] = call_timeout
        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["cloud_server_type"] = cloud_server_type
        __props__.__dict__["fortinetone_cloud_authentication"] = fortinetone_cloud_authentication
        __props__.__dict__["https_port"] = https_port
        __props__.__dict__["name"] = name
        __props__.__dict__["pull_avatars"] = pull_avatars
        __props__.__dict__["pull_sysinfo"] = pull_sysinfo
        __props__.__dict__["pull_tags"] = pull_tags
        __props__.__dict__["pull_vulnerabilities"] = pull_vulnerabilities
        __props__.__dict__["serial_number"] = serial_number
        __props__.__dict__["server"] = server
        __props__.__dict__["source_ip"] = source_ip
        __props__.__dict__["vdomparam"] = vdomparam
        return EndpointControlFctems(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminPassword")
    def admin_password(self) -> pulumi.Output[Optional[str]]:
        """
        FortiClient EMS admin password.
        """
        return pulumi.get(self, "admin_password")

    @property
    @pulumi.getter(name="adminUsername")
    def admin_username(self) -> pulumi.Output[str]:
        """
        FortiClient EMS admin username.
        """
        return pulumi.get(self, "admin_username")

    @property
    @pulumi.getter(name="callTimeout")
    def call_timeout(self) -> pulumi.Output[int]:
        """
        FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
        """
        return pulumi.get(self, "call_timeout")

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[str]:
        """
        FortiClient EMS certificate.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="cloudServerType")
    def cloud_server_type(self) -> pulumi.Output[str]:
        """
        Cloud server type. Valid values: `production`, `alpha`, `beta`.
        """
        return pulumi.get(self, "cloud_server_type")

    @property
    @pulumi.getter(name="fortinetoneCloudAuthentication")
    def fortinetone_cloud_authentication(self) -> pulumi.Output[str]:
        """
        Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fortinetone_cloud_authentication")

    @property
    @pulumi.getter(name="httpsPort")
    def https_port(self) -> pulumi.Output[int]:
        """
        FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
        """
        return pulumi.get(self, "https_port")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        FortiClient Enterprise Management Server (EMS) name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pullAvatars")
    def pull_avatars(self) -> pulumi.Output[str]:
        """
        Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "pull_avatars")

    @property
    @pulumi.getter(name="pullSysinfo")
    def pull_sysinfo(self) -> pulumi.Output[str]:
        """
        Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "pull_sysinfo")

    @property
    @pulumi.getter(name="pullTags")
    def pull_tags(self) -> pulumi.Output[str]:
        """
        Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "pull_tags")

    @property
    @pulumi.getter(name="pullVulnerabilities")
    def pull_vulnerabilities(self) -> pulumi.Output[str]:
        """
        Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "pull_vulnerabilities")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> pulumi.Output[str]:
        """
        FortiClient EMS Serial Number.
        """
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[str]:
        """
        FortiClient EMS FQDN or IPv4 address.
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[str]:
        """
        REST API call source IP.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
