# terraform-provider-p2pub

Terraform custom provider for [IIJ GIO P2 Public Resource](https://www.iij.ad.jp/biz/p2/public/). 

[日本語](README-ja.md)

## Download

Latest release: v0.4.3 (2019-02-21)

- **Linux**: [terraform-provider-p2pub-linux-amd64](https://github.com/iij/terraform-provider-p2pub/releases/download/v0.4.3/terraform-provider-p2pub-linux-amd64)
- **Windows**: [terraform-provider-p2pub-windows-amd64](https://github.com/iij/terraform-provider-p2pub/releases/download/v0.4.3/terraform-provider-p2pub-windows-amd64)
- **MacOS**: [terraform-provider-p2pub-darwin-amd64](https://github.com/iij/terraform-provider-p2pub/releases/download/v0.4.3/terraform-provider-p2pub-darwin-amd64)

## Installation

[Official Documentaion](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin).

### Linux/MacOS

1. Download this plugin and copy to appropriate location.
1. Add path to this provider in ```.terraformrc```
   ```
   $ cat ~/.terraformrc
   providers {
       p2pub = "/home/sishihara/.terraform.d/plugins/terraform-provider-p2pub"
   }
   ```

### Windows

1. Download this plugin and copy to appropriate location.
1. Add path to this provider in ```terraform.rc```
   ```
   > type C:\Users\sishihara\AppData\Roaming\terraform.rc
   providers {
      p2pub = "C:¥¥Users¥¥sishihara¥¥AppData¥¥Roaming¥¥terraform.d¥¥plugins¥¥windows_amd64¥¥terraform-provider-p2pub-windows-amd64.exe"
   }
   ```

## Usage

Before using this provider, you need to get an API key (available at [here](https://help.api.iij.jp/access_keys)).
Note that this provider requires an authority can sign up IIJ services to create P2PUB resources. Make sure your account has a role of "Service Group Administrator". check it at [IIJ Service Online](https://help.iij.ad.jp/).

### Example

In this example, the provider creates a virtual server which is accessible with ssh throught the internet:

```
provider "p2pub" {
    access_key_id = "<YOUR ACCESS KEY ID>"
    secret_access_key = "<YOUR SECRET ACCESS KEY>"
    gis_service_code = "<YOUR GIS SERVICECODE>"
}


resource "p2pub_system_storage" "storage" {
    type = "S30GB_UBUNTU16_64"
    root_ssh_key = "${file("~/.ssh/id_rsa.pub")}"
}

resource "p2pub_virtual_server" "server" {
    type = "VB0-1"
    os_type = "Linux"
    system_storage = "${p2pub_system_storage.storage.id}"
    enable_global_ip = true
}
```

After applying that, you can get the virtual server's IP address from 'terraform show' output.

### Provider configuration

This provider has 4 attributes. You can also set these attributes by using environment variables.

- ```access_key_id```: Access key id (required, ```$IIJAPI_ACCESS_KEY```)
- ```secret_access_key```: Secret access key (required, ```$IIJAPI_SECRET_KEY```)
- ```gis_service_code```: gis service code (required, ```$GISSERVICECODE```)
- ```endpoint```: P2PUB API endpoint. currently in depvelopers use only

### Resource list

#### ```p2pub_virtual_server```: [Virtual Server](http://manual.iij.jp/p2/pub/b-1.html)

| key | value | required |
|-|-|-|
|```type```| [Server type](http://manual.iij.jp/p2/pubapi/59949011.html) | o |
|```os_type```| OS type use in the virtual server. ```Linux``` or ```Windows``` | o |
|```server_group```| [Server group](http://manual.iij.jp/p2/pub/b-1-5.html). ```A``` or ```B``` | |
|```label```| | |
|```system_storage```| System Storage service code attached the virtual server | |
|```data_storage```| List of Additional Storage service code attached the virtual server | |
|```private_network```| List of Private Netowrk/V service code connected the virtual server | |
|```enable_global_ip```| true if you attach a global IP address to the server. default is false | |

**Example**

```
resource "p2pub_virtual_server" "server" {
    type = "VB0-1"
    os_type = "Linux"
    server_group = "A"
    label = "my server"
    system_storage = "iba01234567"
    data_storage = ["ibb01234567", "ibg01234567"]
    private_network = ["ivl01234567"]
    enable_global_ip = true
}
```

#### ```p2pub_system_storage```: [System Storage](http://manual.iij.jp/p2/pub/b-3-1.html)

| key | value | required |
|-|-|-|
|```type```| [Storage type (system storage)](http://manual.iij.jp/p2/pubapi/59949023.html) | o |
|```label```| | |
|```root_password```| root password in plain text | |
|```root_ssh_key```| ssh public key for root user | |
|```userdata```| Base64-encoded UserData string | |
|```encryption```| enable encryption. ```Yes``` / ```No``` | required in use of type-X storage. [detail](http://manual.iij.jp/p2/pubapi/59939812.html) |
|```source_image```| set this when you create the storage by restoring from Storage Archive | |
|```source_image.gis_service_code```| P2 service code source image is located in | |
|```source_image.iar_service_code```| Storage Archive service code source image is located in | |
|```source_image.image_id```| source image's id | |

**Example**

```
resource "p2pub_system_storage" "system_storage" {
    type = "S30GB_UBUNTU16_64"
    label = "my system storage"
    root_ssh_key = "${file("~/.ssh/id_rsa.pub")}"
    userdata = "${base64encode(userdata)}"
    source_image {
        gis_service_code = "gis99999999"
        iar_service_code = "iar99999999"
        image_id = "999999"
    }
}
```

#### ```p2pub_additional_storage```: [Additional Storage](http://manual.iij.jp/p2/pub/b-3-2.html)

| key | value | required |
|-|-|-|
|```type```| [Storage type (additional storage)](http://manual.iij.jp/p2/pubapi/59949023.html) | o |
|```encryption```| enable encryption. ```Yes``` / ```No``` | required in use of type-X storage. [detail](http://manual.iij.jp/p2/pubapi/59940088.html) |
|```label```| | |

**Example**

```
resource "p2pub_additional_storage" "additional_storage" {
    type = "B1000GB"
    label = "my additional storage"
}
```

#### ```p2pub_private_network```: [Private Network/V](http://manual.iij.jp/p2/pub/b-5.html)

| key | value | required |
|-|-|-|
|```label```| | |

#### ```p2pub_storage_archive```: [Storage Archive](http://manual.iij.jp/p2/pub/b-4.html)

| key | value | required |
|-|-|-|
|```archive_size```| capacity for archived images in GB. Need to set multiple of 10. | o |

#### ```p2pub_global_ip_address```: [Global IP Address/V](http://manual.iij.jp/p2/pub/b-5.html)

| key | value | required |
|-|-|-|
|```address_num```| amount of global ip addresses additionally allocate the contract (0~15) | |

#### ```p2pub_load_balancer```: [FW+LB dedicated type](http://manual.iij.jp/p2/pub/b-6-7.html)

|key|value||required|
|-|-|-|-|
|```type```|type|D10M, D100M, D150M, D1000M|o|
|```redundant```|redundancy|"Yes" "No"|o|
|```external_type```|network type|"Global", "PrivateStandard", "Private"|o|
|```external_servicecode```|network servicecode|ivlServicecode||
|```external_masterhost_address```|address of master host|ipaddr||
|```external_slavehost_address```|address of slave host|ipaddr||
|```external_netmask```|netmask of host|mask||
|```internal_type```|network type|"PrivateStandard", "Private"|o|
|```internal_servicecode```|network servicecode|ivlServicecode||
|```internal_masterhost_address```|address of master host|ipaddr||
|```internal_slavehost_address```|address of slave host|ipaddr||
|```internal_netmask```|netmask of host|mask||
|```trafficip_list```|list of trafficips|array|o|
|```trafficip_list.ipv4_name```|name of trafficip|string|o|
|```filter_in_list```|rules of firewall (in)|array||
|```filter_in_list.source_network```|source network|ipaddr/mask, ANY||
|```filter_in_list.destination_network```|destination network|ipaddr/mask, ANY||
|```filter_in_list.destination_port```|destination port|number, ANY||
|```filter_in_list.protocol```|protocol|TCP, UDP||
|```filter_in_list.action```|action|ACCEPT, DROP, REJECT||
|```filter_in_list.label```|label|string||
|```filter_out_list```|rules of firewall (out)|array||
|```filter_out_list.source_network```|source network|ipaddr/mask, ANY||
|```filter_out_list.destination_network```|destination network|ipaddr/mask, ANY||
|```filter_out_list.destination_port```|destination port|number, ANY||
|```filter_out_list.protocol```|protocol|TCP, UDP||
|```filter_out_list.action```|action|ACCEPT, DROP, REJECT||
|```filter_out_list.label```|label|string||
|```administration_server_allow_network_list```|acl for control panel of load balancer|array of ip addresses||


**Example**
```
resource "p2pub_load_balancer" "vtm1" {
    type = "D10M"
    redundant = "No"

    external_type = "Global"
    internal_type = "PrivateStandard"

    trafficip_list = [
        { ipv4_name = "TRAFFICIP1" }
    ]

    filter_in_list = [
        {
            source_network = "ANY"
            destination_network = "ANY"
            destination_port = "80"
            protocol = "TCP"
            action = "ACCEPT"
            label = "ALLOW HTTP"
        },
        {
            source_network = "ANY"
            destination_network = "ANY"
            destination_port = "443"
            protocol = "TCP"
            action = "ACCEPT"
            label = "ALLOW HTTPS"
        }
    ]

    filter_out_list = [
        {
            source_network = "ANY"
            destination_network = "ANY"
            destination_port = "ANY"
            protocol = "TCP"
            action = "ACCEPT"
            label = "ALLOW ALL TCP"
        }
    ]
    administration_server_allow_network_list = [
        "192.0.2.0/24",
        "198.51.100.0/24",
        "203.0.113.0/24"
    ]
}
```

## Developing this provider

### Build from source

[dep](https://github.com/golang/dep) is required.

```
$ make build
```


## References

- IIJ GIO P2 Public Resource API Reference : http://manual.iij.jp/p2/pubapi/index.html




Author: Shoma Ishihara (sishihara@iij.ad.jp)
