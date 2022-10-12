use crate::solution::{sum_prime};

#[test]
fn test_sum_prime(){
    // 2, 3, 5, 7, 11 = 28
    let n: i32 = 5;
    assert_eq!(sum_prime(12), 28);
}