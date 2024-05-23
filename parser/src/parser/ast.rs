#[derive(Debug, PartialEq)]
pub enum Expr {
    Number(String, String), // (value, unit)
    WithUnit(Box<Expr>, String),
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

fn sum_units(unit1: String, unit2: String) -> String {
    if unit1 == unit2 {
        unit1
    } else if unit1.is_empty() {
        unit2
    } else if unit2.is_empty() {
        unit1
    } else {
        panic!("incompatible units: {} and {}", unit1, unit2);
    }
}

fn multiply_units(unit1: String, unit2: String) -> String {
    if unit1.is_empty() || unit2.is_empty() {
        "".to_string()
    } else {
        format!("{}*{}", unit1, unit2)
    }
}

fn divide_units(unit1: String, unit2: String) -> String {
    if unit1.is_empty() || unit2.is_empty() {
        "".to_string()
    } else {
        format!("{}/{}", unit1, unit2)
    }
}

impl Expr {
    pub fn eval(&self) -> (f64, String) {
        match self {
            Expr::Number(n, unit) => (Self::parse_number(n.clone()), unit.clone()),
            Expr::WithUnit(e, unit) => {
                let (_e, _unit) = e.eval();
                if !_unit.is_empty() {
                    panic!("unit already specified");
                }
                (_e, unit.clone())
            }
            Expr::Percent(p) => (Self::parse_number(p.clone()) / 100.0, "".to_string()),
            Expr::Add(v1, v2) => {
                if let Expr::Percent(n) = v2.as_ref() {
                    let (_v1, unit) = v1.eval();
                    let _v2 = 1.0 + Self::parse_number(n.clone()) / 100.0;

                    (_v1 * _v2, unit)
                } else {
                    let (_v1, unit1) = v1.eval();
                    let (_v2, unit2) = v2.eval();

                    (_v1 + _v2, sum_units(unit1, unit2))
                }
            }
            Expr::Subtract(v1, v2) => {
                if let Expr::Percent(n) = v2.as_ref() {
                    let (_v1, unit) = v1.eval();
                    let _v2 = 1.0 - Self::parse_number(n.clone()) / 100.0;

                    (_v1 * _v2, unit)
                } else {
                    let (_v1, unit1) = v1.eval();
                    let (_v2, unit2) = v2.eval();

                    (_v1 - _v2, sum_units(unit1, unit2))
                }
            }
            Expr::Multiply(v1, v2) => {
                let (_v1, unit1) = v1.eval();
                let (_v2, unit2) = v2.eval();

                (_v1 * _v2, multiply_units(unit1, unit2))
            }
            Expr::Divide(v1, v2) => {
                let (_v1, unit1) = v1.eval();
                let (_v2, unit2) = v2.eval();

                (_v1 / _v2, divide_units(unit1, unit2))
            }
            Expr::Modulo(v1, v2) => {
                let (_v1, unit1) = v1.eval();
                let (_v2, unit2) = v2.eval();

                (_v1 % _v2, sum_units(unit1, unit2))
            }
            Expr::Power(v1, v2) => {
                let (_v1, unit1) = v1.eval();
                let (_v2, unit2) = v2.eval();

                if !unit2.is_empty() {
                    panic!("unit not allowed for exponentiation");
                } else {
                    (_v1.powf(_v2), unit1)
                }
            }
            Expr::And(v1, v2) => {
                let (_v1, unit1) = v1.eval();
                let (_v2, unit2) = v2.eval();
                let res = ((_v1 as u64) & (_v2 as u64)) as f64;

                if unit1.is_empty() {
                    (res, unit2)
                } else {
                    (res, unit1)
                }
            }
            Expr::Or(v1, v2) => {
                let (_v1, unit1) = v1.eval();
                let (_v2, unit2) = v2.eval();
                let res = ((_v1 as u64) | (_v2 as u64)) as f64;

                if unit1.is_empty() {
                    (res, unit2)
                } else {
                    (res, unit1)
                }
            }
            Expr::Xor(v1, v2) => {
                let (_v1, unit1) = v1.eval();
                let (_v2, unit2) = v2.eval();
                let res = ((_v1 as u64) ^ (_v2 as u64)) as f64;

                if unit1.is_empty() {
                    (res, unit2)
                } else {
                    (res, unit1)
                }
            }
            Expr::ShiftLeft(v1, v2) => {
                let (_v1, unit1) = v1.eval();
                let (_v2, unit2) = v2.eval();
                let res = ((_v1 as u64) << (_v2 as u64)) as f64;

                if !unit2.is_empty() {
                    panic!("unit not allowed for left shift");
                } else {
                    (res, unit1)
                }
            }
            Expr::ShiftRight(v1, v2) => {
                let (_v1, unit1) = v1.eval();
                let (_v2, unit2) = v2.eval();
                let res = ((_v1 as u64) >> (_v2 as u64)) as f64;

                if !unit2.is_empty() {
                    panic!("unit not allowed for right shift");
                } else {
                    (res, unit1)
                }
            }
            Expr::RotateLeft(v1, v2) => {
                let (_v1, unit1) = v1.eval();
                let (_v2, unit2) = v2.eval();
                let n = _v1 as u64;
                let m = _v2 as u64;
                let s = 64 - m;
                let res = ((n << m) | (n >> s)) as f64;

                if !unit2.is_empty() {
                    panic!("unit not allowed for left rotation");
                } else {
                    (res, unit1)
                }
            }
            Expr::RotateRight(v1, v2) => {
                let (_v1, unit1) = v1.eval();
                let (_v2, unit2) = v2.eval();
                let n = _v1 as u64;
                let m = _v2 as u64;
                let s = 64 - m;
                let res = ((n >> m) | (n << s)) as f64;

                if !unit2.is_empty() {
                    panic!("unit not allowed for right rotation");
                } else {
                    (res, unit1)
                }
            }
            Expr::UnaryFactorial(v) => {
                let (_v, unit) = v.eval();
                let n = _v as u64;
                let mut result = 1;
                for i in 1..=n {
                    result *= i;
                }
                (result as f64, unit)
            }
            Expr::UnaryNegation(v) => {
                let (_v, unit) = v.eval();
                (-_v, unit)
            }
            Expr::UnaryNot(v) => {
                let (_v, unit) = v.eval();
                ((!(_v as u64)) as f64, unit)
            }
            Expr::FunctionCall(name, args) => match name.as_str() {
                "abs" => {
                    let (_v, unit) = args[0].eval();
                    (_v.abs(), unit)
                }
                "acos" => {
                    let (_v, unit) = args[0].eval();
                    (_v.acos(), unit)
                }
                "acosh" => {
                    let (_v, unit) = args[0].eval();
                    (_v.acosh(), unit)
                }
                "asin" => {
                    let (_v, unit) = args[0].eval();
                    (_v.asin(), unit)
                }
                "asinh" => {
                    let (_v, unit) = args[0].eval();
                    (_v.asinh(), unit)
                }
                "atan" => {
                    let (_v, unit) = args[0].eval();
                    (_v.atan(), unit)
                }
                "atanh" => {
                    let (_v, unit) = args[0].eval();
                    (_v.atanh(), unit)
                }
                "cbrt" => {
                    let (_v, unit) = args[0].eval();
                    (_v.cbrt(), unit)
                }
                "ceil" => {
                    let (_v, unit) = args[0].eval();
                    (_v.ceil(), unit)
                }
                "cos" => {
                    let (_v, unit) = args[0].eval();
                    (_v.cos(), unit)
                }
                "cosh" => {
                    let (_v, unit) = args[0].eval();
                    (_v.cosh(), unit)
                }
                "exp" => {
                    let (_v, unit) = args[0].eval();
                    (_v.exp(), unit)
                }
                "floor" => {
                    let (_v, unit) = args[0].eval();
                    (_v.floor(), unit)
                }
                "ln" => {
                    let (_v, unit) = args[0].eval();
                    (_v.ln(), unit)
                }
                "log" => {
                    let (_v, unit) = args[0].eval();
                    (_v.log10(), unit)
                }
                "log2" => {
                    let (_v, unit) = args[0].eval();
                    (_v.log2(), unit)
                }
                "rand" => {
                    let v = rand::random::<f64>();
                    (v, "".to_string())
                }
                "round" => {
                    let (_v, unit) = args[0].eval();
                    (_v.round(), unit)
                }
                "sin" => {
                    let (_v, unit) = args[0].eval();
                    (_v.sin(), unit)
                }
                "sinh" => {
                    let (_v, unit) = args[0].eval();
                    (_v.sinh(), unit)
                }
                "sqrt" => {
                    let (_v, unit) = args[0].eval();
                    (_v.sqrt(), unit)
                }
                "tan" => {
                    let (_v, unit) = args[0].eval();
                    (_v.tan(), unit)
                }
                "tanh" => {
                    let (_v, unit) = args[0].eval();
                    (_v.tanh(), unit)
                }
                "trunc" => {
                    let (_v, unit) = args[0].eval();
                    (_v.trunc(), unit)
                }
                _ => panic!("unknown function: {}", name),
            },
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
        assert_eq!(crate::parse("1+2").unwrap().eval(), (3.0, "".to_string()));

        assert_eq!(crate::parse("123").unwrap().eval(), (123.0, "".to_string()));

        assert_eq!(
            crate::parse("123.456").unwrap().eval(),
            (123.456, "".to_string())
        );

        // assert_eq!(
        //     crate::parse("123,45,.6").unwrap().eval(),
        //     0.0
        // );

        assert_eq!(
            crate::parse("+123").unwrap().eval(),
            (123.0, "".to_string())
        );

        assert_eq!(
            crate::parse("-123").unwrap().eval(),
            (-123.0, "".to_string())
        );

        assert_eq!(
            crate::parse("~123").unwrap().eval(),
            (((!(123 as u64)) as f64), "".to_string())
        );

        assert_eq!(
            crate::parse("1 - -2").unwrap().eval(),
            (3.0, "".to_string())
        );

        assert_eq!(
            crate::parse("3 - ~4").unwrap().eval(),
            (3.0 - (!(4 as u64)) as f64, "".to_string())
        );

        assert_eq!(crate::parse("123%").unwrap().eval(), (1.23, "".to_string()));

        // assert_eq!(
        //     crate::parse("1,23.456%").unwrap().eval(),
        //     0.0
        // );

        assert_eq!(
            crate::parse("123 % 4").unwrap().eval(),
            (3.0, "".to_string())
        );

        assert_eq!(
            crate::parse("(123%)").unwrap().eval(),
            (1.23, "".to_string())
        );

        assert_eq!(
            crate::parse("123%pi").unwrap().eval(),
            (1.23 * std::f64::consts::PI, "".to_string())
        );

        assert_eq!(
            crate::parse("1 + 123% of (3 + 4)").unwrap().eval(),
            (9.61, "".to_string())
        );

        assert_eq!(crate::parse("5!").unwrap().eval(), (120.0, "".to_string()));

        assert_eq!(
            crate::parse("5!*3").unwrap().eval(),
            (360.0, "".to_string())
        );

        assert_eq!(crate::parse("1 + 2").unwrap().eval(), (3.0, "".to_string()));

        assert_eq!(
            crate::parse("1 - 2").unwrap().eval(),
            (-1.0, "".to_string())
        );

        assert_eq!(crate::parse("3 * 2").unwrap().eval(), (6.0, "".to_string()));

        assert_eq!(crate::parse("6 / 4").unwrap().eval(), (1.5, "".to_string()));

        assert_eq!(crate::parse("3 ^ 2").unwrap().eval(), (9.0, "".to_string()));

        assert_eq!(crate::parse("9 % 2").unwrap().eval(), (1.0, "".to_string()));

        assert_eq!(
            crate::parse("1 + 2 + 3").unwrap().eval(),
            (6.0, "".to_string())
        );

        assert_eq!(
            crate::parse("1 + 2 * 3").unwrap().eval(),
            (7.0, "".to_string())
        );

        assert_eq!(
            crate::parse("3 ** 2").unwrap().eval(),
            (9.0, "".to_string())
        );

        assert_eq!(
            crate::parse("4 ^ 2 ^ 3").unwrap().eval(),
            (65536.0, "".to_string())
        );

        assert_eq!(
            crate::parse("-3 + 4").unwrap().eval(),
            (1.0, "".to_string())
        );

        assert_eq!(crate::parse("3 & 4").unwrap().eval(), (0.0, "".to_string()));

        assert_eq!(crate::parse("3 | 4").unwrap().eval(), (7.0, "".to_string()));

        assert_eq!(
            crate::parse("3 xor 4").unwrap().eval(),
            (7.0, "".to_string())
        );

        assert_eq!(
            crate::parse("3 << 4").unwrap().eval(),
            (48.0, "".to_string())
        );

        assert_eq!(
            crate::parse("3 >> 4").unwrap().eval(),
            (0.0, "".to_string())
        );

        assert_eq!(
            crate::parse("3 rol 4").unwrap().eval(),
            (48.0, "".to_string())
        );

        assert_eq!(
            crate::parse("3 ror 4").unwrap().eval(),
            (0x3000000000000000i64 as f64, "".to_string())
        );

        assert_eq!(
            crate::parse("3 mod 5").unwrap().eval(),
            (3.0, "".to_string())
        );

        assert_eq!(
            // (4 << 5 >> 6 rol 7 ror 8) has precedence over (1 & 2 | 3) at the xor operator
            crate::parse("1 & 2 | 3 xor 4 << 5 >> 6 rol 7 ror 8")
                .unwrap()
                .eval(),
            (2.0, "".to_string())
        );

        assert_eq!(
            crate::parse("(1 + 2) * 3").unwrap().eval(),
            (9.0, "".to_string())
        );

        assert_eq!(
            crate::parse("1 + (2 * 3)").unwrap().eval(),
            (7.0, "".to_string())
        );

        assert_eq!(
            crate::parse("3%(3 + 4)").unwrap().eval(), // modulo
            (3.0, "".to_string())
        );

        assert_eq!(
            crate::parse("3(5 + 4)").unwrap().eval(),
            (27.0, "".to_string())
        );

        assert_eq!(
            crate::parse("(3 + 2   )").unwrap().eval(),
            (5.0, "".to_string())
        );

        assert_eq!(
            crate::parse("-(3 + 2)").unwrap().eval(),
            (-5.0, "".to_string())
        );

        assert_eq!(
            crate::parse("pi").unwrap().eval(),
            (std::f64::consts::PI, "".to_string())
        );

        assert_eq!(
            crate::parse("e").unwrap().eval(),
            (std::f64::consts::E, "".to_string())
        );

        assert_eq!(
            crate::parse("sin(3)").unwrap().eval(),
            (0.1411200080598672, "".to_string())
        );

        assert_eq!(
            crate::parse("cos(3)").unwrap().eval(),
            (-0.9899924966004454, "".to_string())
        );

        assert_eq!(
            crate::parse("tan3").unwrap().eval(),
            (-0.1425465430742778, "".to_string())
        );

        assert_eq!(
            crate::parse("tan 3").unwrap().eval(),
            (-0.1425465430742778, "".to_string())
        );

        assert_eq!(
            crate::parse("log(3)").unwrap().eval(),
            (0.47712125471966244, "".to_string())
        );

        assert_eq!(
            crate::parse("ln(3)").unwrap().eval(),
            (1.0986122886681098, "".to_string())
        );

        assert_eq!(
            crate::parse("sin(cos(pi))").unwrap().eval(),
            (-0.8414709848078965, "".to_string())
        );

        assert_eq!(
            crate::parse("sin cos pi").unwrap().eval(),
            (-0.8414709848078965, "".to_string())
        );

        assert_eq!(
            crate::parse("3pi").unwrap().eval(),
            (3.0 * std::f64::consts::PI, "".to_string())
        );

        assert_eq!(
            crate::parse("3e").unwrap().eval(),
            (3.0 * std::f64::consts::E, "".to_string())
        );

        assert_eq!(
            crate::parse("3sin(3)").unwrap().eval(),
            (3.0 * 0.1411200080598672, "".to_string())
        );

        assert_eq!(
            crate::parse("3cos(3)").unwrap().eval(),
            (3.0 * -0.9899924966004454, "".to_string())
        );

        assert_eq!(
            crate::parse("3tan3").unwrap().eval(),
            (3.0 * -0.1425465430742778, "".to_string())
        );

        assert_eq!(
            crate::parse("3tan 3").unwrap().eval(),
            (3.0 * -0.1425465430742778, "".to_string())
        );

        assert_eq!(
            crate::parse("3 * pi").unwrap().eval(),
            (3.0 * std::f64::consts::PI, "".to_string())
        );

        assert_eq!(
            crate::parse("3 + e").unwrap().eval(),
            (3.0 + std::f64::consts::E, "".to_string())
        );

        assert_eq!(
            crate::parse("3 - sin3").unwrap().eval(),
            (3.0 - 0.1411200080598672, "".to_string())
        );

        assert_eq!(
            crate::parse("pi sin 3").unwrap().eval(),
            (std::f64::consts::PI * 0.1411200080598672, "".to_string())
        );

        assert_eq!(
            crate::parse("sin 3^2").unwrap().eval(),
            (9.0f64.sin(), "".to_string())
        );

        assert_eq!(
            crate::parse("sin 3^2^3").unwrap().eval(),
            (6561.0f64.sin(), "".to_string())
        );

        assert_eq!(
            crate::parse("sin 3+2").unwrap().eval(),
            (2.1411200080598672, "".to_string())
        );

        assert_eq!(
            crate::parse("log2 256").unwrap().eval(),
            (8.0, "".to_string())
        );

        assert_eq!(
            crate::parse("log 2 + 256").unwrap().eval(),
            (256.3010299956639812, "".to_string())
        );

        assert_eq!(
            crate::parse("e ^ 2 - 3 * pi + 10%").unwrap().eval(),
            (-2.239294048022603, "".to_string())
        );

        assert_eq!(
            crate::parse("2(3 + sin pi - 4^2) / e").unwrap().eval(),
            (-9.564865470457502, "".to_string())
        );

        assert_eq!(
            crate::parse("sin(cos(tan 45)) + pi * 2e").unwrap().eval(),
            (17.030528719035757, "".to_string())
        );

        assert_eq!(
            crate::parse("2 * pi sin(30)").unwrap().eval(),
            (-6.207985783529054, "".to_string())
        );

        assert_eq!(
            crate::parse("pi e + 2pi * 3e").unwrap().eval(),
            (59.77813955871496, "".to_string())
        );

        assert_eq!(
            crate::parse("pi% * e").unwrap().eval(),
            (0.08539734222673567, "".to_string())
        );

        assert_eq!(
            crate::parse("sin((45 + (30 * 2)) / (3 ^ 2))")
                .unwrap()
                .eval(),
            (-0.78314284623659, "".to_string())
        );

        assert_eq!(
            crate::parse("$1 + 2$").unwrap().eval(),
            (3.0, "$".to_string())
        );
    }
}
