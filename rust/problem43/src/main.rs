use std::collections::HashMap;
use std::time::Instant;

fn is_pandigital(x: i64) -> bool{
    let mut digits_set: HashMap<char, ()> = HashMap::new();

    for c in x.to_string().chars() {
        digits_set.insert(c, ());
    }

    digits_set.len() == 10
}

fn has_divisibility_property(x: i64) -> bool{
    let x_str = x.to_string();
    if !is_pandigital(x) {return false}
    for (i, factor) in [2, 3, 5, 7, 11, 13, 17].iter().enumerate(){
        let sub_string = &x_str[i+1..i+4];
        let num: i64 = sub_string.parse().expect("couldn't convert str to num!");
        if num % factor != 0 {return false}
    }
    return true
}

fn get_permutations(digits: Vec<char>) -> Vec<String>{
    let mut permutations: Vec<String> = Vec::new();

    fn perms(current: Vec<char>, remaining: Vec<char>, permutations: &mut Vec<String>){
        if remaining.len() == 0 {
            let perm: String = current.iter().collect();
            permutations.push(perm);
        }

        for (index, &item) in remaining.iter().enumerate(){
            let mut copied_remaining = remaining.clone();
            copied_remaining.remove(index);
            let mut copied_current = current.clone();
            copied_current.push(item);
            perms(copied_current, copied_remaining, permutations); 
        }
    }

    perms(Vec::new(), digits, &mut permutations);
    return permutations
}

fn main() {
    let start = Instant::now();
    let digits = vec!['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'];
    let mut count = 0;
    let mut sum = 0;
    for p in get_permutations(digits){
       let num: i64 = p.parse().expect("cannot convert p to num");
       if has_divisibility_property(num){
           count += 1;
           sum += num;
       }
    }
    let stop = Instant::now();
    let elapsed = stop - start;
    println!("found {count} pandigital number that has divisibility property with sum = {sum} (elapsed={elapsed:?})")
}
