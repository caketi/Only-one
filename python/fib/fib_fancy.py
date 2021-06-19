def fib_fancy(n):
    prev1 = 1
    prev2 = 1
    for i in range(2, n+1):
        curr = prev1 + prev2
        prev2 = prev1
        prev1 = curr
    return curr
print(fib_fancy(22))