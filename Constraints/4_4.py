#! /usr/bin/python3

# Napisati program koji ređa brojeve u magičan kvadrat. Magičan kvadrat
# je kvadrat dimenzija 3x3 takav da je suma svih brojeva u svakom redu, svakoj koloni i svakoj
# dijagonali jednak 15 i svi brojevi različiti.

import constraint

p = constraint.Problem()

p.addVariable('A1', range(1,10))
p.addVariable('A2', range(1,10))
p.addVariable('A3', range(1,10))
p.addVariable('A4', range(1,10))
p.addVariable('A5', range(1,10))
p.addVariable('A6', range(1,10))
p.addVariable('A7', range(1,10))
p.addVariable('A8', range(1,10))
p.addVariable('A9', range(1,10))

p.addConstraint(constraint.AllDifferentConstraint())
p.addConstraint(constraint.ExactSumConstraint(15), ["A1","A2","A3"])
p.addConstraint(constraint.ExactSumConstraint(15), ["A4","A5","A6"])
p.addConstraint(constraint.ExactSumConstraint(15), ["A7","A8","A9"])
p.addConstraint(constraint.ExactSumConstraint(15), ["A1","A4","A7"])
p.addConstraint(constraint.ExactSumConstraint(15), ["A5","A2","A8"])
p.addConstraint(constraint.ExactSumConstraint(15), ["A3","A6","A9"])
p.addConstraint(constraint.ExactSumConstraint(15), ["A1","A5","A9"])
p.addConstraint(constraint.ExactSumConstraint(15), ["A7","A5","A3"])

resenja = p.getSolutions()
for r in resenja:
    # formatiramo kvadrate
    print(" ------- ")
    print("| {0:d} {1:d} {2:d} |".format(r['A1'],r['A2'],r['A3']))
    print("| {0:d} {1:d} {2:d} |".format(r['A4'],r['A5'],r['A6']))
    print("| {0:d} {1:d} {2:d} |".format(r['A7'],r['A8'],r['A9']))
    print(" ------- ")

