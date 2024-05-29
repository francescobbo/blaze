use lalrpop_util::{lalrpop_mod, ParseError};
lalrpop_mod!(pub calculator, "/parser/calculator.rs");
lalrpop_mod!(pub units, "/parser/units.rs");

pub mod ast;
pub mod tokenizer;
pub mod unit_tokenizer;

pub fn parse(
    input: &str,
) -> Result<Box<ast::Expr>, ParseError<usize, tokenizer::Token, tokenizer::Error>> {
    let tok = tokenizer::Tokenizer::new(input).run_without_eof();
    calculator::ExprParser::new().parse(tok)
}

#[cfg(test)]
mod tests {
    use crate::{ast::UnitExpr, value_unit::ValueUnit};

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
            ast::Expr::Number("123".to_string(), ValueUnit::empty())
        );

        assert_eq!(
            run("123.456"),
            ast::Expr::Number("123.456".to_string(), ValueUnit::empty())
        );

        assert_eq!(
            run("123,45,.6"),
            ast::Expr::Number("123,45,.6".to_string(), ValueUnit::empty())
        );

        assert_eq!(
            run("+123"),
            ast::Expr::Number("123".to_string(), ValueUnit::empty())
        );

        assert_eq!(
            run("-123"),
            ast::Expr::UnaryNegation(Box::new(ast::Expr::Number(
                "123".to_string(),
                ValueUnit::empty()
            )))
        );

        assert_eq!(
            run("~123"),
            ast::Expr::UnaryNot(Box::new(ast::Expr::Number(
                "123".to_string(),
                ValueUnit::empty()
            )))
        );

        assert_eq!(
            run("1 - -2"),
            ast::Expr::Subtract(
                Box::new(ast::Expr::Number("1".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::UnaryNegation(Box::new(ast::Expr::Number(
                    "2".to_string(),
                    ValueUnit::empty()
                ))))
            )
        );

        assert_eq!(
            run("3 - ~4"),
            ast::Expr::Subtract(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::UnaryNot(Box::new(ast::Expr::Number(
                    "4".to_string(),
                    ValueUnit::empty()
                ))))
            )
        );
    }

    #[test]
    fn test_percent() {
        assert_eq!(run("123%"), ast::Expr::Percent("123".to_string()));

        assert_eq!(run("1,23.456%"), ast::Expr::Percent("1,23.456".to_string()));

        assert_eq!(
            run("123 % 4"),
            ast::Expr::Modulo(
                Box::new(ast::Expr::Number("123".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("4".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(run("(123%)"), ast::Expr::Percent("123".to_string()));

        assert_eq!(
            run("123%pi"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Percent("123".to_string())),
                Box::new(ast::Expr::Number("pi".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("123% of (3 + 4)"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Percent("123".to_string())),
                Box::new(ast::Expr::Add(
                    Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                    Box::new(ast::Expr::Number("4".to_string(), ValueUnit::empty()))
                ))
            )
        );
    }

    #[test]
    fn test_factorial() {
        assert_eq!(
            run("5!"),
            ast::Expr::UnaryFactorial(Box::new(ast::Expr::Number("5".to_string(), ValueUnit::empty())))
        );

        assert_eq!(
            run("5!*3"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::UnaryFactorial(Box::new(ast::Expr::Number(
                    "5".to_string(),
                    ValueUnit::empty()
                )))),
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty()))
            )
        )
    }

    #[test]
    fn test_binary_op() {
        assert_eq!(
            run("1 + 2"),
            ast::Expr::Add(
                Box::new(ast::Expr::Number("1".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("1 - 2"),
            ast::Expr::Subtract(
                Box::new(ast::Expr::Number("1".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("1 * 2"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("1".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("1 / 2"),
            ast::Expr::Divide(
                Box::new(ast::Expr::Number("1".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("1 ^ 2"),
            ast::Expr::Power(
                Box::new(ast::Expr::Number("1".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("1 % 2"),
            ast::Expr::Modulo(
                Box::new(ast::Expr::Number("1".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("1 + 2 + 3"),
            ast::Expr::Add(
                Box::new(ast::Expr::Add(
                    Box::new(ast::Expr::Number("1".to_string(), ValueUnit::empty())),
                    Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
                )),
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("1 + 2 * 3"),
            ast::Expr::Add(
                Box::new(ast::Expr::Number("1".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Multiply(
                    Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty())),
                    Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty()))
                ))
            )
        );

        assert_eq!(
            run("1 ** 2"),
            ast::Expr::Power(
                Box::new(ast::Expr::Number("1".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("1 ^ 2 ^ 3"), // Test right associativity of power operator
            ast::Expr::Power(
                Box::new(ast::Expr::Number("1".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Power(
                    Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty())),
                    Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty()))
                ))
            )
        );

        assert_eq!(
            run("-3 + 4"),
            ast::Expr::Add(
                Box::new(ast::Expr::UnaryNegation(Box::new(ast::Expr::Number(
                    "3".to_string(),
                    ValueUnit::empty()
                )))),
                Box::new(ast::Expr::Number("4".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("3 & 4"),
            ast::Expr::And(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("4".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("3 | 4"),
            ast::Expr::Or(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("4".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("3 xor 4"),
            ast::Expr::Xor(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("4".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("3 << 4"),
            ast::Expr::ShiftLeft(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("4".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("3 >> 4"),
            ast::Expr::ShiftRight(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("4".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("3 rol 4"),
            ast::Expr::RotateLeft(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("4".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("3 ror 4"),
            ast::Expr::RotateRight(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("4".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("3 mod 5"),
            ast::Expr::Modulo(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("5".to_string(), ValueUnit::empty()))
            )
        );

        // Test precedence
        assert_eq!(
            run("1 & 2 | 3 xor 4 << 5 >> 6 rol 7 ror 8"),
            ast::Expr::Or(
                Box::new(ast::Expr::And(
                    Box::new(ast::Expr::Number("1".to_string(), ValueUnit::empty())),
                    Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
                )),
                Box::new(ast::Expr::Xor(
                    Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                    Box::new(ast::Expr::RotateRight(
                        Box::new(ast::Expr::RotateLeft(
                            Box::new(ast::Expr::ShiftRight(
                                Box::new(ast::Expr::ShiftLeft(
                                    Box::new(ast::Expr::Number("4".to_string(), ValueUnit::empty())),
                                    Box::new(ast::Expr::Number("5".to_string(), ValueUnit::empty()))
                                )),
                                Box::new(ast::Expr::Number("6".to_string(), ValueUnit::empty()))
                            )),
                            Box::new(ast::Expr::Number("7".to_string(), ValueUnit::empty()))
                        )),
                        Box::new(ast::Expr::Number("8".to_string(), ValueUnit::empty()))
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
                    Box::new(ast::Expr::Number("1".to_string(), ValueUnit::empty())),
                    Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
                )),
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("1 + (2 * 3)"),
            ast::Expr::Add(
                Box::new(ast::Expr::Number("1".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Multiply(
                    Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty())),
                    Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty()))
                ))
            )
        );

        assert_eq!(
            run("3%(3 + 4)"), // This is a modulo, not an implicit multiplication with a percentage
            ast::Expr::Modulo(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Add(
                    Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                    Box::new(ast::Expr::Number("4".to_string(), ValueUnit::empty()))
                ))
            )
        );

        assert_eq!(
            run("3(5 + 4)"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Add(
                    Box::new(ast::Expr::Number("5".to_string(), ValueUnit::empty())),
                    Box::new(ast::Expr::Number("4".to_string(), ValueUnit::empty()))
                ))
            )
        );

        assert_eq!(
            run("(3 + 2   )"),
            ast::Expr::Add(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("-(3 + 2)"),
            ast::Expr::UnaryNegation(Box::new(ast::Expr::Add(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
            )))
        );
    }

    #[test]
    fn test_constants_and_functions() {
        assert_eq!(
            run("pi"),
            ast::Expr::Number("pi".to_string(), ValueUnit::empty())
        );

        assert_eq!(run("e"), ast::Expr::Number("e".to_string(), ValueUnit::empty()));

        assert_eq!(
            run("sin(3)"),
            ast::Expr::FunctionCall(
                "sin".to_string(),
                vec![ast::Expr::Number("3".to_string(), ValueUnit::empty())]
            )
        );

        assert_eq!(
            run("cos(3)"),
            ast::Expr::FunctionCall(
                "cos".to_string(),
                vec![ast::Expr::Number("3".to_string(), ValueUnit::empty())]
            )
        );

        assert_eq!(
            run("tan3"),
            ast::Expr::FunctionCall(
                "tan".to_string(),
                vec![ast::Expr::Number("3".to_string(), ValueUnit::empty())]
            )
        );

        assert_eq!(
            run("tan 3"),
            ast::Expr::FunctionCall(
                "tan".to_string(),
                vec![ast::Expr::Number("3".to_string(), ValueUnit::empty())]
            )
        );

        assert_eq!(
            run("log(3)"),
            ast::Expr::FunctionCall(
                "log".to_string(),
                vec![ast::Expr::Number("3".to_string(), ValueUnit::empty())]
            )
        );

        assert_eq!(
            run("ln(3)"),
            ast::Expr::FunctionCall(
                "ln".to_string(),
                vec![ast::Expr::Number("3".to_string(), ValueUnit::empty())]
            )
        );

        assert_eq!(
            run("sin(cos(pi))"),
            ast::Expr::FunctionCall(
                "sin".to_string(),
                vec![ast::Expr::FunctionCall(
                    "cos".to_string(),
                    vec![ast::Expr::Number("pi".to_string(), ValueUnit::empty())]
                )]
            )
        );

        assert_eq!(
            run("sin cos pi"),
            ast::Expr::FunctionCall(
                "sin".to_string(),
                vec![ast::Expr::FunctionCall(
                    "cos".to_string(),
                    vec![ast::Expr::Number("pi".to_string(), ValueUnit::empty())]
                )]
            )
        );

        assert_eq!(
            run("3pi"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("pi".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("3e"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("e".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("3sin(3)"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::FunctionCall(
                    "sin".to_string(),
                    vec![ast::Expr::Number("3".to_string(), ValueUnit::empty())]
                ))
            )
        );

        assert_eq!(
            run("3cos(3)"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::FunctionCall(
                    "cos".to_string(),
                    vec![ast::Expr::Number("3".to_string(), ValueUnit::empty())]
                ))
            )
        );

        assert_eq!(
            run("3tan3"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::FunctionCall(
                    "tan".to_string(),
                    vec![ast::Expr::Number("3".to_string(), ValueUnit::empty())]
                ))
            )
        );

        assert_eq!(
            run("3tan 3"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::FunctionCall(
                    "tan".to_string(),
                    vec![ast::Expr::Number("3".to_string(), ValueUnit::empty())]
                ))
            )
        );

        assert_eq!(
            run("3 * pi"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("pi".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("3 + e"),
            ast::Expr::Add(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::Number("e".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("3 - sin3"),
            ast::Expr::Subtract(
                Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::FunctionCall(
                    "sin".to_string(),
                    vec![ast::Expr::Number("3".to_string(), ValueUnit::empty())]
                ))
            )
        );

        assert_eq!(
            run("pi sin 3"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Number("pi".to_string(), ValueUnit::empty())),
                Box::new(ast::Expr::FunctionCall(
                    "sin".to_string(),
                    vec![ast::Expr::Number("3".to_string(), ValueUnit::empty())]
                ))
            )
        );

        assert_eq!(
            run("sin 3^2"), // Test precedence of power operator over function call
            ast::Expr::FunctionCall(
                "sin".to_string(),
                vec![ast::Expr::Power(
                    Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                    Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
                )]
            )
        );

        assert_eq!(
            run("sin 3^2^3"),
            ast::Expr::FunctionCall(
                "sin".to_string(),
                vec![ast::Expr::Power(
                    Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                    Box::new(ast::Expr::Power(
                        Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty())),
                        Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty()))
                    ))
                )]
            )
        );

        assert_eq!(
            run("sin 3+2"), // Test precedence of function call over remaining operators
            ast::Expr::Add(
                Box::new(ast::Expr::FunctionCall(
                    "sin".to_string(),
                    vec![ast::Expr::Number("3".to_string(), ValueUnit::empty())]
                )),
                Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("log2 256"), // Test function name with number
            ast::Expr::FunctionCall(
                "log2".to_string(),
                vec![ast::Expr::Number("256".to_string(), ValueUnit::empty())]
            )
        );

        assert_eq!(
            run("log 2 + 256"),
            ast::Expr::Add(
                Box::new(ast::Expr::FunctionCall(
                    "log".to_string(),
                    vec![ast::Expr::Number("2".to_string(), ValueUnit::empty())]
                )),
                Box::new(ast::Expr::Number("256".to_string(), ValueUnit::empty()))
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
                        Box::new(ast::Expr::Number("e".to_string(), ValueUnit::empty())),
                        Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
                    )),
                    Box::new(ast::Expr::Multiply(
                        Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                        Box::new(ast::Expr::Number("pi".to_string(), ValueUnit::empty()))
                    ))
                )),
                Box::new(ast::Expr::Percent("10".to_string()))
            )
        );

        assert_eq!(
            run("2(3 + sin pi - 4^2) / e"),
            ast::Expr::Divide(
                Box::new(ast::Expr::Multiply(
                    Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty())),
                    Box::new(ast::Expr::Subtract(
                        Box::new(ast::Expr::Add(
                            Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                            Box::new(ast::Expr::FunctionCall(
                                "sin".to_string(),
                                vec![ast::Expr::Number("pi".to_string(), ValueUnit::empty())]
                            ))
                        )),
                        Box::new(ast::Expr::Power(
                            Box::new(ast::Expr::Number("4".to_string(), ValueUnit::empty())),
                            Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
                        ))
                    ))
                )),
                Box::new(ast::Expr::Number("e".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("sin(cos(tan 45)) + pi * 2e"),
            ast::Expr::Add(
                Box::new(ast::Expr::FunctionCall(
                    "sin".to_string(),
                    vec![ast::Expr::FunctionCall(
                        "cos".to_string(),
                        vec![ast::Expr::FunctionCall(
                            "tan".to_string(),
                            vec![ast::Expr::Number("45".to_string(), ValueUnit::empty())]
                        )]
                    )]
                )),
                Box::new(ast::Expr::Multiply(
                    Box::new(ast::Expr::Multiply(
                        Box::new(ast::Expr::Number("pi".to_string(), ValueUnit::empty())),
                        Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
                    )),
                    Box::new(ast::Expr::Number("e".to_string(), ValueUnit::empty()))
                ))
            )
        );

        assert_eq!(
            run("2 * pi sin(30)"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Multiply(
                    Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty())),
                    Box::new(ast::Expr::Number("pi".to_string(), ValueUnit::empty()))
                )),
                Box::new(ast::Expr::FunctionCall(
                    "sin".to_string(),
                    vec![ast::Expr::Number("30".to_string(), ValueUnit::empty())]
                ))
            )
        );

        assert_eq!(
            run("pi e + 2pi * 3e"),
            ast::Expr::Add(
                Box::new(ast::Expr::Multiply(
                    Box::new(ast::Expr::Number("pi".to_string(), ValueUnit::empty())),
                    Box::new(ast::Expr::Number("e".to_string(), ValueUnit::empty()))
                )),
                Box::new(ast::Expr::Multiply(
                    Box::new(ast::Expr::Multiply(
                        Box::new(ast::Expr::Multiply(
                            Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty())),
                            Box::new(ast::Expr::Number("pi".to_string(), ValueUnit::empty()))
                        )),
                        Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty()))
                    )),
                    Box::new(ast::Expr::Number("e".to_string(), ValueUnit::empty()))
                ))
            )
        );

        assert_eq!(
            run("pi% * e"),
            ast::Expr::Multiply(
                Box::new(ast::Expr::Percent("pi".to_string())),
                Box::new(ast::Expr::Number("e".to_string(), ValueUnit::empty()))
            )
        );

        assert_eq!(
            run("sin((45 + (30 * 2)) / (3 ^ 2))"),
            ast::Expr::FunctionCall(
                "sin".to_string(),
                vec![ast::Expr::Divide(
                    Box::new(ast::Expr::Add(
                        Box::new(ast::Expr::Number("45".to_string(), ValueUnit::empty())),
                        Box::new(ast::Expr::Multiply(
                            Box::new(ast::Expr::Number("30".to_string(), ValueUnit::empty())),
                            Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
                        ))
                    )),
                    Box::new(ast::Expr::Power(
                        Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
                        Box::new(ast::Expr::Number("2".to_string(), ValueUnit::empty()))
                    ))
                )]
            )
        );
    }

    // #[test]
    // fn test_unit_single_letter() {
    //     assert_eq!(
    //         run("3m"),
    //         ast::Expr::Unit(
    //             Box::new(ast::Expr::Number("3".to_string(), ValueUnit::empty())),
    //             "m".to_string()
    //         )
    //     );
    // }

    // #[test]
    // fn test_unit_multiple_letters() {
    //     assert_eq!(
    //         run("3kg"),
    //         vec![Token::Number("3".to_string()), Token::Text("kg".to_string())]
    //     );
    // }

    // #[test]
    // fn test_unit_with_space() {
    //     assert_eq!(
    //         run("3 kg"),
    //         vec![Token::Number("3".to_string()), Token::Text("kg".to_string())]
    //     );
    // }

    // #[test]
    // fn test_unit_with_implicit_pow() {
    //     assert_eq!(
    //         run("3m2"),
    //         vec![
    //             Token::Number("3".to_string()),
    //             Token::Text("m2".to_string()),
    //         ]
    //     );
    // }

    // #[test]
    // fn test_unit_with_explicit_pow() {
    //     assert_eq!(
    //         run("3m^2"),
    //         vec![
    //             Token::Number("3".to_string()),
    //             Token::Text("m".to_string()),
    //             Token::OpPow,
    //             Token::Number("2".to_string()),
    //         ]
    //     );
    // }

    // #[test]
    // fn test_unit_with_explicit_pow_and_space() {
    //     assert_eq!(
    //         run("3 m ^ 2"),
    //         vec![
    //             Token::Number("3".to_string()),
    //             Token::Text("m".to_string()),
    //             Token::OpPow,
    //             Token::Number("2".to_string()),
    //         ]
    //     );
    // }

    // #[test]
    // fn test_unit_with_explicit_pow_and_space_and_op() {
    //     assert_eq!(
    //         run("3 m ^ 2 + 4"),
    //         vec![
    //             Token::Number("3".to_string()),
    //             Token::Text("m".to_string()),
    //             Token::OpPow,
    //             Token::Number("2".to_string()),
    //             Token::OpAdd,
    //             Token::Number("4".to_string()),
    //         ]
    //     );
    // }

    // #[test]
    // fn test_multiplied_unit() {
    //     assert_eq!(
    //         run("3m*s"),
    //         vec![
    //             Token::Number("3".to_string()),
    //             Token::Text("m".to_string()),
    //             Token::OpMul,
    //             Token::Text("s".to_string()),
    //         ]
    //     );
    // }

    // #[test]
    // fn test_divided_unit() {
    //     assert_eq!(
    //         run("3m/s"),
    //         vec![
    //             Token::Number("3".to_string()),
    //             Token::Text("m".to_string()),
    //             Token::OpDiv,
    //             Token::Text("s".to_string()),
    //         ]
    //     );
    // }

    // #[test]
    // fn test_unit_with_division_implicit_pow() {
    //     assert_eq!(
    //         run("3m/s2"),
    //         vec![
    //             Token::Number("3".to_string()),
    //             Token::Text("m".to_string()),
    //             Token::OpDiv,
    //             Token::Text("s2".to_string()),
    //         ]
    //     );
    // }

    // #[test]
    // fn test_unit_with_division_explicit_pow() {
    //     assert_eq!(
    //         run("3m/s^2"),
    //         vec![
    //             Token::Number("3".to_string()),
    //             Token::Text("m".to_string()),
    //             Token::OpDiv,
    //             Token::Text("s".to_string()),
    //             Token::OpPow,
    //             Token::Number("2".to_string()),
    //         ]
    //     );
    // }

    // #[test]
    // fn test_unit_with_division_expression() {
    //     assert_eq!(
    //         run("3m/(kg*s)"),
    //         vec![
    //             Token::Number("3".to_string()),
    //             Token::Text("m".to_string()),
    //             Token::OpDiv,
    //             Token::LParen,
    //             Token::Text("kg".to_string()),
    //             Token::OpMul,
    //             Token::Text("s".to_string()),
    //             Token::RParen,
    //         ]
    //     );
    // }

    // #[test]
    // fn test_currency_and_divided_unit() {
    //     assert_eq!(
    //         run("$3/kg"),
    //         vec![
    //             Token::CurrencySymbol("$".to_string()),
    //             Token::Number("3".to_string()),
    //             Token::OpDiv,
    //             Token::Text("kg".to_string()),
    //         ]
    //     );
    // }
    // }

    #[test]
    fn test_unit_single() {
        assert_eq!(
            units::ExprParser::new().parse("m").unwrap(),
            UnitExpr::Unit("m".to_string(), 1)
        );
    }

    #[test]
    fn test_unit_pow() {
        assert_eq!(
            units::ExprParser::new().parse("m^2").unwrap(),
            UnitExpr::Unit("m".to_string(), 2)
        );
    }

    #[test]
    fn test_unit_neg_pow() {
        assert_eq!(
            units::ExprParser::new().parse("m^-2").unwrap(),
            UnitExpr::Unit("m".to_string(), -2)
        );
    }

    #[test]
    fn test_unit_mul() {
        assert_eq!(
            units::ExprParser::new().parse("m^2 * s^3").unwrap(),
            UnitExpr::Mul(
                Box::new(UnitExpr::Unit("m".to_string(), 2)),
                Box::new(UnitExpr::Unit("s".to_string(), 3))
            )
        );
    }

    #[test]
    fn test_unit_pow_mul() {
        assert_eq!(
            units::ExprParser::new().parse("(m * s)^2").unwrap(),
            UnitExpr::Mul(
                Box::new(UnitExpr::Unit("m".to_string(), 2)),
                Box::new(UnitExpr::Unit("s".to_string(), 2))
            )
        );
    }

    #[test]
    fn test_unit_mul_pow() {
        assert_eq!(
            units::ExprParser::new().parse("m * s^2").unwrap(),
            UnitExpr::Mul(
                Box::new(UnitExpr::Unit("m".to_string(), 1)),
                Box::new(UnitExpr::Unit("s".to_string(), 2))
            )
        );
    }

    #[test]
    fn test_unit_pow_mul_pow() {
        assert_eq!(
            units::ExprParser::new().parse("m^2 * s^2").unwrap(),
            UnitExpr::Mul(
                Box::new(UnitExpr::Unit("m".to_string(), 2)),
                Box::new(UnitExpr::Unit("s".to_string(), 2))
            )
        );
    }

    #[test]
    fn test_unit_div() {
        assert_eq!(
            units::ExprParser::new().parse("m/s^2").unwrap(),
            UnitExpr::Mul(
                Box::new(UnitExpr::Unit("m".to_string(), 1)),
                Box::new(UnitExpr::Unit("s".to_string(), -2))
            )
        );
    }

    #[test]
    fn test_unit_pow_div() {
        assert_eq!(
            units::ExprParser::new().parse("m^2/s^2").unwrap(),
            UnitExpr::Mul(
                Box::new(UnitExpr::Unit("m".to_string(), 2)),
                Box::new(UnitExpr::Unit("s".to_string(), -2))
            )
        );
    }

    #[test]
    fn test_unit_mul_div() {
        assert_eq!(
            units::ExprParser::new().parse("m * s^2/t").unwrap(),
            UnitExpr::Mul(
                Box::new(UnitExpr::Mul(
                    Box::new(UnitExpr::Unit("m".to_string(), 1)),
                    Box::new(UnitExpr::Unit("s".to_string(), 2))
                )),
                Box::new(UnitExpr::Unit("t".to_string(), -1))
            )
        );
    }

    #[test]
    fn test_unit_div_mul() {
        assert_eq!(
            units::ExprParser::new().parse("m/s^2 * t").unwrap(),
            UnitExpr::Mul(
                Box::new(UnitExpr::Mul(
                    Box::new(UnitExpr::Unit("m".to_string(), 1)),
                    Box::new(UnitExpr::Unit("s".to_string(), -2))
                )),
                Box::new(UnitExpr::Unit("t".to_string(), 1))
            )
        );
    }

    #[test]
    fn test_unit_div_div() {
        assert_eq!(
            units::ExprParser::new().parse("m/s^2/t").unwrap(),
            UnitExpr::Mul(
                Box::new(UnitExpr::Mul(
                    Box::new(UnitExpr::Unit("m".to_string(), 1)),
                    Box::new(UnitExpr::Unit("s".to_string(), -2))
                )),
                Box::new(UnitExpr::Unit("t".to_string(), -1))
            )
        );
    }

    #[test]
    fn test_unit_div_brackets() {
        assert_eq!(
            units::ExprParser::new().parse("m/(s^2)^3").unwrap(),
            UnitExpr::Mul(
                Box::new(UnitExpr::Unit("m".to_string(), 1)),
                Box::new(UnitExpr::Unit("s".to_string(), -6))
            )
        );
    }

    #[test]
    fn test_unit_g() {
        let res = units::ExprParser::new().parse("N*m^2/kg^2").unwrap();

        assert_eq!(
            res,
            UnitExpr::Mul(
                Box::new(UnitExpr::Mul(
                    Box::new(UnitExpr::Unit("N".to_string(), 1)),
                    Box::new(UnitExpr::Unit("m".to_string(), 2))
                )),
                Box::new(UnitExpr::Unit("kg".to_string(), -2))
            )
        );

        assert_eq!(
            res.flatten(),
            vec![
                ("N".to_string(), 1),
                ("m".to_string(), 2),
                ("kg".to_string(), -2)
            ]
        );
    }

    #[test]
    fn test_unary_division() {
        let res = units::ExprParser::new().parse("/s").unwrap();

        assert_eq!(res, UnitExpr::Unit("s".to_string(), -1));

        assert_eq!(res.flatten(), vec![("s".to_string(), -1),]);
    }

    #[test]
    fn test_weird_division() {
        // This may happen when a user types $3/s

        let res = units::ExprParser::new().parse("usd*/s").unwrap();

        assert_eq!(
            res,
            UnitExpr::Mul(
                Box::new(UnitExpr::Unit("usd".to_string(), 1)),
                Box::new(UnitExpr::Unit("s".to_string(), -1))
            )
        );

        assert_eq!(
            res.flatten(),
            vec![("usd".to_string(), 1), ("s".to_string(), -1),]
        );
    }
}
