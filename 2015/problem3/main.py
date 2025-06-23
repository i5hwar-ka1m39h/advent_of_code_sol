
    
def count_houses(direactions):
    x, y = 0, 0
    houses = set()
    houses.add((x, y))
    
    for dir in direactions:
        if dir == "^":
            y += 1
        elif dir == "v":
            y-=1
        elif dir == "<":
            x -= 1
        else:
            x +=1
        houses.add((x,y))
    
    return len(houses)    

with open("input.txt") as f:
    print(count_houses(f.read()))
    