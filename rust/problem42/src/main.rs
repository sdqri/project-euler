use std::fs;
use std::time::Instant;

fn is_valid_triangle_number(x: i64) -> bool {
    // n^2 + n - 2x = 0
    // discriminant = b^2 - 4ac = 1 - 4 * 1 * -2x
    let discriminant = 1 + 8 * x;
    let root = (discriminant as f64).sqrt() as i64;
    root * root == discriminant
}

fn get_word_value(s: &str) -> i64{
   s.bytes().map(|b| i64::from(b) - 64).sum()
}

fn read_words(path: &str) -> Vec<String> {
    let contents = fs::read_to_string(path)
        .expect("Failed to read file contents");

    let words: Vec<String> = contents.split(',')
        .map(|word| {
            let word = word.trim();
            if word.starts_with('"') && word.ends_with('"') {
                word[1..word.len() - 1].to_string()
            } else {
                word.to_string()
            }
        })
        .collect();

    words
}


fn main(){
    let start = Instant::now();
    let words = read_words("./data/0042_words.txt");
    let triangle_words_count = words.iter().filter(|w| is_valid_triangle_number(get_word_value(w))).count();
    let stop = Instant::now();
    let elapsed = stop - start;
    println!("triangel words count = {} (elapsed={:?})", triangle_words_count, elapsed);
}
