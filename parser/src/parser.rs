use std::str::Chars;

pub type LalrpopToken = Result<(usize, Token, usize), Error>;

pub struct Tokenizer<'input> {
    state: State,
    input: Chars<'input>,
    pos: usize,

    tokens: Vec<LalrpopToken>,

    buffer: String,

    // The tokenizer needs help recognizing known functions or constants
    is_const: fn(&str) -> bool,
    is_func: fn(&str) -> bool,
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
    UnaryNeg,
    UnaryNot,
    BinaryOp(OpCode),
    Factorial,
    LParen,
    RParen,
    Eof,
}

#[derive(Clone, Debug, PartialEq)]
pub enum OpCode {
    Add, Sub,
    Mul, Div,
    Pow, Mod,
    Or, And, Xor,
    Shl, Shr, Rol, Ror,
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

            // Later we'll plug in the actual functions
            is_const: |s| s == "pi" || s == "e",
            is_func: |s| s == "sin" || s == "cos" || s == "tan" || s == "log" || s == "ln",
        }
    }

    pub fn run(&mut self) -> Vec<LalrpopToken> {
        while self.state != State::Eof {
            self.step();
        }

        self.tokens.clone()
    }

    pub fn step(&mut self) {
        match self.state {
            State::BeforeExpr => {
                let next = self.peek();

                match next {
                    Some(c) if c.is_digit(10) || c == '.' || c == ',' => {
                        self.state = State::Number;
                    },
                    Some(c) if c.is_alphabetic() => {
                        self.state = State::BeforeExprText;
                    },
                    Some(c) if c == '-' => {
                        self.consume();
                        self.produce(Token::UnaryNeg)
                    },
                    Some(c) if c == '+' => {
                        self.consume(); // Ignore a prefix '+'
                    },
                    Some(c) if c == '~' => {
                        self.consume();
                        self.produce(Token::UnaryNot);
                    },
                    Some(c) if c == '(' => {
                        self.consume();
                        self.produce(Token::LParen)
                    },
                    Some(c) if c.is_whitespace() => {
                        self.consume(); // Ignore whitespace at the start
                    },
                    Some(c) => {
                        self.produce_error(Error::UnexpectedCharacter(c));
                        self.produce(Token::Eof);
                    },
                    None => self.produce(Token::Eof)
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
                    },
                    Some(c) if c == '%' => {
                        self.consume();
                        self.state = State::AmbiguousPercent;
                    },
                    Some(c) if c == '!' => {
                        self.produce(Token::Number(self.buffer.clone()));
                        self.consume();
                        self.produce(Token::Factorial);

                        self.state = State::BeforeOp;
                    },
                    Some(c) if c == '+' || c == '-' || c == '*' || c == '/' || c == '^' => {
                        self.produce(Token::Number(self.buffer.clone()));
                        self.state = State::BeforeOp;
                    },
                    Some(c) if c == '(' => {
                        // This is an implicit multiplication (e.g. 2(3+4)), so insert a multiplication operator
                        self.consume();
                        self.produce(Token::Number(self.buffer.clone()));
                        self.produce(Token::BinaryOp(OpCode::Mul));
                        self.produce(Token::LParen);

                        self.state = State::BeforeExpr;
                    },
                    Some(c) if c == ')' => {
                        self.produce(Token::Number(self.buffer.clone()));
                        self.consume();
                        self.produce(Token::RParen);

                        self.state = State::BeforeOp;
                    },
                    Some(c) if c.is_alphabetic() => {
                        self.produce(Token::Number(self.buffer.clone()));
                        self.state = State::AfterNumberText;
                    },
                    Some(c) => {
                        self.produce(Token::Number(self.buffer.clone()));
                        self.produce_error(Error::UnexpectedCharacter(c));

                        self.produce(Token::Eof);
                    },
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
                    Some(c) if c == '+' || c == '-' || c == '*' || c == '/' || c == '^' || c == '%' => {
                        self.state = State::BeforeOp;
                    },
                    Some(c) if c == '(' => {
                        // This is an implicit multiplication (e.g. 2 (3+4)), so insert a multiplication operator
                        // However, it's kind of weird to have a space before an opening parenthesis, so we'll also produce an error
                        self.consume();
                        self.produce(Token::BinaryOp(OpCode::Mul));
                        self.produce(Token::LParen);
                        self.produce_error(Error::UnexpectedCharacter(c));
                        self.state = State::BeforeExpr;
                    },
                    Some(c) if c == ')' => {
                        self.consume();
                        self.produce(Token::RParen);
                    },
                    Some(c) if c.is_alphabetic() => {
                        self.state = State::AfterNumberText;
                    },
                    Some(c) if c.is_whitespace() => {
                        self.consume();
                    },
                    Some(c) if c.is_digit(10) => {
                        // A number following a number is never valid, so this is a critical error
                        self.produce_error(Error::UnexpectedCharacter(c));
                        self.produce(Token::Eof);
                    },
                    Some(c) => {
                        self.produce_error(Error::UnexpectedCharacter(c));
                        self.produce(Token::Eof);
                    },
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
                    Some(c) if c == '+' || c == '-' || c == '*' || c == '/' || c == '^' => {
                        self.produce(Token::Percent(self.buffer.clone()));
                        self.state = State::BeforeOp;
                    },
                    Some(c) if c.is_alphabetic() => {
                        self.produce(Token::Percent(self.buffer.clone()));
                        self.state = State::AfterNumberText;
                    },
                    Some(c) if c == ')' => {
                        self.consume();
                        self.produce(Token::Percent(self.buffer.clone()));
                        self.produce(Token::RParen);
                        self.state = State::BeforeOp;
                    },
                    Some(c) if c == '(' => {
                        self.consume();

                        self.produce(Token::Number(self.buffer.clone()));
                        self.produce(Token::BinaryOp(OpCode::Mod));
                        self.produce(Token::LParen);

                        self.state = State::BeforeExpr;
                    },
                    Some(c) if c.is_whitespace() => {
                        self.consume();
                    },
                    Some(c) => {
                        self.produce(Token::Percent(self.buffer.clone()));
                        self.produce_error(Error::UnexpectedCharacter(c));
                        self.produce(Token::Eof);
                    },
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
                        self.produce(Token::BinaryOp(OpCode::Add));
                        self.state = State::BeforeExpr;
                    },
                    Some(c) if c == '-' => {
                        self.consume();
                        self.produce(Token::BinaryOp(OpCode::Sub));
                        self.state = State::BeforeExpr;
                    },
                    Some(c) if c == '*' => {
                        self.consume();

                        if self.peek() == Some('*') {
                            self.consume();
                            self.produce(Token::BinaryOp(OpCode::Pow));
                        } else {
                            self.produce(Token::BinaryOp(OpCode::Mul));
                        }
                        self.state = State::BeforeExpr;
                    },
                    Some(c) if c == '/' => {
                        self.consume();
                        self.produce(Token::BinaryOp(OpCode::Div));
                        self.state = State::BeforeExpr;
                    },
                    Some(c) if c == '^' => {
                        self.consume();
                        self.produce(Token::BinaryOp(OpCode::Pow));
                        self.state = State::BeforeExpr;
                    },
                    Some(c) if c == '%' => {
                        self.consume();
                        self.produce(Token::BinaryOp(OpCode::Mod));
                        self.state = State::BeforeExpr;
                    }
                    Some(c) if c == ')' => {
                        self.consume();
                        self.produce(Token::RParen);
                    },
                    Some(c) if c.is_whitespace() => {
                        self.consume();
                    },
                    Some(c) => {
                        self.produce_error(Error::UnexpectedCharacter(c));
                        self.produce(Token::Eof);
                    },
                    None => {
                        self.produce(Token::Eof);
                    }
                }
            }
            State::BeforeExprText => {
                // If an expression starts with a word, it can only be a function or a constant

                self.buffer = self.read_word();
                
                if (self.is_const)(&self.buffer) {
                    // A constant is a number in disguise
                    self.state = State::AfterNumber;
                } else if (self.is_func)(&self.buffer) {
                    self.produce(Token::FunctionCall(self.buffer.clone()));
                    self.state = State::BeforeExpr;
                } else {
                    self.produce_error(Error::UnexpectedCharacter(self.buffer.chars().nth(0).unwrap()));
                    self.produce(Token::Eof);
                }
            }
            State::AfterNumberText => {
                // If a number is followed by a word, it may be a function or a constant (implicit multiplication: 3pi)
                // or it could be an unit (3kg, 3m/s)

                self.buffer = self.read_word();
                
                if (self.is_const)(&self.buffer) {
                    // A constant is a number in disguise
                    self.produce(Token::BinaryOp(OpCode::Mul));
                    self.produce(Token::Number(self.buffer.clone()));
                    self.state = State::BeforeOp;
                } else if (self.is_func)(&self.buffer) {
                    self.produce(Token::BinaryOp(OpCode::Mul));
                    self.produce(Token::FunctionCall(self.buffer.clone()));
                    self.state = State::BeforeExpr;
                } else {
                    // println!("This might be an unit: {:?}", self.buffer);
                    self.state = State::BeforeOp;
                }
            },
            State::Eof => {
                // Do nothing
            }
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
            Token::Eof => 0,
            _ => 1
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

    fn read_word(&mut self) -> String {
        let mut text = String::new();

        while let Some(c) = self.peek() {
            if c.is_alphabetic() {
                text.push(self.consume().unwrap());
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

    pub fn peek(&self) -> Option<char> {
        self.input.clone().next()
    }

    // pub fn peek_nth(&mut self, n: usize) -> Option<char> {
    //     self.input.chars().nth(n)
    // }
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
        assert_eq!(
            run("123"),
            vec![Token::Number("123".to_string())]
        );

        assert_eq!(
            run("123.456"),
            vec![Token::Number("123.456".to_string())]
        );

        // This might be an error, but we don't assume a specific decimal separator at this stage, so well allow weird
        // numbers
        assert_eq!(
            run("123,45,.6"),
            vec![Token::Number("123,45,.6".to_string())]
        );
    }

    #[test]
    fn test_prefix_plus() {
        assert_eq!(
            run("+123"),
            vec![Token::Number("123".to_string())]
        );
    }

    #[test]
    fn test_prefix_minus() {
        assert_eq!(
            run("-123"),
            vec![Token::UnaryNeg, Token::Number("123".to_string())]
        );

        assert_eq!(
            run("1 - -2"),
            vec![
                Token::Number("1".to_string()),
                Token::BinaryOp(OpCode::Sub),
                Token::UnaryNeg,
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("3 - ~4"),
            vec![
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Sub),
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
        assert_eq!(
            run("123%"),
            vec![Token::Percent("123".to_string())]
        );

        assert_eq!(
            run("1,23.456%"),
            vec![Token::Percent("1,23.456".to_string())]
        );

        assert_eq!(
            run("123 %"),
            vec![
                Token::Number("123".to_string()),
                Token::BinaryOp(OpCode::Mod)
            ]
        );

        assert_eq!(
            run("123%)"),
            vec![
                Token::Percent("123".to_string()),
                Token::RParen
            ]
        );

        assert_eq!(
            run("123%pi"),
            vec![
                Token::Percent("123".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::Number("pi".to_string())
            ]
        );
    }

    #[test]
    fn test_factorial() {
        assert_eq!(
            run("5!"),
            vec![
                Token::Number("5".to_string()),
                Token::Factorial
            ]
        );

        assert_eq!(
            run("5!*3"),
            vec![
                Token::Number("5".to_string()),
                Token::Factorial,
                Token::BinaryOp(OpCode::Mul),
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
                Token::BinaryOp(OpCode::Add),
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("1 - 2"),
            vec![
                Token::Number("1".to_string()),
                Token::BinaryOp(OpCode::Sub),
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("1 * 2"),
            vec![
                Token::Number("1".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("1 / 2"),
            vec![
                Token::Number("1".to_string()),
                Token::BinaryOp(OpCode::Div),
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("1 ^ 2"),
            vec![
                Token::Number("1".to_string()),
                Token::BinaryOp(OpCode::Pow),
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("1 % 2"),
            vec![
                Token::Number("1".to_string()),
                Token::BinaryOp(OpCode::Mod),
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("1 + 2 + 3"),
            vec![
                Token::Number("1".to_string()),
                Token::BinaryOp(OpCode::Add),
                Token::Number("2".to_string()),
                Token::BinaryOp(OpCode::Add),
                Token::Number("3".to_string())
            ]
        );

        assert_eq!(
            run("1 + 2 * 3"),
            vec![
                Token::Number("1".to_string()),
                Token::BinaryOp(OpCode::Add),
                Token::Number("2".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::Number("3".to_string())
            ]
        );

        assert_eq!(
            run("1 ** 2"),
            vec![
                Token::Number("1".to_string()),
                Token::BinaryOp(OpCode::Pow),
                Token::Number("2".to_string())
            ]
        );

        assert_eq!(
            run("-3 + 4"),
            vec![
                Token::UnaryNeg,
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Add),
                Token::Number("4".to_string())
            ]
        )
    }

    #[test]
    fn test_brakets() {
        assert_eq!(
            run("(1 + 2) * 3"),
            vec![
                Token::LParen,
                Token::Number("1".to_string()),
                Token::BinaryOp(OpCode::Add),
                Token::Number("2".to_string()),
                Token::RParen,
                Token::BinaryOp(OpCode::Mul),
                Token::Number("3".to_string())
            ]
        );

        assert_eq!(
            run("1 + (2 * 3)"),
            vec![
                Token::Number("1".to_string()),
                Token::BinaryOp(OpCode::Add),
                Token::LParen,
                Token::Number("2".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::Number("3".to_string()),
                Token::RParen
            ]
        );

        assert_eq!(
            run("3%(3 + 4)"),
            vec![
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Mod),
                Token::LParen,
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Add),
                Token::Number("4".to_string()),
                Token::RParen
            ]
        );

        assert_eq!(
            run("3(5 + 4)"),
            vec![
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::LParen,
                Token::Number("5".to_string()),
                Token::BinaryOp(OpCode::Add),
                Token::Number("4".to_string()),
                Token::RParen
            ]
        );

        assert_eq!(
            run_fallible("3 (5 + 4) + 2").0,
            vec![
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::LParen,
                Token::Number("5".to_string()),
                Token::BinaryOp(OpCode::Add),
                Token::Number("4".to_string()),
                Token::RParen,
                Token::BinaryOp(OpCode::Add),
                Token::Number("2".to_string())
            ]
        );
        
        assert_eq!(
            run("(3 + 2   )"),
            vec![
                Token::LParen,
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Add),
                Token::Number("2".to_string()),
                Token::RParen
            ]
        );
    }

    #[test]
    fn test_words_at_beginning() {
        assert_eq!(
            run("pi"),
            vec![Token::Number("pi".to_string())]
        );

        assert_eq!(
            run("e"),
            vec![Token::Number("e".to_string())]
        );

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
                Token::BinaryOp(OpCode::Mul),
                Token::Number("pi".to_string())
            ]
        );

        assert_eq!(
            run("3e"),
            vec![
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::Number("e".to_string())
            ]
        );

        assert_eq!(
            run("3sin(3)"),
            vec![
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Mul),
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
                Token::BinaryOp(OpCode::Mul),
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
                Token::BinaryOp(OpCode::Mul),
                Token::FunctionCall("tan".to_string()),
                Token::Number("3".to_string()),
            ]
        );

        assert_eq!(
            run("3tan 3"),
            vec![
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Mul),
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
                Token::BinaryOp(OpCode::Mul),
                Token::Number("pi".to_string())
            ]
        );

        assert_eq!(
            run("3 + e"),
            vec![
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Add),
                Token::Number("e".to_string())
            ]
        );

        assert_eq!(
            run("3 - sin3"),
            vec![
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Sub),
                Token::FunctionCall("sin".to_string()),
                Token::Number("3".to_string()),
            ]
        );

        assert_eq!(
            run("pi sin 3"),
            vec![
                Token::Number("pi".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::FunctionCall("sin".to_string()),
                Token::Number("3".to_string()),
            ]
        );
    }

    #[test]
    fn test_hard_unitless() {
        assert_eq!(
            run("e ^ 2 - 3 * pi + 10%"),
            vec![
                Token::Number("e".to_string()),
                Token::BinaryOp(OpCode::Pow),
                Token::Number("2".to_string()),
                Token::BinaryOp(OpCode::Sub),
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::Number("pi".to_string()),
                Token::BinaryOp(OpCode::Add),
                Token::Percent("10".to_string())
            ]
        );

        assert_eq!(
            run("2(3 + sin pi - 4^2) / e"),
            vec![
                Token::Number("2".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::LParen,
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Add),
                Token::FunctionCall("sin".to_string()),
                Token::Number("pi".to_string()),
                Token::BinaryOp(OpCode::Sub),
                Token::Number("4".to_string()),
                Token::BinaryOp(OpCode::Pow),
                Token::Number("2".to_string()),
                Token::RParen,
                Token::BinaryOp(OpCode::Div),
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
                Token::BinaryOp(OpCode::Add),
                Token::Number("pi".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::Number("2".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::Number("e".to_string())
            ]
        );

        assert_eq!(
            run("2 * pi sin(30)"),
            vec![
                Token::Number("2".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::Number("pi".to_string()),
                Token::BinaryOp(OpCode::Mul),
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
                Token::BinaryOp(OpCode::Mul),
                Token::Number("e".to_string()),
                Token::BinaryOp(OpCode::Add),
                Token::Number("2".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::Number("pi".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::Number("e".to_string())
            ]
        );

        assert_eq!(
            run("pi% * e"),
            vec![
                Token::Percent("pi".to_string()),
                Token::BinaryOp(OpCode::Mul),
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
                Token::BinaryOp(OpCode::Add),
                Token::LParen,
                Token::Number("30".to_string()),
                Token::BinaryOp(OpCode::Mul),
                Token::Number("2".to_string()),
                Token::RParen,
                Token::RParen,
                Token::BinaryOp(OpCode::Div),
                Token::LParen,
                Token::Number("3".to_string()),
                Token::BinaryOp(OpCode::Pow),
                Token::Number("2".to_string()),
                Token::RParen,
                Token::RParen
            ]
        )
    }

    #[test]
    fn test_easy_failures() {
        assert_eq!(
            run_fallible("@"),
            (
                vec![],
                vec![Error::UnexpectedCharacter('@')]
            )
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
                vec![
                    Token::Number("1".to_string()),
                ],
                vec![Error::UnexpectedCharacter('@')]
            )
        );

        assert_eq!(
            run_fallible("3%@"),
            (
                vec![
                    Token::Percent("3".to_string()),
                ],
                vec![Error::UnexpectedCharacter('@')]
            )
        );

        assert_eq!(
            run_fallible("3 % @"),
            (
                vec![
                    Token::Number("3".to_string()),
                    Token::BinaryOp(OpCode::Mod),
                ],
                vec![Error::UnexpectedCharacter('@')]
            )
        );

        assert_eq!(
            run_fallible("hello"),
            (
                vec![],
                vec![Error::UnexpectedCharacter('h')]
            )
        );

        assert_eq!(
            run_fallible("(3 + 2) hello"),
            (
                vec![
                    Token::LParen,
                    Token::Number("3".to_string()),
                    Token::BinaryOp(OpCode::Add),
                    Token::Number("2".to_string()),
                    Token::RParen
                ],
                vec![Error::UnexpectedCharacter('h')]
            )
        );
    }
}