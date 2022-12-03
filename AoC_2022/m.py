
i = 6

while i <= 25:
    path = f"day{i}/main.go"
    contents = """package main
    
import \"fmt\"
    
func main() {
    fmt.Println(\"Welcome to day {i}\")
}\n"""

    with open(path, "w") as f:
        f.write(contents)
    txtpath = f"day{i}/input.txt"
    with open(txtpath, "w") as f:
        pass
    i+=1