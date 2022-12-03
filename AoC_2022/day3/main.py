
print("Part 1 : ", sum([ord(min(n))-96 if min(n).islower() else ord(min(n))-38 for n in [set(l[:len(l)//2]).intersection(l[len(l)//2:]) for l in [n.rstrip() for n in open("input.txt").readlines()]]]))
print("Part 2 : ", sum([ord(min(n))-96 if min(n).islower() else ord(min(n))-38 for n in [''.join(set([n.rstrip() for n in open("input.txt").readlines()][i]).intersection([n.rstrip() for n in open("input.txt").readlines()][i+1]).intersection([n.rstrip() for n in open("input.txt").readlines()][i+2])) for i in range(0,len([n.rstrip() for n in open("input.txt").readlines()]), 3)]]))
