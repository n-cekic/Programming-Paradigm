# U fajlu automobili.csv su dati podaci o prosečnoj potrošnji različitih modela
# automobila (model, potrošnja po gradu i autoputu redom po kolonama). Učitati tabelu iz fajla
# i prikazati sumarne informacije o tabeli. Dodati novu kolonu sa prosečnom potrošnjom za svaki
# model i po njoj sortirati tabelu opadajuće i nakon toga prikazati pet modela sa najvećom potrošnjom goriva. Dodatno, izvući iz tabele vrste u kojima je prosečna potrošnja između 5 i 7 litara i
# rezultat filtriranja sačuvati u tabeli ekonomicni.csv.

import pandas 

# funkcija read_csv ucitava fajl i vraca DataFrame objekat
df = pandas.read_csv('automobili.csv')

# head daje prvih pet redova rezultata
print(df.head())
print('-------------------------------------')
# tail daje poslednjih pet redova rezultata
# zadavanjem argumenta menja se podrazumevana vrednost za obe funkcije
print(df.tail(3))
print('-------------------------------------')

# ukoliko je potrebno koliko redova/kolona ima ucitan fajl
print(df.shape)

# detaljnije informacije o tipu podataka, koliko memorije zauzimaju itd.
print(df.info())

# dodavanje nove kolone - prosecne potrosnje 
df['prosecna potrosnja'] = (df['KML grad'] + df['KML autoput'])/2

# sortiranje (drugi argument je podrazumevano True)
# ascending = False ukoliko nam je potrebno opadajuce sortiranje 
df = df.sort_values(by='prosecna potrosnja', ascending=True)
print(df.head())

# filtriranje automobila po potrosnji goriva
# ako je naziv kolone iz vise reci, naziv stavljamo pod kose navodnike u upitu
f = df.query('5 <= `prosecna potrosnja` <= 7')

# eksportovanje rezultata u csv datoteku
f.to_csv('ekonomicni.csv')
