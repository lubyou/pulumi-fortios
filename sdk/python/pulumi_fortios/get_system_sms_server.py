# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemSmsServerResult',
    'AwaitableGetSystemSmsServerResult',
    'get_system_sms_server',
    'get_system_sms_server_output',
]

@pulumi.output_type
class GetSystemSmsServerResult:
    """
    A collection of values returned by GetSystemSmsServer.
    """
    def __init__(__self__, id=None, mail_server=None, name=None, vdomparam=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mail_server and not isinstance(mail_server, str):
            raise TypeError("Expected argument 'mail_server' to be a str")
        pulumi.set(__self__, "mail_server", mail_server)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="mailServer")
    def mail_server(self) -> str:
        return pulumi.get(self, "mail_server")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemSmsServerResult(GetSystemSmsServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemSmsServerResult(
            id=self.id,
            mail_server=self.mail_server,
            name=self.name,
            vdomparam=self.vdomparam)


def get_system_sms_server(name: Optional[str] = None,
                          vdomparam: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemSmsServerResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemSmsServer:GetSystemSmsServer', __args__, opts=opts, typ=GetSystemSmsServerResult).value

    return AwaitableGetSystemSmsServerResult(
        id=pulumi.get(__ret__, 'id'),
        mail_server=pulumi.get(__ret__, 'mail_server'),
        name=pulumi.get(__ret__, 'name'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_system_sms_server)
def get_system_sms_server_output(name: Optional[pulumi.Input[str]] = None,
                                 vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemSmsServerResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
