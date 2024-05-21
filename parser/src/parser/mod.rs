use lalrpop_util::{lalrpop_mod, ParseError};
lalrpop_mod!(pub calculator, "/parser/calculator.rs");

pub mod tokenizer;

use crate::ast;

pub fn parse(input: &str) -> Result<Box<ast::Expr>, ParseError<usize, tokenizer::Token, tokenizer::Error>> {
    let tok = tokenizer::Tokenizer::new(input).run_without_eof();
    calculator::ExprParser::new().parse(tok)
}

#[cfg(test)]
mod tests {
    use super::*;

    fn run(input: &str) -> ast::Expr {
        let res = parse(input);
        match res {
            Ok(expr) => *expr,
            Err(e) => panic!("Error: {:?}", e),
        }
    }

    #[test]
    fn test_number() {
        assert_eq!(
            run("123"),
            ast::Expr::Number("123".to_string())
        );

        assert_eq!(
            run("123.456"),
            ast::Expr::Number("123.456".to_string())
        );

        assert_eq!(
            run("123,45,.6"),
            ast::Expr::Number("123,45,.6".to_string())
        );

        assert_eq!(
            run("+123"),
            ast::Expr::Number("123".to_string())
        );

        assert_eq!(
            run("-123"),
            ast::Expr::UnaryNegation(
                Box::new(ast::Expr::Number("123".to_string()))
            )
        );

        assert_eq!(
            run("~123"),
            ast::Expr::UnaryNot(
                Box::new(ast::Expr::Number("123".to_string()))
            )
        );

        assert_eq!(
            run("1 - -2"),
            ast::Expr::Subtract(
                Box::new(ast::Expr::Number("1".to_string())),
                Box::new(ast::Expr::UnaryNegation(
                    Box::new(ast::Expr::Number("2".to_string()))
                ))
            )
        );

        assert_eq!(
            run("3 - ~4"),
            ast::Expr::Subtract(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::UnaryNot(
                    Box::new(ast::Expr::Number("4".to_string()))
                ))
            )
        );
    }

    #[test]
    fn test_percent() {
        assert_eq!(
            run("123%"),
            ast::Expr::Percent("123".to_string())
        );

        assert_eq!(
            run("1,23.456%"),
            ast::Expr::Percent("1,23.456".to_string())
        );

        assert_eq!(
            run("123 % 4"),
            ast::Expr::Modulo(
                Box::new(ast::Expr::Number("123".to_string())),
                Box::new(ast::Expr::Number("4".to_string()))
            )
        );

        assert_eq!(
            run("(123%)"),
            ast::Expr::Percent("123".to_string())
        );

        assert_eq!(
            run("123%pi"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Percent("123".to_string())),
                Box::new(ast::Expr::Number("pi".to_string()))
            )
        );    
    }

    #[test]
    fn test_factorial() {
        assert_eq!(
            run("5!"),
            ast::Expr::UnaryFactorial(Box::new(ast::Expr::Number("5".to_string())))
        );

        assert_eq!(
            run("5!*3"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::UnaryFactorial(Box::new(ast::Expr::Number("5".to_string())))),
                Box::new(ast::Expr::Number("3".to_string()))
            )
        )
    }

    #[test]
    fn test_binary_op() {
        assert_eq!(
            run("1 + 2"),
            ast::Expr::Add(
                Box::new(ast::Expr::Number("1".to_string())),
                Box::new(ast::Expr::Number("2".to_string()))
            )
        );

        assert_eq!(
            run("1 - 2"),
            ast::Expr::Subtract(
                Box::new(ast::Expr::Number("1".to_string())),
                Box::new(ast::Expr::Number("2".to_string()))
            )
        );

        assert_eq!(
            run("1 * 2"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("1".to_string())),
                Box::new(ast::Expr::Number("2".to_string()))
            )
        );

        assert_eq!(
            run("1 / 2"),
            ast::Expr::Divide(
                Box::new(ast::Expr::Number("1".to_string())),
                Box::new(ast::Expr::Number("2".to_string()))
            )
        );

        assert_eq!(
            run("1 ^ 2"),
            ast::Expr::Power(
                Box::new(ast::Expr::Number("1".to_string())),
                Box::new(ast::Expr::Number("2".to_string()))
            )
        );

        assert_eq!(
            run("1 % 2"),
            ast::Expr::Modulo(
                Box::new(ast::Expr::Number("1".to_string())),
                Box::new(ast::Expr::Number("2".to_string()))
            )
        );

        assert_eq!(
            run("1 + 2 + 3"),
            ast::Expr::Add(
                Box::new(ast::Expr::Add(
                    Box::new(ast::Expr::Number("1".to_string())),
                    Box::new(ast::Expr::Number("2".to_string()))
                )),
                Box::new(ast::Expr::Number("3".to_string()))
            )
        );

        assert_eq!(
            run("1 + 2 * 3"),
            ast::Expr::Add(
                Box::new(ast::Expr::Number("1".to_string())),
                Box::new(ast::Expr::Multiply(
                    Box::new(ast::Expr::Number("2".to_string())),
                    Box::new(ast::Expr::Number("3".to_string()))
                ))
            )
        );

        assert_eq!(
            run("1 ** 2"),
            ast::Expr::Power(
                Box::new(ast::Expr::Number("1".to_string())),
                Box::new(ast::Expr::Number("2".to_string()))
            )
        );

        assert_eq!(
            run("1 ^ 2 ^ 3"), // Test right associativity of power operator
            ast::Expr::Power(
                Box::new(ast::Expr::Number("1".to_string())),
                Box::new(ast::Expr::Power(
                    Box::new(ast::Expr::Number("2".to_string())),
                    Box::new(ast::Expr::Number("3".to_string()))
                ))
            )
        );

        assert_eq!(
            run("-3 + 4"),
            ast::Expr::Add(
                Box::new(ast::Expr::UnaryNegation(
                    Box::new(ast::Expr::Number("3".to_string()))
                )),
                Box::new(ast::Expr::Number("4".to_string()))
            )
        );

        assert_eq!(
            run("3 & 4"),
            ast::Expr::And(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::Number("4".to_string()))
            )
        );

        assert_eq!(
            run("3 | 4"),
            ast::Expr::Or(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::Number("4".to_string()))
            )
        );

        assert_eq!(
            run("3 xor 4"),
            ast::Expr::Xor(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::Number("4".to_string()))
            )
        );

        assert_eq!(
            run("3 << 4"),
            ast::Expr::ShiftLeft(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::Number("4".to_string()))
            )
        );

        assert_eq!(
            run("3 >> 4"),
            ast::Expr::ShiftRight(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::Number("4".to_string()))
            )
        );

        assert_eq!(
            run("3 rol 4"),
            ast::Expr::RotateLeft(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::Number("4".to_string()))
            )
        );

        assert_eq!(
            run("3 ror 4"),
            ast::Expr::RotateRight(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::Number("4".to_string()))
            )
        );

        assert_eq!(
            run("3 mod 5"),
            ast::Expr::Modulo(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::Number("5".to_string()))
            )
        );

        // Test precedence
        assert_eq!(
            run("1 & 2 | 3 xor 4 << 5 >> 6 rol 7 ror 8"),
            ast::Expr::Or(
                Box::new(ast::Expr::And(
                    Box::new(ast::Expr::Number("1".to_string())),
                    Box::new(ast::Expr::Number("2".to_string()))
                )),
                Box::new(ast::Expr::Xor(
                    Box::new(ast::Expr::Number("3".to_string())),
                    Box::new(ast::Expr::RotateRight(
                        Box::new(ast::Expr::RotateLeft(
                            Box::new(ast::Expr::ShiftRight(
                                Box::new(ast::Expr::ShiftLeft(
                                    Box::new(ast::Expr::Number("4".to_string())),
                                    Box::new(ast::Expr::Number("5".to_string()))
                                )),
                                Box::new(ast::Expr::Number("6".to_string()))
                            )),
                            Box::new(ast::Expr::Number("7".to_string()))
                        )),
                        Box::new(ast::Expr::Number("8".to_string()))
                    ))
                ))
            )
        );
    }

    #[test]
    fn test_brakets() {
        assert_eq!(
            run("(1 + 2) * 3"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Add(
                    Box::new(ast::Expr::Number("1".to_string())),
                    Box::new(ast::Expr::Number("2".to_string()))
                )),
                Box::new(ast::Expr::Number("3".to_string()))
            )
        );

        assert_eq!(
            run("1 + (2 * 3)"),
            ast::Expr::Add(
                Box::new(ast::Expr::Number("1".to_string())),
                Box::new(ast::Expr::Multiply(
                    Box::new(ast::Expr::Number("2".to_string())),
                    Box::new(ast::Expr::Number("3".to_string()))
                ))
            )
        );

        assert_eq!(
            run("3%(3 + 4)"), // This is a modulo, not an implicit multiplication with a percentage
            ast::Expr::Modulo(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::Add(
                    Box::new(ast::Expr::Number("3".to_string())),
                    Box::new(ast::Expr::Number("4".to_string()))
                ))
            )
        );

        assert_eq!(
            run("3(5 + 4)"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::Add(
                    Box::new(ast::Expr::Number("5".to_string())),
                    Box::new(ast::Expr::Number("4".to_string()))
                ))
            )
        );

        assert_eq!(
            run("(3 + 2   )"),
            ast::Expr::Add(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::Number("2".to_string()))
            )
        );

        assert_eq!(
            run("-(3 + 2)"),
            ast::Expr::UnaryNegation(
                Box::new(ast::Expr::Add(
                    Box::new(ast::Expr::Number("3".to_string())),
                    Box::new(ast::Expr::Number("2".to_string()))
                ))
            )
        );
    }

    #[test]
    fn test_constants_and_functions() {
        assert_eq!(
            run("pi"),
            ast::Expr::Number("pi".to_string())
        );

        assert_eq!(
            run("e"),
            ast::Expr::Number("e".to_string())
        );

        assert_eq!(
            run("sin(3)"),
            ast::Expr::FunctionCall("sin".to_string(), vec![ast::Expr::Number("3".to_string())])
        );

        assert_eq!(
            run("cos(3)"),
            ast::Expr::FunctionCall("cos".to_string(), vec![ast::Expr::Number("3".to_string())])
        );

        assert_eq!(
            run("tan3"),
            ast::Expr::FunctionCall("tan".to_string(), vec![ast::Expr::Number("3".to_string())])
        );

        assert_eq!(
            run("tan 3"),
            ast::Expr::FunctionCall("tan".to_string(), vec![ast::Expr::Number("3".to_string())])
        );

        assert_eq!(
            run("log(3)"),
            ast::Expr::FunctionCall("log".to_string(), vec![ast::Expr::Number("3".to_string())])
        );

        assert_eq!(
            run("ln(3)"),
            ast::Expr::FunctionCall("ln".to_string(), vec![ast::Expr::Number("3".to_string())])
        );

        assert_eq!(
            run("sin(cos(pi))"),
            ast::Expr::FunctionCall("sin".to_string(), vec![
                ast::Expr::FunctionCall("cos".to_string(), vec![
                    ast::Expr::Number("pi".to_string())
                ])
            ])
        );

        assert_eq!(
            run("sin cos pi"),
            ast::Expr::FunctionCall("sin".to_string(), vec![
                ast::Expr::FunctionCall("cos".to_string(), vec![
                    ast::Expr::Number("pi".to_string())
                ])
            ])
        );

        assert_eq!(
            run("3pi"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::Number("pi".to_string()))
            )
        );

        assert_eq!(
            run("3e"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::Number("e".to_string()))
            )
        );

        assert_eq!(
            run("3sin(3)"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::FunctionCall("sin".to_string(), vec![
                    ast::Expr::Number("3".to_string())
                ])
            ))
        );

        assert_eq!(
            run("3cos(3)"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::FunctionCall("cos".to_string(), vec![
                    ast::Expr::Number("3".to_string())
                ])
            ))
        );

        assert_eq!(
            run("3tan3"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::FunctionCall("tan".to_string(), vec![
                    ast::Expr::Number("3".to_string())
                ])
            ))
        );

        assert_eq!(
            run("3tan 3"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::FunctionCall("tan".to_string(), vec![
                    ast::Expr::Number("3".to_string())
                ])
            ))
        );

        assert_eq!(
            run("3 * pi"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::Number("pi".to_string()))
            )
        );

        assert_eq!(
            run("3 + e"),
            ast::Expr::Add(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::Number("e".to_string()))
            )
        );

        assert_eq!(
            run("3 - sin3"),
            ast::Expr::Subtract(
                Box::new(ast::Expr::Number("3".to_string())),
                Box::new(ast::Expr::FunctionCall("sin".to_string(), vec![
                    ast::Expr::Number("3".to_string())
                ])
            ))
        );

        assert_eq!(
            run("pi sin 3"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("pi".to_string())),
                Box::new(ast::Expr::FunctionCall("sin".to_string(), vec![
                    ast::Expr::Number("3".to_string())
                ])
            ))
        );

        assert_eq!(
            run("sin 3^2"), // Test precedence of power operator over function call
            ast::Expr::FunctionCall("sin".to_string(), vec![
                ast::Expr::Power(
                    Box::new(ast::Expr::Number("3".to_string())),
                    Box::new(ast::Expr::Number("2".to_string()))
                )
            ])
        );

        assert_eq!(
            run("sin 3^2^3"),
            ast::Expr::FunctionCall("sin".to_string(), vec![
                ast::Expr::Power(
                    Box::new(ast::Expr::Number("3".to_string())),
                    Box::new(ast::Expr::Power(
                        Box::new(ast::Expr::Number("2".to_string())),
                        Box::new(ast::Expr::Number("3".to_string()))
                    )
                ))
            ])
        );

        assert_eq!(
            run("sin 3+2"), // Test precedence of function call over remaining operators
            ast::Expr::Add(
                Box::new(ast::Expr::FunctionCall("sin".to_string(), vec![
                    ast::Expr::Number("3".to_string())
                ])),
                Box::new(ast::Expr::Number("2".to_string()))
            )
        );

        assert_eq!(
            run("log2 256"), // Test function name with number
            ast::Expr::FunctionCall("log2".to_string(), vec![
                ast::Expr::Number("256".to_string())
            ])
        );

        assert_eq!(
            run("log 2 + 256"),
            ast::Expr::Add(
                Box::new(ast::Expr::FunctionCall("log".to_string(), vec![
                    ast::Expr::Number("2".to_string())
                ])),
                Box::new(ast::Expr::Number("256".to_string()))
            )
        );
    }

    #[test]
    fn test_mixed() {
        assert_eq!(
            run("e ^ 2 - 3 * pi + 10%"),
            ast::Expr::Add(
                Box::new(ast::Expr::Subtract(
                    Box::new(ast::Expr::Power(
                        Box::new(ast::Expr::Number("e".to_string())),
                        Box::new(ast::Expr::Number("2".to_string()))
                    )),
                    Box::new(ast::Expr::Multiply(
                        Box::new(ast::Expr::Number("3".to_string())),
                        Box::new(ast::Expr::Number("pi".to_string()))
                    ))
                )),
                Box::new(ast::Expr::Percent("10".to_string()))
            )
        );

        assert_eq!(
            run("2(3 + sin pi - 4^2) / e"),
            ast::Expr::Divide(
                Box::new(
                    ast::Expr::Multiply(
                    Box::new(ast::Expr::Number("2".to_string())),
                    Box::new(ast::Expr::Subtract(
                        Box::new(ast::Expr::Add(
                            Box::new(ast::Expr::Number("3".to_string())),
                            Box::new(ast::Expr::FunctionCall("sin".to_string(), vec![
                                ast::Expr::Number("pi".to_string())
                            ])
                        ))),
                        Box::new(ast::Expr::Power(
                            Box::new(ast::Expr::Number("4".to_string())),
                            Box::new(ast::Expr::Number("2".to_string()))
                        ))
                    ))
                )),
                Box::new(ast::Expr::Number("e".to_string()))
            )
        );

        assert_eq!(
            run("sin(cos(tan 45)) + pi * 2e"),
            ast::Expr::Add(
                Box::new(ast::Expr::FunctionCall("sin".to_string(), vec![
                    ast::Expr::FunctionCall("cos".to_string(), vec![
                        ast::Expr::FunctionCall("tan".to_string(), vec![
                            ast::Expr::Number("45".to_string())
                        ])
                    ])
                ])),
                Box::new(ast::Expr::Multiply(
                    Box::new(ast::Expr::Multiply(
                        Box::new(ast::Expr::Number("pi".to_string())),
                        Box::new(ast::Expr::Number("2".to_string()))
                    )),
                    Box::new(ast::Expr::Number("e".to_string()))
                ))
            )
        );

        assert_eq!(
            run("2 * pi sin(30)"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Multiply(
                    Box::new(ast::Expr::Number("2".to_string())),
                    Box::new(ast::Expr::Number("pi".to_string()))
                )),
                Box::new(ast::Expr::FunctionCall("sin".to_string(), vec![
                    ast::Expr::Number("30".to_string())
                ]))
            )
        );

        assert_eq!(
            run("pi e + 2pi * 3e"),
            ast::Expr::Add(
                Box::new(ast::Expr::Multiply(
                    Box::new(ast::Expr::Number("pi".to_string())),
                    Box::new(ast::Expr::Number("e".to_string()))
                )),
                Box::new(ast::Expr::Multiply(
                    Box::new(ast::Expr::Multiply(
                        Box::new(ast::Expr::Multiply(
                            Box::new(ast::Expr::Number("2".to_string())),
                            Box::new(ast::Expr::Number("pi".to_string()))
                        )),
                        Box::new(ast::Expr::Number("3".to_string()))
                    )),
                    Box::new(ast::Expr::Number("e".to_string())
                ))
            ))
        );

        assert_eq!(
            run("pi% * e"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Percent("pi".to_string())),
                Box::new(ast::Expr::Number("e".to_string()))
            )
        );

        assert_eq!(
            run("sin((45 + (30 * 2)) / (3 ^ 2))"),
            ast::Expr::FunctionCall("sin".to_string(), vec![
                ast::Expr::Divide(
                    Box::new(ast::Expr::Add(
                        Box::new(ast::Expr::Number("45".to_string())),
                        Box::new(ast::Expr::Multiply(
                            Box::new(ast::Expr::Number("30".to_string())),
                            Box::new(ast::Expr::Number("2".to_string()))
                        ))
                    )),
                    Box::new(ast::Expr::Power(
                        Box::new(ast::Expr::Number("3".to_string())),
                        Box::new(ast::Expr::Number("2".to_string()))
                    ))
                )
            ])
        );
    }
}