extern crate tonic_build;

use std::env;
use std::error::Error;
use env::current_dir;

fn main() -> Result<(), Box<dyn Error>> {
    let path = current_dir()?;
    println!(">> {:?}", path.to_str());
    tonic_build::configure()
        .build_server(true)
        .build_client(true)
        .out_dir(path.join("src").join("basic"))
        // .compile(&["../proto/Basic.proto"], &["../proto/"])?;
        .compile(&["Basic.proto"], &["."])?;
    Ok(())
}
