use crate::solution::sum_prime;

mod tests;
mod solution;

fn main() {
    println!("Sum of all the primes below two million = {}", sum_prime(2000000));
}
