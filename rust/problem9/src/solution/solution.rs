pub fn find_special_pythagorean_triplet(sum: i32) -> (i32, i32, i32) {
    let r = (1..sum)
        .flat_map(|a| (1..sum).map(move |b| (a, b)))
        .map(|ab| (ab.0, ab.1, sum-(ab.0+ab.1)))
        .filter(|abc| abc.0<abc.1 && abc.1<abc.2)
        .filter(|abc|i32::pow(abc.0, 2)+i32::pow(abc.1, 2)==i32::pow(abc.2, 2))
        .nth(0).unwrap();
    return r;
}
