# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemAutoupdateTunnelingResult',
    'AwaitableGetSystemAutoupdateTunnelingResult',
    'get_system_autoupdate_tunneling',
]

@pulumi.output_type
class GetSystemAutoupdateTunnelingResult:
    """
    A collection of values returned by GetSystemAutoupdateTunneling.
    """
    def __init__(__self__, address=None, id=None, password=None, port=None, status=None, username=None, vdomparam=None):
        if address and not isinstance(address, str):
            raise TypeError("Expected argument 'address' to be a str")
        pulumi.set(__self__, "address", address)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def address(self) -> str:
        """
        Web proxy IP address or FQDN.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        Web proxy password.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        Web proxy port.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable web proxy tunnelling.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        Web proxy username.
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemAutoupdateTunnelingResult(GetSystemAutoupdateTunnelingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemAutoupdateTunnelingResult(
            address=self.address,
            id=self.id,
            password=self.password,
            port=self.port,
            status=self.status,
            username=self.username,
            vdomparam=self.vdomparam)


def get_system_autoupdate_tunneling(vdomparam: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemAutoupdateTunnelingResult:
    """
    Use this data source to get information on fortios systemautoupdate tunneling


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemAutoupdateTunneling:GetSystemAutoupdateTunneling', __args__, opts=opts, typ=GetSystemAutoupdateTunnelingResult).value

    return AwaitableGetSystemAutoupdateTunnelingResult(
        address=__ret__.address,
        id=__ret__.id,
        password=__ret__.password,
        port=__ret__.port,
        status=__ret__.status,
        username=__ret__.username,
        vdomparam=__ret__.vdomparam)