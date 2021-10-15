/*jeVeci(a,b) <=> a je veci od b*/
zivotinja(slon).
zivotinja(zebra).
zivotinja(pas).
zivotinja(macka).

jeVeci(slon, macka).
jeVeci(slon, pas).
jeVeci(slon, zebra).
jeVeci(pas, macka).
jeVeci(zebra, pas).
jeVeci(zebra, macka).

veci(X, Y) :- jeVeci(X, Z), jeVeci(Z, Y).