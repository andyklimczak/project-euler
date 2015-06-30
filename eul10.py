def isPrime(i):
    primeTest = True
    if i != 2:
        for x in range(2,i):
            if i % x == 0:
                primeTest = False
    return primeTest

sum = 0
n = 2000001
for i in range(0, n):
    if isPrime(i):
        sum = i + sum
    print(i) if i % 10000 == 0 else None
print("sum is: ", sum)