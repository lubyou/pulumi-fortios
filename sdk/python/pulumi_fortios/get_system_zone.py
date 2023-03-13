# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetSystemZoneResult',
    'AwaitableGetSystemZoneResult',
    'get_system_zone',
    'get_system_zone_output',
]

@pulumi.output_type
class GetSystemZoneResult:
    """
    A collection of values returned by GetSystemZone.
    """
    def __init__(__self__, description=None, id=None, interfaces=None, intrazone=None, name=None, taggings=None, vdomparam=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interfaces and not isinstance(interfaces, list):
            raise TypeError("Expected argument 'interfaces' to be a list")
        pulumi.set(__self__, "interfaces", interfaces)
        if intrazone and not isinstance(intrazone, str):
            raise TypeError("Expected argument 'intrazone' to be a str")
        pulumi.set(__self__, "intrazone", intrazone)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if taggings and not isinstance(taggings, list):
            raise TypeError("Expected argument 'taggings' to be a list")
        pulumi.set(__self__, "taggings", taggings)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interfaces(self) -> Sequence['outputs.GetSystemZoneInterfaceResult']:
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter
    def intrazone(self) -> str:
        return pulumi.get(self, "intrazone")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def taggings(self) -> Sequence['outputs.GetSystemZoneTaggingResult']:
        return pulumi.get(self, "taggings")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemZoneResult(GetSystemZoneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemZoneResult(
            description=self.description,
            id=self.id,
            interfaces=self.interfaces,
            intrazone=self.intrazone,
            name=self.name,
            taggings=self.taggings,
            vdomparam=self.vdomparam)


def get_system_zone(name: Optional[str] = None,
                    vdomparam: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemZoneResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemZone:GetSystemZone', __args__, opts=opts, typ=GetSystemZoneResult).value

    return AwaitableGetSystemZoneResult(
        description=__ret__.description,
        id=__ret__.id,
        interfaces=__ret__.interfaces,
        intrazone=__ret__.intrazone,
        name=__ret__.name,
        taggings=__ret__.taggings,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_system_zone)
def get_system_zone_output(name: Optional[pulumi.Input[str]] = None,
                           vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemZoneResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
