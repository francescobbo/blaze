#![no_main]

use parser::tokenizer::Tokenizer;

use libfuzzer_sys::fuzz_target;

fuzz_target!(|data: &[u8]| {
    if let Ok(s) = std::str::from_utf8(data) {
        let _ = Tokenizer::new(s).run();
    }
});
