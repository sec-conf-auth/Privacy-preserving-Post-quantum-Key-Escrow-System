# PqAuthKEM4Chain
# About the Function of Chaincodes
The main implementation of the PQ user authentication and key
exchange system consists of two chaincodes, namely `PqCa` and
`PqUser`. The `PqCa` chaincode provides
APIs to generate the PQ key pairs together with the related
certificates (for the CA, the user identity key and the user KEM
key), query one certificate (based on a given serial number),
verify one certificate (and its parent certificates in the certificate
chain), revoke one certificate, encapsulate and decapsulate the
shared secret between users. While the `PqUser` chaincode
invokes the corresponding APIs in `PqCa` to query/verify one
certificate and provides APIs to perform the encapsulation and
decapsulation operations for the users.

The data structure in the off-chain private databases is
designed relatively simply for further development. In the off-
chain private databases (e.g., MySQL), one key pair, the
corresponding certificate and related information (e.g., password)
are all stored in one line, and the certificate serial number is set as
the key of the line.

# About the Function of Client Codes
Besides the chaincodes, we also develop client codes to help
developers create their applications based on the PQ user
authentication and key exchange system. 
There are two client codes, namely `PqCa_Cli` and `PqUser_Cli`,
which are written in Go and used to invoke the APIs in the
corresponding chaincodes. The configuration file (e.g.,`connection-profile.yaml`)for the client
codes to invoke the chaincodes is also provided.

# About the Use of Docker Image
In the experiments , we utilize the liboqs 0.6.0 library [1] together with its Go wrapper [2] to generate public/private key pair, encapsulate and decapsulate the shared secret for all the related post-quantum KEM algorithms. Since the liboqs library is incompatible with the
native Hyperledger Fabric docker image for chaincode execution,
we build a new docker image integrated with liboqs based on Ubuntu 18.04 and use it as the execution environment of our
chaincodes. The versions of docker and docker-compose we use
are 19.03.13 and 1.25.0-rc1. For developers, the configuration
files for the client codes, the docker file to generate the new
docker image and information about how to use the new image
are also available.

1.In ccenv dockerfile, use the `Ubuntu18.04` as the docker image environment, download and install `liboqs library` required for the experiment, install and configure the Go env(`go1.15.7.linux-amd64.tar.gz`).

2.Build the dockerfile.

   ```javascript
   docker built -t ccenv:latest ./ 
   ```

3.Modify core.yaml.
Set the `chaincode.builder` and `chaincode.golang.runtime` as `ccenv:latest`, and set the `chaincode.golang.dynamicLink` as `true`. The part of the code that needs to be modified in core.yaml is set as follows:

```html
###############################################################################
#
#    Chaincode section
#
###############################################################################
chaincode:

    # The id is used by the Chaincode stub to register the executing Chaincode
    # ID with the Peer and is generally supplied through ENV variables
    # the `path` form of ID is provided when installing the chaincode.
    # The `name` is used for all other requests and can be any string.
    id:
        path:
        name:

    # Generic builder environment, suitable for most chaincode types
    builder: ccenv:latest

    # Enables/disables force pulling of the base docker images (listed below)
    # during user chaincode instantiation.
    # Useful when using moving image tags (such as :latest)
    pull: false

    golang:
        # golang will never need more than baseos
        runtime: ccenv:latest
        # whether or not golang chaincode should be linked dynamically
        dynamicLink: true
   ```
   
4.Start the fabric network and deploy chaincodes.
 
# About the Experiment Data
Our experiments involve three VMs on the local server, each one of which is based on Ubuntu 18.04 configured with eight CPU cores of Intel Xeon Gold 6131 throttled to 2.6GHz and 16GB memory, and we deploy one CA peer (for CA administrator) and two user peers (for users Alice and Bob) in one channel separately running on the three VMs. The system is built on Hyperledger Fabric 2.2.3, and the chaincodes together with the client codes are written in Go 1.15.7. 

We divide all the PQ algorithms into three groups, namely level 1-2, level 3, level 4-5, which separately indicate low, medium and high security levels. The specified PQ algorithm names are summarized by groups in Table 1. In our experiments, we choose one PQ signature algorithm and one KEM algorithm in the same security group (i.e., the same line of Table 1).

       Table 1：Specific PQ Algorithm names in our analysis grouped by NIST security levels.
|Levels | Sig. Alg. Names|KEM Alg. Names|
|--|--|--|
|Level 1~2| (1).Dilithium2-AES<br> (2).Falcon-512<br> (3).picnic3_L1 <br> (4).SPHINCS+-SHA256-128f-robust <br>(5).SPHINCS+-SHA256-128f-simple <br>| ▪Kyber512-90s<br> ▪NTRU-HPS-2048-509<br>▪LightSaber-KEM<br>▪BIKE1-L1-FO<br>▪FrodoKEM-640-SHAKE<br>▪HQC-128<br>▪ntrulpr653<br>▪sntrup653<br>▪SIKE-p434<br>▪SIKE-p503<br>|
| Level 3 | (6).Dilithium3-AES<br>(7).picnic3_L3<br>(8).SPHINCS+-SHA256-192f-robust<br>(9).SPHINCS+-SHA256-192f-simple<br>    | ▪Kyber768-90s<br>▪NTRU-HPS-2048-677<br>▪Saber-KEM<br>▪BIKE1-L3-FO<br>▪FrodoKEM-976-SHAKE<br>▪HQC-192<br>▪ntrulpr761<br>▪sntrup761<br>▪SIKE-p610<br> |
| Level 4~5| (10).Dilithium5-AES<br>(11).Falcon-1024<br>(12).picnic3_L5<br>(13).SPHINCS+-SHA256-256f-robust<br>(14).SPHINCS+-SHA256-256f-simple<br>| ▪Kyber1024-90s<br>▪NTRU-HPS-4096-821<br>▪FireSaber-KEM<br>▪FrodoKEM-1344-SHAKE<br>▪HQC-256<br>▪ntrulpr857<br>▪sntrup857<br>▪SIKE-p751|

                         
The full list of the experiment data and related explanation can be found in the folder "ExperimentData".

# About the PQ Algorithms
In the third (current) round of NIST call for PQ public-key algorithms, there are three digital signature (i.e., CRYSTALS-DILITHIUM , FALCON , Rainbow ) and four KEM (i.e., Classic McEliece , CRYSTALS-KYBER , NTRU , SABER ) finalist algorithms, which will continue to be reviewed for consideration for standardization at the conclusion of the third round. While at the same time, there are also three digital signature (i.e., GeMSS , Picnic , SPHINCS+ ) and five KEM (i.e., BIKE , NTRU Prime , FrodoKEM , SIKE , HQC ) alternate candidate algorithms, which are potentially standardized but most likely will not occur at the end of the third round. These alternate candidate algorithms may be reconsidered for various reasons (e.g., better performance, higher security level, broader range of hardness assumptions) in a fourth round of evaluation held by NIST.All of these are summarized in PQ Algorithms.

# References
[1] Liboqs, https://github.com/open-quantum-safe/liboqs

[2] Liboqs-go, https://github.com/open-quantum-safe/liboqs-go
