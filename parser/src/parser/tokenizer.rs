use std::str::Chars;

pub type LalrpopToken = Result<(usize, Token, usize), Error>;

pub struct Tokenizer<'input> {
    state: State,
    input: Chars<'input>,
    pos: usize,

    tokens: Vec<LalrpopToken>,

    buffer: String,
    unit_buffer: String,

    // The tokenizer needs help recognizing known functions or constants
    // TODO: allow for custom functions
    consts: Vec<&'static str>,
    funcs: Vec<&'static str>,
    is_currency_symbol: fn(&str) -> bool,
    of_operator: fn(&str) -> bool,
}

#[derive(Clone, Debug, PartialEq)]
pub enum Error {
    UnexpectedCharacter(char),
}

#[derive(Clone, Debug, PartialEq)]
pub enum Token {
    Number(String),
    Text(String),
    Percent(String),
    FunctionCall(String),
    CurrencySymbol(String),
    UnaryNeg,
    UnaryNot,
    Factorial,
    LParen,
    RParen,

    OpAdd,
    OpSub,
    OpMul,
    OpDiv,
    OpPow,
    OpMod,
    OpOr,
    OpAnd,
    OpXor,
    OpShl,
    OpShr,
    OpRol,
    OpRor,
    OpOf,

    Eof,
}

#[derive(Debug, PartialEq)]
enum State {
    BeforeExpr,
    Number,
    AfterNumber,
    AfterNumberSpace,
    AmbiguousPercent,
    BeforeOp,
    BeforeExprText,
    AfterNumberText,
    InUnit,
    InUnitAmbiguousOp,
    InUnitAfterAmbiguousOp,
    InUnitPower,
    Eof,
}

impl<'input> Tokenizer<'input> {
    pub fn new(input: &'input str) -> Self {
        Tokenizer {
            input: input.chars(),
            pos: 0,
            state: State::BeforeExpr,
            tokens: Vec::new(),

            buffer: String::new(),
            unit_buffer: String::new(),

            // Later we'll plug in the actual functions
            consts: vec!["pi", "e"],
            funcs: vec![
                "abs", "acos", "acosh", "asin", "asinh", "atan", "atanh", "cbrt", "ceil", "cos",
                "cosh", "exp", "floor", "ln", "log", "log2", "rand", "round", "sin", "sinh", "sqrt", "tan",
                "tanh", "trunc",
            ],
            is_currency_symbol: |s| vec!["$", "€", "£", "¥", "₹", "₽", "₿"].contains(&s),
            of_operator: |s| s == "of",
        }
    }

    pub fn run(&mut self) -> Vec<LalrpopToken> {
        while self.state != State::Eof {
            self.step();
        }

        self.tokens.clone()
    }

    pub fn run_without_eof(&mut self) -> Vec<LalrpopToken> {
        self.run()
            .into_iter()
            .take_while(|t| t.is_ok() && t.as_ref().unwrap().1 != Token::Eof)
            .collect()
    }

    pub fn step(&mut self) {
        match self.state {
            State::BeforeExpr => {
                let next = self.peek();

                match next {
                    Some(c) if c.is_digit(10) || c == '.' || c == ',' => {
                        self.state = State::Number;
                    }
                    Some(c) if c == '-' => {
                        self.consume();
                        self.produce(Token::UnaryNeg)
                    }
                    Some(c) if c == '+' => {
                        self.consume(); // Ignore a prefix '+'
                    }
                    Some(c) if c == '~' => {
                        self.consume();
                        self.produce(Token::UnaryNot);
                    }
                    Some(c) if c == '(' => {
                        self.consume();
                        self.produce(Token::LParen)
                    }
                    Some(c) if c == ')' => {
                        self.consume();
                        self.produce(Token::RParen);
                        self.state = State::BeforeOp;
                    }
                    Some(c) if c.is_whitespace() => {
                        self.consume(); // Ignore whitespace at the start
                    }
                    Some(_) => {
                        self.state = State::BeforeExprText;
                    }
                    None => self.produce(Token::Eof),
                }
            }
            State::Number => {
                self.buffer = self.read_numeric();
                self.state = State::AfterNumber;
            }
            State::AfterNumber => {
                // In a valid expression, after a number you can have:
                // - a unit (e.g. 2 kg)
                // - a function (e.g. 2 sin(3))
                // - a binary operator (e.g. 2 + 3),
                // - an opening parenthesis (e.g. 2(3 + 4))
                // - a closing parenthesis (e.g. 2 * (3 + 4))
                // - a percent sign (e.g. 33%) - this could be a percent or a modulo (but not with a space: 33 %)
                // - nothing (e.g. 2)

                let next = self.peek();

                match next {
                    Some(c) if c.is_whitespace() => {
                        self.produce(Token::Number(self.buffer.clone()));
                        self.state = State::AfterNumberSpace;
                        self.consume();
                    }
                    Some(c) if c == '%' => {
                        self.consume();
                        self.state = State::AmbiguousPercent;
                    }
                    Some(c) if c == '!' => {
                        self.produce(Token::Number(self.buffer.clone()));
                        self.consume();
                        self.produce(Token::Factorial);

                        self.state = State::BeforeOp;
                    }
                    Some(c)
                        if c == '+'
                            || c == '-'
                            || c == '*'
                            || c == '/'
                            || c == '^'
                            || c == '&'
                            || c == '|'
                            || c == '<'
                            || c == '>' =>
                    {
                        self.produce(Token::Number(self.buffer.clone()));
                        self.state = State::BeforeOp;
                    }
                    Some(c) if c == '(' => {
                        // This is an implicit multiplication (e.g. 2(3+4)), so insert a multiplication operator
                        self.consume();
                        self.produce(Token::Number(self.buffer.clone()));
                        self.produce(Token::OpMul);
                        self.produce(Token::LParen);

                        self.state = State::BeforeExpr;
                    }
                    Some(c) if c == ')' => {
                        self.produce(Token::Number(self.buffer.clone()));
                        self.consume();
                        self.produce(Token::RParen);

                        self.state = State::BeforeOp;
                    }
                    Some(_) => {
                        self.produce(Token::Number(self.buffer.clone()));
                        self.state = State::AfterNumberText;
                    }
                    None => {
                        self.produce(Token::Number(self.buffer.clone()));
                        self.produce(Token::Eof)
                    }
                }
            }
            State::AfterNumberSpace => {
                // The only difference here is that a '%' is not ambiguous anymore, it can only be a modulo

                let next = self.peek();

                match next {
                    Some(c)
                        if c == '+'
                            || c == '-'
                            || c == '*'
                            || c == '/'
                            || c == '^'
                            || c == '%'
                            || c == '&'
                            || c == '|'
                            || c == '<'
                            || c == '>' =>
                    {
                        self.state = State::BeforeOp;
                    }
                    Some(c) if c == '(' => {
                        // This is an implicit multiplication (e.g. 2 (3+4)), so insert a multiplication operator
                        // However, it's kind of weird to have a space before an opening parenthesis, so we'll also produce an error
                        self.consume();
                        self.produce(Token::OpMul);
                        self.produce(Token::LParen);
                        self.produce_error(Error::UnexpectedCharacter(c));
                        self.state = State::BeforeExpr;
                    }
                    Some(c) if c == ')' => {
                        self.consume();
                        self.produce(Token::RParen);
                    }
                    Some(c) if c.is_whitespace() => {
                        self.consume();
                    }
                    Some(c) if c.is_digit(10) => {
                        // A number following a number is never valid, so this is a critical error
                        self.produce_error(Error::UnexpectedCharacter(c));
                        self.produce(Token::Eof);
                    }
                    Some(_) => {
                        self.state = State::AfterNumberText;
                    }
                    None => {
                        self.produce(Token::Eof);
                    }
                }
            }
            State::AmbiguousPercent => {
                // We have a percent after a number (without space), could it be a modulo?
                // Ignore all the whitespace, then check the next character. It is a percent if:
                // - there's another operator (e.g. 33% * 4, 33% + 4, or even 33% % 5)
                // - there's text (e.g. 33% to kg, 33% of 4, 33% sin(2), 33%pi). Ok you want 33 modulo pi? Do 33 % pi or 33 mod pi
                // - there's a closing parenthesis (e.g. 1 + (2 * 33%))
                // - there's an EOF (e.g. 1+33%)
                // - what about an opening parenthesis? (e.g. 33%(2 + 3)) - I say this is a modulo. If you want a percent, multiply: 33% * (2 + 3)

                let next = self.peek();

                match next {
                    Some(c)
                        if c == '+'
                            || c == '-'
                            || c == '*'
                            || c == '/'
                            || c == '^'
                            || c == '%'
                            || c == '&'
                            || c == '|'
                            || c == '<'
                            || c == '>' =>
                    {
                        self.produce(Token::Percent(self.buffer.clone()));
                        self.state = State::BeforeOp;
                    }
                    Some(c) if c == ')' => {
                        self.consume();
                        self.produce(Token::Percent(self.buffer.clone()));
                        self.produce(Token::RParen);
                        self.state = State::BeforeOp;
                    }
                    Some(c) if c == '(' => {
                        self.consume();

                        self.produce(Token::Number(self.buffer.clone()));
                        self.produce(Token::OpMod);
                        self.produce(Token::LParen);

                        self.state = State::BeforeExpr;
                    }
                    Some(c) if c.is_whitespace() => {
                        self.consume();
                    }
                    Some(_) => {
                        self.produce(Token::Percent(self.buffer.clone()));
                        self.state = State::AfterNumberText;
                    }
                    None => {
                        self.produce(Token::Percent(self.buffer.clone()));
                        self.produce(Token::Eof);
                    }
                }
            }
            State::BeforeOp => {
                let next = self.peek();

                match next {
                    Some(c) if c == '+' => {
                        self.consume();
                        self.produce(Token::OpAdd);
                        self.state = State::BeforeExpr;
                    }
                    Some(c) if c == '-' => {
                        self.consume();
                        self.produce(Token::OpSub);
                        self.state = State::BeforeExpr;
                    }
                    Some(c) if c == '*' => {
                        self.consume();

                        if self.peek() == Some('*') {
                            self.consume();
                            self.produce(Token::OpPow);
                        } else {
                            self.produce(Token::OpMul);
                        }
                        self.state = State::BeforeExpr;
                    }
                    Some(c) if c == '/' => {
                        self.consume();
                        self.produce(Token::OpDiv);
                        self.state = State::BeforeExpr;
                    }
                    Some(c) if c == '^' => {
                        self.consume();
                        self.produce(Token::OpPow);
                        self.state = State::BeforeExpr;
                    }
                    Some(c) if c == '%' => {
                        self.consume();
                        self.produce(Token::OpMod);
                        self.state = State::BeforeExpr;
                    }
                    Some(c) if c == '&' => {
                        self.consume();
                        self.produce(Token::OpAnd);
                        self.state = State::BeforeExpr;
                    }
                    Some(c) if c == '|' => {
                        self.consume();
                        self.produce(Token::OpOr);
                        self.state = State::BeforeExpr;
                    }
                    Some(c) if c == '<' => {
                        self.consume();
                        if self.peek() == Some('<') {
                            self.consume();
                            self.produce(Token::OpShl);
                        } else {
                            self.produce_error(Error::UnexpectedCharacter(c));
                            self.produce(Token::Eof);
                        }
                        self.state = State::BeforeExpr;
                    }
                    Some(c) if c == '>' => {
                        self.consume();
                        if self.peek() == Some('>') {
                            self.consume();
                            self.produce(Token::OpShr);
                        } else {
                            self.produce_error(Error::UnexpectedCharacter(c));
                            self.produce(Token::Eof);
                        }
                        self.state = State::BeforeExpr;
                    }
                    Some(c) if c == ')' => {
                        self.consume();
                        self.produce(Token::RParen);
                    }
                    Some(c) if c.is_whitespace() => {
                        self.consume();
                    }
                    Some(_) => {
                        let word = self.read_word();

                        if word.len() > 0 {
                            self.skip(word.chars().count());
                            self.produce(Token::Text(word));
                        } else {
                            self.produce_error(Error::UnexpectedCharacter(self.peek().unwrap()));
                            self.produce(Token::Eof);
                        }
                    }
                    None => {
                        self.produce(Token::Eof);
                    }
                }
            }
            State::BeforeExprText => {
                // If an expression starts with a word, it can only be a function or a constant

                self.buffer = self.read_word();

                if self.consts.contains(&self.buffer.as_str()) {
                    // A constant is a number in disguise
                    self.state = State::AfterNumber;
                    self.skip(self.buffer.chars().count());
                } else {
                    // Some functions have numbers in their names. We'll greedily match the longest function name
                    // against the buffer which contains "[a-z][a-z0-9]*"
                    while self.buffer.len() > 0 {
                        if self.funcs.contains(&self.buffer.as_str()) {
                            self.skip(self.buffer.chars().count());
                            self.produce(Token::FunctionCall(self.buffer.clone()));
                            self.state = State::BeforeExpr;
                            break;
                        } else if (self.is_currency_symbol)(&self.buffer) {
                            self.skip(self.buffer.chars().count());
                            self.produce(Token::CurrencySymbol(self.buffer.clone()));
                            self.state = State::BeforeExpr;
                            break;
                        }

                        self.buffer.pop();
                    }

                    if self.buffer.len() == 0 {
                        // We didn't find a function, so it must be arbitrary text (a unit, a variable, etc.)
                        // re-read the buffer and produce a text token
                        let word = self.read_word();

                        if word.len() > 0 {
                            self.skip(word.chars().count());
                            self.produce(Token::Text(word));

                            // After a unit or a variable, we expect an operator
                            self.state = State::BeforeOp;
                        } else {
                            // This is text with non-alphanumeric characters
                            self.produce_error(Error::UnexpectedCharacter(self.peek().unwrap()));
                            self.produce(Token::Eof);
                        }
                    }
                }
            }
            State::AfterNumberText => {
                // If a number is followed by a word, it may be:
                // - a function or a constant (implicit multiplication: 3pi)
                // - a textual operator: mod, xor, rol, ror (e.g. 3 mod 4)
                // - a unit (e.g. 3 m/s)

                self.buffer = self.read_word();

                if self.consts.contains(&self.buffer.as_str()) {
                    // A constant is a number in disguise
                    self.skip(self.buffer.chars().count());
                    self.produce(Token::OpMul);
                    self.produce(Token::Number(self.buffer.clone()));
                    self.state = State::BeforeOp;
                } else {
                    while self.buffer.len() > 0 {
                        if self.buffer == "mod" {
                            self.skip(3);
                            self.produce(Token::OpMod);
                            self.state = State::BeforeExpr;
                            break;
                        } else if self.buffer == "xor" {
                            self.skip(3);
                            self.produce(Token::OpXor);
                            self.state = State::BeforeExpr;
                            break;
                        } else if self.buffer == "rol" {
                            self.skip(3);
                            self.produce(Token::OpRol);
                            self.state = State::BeforeExpr;
                            break;
                        } else if self.buffer == "ror" {
                            self.skip(3);
                            self.produce(Token::OpRor);
                            self.state = State::BeforeExpr;
                            break;
                        } else if self.funcs.contains(&self.buffer.as_str()) {
                            self.skip(self.buffer.chars().count());
                            self.produce(Token::OpMul);
                            self.produce(Token::FunctionCall(self.buffer.clone()));
                            self.state = State::BeforeExpr;
                            break;
                        } else if (self.of_operator)(&self.buffer) {
                            self.skip(self.buffer.chars().count());
                            self.produce(Token::OpOf);
                            self.state = State::BeforeExpr;
                            break;
                        } else if (self.is_currency_symbol)(&self.buffer) {
                            self.skip(self.buffer.chars().count());
                            self.produce(Token::CurrencySymbol(self.buffer.clone()));
                            self.state = State::AfterNumberSpace;
                            break;
                        }
                        
                        self.buffer.pop();
                    }

                    if self.buffer.len() == 0 {
                        // We didn't find a function, so it must be an unit or a variable
                        
                        let word = self.read_word();
                        if word.len() > 0 {
                            self.skip(word.chars().count());
                            self.produce(Token::Text(word));

                            // After a unit or a variable, we expect an operator
                            self.state = State::BeforeOp;
                        } else {
                            // This is text with non-alphanumeric characters
                            self.produce_error(Error::UnexpectedCharacter(self.peek().unwrap()));
                            self.produce(Token::Eof);
                        }
                    }
                }
            }
            State::InUnit => {
                match self.peek() {
                    Some(c) if c.is_alphanumeric() => {
                        let ch = self.consume().unwrap();
                        self.buffer.push(ch);
                    }
                    Some(c) if c.is_whitespace() => {
                        self.consume(); // Ignore whitespace
                    }
                    Some(c) if c == '*' || c == '/' => {
                        self.state = State::InUnitAmbiguousOp;
                    }
                    Some('^') => {
                        self.consume();
                        self.buffer.push('^');
                        self.state = State::InUnitPower;
                    }
                    Some(_) => {
                        self.produce(Token::Text(self.buffer.clone()));
                        self.state = State::BeforeOp;
                    }
                    None => {
                        self.produce(Token::Text(self.buffer.clone()));
                        self.produce(Token::Eof);
                    }
                }
            }
            State::InUnitAmbiguousOp => {
                let next = self.consume().unwrap();

                // match next {
                //     Some('*') => {
                //         self.unit_buffer.push(self.consume().unwrap());
                //     }
                //     Some('/') => {
                //         self.unit_buffer.push(self.consume().unwrap());
                //         self.produce(Token::Text(self.unit_buffer.clone()));
                //         self.state = State::BeforeOp;
                //     }
                //     Some(_) | None => unreachable!(),
                // }
            }
            State::InUnitPower => {
                let next = self.peek();

                match next {
                    Some(c) if c.is_digit(10) || c == '.' || c == ',' || c == '-' => {
                        let ch = self.consume().unwrap();
                        self.buffer.push(ch);
                    }
                    Some(c) if c.is_whitespace() => {
                        self.consume(); // Ignore whitespace
                    }
                    _ => {
                        self.produce(Token::Text(self.buffer.clone()));
                        self.state = State::InUnit;
                    }
                }
            }
            State::Eof => {
                // Do nothing
            }
            _ => panic!("Unexpected state: {:?}", self.state),
        }
    }

    fn produce(&mut self, token: Token) {
        if token == Token::Eof {
            self.state = State::Eof;
        }

        let end = self.pos;
        let size = match token {
            Token::Number(ref s) => s.len(),
            Token::Text(ref s) => s.len(),
            Token::Percent(ref s) => s.len() + 1,
            Token::FunctionCall(ref s) => s.len(),
            Token::CurrencySymbol(ref s) => s.len(),
            Token::OpShl | Token::OpShr => 2,
            Token::OpXor | Token::OpRol | Token::OpRor => 3,
            Token::Eof => 0,
            _ => 1,
        };

        let start = end - size;

        self.tokens.push(Ok((start, token, end)));
    }

    fn produce_error(&mut self, error: Error) {
        self.tokens.push(Err(error));
    }

    fn read_numeric(&mut self) -> String {
        let mut number = String::new();

        while let Some(c) = self.peek() {
            if c.is_digit(10) || c == '.' || c == ',' {
                number.push(self.consume().unwrap());
            } else {
                break;
            }
        }

        number
    }

    // Reads the following word by peeking at the next characters
    fn read_word(&mut self) -> String {
        let mut text = String::new();

        let mut i = 0;
        while let Some(c) = self.peek_nth(i) {
            if (!c.is_whitespace() && !c.is_ascii_punctuation()) || c == '$' {
                text.push(c);
                i += 1;
            } else {
                break;
            }
        }

        text
    }

    pub fn consume(&mut self) -> Option<char> {
        self.pos += 1;
        self.input.next()
    }

    pub fn skip(&mut self, n: usize) {
        self.pos += n;
        self.input.by_ref().take(n).for_each(drop);
    }

    pub fn peek(&self) -> Option<char> {
        self.input.clone().next()
    }

    pub fn peek_nth(&mut self, n: usize) -> Option<char> {
        self.input.clone().nth(n)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn run(input: &str) -> Vec<Token> {
        let mut tokenizer = Tokenizer::new(input);
        tokenizer.run();

        if tokenizer.tokens.iter().any(|t| t.is_err()) {
            panic!("Unexpected tokenizer failure: {:?}", tokenizer.tokens);
        }

        tokenizer
            .tokens
            .iter()
            .map(|t| t.as_ref().unwrap().1.clone())
            .take_while(|t| *t != Token::Eof)
            .collect()
    }

    fn run_fallible(input: &str) -> (Vec<Token>, Vec<Error>) {
        let mut tokenizer = Tokenizer::new(input);
        tokenizer.run();

        let tokens: Vec<Token> = tokenizer
            .tokens
            .iter()
            .filter_map(|t| t.as_ref().ok())
            .map(|t| t.1.clone())
            .take_while(|t| *t != Token::Eof)
            .collect();
        let errors = tokenizer
            .tokens
            .iter()
            .filter_map(|t| t.as_ref().err())
            .cloned()
            .collect();

        (tokens, errors)
    }

    #[test]
    fn test_number() {
        assert_eq!(run("123"), vec![Token::Number("123".to_string())]);
    }

    #[test]
    fn test_number_with_decimal() {
        assert_eq!(run("123.456"), vec![Token::Number("123.456".to_string())]);
    }

    #[test]
    fn test_number_with_comma() {
        // This might be an error, but we don't assume a specific decimal separator at this stage, so well allow weird
        // numbers
        assert_eq!(
            run("123,45,.6"),
            vec![Token::Number("123,45,.6".to_string())]
        );
    }

    #[test]
    fn test_prefix_plus() {
        assert_eq!(run("+123"), vec![Token::Number("123".to_string())]);
    }

    #[test]
    fn test_prefix_minus() {
        assert_eq!(
            run("-123"),
            vec![Token::UnaryNeg, Token::Number("123".to_string())]
        );
    }

    #[test]
    fn test_prefix_minus_after_op() {
        assert_eq!(
            run("1 - -2"),
            vec![
                Token::Number("1".to_string()),
                Token::OpSub,
                Token::UnaryNeg,
                Token::Number("2".to_string())
            ]
        );
    }

    #[test]
    fn test_prefix_not() {
        assert_eq!(
            run("3 - ~4"),
            vec![
                Token::Number("3".to_string()),
                Token::OpSub,
                Token::UnaryNot,
                Token::Number("4".to_string())
            ]
        )
    }

    #[test]
    fn test_invalid_number() {
        assert_eq!(
            run_fallible("123 456"),
            (
                vec![Token::Number("123".to_string())],
                vec![Error::UnexpectedCharacter('4')]
            )
        );
    }

    #[test]
    fn test_percent() {
        assert_eq!(run("123%"), vec![Token::Percent("123".to_string())]);
    }

    #[test]
    fn test_weird_percent() {
        assert_eq!(
            run("1,23.456%"),
            vec![Token::Percent("1,23.456".to_string())]
        );
    }

    #[test]
    fn test_percent_with_space_is_mod() {
        assert_eq!(
            run("123 %"),
            vec![Token::Number("123".to_string()), Token::OpMod]
        );
    }

    #[test]
    fn test_percent_at_end_of_expression() {
        assert_eq!(
            run("123%)"),
            vec![Token::Percent("123".to_string()), Token::RParen]
        );
    }

    #[test]
    fn test_percent_implicit_mul() {
        assert_eq!(
            run("123%pi"),
            vec![
                Token::Percent("123".to_string()),
                Token::OpMul,
                Token::Number("pi".to_string())
            ]
        );
    }

    #[test]
    fn test_percent_of() {
        assert_eq!(
            run("123% of 100"),
            vec![
                Token::Percent("123".to_string()),
                Token::OpOf,
                Token::Number("100".to_string())
            ]
        );
    }

    #[test]
    fn test_factorial() {
        assert_eq!(
            run("5!"),
            vec![Token::Number("5".to_string()), Token::Factorial]
        );
    }

    #[test]
    fn test_factorial_with_op() {
        assert_eq!(
            run("5!*3"),
            vec![
                Token::Number("5".to_string()),
                Token::Factorial,
                Token::OpMul,
                Token::Number("3".to_string())
            ]
        )
    }

    #[test]
    fn test_binary_op() {
        assert_eq!(
            run("1 + 2"),
            vec![
                Token::Number("1".to_string()),
                Token::OpAdd,
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("1 - 2"),
            vec![
                Token::Number("1".to_string()),
                Token::OpSub,
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("1 * 2"),
            vec![
                Token::Number("1".to_string()),
                Token::OpMul,
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("1 / 2"),
            vec![
                Token::Number("1".to_string()),
                Token::OpDiv,
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("1 ^ 2"),
            vec![
                Token::Number("1".to_string()),
                Token::OpPow,
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("1 % 2"),
            vec![
                Token::Number("1".to_string()),
                Token::OpMod,
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("1 + 2 + 3"),
            vec![
                Token::Number("1".to_string()),
                Token::OpAdd,
                Token::Number("2".to_string()),
                Token::OpAdd,
                Token::Number("3".to_string())
            ]
        );

        assert_eq!(
            run("1 + 2 * 3"),
            vec![
                Token::Number("1".to_string()),
                Token::OpAdd,
                Token::Number("2".to_string()),
                Token::OpMul,
                Token::Number("3".to_string())
            ]
        );

        assert_eq!(
            run("1 ** 2"),
            vec![
                Token::Number("1".to_string()),
                Token::OpPow,
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("-3 + 4"),
            vec![
                Token::UnaryNeg,
                Token::Number("3".to_string()),
                Token::OpAdd,
                Token::Number("4".to_string())
            ]
        );

        assert_eq!(
            run("3&4"),
            vec![
                Token::Number("3".to_string()),
                Token::OpAnd,
                Token::Number("4".to_string())
            ]
        );

        assert_eq!(
            run("3|4"),
            vec![
                Token::Number("3".to_string()),
                Token::OpOr,
                Token::Number("4".to_string())
            ]
        );

        assert_eq!(
            run("3<<4"),
            vec![
                Token::Number("3".to_string()),
                Token::OpShl,
                Token::Number("4".to_string())
            ]
        );

        assert_eq!(
            run("3>>4"),
            vec![
                Token::Number("3".to_string()),
                Token::OpShr,
                Token::Number("4".to_string())
            ]
        );

        assert_eq!(
            run("3xor4"),
            vec![
                Token::Number("3".to_string()),
                Token::OpXor,
                Token::Number("4".to_string())
            ]
        );

        assert_eq!(
            run("3rol4"),
            vec![
                Token::Number("3".to_string()),
                Token::OpRol,
                Token::Number("4".to_string())
            ]
        );

        assert_eq!(
            run("3ror4"),
            vec![
                Token::Number("3".to_string()),
                Token::OpRor,
                Token::Number("4".to_string())
            ]
        );

        assert_eq!(
            run("3mod4"),
            vec![
                Token::Number("3".to_string()),
                Token::OpMod,
                Token::Number("4".to_string())
            ]
        );

        assert_eq!(
            run("3 xor 4"),
            vec![
                Token::Number("3".to_string()),
                Token::OpXor,
                Token::Number("4".to_string())
            ]
        );

        assert_eq!(
            run("3 mod 4"),
            vec![
                Token::Number("3".to_string()),
                Token::OpMod,
                Token::Number("4".to_string())
            ]
        );
    }

    #[test]
    fn test_brakets() {
        assert_eq!(
            run("(1 + 2) * 3"),
            vec![
                Token::LParen,
                Token::Number("1".to_string()),
                Token::OpAdd,
                Token::Number("2".to_string()),
                Token::RParen,
                Token::OpMul,
                Token::Number("3".to_string())
            ]
        );

        assert_eq!(
            run("1 + (2 * 3)"),
            vec![
                Token::Number("1".to_string()),
                Token::OpAdd,
                Token::LParen,
                Token::Number("2".to_string()),
                Token::OpMul,
                Token::Number("3".to_string()),
                Token::RParen
            ]
        );

        assert_eq!(
            run("3%(3 + 4)"),
            vec![
                Token::Number("3".to_string()),
                Token::OpMod,
                Token::LParen,
                Token::Number("3".to_string()),
                Token::OpAdd,
                Token::Number("4".to_string()),
                Token::RParen
            ]
        );

        assert_eq!(
            run("3(5 + 4)"),
            vec![
                Token::Number("3".to_string()),
                Token::OpMul,
                Token::LParen,
                Token::Number("5".to_string()),
                Token::OpAdd,
                Token::Number("4".to_string()),
                Token::RParen
            ]
        );

        assert_eq!(
            run_fallible("3 (5 + 4) + 2").0,
            vec![
                Token::Number("3".to_string()),
                Token::OpMul,
                Token::LParen,
                Token::Number("5".to_string()),
                Token::OpAdd,
                Token::Number("4".to_string()),
                Token::RParen,
                Token::OpAdd,
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("(3 + 2   )"),
            vec![
                Token::LParen,
                Token::Number("3".to_string()),
                Token::OpAdd,
                Token::Number("2".to_string()),
                Token::RParen
            ]
        );
    }

    #[test]
    fn test_words_at_beginning() {
        assert_eq!(run("pi"), vec![Token::Number("pi".to_string())]);

        assert_eq!(run("e"), vec![Token::Number("e".to_string())]);

        assert_eq!(
            run("sin(3)"),
            vec![
                Token::FunctionCall("sin".to_string()),
                Token::LParen,
                Token::Number("3".to_string()),
                Token::RParen
            ]
        );

        assert_eq!(
            run("cos(3)"),
            vec![
                Token::FunctionCall("cos".to_string()),
                Token::LParen,
                Token::Number("3".to_string()),
                Token::RParen
            ]
        );

        assert_eq!(
            run("tan3"),
            vec![
                Token::FunctionCall("tan".to_string()),
                Token::Number("3".to_string()),
            ]
        );

        assert_eq!(
            run("tan 3"),
            vec![
                Token::FunctionCall("tan".to_string()),
                Token::Number("3".to_string()),
            ]
        );

        assert_eq!(
            run("log(3)"),
            vec![
                Token::FunctionCall("log".to_string()),
                Token::LParen,
                Token::Number("3".to_string()),
                Token::RParen
            ]
        );

        assert_eq!(
            run("ln(3)"),
            vec![
                Token::FunctionCall("ln".to_string()),
                Token::LParen,
                Token::Number("3".to_string()),
                Token::RParen
            ]
        );

        assert_eq!(
            run("sin(cos(pi))"),
            vec![
                Token::FunctionCall("sin".to_string()),
                Token::LParen,
                Token::FunctionCall("cos".to_string()),
                Token::LParen,
                Token::Number("pi".to_string()),
                Token::RParen,
                Token::RParen
            ]
        );

        assert_eq!(
            run("sin cos pi"),
            vec![
                Token::FunctionCall("sin".to_string()),
                Token::FunctionCall("cos".to_string()),
                Token::Number("pi".to_string())
            ]
        );
    }

    #[test]
    fn test_words_after_number() {
        assert_eq!(
            run("3pi"),
            vec![
                Token::Number("3".to_string()),
                Token::OpMul,
                Token::Number("pi".to_string())
            ]
        );

        assert_eq!(
            run("3e"),
            vec![
                Token::Number("3".to_string()),
                Token::OpMul,
                Token::Number("e".to_string())
            ]
        );

        assert_eq!(
            run("3sin(3)"),
            vec![
                Token::Number("3".to_string()),
                Token::OpMul,
                Token::FunctionCall("sin".to_string()),
                Token::LParen,
                Token::Number("3".to_string()),
                Token::RParen
            ]
        );

        assert_eq!(
            run("3cos(3)"),
            vec![
                Token::Number("3".to_string()),
                Token::OpMul,
                Token::FunctionCall("cos".to_string()),
                Token::LParen,
                Token::Number("3".to_string()),
                Token::RParen
            ]
        );

        assert_eq!(
            run("3tan3"),
            vec![
                Token::Number("3".to_string()),
                Token::OpMul,
                Token::FunctionCall("tan".to_string()),
                Token::Number("3".to_string()),
            ]
        );

        assert_eq!(
            run("3tan 3"),
            vec![
                Token::Number("3".to_string()),
                Token::OpMul,
                Token::FunctionCall("tan".to_string()),
                Token::Number("3".to_string()),
            ]
        );
    }

    #[test]
    fn test_words_after_symbol() {
        assert_eq!(
            run("3 * pi"),
            vec![
                Token::Number("3".to_string()),
                Token::OpMul,
                Token::Number("pi".to_string())
            ]
        );

        assert_eq!(
            run("3 + e"),
            vec![
                Token::Number("3".to_string()),
                Token::OpAdd,
                Token::Number("e".to_string())
            ]
        );

        assert_eq!(
            run("3 - sin3"),
            vec![
                Token::Number("3".to_string()),
                Token::OpSub,
                Token::FunctionCall("sin".to_string()),
                Token::Number("3".to_string()),
            ]
        );

        assert_eq!(
            run("pi sin 3"),
            vec![
                Token::Number("pi".to_string()),
                Token::OpMul,
                Token::FunctionCall("sin".to_string()),
                Token::Number("3".to_string()),
            ]
        );
    }

    #[test]
    fn test_function_with_number_in_name() {
        assert_eq!(
            run("log210 + log2(10) + log2pi + pi*log2(10)"),
            vec![
                Token::FunctionCall("log2".to_string()),
                Token::Number("10".to_string()),
                Token::OpAdd,
                Token::FunctionCall("log2".to_string()),
                Token::LParen,
                Token::Number("10".to_string()),
                Token::RParen,
                Token::OpAdd,
                Token::FunctionCall("log2".to_string()),
                Token::Number("pi".to_string()),
                Token::OpAdd,
                Token::Number("pi".to_string()),
                Token::OpMul,
                Token::FunctionCall("log2".to_string()),
                Token::LParen,
                Token::Number("10".to_string()),
                Token::RParen
            ]
        );
    }

    #[test]
    fn test_hard_unitless() {
        assert_eq!(
            run("e ^ 2 - 3 * pi + 10%"),
            vec![
                Token::Number("e".to_string()),
                Token::OpPow,
                Token::Number("2".to_string()),
                Token::OpSub,
                Token::Number("3".to_string()),
                Token::OpMul,
                Token::Number("pi".to_string()),
                Token::OpAdd,
                Token::Percent("10".to_string())
            ]
        );

        assert_eq!(
            run("2(3 + sin pi - 4^2) / e"),
            vec![
                Token::Number("2".to_string()),
                Token::OpMul,
                Token::LParen,
                Token::Number("3".to_string()),
                Token::OpAdd,
                Token::FunctionCall("sin".to_string()),
                Token::Number("pi".to_string()),
                Token::OpSub,
                Token::Number("4".to_string()),
                Token::OpPow,
                Token::Number("2".to_string()),
                Token::RParen,
                Token::OpDiv,
                Token::Number("e".to_string())
            ]
        );

        assert_eq!(
            run("sin(cos(tan 45)) + pi * 2e"),
            vec![
                Token::FunctionCall("sin".to_string()),
                Token::LParen,
                Token::FunctionCall("cos".to_string()),
                Token::LParen,
                Token::FunctionCall("tan".to_string()),
                Token::Number("45".to_string()),
                Token::RParen,
                Token::RParen,
                Token::OpAdd,
                Token::Number("pi".to_string()),
                Token::OpMul,
                Token::Number("2".to_string()),
                Token::OpMul,
                Token::Number("e".to_string())
            ]
        );

        assert_eq!(
            run("2 * pi sin(30)"),
            vec![
                Token::Number("2".to_string()),
                Token::OpMul,
                Token::Number("pi".to_string()),
                Token::OpMul,
                Token::FunctionCall("sin".to_string()),
                Token::LParen,
                Token::Number("30".to_string()),
                Token::RParen
            ]
        );

        assert_eq!(
            run("pi e + 2pi * 3e"),
            vec![
                Token::Number("pi".to_string()),
                Token::OpMul,
                Token::Number("e".to_string()),
                Token::OpAdd,
                Token::Number("2".to_string()),
                Token::OpMul,
                Token::Number("pi".to_string()),
                Token::OpMul,
                Token::Number("3".to_string()),
                Token::OpMul,
                Token::Number("e".to_string())
            ]
        );

        assert_eq!(
            run("pi% * e"),
            vec![
                Token::Percent("pi".to_string()),
                Token::OpMul,
                Token::Number("e".to_string())
            ]
        );

        assert_eq!(
            run("sin((45 + (30 * 2)) / (3 ^ 2))"),
            vec![
                Token::FunctionCall("sin".to_string()),
                Token::LParen,
                Token::LParen,
                Token::Number("45".to_string()),
                Token::OpAdd,
                Token::LParen,
                Token::Number("30".to_string()),
                Token::OpMul,
                Token::Number("2".to_string()),
                Token::RParen,
                Token::RParen,
                Token::OpDiv,
                Token::LParen,
                Token::Number("3".to_string()),
                Token::OpPow,
                Token::Number("2".to_string()),
                Token::RParen,
                Token::RParen
            ]
        );

        assert_eq!(
            run("rand() + 2"),
            vec![
                Token::FunctionCall("rand".to_string()),
                Token::LParen,
                Token::RParen,
                Token::OpAdd,
                Token::Number("2".to_string())
            ]
        );
    }

    #[test]
    fn test_currency_symbol() {
        assert_eq!(
            run("$10"),
            vec![Token::CurrencySymbol("$".to_string()), Token::Number("10".to_string())]
        );

        assert_eq!(
            run("10$ + 20€"),
            vec![
                Token::Number("10".to_string()),
                Token::CurrencySymbol("$".to_string()),
                Token::OpAdd,
                Token::Number("20".to_string()),
                Token::CurrencySymbol("€".to_string())
            ]
        );

        assert_eq!(
            run("10 $ + € 20 + 30£"),
            vec![
                Token::Number("10".to_string()),
                Token::CurrencySymbol("$".to_string()),
                Token::OpAdd,
                Token::CurrencySymbol("€".to_string()),
                Token::Number("20".to_string()),
                Token::OpAdd,
                Token::Number("30".to_string()),
                Token::CurrencySymbol("£".to_string())
            ]
        );
    }

    #[test]
    fn test_unit_single_letter() {
        assert_eq!(
            run("3m"),
            vec![Token::Number("3".to_string()), Token::Text("m".to_string())]
        );
    }

    #[test]
    fn test_unit_multiple_letters() {
        assert_eq!(
            run("3kg"),
            vec![Token::Number("3".to_string()), Token::Text("kg".to_string())]
        );
    }

    #[test]
    fn test_unit_with_space() {
        assert_eq!(
            run("3 kg"),
            vec![Token::Number("3".to_string()), Token::Text("kg".to_string())]
        );
    }

    #[test]
    fn test_unit_with_implicit_pow() {
        assert_eq!(
            run("3m2"),
            vec![
                Token::Number("3".to_string()),
                Token::Text("m2".to_string()),
            ]
        );
    }

    #[test]
    fn test_unit_with_explicit_pow() {
        assert_eq!(
            run("3m^2"),
            vec![
                Token::Number("3".to_string()),
                Token::Text("m".to_string()),
                Token::OpPow,
                Token::Number("2".to_string()),
            ]
        );
    }

    #[test]
    fn test_unit_with_explicit_pow_and_space() {
        assert_eq!(
            run("3 m ^ 2"),
            vec![
                Token::Number("3".to_string()),
                Token::Text("m".to_string()),
                Token::OpPow,
                Token::Number("2".to_string()),
            ]
        );
    }

    #[test]
    fn test_unit_with_explicit_pow_and_space_and_op() {
        assert_eq!(
            run("3 m ^ 2 + 4"),
            vec![
                Token::Number("3".to_string()),
                Token::Text("m".to_string()),
                Token::OpPow,
                Token::Number("2".to_string()),
                Token::OpAdd,
                Token::Number("4".to_string()),
            ]
        );
    }

    #[test]
    fn test_multiplied_unit() {
        assert_eq!(
            run("3m*s"),
            vec![
                Token::Number("3".to_string()),
                Token::Text("m".to_string()),
                Token::OpMul,
                Token::Text("s".to_string()),
            ]
        );
    }

    #[test]
    fn test_divided_unit() {
        assert_eq!(
            run("3m/s"),
            vec![
                Token::Number("3".to_string()),
                Token::Text("m".to_string()),
                Token::OpDiv,
                Token::Text("s".to_string()),
            ]
        );
    }

    #[test]
    fn test_unit_with_division_implicit_pow() {
        assert_eq!(
            run("3m/s2"),
            vec![
                Token::Number("3".to_string()),
                Token::Text("m".to_string()),
                Token::OpDiv,
                Token::Text("s2".to_string()),
            ]
        );
    }

    #[test]
    fn test_unit_with_division_explicit_pow() {
        assert_eq!(
            run("3m/s^2"),
            vec![
                Token::Number("3".to_string()),
                Token::Text("m".to_string()),
                Token::OpDiv,
                Token::Text("s".to_string()),
                Token::OpPow,
                Token::Number("2".to_string()),
            ]
        );
    }

    #[test]
    fn test_unit_with_division_expression() {
        assert_eq!(
            run("3m/(kg*s)"),
            vec![
                Token::Number("3".to_string()),
                Token::Text("m".to_string()),
                Token::OpDiv,
                Token::LParen,
                Token::Text("kg".to_string()),
                Token::OpMul,
                Token::Text("s".to_string()),
                Token::RParen,
            ]
        );
    }

    #[test]
    fn test_currency_and_divided_unit() {
        assert_eq!(
            run("$3/kg"),
            vec![
                Token::CurrencySymbol("$".to_string()),
                Token::Number("3".to_string()),
                Token::OpDiv,
                Token::Text("kg".to_string()),
            ]
        );
    }

    #[test]
    fn test_easy_failures() {
        assert_eq!(
            run_fallible("@"),
            (vec![], vec![Error::UnexpectedCharacter('@')])
        );

        assert_eq!(
            run_fallible("1@"),
            (
                vec![Token::Number("1".to_string())],
                vec![Error::UnexpectedCharacter('@')]
            )
        );

        assert_eq!(
            run_fallible("1 @"),
            (
                vec![Token::Number("1".to_string()),],
                vec![Error::UnexpectedCharacter('@')]
            )
        );

        assert_eq!(
            run_fallible("3%@"),
            (
                vec![Token::Percent("3".to_string()),],
                vec![Error::UnexpectedCharacter('@')]
            )
        );

        assert_eq!(
            run_fallible("3 % @"),
            (
                vec![Token::Number("3".to_string()), Token::OpMod,],
                vec![Error::UnexpectedCharacter('@')]
            )
        );

        assert_eq!(
            run("hello"),
            vec![
                Token::Text("hello".to_string()),
            ]
        );

        assert_eq!(
            run("(3 + 2) hello"),
            vec![
                Token::LParen,
                Token::Number("3".to_string()),
                Token::OpAdd,
                Token::Number("2".to_string()),
                Token::RParen,
                Token::Text("hello".to_string()),
            ],
        );
    }
}
