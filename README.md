<p align="center">
<img src="https://user-images.githubusercontent.com/52234994/165200623-c60e956b-5805-4088-bf58-f97ebd8ae8b4.png" 
    width="30%" border="0" alt="CA">
</p>

# DeCA
DeCA is a decentralized PKI framework compatible with X.509

DeCA can perform all the key functions of X.509 PKI standard, namely, registering, confirming, revoking and verifying TLS certificates.

DeCA is compatible with the existing PKI standard, namely X.509 It stores, issues and verifies the certificate content in X.509 format instead of creating its custom implementation.

# Background
The security of PKI largely depends on the reliability of these third-party CAs, which is a single point of failure for PKI.
There have been many popular CA violations in the past, among which the centralized operation mode of CA was triggered by the spread of rogue certificates.

Our goal is to completely decentralize the CA pool, and at the same time build our decentralized solution and the established
PKI standard (i.e. X.509) to achieve effective real-world integration.

# System architecture
DeCA proposed a decentralized PKI framework named DeCA by using IPFS and FVM technology,This framework provides data synchronization, hidden synchronization strategy and low-latency synchronization of basic data among decentralized CA groups.

The certificate is stored in decentralize storage by using the characteristics of IPFS and FVM, and cannot be tampered with, thus effectively preventing the third party from attacking.

![image](https://user-images.githubusercontent.com/52234994/192089294-d5891f90-16ac-497d-9efe-a09eb38b0ced.png)

## EVM Practice
### Deploy contract
```
$ bin/deca deploy -u https://goerli.infura.io/v3/81fec17192394537a0408c47ef5ec0b8 -k <private key>
```
### Configuration
Modify the contract information in the configuration file
```
# Contract Configuration
[Contract]
# Contract address
Address = "0x7396fbfa3192325162ca39da7ab7b43bd587750a"
# RPC Address
RpcUrl = "https://goerli.infura.io/v3/81fec17192394537a0408c47ef5ec0b8"
# Account private key - used for signing requests
PrivateKey = "<private key>"
```

### Start CA Service
```
$ bin/deca tls -c configs/config.toml
```
### DPKI SDK Demo
Support decentralized certificate application, certificate online status query, certificate chain verification

[Demo](https://github.com/FlowShield/casdk/tree/dpki/examples/vm-mtls)