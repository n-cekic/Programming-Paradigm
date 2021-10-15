clan(X, [X|_]).
clan(X, [_|R]) :- clan(X, R).


pre(X,Y,[X,Y|_]).
pre(X,Y,[X|R]):- clan(Y,R).
pre(X,Y,[_|R]):- pre(X,Y,R).

parovi(L) :- L = [p(X,_,_,_), p(_,_,_,_), p(marko,_,_,macka), p(_,_,_,_)],
    clan(p(_,medved,_,_), L),
    X \== vasa,
    pre(p(X,_,_,_), p(_,princ,_,_), L),
    clan(p(_,princ,_,_), L),
    clan(p(pera, pajapatak, Z1, vestica), L),
    Z1 \== bojana,
    pre(p(laza,_,_,_), p(_,_,marija,_), L),
    pre(p(laza,_,_,_), p(_,_,bojana,_), L),
    pre(p(_,_,marija,_), p(_,_,bojana,_), L),
    pre(p(_,Z,_,ciganka), p(_,W,ana,_), L),
    Z \== betmen, W \== betmen,
    pre(p(_,_,ivana,_), p(_,_,_,snezana), L),
    clan(p(vasa,_,_,_), L),
    clan(p(_,betmen,_,_),L).