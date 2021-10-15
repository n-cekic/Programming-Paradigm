# U fajlu tacke.csv su dati podaci o tačkama: naziv, dve koordinate i boja redom.
# Učitati tabelu iz fajla i prikazati naredne informacije:
# a) za koliko tačaka su učitani podaci
# b) koliko tačaka ima prvu koordinatu manju od 30 i boju koja nije NaN (nepoznata vrednost)
# c) izvdojiti one tačke koje su crvene ili čije su obe koordinate manje od 10
# d) koliko ima tačaka obojenih određenom bojom i dati grafički prikaz koristeći kružni grafik uz
# koji treba ispisati nazive boja
# e) prikazati grafički prvih 15 tačaka na osnovu njihovih koordinata
# f) tačke koje su crvene ili plave sačuvati u poseban csv fajl

import pandas
from matplotlib import pyplot as plt

tacke = pandas.read_csv('tacke.csv')
# "duzina" tabele odgovara broju unosa, tj. redova, ne racunajuci zaglavlje 
print("CSV fajl sadrzi ", len(tacke), " unosa.")

# pravimo seriju buleanskih vrednosti 
# element serije je True ako je odgovarajuca x koordinata manja od 30, inace False
t1 = tacke['x'] < 30
print(t1)
# izdvajamo iz tabele one redove gde je u prethodnom rezultatu True
# zatim iz te podtabele izdvajamo kolonu "Color" i sa count prebrojavamo vrednosti koje nisu NaN
print(tacke[t1]['Color'].count())

crvene = tacke['Color'] == 'red'
xmanja = tacke['x'] < 10
ymanja = tacke['y'] < 10
# podtabela - redovi gde je boja crvena ili koordinate manje od 10
print(tacke[crvene | xmanja & ymanja])

# prebrojavamo koliko ima duplikata u koloni "Color"
# dobijamo seriju gde su indeksi konkretne boje, a vrednosti rezultat prebrojavanja te boje
vrednosti = tacke['Color'].value_counts()
print(vrednosti)

plt.title('Color of points')
# koristimo index iz serije (nazive boja) kao labele 
# ujedno ti nazivi mogu da se iskoriste i za boje kojim ce se obojiti odgovarajuci isecak kruga
plt.pie(vrednosti, labels = vrednosti.index, colors = vrednosti.index)
plt.show()

# izvlacimo samo kolone za koordinate i to prvih 15 redova
filter = tacke[['x','y']][:15]
print(filter)

plt.plot(filter['x'],filter['y'],'.')
plt.xlabel('x')
plt.ylabel('y')
plt.show()

zute = tacke['Color'] == 'yellow'
plave = tacke['Color'] == 'blue'
# izvlacimo one redove gde je vrednost u koloni 'Color' zuta ili plava
rezultat = tacke[zute | plave].to_csv('zute_plave.csv')




