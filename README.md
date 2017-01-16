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

## Details

The current idea is to keep shared secrets between each client (C) and the server
(S), thus for N clients the server will have the shared-secrets {S-C<sub>1</sub>
S-C<sub>2</sub> S-C<sub>3</sub> ... S-C<sub>N</sub> }.

This is pretty much the simplest thing ever, but I'm only implementing this stuff
for sketching out some stuff for a future project. I don't expect this library to
actually be useful, but who knows, maybe it will help you with something.

## Protocol
The connection is pretty simple and occurs between clients and a server

1. Connection request from client which can include out of band credentials or not
(this is more for a future project and if you do send credentials, ensure you
  already have TLS for this exchange)

2. Connection acceptance or rejection from server.

3. * If rejected: end of communication
   * If accepted: client sends public signing key request along with client public
  key
4. Server responds with its own public key for signing and point-wise multiplies
  its own private key with the Client public key.

5. Client receives server public key and point-wise multiplies its own private key
  with the Server public key.

6. Both clients now have equal shared-secret signing keys.
