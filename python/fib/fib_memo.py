def fib_memo(n):
     a = [None] * (n+1)
     a[0] = 1
     a[1] = 1
     for i in range(2, n+1):
         a[i] = a[i-1] + a[i-2]
     return a[n]
print(fib_memo(500))