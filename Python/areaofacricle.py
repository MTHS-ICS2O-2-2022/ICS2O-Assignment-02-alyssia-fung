import math

function def calculate():
    radius = float(input("Enter the radius of the circle: "))

    area = math.pi * radius ** 2

    print("The area of the circle is: ", str(area) + " cm²")

calculate()
