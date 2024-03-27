function isPandigital(x: number, n: number): boolean {
    if (n < 1 || n > 9){
        throw new Error("Invalid value for n, It should be a number between 1 and 9.")
    }
    
    const xStr: string = x.toString()
    if (xStr.length != n){
        return false
    }

    for (let i=1;i<=n;i++){
        if (!xStr.includes(i.toString())) return false
    }

    return true; 
}

function isPrime(x: number): boolean {
    for (let i = 2; i<=Math.sqrt(x);i++){
        if (x%i==0) {
            return false
        }
    }
    return true
}

function arrayFromOneToN(n: number): number[] {
    const result: number[] = [];
    for (let i = 1; i <= n; i++) {
        result.push(i);
    }
    return result;
}

function arrayToNumber(arr: number[], base: number): number{
    let value: number = 0
    for (let i=arr.length-1; i>=0; i--){
        const position: number = (arr.length-1) - i 
        value += arr[i] * (base**position)
    }
    return value
}

function permute(digits: number[]): number[]{
    const permutations: number[][] = []

    function _permute(current: number[], remaining: number[]): void{
        if (remaining.length==0){
            permutations.push([...current])
            return
        }

        for (let i=0; i< remaining.length; i++){
            current.push(remaining[i])    
            const nextRemaining: number[] = [...remaining.slice(0, i), ...remaining.slice(i+1)]    
            _permute(current, nextRemaining)
            current.pop()
        }

        return
    }

    _permute([], digits)

    const results: number[] = []
    for (let i=0; i< permutations.length; i++){
        results.push(arrayToNumber(permutations[i], 10))
    }

    return results
}

let max: number = 1;

const startTime: number = performance.now()

for (let i=1; i<=9; i++){
    const digits: number[] = arrayFromOneToN(i)
    const numbers = permute(digits)
    for (let num of numbers){
        if (isPrime(num) && isPandigital(num, i)){
            if (num > max){
                max = num
            }
        } 
    }
}

const stopTime: number = performance.now()

const elapsed: number = stopTime - startTime

console.log(`biggest n digit prime & pandigital number = ${max} (elapsed = ${elapsed}ms)`)

