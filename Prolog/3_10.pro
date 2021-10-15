sadrzi(X, [X|_]) :- !.
sadrzi(X, [G|R]) :- G \== X, sadrzi(X, R).


duzina([_], 1) :- !.
duzina([_|R], D) :- duzina(R, D1), D is D1 + 1.


suma([], 0) :- !.
suma([G|R], S) :- suma(R, S1), S is G + S1.


arsr([], 0):- !.
arsr(L, S) :- suma(L, S1), duzina(L, D1), D1 =\= 0, S is S1 / D1, !.