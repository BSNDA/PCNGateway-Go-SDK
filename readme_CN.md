# SDK调用参考


## Go语言SDK

## 框架及秘钥支持

> Go语言SDK目前支持官网所有的框架和秘钥组合应用的调用

具体如下：
<escape>
<table >
<tr>
<th rowspan="2">框架</th><th colspan="3">秘钥上传</th><th colspan="3">秘钥托管</th>
</tr>
<tr>
<th>secp256r1</th><th>secp256k1</th><th>SM2</th><th>secp256r1</th><th>secp256k1</th><th>SM2</th>
</tr>
<tr>
<td>Fabric</td><td>支持</td><td></td><td>支持</td><td>支持</td><td></td><td>支持</td>
</tr>
<tr>
<td>FISCO-BCOS</td><td></td><td>支持</td><td>支持</td><td></td><td>支持</td><td>支持</td>
</tr>
<tr>
<td>XuperChain</td><td></td><td></td><td>支持</td><td></td><td></td><td>支持</td>
</tr>
<tr>
<td>CITA</td><td></td><td></td><td></td><td></td><td></td><td>支持</td>
</tr>
</table>
</escape>

* fabric框架应用使用secp256r1、SM2 秘钥的秘钥托管和秘钥上传两种模式;
* FISCO-BCOS框架应用使用secp256k1、SM2 秘钥的秘钥托管和秘钥上传两种模式;
* XuperChain框架应用使用SM2 秘钥的秘钥托管和秘钥上传两种模式;
* CITA框架应用仅支持SM2秘钥的秘钥托管模式；

### 1. 调用前准备

#### 应用参数
> 应用参数是用户在参与应用成功之后在应用详情页面获取，或者由本地设置的一些参数，具体包含以下参数
 * __节点网关接口地址：__ 参与的城市节点的节点网关的调用地址
 * __用户编号：__ 用户的编号
 * __应用编号：__ 参与的应用的编号
 * __应用公钥：__ 用户参与成功之后下载的节点网关的应用公钥
 * __应用私钥：__ 托管类型应用再参与成功之后由BSN生成的应用公钥，非托管应用为在参与应用时上传的公钥所对应的私钥
 * __Https证书：__ 调用https网关接口时使用的https证书

 #### 本地参数
 * __证书存储目录：__ 用来存储非托管应用在调用用户证书登记时生成的用户私钥和证书的目录

### 2. 准备调用

#### 导入sdk包
Fabric 需要引入下面的包
```
import (
    "github.com/BSNDA/PCNGateway-Go-SDK/pkg/client/fabric"
    "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	)
```
FISCO-BCOS 需要引入下面的包
```
import (
    "github.com/BSNDA/PCNGateway-Go-SDK/pkg/client/fisco-bcos"
    "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	)
```
XuperChain 需要引入下面的包
```
import (
    "github.com/BSNDA/PCNGateway-Go-SDK/pkg/client/xuperchain"
    "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	)
```

#### 初始化config
可以初始化一个存储所有配置的对象，这些具体的配置信息应当由调用者根据各自的项目配置或者读取之后，在调用时传入，
在config的`Init`方法中实现了获取一个App基础信息的操作，该操作请不要频繁的调用，该接口将占用您的TPS和流量，可以在项目使用一个静态对象存储`config`在需要时使用。
值得注意的是，在配置证书的时候，应用的证书（即用来签名和验签的证书）是直接传证书内容，而Https的证书是证书的文件路径。
```
	api:="" //节点网关地址
	userCode:="" //用户编号
	appCode :="" //应用编号
	puk :="" //应用公钥
	prk :="" //应用私钥
	mspDir:="" //证书存储目录
	cert :="" //证书
	config,err :=config.NewConfig(api, userCode, appCode, puk, prk, mspDir, cert )
	if err !=nil{
	    log.Fatal(err)
	}
```
#### 初始化Client
使用已经生成的配置对象，调用以下代码可以创建一个Client对象，用来调用节点网关

```
	client,err :=fabric.InitFabricClient(config)
	//client,err :=fisco_bcos.NewFiscoBcosClient(config) //FISCO-BCOS Client
	//client,err :=xuperchain.NewXuperChainClient(config) //XuperChain Client
	if err !=nil{
	    log.Fatal(err)
	}
```

####   调用接口
每一个网关接口已经封装了请求和响应的参数对象，只需要赋值就可直接调用，方法内已经实现了签名和验签的操作。
以下为注册子用户的调用操作，其他类似。
```
	req :=user.RegisterReqDataBody{
	    Name:"abc",
	    Secret:"123456",
 	}

 	res,err :=client.RegisterUser(req)
 	if err !=nil{
	    log.Fatal(err)
 	}

 	if res.Header.Code != 0{
	    log.Fatal( res.Header.Msg)
 	}
```

### 3.一些其他说明

#### 非托管应用的用户身份证书的说明
由于`Fabric`框架的非托管的应用在调用网关进行交易的时候所需要的用户证书需要用户自己生成，其流程是：注册用户->登记用户证书 。在登记用户证书的操作中，会由本地生成一对秘钥，然后通过秘钥导出证书的CSR文件（证书申请文件），调用用户证书
登记接口获取一个有效的证书，使用该证书才能在通过托管应用交易处理接口中正常的发起交易。
需要注意的是在CSR文件中设置CN时，并不直接是注册的Name，而是由Name和AppCode拼接的名称，格式为`Name@AppCode` 。
该操作是在 `FabricClient`的`EnrollUser`方法中实现的。  
而`FISCO-BCOS`框架的非托管应用在进行交易时只需要在本地生成一对符合框架算法的密钥对即可，无需其他操作。

__证书的存储__ 是通过 `util`中的`keystore`和`userstore`实现的，该方法只存储本地文件形式的证书，如果需要其
他形式的证书存储方式。是需要实现具体的接口即可，详细请参考具体的代码。
`keystore`的实现参考了`fabric-sdk-go`中的实现方式，通过计算证书的SKI区分私钥，也可以通过其他方式区分证书和私钥的关系

#### 关于加密
为方便在进行数据交易的上链操作中对数据进行加密解密，SDK中实现了一种对称加密`AES`和一种非对称加密`SM2`算法
其中对称加密为`AES`具体调用如下
```
	data :=[]byte("abc")
	key :=[]byte("123456")

	//CBC模式，秘钥不足16位 PKCS7填充秘钥
	key = keystore.Pkcs7PaddingKey(key)
	//加密
	cr ,err :=keystore.AESCBCPKCS7Encrypt(key,data)
	if err !=nil{
	    t.Fatal(err)
	}

	//转hex输出
	fmt.Println("加密后：",hex.EncodeToString(cr))

	//解密
	data,err = keystore.AESCBCPKCS7Decrypt(key,cr)
	if err !=nil{
	    t.Fatal(err)
	}

	fmt.Println("解密后：",string(data))
```
非对称加密`SM2`，具体如下,在该方法中同时实现了SM2的签名和验签
>非对称加密中由公钥加密，私钥进行解密
```
	puk := ``//公钥
	prik := ``//私钥
	sm, err := sm2.NewSM2Handle(puk, prik)
	if err != nil {
	    t.Fatal(err)
	}
	data :=[]byte("abc")
	cr ,err :=sm.Encrypt(data)
	if err != nil {
	    t.Fatal(err)
	}
	fmt.Println("加密后：",hex.EncodeToString(cr))
	data,err = sm.Decrypt(cr)
	if err != nil {
	    t.Fatal(err)
	}
	fmt.Println("解密后：",string(data))
```
#### 关于秘钥生成
在BSN中，`fabric`框架的密钥格式为`ECDSA`的`secp256r1`曲线，而`fisco-bcos`框架的密钥格式为`SM2`
在用户参与非托管应用时需要生成对应格式的密钥并上传。
下面介绍这两种密钥的生成，秘钥的生成是使用`openssl`生成的，其中`SM2`秘钥的生成需要`openssl`的`1.1.1`及以上版本
> 注：以下命令是在linux环境下执行的
##### 1. ECDSA(secp256r1)的密钥生成
- 生成私钥
```
openssl ecparam -name prime256v1 -genkey -out key.pem
```
- 导出公钥
```
openssl ec -in key.pem -pubout -out pub.pem
```
- 导出pkcs8格式私钥
> 由于部分语言中使用pkcs8格式的密钥比较方便，可以使用下面的命令导出pkcs8格式私钥
> 在本sdk中使用的私钥即为pkcs8格式
```
openssl pkcs8 -topk8 -inform PEM -in key.pem -outform PEM -nocrypt -out key_pkcs8.pem
```
通过以上命令可以生成三个文件
__`key.pem`__ :私钥
__`pub.pem`__ :公钥
__`key_pkcs8.pem`__ :pkcs8格式私钥

##### 2. ECDSA(secp256k1)的密钥生成
- 生成私钥
```
openssl ecparam -name secp256k1 -genkey -out key.pem
```
- 导出公钥
```
openssl ec -in key.pem -pubout -out pub.pem
```
- 导出pkcs8格式私钥
> 由于部分语言中使用pkcs8格式的密钥比较方便，可以使用下面的命令导出pkcs8格式私钥
> 在本sdk中使用的私钥即为pkcs8格式
```
openssl pkcs8 -topk8 -inform PEM -in key.pem -outform PEM -nocrypt -out key_pkcs8.pem
```
通过以上命令可以生成三个文件
__`key.pem`__ :私钥
__`pub.pem`__ :公钥
__`key_pkcs8.pem`__ :pkcs8格式私钥

##### 3.`SM2`格式秘钥生成
首先需要检查`openssl`的版本是否支持`SM2`格式秘钥生成，可以使用下面的命令
```
openssl ecparam -list_curves | grep SM2
```
如果输出以下内容，则表示支持，
```
SM2       : SM2 curve over a 256 bit prime field
```
否则需要去官网下载`1.1.1`或者以上版本，
这是使用的为`1.1.1d`版本，
官网下载地址：[https://www.openssl.org/source/openssl-1.1.1d.tar.gz](https://www.openssl.org/source/openssl-1.1.1d.tar.gz])

- 生成私钥
```
openssl ecparam -genkey -name SM2 -out sm2PriKey.pem
```
- 导出公钥
```
openssl ec -in sm2PriKey.pem -pubout -out sm2PubKey.pem
```
- 导出pkcs8格式私钥
> 由于部分语言中使用pkcs8格式的密钥比较方便，可以使用下面的命令导出pkcs8格式私钥
> 在本sdk中使用的私钥即为pkcs8格式
```
openssl pkcs8 -topk8 -inform PEM -in sm2PriKey.pem -outform pem -nocrypt -out sm2PriKeyPkcs8.pem
```
通过以上命令可以生成三个文件
__`sm2PriKey.pem`__ :私钥
__`sm2PubKey.pem`__ :公钥
__`sm2PriKeyPkcs8.pem`__ :pkcs8格式私钥
