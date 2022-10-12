pub fn sum_prime(n: i32) -> i64 {
    let mut x = 2;
    let mut prime_list: Vec<i64> = Vec::new();
    while x<n {
        if is_prime(x){
            prime_list.push(x as i64);
        }
        x+=1;
    }
    return prime_list.iter().sum();
}

pub fn is_prime(n: i32) -> bool {
    for i in 2..((n as f32).sqrt() as i32 + 1){
        if n%i==0{
            return false;
        }
    }
    return true;
}
