mod config;

use std::path::PathBuf;
use structopt::StructOpt;

#[derive(StructOpt, Debug)]
#[structopt(name = "tambourine")]
struct Opt {

    // The number of occurrences of the `v/verbose` flag
    /// Verbose mode (-v, -vv, -vvv, etc.)
    #[structopt(short, long, parse(from_occurrences))]
    verbose: u8,

    /// Watch directory
    #[structopt(short, long, parse(from_os_str))]
    watch: PathBuf,

}

fn main() {
    let opt = Opt::from_args();

    let service = true;
    while service == true {
        // Service here

        // Service here
    }

}