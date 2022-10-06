use more_asserts as ma;
use crate::solution::{find_special_pythagorean_triplet};

#[test]
fn test_find_special_pythagorean_triplet() {
    let (a, b, c) =  find_special_pythagorean_triplet(12);
    ma::assert_le!(a, b);
    ma::assert_le!(b, c);
    assert_eq!((a, b,c), (3, 4, 5));
}
