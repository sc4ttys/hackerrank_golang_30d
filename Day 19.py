class AdvancedArithmetic(object):
    def divisorSum(n):
        raise NotImplementedError

class Calculator(AdvancedArithmetic):
    def divisorSum(self, n):
        divisors = []
        sum = 0
        for i in range (n+1):
            if i == 0:
                continue
            if (n) % (i) == 0:
                divisors.append(i)
        for d in divisors:
            sum += d
        return sum


n = int(input())
my_calculator = Calculator()
s = my_calculator.divisorSum(n)
print("I implemented: " + type(my_calculator).__bases__[0].__name__)
print(s)