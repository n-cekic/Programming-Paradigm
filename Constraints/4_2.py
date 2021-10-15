#! /usr/bin/python3

# Napisati program koji pronalazi petocifren broj ABCDE za koji je izraz A+2*B-3*C+4*D-5*E
# minimalan i A, B, C, D i E su različite cifre.

import constraint

problem = constraint.Problem()

problem.addVariables("BCDE", range(10))
problem.addVariables('A', range(1,10))
problem.addConstraint(constraint.AllDifferentConstraint())

resenja = problem.getSolutions()

min_izraz = 63
min_resenje = {}

for r in resenja:
    a = r['A']
    b = r['B']
    c = r['C']
    d = r['D']
    e = r['E']

    vrednost = a+2*b-3*c+4*d-5*e

    if(vrednost < min_izraz):
        min_izraz = vrednost
        min_resenje = r

print(min_resenje['A']*10000+min_resenje['B']*1000+
        min_resenje['C']*100+min_resenje['D']*10+
        min_resenje['E'])