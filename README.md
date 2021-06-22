# PqAuthKEM4Chain
# About The Function of Chaincodes
The main implementation of the PQ user authentication and key
exchange system consists of two chaincodes, namely PqCa and
PqUser. As shown in Table 2, the PqCa chaincode provides
APIs to generate the PQ key pairs together with the related
certificates (for the CA, the user identity key and the user KEM
key), query one certificate (based on a given serial number),
verify one certificate (and its parent certificates in the certificate
chain), revoke one certificate, encapsulate and decapsulate the
shared secret between users. While the PqUser chaincode
invokes the corresponding APIs in PqCa to query/verify one
certificate and provides APIs to perform the encapsulation and
decapsulation operations for the users.
########
The data structure in the off-chain private databases is
designed relatively simply for further development. In the off-
chain private databases (e.g., MySQL), one key pair, the
corresponding certificate and related information (e.g., password)
are all stored in one line, and the certificate serial number is set as
the key of the line.
#######
Besides the chaincodes, we also develop client codes to help
developers create their applications based on the PQ user
authentication and key exchange system. As shown in Table 2,
there are two client codes, namely PqCa_Cli and PqUser_Cli,
which are written in Go and used to invoke the APIs in the
corresponding chaincodes. The configuration files for the client
codes to invoke the chaincodes are also provided
