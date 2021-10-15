% porodicno stablo 

% svojstva
musko(mihajlo).
musko(stevan).
musko(petar).
musko(mladen).
musko(rajko).
zensko(milena).
zensko(milica).
zensko(jelena).
zensko(senka).
zensko(mina).
zensko(maja).

% odnosi
roditelj(mihajlo,milica).
roditelj(mihajlo,senka).
roditelj(milena,rajko).
roditelj(maja,petar).
roditelj(maja,mina).
roditelj(stevan,mladen).
roditelj(stevan,jelena).
roditelj(milica,mladen).
roditelj(milica,jelena).

majka(X, Y) :- zensko(X), roditelj(X, Y).
otac(X, Y) :- musko(X), roditelj(X, Y).
brat(X, Y) :- musko(X), roditelj(Z, X), roditelj(Z, Y), X \= Y.
sestra(X, Y) :- zensko(X), roditelj(Z, X), roditelj(Z, Y), X \= Y.
ujak(X, Y) :- majka(Z, Y), brat(X, Z).
tetka(X, Y) :- sestra(X, Z), otac(Z, Y).

predak(X, Y) :- roditelj(X, Y).
predak(X, Y) :- roditelj(X, Z), predak(Z, Y).