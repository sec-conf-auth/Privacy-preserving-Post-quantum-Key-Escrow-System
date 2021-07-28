To avoid the Shor’s attack, researchers have developed PQ public-key algorithms relying on different mathematics problems other than the IF, DL or ECDL problems. Based on the mathematics problems, the PQ public-key algorithms can be classified into different families introduced as follows.
2.3.1 PQ Digital Signature Algorithms. In the current round of NIST call, there are six PQ digital signature algorithms, which can be divided into three following groups.

●	Lattice-based cryptosystems. This kind of cryptographic schemes are based on lattices, which are sets of points in n-dimensional spaces with a periodic structure. Lattice-based security schemes rely on presumed hardness of basic lattice problems like the Shortest Vector Problem (SVP) (i.e., find the shortest non-zero vector within a lattice), the Closest Vector Problem (CVP) (i.e., find a lattice vector that minimizes the distance from another target lattice) and the Shortest Independent Vectors Problem (SIVP) derived from SVP and CVP. In the current round, FALCON and CRYSTALS-DILITHIUM are lattice-based candidate algorithms and separately based on the Short Integer Solution problem (SIS) [39] over NTRU lattices and “Fiat-Shamir with Aborts” technique [40]. Lattice-based algorithms are promising quantum-resistant solutions with relatively efficient implementations, balanced key/signature sizes and strong security properties.

●	Hash-based cryptosystems. The security of these schemes depends on the security of underlying hash function instead of on the hardness of a mathematical problem. Currently, Picnic is a hash-based algorithm and makes use of hash-based zero-knowledge proof technique [41], while SPHINCS+ is a stateless hash-based signature algorithm built on Merkle tree. Normally, hash-based schemes have very small key sizes but suffer from large signature sizes.

●	Multivariate-based cryptosystems. Another family of problems that are used by some NIST PQ signature algorithms is related to solving multivariate quadratic equations over finite fields which is an NP-hard problem. Multivariate PQ schemes often lead to excessive key sizes, and the NIST multivariate-based signature algorithms are Rainbow and GeMSS.
As shown in Table 1 attached as Appendix A, we summarize the details of all the six PQ digital signature algorithms. By using different key sizes and parameter sets, all the algorithms can achieve different NIST security levels (i.e., bits-of-security level), which is defined as the effort required by a classical computer to perform a brute-force attack on a given-length cryptographic key. Normally NIST security levels 1~5 approximately imply 128/160/192/224/256-bits-of security levels.
2.3.2 PQ KEM Algorithms. In the current round of NIST call, there are nine PQ KEM algorithms, which can be divided into three following groups.

●	Lattice-based cryptosystems. The lattice problems that NIST PQ KEM algorithms depend on are Learning With Errors (LWE) [42], Ring Learning With Errors (Ring-LWE) [43] and Module Learning with Rounding (MLWR) problem [44] with different lattices. And the list of NIST PQ KEM algorithms includes KYBER, NTRU, SABER, FrodoKEM and NTRU Prime, which have similar advantages just like lattice-based PQ signature algorithms.

●	Code-based cryptosystems. This family of algorithms are essentially based on the theory that supports error-correction codes. In the current round of NIST call, Classic McEliece is an example of code-based cryptosystem that dates back from 1970s and whose security is based on the syndrome decoding problem [45]. Classic McEliece provides fast encapsulation and relatively fast decapsulation, but is burdened with huge sizes of public and private key sizes. To ease this issue, the PQ KEM algorithms BIKE and HQC based on different codes have been proposed and selected as the alternate candidate algorithms in the current round of NIST.

●	Supersingular Elliptic Curve Isogeny (SECI) cryptosystems. The SECI cryptosystems are based on the isogeny protocol for ordinary elliptic curves and enhanced to withstand the quantum attack detailed in [46]. The only one isogeny-based PQ KEM algorithm passed to the current round of NIST call is SIKE, which is founded on pseudo-random walks in supersingular isogeny graphs.

We also summarize the details of all the nine PQ KEM algorithms in Table 2 attached as Appendix A. Just like the PQ signature algorithms, all the PQ KEM algorithms can achieve different NIST security levels by using different key lengths and related parameters.

<table>
   <tr>
      <td>Alg.</td>
      <td>Algorithms names</td>
      <td>Alg.</td>
      <td>Claimed NIST security level</td>
      <td>Public key size (bytes)</td>
      <td>Private key size (bytes)</td>
      <td>Signature</td>
   </tr>
   <tr>
      <td>types</td>
      <td></td>
      <td>subtypes</td>
      <td></td>
      <td></td>
      <td></td>
      <td>size (bytes)</td>
   </tr>
   <tr>
      <td>Digital signature</td>
      <td>CRYSTALS-DILITHIUM</td>
      <td>Lattice-based</td>
      <td>2, 3, 5</td>
      <td>1312~2592</td>
      <td>2528~4864</td>
      <td>2420~4595</td>
   </tr>
   <tr>
      <td></td>
      <td>FALCON</td>
      <td>Lattice-based</td>
      <td>1, 5</td>
      <td>897~1793</td>
      <td>1281~2305</td>
      <td>690~1330</td>
   </tr>
   <tr>
      <td></td>
      <td>Rainbow</td>
      <td>Multivariate-based</td>
      <td>1, 3, 5</td>
      <td>60192~1930600</td>
      <td>64~1408736</td>
      <td>66~212</td>
   </tr>
   <tr>
      <td></td>
      <td>Picnic</td>
      <td>Hash-based</td>
      <td>1, 3, 5</td>
      <td>33~65</td>
      <td>49~97</td>
      <td>34036~209510</td>
   </tr>
   <tr>
      <td></td>
      <td>SPHINCS+</td>
      <td>Hash-based</td>
      <td>1, 3, 5</td>
      <td>32~64</td>
      <td>64~128</td>
      <td>7856~49856</td>
   </tr>
   <tr>
      <td></td>
      <td>GeMSS</td>
      <td>Multivariate-based</td>
      <td>1, 3, 5</td>
      <td>352190~3135590</td>
      <td>128~256</td>
      <td>258~600</td>
   </tr>
   <tr>
      <td></td>
   </tr>
</table>


