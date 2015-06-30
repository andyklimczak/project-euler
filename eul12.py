import sys

def factors(num):
    count = 2
    for i in range(1, (num+1) // 2):
        if num % i == 0:
            count = count + 1
    return count

for i in range(1,sys.maxsize):
    if factors(sum(range(i+1))) > 5:
        print("number: ", i)
        break