d = [int(n.rstrip()) if n.rstrip()!='' else 0 for n in open('part1.txt')]

tmp, _max = 0, 0
for n in d:
    if n == 0:
        if tmp > _max:
            _max = tmp
        tmp = 0
    tmp += n

print(_max)

new = []

tmp = 0
for n in d:
    tmp += n
    if n == 0:
        new.append(tmp)
        tmp = 0

print(sorted(new)[-1]+sorted(new)[-2]+sorted(new)[-3])
