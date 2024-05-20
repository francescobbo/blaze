#[derive(Debug)]
pub enum Expr {
    Number(i32),
    Percent(i32),
    Neg(Box<Expr>),
    Conversion(Box<Expr>, String),
    Op(Box<Expr>, Opcode, Box<Expr>),
}

#[derive(Debug)]
pub enum Opcode {
    Add,
    Sub,
    Mul,
    Div,
    Pow,
    Mod,
}
