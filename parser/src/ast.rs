use astro_float::BigFloat;

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
                    v1.eval() + (v1.eval() * Self::parse_number(n.clone()) / 100.0)
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
            Expr::Divide(v1, v2) => v1.eval() / v2.eval(),
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
                    "log" => args[0].eval().log(10.0),
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