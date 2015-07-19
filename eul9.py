#comment
for c in range(0, 1000):
    for b in range(0, c):
        for a in range(0, b):
            if a**2 + b**2 == c**2 and a + b + c == 1000:
                print(a**2, ' + ', b**2, ' == ', c**2, ' and ', a, ' + ', b, c, ' == 1000 and the product is:', a*b*c)
                break
            elif a**2 + b**2 == c**2:
                print('a^2 + b^2 = c^2: ', a**2, ' + ', b**2, ' == ', c**2)