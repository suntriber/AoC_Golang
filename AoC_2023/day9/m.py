import numpy as np
data = [np.fromstring(l,int,sep=' ') for l in open('input.txt')]

ret = 0
for d in data:
    stack = []
    while (d!=0).any():
        stack.append(d[0])
        d = np.diff(d)
    d = np.append(d, 0)
    while stack:
        p = stack.pop()
        d = np.cumsum(np.concatenate(([p], d)))
    ret += d[-1]

print(ret)

f = open('input.txt')
def extrapolate(nums):
    if len(set(nums)) == 1: return nums[0]
    return nums[0] - extrapolate([nums[i+1] - nums[i] for i in range(len(nums)-1)])
s = 0
while k := f.readline().strip('\n'):
    num = list(map(int, k.split()))
    r = extrapolate(num)
    s += r

print(s)