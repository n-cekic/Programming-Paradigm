# U fajlu podaci.json su dati podaci o jednoj grupi dece u obliku liste listi. Za
# svako dete su zasebnoj listi dati ime, pol, starost (u godinama), masa (u kilogramima) i visina (u
# centimetrima) redom. Učitati podatke iz fajla u tabelu i dodati odgovarajuće nazive kolonama.

import json
import pandas
from matplotlib import pyplot as plt

f = open('podaci.json', 'r')
podaci = json.load(f)

# pravimo tabelu od liste listi
tabela = pandas.DataFrame(podaci)
# dodajemo imena kolonama (podrazumevano ce imena biti indeksi od 0 do broj kolona - 1)
tabela.columns = ['ime','pol','starost','masa','visina']
print(tabela)

# izdvajamo odgovarajucu kolonu (seriju) iz tabele i racunamo 
# redom minimum, maksimum i prosek
print(tabela['visina'].min())
print(tabela['starost'].max())
print(tabela['visina'].mean())

# sabiramo sve elemente kolone sa podacima o masi 
if tabela['masa'].sum() <= 600:
    print("Mogu svi da stanu u lift.")
else:
    print("Ne. Zajedno su preteski.")

# definisemo velicinu grafika
plt.figure(figsize = (10,5))
# x osa - imena, y osa - podaci o visini
plt.bar(tabela['ime'],tabela['visina'])
plt.title('Visina dece u grupi')
plt.show()
