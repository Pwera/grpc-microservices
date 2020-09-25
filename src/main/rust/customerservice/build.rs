extern crate tonic_build;

fn main() {
    tonic_build::compile_protos("../../proto/customer/Customer.proto").unwrap();
}
