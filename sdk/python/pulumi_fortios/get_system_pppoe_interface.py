# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemPppoeInterfaceResult',
    'AwaitableGetSystemPppoeInterfaceResult',
    'get_system_pppoe_interface',
    'get_system_pppoe_interface_output',
]

@pulumi.output_type
class GetSystemPppoeInterfaceResult:
    """
    A collection of values returned by GetSystemPppoeInterface.
    """
    def __init__(__self__, ac_name=None, auth_type=None, device=None, dial_on_demand=None, disc_retry_timeout=None, id=None, idle_timeout=None, ipunnumbered=None, ipv6=None, lcp_echo_interval=None, lcp_max_echo_fails=None, name=None, padt_retry_timeout=None, password=None, pppoe_unnumbered_negotiate=None, service_name=None, username=None, vdomparam=None):
        if ac_name and not isinstance(ac_name, str):
            raise TypeError("Expected argument 'ac_name' to be a str")
        pulumi.set(__self__, "ac_name", ac_name)
        if auth_type and not isinstance(auth_type, str):
            raise TypeError("Expected argument 'auth_type' to be a str")
        pulumi.set(__self__, "auth_type", auth_type)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if dial_on_demand and not isinstance(dial_on_demand, str):
            raise TypeError("Expected argument 'dial_on_demand' to be a str")
        pulumi.set(__self__, "dial_on_demand", dial_on_demand)
        if disc_retry_timeout and not isinstance(disc_retry_timeout, int):
            raise TypeError("Expected argument 'disc_retry_timeout' to be a int")
        pulumi.set(__self__, "disc_retry_timeout", disc_retry_timeout)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if idle_timeout and not isinstance(idle_timeout, int):
            raise TypeError("Expected argument 'idle_timeout' to be a int")
        pulumi.set(__self__, "idle_timeout", idle_timeout)
        if ipunnumbered and not isinstance(ipunnumbered, str):
            raise TypeError("Expected argument 'ipunnumbered' to be a str")
        pulumi.set(__self__, "ipunnumbered", ipunnumbered)
        if ipv6 and not isinstance(ipv6, str):
            raise TypeError("Expected argument 'ipv6' to be a str")
        pulumi.set(__self__, "ipv6", ipv6)
        if lcp_echo_interval and not isinstance(lcp_echo_interval, int):
            raise TypeError("Expected argument 'lcp_echo_interval' to be a int")
        pulumi.set(__self__, "lcp_echo_interval", lcp_echo_interval)
        if lcp_max_echo_fails and not isinstance(lcp_max_echo_fails, int):
            raise TypeError("Expected argument 'lcp_max_echo_fails' to be a int")
        pulumi.set(__self__, "lcp_max_echo_fails", lcp_max_echo_fails)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if padt_retry_timeout and not isinstance(padt_retry_timeout, int):
            raise TypeError("Expected argument 'padt_retry_timeout' to be a int")
        pulumi.set(__self__, "padt_retry_timeout", padt_retry_timeout)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if pppoe_unnumbered_negotiate and not isinstance(pppoe_unnumbered_negotiate, str):
            raise TypeError("Expected argument 'pppoe_unnumbered_negotiate' to be a str")
        pulumi.set(__self__, "pppoe_unnumbered_negotiate", pppoe_unnumbered_negotiate)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="acName")
    def ac_name(self) -> str:
        """
        PPPoE AC name.
        """
        return pulumi.get(self, "ac_name")

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> str:
        """
        PPP authentication type to use.
        """
        return pulumi.get(self, "auth_type")

    @property
    @pulumi.getter
    def device(self) -> str:
        """
        Name for the physical interface.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="dialOnDemand")
    def dial_on_demand(self) -> str:
        """
        Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface.
        """
        return pulumi.get(self, "dial_on_demand")

    @property
    @pulumi.getter(name="discRetryTimeout")
    def disc_retry_timeout(self) -> int:
        """
        PPPoE discovery init timeout value in (0-4294967295 sec).
        """
        return pulumi.get(self, "disc_retry_timeout")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idleTimeout")
    def idle_timeout(self) -> int:
        """
        PPPoE auto disconnect after idle timeout (0-4294967295 sec).
        """
        return pulumi.get(self, "idle_timeout")

    @property
    @pulumi.getter
    def ipunnumbered(self) -> str:
        """
        PPPoE unnumbered IP.
        """
        return pulumi.get(self, "ipunnumbered")

    @property
    @pulumi.getter
    def ipv6(self) -> str:
        """
        Enable/disable IPv6 Control Protocol (IPv6CP).
        """
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter(name="lcpEchoInterval")
    def lcp_echo_interval(self) -> int:
        """
        PPPoE LCP echo interval in (0-4294967295 sec, default = 5).
        """
        return pulumi.get(self, "lcp_echo_interval")

    @property
    @pulumi.getter(name="lcpMaxEchoFails")
    def lcp_max_echo_fails(self) -> int:
        """
        Maximum missed LCP echo messages before disconnect (0-4294967295, default = 3).
        """
        return pulumi.get(self, "lcp_max_echo_fails")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the PPPoE interface.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="padtRetryTimeout")
    def padt_retry_timeout(self) -> int:
        """
        PPPoE terminate timeout value in (0-4294967295 sec).
        """
        return pulumi.get(self, "padt_retry_timeout")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        Enter the password.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="pppoeUnnumberedNegotiate")
    def pppoe_unnumbered_negotiate(self) -> str:
        """
        Enable/disable PPPoE unnumbered negotiation.
        """
        return pulumi.get(self, "pppoe_unnumbered_negotiate")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        PPPoE service name.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        User name.
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemPppoeInterfaceResult(GetSystemPppoeInterfaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemPppoeInterfaceResult(
            ac_name=self.ac_name,
            auth_type=self.auth_type,
            device=self.device,
            dial_on_demand=self.dial_on_demand,
            disc_retry_timeout=self.disc_retry_timeout,
            id=self.id,
            idle_timeout=self.idle_timeout,
            ipunnumbered=self.ipunnumbered,
            ipv6=self.ipv6,
            lcp_echo_interval=self.lcp_echo_interval,
            lcp_max_echo_fails=self.lcp_max_echo_fails,
            name=self.name,
            padt_retry_timeout=self.padt_retry_timeout,
            password=self.password,
            pppoe_unnumbered_negotiate=self.pppoe_unnumbered_negotiate,
            service_name=self.service_name,
            username=self.username,
            vdomparam=self.vdomparam)


def get_system_pppoe_interface(name: Optional[str] = None,
                               vdomparam: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemPppoeInterfaceResult:
    """
    Use this data source to get information on an fortios system pppoeinterface


    :param str name: Specify the name of the desired system pppoeinterface.
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
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemPppoeInterface:GetSystemPppoeInterface', __args__, opts=opts, typ=GetSystemPppoeInterfaceResult).value

    return AwaitableGetSystemPppoeInterfaceResult(
        ac_name=__ret__.ac_name,
        auth_type=__ret__.auth_type,
        device=__ret__.device,
        dial_on_demand=__ret__.dial_on_demand,
        disc_retry_timeout=__ret__.disc_retry_timeout,
        id=__ret__.id,
        idle_timeout=__ret__.idle_timeout,
        ipunnumbered=__ret__.ipunnumbered,
        ipv6=__ret__.ipv6,
        lcp_echo_interval=__ret__.lcp_echo_interval,
        lcp_max_echo_fails=__ret__.lcp_max_echo_fails,
        name=__ret__.name,
        padt_retry_timeout=__ret__.padt_retry_timeout,
        password=__ret__.password,
        pppoe_unnumbered_negotiate=__ret__.pppoe_unnumbered_negotiate,
        service_name=__ret__.service_name,
        username=__ret__.username,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_pppoe_interface)
def get_system_pppoe_interface_output(name: Optional[pulumi.Input[str]] = None,
                                      vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemPppoeInterfaceResult]:
    """
    Use this data source to get information on an fortios system pppoeinterface


    :param str name: Specify the name of the desired system pppoeinterface.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
