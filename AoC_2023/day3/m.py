# p1
from itertools import product
lines = open("input.txt", "r").readlines()
p1 = 0
seen = set()
for y, line in enumerate(lines):
    for x, n in enumerate(line):
        if not n.isdigit():
            continue

        for dx, dy in product(range(-1, 2), repeat=2):
            if dx == dy == 0:
                continue

            nx, ny = x+dx, y+dy
            if nx < 0 or nx >= len(line) or ny < 0 or ny >= len(lines):
                continue

            if lines[ny][nx] != '.' and not lines[ny][nx].isdigit():
                break
        else:
            continue

        while x >= 0 and line[x].isdigit():
            x -= 1

        x += 1

        if (x, y) in seen:
            continue

        seen.add((x, y))

        c = 0
        while x < len(line) and line[x].isdigit():
            c = c*10 + int(line[x])
            x += 1

        print(c)

        p1 += c

print(seen)
print(p1)

# p2
from itertools import product
from re import finditer

from collections import defaultdict
ratios = defaultdict(list)
for y, line in enumerate(lines):
    for m in finditer("\d+", line):
        number = int(m.group())
        gears = set()
        for x in range(*m.span()):
            for dx, dy in product(range(-1, 2), repeat=2):
                if dx == dy == 0:
                    continue

                nx, ny = x+dx, y+dy
                if nx < 0 or nx >= len(line) or ny < 0 or ny >= len(lines):
                    continue

                if (nx, ny) in gears:
                    continue

                if lines[ny][nx] == '*':
                    gears.add((nx, ny))
                    ratios[(nx, ny)].append(number)

from math import prod
p1 = 0
for a, b in ratios.items():
    if len(b) == 2:
        p1 += prod(b)

print(p1)