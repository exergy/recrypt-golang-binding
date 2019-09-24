extern crate libc;
extern crate recrypt;
use std::ffi::CStr;
// use recrypt::prelude::*;

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

#[no_mangle]
pub extern "C" fn get_proof_of_concept_string() -> &'static CStr {
    let val = CStr::from_bytes_with_nul(b"this string is from Rust\0").unwrap();
    return val;
}
