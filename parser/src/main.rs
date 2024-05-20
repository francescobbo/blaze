// mod ast;
pub mod parser;

// fn main() {
//     let r1 = parser::Tokenizer::new("ʹ").run();

//     println!("{:?}", r1);
//     // let r2 = parser::Tokenizer::new().tokenize("33 % 4 to kg").unwrap();
//     // println!("{:?}, {:?}", r1, r2);
// }

#[macro_use]
extern crate afl;

fn main() {
    fuzz!(|data: &[u8]| {
        if let Ok(s) = std::str::from_utf8(data) {
            let _ = parser::Tokenizer::new(s).run();
        }
    });
}