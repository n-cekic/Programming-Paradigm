dodajPocetak(X, L, [X|L]).


dodajKraj(X, [], [X]) :- !.
dodajKraj(X, [G|R], [G|NL1]) :- dodajKraj(X, R, NL1).


obrisiPrvi([], []).
obrisiPrvi([_|R], R).


obrisiPoslednji([], []) :- !.
obrisiPoslednji([X,_], [X]) :- !.
obrisiPoslednji([G|R], [G|NL1]) :- obrisiPoslelednji(R, NL1).


izbaci(_, [], []) :- !.
izbaci(X, [G|R], NL1) :- G == X, izbaci(X, R, NL1).
izbaci(X, [G|R], [G|NL1]) :- G \== X, izbaci(X, R, NL1).


izbaci(_, [], []) :- !.
izbaciPrvo(X, [G|R], R):- G == X, !.
izbaciPrvo(X, [G|R], [G|N]) :- G \== X, izbaciPrvo(X, R, N).


obrisiKti([], _, []):- !.
obrisiKti([_|R], 1, R):- !.
obrisiKti([G|R], K, [G|NL1]) :- K1 is K - 1, obrisiKti(R, K1, NL1).