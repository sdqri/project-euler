#Multiples of 3 and 5
#Problem 1
'''
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
'''

m3 = list(range(0, 1000, 3)) # multiples of 3
m5 = list(range(0, 1000, 5)) # multiples of 5
m = list(set(m3 + m5))
ans = sum(m)

print(ans)




