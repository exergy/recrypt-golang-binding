#include <stdlib.h>
#include <stdint.h>

void hello(char *name);
char* get_proof_of_concept_string();


struct CPublicKey {
  char x[32];
  char y[32];
};
struct CKeyPair {
   struct CPublicKey public_key;
   char private_key[32];
};

struct CSigningKeyPair {
  char public_key[32];
  char private_key[64];
};

struct CPlaintext {
  char bytes[384];
};

struct CKeyPair generate_keypair();
struct CSigningKeyPair generate_ed25519_keypair();
struct CPlaintext generate_plaintext();
