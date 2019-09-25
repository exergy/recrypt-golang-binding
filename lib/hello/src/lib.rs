extern crate libc;
extern crate recrypt;
use std::ffi::CStr;
use recrypt::prelude::*;
use std::mem;

#[repr(C)]
pub struct CPublicKey {
    x: [u8; 32],
    y: [u8; 32]
}

#[repr(C)]
pub struct CPrivateKey {
    bytes: [u8; 32]
}

#[repr(C)]
pub struct CKeyPair {
    public_key: CPublicKey,
    private_key: CPrivateKey
}

#[no_mangle]
pub extern "C" fn hello(name: *const libc::c_char) {
    let buf_name = unsafe {
        CStr::from_ptr(name).to_bytes()
    };
    let str_name = String::from_utf8(buf_name.to_vec()).unwrap();
    println!("A big hello from Rust to {}!", str_name);
}

// #[no_mangle]
// pub extern "C" fn generate_key_pair() -> CKeyPair {
//     let recrypt = Recrypt::new();
//     let (priv_key, pub_key) = recrypt.generate_key_pair().unwrap();
//     let (pub_x, pub_y) = pub_key.bytes_x_y();
//
//     CKeyPair {
//         private_key: CPrivateKey { bytes: *priv_key.bytes() },
//         public_key: CPublicKey { x: *pub_x, y: *pub_y }
//     }
// }

#[no_mangle]
pub extern "C" fn generate_key_pair() -> *const char {
    let recrypt = Recrypt::new();
    let (priv_key, _) = recrypt.generate_key_pair().unwrap();

    &(priv_key.bytes()[0] as char)
}
