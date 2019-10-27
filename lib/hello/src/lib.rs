#[macro_use]
extern crate lazy_static;

extern crate libc;
extern crate recrypt;
use std::ffi::CStr;

use recrypt::prelude::*;
use recrypt::api::DefaultRng;

lazy_static! {
    static ref R:Recrypt<Sha256, Ed25519, RandomBytes<DefaultRng>> = Recrypt::new();
}

#[repr(C)]
pub struct CPublicKey {
    x: [u8; 32],
    y: [u8; 32],
}

#[repr(C)]
pub struct CKeyPair {
    public_key: CPublicKey,
    private_key: [u8; 32],
}

#[repr(C)]
pub struct CSigningKeyPair {
    pub public_key: [u8; 32],
    pub private_key: [u8; 64],
}

#[repr(C)]
pub struct CPlaintext {
    bytes: [u8; 384],
}

#[no_mangle]
pub extern "C" fn generate_keypair() -> CKeyPair {
    let (private_key, public_key) = R.generate_key_pair().unwrap();
    let private_key = *private_key.bytes();
    
    let (x, y) = public_key.bytes_x_y();
    let public_key = CPublicKey {
        x: *x,
        y: *y,
    };

    CKeyPair {
        public_key,
        private_key,
    }
}

#[no_mangle]
pub extern "C" fn generate_ed25519_keypair() -> CSigningKeyPair {
    let keypair = R.generate_ed25519_key_pair();
    let public_key = *keypair.public_key().bytes();
    let private_key = *keypair.bytes();

    CSigningKeyPair {
        public_key,
        private_key,
    }
}

/*
#[no_mangle]
pub extern "C" fn ed25519_sign(private_key:[u8; 32], message:&[u8]) -> [u8; 64] {
    
}
*/
/*
#[no_mangle]
pub extern "C" fn compute_ed25519_public_key(private_key:[u8; 64]) -> [u8; 32] {
    
}
*/


#[no_mangle]
pub extern "C" fn generate_plaintext() -> CPlaintext {
    CPlaintext {
        bytes: *R.gen_plaintext().bytes(),
    }
}

#[no_mangle]
pub extern "C" fn hello(name: *const libc::c_char) {
    let buf_name = unsafe {
        CStr::from_ptr(name).to_bytes()
    };
    let str_name = String::from_utf8(buf_name.to_vec()).unwrap();
    println!("A big hello from Rust to {}!", str_name);
}

#[no_mangle]
pub extern "C" fn get_proof_of_concept_string() -> &'static CStr {
    let val = CStr::from_bytes_with_nul(b"this string is from Rust\0").unwrap();
    return val;
}
