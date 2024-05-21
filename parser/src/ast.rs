#[derive(Debug)]
pub enum Expr {
    Number(String),
    Percent(String),
    Add(Box<Expr>, Box<Expr>),
    Subtract(Box<Expr>, Box<Expr>),
    Multiply(Box<Expr>, Box<Expr>),
    Divide(Box<Expr>, Box<Expr>),
    Modulo(Box<Expr>, Box<Expr>),
    Power(Box<Expr>, Box<Expr>),
    And(Box<Expr>, Box<Expr>),
    Or(Box<Expr>, Box<Expr>),
    Xor(Box<Expr>, Box<Expr>),
    ShiftLeft(Box<Expr>, Box<Expr>),
    ShiftRight(Box<Expr>, Box<Expr>),
    RotateLeft(Box<Expr>, Box<Expr>),
    RotateRight(Box<Expr>, Box<Expr>),
    UnaryFactorial(Box<Expr>),
    UnaryNegation(Box<Expr>),
    UnaryNot(Box<Expr>),
    FunctionCall(String, Vec<Expr>),
}

impl PartialEq for Expr {
    fn eq(&self, other: &Self) -> bool {
        match (self, other) {
            (Expr::Number(a), Expr::Number(b)) => a == b,
            (Expr::Percent(a), Expr::Percent(b)) => a == b,
            (Expr::Add(a1, b1), Expr::Add(a2, b2)) => a1 == a2 && b1 == b2,
            (Expr::Subtract(a1, b1), Expr::Subtract(a2, b2)) => a1 == a2 && b1 == b2,
            (Expr::Multiply(a1, b1), Expr::Multiply(a2, b2)) => a1 == a2 && b1 == b2,
            (Expr::Divide(a1, b1), Expr::Divide(a2, b2)) => a1 == a2 && b1 == b2,
            (Expr::Modulo(a1, b1), Expr::Modulo(a2, b2)) => a1 == a2 && b1 == b2,
            (Expr::Power(a1, b1), Expr::Power(a2, b2)) => a1 == a2 && b1 == b2,
            (Expr::And(a1, b1), Expr::And(a2, b2)) => a1 == a2 && b1 == b2,
            (Expr::Or(a1, b1), Expr::Or(a2, b2)) => a1 == a2 && b1 == b2,
            (Expr::Xor(a1, b1), Expr::Xor(a2, b2)) => a1 == a2 && b1 == b2,
            (Expr::ShiftLeft(a1, b1), Expr::ShiftLeft(a2, b2)) => a1 == a2 && b1 == b2,
            (Expr::ShiftRight(a1, b1), Expr::ShiftRight(a2, b2)) => a1 == a2 && b1 == b2,
            (Expr::RotateLeft(a1, b1), Expr::RotateLeft(a2, b2)) => a1 == a2 && b1 == b2,
            (Expr::RotateRight(a1, b1), Expr::RotateRight(a2, b2)) => a1 == a2 && b1 == b2,
            (Expr::UnaryFactorial(a1), Expr::UnaryFactorial(a2)) => a1 == a2,
            (Expr::UnaryNegation(a1), Expr::UnaryNegation(a2)) => a1 == a2,
            (Expr::UnaryNot(a1), Expr::UnaryNot(a2)) => a1 == a2,
            (Expr::FunctionCall(a1, b1), Expr::FunctionCall(a2, b2)) => a1 == a2 && b1 == b2,
            _ => false,
        }
    }
}

impl Expr {
    pub fn eval(&self) -> f64 {
        match self {
            Expr::Number(n) => Self::parse_number(n.clone()),
            Expr::Percent(p) => Self::parse_number(p.clone()) / 100.0,
            Expr::Add(v1, v2) => {
                if let Expr::Percent(n) = v2.as_ref() {
                    v1.eval() * (1.0 + Self::parse_number(n.clone()) / 100.0)
                } else {
                    v1.eval() + v2.eval()
                }
            }
            Expr::Subtract(v1, v2) => {
                if let Expr::Percent(n) = v2.as_ref() {
                    v1.eval() - (v1.eval() * Self::parse_number(n.clone()) / 100.0)
                } else {
                    v1.eval() - v2.eval()
                }
            }
            Expr::Multiply(v1, v2) => v1.eval() * v2.eval(),
            Expr::Divide(v1, v2) => {
                v1.eval() / v2.eval()
            }
            Expr::Modulo(v1, v2) => v1.eval() % v2.eval(),
            Expr::Power(v1, v2) => v1.eval().powf(v2.eval()),
            Expr::And(v1, v2) => (v1.eval() as u64 & v2.eval() as u64) as f64,
            Expr::Or(v1, v2) => (v1.eval() as u64 | v2.eval() as u64) as f64,
            Expr::Xor(v1, v2) => (v1.eval() as u64 ^ v2.eval() as u64) as f64,
            Expr::ShiftLeft(v1, v2) => ((v1.eval() as u64) << (v2.eval() as u64)) as f64,
            Expr::ShiftRight(v1, v2) => ((v1.eval() as u64) >> (v2.eval() as u64)) as f64,
            Expr::RotateLeft(v1, v2) => {
                let n = v1.eval() as u64;
                let m = v2.eval() as u64;
                let s = 64 - m;
                ((n << m) | (n >> s)) as f64
            }
            Expr::RotateRight(v1, v2) => {
                let n = v1.eval() as u64;
                let m = v2.eval() as u64;
                let s = 64 - m;
                ((n >> m) | (n << s)) as f64
            }
            Expr::UnaryFactorial(v) => {
                let n = v.eval() as u64;
                let mut result = 1;
                for i in 1..=n {
                    result *= i;
                }
                result as f64
            }
            Expr::UnaryNegation(v) => -v.eval(),
            Expr::UnaryNot(v) => !(v.eval() as u64) as f64,
            Expr::FunctionCall(name, args) => {
                match name.as_str() {
                    "abs" => args[0].eval().abs(),
                    "acos" => args[0].eval().acos(),
                    "acosh" => args[0].eval().acosh(),
                    "asin" => args[0].eval().asin(),
                    "asinh" => args[0].eval().asinh(),
                    "atan" => args[0].eval().atan(),
                    "atanh" => args[0].eval().atanh(),
                    "cbrt" => args[0].eval().cbrt(),
                    "ceil" => args[0].eval().ceil(),
                    "cos" => args[0].eval().cos(),
                    "cosh" => args[0].eval().cosh(),
                    "exp" => args[0].eval().exp(),
                    "floor" => args[0].eval().floor(),
                    "ln" => args[0].eval().ln(),
                    "log" => args[0].eval().log10(),
                    "log2" => args[0].eval().log2(),
                    "round" => args[0].eval().round(),
                    "sin" => args[0].eval().sin(),
                    "sinh" => args[0].eval().sinh(),
                    "sqrt" => args[0].eval().sqrt(),
                    "tan" => args[0].eval().tan(),
                    "tanh" => args[0].eval().tanh(),
                    "trunc" => args[0].eval().trunc(),
                    _ => panic!("unknown function: {}", name),
                }
            }
        }
    }

    // TODO: handle thousands separator, custom decimal separator, and scientific notation
    fn parse_number(n: String) -> f64 {
        if n == "pi" {
            std::f64::consts::PI
        } else if n == "e" {
            std::f64::consts::E
        } else {
            n.parse().unwrap()
        }
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_eval() {
        assert_eq!(
            crate::parse("1+2").unwrap().eval(),
            3.0
        );

        assert_eq!(
            crate::parse("123").unwrap().eval(),
            123.0
        );

        assert_eq!(
            crate::parse("123.456").unwrap().eval(),
            123.456
        );

        // assert_eq!(
        //     crate::parse("123,45,.6").unwrap().eval(),
        //     0.0
        // );

        assert_eq!(
            crate::parse("+123").unwrap().eval(),
            123.0
        );

        assert_eq!(
            crate::parse("-123").unwrap().eval(),
            -123.0
        );

        assert_eq!(
            crate::parse("~123").unwrap().eval(),
            (!(123 as u64)) as f64
        );

        assert_eq!(
            crate::parse("1 - -2").unwrap().eval(),
            3.0
        );

        assert_eq!(
            crate::parse("3 - ~4").unwrap().eval(),
            3.0 - (!(4 as u64)) as f64
        );

        assert_eq!(
            crate::parse("123%").unwrap().eval(),
            1.23
        );

        // assert_eq!(
        //     crate::parse("1,23.456%").unwrap().eval(),
        //     0.0
        // );

        assert_eq!(
            crate::parse("123 % 4").unwrap().eval(),
            3.0
        );

        assert_eq!(
            crate::parse("(123%)").unwrap().eval(),
            1.23
        );

        assert_eq!(
            crate::parse("123%pi").unwrap().eval(),
            1.23 * std::f64::consts::PI
        );

        assert_eq!(
            crate::parse("5!").unwrap().eval(),
            120.0
        );

        assert_eq!(
            crate::parse("5!*3").unwrap().eval(),
            360.0
        );

        assert_eq!(
            crate::parse("1 + 2").unwrap().eval(),
            3.0
        );

        assert_eq!(
            crate::parse("1 - 2").unwrap().eval(),
            -1.0
        );

        assert_eq!(
            crate::parse("3 * 2").unwrap().eval(),
            6.0
        );

        assert_eq!(
            crate::parse("6 / 4").unwrap().eval(),
            1.5
        );

        assert_eq!(
            crate::parse("3 ^ 2").unwrap().eval(),
            9.0
        );

        assert_eq!(
            crate::parse("9 % 2").unwrap().eval(),
            1.0
        );

        assert_eq!(
            crate::parse("1 + 2 + 3").unwrap().eval(),
            6.0
        );

        assert_eq!(
            crate::parse("1 + 2 * 3").unwrap().eval(),
            7.0
        );

        assert_eq!(
            crate::parse("3 ** 2").unwrap().eval(),
            9.0
        );

        assert_eq!(
            crate::parse("4 ^ 2 ^ 3").unwrap().eval(),
            65536.0
        );

        assert_eq!(
            crate::parse("-3 + 4").unwrap().eval(),
            1.0
        );

        assert_eq!(
            crate::parse("3 & 4").unwrap().eval(),
            0.0
        );

        assert_eq!(
            crate::parse("3 | 4").unwrap().eval(),
            7.0
        );

        assert_eq!(
            crate::parse("3 xor 4").unwrap().eval(),
            7.0
        );

        assert_eq!(
            crate::parse("3 << 4").unwrap().eval(),
            48.0
        );

        assert_eq!(
            crate::parse("3 >> 4").unwrap().eval(),
            0.0
        );

        assert_eq!(
            crate::parse("3 rol 4").unwrap().eval(),
            48.0
        );

        assert_eq!(
            crate::parse("3 ror 4").unwrap().eval(),
            0x3000000000000000i64 as f64
        );

        assert_eq!(
            crate::parse("3 mod 5").unwrap().eval(),
            3.0
        );

        assert_eq!(
            // (4 << 5 >> 6 rol 7 ror 8) has precedence over (1 & 2 | 3) at the xor operator
            crate::parse("1 & 2 | 3 xor 4 << 5 >> 6 rol 7 ror 8").unwrap().eval(),
            2.0
        );

        assert_eq!(
            crate::parse("(1 + 2) * 3").unwrap().eval(),
            9.0
        );

        assert_eq!(
            crate::parse("1 + (2 * 3)").unwrap().eval(),
            7.0
        );

        assert_eq!(
            crate::parse("3%(3 + 4)").unwrap().eval(), // modulo
            3.0
        );

        assert_eq!(
            crate::parse("3(5 + 4)").unwrap().eval(),
            27.0
        );

        assert_eq!(
            crate::parse("(3 + 2   )").unwrap().eval(),
            5.0
        );

        assert_eq!(
            crate::parse("-(3 + 2)").unwrap().eval(),
            -5.0
        );

        assert_eq!(
            crate::parse("pi").unwrap().eval(),
            std::f64::consts::PI
        );

        assert_eq!(
            crate::parse("e").unwrap().eval(),
            std::f64::consts::E
        );

        assert_eq!(
            crate::parse("sin(3)").unwrap().eval(),
            0.1411200080598672
        );

        assert_eq!(
            crate::parse("cos(3)").unwrap().eval(),
            -0.9899924966004454
        );

        assert_eq!(
            crate::parse("tan3").unwrap().eval(),
            -0.1425465430742778
        );

        assert_eq!(
            crate::parse("tan 3").unwrap().eval(),
            -0.1425465430742778
        );

        assert_eq!(
            crate::parse("log(3)").unwrap().eval(),
            0.47712125471966244
        );

        assert_eq!(
            crate::parse("ln(3)").unwrap().eval(),
            1.0986122886681098
        );

        assert_eq!(
            crate::parse("sin(cos(pi))").unwrap().eval(),
            -0.8414709848078965
        );

        assert_eq!(
            crate::parse("sin cos pi").unwrap().eval(),
            -0.8414709848078965
        );

        assert_eq!(
            crate::parse("3pi").unwrap().eval(),
            3.0 * std::f64::consts::PI
        );

        assert_eq!(
            crate::parse("3e").unwrap().eval(),
            3.0 * std::f64::consts::E
        );

        assert_eq!(
            crate::parse("3sin(3)").unwrap().eval(),
            3.0 * 0.1411200080598672
        );

        assert_eq!(
            crate::parse("3cos(3)").unwrap().eval(),
            3.0 * -0.9899924966004454
        );

        assert_eq!(
            crate::parse("3tan3").unwrap().eval(),
            3.0 * -0.1425465430742778
        );

        assert_eq!(
            crate::parse("3tan 3").unwrap().eval(),
            3.0 * -0.1425465430742778
        );

        assert_eq!(
            crate::parse("3 * pi").unwrap().eval(),
            3.0 * std::f64::consts::PI
        );

        assert_eq!(
            crate::parse("3 + e").unwrap().eval(),
            3.0 + std::f64::consts::E
        );

        assert_eq!(
            crate::parse("3 - sin3").unwrap().eval(),
            3.0 - 0.1411200080598672
        );

        assert_eq!(
            crate::parse("pi sin 3").unwrap().eval(),
            std::f64::consts::PI * 0.1411200080598672
        );

        assert_eq!(
            crate::parse("sin 3^2").unwrap().eval(),
            9.0f64.sin()
        );

        assert_eq!(
            crate::parse("sin 3^2^3").unwrap().eval(),
            6561.0f64.sin()
        );

        assert_eq!(
            crate::parse("sin 3+2").unwrap().eval(),
            2.1411200080598672
        );

        assert_eq!(
            crate::parse("log2 256").unwrap().eval(),
            8.0
        );

        assert_eq!(
            crate::parse("log 2 + 256").unwrap().eval(),
            256.3010299956639812
        );

        assert_eq!(
            crate::parse("e ^ 2 - 3 * pi + 10%").unwrap().eval(),
            -2.239294048022603
        );

        assert_eq!(
            crate::parse("2(3 + sin pi - 4^2) / e").unwrap().eval(),
            -9.564865470457502
        );

        assert_eq!(
            crate::parse("sin(cos(tan 45)) + pi * 2e").unwrap().eval(),
            17.030528719035757
        );

        assert_eq!(
            crate::parse("2 * pi sin(30)").unwrap().eval(),
            -6.207985783529054
        );

        assert_eq!(
            crate::parse("pi e + 2pi * 3e").unwrap().eval(),
            59.77813955871496
        );

        assert_eq!(
            crate::parse("pi% * e").unwrap().eval(),
            0.08539734222673567
        );

        assert_eq!(
            crate::parse("sin((45 + (30 * 2)) / (3 ^ 2))").unwrap().eval(),
            -0.78314284623659
        );
    }
}
