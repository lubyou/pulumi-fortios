# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemFssoPollingResult',
    'AwaitableGetSystemFssoPollingResult',
    'get_system_fsso_polling',
    'get_system_fsso_polling_output',
]

@pulumi.output_type
class GetSystemFssoPollingResult:
    """
    A collection of values returned by GetSystemFssoPolling.
    """
    def __init__(__self__, auth_password=None, authentication=None, id=None, listening_port=None, status=None, vdomparam=None):
        if auth_password and not isinstance(auth_password, str):
            raise TypeError("Expected argument 'auth_password' to be a str")
        pulumi.set(__self__, "auth_password", auth_password)
        if authentication and not isinstance(authentication, str):
            raise TypeError("Expected argument 'authentication' to be a str")
        pulumi.set(__self__, "authentication", authentication)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if listening_port and not isinstance(listening_port, int):
            raise TypeError("Expected argument 'listening_port' to be a int")
        pulumi.set(__self__, "listening_port", listening_port)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="authPassword")
    def auth_password(self) -> str:
        """
        Password to connect to FSSO Agent.
        """
        return pulumi.get(self, "auth_password")

    @property
    @pulumi.getter
    def authentication(self) -> str:
        """
        Enable/disable FSSO Agent Authentication.
        """
        return pulumi.get(self, "authentication")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="listeningPort")
    def listening_port(self) -> int:
        """
        Listening port to accept clients (1 - 65535).
        """
        return pulumi.get(self, "listening_port")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable FSSO Polling Mode.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemFssoPollingResult(GetSystemFssoPollingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemFssoPollingResult(
            auth_password=self.auth_password,
            authentication=self.authentication,
            id=self.id,
            listening_port=self.listening_port,
            status=self.status,
            vdomparam=self.vdomparam)


def get_system_fsso_polling(vdomparam: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemFssoPollingResult:
    """
    Use this data source to get information on fortios system fssopolling


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemFssoPolling:GetSystemFssoPolling', __args__, opts=opts, typ=GetSystemFssoPollingResult).value

    return AwaitableGetSystemFssoPollingResult(
        auth_password=__ret__.auth_password,
        authentication=__ret__.authentication,
        id=__ret__.id,
        listening_port=__ret__.listening_port,
        status=__ret__.status,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_fsso_polling)
def get_system_fsso_polling_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemFssoPollingResult]:
    """
    Use this data source to get information on fortios system fssopolling


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
