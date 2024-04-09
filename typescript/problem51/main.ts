function sieveOfEratosthenes(n: number): number[] {
    const root = Math.sqrt(n)
    const notPrimes: boolean[] = []
    for (let i = 0; i < n; i++) {
        notPrimes.push(false)
    }

    notPrimes[0] = true
    for (let i = 2; i <= root; i++) {
        let factor = 2 * i
        while (factor <= n) {
            notPrimes[factor - 1] = true
            factor += i
        }
    }

    const primes: number[] = []
    for (let i = 0; i < notPrimes.length; i++) {
        if (!notPrimes[i]) {
            primes.push(i + 1)
        }
    }

    return primes
}

function generatePattern(num: number): string[] {
    const patterns: string[] = []
    const digits: Map<string, number> = new Map()
    const numStr = Math.abs(num).toString()
    for (const char of numStr) {
        if (digits.has(char)) {
            digits.set(char, digits.get(char)! + 1)
        } else {
            digits.set(char, 1)
        }
    }

    digits.forEach((value, key) => {
        for (let i = 1; i < (1 << value); i++) {
            let pattern = ""
            let bitIndex = 0
            for (let j = 0; j < numStr.length; j++) {
                if (numStr[j] != key) {
                    pattern = pattern + numStr[j]
                } else {
                    if ((i & (1 << bitIndex)) != 0) {
                        pattern = pattern + "*"
                    } else {
                        pattern = pattern + numStr[j]
                    }
                    bitIndex++
                }
            }
            patterns.push(pattern)
        }
    })

    return patterns
}

interface PatternValue {
    Count: number
    Creator: number
}

const n = 1_000_000
const primes = sieveOfEratosthenes(n)
const primesSet: Set<number> = new Set()
for (let i = 0; i < n; i++) {
    primesSet.add(primes[i])
}

const start = performance.now()

const families = new Map<string, PatternValue>();
for (let i = 1; i < primes.length; i++) {
    const patterns = generatePattern(primes[i])
    for (const pattern of patterns) {
        if (families.has(pattern)) {
            const patternValue = families.get(pattern)!
            patternValue.Count++
            families.set(pattern, patternValue)
        } else {
            const patternValue = {
                Count: 1,
                Creator: primes[i]
            } as PatternValue
            families.set(pattern, patternValue)
        }

    }
}

interface MinKeyValue {
    Key: string
    Value: PatternValue
}

let min8ValueFamily: MinKeyValue = { Key: "", Value: { Count: 0, Creator: Infinity } }
families.forEach((value, key) => {
    if (value.Count == 8) {
        if (value.Count < min8ValueFamily.Value.Creator) {
            min8ValueFamily.Key = key
            min8ValueFamily.Value = value
        }
    }
})

const end = performance.now()
const elapsed = end - start

console.log("minimum 8 value family = ", min8ValueFamily)
console.log(`(elapsed=${elapsed}ms)`)
