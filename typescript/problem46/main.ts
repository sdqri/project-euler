function is_prime(x: number): boolean {
    if (x <= 1) {
        return false
    }
    const root: number = Math.sqrt(x)
    for (let i = 2; i <= root; i++) {
        if (x % i == 0) {
            return false
        }
    }
    return true
}

function is_valid_goldbachs_other_conjecture(x: number): boolean{
    for (let i=1;(i**2)*2<x;i++){
        let twiceasquare = (i**2)*2
        if (is_prime(x - twiceasquare)){
            return true
        }
    }
    return false
}

const start = performance.now()
let invalid_value = 0
let x = 3
while(true){
    if (!is_prime(x)) {
        if (!is_valid_goldbachs_other_conjecture(x)){
            invalid_value = x
            break
        }
    }
    x += 2
}
const end = performance.now()
const elapsed = end - start

console.log(`Smallest counterexample = ${invalid_value} (elapsed=${elapsed.toFixed(3)}ms)`)
