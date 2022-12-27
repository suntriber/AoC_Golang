import numpy as np
# raw_data = open(0).read()

lines = [n.rstrip() for n in open("test.txt").readlines()]

def v2(x, y):
    return np.array((x, y))

lengths = [v2(0, 0) for _ in range(10)]
head = lengths[0]
tail = lengths[-1]
visited = set()
visited2 = set()
# print(lengths)
# print(head, tail)
for line in lines:
    d, a = line.split()
    a = int(a)
    if d=='D':
        d = v2(0, 1)
    elif d=='U':
        d = v2(0, -1)
    elif d == 'L':
        d = v2(-1, 0)
    else:
        d = v2(1, 0)
    for _ in range(a):
        head += d
        # print("head : ", head)
        for i in range(9):
            front, back = lengths[i], lengths[i+1]
            delta = front - back
            # print("f, b, d : ", front, back, delta)
            if abs(delta).max() > 1:
                # print("back before clip: ",back)
                # print("delta is ", delta)
                # print("delta.clip is ", delta.clip(-1, 1))
                back += delta.clip(-1, 1)
                # print("back after clip: ",back)
        # print(lengths[1])
        # print(tail)
        print(lengths)
        visited.add(tuple(lengths[1]))
        visited2.add(tuple(tail))
        # print("visited: ", visited)
print(len(visited),len(visited2))