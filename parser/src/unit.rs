/// A unit is a collection its factors and their exponents.
/// For example, the unit "m/s^2" would be represented as
/// [("m", 1), ("s", -2)].
#[derive(Clone, Debug)]
pub struct Unit {
    pub parts: Vec<(String, isize)>,
}

impl Unit {
    /// Create a new unit from a list of factors and exponents.
    fn new(parts: Vec<(String, isize)>) -> Unit {
        Unit { parts }
    }

    pub fn empty() -> Unit {
        Unit { parts: vec![] }
    }

    pub fn from_str(s: &str) -> Unit {
        let parts = crate::parser::units::ExprParser::new().parse(s).unwrap().flatten();

        Unit::new(parts).normalize()
    }

    pub fn blank(&self) -> bool {
        self.parts.is_empty()
    }

    /// Multiply two units together.
    pub fn multiply(&self, other: &Unit) -> Unit {
        let mut parts = self.parts.clone();
        for (factor, exponent) in &other.parts {
            parts.push((factor.clone(), *exponent));
        }
        Unit::new(parts).normalize()
    }

    /// Divide two units.
    pub fn divide(&self, other: &Unit) -> Unit {
        let mut parts = self.parts.clone();
        for (factor, exponent) in &other.parts {
            parts.push((factor.clone(), -*exponent));
        }
        Unit::new(parts).normalize()
    }

    /// Raise a unit to a power.
    pub fn pow(&self, power: i32) -> Unit {
        let parts = self
            .parts
            .iter()
            .map(|(factor, exponent)| (factor.clone(), (*exponent as i32 * power) as isize))
            .collect();
        Unit::new(parts).normalize()
    }

    /// Clean up a unit by combining like terms.
    fn normalize(&self) -> Unit {
        if self.parts.is_empty() {
            return self.clone();
        }

        let mut parts = self.parts.clone();
        parts.sort_by(|a, b| a.0.cmp(&b.0));
        let mut i = 0;
        while i < parts.len() - 1 {
            if parts[i].0 == parts[i + 1].0 {
                parts[i].1 += parts[i + 1].1;
                parts.remove(i + 1);
            } else {
                i += 1;
            }
        }

        // Filter out any terms with a zero exponent.
        parts.retain(|(_, exponent)| *exponent != 0);

        Unit { parts }
    }

    /// Convert a unit to a string.
    fn to_string(&self) -> String {
        let positives = self
            .parts
            .iter()
            .filter(|(_, exponent)| *exponent > 0)
            .map(|(factor, exponent)| {
                if *exponent == 1 {
                    factor.clone()
                } else {
                    format!("{}^{}", factor, exponent)
                }
            })
            .collect::<Vec<String>>()
            .join(" ");
        let negatives = self
            .parts
            .iter()
            .filter(|(_, exponent)| *exponent < 0)
            .map(|(factor, exponent)| {
                if *exponent == -1 {
                    factor.clone()
                } else {
                    format!("{}^{}", factor, -exponent)
                }
            })
            .collect::<Vec<String>>()
            .join(" ");

        if positives.is_empty() {
            if negatives.is_empty() {
                "".to_string()
            } else {
                "/".to_string() + &negatives
            }
        } else if negatives.is_empty() {
            positives
        } else {
            positives + "/" + &negatives
        }
    }
}

impl PartialEq for Unit {
    /// Equality is determined by the normalized form of the unit.
    fn eq(&self, other: &Self) -> bool {
        self.normalize().parts == other.normalize().parts
    }
}

impl std::fmt::Display for Unit {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "{}", self.to_string())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_unit_multiply() {
        let unit1 = Unit::new(vec![("m".to_string(), 1), ("s".to_string(), -2)]);
        let unit2 = Unit::new(vec![("s".to_string(), 1), ("kg".to_string(), 1)]);
        let result = unit1.multiply(&unit2);

        assert_eq!(
            result,
            Unit::new(vec![
                ("m".to_string(), 1),
                ("s".to_string(), -1),
                ("kg".to_string(), 1),
            ])
        );
    }

    #[test]
    fn test_unit_divide() {
        let unit1 = Unit::new(vec![("m".to_string(), 1), ("s".to_string(), -2)]);
        let unit2 = Unit::new(vec![("s".to_string(), 1), ("kg".to_string(), 1)]);
        let result = unit1.divide(&unit2);

        assert_eq!(
            result,
            Unit::new(vec![
                ("m".to_string(), 1),
                ("s".to_string(), -3),
                ("kg".to_string(), -1),
            ])
        );
    }

    #[test]
    fn test_unit_pow() {
        let unit = Unit::new(vec![("m".to_string(), 1), ("s".to_string(), -2)]);
        let result = unit.pow(3);

        assert_eq!(
            result,
            Unit::new(vec![("m".to_string(), 3), ("s".to_string(), -6),])
        );
    }
}
