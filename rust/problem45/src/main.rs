use std::time::Instant;

fn get_triangle_number(n: i64) -> i64{
    n*(n+1)/2
}


fn is_triangle(t: i64) -> bool{
    let discriminant: i64 = 1 + 8*t;
    let root: i64 = (discriminant as f64).sqrt() as i64;
    root*root==discriminant && (1 + root) % 2 == 0
}

fn is_pentagonal(p: i64) -> bool{
    let discriminant: i64 = 1 + 24*p;
    let root: i64 = (discriminant as f64).sqrt() as i64;
    root*root==discriminant && (1 + root) % 6 == 0
}

fn is_hexagonal(h: i64) -> bool{
    let discriminant: i64 = 1 + 8*h;
    let root: i64 = (discriminant as f64).sqrt() as i64;
    root*root==discriminant && (1 + root) % 4 == 0
}

fn main() {
    let start = Instant::now();
    let mut i: i64 = 286;
    let mut n: i64;

    loop{
        n = get_triangle_number(i);
        if is_pentagonal(n) && is_hexagonal(n){
            break;
        }
        i += 1;
    }
    let stop = Instant::now();
    let elapsed = stop - start;
    println!("T({i})={n}=P(?)=H(?) (elapsed={elapsed:?})");
}
