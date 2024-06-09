<p align="center">
<img src="https://user-images.githubusercontent.com/52234994/165200623-c60e956b-5805-4088-bf58-f97ebd8ae8b4.png" 
    width="30%" border="0" alt="DeCA Logo">
</p>

<h1 align="center">DeCA: Decentralized PKI Framework</h1>

DeCA is a decentralized Public Key Infrastructure (PKI) framework fully compatible with the X.509 standard. It performs all essential PKI functions, including registering, confirming, revoking, and verifying TLS certificates, while leveraging decentralized technologies to enhance security and resilience.

## Key Features

- **Full Compatibility:** DeCA integrates seamlessly with the existing X.509 PKI standard, ensuring smooth interoperability by storing, issuing, and verifying certificates in X.509 format.
- **Decentralized Storage:** Utilizes IPFS and FVM for tamper-proof, decentralized storage of certificates, protecting against third-party attacks and rogue certificates.
- **High Security:** Eliminates single points of failure by decentralizing the CA pool, enhancing the overall security of the PKI system.

## Background

Traditional PKI systems rely heavily on centralized Certificate Authorities (CAs), which pose significant security risks as single points of failure. Historical breaches and the spread of rogue certificates highlight the vulnerabilities of centralized CAs.

DeCA addresses these issues by decentralizing the CA pool while maintaining compatibility with the established X.509 PKI standard, ensuring robust and secure real-world integration.

## System Architecture

The DeCA framework utilizes IPFS and FVM technologies to provide a decentralized PKI system. Key components include:

- **Data Synchronization:** Ensures efficient and low-latency synchronization of basic data among decentralized CA groups.
- **Decentralized Storage:** Certificates are stored in a tamper-proof, decentralized manner, leveraging IPFS and FVM to prevent third-party attacks.

<p align="center">
<img src="https://github.com/FlowShield/DeCA/assets/34047788/5d843aad-6d2f-4c64-a08d-2a17b07bd82e" 
    width="90%" border="0" alt="DeCA System Architecture">
</p>

## Getting Started

### Installation

To install DeCA, use the following commands:

```sh
$ go get github.com/FlowShield/deca
$ make
$ bin/ca tls -c configs/config.toml
```

### SDK Installation

To install the DeCA SDK, use:

```sh
$ go get github.com/FlowShield/deca
```

### Usage

The classic use case of the DeCA SDK involves clients and servers utilizing certificates issued by the CA Center for encrypted communication. For detailed usage, refer to our [SDK Examples](https://github.com/FlowShield/casdk/tree/main/examples).

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

## Contributing

We welcome contributions from the community! Please refer to our [CONTRIBUTING](CONTRIBUTING.md) guidelines for more information.
