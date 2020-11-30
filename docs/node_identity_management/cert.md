<!--
order: 1
-->

# 证书

IRITA 是许可区块链平台，采用作为公钥基础设施的 `CA` 机制来确保认证性、完整性、与不可抵赖性。机构节点在获准加入 IRITA 网络之前，必须经由指定的 `CA` 签发相应的证书。

## 证书结构

IRITA 使用 `x509协议` 的证书格式，采用两层结构，由上至下依次是：`CA` 根证书、机构节点证书。

## 节点证书生成

### 证书请求生成(节点)

机构节点在申请证书之前，需用 `OpenSSL` 命令生成一个证书请求文件（`.csr` 文件）。

```bash
openssl req -new -key priv.pem -out req.csr -sm3 -sigopt "distid:1234567812345678"
```

::tip

当为验证人创建证书请求时，`key` 的生成方式为：

```bash
irita genkey --out-file=<output-file> --home=<home>
```

当为节点创建证书请求时，`key` 的生成方式为：

```bash
irita genkey --type node --out-file=<output-file> --home=<home>
```

::

示例：

```bash
openssl req -new -key priv.pem -out req.csr -sm3 -sigopt "distid:1234567812345678"
```

```text
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:CN
State or Province Name (full name) [Some-State]:SH
Locality Name (eg, city) []:PD
Organization Name (eg, company) [Internet Widgits Pty Ltd]:TJ
Organizational Unit Name (eg, section) []:TJ
Common Name (e.g. server FQDN or YOUR name) []:IRITA
Email Address []:

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:
An optional company name []:
```

使用以下命令查看证书请求详情：

```bash
openssl req -in req.csr -text
```

结果如下：

```text
Certificate Request:
    Data:
        Version: 1 (0x0)
        Subject: C = CN, ST = SH, L = PD, O = TJ, OU = TJ, CN = IRITA
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:99:91:65:ee:9b:ec:ed:f1:38:39:18:d6:28:33:
                    f7:15:c6:75:24:42:1a:46:59:96:f7:da:a6:ce:f2:
                    ef:b3:5c:0d:5f:35:dd:ea:a1:c0:46:76:e8:77:f7:
                    4c:7c:ab:31:20:96:38:d7:04:27:82:2a:76:b8:41:
                    76:79:8b:59:42
                ASN1 OID: SM2
        Attributes:
            a0:00
    Signature Algorithm: SM2-with-SM3
    Signature Value:
        30:45:02:20:75:b5:0b:d3:35:d2:d5:f2:70:e6:2f:f8:41:93:
        c4:11:07:81:1c:85:53:f7:30:aa:4a:30:68:d7:12:b8:1c:1d:
        02:21:00:9a:c5:14:cd:ae:79:06:df:f9:66:ce:c2:a3:c6:c3:
        c0:13:3a:98:35:28:e9:57:60:bd:ef:b0:73:d3:f1:a0:89
-----BEGIN CERTIFICATE REQUEST-----
MIIBCzCBsgIBADBQMQswCQYDVQQGEwJDTjELMAkGA1UECAwCU0gxCzAJBgNVBAcM
AlBEMQswCQYDVQQKDAJUSjELMAkGA1UECwwCVEoxDTALBgNVBAMMBENTUkIwWTAT
BgcqhkjOPQIBBggqgRzPVQGCLQNCAASZkWXum+zt8Tg5GNYoM/cVxnUkQhpGWZb3
2qbO8u+zXA1fNd3qocBGduh390x8qzEgljjXBCeCKna4QXZ5i1lCoAAwCgYIKoEc
z1UBg3UDSAAwRQIgdbUL0zXS1fJw5i/4QZPEEQeBHIVT9zCqSjBo1xK4HB0CIQCa
xRTNrnkG3/lmzsKjxsPAEzqYNSjpV2C977Bz0/GgiQ==
-----END CERTIFICATE REQUEST-----
```

### 根证书生成(CA机构)

根证书的生成由CA机构操作，如果机构节点请求第三方签发证书，而不是自签证书，可以忽略该步骤。

```bash
##生成根证书秘钥
openssl ecparam -genkey -name SM2 -out root.key
##生成根证书
openssl req -new -x509 -sm3 -sigopt "distid:1234567812345678" -key root.key -out root.crt -days 365
```

输出如下：

```text
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:CN
State or Province Name (full name) [Some-State]:SH
Locality Name (eg, city) []:SH
Organization Name (eg, company) [Internet Widgits Pty Ltd]:TJ
Organizational Unit Name (eg, section) []:TJ
Common Name (e.g. server FQDN or YOUR name) []:IRITA
Email Address []:
```

使用以下命令查看证书请求详情：

```bash
openssl x509 -in root.crt --text
```

输出如下：

```text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            53:1c:39:a7:1b:56:2b:c3:ce:56:7e:a8:7b:76:a5:7c:fe:54:03:f0
        Signature Algorithm: SM2-with-SM3
        Issuer: C = CN, ST = SH, L = SH, O = TJ, OU = TJ, CN = IRITA
        Validity
            Not Before: Jun 29 02:48:37 2020 GMT
            Not After : Jun 29 02:48:37 2021 GMT
        Subject: C = CN, ST = SH, L = SH, O = TJ, OU = TJ, CN = IRITA
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:72:99:1d:8d:b0:0a:4e:ee:80:55:b5:ac:28:2b:
                    79:bd:f7:61:5a:46:a0:3d:3d:dd:2f:0d:84:87:ca:
                    26:78:8b:31:70:f3:45:9e:c4:cf:94:4c:d9:2c:f9:
                    a0:0b:5f:6a:2e:73:b2:15:87:fa:a7:31:fa:44:ee:
                    cd:32:a7:6e:b7
                ASN1 OID: SM2
        X509v3 extensions:
            X509v3 Subject Key Identifier:
                F2:B7:2D:DD:95:DA:25:D6:B4:CF:6B:16:89:A7:BB:07:CB:34:FB:D4
            X509v3 Authority Key Identifier:
                F2:B7:2D:DD:95:DA:25:D6:B4:CF:6B:16:89:A7:BB:07:CB:34:FB:D4
            X509v3 Basic Constraints: critical
                CA:TRUE
    Signature Algorithm: SM2-with-SM3
    Signature Value:
        30:45:02:21:00:e3:58:89:0f:37:57:06:76:23:f1:60:ce:da:
        7e:8a:6c:5e:81:fb:ca:76:23:c9:31:e3:95:d9:ef:c8:9e:59:
        6d:02:20:3d:e7:71:00:2c:65:bf:b8:e0:67:02:bd:d1:d1:7f:
        8d:61:a4:2d:db:26:f7:40:fb:a8:d9:f8:ff:3c:77:7e:3b
-----BEGIN CERTIFICATE-----
MIIB9TCCAZugAwIBAgIUUxw5pxtWK8POVn6oe3alfP5UA/AwCgYIKoEcz1UBg3Uw
UDELMAkGA1UEBhMCQ04xCzAJBgNVBAgMAlNIMQswCQYDVQQHDAJTSDELMAkGA1UE
CgwCVEoxCzAJBgNVBAsMAlRKMQ0wCwYDVQQDDARDU1JCMB4XDTIwMDYyOTAyNDgz
N1oXDTIxMDYyOTAyNDgzN1owUDELMAkGA1UEBhMCQ04xCzAJBgNVBAgMAlNIMQsw
CQYDVQQHDAJTSDELMAkGA1UECgwCVEoxCzAJBgNVBAsMAlRKMQ0wCwYDVQQDDARD
U1JCMFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEcpkdjbAKTu6AVbWsKCt5vfdh
WkagPT3dLw2Eh8omeIsxcPNFnsTPlEzZLPmgC19qLnOyFYf6pzH6RO7NMqdut6NT
MFEwHQYDVR0OBBYEFPK3Ld2V2iXWtM9rFomnuwfLNPvUMB8GA1UdIwQYMBaAFPK3
Ld2V2iXWtM9rFomnuwfLNPvUMA8GA1UdEwEB/wQFMAMBAf8wCgYIKoEcz1UBg3UD
SAAwRQIhAONYiQ83VwZ2I/Fgztp+imxegfvKdiPJMeOV2e/InlltAiA953EALGW/
uOBnAr3R0X+NYaQt2yb3QPuo2fj/PHd+Ow==
-----END CERTIFICATE-----
```

### 证书签发

证书的签发有两种渠道，一种是用户使用自签证书部署多节点网络，一种是由第三方CA机构签发证书：

- 自签证书

  ```bash
  openssl x509 -req -in <req.csr> -out node0.crt -sm3 -sigopt "distid:1234567812345678" -vfyopt "distid:1234567812345678" -CA <root.crt> -CAkey <root.key> -CAcreateserial
  ```

- 第三方CA机构签发

  由节点机构从CA机构获取证书。

## 证书续期

当节点证书过期后，需重新签发新的证书。流程与上述证书生成过程一致。
