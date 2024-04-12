



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




