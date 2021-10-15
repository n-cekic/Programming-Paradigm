podeli([], [], []):- !.
podeli([G|R], [G|LP1], LN1) :- G >= 0, podeli(R, LP1, LN1).
podeli([G|R], LP1, [G|LN1]) :- G < 0, podeli(R, LP1, LN1).


dupliraj([], []) :- !.
dupliraj([G|R], [G|NL1]) :- G >= 0, dupliraj(R, NL1).
dupliraj([G|R], [G,G|NL1]) :- G < 0, dupliraj(R, NL1).

zameni(_, _, [], []) :- !.
zameni(X, Y, [G|R], [Y|NL1]) :- G == X, zameni(X, Y, R, NL1), !.
zameni(X, Y, [G|R], [X|NL1]) :- G == Y, zameni(X, Y, R, NL1), !.
zameni(X, Y, [G|R], [G|NL1]) :- zameni(X, Y, R, NL1). 


uBrojNazad([], 0) :- !.
uBrojNazad([G|R], B) :- uBroj(R, B1), B is 10 * B1 + G.

poslednji([X], X, []) :- !.
poslednji([G|R], X, [G|NL1]) :- poslednji(R, X, NL1).

pretvori([], 0):- !.
pretvori(L, X) :- poslednji(L, P, NL), pretvori(NL, X1), X is P + 10 * X1.


veci(X, Y, X) :- X >= Y, !.
veci(X, Y, Y) :- X < Y, !.

maxEl([X], X) :- !. 
maxEl([G|R], G) :- maxEl(R, M1), G >= M1, !.
maxEl([G|R], M1) :- maxEl(R, M1), G < M1, !.