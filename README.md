# group_ecdhm
Group Elliptic Curve Diffie-Hellman-Merkle key exchange implementation in golang

The aim of this library is to provide a method of satisfying key exchange with
groups of greater than 2 participants. This key exchange is mediated via a server
node, therefore the minimum amount of participants in key exchange between two
parties is 3 (party A, party B, and the server S).

Because I like elliptic curves, I am implementing this with the expectation that
elliptical curve cryptography is used. Any curve can be used, but I will be
implementing this by expecting that all curves satisfy the `elliptic.Curve`
interface from the golang crypto package.

For now, I will NOT include the setting up of TLS for intercommunication.
Perhaps later, but for now I'm just playing with key exchange protocols.

Cheers :)
