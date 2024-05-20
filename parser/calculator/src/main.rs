extern crate parser;

fn main() {
    // Get the first command line argument
    let expression = std::env::args().nth(1).expect("no expression provided");

    let res = parser::parser::parse(&expression);
    println!("{:?}", res.unwrap().eval());
}
