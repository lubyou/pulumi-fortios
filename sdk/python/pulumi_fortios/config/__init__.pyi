# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

cabundlecontent: Optional[str]
"""
CA Bundle file content
"""

cabundlefile: Optional[str]
"""
CA Bundle file
"""

cacert: Optional[str]
"""
CA certtificate(Optional)
"""

clientcert: Optional[str]
"""
User certificate
"""

clientkey: Optional[str]
"""
User private key
"""

fmgCabundlefile: Optional[str]
"""
CA Bundle file
"""

fmgHostname: Optional[str]
"""
Hostname/IP address of the FortiManager to connect to
"""

fmgInsecure: Optional[bool]

fmgPasswd: Optional[str]

fmgUsername: Optional[str]

hostname: Optional[str]
"""
The hostname/IP address of the FortiOS to be connected
"""

httpProxy: Optional[str]
"""
HTTP proxy address
"""

insecure: Optional[bool]

peerauth: Optional[str]
"""
Enable/disable peer authentication, can be 'enable' or 'disable'
"""

token: Optional[str]

vdom: Optional[str]

