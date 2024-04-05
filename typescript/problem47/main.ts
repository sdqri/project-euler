function factor(x: number): number[] {
    const factors: number[] = []
    let i = 2
    while (x != 1) {
        if (x % i == 0) {
            factors.push(i)
            x /= i
        } else {
            i++
        }
    }
    return factors
}

function getFactorsSet(x: number): Set<number> {
    const factorsSet: Set<number> = new Set()
    const factors = factor(x)
    for (let i = 0; i < factors.length; i++) {
        factorsSet.add(factors[i])
    }

    return factorsSet
}

const start = performance.now()
let i = 1, j = 2, k = 3, l = 4;
let factorsSet_i: Set<number>
let factorsSet_j: Set<number>
let factorsSet_k: Set<number>
let factorsSet_l: Set<number>
let union: Set<number>
do {
    i++
    j++
    k++
    l++
    factorsSet_i = getFactorsSet(i)
    factorsSet_j = getFactorsSet(j)
    factorsSet_k = getFactorsSet(k)
    factorsSet_l = getFactorsSet(l)
   } while (
    !(factorsSet_i.size == 4 &&
        factorsSet_j.size == 4 &&
        factorsSet_k.size == 4 &&
        factorsSet_l.size == 4)
)

let end = performance.now()
let elapsed = end - start

console.log(`i=${i} (elapsed=${elapsed}ms)`)

