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
    'GetSystemAutomationDestinationResult',
    'AwaitableGetSystemAutomationDestinationResult',
    'get_system_automation_destination',
]

@pulumi.output_type
class GetSystemAutomationDestinationResult:
    """
    A collection of values returned by GetSystemAutomationDestination.
    """
    def __init__(__self__, destinations=None, ha_group_id=None, id=None, name=None, type=None, vdomparam=None):
        if destinations and not isinstance(destinations, list):
            raise TypeError("Expected argument 'destinations' to be a list")
        pulumi.set(__self__, "destinations", destinations)
        if ha_group_id and not isinstance(ha_group_id, int):
            raise TypeError("Expected argument 'ha_group_id' to be a int")
        pulumi.set(__self__, "ha_group_id", ha_group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def destinations(self) -> Sequence['outputs.GetSystemAutomationDestinationDestinationResult']:
        """
        Destinations. The structure of `destination` block is documented below.
        """
        return pulumi.get(self, "destinations")

    @property
    @pulumi.getter(name="haGroupId")
    def ha_group_id(self) -> int:
        """
        Cluster group ID set for this destination (default = 0).
        """
        return pulumi.get(self, "ha_group_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Destination.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Destination type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemAutomationDestinationResult(GetSystemAutomationDestinationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemAutomationDestinationResult(
            destinations=self.destinations,
            ha_group_id=self.ha_group_id,
            id=self.id,
            name=self.name,
            type=self.type,
            vdomparam=self.vdomparam)


def get_system_automation_destination(name: Optional[str] = None,
                                      vdomparam: Optional[str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemAutomationDestinationResult:
    """
    Use this data source to get information on an fortios system automationdestination


    :param str name: Specify the name of the desired system automationdestination.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemAutomationDestination:GetSystemAutomationDestination', __args__, opts=opts, typ=GetSystemAutomationDestinationResult).value

    return AwaitableGetSystemAutomationDestinationResult(
        destinations=__ret__.destinations,
        ha_group_id=__ret__.ha_group_id,
        id=__ret__.id,
        name=__ret__.name,
        type=__ret__.type,
        vdomparam=__ret__.vdomparam)