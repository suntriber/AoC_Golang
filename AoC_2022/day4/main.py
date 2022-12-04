# lines = open("input.txt").read().splitlines()

# s = 0
# for line in lines:
#     l1, l2 = line.split(",")
#     (min1, max1), (min2, max2) = l1.split("-"), l2.split("-")
#     min1, max1, min2, max2 = list(map(int, [min1, max1, min2, max2]))
#     # print(l1, l2)
#     # print(min1, max1, min2, max2)
#     if (min1 <= min2 <= max1) or (min2 <= min1 <= max2):
#         s += 1


# print(s)

import re

data = open("input.txt").read()

lines = data.splitlines()

part1 = 0
part2 = 0
for line in lines:
    a, b, c, d = list(int(i) for i in re.findall(r'\d+', line))
    e, f = set(range(a, b+1)), set(range(c, d+1))
    if e <= f or f <= e:
        part1 += 1
    if e & f:
        part2 += 1

print(part1, part2)


print(sum(1 for string in open("input.txt", "r").read().split('\n') if ((int(string.split(',')[0].split('-')[0]) >= int(string.split(',')[1].split('-')[0]) and int(string.split(',')[0].split('-')[1]) <= int(string.split(',')[1].split('-')[1])) or (int(string.split(',')[1].split('-')[0]) >= int(string.split(',')[0].split('-')[0]) and int(string.split(',')[1].split('-')[1]) <= int(string.split(',')[0].split('-')[1])))))