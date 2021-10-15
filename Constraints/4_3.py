#! /usr/bin/python3

# Dati su novčići od 1, 2, 5, 10, 20 dinara. Napisati program koji pronalazi sve
# moguće kombinacije tako da zbir svih novčića bude 50 i da se svaki novčič pojavljuje bar jednom
# u kombinaciji.

import constraint

problem = constraint.Problem()

problem.addVariable('A', range(1,51))
problem.addVariable('B', range(1,26))
problem.addVariable('C', range(1,11))
problem.addVariable('D', range(1,6))
problem.addVariable('E', range(1,3))

problem.addConstraint(constraint.ExactSumConstraint(50, [1,2,5,10,20]), 'ABCDE')

resenje = problem.getSolutions()

for r in resenje:
    print(r)

