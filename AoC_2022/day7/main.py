d = [n.rstrip() for n in open("input.txt").readlines()]

fs = {}
dirs = []
for i in d:
    if i[0] == '$':
        cmd = i.split()
        if cmd[1] == 'cd':
            if cmd[2] == '/':
                dirs = [fs]
            elif cmd[2] == '..':
                dirs.pop()
            else:
                if cmd[2] not in dirs[-1]:
                    dirs[-1][cmd[2]] = {}
                dirs.append(dirs[-1][cmd[2]])
    else:
        size, fn = i.split()

        if size == 'dir':
            size = {}
        else:
            size = int(size)

        dirs[-1][fn] = size

def size(fs):
    total = 0
    for f, s in fs.items():
        if type(s) == dict:
            total += size(s)
        else:
            total += s
    return total

def findsizes(fs, min, max):
    sizes = []
    for f, s in fs.items():
        if type(s) == dict:
            sizes.extend(findsizes(s, min, max))
            total = size(s)
            if total <= min or total >= max:
                continue
            sizes.append(total)
    return sizes

dirs = findsizes(fs, 0, 100000)

used = size(fs)
free = 70000000 - used
need = 30000000 - free
delete = findsizes(fs, need, used)

print(sum(dirs), min(delete))