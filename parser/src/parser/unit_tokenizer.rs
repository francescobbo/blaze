use std::str::Chars;

pub struct UnitTokenizer<'input> {
    state: State,
    input: Chars<'input>,
    pos: usize,

    matched: String,
    current_word: String,

    non_units: Vec<&'input str>,

    nesting: usize,
}

#[derive(Debug, PartialEq)]
enum State {
    BeforeUnit,
    Unit,
    ImplicitPower,
    Operator,
    BeforePower,
    InPower,
    SpaceAfterUnit,
    End
}

/// The unit tokenizer generates a single token from a string, greedily trying to match the longest possible unit.
/// It expects the input to start with a viable unit. 
/// The non_units parameter is a list of strings that should not be considered units and halt the matching process if encountered.
impl<'input> UnitTokenizer<'input> {
    pub fn new(input: Chars<'input>, non_units: Vec<&'input str>) -> Self {
        Self {
            state: State::BeforeUnit,
            input,
            pos: 0,

            matched: String::new(),
            current_word: String::new(),

            non_units,

            nesting: 0
        }
    }

    fn new_nested(&self, input: Chars<'input>) -> Self {
        Self {
            state: State::BeforeUnit,
            input,
            pos: 0,

            matched: String::new(),
            current_word: String::new(),

            non_units: self.non_units.clone(),

            nesting: self.nesting + 1
        }
    }

    pub fn step(&mut self) {
        if self.nesting > 20 {
            self.state = State::End;
            return;
        }

        match self.state {
            State::BeforeUnit => {
                match self.peek() {
                    Some(c) if c.is_alphabetic() => {
                        self.state = State::Unit;
                    }
                    Some(c) if c.is_whitespace() => {
                        self.consume();
                    }
                    Some('/') => {
                        self.state = State::Operator;
                    }
                    _ => {
                        self.state = State::End;
                    }
                }
            }
            State::Unit => {
                match self.peek() {
                    Some(c) if c.is_alphabetic() => {
                        self.consume();
                        self.current_word.push(c);

                        if self.non_units.contains(&self.current_word.as_str()) {
                            self.state = State::End;
                        }
                    }
                    Some(c) if c.is_digit(10) => {
                        self.current_word.push('^');
                        self.state = State::ImplicitPower;
                    }
                    Some(' ') => {
                        self.matched.push_str(&self.current_word);
                        self.state = State::SpaceAfterUnit;
                    }
                    Some('*') | Some('/') | Some('^') => {
                        self.matched.push_str(&self.current_word);
                        self.current_word.clear();

                        self.state = State::Operator;
                    }
                    _ => {
                        self.matched.push_str(&self.current_word);
                        self.state = State::End;
                    }
                }
            }
            State::ImplicitPower => {
                match self.peek() {
                    Some(c) if c.is_digit(10) => {
                        self.consume();
                        self.current_word.push(c);
                    }
                    Some(c) if c.is_whitespace() => {
                        self.consume();
                        self.state = State::SpaceAfterUnit;
                    }
                    Some('*') | Some('/') => {
                        self.matched.push_str(&self.current_word);
                        self.current_word.clear();

                        self.state = State::Operator;
                    }
                    _ => {
                        self.matched.push_str(&self.current_word);
                        self.current_word.clear();

                        self.state = State::End;
                    }
                }
            }
            State::Operator => {
                match self.peek().unwrap() {
                    '^' => {
                        self.consume();
                        self.matched.push('^');

                        self.state = State::BeforePower;
                    }
                    '*' => {
                        // Try to match more units recursively
                        let mut remainder = self.input.clone();
                        remainder.next();

                        let mut nested = self.new_nested(remainder);
                        let rhs = nested.run();

                        if rhs != "" {
                            self.matched.push('*');
                            self.matched.push_str(&rhs);

                            self.skip(nested.pos + 1);

                            self.state = State::Unit;
                        } else {
                            // The right hand side is not a unit, so this is a regular multiplication (eg: "2kg * 3")
                            self.state = State::End;
                        }
                    }
                    '/' => {
                        // Try to match more units recursively
                        let mut remainder = self.input.clone();
                        remainder.next();

                        let is_braketed = remainder.clone().nth(0) == Some('(');
                        if is_braketed {
                            remainder.next();
                        }

                        let mut nested = self.new_nested(remainder);
                        let rhs = nested.run();

                        if rhs != "" {
                            if rhs.chars().nth(0).unwrap() == '/' {
                                // Something like "m/(/", so we stop the matching process
                                self.state = State::End;
                                return;
                            }

                            self.skip(nested.pos + 1 + if is_braketed { 1 } else { 0 });
                            if is_braketed && self.peek() != Some(')') {
                                // Something like "m/(s+2", so we stop the matching process
                                self.state = State::End;
                                return;
                            }

                            if is_braketed {
                                self.matched.push_str("/(");
                            } else {
                                self.matched.push('/');
                            }

                            self.matched.push_str(&rhs);

                            if is_braketed {
                                self.matched.push(')');
                                self.consume();
                            }

                            self.state = State::Unit;
                        } else {
                            // The right hand side is not a unit, so this is a regular multiplication (eg: "2kg * 3")
                            self.state = State::End;
                        }
                    }
                    _ => unreachable!()
                }
            }
            State::BeforePower => {
                match self.peek() {
                    Some(c) if c.is_digit(10) || c == '-' => {
                        // Eg: "m^2"
                        self.state = State::InPower;
                    }
                    Some(c) if c.is_whitespace() => {
                        // Eg: "m^ "
                        self.consume();
                    }
                    _ => {
                        self.state = State::End;
                    }
                }
            }
            State::InPower => {
                match self.peek() {
                    Some(c) if c.is_digit(10) || c == '-' => {
                        // Eg: "m^2"
                        self.consume();
                        self.matched.push(c);
                    }
                    Some(c) if c.is_whitespace() => {
                        // Eg: "m^2 "
                        self.state = State::SpaceAfterUnit;
                    }
                    Some(c) if c.is_alphabetic() => {
                        // Eg: "m^2kg"
                        // We don't support implicit multiplication here, so we stop the matching process
                        self.state = State::End;
                    }
                    _ => {
                        self.matched.push_str(&self.current_word);
                        self.current_word.clear();

                        self.state = State::End;
                    }
                }
            }
            State::SpaceAfterUnit => {
                match self.peek() {
                    Some(c) if c.is_whitespace() => {
                        self.consume();
                    }
                    Some(c) if c.is_alphabetic() => {
                        // Eg: "m kg" (implicit multiplication). Stop the matching process
                        self.state = State::End;
                    }
                    Some('*') | Some('/') => {
                        self.current_word.clear();
                        self.state = State::Operator;
                    }
                    _ => {
                        self.state = State::End;
                    }
                }
            }
            State::End => {}
        }
    }

    pub fn run(&mut self) -> String {
        while self.state != State::End {
            self.step();
        }

        self.matched.clone()
    }

    pub fn consume(&mut self) -> Option<char> {
        self.pos += 1;
        self.input.next()
    }

    pub fn peek(&self) -> Option<char> {
        self.input.clone().next()
    }

    pub fn skip(&mut self, n: usize) {
        self.pos += n;
        self.input.by_ref().take(n).for_each(drop);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn run(input: &str) -> String {
        UnitTokenizer::new(input.chars(), vec!["sin", "pi"]).run()
    }

    #[test]
    fn test_unit_tokenizer() {
        assert_eq!(
            run("kg"),
            "kg"
        );
    }

    #[test]
    fn test_unit_tokenizer_with_whitespace() {
        assert_eq!(
            run("kg "),
            "kg"
        );
    }

    #[test]
    fn test_unit_tokenizer_with_number() {
        assert_eq!(
            run("m2"),
            "m^2"
        );
    }

    #[test]
    fn test_unit_tokenizer_with_power() {
        assert_eq!(
            run("m^2"),
            "m^2"
        );

        assert_eq!(
            run("m^-1"),
            "m^-1"
        );
    }

    #[test]
    fn test_unit_tokenizer_with_division() {
        assert_eq!(
            run("m/s"),
            "m/s"
        );
    }

    #[test]
    fn test_unit_tokenizer_with_parentheses() {
        assert_eq!(
            run("m/(s/(kg))"),
            "m/(s/(kg))"
        );
    }

    #[test]
    fn test_unit_tokenizer_with_parentheses_power_and_spaces() {
        assert_eq!(
            run("m^2 /( s * kg ) "),
            "m^2/(s*kg)"
        );
    }

    #[test]
    fn test_unit_followed_by_standard_multiplication() {
        assert_eq!(
            run("kg * 3"),
            "kg"
        );

        assert_eq!(
            run("kg*3"),
            "kg"
        );
    }

    #[test]
    fn test_unit_followed_by_standard_division() {
        assert_eq!(
            run("kg / 3"),
            "kg"
        );

        assert_eq!(
            run("kg/3"),
            "kg"
        );
    }

    #[test]
    fn test_unit_followed_by_multiplication_constant() {
        assert_eq!(
            run("kg * pi"),
            "kg"
        );
        
        assert_eq!(
            run("kg*pi"),
            "kg"
        );
    }

    #[test]
    fn test_unit_followed_by_division_constant() {
        assert_eq!(
            run("kg / pi"),
            "kg"
        );
    }

    #[test]
    fn test_unit_followed_by_multiplication_function() {
        assert_eq!(
            run("kg * sin(3)"),
            "kg"
        );
    }

    #[test]
    fn test_unit_followed_by_division_function() {
        assert_eq!(
            run("kg / sin(3)"),
            "kg"
        );
    }

    #[test]
    fn test_unit_starting_with_division() {
        assert_eq!(
            run("/kg + 2"),
            "/kg"
        );
    }

}