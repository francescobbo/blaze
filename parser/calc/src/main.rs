extern crate parser;

fn main() {
    // Get the first command line argument
    let expression = std::env::args().nth(1).expect("no expression provided");

    let resf = parser::parse(&expression).unwrap().eval();
    println!("{} {}", resf.0, resf.1);
}
