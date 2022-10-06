use solution::find_special_pythagorean_triplet;

mod tests;
mod solution;

fn main() {
    let sum = 1000;
    let (a, b, c) = find_special_pythagorean_triplet(sum);
    println!("product of abc = {}", a*b*c)
}