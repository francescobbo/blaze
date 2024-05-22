#![no_main]

use parser::unit_tokenizer::UnitTokenizer;

use libfuzzer_sys::fuzz_target;

fuzz_target!(|data: &[u8]| {
    if let Ok(s) = std::str::from_utf8(data) {
        let it = s.chars();
        let _ = UnitTokenizer::new(it, vec!["sin", "cos", "tan", "pi", "e"]).run();
    }
});
