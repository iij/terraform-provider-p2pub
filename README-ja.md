# terraform-provider-p2pub

Terraform Custom Provider for P2 PUB

# 使い方

## セットアップ

### ビルドとインストール

make してバイナリを作成した後、.terraformrc にパスを追加します

```
$ make build
$ cp terraform-provider-p2pub /path/to/install/dir
```

- ~/.terraformrc

```
providers {
    p2pub = "/path/to/install/dir/terraform-provider-p2pub"
}
```

### プロバイダーの設定
P2 PUBをTerraformで操作するための運用管理担当者アカウント（マスターID）を用意します。

[IIJサービスオンライン](https://help.iij.ad.jp/) にて、
サービス契約ID(gisサービスコード)に対して、下記両方の役割として登録された担当者を追加します。
- 「サービスグループの運用管理担当者」
- 「サービスの運用管理担当者」

担当者の追加方法はIIJサービスオンラインのマニュアルを参照してください。

[IIJ：ご利用にあたって](https://help.iij.ad.jp/admin/guidance/termsofuse/index.cfm)

次に、マスターIDに紐づくAPIキーを発行します。APIキーの発行方法は下記マニュアルを参照してください。

[AccessKey](http://manual.iij.jp/p2/pubapi/59950199.html)

最後に、tf ファイルに API キーと gis サービスコードを記述します。Terraform で管理されるリソース(仮想サーバー等)はここで指定した gis サービスコード配下の契約になります。

```
provider "p2pub" {
    access_key_id = "<ACCESSKEY>"
    secret_access_key = "<SECRETKEY>"
    gis_service_code = "<SERVICECODE>"
}
```

設定値は環境変数から与えることもできます。

|項目|環境変数名|
|-|-|
|```access_key_id```|```IIJAPI_ACCESS_KEY```|
|```secret_access_key```|```IIJAPI_SECRET_KEY```|
|```gis_service_code```|```GISSERVICECODE```|

## Terraform 実行

固有の手順は特にありません。通常通り ```terraform plan```, ```terraform apply```, ```terraform destroy```, etc... を実行してください。

設定例は ```example/``` を参照してください。

## リソース一覧

### ```p2pub_virtual_server```

[仮想サーバー](http://manual.iij.jp/p2/pub/b-1.html)

|項目|内容|値|必須|
|-|-|-|-|
|```type```|仮想サーバー品目| VBxx-xx, VGxx-xx, VDxx-xx |◯|
|```os_type```|OS種別|Linux, Windows|◯|
|```label```|ラベル|任意の文字列||
|```system_storage```|接続するブートデバイス|ibaサービスコード||
|```data_storage```|接続する追加ストレージ|ibb/ibgサービスコードのリスト||
|```private_network```|接続するプライベートネットワーク/V|ivlサービスコードのリスト||
|```enable_global_ip```|グローバルIPアドレス使用の有無|true/false||


```
resource "p2pub_virtual_server" "vs1" {
    type = "VB0-1"
    os_type = "Linux"
    label = "vs-label"
    system_storage = "iba########"
    data_storage = ["ibg#######"]
    private_network = ["ivl########"]
    enable_global_ip = true
}
```

### ```p2pub_system_storage```

[システムストレージ](http://manual.iij.jp/p2/pub/b-3-1.html)

|項目|内容|値|必須|
|-|-|-|-|
|```type```|システムストレージ品目| http://manual.iij.jp/p2/pubapi/59949023.html |◯|
|```label```|ラベル|任意の文字列||
|```root_ssh_key```|rootのSSH公開鍵|||
|```root_password```|rootパスワード|||

```
resource "p2pub_system_storage" "ss1" {
    type = "S30GB_CENTOS7_64"
    label = "ss-label"
    root_ssh_key = "<SSH_PUBLIC_KEY>"
}
```

### ```p2pub_additional_storage```

[追加ストレージ](http://manual.iij.jp/p2/pub/b-3-1.html)

|項目|内容|値|必須|
|-|-|-|-|
|```type```|システムストレージ品目| http://manual.iij.jp/p2/pubapi/59949023.html |◯|
|```label```|ラベル|任意の文字列||

```
resource "p2pub_additional_storage" "as1" {
    type = "B1000GB"
    label = "as-label"
}
```

### ```p2pub_private_network```

[プライベートネットワーク/V](http://manual.iij.jp/p2/pub/b-5-1-1.html)

|項目|内容|値|必須|
|-|-|-|-|
|```label```|ラベル|任意の文字列||

```
resource "p2pub_private_network" "net1" {
    label = "pn-label"
}
```

### ```p2pub_storage_archive```

[ストレージアーカイブ](http://manual.iij.jp/p2/pub/b-4.html)


|項目|内容|値|必須|
|-|-|-|-|
|```archive_size```| アーカイブの容量 | 10〜100 (10GB単位) | ○ |

```
resource "p2pub_storage_archive" "sa1" {}
```


### ```p2pub_global_ip_address```

[グローバルIPアドレス](http://manual.iij.jp/p2/pub/b-5-1-2.html)

|項目|内容|値|必須|
|-|-|-|-|
|```address_num```|IPアドレスの契約数|1~20までの整数|◯|

```
resource "p2pub_global_ip_address" "ip1" {
    address_num = 10
}
```

### ```p2pub_load_balancer```

[FW+LB専有タイプ](http://manual.iij.jp/p2/pub/b-6-7.html)

|項目|内容|値|必須|
|-|-|-|-|
|```type```|FW+LB 専有タイプ品目||◯|
|```redundant```|冗長構成有無|"Yes" "No"|◯|
|```external_type```|ネットワーク種別|"Global", "PrivateStandard", "Private"|◯|
|```external_servicecode```|サービスコード|ivlServicecode||
|```external_masterhost_address```|マスターホストのアドレス|ipaddr||
|```external_slavehost_address```|スレーブホストのアドレス|ipaddr||
|```external_netmask```|ネットマスク|mask||
|```internal_type```|ネットワーク種別|"PrivateStandard", "Private"|◯|
|```internal_servicecode```|サービスコード|ivlServicecode||
|```internal_masterhost_address```|マスターホストのアドレス|ipaddr||
|```internal_slavehost_address```|スレーブホストのアドレス|ipaddr||
|```internal_netmask```|ネットマスク|mask||
|```trafficip_list```|トラフィックIPの一覧|配列|◯|
|```trafficip_list.ipv4_name```|トラフィックIPの名前|"文字列"|◯|
|```filter_in_list```|ファイアウォールのルール一覧（IN）|配列||
|```filter_in_list.source_network```|ソースネットワーク|"IPアドレス/マスク長" "ANY"||
|```filter_in_list.destination_network```|デスティネーションネットワーク|"IPアドレス/マスク長" "ANY"||
|```filter_in_list.destination_port```|デスティネーションポート番号|"数字" "ANY"||
|```filter_in_list.protocol```|プロトコル|"TCP" "UDP"||
|```filter_in_list.action```|ルールにマッチしたパケットに対する処理|"ACCEPT"（許可） "DROP"（破棄） "REJECT"（拒否）||
|```filter_in_list.label```|ラベル|"文字列"||
|```filter_out_list```|ファイアウォールのルール一覧（OUT）|配列||
|```filter_out_list.source_network```|ソースネットワーク|"IPアドレス/マスク長" "ANY"||
|```filter_out_list.destination_network```|デスティネーションネットワーク|"IPアドレス/マスク長" "ANY"||
|```filter_out_list.destination_port```|デスティネーションポート番号|"数字" "ANY"||
|```filter_out_list.protocol```|プロトコル|"TCP" "UDP"||
|```filter_out_list.action```|ルールにマッチしたパケットに対する処理|"ACCEPT"（許可） "DROP"（破棄） "REJECT"（拒否）||
|```filter_out_list.label```|ラベル|"文字列"||
|```administration_server_allow_network_list```|管理画面へのアクセスを許可するIPアドレス|IPアドレスの配列||


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
