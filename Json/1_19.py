# izracunati cenu spiska za kupovinu, u odnosu na razlicite prodavnice
import json

maxi_file = open('1_19_maxi_cene.json', 'r')
maxi_cene = json.load(maxi_file)
maxi_file.close()

idea_file = open('1_19_idea_cene.json', 'r')
idea_cene = json.load(idea_file)
idea_file.close()

shopngo_file = open('1_19_shopngo_cene.json', 'r')
shopngo_cene = json.load(shopngo_file)
shopngo_file.close()

korpa_file = open('1_19_korpa.json', 'r')
korpa_cene = json.load(korpa_file)
korpa_file.close()

maxi = 0
idea = 0
shopngo = 0

def cena(prodavnica, ime):
    for vocka in prodavnica:
        if (vocka['ime'] == ime):
            return vocka['cena']
    return

from functools import reduce


lista_cena = [vocka['kolicina'] * cena(maxi_cene, vocka['ime']) for vocka in korpa_cene]
print(reduce(lambda x, y: x + y, lista_cena))

lista_cena = [vocka['kolicina'] * cena(idea_cene, vocka['ime']) for vocka in korpa_cene]
print(reduce(lambda x, y: x + y, lista_cena))

lista_cena = [vocka['kolicina'] * cena(shopngo_cene, vocka['ime']) for vocka in korpa_cene]
print(reduce(lambda x, y: x + y, lista_cena))