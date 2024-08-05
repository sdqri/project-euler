// Cubic Permutations
function getFingerprint(x: number): string{
    return x.toString().split("").sort().join("")
}

const start = performance.now()
let i = 0
let loop = true;
let result: number = 1;
let resultRoot: number = 1;
const myMap: Map<string, Array<number>> = new Map();
while (loop) {
    i += 1
    const num = i**3
    const fp = getFingerprint(num)

    let val = myMap.get(fp) ?? [];
    val?.push(num)
    if (val?.length == 5){
        result = val[0]
        loop = false
    }
    myMap.set(fp, val)
}

const end = performance.now()
const elapsed = end - start

console.log(`smallest cube with five cube permutations = ${result}(${Math.cbrt(result)}^3) (elapsed=${elapsed}ms)`)
