import re 

storage = [] 

def get_value(parameters):
    sides = re.findall(r"\d+", parameters)
    return sides

with open("input.txt") as file:
    for line in file:
        storage.append(get_value(line.strip()))
        

        
def get_paper_area(sideArr):
    totalArea = 0
    
    a,b, c = sideArr
    face1 = int(a)*int(b)
    face2 = int(b)*int(c)
    face3 = int(c)*int(a)
    
    slack = min(face1, face2, face3)
    totalArea = 2*(face1 + face2 + face3) + slack
    return totalArea
    
def get_Grand(grandArr):
    paper_area = 0
    for i in grandArr:
        paper_area += get_paper_area(i)
    return paper_area


def get_ribbon(parameters):
    a, b, c = parameters
    side1 = 2*int(a) + 2*int(b)
    side2 = 2*int(b) + 2*int(c)
    side3 = 2*int(a) + 2*int(c)
    smallest = min(side1, side2, side3)
    vol = int(a)*int(b)*int(c)
    ribbon_length = vol + smallest
    return ribbon_length

def get_total_ribbon(grandArr):
    lenght = 0
    for i in grandArr:
        lenght += get_ribbon(i)
    return lenght

print(get_total_ribbon(storage))

    
        