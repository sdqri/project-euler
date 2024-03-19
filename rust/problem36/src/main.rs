use std::time::Instant;

fn is_double_base_palindrome(number: i64) -> bool{
    let number_str: String = number.to_string();
    let binary_number_str: String = format!("{:b}", number);
    let reversed_number_str: String = number_str.chars().rev().collect();
    let reversed_binary_number_str: String = binary_number_str.chars().rev().collect();
    number_str == reversed_number_str && binary_number_str == reversed_binary_number_str
}

fn main() {
    let start = Instant::now();
    let mut sum: i64 = 0;
    for i in 0..1_000_001{
        if is_double_base_palindrome(i){
            sum += i;
        }
    }
    let end = Instant::now();
    let elapsed = end.duration_since(start);

    println!("sum of double base palindrome numbers <= 1000000 = {sum} (elapsed = {elapsed:?})");
}
