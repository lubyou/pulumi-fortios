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
    'GetSystemDdnsResult',
    'AwaitableGetSystemDdnsResult',
    'get_system_ddns',
    'get_system_ddns_output',
]

@pulumi.output_type
class GetSystemDdnsResult:
    """
    A collection of values returned by GetSystemDdns.
    """
    def __init__(__self__, addr_type=None, bound_ip=None, clear_text=None, ddns_auth=None, ddns_domain=None, ddns_key=None, ddns_keyname=None, ddns_password=None, ddns_server=None, ddns_server_addrs=None, ddns_server_ip=None, ddns_sn=None, ddns_ttl=None, ddns_username=None, ddns_zone=None, ddnsid=None, id=None, monitor_interfaces=None, server_type=None, ssl_certificate=None, update_interval=None, use_public_ip=None, vdomparam=None):
        if addr_type and not isinstance(addr_type, str):
            raise TypeError("Expected argument 'addr_type' to be a str")
        pulumi.set(__self__, "addr_type", addr_type)
        if bound_ip and not isinstance(bound_ip, str):
            raise TypeError("Expected argument 'bound_ip' to be a str")
        pulumi.set(__self__, "bound_ip", bound_ip)
        if clear_text and not isinstance(clear_text, str):
            raise TypeError("Expected argument 'clear_text' to be a str")
        pulumi.set(__self__, "clear_text", clear_text)
        if ddns_auth and not isinstance(ddns_auth, str):
            raise TypeError("Expected argument 'ddns_auth' to be a str")
        pulumi.set(__self__, "ddns_auth", ddns_auth)
        if ddns_domain and not isinstance(ddns_domain, str):
            raise TypeError("Expected argument 'ddns_domain' to be a str")
        pulumi.set(__self__, "ddns_domain", ddns_domain)
        if ddns_key and not isinstance(ddns_key, str):
            raise TypeError("Expected argument 'ddns_key' to be a str")
        pulumi.set(__self__, "ddns_key", ddns_key)
        if ddns_keyname and not isinstance(ddns_keyname, str):
            raise TypeError("Expected argument 'ddns_keyname' to be a str")
        pulumi.set(__self__, "ddns_keyname", ddns_keyname)
        if ddns_password and not isinstance(ddns_password, str):
            raise TypeError("Expected argument 'ddns_password' to be a str")
        pulumi.set(__self__, "ddns_password", ddns_password)
        if ddns_server and not isinstance(ddns_server, str):
            raise TypeError("Expected argument 'ddns_server' to be a str")
        pulumi.set(__self__, "ddns_server", ddns_server)
        if ddns_server_addrs and not isinstance(ddns_server_addrs, list):
            raise TypeError("Expected argument 'ddns_server_addrs' to be a list")
        pulumi.set(__self__, "ddns_server_addrs", ddns_server_addrs)
        if ddns_server_ip and not isinstance(ddns_server_ip, str):
            raise TypeError("Expected argument 'ddns_server_ip' to be a str")
        pulumi.set(__self__, "ddns_server_ip", ddns_server_ip)
        if ddns_sn and not isinstance(ddns_sn, str):
            raise TypeError("Expected argument 'ddns_sn' to be a str")
        pulumi.set(__self__, "ddns_sn", ddns_sn)
        if ddns_ttl and not isinstance(ddns_ttl, int):
            raise TypeError("Expected argument 'ddns_ttl' to be a int")
        pulumi.set(__self__, "ddns_ttl", ddns_ttl)
        if ddns_username and not isinstance(ddns_username, str):
            raise TypeError("Expected argument 'ddns_username' to be a str")
        pulumi.set(__self__, "ddns_username", ddns_username)
        if ddns_zone and not isinstance(ddns_zone, str):
            raise TypeError("Expected argument 'ddns_zone' to be a str")
        pulumi.set(__self__, "ddns_zone", ddns_zone)
        if ddnsid and not isinstance(ddnsid, int):
            raise TypeError("Expected argument 'ddnsid' to be a int")
        pulumi.set(__self__, "ddnsid", ddnsid)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if monitor_interfaces and not isinstance(monitor_interfaces, list):
            raise TypeError("Expected argument 'monitor_interfaces' to be a list")
        pulumi.set(__self__, "monitor_interfaces", monitor_interfaces)
        if server_type and not isinstance(server_type, str):
            raise TypeError("Expected argument 'server_type' to be a str")
        pulumi.set(__self__, "server_type", server_type)
        if ssl_certificate and not isinstance(ssl_certificate, str):
            raise TypeError("Expected argument 'ssl_certificate' to be a str")
        pulumi.set(__self__, "ssl_certificate", ssl_certificate)
        if update_interval and not isinstance(update_interval, int):
            raise TypeError("Expected argument 'update_interval' to be a int")
        pulumi.set(__self__, "update_interval", update_interval)
        if use_public_ip and not isinstance(use_public_ip, str):
            raise TypeError("Expected argument 'use_public_ip' to be a str")
        pulumi.set(__self__, "use_public_ip", use_public_ip)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="addrType")
    def addr_type(self) -> str:
        """
        Address type of interface address in DDNS update.
        """
        return pulumi.get(self, "addr_type")

    @property
    @pulumi.getter(name="boundIp")
    def bound_ip(self) -> str:
        """
        Bound IP address.
        """
        return pulumi.get(self, "bound_ip")

    @property
    @pulumi.getter(name="clearText")
    def clear_text(self) -> str:
        """
        Enable/disable use of clear text connections.
        """
        return pulumi.get(self, "clear_text")

    @property
    @pulumi.getter(name="ddnsAuth")
    def ddns_auth(self) -> str:
        """
        Enable/disable TSIG authentication for your DDNS server.
        """
        return pulumi.get(self, "ddns_auth")

    @property
    @pulumi.getter(name="ddnsDomain")
    def ddns_domain(self) -> str:
        """
        Your fully qualified domain name (for example, yourname.DDNS.com).
        """
        return pulumi.get(self, "ddns_domain")

    @property
    @pulumi.getter(name="ddnsKey")
    def ddns_key(self) -> str:
        """
        DDNS update key (base 64 encoding).
        """
        return pulumi.get(self, "ddns_key")

    @property
    @pulumi.getter(name="ddnsKeyname")
    def ddns_keyname(self) -> str:
        """
        DDNS update key name.
        """
        return pulumi.get(self, "ddns_keyname")

    @property
    @pulumi.getter(name="ddnsPassword")
    def ddns_password(self) -> str:
        """
        DDNS password.
        """
        return pulumi.get(self, "ddns_password")

    @property
    @pulumi.getter(name="ddnsServer")
    def ddns_server(self) -> str:
        """
        Select a DDNS service provider.
        """
        return pulumi.get(self, "ddns_server")

    @property
    @pulumi.getter(name="ddnsServerAddrs")
    def ddns_server_addrs(self) -> Sequence['outputs.GetSystemDdnsDdnsServerAddrResult']:
        """
        Generic DDNS server IP/FQDN list. The structure of `ddns_server_addr` block is documented below.
        """
        return pulumi.get(self, "ddns_server_addrs")

    @property
    @pulumi.getter(name="ddnsServerIp")
    def ddns_server_ip(self) -> str:
        """
        Generic DDNS server IP.
        """
        return pulumi.get(self, "ddns_server_ip")

    @property
    @pulumi.getter(name="ddnsSn")
    def ddns_sn(self) -> str:
        """
        DDNS Serial Number.
        """
        return pulumi.get(self, "ddns_sn")

    @property
    @pulumi.getter(name="ddnsTtl")
    def ddns_ttl(self) -> int:
        """
        Time-to-live for DDNS packets.
        """
        return pulumi.get(self, "ddns_ttl")

    @property
    @pulumi.getter(name="ddnsUsername")
    def ddns_username(self) -> str:
        """
        DDNS user name.
        """
        return pulumi.get(self, "ddns_username")

    @property
    @pulumi.getter(name="ddnsZone")
    def ddns_zone(self) -> str:
        """
        Zone of your domain name (for example, DDNS.com).
        """
        return pulumi.get(self, "ddns_zone")

    @property
    @pulumi.getter
    def ddnsid(self) -> int:
        """
        DDNS ID.
        """
        return pulumi.get(self, "ddnsid")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="monitorInterfaces")
    def monitor_interfaces(self) -> Sequence['outputs.GetSystemDdnsMonitorInterfaceResult']:
        """
        Monitored interface. The structure of `monitor_interface` block is documented below.
        """
        return pulumi.get(self, "monitor_interfaces")

    @property
    @pulumi.getter(name="serverType")
    def server_type(self) -> str:
        """
        Address type of the DDNS server.
        """
        return pulumi.get(self, "server_type")

    @property
    @pulumi.getter(name="sslCertificate")
    def ssl_certificate(self) -> str:
        """
        Name of local certificate for SSL connections.
        """
        return pulumi.get(self, "ssl_certificate")

    @property
    @pulumi.getter(name="updateInterval")
    def update_interval(self) -> int:
        """
        DDNS update interval (60 - 2592000 sec, default = 300).
        """
        return pulumi.get(self, "update_interval")

    @property
    @pulumi.getter(name="usePublicIp")
    def use_public_ip(self) -> str:
        """
        Enable/disable use of public IP address.
        """
        return pulumi.get(self, "use_public_ip")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemDdnsResult(GetSystemDdnsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemDdnsResult(
            addr_type=self.addr_type,
            bound_ip=self.bound_ip,
            clear_text=self.clear_text,
            ddns_auth=self.ddns_auth,
            ddns_domain=self.ddns_domain,
            ddns_key=self.ddns_key,
            ddns_keyname=self.ddns_keyname,
            ddns_password=self.ddns_password,
            ddns_server=self.ddns_server,
            ddns_server_addrs=self.ddns_server_addrs,
            ddns_server_ip=self.ddns_server_ip,
            ddns_sn=self.ddns_sn,
            ddns_ttl=self.ddns_ttl,
            ddns_username=self.ddns_username,
            ddns_zone=self.ddns_zone,
            ddnsid=self.ddnsid,
            id=self.id,
            monitor_interfaces=self.monitor_interfaces,
            server_type=self.server_type,
            ssl_certificate=self.ssl_certificate,
            update_interval=self.update_interval,
            use_public_ip=self.use_public_ip,
            vdomparam=self.vdomparam)


def get_system_ddns(ddnsid: Optional[int] = None,
                    vdomparam: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemDdnsResult:
    """
    Use this data source to get information on an fortios system ddns


    :param int ddnsid: Specify the ddnsid of the desired system ddns.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['ddnsid'] = ddnsid
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemDdns:GetSystemDdns', __args__, opts=opts, typ=GetSystemDdnsResult).value

    return AwaitableGetSystemDdnsResult(
        addr_type=__ret__.addr_type,
        bound_ip=__ret__.bound_ip,
        clear_text=__ret__.clear_text,
        ddns_auth=__ret__.ddns_auth,
        ddns_domain=__ret__.ddns_domain,
        ddns_key=__ret__.ddns_key,
        ddns_keyname=__ret__.ddns_keyname,
        ddns_password=__ret__.ddns_password,
        ddns_server=__ret__.ddns_server,
        ddns_server_addrs=__ret__.ddns_server_addrs,
        ddns_server_ip=__ret__.ddns_server_ip,
        ddns_sn=__ret__.ddns_sn,
        ddns_ttl=__ret__.ddns_ttl,
        ddns_username=__ret__.ddns_username,
        ddns_zone=__ret__.ddns_zone,
        ddnsid=__ret__.ddnsid,
        id=__ret__.id,
        monitor_interfaces=__ret__.monitor_interfaces,
        server_type=__ret__.server_type,
        ssl_certificate=__ret__.ssl_certificate,
        update_interval=__ret__.update_interval,
        use_public_ip=__ret__.use_public_ip,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_ddns)
def get_system_ddns_output(ddnsid: Optional[pulumi.Input[int]] = None,
                           vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemDdnsResult]:
    """
    Use this data source to get information on an fortios system ddns


    :param int ddnsid: Specify the ddnsid of the desired system ddns.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
