To avoid the Shor’s attack, researchers have developed PQ public-key algorithms relying on different mathematics problems other than the IF, DL or ECDL problems. Based on the mathematics problems, the PQ public-key algorithms can be classified into different families introduced as follows.
2.3.1 PQ Digital Signature Algorithms. In the current round of NIST call, there are six PQ digital signature algorithms, which can be divided into three following groups.

●	Lattice-based cryptosystems. This kind of cryptographic schemes are based on lattices, which are sets of points in n-dimensional spaces with a periodic structure. Lattice-based security schemes rely on presumed hardness of basic lattice problems like the Shortest Vector Problem (SVP) (i.e., find the shortest non-zero vector within a lattice), the Closest Vector Problem (CVP) (i.e., find a lattice vector that minimizes the distance from another target lattice) and the Shortest Independent Vectors Problem (SIVP) derived from SVP and CVP. In the current round, FALCON and CRYSTALS-DILITHIUM are lattice-based candidate algorithms and separately based on the Short Integer Solution problem (SIS) [1] over NTRU lattices and “Fiat-Shamir with Aborts” technique [2]. Lattice-based algorithms are promising quantum-resistant solutions with relatively efficient implementations, balanced key/signature sizes and strong security properties.

●	Hash-based cryptosystems. The security of these schemes depends on the security of underlying hash function instead of on the hardness of a mathematical problem. Currently, Picnic is a hash-based algorithm and makes use of hash-based zero-knowledge proof technique [3], while SPHINCS+ is a stateless hash-based signature algorithm built on Merkle tree. Normally, hash-based schemes have very small key sizes but suffer from large signature sizes.

●	Multivariate-based cryptosystems. Another family of problems that are used by some NIST PQ signature algorithms is related to solving multivariate quadratic equations over finite fields which is an NP-hard problem. Multivariate PQ schemes often lead to excessive key sizes, and the NIST multivariate-based signature algorithms are Rainbow and GeMSS.
As shown in Table 1 attached as Appendix A, we summarize the details of all the six PQ digital signature algorithms. By using different key sizes and parameter sets, all the algorithms can achieve different NIST security levels (i.e., bits-of-security level), which is defined as the effort required by a classical computer to perform a brute-force attack on a given-length cryptographic key. Normally NIST security levels 1~5 approximately imply 128/160/192/224/256-bits-of security levels.
2.3.2 PQ KEM Algorithms. In the current round of NIST call, there are nine PQ KEM algorithms, which can be divided into three following groups.

●	Lattice-based cryptosystems. The lattice problems that NIST PQ KEM algorithms depend on are Learning With Errors (LWE) [4], Ring Learning With Errors (Ring-LWE) [5] and Module Learning with Rounding (MLWR) problem [6] with different lattices. And the list of NIST PQ KEM algorithms includes KYBER, NTRU, SABER, FrodoKEM and NTRU Prime, which have similar advantages just like lattice-based PQ signature algorithms.

●	Code-based cryptosystems. This family of algorithms are essentially based on the theory that supports error-correction codes. In the current round of NIST call, Classic McEliece is an example of code-based cryptosystem that dates back from 1970s and whose security is based on the syndrome decoding problem [7]. Classic McEliece provides fast encapsulation and relatively fast decapsulation, but is burdened with huge sizes of public and private key sizes. To ease this issue, the PQ KEM algorithms BIKE and HQC based on different codes have been proposed and selected as the alternate candidate algorithms in the current round of NIST.

●	Supersingular Elliptic Curve Isogeny (SECI) cryptosystems. The SECI cryptosystems are based on the isogeny protocol for ordinary elliptic curves and enhanced to withstand the quantum attack detailed in [8]. The only one isogeny-based PQ KEM algorithm passed to the current round of NIST call is SIKE, which is founded on pseudo-random walks in supersingular isogeny graphs.

We also summarize the details of all the nine PQ KEM algorithms in Table 2 attached as Appendix A. Just like the PQ signature algorithms, all the PQ KEM algorithms can achieve different NIST security levels by using different key lengths and related parameters.

|Alg.types| Algorithms names | Alg.subtypes| Claimed NIST security level| Public key size (bytes)|Private key size (bytes) | Signature size (bytes)|
|--|--|--|--|--|--|--|
| Digital signature |CRYSTALS-DILITHIUM  | Lattice-based     | 2, 3, 5|1312~2592     |2528~4864  |2420~4595    |
| Digital signature | FALCON             | Lattice-based     | 1, 5   |897~1793      |1281~2305  |690~1330     |
| Digital signature | Rainbow            | Multivariate-based| 1, 3, 5|60192~1930600 |64~1408736 |66~212       |
| Digital signature | Picnic             | Hash-based        | 1, 3, 5|33~65         |49~97      |34036~209510 |
| Digital signature | SPHINCS+           | Hash-based        | 1, 3, 5|32~64         |64~128     |7856~49856   |
| Digital signature | GeMSS              | Multivariate-based| 1, 3, 5|352190~3135590|128~256    |258~600      |
                                                       Table 1
                                                       

|Alg.types| Algorithms names | Alg.subtypes| Claimed NIST security level| Public key size (bytes)|Private key size (bytes) | Ciphertext size (bytes)|Shared secret size (bytes)|
|--|--|--|--|--|--|--|--|
| KEM |Classic McEliece  | Code-based        | 1, 3, 5   |261120~1357824   |6452~14080  |128~240    |32         |
| KEM |CRYSTALS-KYBER    | Lattice-based     | 1, 3, 5   |800~1568         |1632~3168   |768~1568   |32         |
| KEM |NTRU              | Lattice-based     | 1, 3, 5   |699~1138         |935~1450    |699~1138   |32         |
| KEM |SABER             | Lattice-based     | 1, 3, 5   |672~1312         |1568~3040   |736~1472   |32         |
| KEM |BIKE              | Code-based        | 1, 3      |2542~6206        |3110~13236  |2542~6206  |32         |
| KEM |FrodoKEM          | Lattice-based     | 1, 3, 5   |9616~21520       |19888~43088 |9720~21632 |16, 24, 32 |
| KEM |HQC               | Code-based        | 1, 3, 5   |2249~7425        |2289~7285   |4481~14469 |64         |
| KEM |NTRU Prime        | Lattice-based     | 2, 3, 4   |897~1322         |1125~1999   |897~1184   |32         |
| KEM |SIKE              | SECI-based        | 1, 2, 3, 5|197~564          |350~644     |236~596    |16, 24, 32 |
                                                       Table 2
# References
[1]	M. Ajtai. 1996. Generating hard instances of lattice problems (extended abstract). In STOC. Philadelphia, PA, USA, 99–108.

[2]	Vadim Lyubashevsky. 2009. Fiat-Shamir with Aborts: Applications to Lattice and Factoring-Based Signatures. In ASIACRYPT. Tokyo, Japan, 598–616.

[3]	Melissa Chase, David Derler, Steven Goldfeder, Claudio Orlandi, Sebastian Ramacher, Christian Rechberger, Daniel Slamanig, Greg Zaverucha. 2017. Post-Quantum Zero-Knowledge and Signatures from Symmetric-Key Primitives. In CCS. Dallas, Texas, USA, 1825–1842.

[4]	Oded Regev. 2009. On lattices, learning with errors, random linear codes, and cryptography. Journal of the ACM 56, 6, (2009), 1–40.

[5]	Vadim Lyubashevsky, Chris Peikert, Oded Regev. 2013. On Ideal Lattices and Learning with Errors over Rings. Journal of the ACM 60, 6, (2013), 1–35.

[6]	Adeline Langlois, Damien Stehle. 2015. Worst-case to average-case reductions for module lattices. Designs, Codes and Cryptography 75, (2015), 565–599.

[7]	Elwyn R. Berlekamp, Robert J. McEliece, Henk C. A. Van Tilborg. 1978. On the inherent intractability of certain coding problems (Corresp.). IEEE Transactions on Information Theory 24, 3, (1978), 384–386.

[8]	Andrew Childs, David Jao, Vladimir Soukharev. 2013. Constructing elliptic curve isogenies in quantum subexponential time. Journal of Mathematical Cryptology 8, (2013), 1–29.

