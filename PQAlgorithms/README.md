To avoid the Shor’s attack, researchers have developed PQ public-key algorithms relying on different mathematics problems other than the IF, DL or ECDL problems. Based on the mathematics problems, the PQ public-key algorithms can be classified into different families introduced as follows.



 **PQ KEM Algorithms**：In the current round of NIST call, there are nine PQ KEM algorithms, which can be divided into three following groups.

●	Lattice-based cryptosystems. The lattice problems that NIST PQ KEM algorithms depend on are Learning With Errors (LWE) [4], Ring Learning With Errors (Ring-LWE) [5] and Module Learning with Rounding (MLWR) problem [6] with different lattices. And the list of NIST PQ KEM algorithms includes KYBER, NTRU, SABER, FrodoKEM and NTRU Prime, which have similar advantages just like lattice-based PQ signature algorithms.

●	Code-based cryptosystems. This family of algorithms are essentially based on the theory that supports error-correction codes. In the current round of NIST call, Classic McEliece is an example of code-based cryptosystem that dates back from 1970s and whose security is based on the syndrome decoding problem [7]. Classic McEliece provides fast encapsulation and relatively fast decapsulation, but is burdened with huge sizes of public and private key sizes. To ease this issue, the PQ KEM algorithms BIKE and HQC based on different codes have been proposed and selected as the alternate candidate algorithms in the current round of NIST.

●	Supersingular Elliptic Curve Isogeny (SECI) cryptosystems. The SECI cryptosystems are based on the isogeny protocol for ordinary elliptic curves and enhanced to withstand the quantum attack detailed in [8]. The only one isogeny-based PQ KEM algorithm passed to the current round of NIST call is SIKE, which is founded on pseudo-random walks in supersingular isogeny graphs.

We also summarize the details of all the nine PQ KEM algorithms in Table 2 . Just like the PQ signature algorithms, all the PQ KEM algorithms can achieve different NIST security levels by using different key lengths and related parameters.

                                   
       Table 2：Details of PQ KEM algorithms in the current round of NIST call for national standards.
| Algorithms names | Alg.subtypes| Claimed NIST security level| Public key size (bytes)|Private key size (bytes) | Ciphertext size (bytes)|Shared secret size (bytes)|
|--|--|--|--|--|--|--|
|Classic McEliece  | Code-based        | 1, 3, 5   |261120~1357824   |6452~14080  |128~240    |32         |
|CRYSTALS-KYBER    | Lattice-based     | 1, 3, 5   |800~1568         |1632~3168   |768~1568   |32         |
|NTRU              | Lattice-based     | 1, 3, 5   |699~1138         |935~1450    |699~1138   |32         |
|SABER             | Lattice-based     | 1, 3, 5   |672~1312         |1568~3040   |736~1472   |32         |
|BIKE              | Code-based        | 1, 3      |2542~6206        |3110~13236  |2542~6206  |32         |
|FrodoKEM          | Lattice-based     | 1, 3, 5   |9616~21520       |19888~43088 |9720~21632 |16, 24, 32 |
|HQC               | Code-based        | 1, 3, 5   |2249~7425        |2289~7285   |4481~14469 |64         |
|NTRU Prime        | Lattice-based     | 2, 3, 4   |897~1322         |1125~1999   |897~1184   |32         |
|SIKE              | SECI-based        | 1, 2, 3, 5|197~564          |350~644     |236~596    |16, 24, 32 |
                                                      
# References
[1]	M. Ajtai. 1996. Generating hard instances of lattice problems (extended abstract). In STOC. Philadelphia, PA, USA, 99–108.

[2]	Vadim Lyubashevsky. 2009. Fiat-Shamir with Aborts: Applications to Lattice and Factoring-Based Signatures. In ASIACRYPT. Tokyo, Japan, 598–616.

[3]	Melissa Chase, David Derler, Steven Goldfeder, Claudio Orlandi, Sebastian Ramacher, Christian Rechberger, Daniel Slamanig, Greg Zaverucha. 2017. Post-Quantum Zero-Knowledge and Signatures from Symmetric-Key Primitives. In CCS. Dallas, Texas, USA, 1825–1842.

[4]	Oded Regev. 2009. On lattices, learning with errors, random linear codes, and cryptography. Journal of the ACM 56, 6, (2009), 1–40.

[5]	Vadim Lyubashevsky, Chris Peikert, Oded Regev. 2013. On Ideal Lattices and Learning with Errors over Rings. Journal of the ACM 60, 6, (2013), 1–35.

[6]	Adeline Langlois, Damien Stehle. 2015. Worst-case to average-case reductions for module lattices. Designs, Codes and Cryptography 75, (2015), 565–599.

[7]	Elwyn R. Berlekamp, Robert J. McEliece, Henk C. A. Van Tilborg. 1978. On the inherent intractability of certain coding problems (Corresp.). IEEE Transactions on Information Theory 24, 3, (1978), 384–386.

[8]	Andrew Childs, David Jao, Vladimir Soukharev. 2013. Constructing elliptic curve isogenies in quantum subexponential time. Journal of Mathematical Cryptology 8, (2013), 1–29.

