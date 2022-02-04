# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetIPMaskCIDRResult',
    'AwaitableGetIPMaskCIDRResult',
    'get_ip_mask_cidr',
    'get_ip_mask_cidr_output',
]

@pulumi.output_type
class GetIPMaskCIDRResult:
    """
    A collection of values returned by GetIPMaskCIDR.
    """
    def __init__(__self__, cidr=None, cidrlists=None, id=None, ipmask=None, ipmasklists=None):
        if cidr and not isinstance(cidr, str):
            raise TypeError("Expected argument 'cidr' to be a str")
        pulumi.set(__self__, "cidr", cidr)
        if cidrlists and not isinstance(cidrlists, list):
            raise TypeError("Expected argument 'cidrlists' to be a list")
        pulumi.set(__self__, "cidrlists", cidrlists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ipmask and not isinstance(ipmask, str):
            raise TypeError("Expected argument 'ipmask' to be a str")
        pulumi.set(__self__, "ipmask", ipmask)
        if ipmasklists and not isinstance(ipmasklists, list):
            raise TypeError("Expected argument 'ipmasklists' to be a list")
        pulumi.set(__self__, "ipmasklists", ipmasklists)

    @property
    @pulumi.getter
    def cidr(self) -> str:
        """
        Classless Inter-Domain Routing of the IP/MASK.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def cidrlists(self) -> Sequence[str]:
        """
        Classless Inter-Domain Routing list converted from the IP/MASK list.
        """
        return pulumi.get(self, "cidrlists")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ipmask(self) -> Optional[str]:
        """
        IP/MASK.
        """
        return pulumi.get(self, "ipmask")

    @property
    @pulumi.getter
    def ipmasklists(self) -> Optional[Sequence[str]]:
        """
        IP/MASK list.
        """
        return pulumi.get(self, "ipmasklists")


class AwaitableGetIPMaskCIDRResult(GetIPMaskCIDRResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIPMaskCIDRResult(
            cidr=self.cidr,
            cidrlists=self.cidrlists,
            id=self.id,
            ipmask=self.ipmask,
            ipmasklists=self.ipmasklists)


def get_ip_mask_cidr(ipmask: Optional[str] = None,
                     ipmasklists: Optional[Sequence[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIPMaskCIDRResult:
    """
    Convert IP/Mask to CIDR

    ## Example Usage
    ### Example1

    ```python
    import pulumi
    import pulumi_fortios as fortios

    trname_get_system_interface = fortios.get_system_interface(name="port3")
    trname_get_ip_mask_cidr = fortios.get_ip_mask_cidr(ipmask=trname_get_system_interface.ip)
    pulumi.export("output1", trname_get_ip_mask_cidr.cidr)
    ```
    ### Example2

    ```python
    import pulumi
    import pulumi_fortios as fortios

    trname_get_system_interface = fortios.get_system_interface(name="port3")
    trname_get_ip_mask_cidr = fortios.get_ip_mask_cidr(ipmask=trname_get_system_interface.ip,
        ipmasklists=[
            "21.1.1.1 255.255.255.0",
            "22.1.1.1 255.255.255.240",
            "23.1.1.1 255.255.255.224",
        ])
    pulumi.export("outputConv1", trname_get_ip_mask_cidr.cidr)
    pulumi.export("outputConv2", trname_get_ip_mask_cidr.cidrlists)
    pulumi.export("outputOrignal", trname_get_system_interface.ip)
    ```


    :param str ipmask: Specify IP/MASK.
    :param Sequence[str] ipmasklists: Specify IP/MASK list.
    """
    __args__ = dict()
    __args__['ipmask'] = ipmask
    __args__['ipmasklists'] = ipmasklists
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('fortios:index/getIPMaskCIDR:GetIPMaskCIDR', __args__, opts=opts, typ=GetIPMaskCIDRResult).value

    return AwaitableGetIPMaskCIDRResult(
        cidr=__ret__.cidr,
        cidrlists=__ret__.cidrlists,
        id=__ret__.id,
        ipmask=__ret__.ipmask,
        ipmasklists=__ret__.ipmasklists)


@_utilities.lift_output_func(get_ip_mask_cidr)
def get_ip_mask_cidr_output(ipmask: Optional[pulumi.Input[Optional[str]]] = None,
                            ipmasklists: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIPMaskCIDRResult]:
    """
    Convert IP/Mask to CIDR

    ## Example Usage
    ### Example1

    ```python
    import pulumi
    import pulumi_fortios as fortios

    trname_get_system_interface = fortios.get_system_interface(name="port3")
    trname_get_ip_mask_cidr = fortios.get_ip_mask_cidr(ipmask=trname_get_system_interface.ip)
    pulumi.export("output1", trname_get_ip_mask_cidr.cidr)
    ```
    ### Example2

    ```python
    import pulumi
    import pulumi_fortios as fortios

    trname_get_system_interface = fortios.get_system_interface(name="port3")
    trname_get_ip_mask_cidr = fortios.get_ip_mask_cidr(ipmask=trname_get_system_interface.ip,
        ipmasklists=[
            "21.1.1.1 255.255.255.0",
            "22.1.1.1 255.255.255.240",
            "23.1.1.1 255.255.255.224",
        ])
    pulumi.export("outputConv1", trname_get_ip_mask_cidr.cidr)
    pulumi.export("outputConv2", trname_get_ip_mask_cidr.cidrlists)
    pulumi.export("outputOrignal", trname_get_system_interface.ip)
    ```


    :param str ipmask: Specify IP/MASK.
    :param Sequence[str] ipmasklists: Specify IP/MASK list.
    """
    ...
