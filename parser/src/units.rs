
use std::{sync::RwLock};

struct Unit {
    /// The main representation of the unit. Eg: "m" for meters.
    name: &'static str,
    /// Some units can be considered case insensitive. Eg: "m" and "M" for meters.
    /// Not the case for mW and MW, or kB and KB.
    case_sensitive: bool,
    /// Alternative names for the unit. Eg: "meter", "metre", "meters", "metres"
    aliases: Vec<&'static str>,
    /// The conversion factor to the base unit. Eg: 1 for meters, 1000 for kilometers.
    conversion: f64,

    /// For compound units, the units that make up the compound unit.
    units: Vec<(&'static str, i32)>,
}

lazy_static! {
    static ref LENGTH_UNITS: RwLock<Vec<Unit>> = RwLock::new(vec![
        Unit {
            name: "m",
            case_sensitive: false,
            aliases: vec!["meter", "metre", "meters", "metres"],
            conversion: 1.0,
            units: vec![],
        },
        Unit {
            name: "km",
            case_sensitive: false,
            aliases: vec!["kilometer", "kilometre", "kilometers", "kilometres"],
            conversion: 1000.0,
            units: vec![],
        },
        Unit {
            name: "cm",
            case_sensitive: false,
            aliases: vec!["centimeter", "centimetre", "centimeters", "centimetres"],
            conversion: 0.01,
            units: vec![],
        },
        Unit {
            name: "mm",
            case_sensitive: false,
            aliases: vec!["millimeter", "millimetre", "millimeters", "millimetres"],
            conversion: 0.001,
            units: vec![],
        },
        Unit {
            name: "nm",
            case_sensitive: true, // "nm" is nanometers, "NM" is nautical miles.
            aliases: vec!["nanometer", "nanometre", "nanometers", "nanometres"],
            conversion: 1e-9,
            units: vec![],
        },
        Unit {
            name: "μm",
            case_sensitive: false,
            aliases: vec!["micrometer", "micrometre", "micrometers", "micrometres"],
            conversion: 1e-6,
            units: vec![],
        },
        Unit {
            name: "in",
            case_sensitive: false,
            aliases: vec!["inch", "inches"],
            conversion: 0.0254,
            units: vec![],
        },
        Unit {
            name: "ft",
            case_sensitive: false,
            aliases: vec!["foot", "feet"],
            conversion: 0.3048,
            units: vec![],
        },
        Unit {
            name: "yd",
            case_sensitive: false,
            aliases: vec!["yard", "yards"],
            conversion: 0.9144,
            units: vec![],
        },
        Unit {
            name: "mi",
            case_sensitive: false,
            aliases: vec!["mile", "miles"],
            conversion: 1609.344,
            units: vec![],
        },
        Unit {
            name: "au",
            case_sensitive: false,
            aliases: vec!["astronomical unit", "astronomical units"],
            conversion: 149597870700.0,
            units: vec![],
        },
        Unit {
            name: "ly",
            case_sensitive: false,
            aliases: vec!["light year", "light years"],
            conversion: 9460730472580800.0,
            units: vec![],
        },
        Unit {
            name: "pc",
            case_sensitive: false,
            aliases: vec!["parsec", "parsecs"],
            conversion: 30856775814671900.0,
            units: vec![],
        },
        Unit {
            name: "NM",
            case_sensitive: true,
            aliases: vec!["nmi", "nautical mile", "nautical miles"],
            conversion: 1852.0,
            units: vec![],
        }
    ]);

    static ref MASS_UNITS: RwLock<Vec<Unit>> = RwLock::new(vec![
        Unit {
            name: "g",
            case_sensitive: false,
            aliases: vec!["gram", "grams"],
            conversion: 0.001,
            units: vec![],
        },
        Unit {
            name: "kg",
            case_sensitive: false,
            aliases: vec!["kilogram", "kilograms"],
            conversion: 1.0,
            units: vec![],
        },
        Unit {
            name: "mg",
            case_sensitive: false,
            aliases: vec!["milligram", "milligrams"],
            conversion: 1e-6,
            units: vec![],
        },
        Unit {
            name: "μg",
            case_sensitive: false,
            aliases: vec!["microgram", "micrograms"],
            conversion: 1e-9,
            units: vec![],
        },
        Unit {
            name: "lb",
            case_sensitive: false,
            aliases: vec!["pound", "pounds"],
            conversion: 0.45359237,
            units: vec![],
        },
        Unit {
            name: "oz",
            case_sensitive: false,
            aliases: vec!["ounce", "ounces"],
            conversion: 0.028349523125,
            units: vec![],
        },
        Unit {
            name: "st",
            case_sensitive: false,
            aliases: vec!["stone", "stones"],
            conversion: 6.35029318,
            units: vec![],
        },
        Unit {
            name: "slug",
            case_sensitive: false,
            aliases: vec!["slugs"],
            conversion: 14.5939029,
            units: vec![],
        }
    ]);

    static ref TIME_UNITS: RwLock<Vec<Unit>> = RwLock::new(vec![
        Unit {
            name: "s",
            case_sensitive: false,
            aliases: vec!["second", "seconds"],
            conversion: 1.0,
            units: vec![],
        },
        Unit {
            name: "ms",
            case_sensitive: false,
            aliases: vec!["millisecond", "milliseconds"],
            conversion: 1e-3,
            units: vec![],
        },
        Unit {
            name: "μs",
            case_sensitive: false,
            aliases: vec!["microsecond", "microseconds"],
            conversion: 1e-6,
            units: vec![],
        },
        Unit {
            name: "ns",
            case_sensitive: false,
            aliases: vec!["nanosecond", "nanoseconds"],
            conversion: 1e-9,
            units: vec![],
        },
        Unit {
            name: "min",
            case_sensitive: false,
            aliases: vec!["minute", "minutes"],
            conversion: 60.0,
            units: vec![],
        },
        Unit {
            name: "h",
            case_sensitive: false,
            aliases: vec!["hour", "hours"],
            conversion: 3600.0,
            units: vec![],
        },
        Unit {
            name: "d",
            case_sensitive: false,
            aliases: vec!["day", "days"],
            conversion: 86400.0,
            units: vec![],
        },
        Unit {
            name: "wk",
            case_sensitive: false,
            aliases: vec!["week", "weeks"],
            conversion: 604800.0,
            units: vec![],
        },
        Unit {
            name: "yr",
            case_sensitive: false,
            aliases: vec!["year", "years"],
            conversion: 86400.0 * 365.0, // Does not account for leap years, but it's more useful for general purposes.
            units: vec![],
        }
    ]);

    static ref TEMPERATURE_UNITS: RwLock<Vec<Unit>> = RwLock::new(vec![
        Unit {
            name: "K",
            case_sensitive: false,
            aliases: vec!["kelvin"],
            conversion: 1.0,
            units: vec![],
        },
        Unit {
            name: "C",
            case_sensitive: false,
            aliases: vec!["celsius"],
            conversion: 1.0,
            units: vec![],
        },
        Unit {
            name: "F",
            case_sensitive: false,
            aliases: vec!["fahrenheit"],
            conversion: 5.0 / 9.0,
            units: vec![],
        },
    ]);

    static ref CURRENT_UNITS: RwLock<Vec<Unit>> = RwLock::new(vec![
        Unit {
            name: "A",
            case_sensitive: false,
            aliases: vec!["ampere", "amperes"],
            conversion: 1.0,
            units: vec![],
        },
        Unit {
            name: "mA",
            case_sensitive: false,
            aliases: vec!["milliampere", "milliamperes"],
            conversion: 1e-3,
            units: vec![],
        },
        Unit {
            name: "μA",
            case_sensitive: false,
            aliases: vec!["microampere", "microamperes"],
            conversion: 1e-6,
            units: vec![],
        },
    ]);

    static ref VOLUME_UNITS: RwLock<Vec<Unit>> = RwLock::new(vec![
        Unit {
            name: "m3",
            case_sensitive: false,
            aliases: vec!["cubic meter", "cubic meters", "metre", "metres"],
            conversion: 1.0,
            units: vec![],
        },
        Unit {
            name: "cm3",
            case_sensitive: false,
            aliases: vec!["cubic centimeter", "cubic centimeters", "centimetre", "centimetres"],
            conversion: 1e-6,
            units: vec![],
        },
        Unit {
            name: "mm3",
            case_sensitive: false,
            aliases: vec!["cubic millimeter", "cubic millimeters", "millimetre", "millimetres"],
            conversion: 1e-9,
            units: vec![],
        },
        Unit {
            name: "L",
            case_sensitive: false,
            aliases: vec!["liter", "litre", "liters", "litres"],
            conversion: 1e-3,
            units: vec![],
        },
        Unit {
            name: "mL",
            case_sensitive: false,
            aliases: vec!["milliliter", "millilitre", "milliliters", "millilitres"],
            conversion: 1e-6,
            units: vec![],
        },
        Unit {
            name: "gal",
            case_sensitive: false,
            aliases: vec!["gallon", "gallons"],
            conversion: 0.00378541,
            units: vec![],
        },
        Unit {
            name: "qt",
            case_sensitive: false,
            aliases: vec!["quart", "quarts"],
            conversion: 0.000946353,
            units: vec![],
        },
        Unit {
            name: "pt",
            case_sensitive: false,
            aliases: vec!["pint", "pints"],
            conversion: 0.000473176,
            units: vec![],
        },
        Unit {
            name: "cup",
            case_sensitive: false,
            aliases: vec!["cups"],
            conversion: 0.000236588,
            units: vec![],
        },
        Unit {
            name: "fl oz",
            case_sensitive: false,
            aliases: vec!["fluid ounce","fluid ounces"],
            conversion: 2.95735e-5,
            units: vec![],
        },
        Unit {
            name: "tbsp",
            case_sensitive: false,
            aliases: vec!["tablespoon", "tablespoons"],
            conversion: 1.47868e-5,
            units: vec![],
        },
        Unit {
            name: "tsp",
            case_sensitive: false,
            aliases: vec!["teaspoon", "teaspoons"],
            conversion: 4.92892e-6,
            units: vec![],
        },
    ]);

    static ref ANGLE_UNITS: RwLock<Vec<Unit>> = RwLock::new(vec![
        Unit {
            name: "rad",
            case_sensitive: false,
            aliases: vec!["radian", "radians"],
            conversion: 1.0,
            units: vec![],
        },
        Unit {
            name: "deg",
            case_sensitive: false,
            aliases: vec!["degree", "degrees", "°"],
            conversion: std::f64::consts::PI / 180.0,
            units: vec![],
        },
    ]);

    static ref FORCE_UNITS: RwLock<Vec<Unit>> = RwLock::new(vec![
        Unit {
            name: "N",
            case_sensitive: false,
            aliases: vec!["newton", "newtons"],
            conversion: 1.0,
            units: vec![("kg", 1), ("m", 1), ("s", -2)],
        },
        Unit {
            name: "kN", // kN is kilonewtons, kn is knots.
            case_sensitive: true,
            aliases: vec!["kilonewton", "kilonewtons"],
            conversion: 1e3,
            units: vec![("kg", 1), ("m", 1), ("s", -2)],
        },
        Unit {
            name: "lbf",
            case_sensitive: false,
            aliases: vec!["pound force", "pounds force"],
            conversion: 4.4482216152605,
            units: vec![("kg", 1), ("m", 1), ("s", -2)],
        },
        Unit {
            name: "dyn",
            case_sensitive: false,
            aliases: vec!["dyne", "dynes"],
            conversion: 1e-5,
            units: vec![("kg", 1), ("m", 1), ("s", -2)],
        },
        Unit {
            name: "kgf",
            case_sensitive: false,
            aliases: vec!["kilogram force", "kilograms force"],
            conversion: 9.80665,
            units: vec![("kg", 1), ("m", 1), ("s", -2)],
        }
    ]);

    static ref PRESSURE_UNITS: RwLock<Vec<Unit>> = RwLock::new(vec![
        Unit {
            name: "Pa",
            case_sensitive: false,
            aliases: vec!["pascal", "pascals"],
            conversion: 1.0,
            units: vec![("kg", 1), ("m", -1), ("s", -2)],
        },
        Unit {
            name: "kPa",
            case_sensitive: false,
            aliases: vec!["kilopascal", "kilopascals"],
            conversion: 1e3,
            units: vec![("kg", 1), ("m", -1), ("s", -2)],
        },
        Unit {
            name: "MPa",
            case_sensitive: false,
            aliases: vec!["megapascal", "megapascals"],
            conversion: 1e6,
            units: vec![("kg", 1), ("m", -1), ("s", -2)],
        },
        Unit {
            name: "GPa",
            case_sensitive: false,
            aliases: vec!["gigapascal", "gigapascals"],
            conversion: 1e9,
            units: vec![("kg", 1), ("m", -1), ("s", -2)],
        },
        Unit {
            name: "bar",
            case_sensitive: false,
            aliases: vec!["bars"],
            conversion: 1e5,
            units: vec![("kg", 1), ("m", -1), ("s", -2)],
        },
        Unit {
            name: "atm",
            case_sensitive: false,
            aliases: vec!["atmosphere", "atmospheres"],
            conversion: 101325.0,
            units: vec![("kg", 1), ("m", -1), ("s", -2)],
        },
        Unit {
            name: "psi",
            case_sensitive: false,
            aliases: vec!["pound per square inch", "pounds per square inch"],
            conversion: 6894.757293168361,
            units: vec![("kg", 1), ("m", -1), ("s", -2)],
        },
    ]);

    static ref WORK_UNITS: RwLock<Vec<Unit>> = RwLock::new(vec![
        Unit {
            name: "J",
            case_sensitive: false,
            aliases: vec!["joule", "joules"],
            conversion: 1.0,
            units: vec![("kg", 1), ("m", 2), ("s", -2)],
        },
        Unit {
            name: "kJ",
            case_sensitive: false,
            aliases: vec!["kilojoule", "kilojoules"],
            conversion: 1e3,
            units: vec![("kg", 1), ("m", 2), ("s", -2)],
        },
        Unit {
            name: "MJ",
            case_sensitive: false,
            aliases: vec!["megajoule", "megajoules"],
            conversion: 1e6,
            units: vec![("kg", 1), ("m", 2), ("s", -2)],
        },
        Unit {
            name: "GJ",
            case_sensitive: false,
            aliases: vec!["gigajoule", "gigajoules"],
            conversion: 1e9,
            units: vec![("kg", 1), ("m", 2), ("s", -2)],
        },
        Unit {
            name: "cal",
            case_sensitive: false,
            aliases: vec!["calorie", "calories"],
            conversion: 4.184,
            units: vec![("kg", 1), ("m", 2), ("s", -2)],
        },
        Unit {
            name: "kcal",
            case_sensitive: false,
            aliases: vec!["kilocalorie", "kilocalories"],
            conversion: 4184.0,
            units: vec![("kg", 1), ("m", 2), ("s", -2)],
        },
        Unit {
            name: "Wh",
            case_sensitive: false,
            aliases: vec!["watt hour", "watt hours"],
            conversion: 3600.0,
            units: vec![("kg", 1), ("m", 2), ("s", -2)],
        },
        Unit {
            name: "kWh",
            case_sensitive: false,
            aliases: vec!["kilowatt hour", "kilowatt hours"],
            conversion: 3.6e6,
            units: vec![("kg", 1), ("m", 2), ("s", -2)],
        },
        Unit {
            name: "MWh",
            case_sensitive: false,
            aliases: vec!["megawatt hour", "megawatt hours"],
            conversion: 3.6e9,
            units: vec![("kg", 1), ("m", 2), ("s", -2)],
        },
        Unit {
            name: "GWh",
            case_sensitive: false,
            aliases: vec!["gigawatt hour", "gigawatt hours"],
            conversion: 3.6e12,
            units: vec![("kg", 1), ("m", 2), ("s", -2)],
        },
        Unit {
            name: "BTU",
            case_sensitive: false,
            aliases: vec!["british thermal unit", "british thermal units"],
            conversion: 1055.05585262,
            units: vec![("kg", 1), ("m", 2), ("s", -2)],
        },
    ]);

    static ref POWER_UNITS: RwLock<Vec<Unit>> = RwLock::new(vec![
        Unit {
            name: "W",
            case_sensitive: false,
            aliases: vec!["watt", "watts"],
            conversion: 1.0,
            units: vec![("kg", 1), ("m", 2), ("s", -3)],
        },
        Unit {
            name: "kW",
            case_sensitive: false,
            aliases: vec!["kilowatt", "kilowatts"],
            conversion: 1e3,
            units: vec![("kg", 1), ("m", 2), ("s", -3)],
        },
        Unit {
            name: "MW",
            case_sensitive: false,
            aliases: vec!["megawatt", "megawatts"],
            conversion: 1e6,
            units: vec![("kg", 1), ("m", 2), ("s", -3)],
        },
        Unit {
            name: "GW",
            case_sensitive: false,
            aliases: vec!["gigawatt", "gigawatts"],
            conversion: 1e9,
            units: vec![("kg", 1), ("m", 2), ("s", -3)],
        },
        Unit {
            name: "hp",
            case_sensitive: false,
            aliases: vec!["horsepower", "horsepowers"],
            conversion: 745.69987158227022,
            units: vec![("kg", 1), ("m", 2), ("s", -3)],
        },
    ]);

    static ref DATA_UNITS: RwLock<Vec<Unit>> = RwLock::new(vec![
        Unit {
            name: "bit",
            case_sensitive: false,
            aliases: vec!["bits"],
            conversion: 1.0 / 8.0,
            units: vec![],
        },
        Unit {
            name: "byte",
            case_sensitive: false,
            aliases: vec!["bytes"],
            conversion: 1.0,
            units: vec![],
        },
        Unit {
            name: "kB",
            case_sensitive: false,
            aliases: vec!["kilobyte", "kilobytes"],
            conversion: 1e3,
            units: vec![],
        },
        Unit {
            name: "KiB",
            case_sensitive: false,
            aliases: vec!["kibibyte", "kibibytes"],
            conversion: 1024.0,
            units: vec![],
        },
        Unit {
            name: "MB",
            case_sensitive: false,
            aliases: vec!["megabyte", "megabytes"],
            conversion: 1e6,
            units: vec![],
        },
        Unit {
            name: "MiB",
            case_sensitive: false,
            aliases: vec!["mebibyte", "mebibytes"],
            conversion: 1024.0 * 1024.0,
            units: vec![],
        },
        Unit {
            name: "GB",
            case_sensitive: false,
            aliases: vec!["gigabyte", "gigabytes"],
            conversion: 1e9,
            units: vec![],
        },
        Unit {
            name: "GiB",
            case_sensitive: false,
            aliases: vec!["gibibyte", "gibibytes"],
            conversion: 1024.0 * 1024.0 * 1024.0,
            units: vec![],
        },
        Unit {
            name: "TB",
            case_sensitive: false,
            aliases: vec!["terabyte", "terabytes"],
            conversion: 1e12,
            units: vec![],
        },
        Unit {
            name: "TiB",
            case_sensitive: false,
            aliases: vec!["tebibyte", "tebibytes"],
            conversion: 1024.0 * 1024.0 * 1024.0 * 1024.0,
            units: vec![],
        },
    ]);

    static ref MISC_UNITS: RwLock<Vec<Unit>> = RwLock::new(vec![
        Unit {
            name: "Hz",
            case_sensitive: false,
            aliases: vec!["hertz"],
            conversion: 1.0,
            units: vec![
                ("s", -1),
            ],
        },

        Unit {
            name: "V",
            case_sensitive: false,
            aliases: vec!["volt", "volts"],
            conversion: 1.0,
            units: vec![
                ("kg", 1),
                ("m", 2),
                ("s", -3),
                ("A", -1),
            ],
        }
    ]);

    static ref ALL_UNITS: Vec<&'static RwLock<Vec<Unit>>> = vec![
        &LENGTH_UNITS,
        &MASS_UNITS,
        &TIME_UNITS,
        &TEMPERATURE_UNITS,
        &CURRENT_UNITS,
        &VOLUME_UNITS,
        &ANGLE_UNITS,
        &FORCE_UNITS,
        &PRESSURE_UNITS,
        &WORK_UNITS,
        &POWER_UNITS,
        &DATA_UNITS,
        &MISC_UNITS,
    ];
}
