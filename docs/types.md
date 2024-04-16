



## InstanceArray
InstanceArray object describes a collection of identical instances






<hr />

<div class="dd">

<code>instanceID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the object.

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

description: The label of the object. Must be unique per infrastructure. Must follow DNS naming rules: Pattern: ^[a-zA-Z]{1,1}[a-zA-Z0-9-]{0,61}[a-zA-Z0-9]{1,1}|[a-zA-Z]{1,1}$

</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

description: |
		User editable DNS record that gets created for this instance array in the built-in DNS
		service and associated with all the primary IP address on the WAN network. Must adhere to DNS naming rules such
     as:  only "-", lowercase alphanumeric characters and not start with a number.
     Pattern:


</div>

<hr />

<div class="dd">

<code>bootMethod</code>  <i>string</i>

</div>
<div class="dt">

description: The booth method to use. Note that the pxe_iscsi booth method is deprecated.
values:
	- local_disks
 - pxe_iscsi


</div>

<hr />

<div class="dd">

<code>instanceCount</code>  <i>int</i>

</div>
<div class="dt">

The number of instances in this array.

</div>

<hr />

<div class="dd">

<code>ramGBytes</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum amount of RAM expressed in GB.

</div>

<hr />

<div class="dd">

<code>processorCount</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum amount of CPU sockets.

</div>

<hr />

<div class="dd">

<code>processorCoreMhz</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum GPU frequency.

</div>

<hr />

<div class="dd">

<code>processorCoreCount</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum core count (hyperthreaded).

</div>

<hr />

<div class="dd">

<code>diskCount</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum disk count.

</div>

<hr />

<div class="dd">

<code>diskSizeMBytes</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum disk size.

</div>

<hr />

<div class="dd">

<code>diskTypes</code>  <i>[]string</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this disk type. This assumes all disks are identical.


Valid values:


  - <code>ssh</code>

  - <code>hdd</code>
</div>

<hr />

<div class="dd">

<code>infrastructureID</code>  <i>int</i>

</div>
<div class="dt">

The id of the infrastructure on which this infrastructure is created on.

</div>

<hr />

<div class="dd">

<code>serviceStatus</code>  <i>string</i>

</div>
<div class="dt">

The service status.  Read only.


Valid values:


  - <code>ordered</code>

  - <code>active</code>

  - <code>suspended</code>

  - <code>stopped</code>

  - <code>deleted</code>
</div>

<hr />

<div class="dd">

<code>interfaces</code>  <i>[]<a href="#instancearrayinterface">InstanceArrayInterface</a></i>

</div>
<div class="dt">

The instance array interfaces configuration

</div>

<hr />

<div class="dd">

<code>clusterID</code>  <i>int</i>

</div>
<div class="dt">

The cluster (such as Kubernetes, VMWare vSphere etc) of which this instance array is part of. A vanilla cluster is created for all instance arrays not added to any application cluster.

</div>

<hr />

<div class="dd">

<code>clusterRoleGroup</code>  <i>string</i>

</div>
<div class="dt">

If part of an app cluster this field will receive the role that this instance array has such as `master` or `worker` which is application specific.

</div>

<hr />

<div class="dd">

<code>firewallManaged</code>  <i>bool</i>

</div>
<div class="dt">

description: If set to true, the firewall will be configured based on rules provided in the InstanceArrayFirewallRules field. Note that for this to work the following conditions must be fufilled:
a. The OS template should have the template set capability (for the first configuration of the firewall, at install time)
b. The in-band site controller agent should be enabled with in-band access to the operating system over SSH.


</div>

<hr />

<div class="dd">

<code>firewallRules</code>  <i>[]FirewallRule</i>

</div>
<div class="dt">

The list of firewall rules to configure

</div>

<hr />

<div class="dd">

<code>volumeTemplateID</code>  <i>int</i>

</div>
<div class="dt">

The operating system template to use.

</div>

<hr />

<div class="dd">

<code>operation</code>  <i><a href="#instancearrayoperation">InstanceArrayOperation</a></i>

</div>
<div class="dt">

Used when changing an instance array. It captures the operation that needs to happen on the instance array.

</div>

<hr />

<div class="dd">

<code>additionalWanIPv4</code>  <i>string</i>

</div>
<div class="dt">

Information about additional ips to be assigned to the WAN interfaces. Used internally.

</div>

<hr />

<div class="dd">

<code>customVariables</code>  <i>interface{}</i>

</div>
<div class="dt">

Custom variables and variable overrides to be pushed to the operating system deployment process.

</div>

<hr />

<div class="dd">

<code>firmwarePolicies</code>  <i>[]int</i>

</div>
<div class="dt">

Firmware policies to apply. Deprecated. Use baselines.

</div>

<hr />

<div class="dd">

<code>drive_array_id_boot</code>  <i>int</i>

</div>
<div class="dt">

When iSCSI boot is used this is the id of the drive array that will be the boot device.

</div>

<hr />





## InstanceArrayOperation
InstanceArrayOperation object describes the changes that will be applied to an instance array

Appears in:


- <code><a href="#instancearray">InstanceArray</a>.operation</code>





<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The ID of the object.

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

description: The label of the object. Must be unique per infrastructure. Must follow DNS naming rules: Pattern: ^[a-zA-Z]{1,1}[a-zA-Z0-9-]{0,61}[a-zA-Z0-9]{1,1}|[a-zA-Z]{1,1}$

</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

description: |
		User editable DNS record that gets created for this instance array in the built-in DNS
		service and associated with all the primary IP address on the WAN network. Must adhere to DNS naming rules such
     as:  only "-", lowercase alphanumeric characters and not start with a number.
     Pattern:


</div>

<hr />

<div class="dd">

<code>bootMethod</code>  <i>string</i>

</div>
<div class="dt">

description: The booth method to use. Note that the pxe_iscsi booth method is deprecated.
values:
	- local_disks
 - pxe_iscsi


</div>

<hr />

<div class="dd">

<code>instanceCount</code>  <i>int</i>

</div>
<div class="dt">

The number of instances in this array.

</div>

<hr />

<div class="dd">

<code>ramGBytes</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum amount of RAM expressed in GB.

</div>

<hr />

<div class="dd">

<code>processorCount</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum amount of CPU sockets.

</div>

<hr />

<div class="dd">

<code>processorCoreMhz</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum GPU frequency.

</div>

<hr />

<div class="dd">

<code>processorCoreCount</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum core count (hyperthreaded).

</div>

<hr />

<div class="dd">

<code>diskCount</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum disk count.

</div>

<hr />

<div class="dd">

<code>diskSizeMBytes</code>  <i>int</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this minimum disk size.

</div>

<hr />

<div class="dd">

<code>diskTypes</code>  <i>[]string</i>

</div>
<div class="dt">

When the ServerTypeID is not set on the Instance object this will restrict the search for a matching server to just those with this disk type. This assumes all disks are identical.


Valid values:


  - <code>ssh</code>

  - <code>hdd</code>
</div>

<hr />

<div class="dd">

<code>serviceStatus</code>  <i>string</i>

</div>
<div class="dt">

The service status.  Read only.


Valid values:


  - <code>ordered</code>

  - <code>active</code>

  - <code>suspended</code>

  - <code>stopped</code>

  - <code>deleted</code>
</div>

<hr />

<div class="dd">

<code>interfaces</code>  <i>[]<a href="#instancearrayinterfaceoperation">InstanceArrayInterfaceOperation</a></i>

</div>
<div class="dt">

The instance array interfaces configuration

</div>

<hr />

<div class="dd">

<code>clusterID</code>  <i>int</i>

</div>
<div class="dt">

The cluster (such as Kubernetes, VMWare vSphere etc) of which this instance array is part of. A vanilla cluster is created for all instance arrays not added to any application cluster.

</div>

<hr />

<div class="dd">

<code>clusterRoleGroup</code>  <i>string</i>

</div>
<div class="dt">

If part of an app cluster this field will receive the role that this instance array has such as master or worker which is application specific.

</div>

<hr />

<div class="dd">

<code>firewallManaged</code>  <i>bool</i>

</div>
<div class="dt">

description: If set to true, the firewall will be configured based on rules provided in the InstanceArrayFirewallRules field. Note that for this to work the following conditions must be fufilled:
a. The OS template should have the template set capability (for the first configuration of the firewall, at install time)
b. The in-band site controller agent should be enabled with in-band access to the operating system over SSH.


</div>

<hr />

<div class="dd">

<code>firewallRules</code>  <i>[]FirewallRule</i>

</div>
<div class="dt">

The list of firewall rules to configure

</div>

<hr />

<div class="dd">

<code>volumeTemplateID</code>  <i>int</i>

</div>
<div class="dt">

The operating system template to use.

</div>

<hr />

<div class="dd">

<code>deployType</code>  <i>string</i>

</div>
<div class="dt">

description: The deploy type, one of:
values:
    - create
	   - delete
    - edit
	   - start
    - stop
	   - suspend


</div>

<hr />

<div class="dd">

<code>deployStatus</code>  <i>string</i>

</div>
<div class="dt">

The status of the deployment


Valid values:


  - <code>not_started</code>

  - <code>ongoing</code>

  - <code>finished</code>
</div>

<hr />

<div class="dd">

<code>changeID</code>  <i>int</i>

</div>
<div class="dt">

The id of the change operation. Readonly.

</div>

<hr />

<div class="dd">

<code>additionalWanIPv4</code>  <i>string</i>

</div>
<div class="dt">

Information about additional ips to be assigned to the WAN interfaces. Used internally.

</div>

<hr />

<div class="dd">

<code>customVariables</code>  <i>interface{}</i>

</div>
<div class="dt">

Custom variables and variable overrides to be pushed to the operating system deployment process.

</div>

<hr />

<div class="dd">

<code>firmwarePolicies</code>  <i>[]int</i>

</div>
<div class="dt">

Firmware policies to apply. Deprecated. Use baselines instead.

</div>

<hr />

<div class="dd">

<code>drive_array_id_boot</code>  <i>int</i>

</div>
<div class="dt">

When iSCSI boot is used this is the id of the drive array that will be the boot device.

</div>

<hr />





## InstanceArrayInterface
InstanceArrayInterface describes a network interface of the array.
It's properties will be applied to all InstanceInterfaces of the array's instances.


Appears in:


- <code><a href="#instancearray">InstanceArray</a>.interfaces</code>





<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

Label of the interface

</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

An unique string describing the interface.

</div>

<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

Interface ID. A unique id of the interface.

</div>

<hr />

<div class="dd">

<code>instanceArrayID</code>  <i>int</i>

</div>
<div class="dt">

The instance array to which this interface is associated to

</div>

<hr />

<div class="dd">

<code>networkID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the Network to which this interface is connected to. Can be 0 if the interface is not connected.

</div>

<hr />

<div class="dd">

<code>LAGGIndexes</code>  <i>[]interface{}</i>

</div>
<div class="dt">

Used internally. Readonly.

</div>

<hr />

<div class="dd">

<code>index</code>  <i>int</i>

</div>
<div class="dt">

description: |
	The index of the interface in the server. This is 0-based and configured based on the lexicographic sorting of the switch and switch ports
	thus NOT based on the PCI slots or anything like that. This ensures consistent ordering regardless of cabling and/or NIC seating.


</div>

<hr />

<div class="dd">

<code>serviceStatus</code>  <i>string</i>

</div>
<div class="dt">

The service status.  Read only.


Valid values:


  - <code>ordered</code>

  - <code>active</code>

  - <code>suspended</code>

  - <code>stopped</code>

  - <code>deleted</code>
</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

The creation date and time in ISO 8601 format.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

The last update date and time in ISO 8601 format.

</div>

<hr />

<div class="dd">

<code>operation</code>  <i><a href="#instancearrayinterfaceoperation">InstanceArrayInterfaceOperation</a></i>

</div>
<div class="dt">

The operation object. Must be set to alter the configuration of an interface.

</div>

<hr />

<div class="dd">

<code>instance_array_interface_change_id</code>  <i>int</i>

</div>
<div class="dt">

Ongoing change ID. Readonly.

</div>

<hr />





## InstanceArrayInterfaceOperation
InstanceArrayInterfaceOperation describes changes to a network array interface

Appears in:


- <code><a href="#instancearrayoperation">InstanceArrayOperation</a>.interfaces</code>

- <code><a href="#instancearrayinterface">InstanceArrayInterface</a>.operation</code>





<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

Label of the interface

</div>

<hr />

<div class="dd">

<code>subdomain</code>  <i>string</i>

</div>
<div class="dt">

An unique string describing the interface.

</div>

<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

Interface ID. A unique id of the interface.

</div>

<hr />

<div class="dd">

<code>instanceArrayID</code>  <i>int</i>

</div>
<div class="dt">

The instance array to which this interface is associated to

</div>

<hr />

<div class="dd">

<code>networkID</code>  <i>int</i>

</div>
<div class="dt">

The ID of the Network to which this interface is connected to. Can be 0 if the interface is not connected.

</div>

<hr />

<div class="dd">

<code>LAGGIndexes</code>  <i>[]interface{}</i>

</div>
<div class="dt">

Used internally. Readonly.

</div>

<hr />

<div class="dd">

<code>index</code>  <i>int</i>

</div>
<div class="dt">

description: |
	The index of the interface in the server. This is 0-based and configured based on the lexicographic sorting of the switch and switch ports
	thus NOT based on the PCI slots or anything like that. This ensures consistent ordering regardless of cabling and/or NIC seating.


</div>

<hr />

<div class="dd">

<code>serviceStatus</code>  <i>string</i>

</div>
<div class="dt">

The service status.  Read only.


Valid values:


  - <code>ordered</code>

  - <code>active</code>

  - <code>suspended</code>

  - <code>stopped</code>

  - <code>deleted</code>
</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

The creation date and time in ISO 8601 format.

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

The last update date and time in ISO 8601 format.

</div>

<hr />

<div class="dd">

<code>changeID</code>  <i>int</i>

</div>
<div class="dt">

Ongoing change ID. Readonly.

</div>

<hr />








## SubnetOOB
Subnet represents a subnet for OOB operations



```yaml
id: 10 # The id of the object
label: mysubnet # The label of the object

# # The Netmask to use.
# netmask: 255.255.255.192

# # The Prefix size in CIDR format. Must match the netmask
# size: 26

# # The start of the range
# rangeStart: 192.168.0.10

# # The end of the range.
# rangeEnd: 192.168.0.100
```



<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The id of the object

</div>

<hr />

<div class="dd">

<code>label</code>  <i>string</i>

</div>
<div class="dt">

The label of the object

</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

The type of the object


Valid values:


  - <code>ipv4</code>

  - <code>ipv6</code>
</div>

<hr />

<div class="dd">

<code>useForAutoAllocation</code>  <i>bool</i>

</div>
<div class="dt">

If set to `true` this subnet will be used for auto-allocation of IPs

</div>

<hr />

<div class="dd">

<code>forResourceType</code>  <i>string</i>

</div>
<div class="dt">

What type of resource to allocate this object for


Valid values:


  - <code>server</code>

  - <code>network_equipment</code>

  - <code>any</code>
</div>

<hr />

<div class="dd">

<code>blacklist</code>  <i>[]string</i>

</div>
<div class="dt">

description: Array of IPs that are to be skipped from the interval
examples:
  - value: ['192.168.0.10','192.168.0.22']


</div>

<hr />

<div class="dd">

<code>gatewayHex</code>  <i>string</i>

</div>
<div class="dt">

The Gateway in hexadecimal format. Readonly.

</div>

<hr />

<div class="dd">

<code>gateway</code>  <i>string</i>

</div>
<div class="dt">

description: The Gateway to use when allocating IPs from this subnet.
examples:
  -values:  '"192.168.0.1"'


</div>

<hr />

<div class="dd">

<code>netmaskHex</code>  <i>string</i>

</div>
<div class="dt">

The Netmask in hexadecimal format. Readonly.

</div>

<hr />

<div class="dd">

<code>netmask</code>  <i>string</i>

</div>
<div class="dt">

The Netmask to use.



Examples:


```yaml
netmask: 255.255.255.192
```


</div>

<hr />

<div class="dd">

<code>size</code>  <i>int</i>

</div>
<div class="dt">

The Prefix size in CIDR format. Must match the netmask



Examples:


```yaml
size: 26
```


</div>

<hr />

<div class="dd">

<code>rangeStartHex</code>  <i>string</i>

</div>
<div class="dt">

The start of the range in hexadecimal. Readonly.

</div>

<hr />

<div class="dd">

<code>rangeStart</code>  <i>string</i>

</div>
<div class="dt">

The start of the range



Examples:


```yaml
rangeStart: 192.168.0.10
```


</div>

<hr />

<div class="dd">

<code>rangeEndHex</code>  <i>string</i>

</div>
<div class="dt">

The end of the range in hexadecimal. Readonly

</div>

<hr />

<div class="dd">

<code>rangeEnd</code>  <i>string</i>

</div>
<div class="dt">

The end of the range.



Examples:


```yaml
rangeEnd: 192.168.0.100
```


</div>

<hr />

<div class="dd">

<code>datacenter</code>  <i>string</i>

</div>
<div class="dt">

The data center in which this subnet is valid

</div>

<hr />








## DatacenterWithConfig
A data center object that contains both metadata and configuration








## Datacenter
Datacenter metadata






<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The ID of this datacenter.

</div>

<hr />

<div class="dd">

<code>name</code>  <i>string</i>

</div>
<div class="dt">

The name (label) of this datacenter. Once set it cannot be changed.

</div>

<hr />

<div class="dd">

<code>parentName</code>  <i>string</i>

</div>
<div class="dt">

The name (label) of the parent datacenter. This is useful in hierarchical setups where one datacenter needs to access it's parent's resources.

</div>

<hr />

<div class="dd">

<code>userid</code>  <i>int</i>

</div>
<div class="dt">

The owner of a datacenter.

</div>

<hr />

<div class="dd">

<code>displayname</code>  <i>string</i>

</div>
<div class="dt">

The display name of a data center. Can be changed.

</div>

<hr />

<div class="dd">

<code>ismaster</code>  <i>bool</i>

</div>
<div class="dt">

Deprecated.

</div>

<hr />

<div class="dd">

<code>ismaintenance</code>  <i>bool</i>

</div>
<div class="dt">

If set to true no new operations can happen on this datacenter.

</div>

<hr />

<div class="dd">

<code>type</code>  <i>string</i>

</div>
<div class="dt">

The datacenter type. Deprecated. Currently the only supported value is metal_cloud.


Valid values:


  - <code>metal_cloud</code>
</div>

<hr />

<div class="dd">

<code>createdtimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the datacenter was created.

</div>

<hr />

<div class="dd">

<code>updatedtimestamp</code>  <i>string</i>

</div>
<div class="dt">

ISO 8601 timestamp which holds the date and time when the datacenter was updated.

</div>

<hr />

<div class="dd">

<code>ishidden</code>  <i>bool</i>

</div>
<div class="dt">

If set the datacenter will not be visible in the UI

</div>

<hr />

<div class="dd">

<code>tags</code>  <i>[]string</i>

</div>
<div class="dt">

An array of tags (strings)

</div>

<hr />





## DatacenterConfig
DatacenterConfig - datacenter configuration






<hr />

<div class="dd">

<code>BSIMachinesSubnetIPv4CIDR</code>  <i>string</i>

</div>
<div class="dt">

The ip address of the Global Controller. Deprecated.

</div>

<hr />

<div class="dd">

<code>BSIVRRPListenIPv4</code>  <i>string</i>

</div>
<div class="dt">

The ip address on which all datacenter agents listen for connections. Deprecated.

</div>

<hr />

<div class="dd">

<code>BSIMachineListenIPv4List</code>  <i>[]string</i>

</div>
<div class="dt">

Site Controller's secondary ip addresses. Deprecated.

</div>

<hr />

<div class="dd">

<code>BSIExternallyVisibleIPv4</code>  <i>string</i>

</div>
<div class="dt">

The agent's IP that is visible from the controller. Deprecated.

</div>

<hr />

<div class="dd">

<code>repoURLRoot</code>  <i>string</i>

</div>
<div class="dt">

The repository to use

</div>

<hr />

<div class="dd">

<code>repoURLRootQuarantineNetwork</code>  <i>string</i>

</div>
<div class="dt">

The repository to use during legacy (PXE) provisioning process. Same as repoURLRoot, with an IP address for the hostname, required in networks where DNS is not available.

</div>

<hr />

<div class="dd">

<code>SANRoutedSubnet</code>  <i>string</i>

</div>
<div class="dt">

The SAN subnet in CIDR format.

</div>

<hr />

<div class="dd">

<code>NTPServers</code>  <i>[]string</i>

</div>
<div class="dt">

IP addresses of NTP servers.

</div>

<hr />

<div class="dd">

<code>DNSServers</code>  <i>[]string</i>

</div>
<div class="dt">

IP addresses of DNS servers to be used in the DHCP response.

</div>

<hr />

<div class="dd">

<code>KMS</code>  <i>string</i>

</div>
<div class="dt">

Host (IP:port) of the Windows machine hosting the Key Management Service. Set to empty string to disable.

</div>

<hr />

<div class="dd">

<code>TFTPServerWANVRRPListenIPv4</code>  <i>string</i>

</div>
<div class="dt">

The IP of the Site Controller TFTP service used during the legacy (PXE) deployment process.

</div>

<hr />

<div class="dd">

<code>dataLakeEnabled</code>  <i>bool</i>

</div>
<div class="dt">

If set to true, the datalake service is enabled in this environment. Deprecated


Valid values:


  - <code>true</code>

  - <code>false</code>
</div>

<hr />

<div class="dd">

<code>monitoringGraphitePlainTextSocketHost</code>  <i>string</i>

</div>
<div class="dt">

Graphite host (IPv4:port) for the plain text protocol socket. Set to empty string to disable. Deprecated

</div>

<hr />

<div class="dd">

<code>monitoringGraphiteRenderURLHost</code>  <i>string</i>

</div>
<div class="dt">

Graphite host (IPv4:port) for the HTTP Render URL API. Set to empty string to disable. Deprecated

</div>

<hr />

<div class="dd">

<code>latitude</code>  <i>float64</i>

</div>
<div class="dt">

The Datacenter's latitude. Use negative numbers for the south hemisphere

</div>

<hr />

<div class="dd">

<code>longitude</code>  <i>float64</i>

</div>
<div class="dt">

description: The data center's longitude: Use negative numbers for areas west of Greenwich (UK)

</div>

<hr />

<div class="dd">

<code>address</code>  <i>string</i>

</div>
<div class="dt">

The data center's address

</div>

<hr />

<div class="dd">

<code>serverRegisterUsingGeneratedIPMICredentialsEnabled</code>  <i>bool</i>

</div>
<div class="dt">

If set to true the system will configure a randomly generated username and password on the server's BMC(ILO/IDRAC etc.)


Valid values:


  - <code>true</code>

  - <code>false</code>
</div>

<hr />

<div class="dd">

<code>serverRegisterUsingProvidedIPMICredentialsEnabled</code>  <i>bool</i>

</div>
<div class="dt">

If set to true the system will ask for credentials during server registration.


Valid values:


  - <code>true</code>

  - <code>false</code>
</div>

<hr />

<div class="dd">

<code>switchProvisioner</code>  <i>map[string]interface{}</i>

</div>
<div class="dt">

The provisioner (fabric) to use when provisioning the network on switch devices


Valid values:


  - <code>VLAN</code>

  - <code>EVPNVXLANL2</code>

  - <code>VPLS</code>

  - <code>LAN</code>

  - <code>SDN</code>
</div>

<hr />

<div class="dd">

<code>enableTenantAccessToIPMI</code>  <i>bool</i>

</div>
<div class="dt">

If set to true the tenants will receive credentials for accessing the server's BMC with a special user.

</div>

<hr />

<div class="dd">

<code>allowVLANOverrides</code>  <i>bool</i>

</div>
<div class="dt">

description: Allows the end-user to force a VLAN ID (or EPG in CISCO ACI environments). This enables the user to connect to pre-existing VLANs in the established infrastructure. WARNING: This enables a tenant to access unauthorized VLANs.

</div>

<hr />

<div class="dd">

<code>allowNetworkProfiles</code>  <i>bool</i>

</div>
<div class="dt">

Allows the usage of network profiles for customizing InstanceArray network connections.

</div>

<hr />

<div class="dd">

<code>enableServerRegistrationStartedByInBandDHCP</code>  <i>bool</i>

</div>
<div class="dt">

If set enables in-band triggered registration via the legacy (PXE) mechanism.

</div>

<hr />

<div class="dd">

<code>extraInternalIPsPerSubnet</code>  <i>int</i>

</div>
<div class="dt">

Extra ips to reserve on each subnet for WAN networks. Certain fabrics (such as VRRP-based L3 SVIs need more than one IP to be allocated on each subnet). This option will force the system to reserve this number of IPs from each subnet.

</div>

<hr />

<div class="dd">

<code>extraInternalIPsPerSANSubnet</code>  <i>int</i>

</div>
<div class="dt">

Extra ips to reserve on each subnet for SAN networks. Certain fabrics (such as VRRP-based L3 SVIs need more than one IP to be allocated on each subnet). This option will force the system to reserve this number of IPs from each subnet.

</div>

<hr />

<div class="dd">

<code>serverRAIDConfigurationEnabled</code>  <i>bool</i>

</div>
<div class="dt">

If enabled RAID configurations are set on servers

</div>

<hr />

<div class="dd">

<code>webProxy</code>  <i><a href="#webproxy">WebProxy</a></i>

</div>
<div class="dt">

If configured the proxy will be used by all operations.

</div>

<hr />

<div class="dd">

<code>isKubernetesDeployment</code>  <i>bool</i>

</div>
<div class="dt">

Deprecated.

</div>

<hr />

<div class="dd">

<code>allowInstanceArrayFirmwarePolicies</code>  <i>bool</i>

</div>
<div class="dt">

If set it allows  the use of firmware policies. Note that for baselines to function this needs to be enabled.

</div>

<hr />

<div class="dd">

<code>provisionUsingTheQuarantineNetwork</code>  <i>bool</i>

</div>
<div class="dt">

If set to true, during the legacy registration process (PXE) the system will configure special provisioning VLAN on server ports prior to performing the deployment

</div>

<hr />

<div class="dd">

<code>enableDHCPRelaySecurityForQuarantineNetwork</code>  <i>bool</i>

</div>
<div class="dt">

If set to true, during the legacy registration process (PXE) the system will enforce DHCP option 82 security.

</div>

<hr />

<div class="dd">

<code>enableDHCPRelaySecurityForClientNetworks</code>  <i>bool</i>

</div>
<div class="dt">

If set to true, the DHCP server will ignore requests that do not respect DHCP option 82 for regular networks.

</div>

<hr />

<div class="dd">

<code>enableDHCPBMCMACAddressWhitelist</code>  <i>bool</i>

</div>
<div class="dt">

If enabled, the DHCPBMCMACAddressWhitelist will be used to whitelist certain MAC addresses in order to ensure that only certain servers get registered during the ZTP process.

</div>

<hr />

<div class="dd">

<code>dhcpBMCMACAddressWhitelist</code>  <i>[]string</i>

</div>
<div class="dt">

The mac addresses of the servers that are to be allowed to be registered via ZTP. This is useful during initial testing.

</div>

<hr />

<div class="dd">

<code>defaultServerCleanupPolicyID</code>  <i>int</i>

</div>
<div class="dt">

If set the server cleanup policy will be the policy with the specified id instead of the default one (which is 0)

</div>

<hr />

<div class="dd">

<code>defaultWANNetworkProfileID</code>  <i>int</i>

</div>
<div class="dt">

If set, this will be the default network profile instead of no network profile.

</div>

<hr />

<div class="dd">

<code>defaultDeploymentMechanism</code>  <i>string</i>

</div>
<div class="dt">

Deployment mechanism used in case a server supports both Virtual Media and legacy (PXE).


Valid values:


  - <code>virtual_media</code>

  - <code>pxe</code>
</div>

<hr />

<div class="dd">

<code>defaultCleanupAndRegistrationMechanism</code>  <i>string</i>

</div>
<div class="dt">

The cleanup and register mechanism used in case a server supports both BMC-only and BDK mechanisms. Defaults to BMC.


Valid values:


  - <code>bmc</code>

  - <code>bdk</code>
</div>

<hr />

<div class="dd">

<code>NFSServer</code>  <i>string</i>

</div>
<div class="dt">

The NFS server to use for server OS deployment (the IP of the site controller as seen from the server's BMC). Should be an IP to avoid DNS resolutions.

</div>

<hr />

<div class="dd">

<code>Option82ToIPMapping</code>  <i>Option82ToIPMapping</i>

</div>
<div class="dt">

Can be used to set a mapping between Option82 and IPs that the DHCP server allocates to servers during registration.

</div>

<hr />





## WebProxy
Defines web proxy configuration

Appears in:


- <code><a href="#datacenterconfig">DatacenterConfig</a>.webProxy</code>





<hr />

<div class="dd">

<code>ip</code>  <i>string</i>

</div>
<div class="dt">

Ip fo the web proxy

</div>

<hr />

<div class="dd">

<code>port</code>  <i>int</i>

</div>
<div class="dt">

Port fo the web proxy

</div>

<hr />

<div class="dd">

<code>username</code>  <i>string</i>

</div>
<div class="dt">

Username of the web proxy

</div>

<hr />

<div class="dd">

<code>password</code>  <i>string</i>

</div>
<div class="dt">

Password to use for the web proxy

</div>

<hr />








## Variable
Variable struct defines a Variable type






<hr />

<div class="dd">

<code>id</code>  <i>int</i>

</div>
<div class="dt">

The id of the object

</div>

<hr />

<div class="dd">

<code>ownerID</code>  <i>int</i>

</div>
<div class="dt">

The id of the owner (user object) of the object

</div>

<hr />

<div class="dd">

<code>userIDAuthenticated</code>  <i>int</i>

</div>
<div class="dt">

The id of the user that is currently manipulating the object. Readonly.

</div>

<hr />

<div class="dd">

<code>name</code>  <i>string</i>

</div>
<div class="dt">

The name of the variable

</div>

<hr />

<div class="dd">

<code>usage</code>  <i>string</i>

</div>
<div class="dt">

The usage of a variable

</div>

<hr />

<div class="dd">

<code>json</code>  <i>string</i>

</div>
<div class="dt">

The content of the variable in json encoded format

</div>

<hr />

<div class="dd">

<code>createdTimestamp</code>  <i>string</i>

</div>
<div class="dt">

Timestamp of the variable creation date. Readonly

</div>

<hr />

<div class="dd">

<code>updatedTimestamp</code>  <i>string</i>

</div>
<div class="dt">

Timestamp of the variable last update. Readonly

</div>

<hr />




