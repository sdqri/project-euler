use std::time::Instant;
use fraction;
use num_bigint::BigUint;
use num_traits::Zero;
use num_traits::One;

type F = fraction::BigFraction;

fn sqrt_series(length: u64) -> F {
    return F::from(BigUint::from(1u64)) + 
        (F::from(BigUint::from(1u64)) / sqrt2_series_helper(length))
}

fn sqrt2_series_helper(length: u64) -> F {
    if length == 0 {
        return F::from(BigUint::from(2u64))
    } else {
        return F::from(BigUint::from(2u64)) + 
            (F::from(BigUint::from(1u64)) / sqrt2_series_helper(length - 1))
    }
}

fn digit_length(x: &BigUint) -> u64 {
    x.to_string().chars().count() as u64
}

fn main() {
    let start = Instant::now();
    let mut num = 0;
    for l in 0..1000 {
        println!("{}", l);
        let expansion = sqrt_series(l);
        let numer_length = digit_length(expansion.numer().unwrap());
        let denom_length = digit_length(expansion.denom().unwrap());
        if numer_length > denom_length {
            num += 1;
        }
    }
    let elapsed = Instant::now() - start;
    println!(
        "Number of expansions with more digits in numerator than denominator = {} in 0..1000 (elapsed={elapsed:?})",
        num
    );
}

