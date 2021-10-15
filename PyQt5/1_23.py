# Napisati program kojim se iscrtava kružni grafik na osnovu rezultata ispita
# koji se učitavaju iz datoteke čiji naziv zadaje korisnik

from PyQt5 import QtWidgets
from matplotlib import pyplot as plt

class Gui(QtWidgets.QWidget):
    def __init__(self):
        super().__init__()
        self.initGui()
    
    def initGui(self):
        self.setGeometry(200,200,250,250)
        self.setWindowTitle("python")

        grid=QtWidgets.QGridLayout()
        self.setLayout(grid)
        grid.setSpacing(20)

        self.labela1=QtWidgets.QLabel("rezultati ispita",self)
        grid.addWidget(self.labela1,0,0)

        self.text1=QtWidgets.QLineEdit(self)
        grid.addWidget(self.text1,0,1)

        self.dugme=QtWidgets.QPushButton("DIJAGRAM")
        grid.addWidget(self.dugme,1,0)
        self.dugme.clicked.connect(self.dijagram)

        self.show()
    
    def dijagram(self):
        datoteka = self.text1.text()
        with open(datoteka,'r') as f:
            # procitamo sadrzaj
            # podelimo na niske 
            # svaku nisku kastujemo u broj
            # pretpostavka je da je ispravno zadata datoteka
            ocene = list(map(lambda x: int(x),f.read().split()))
        labele = ['Pali ispit', 'Polozili sa ocenom 6, 7 ili 8', 'Polozili sa ocenom 9 ili 10']
        vrednosti = []
        # prvi element je broj onih koji su pali ispit
        vrednosti.append(ocene[0])
        # naredna tri elementa - broj onih koji su dobili 6, 7 ili 8
        vrednosti.append(ocene[1]+ocene[2]+ocene[3])
        # poslednja dva elementa - broj onih koji su dobili 9 ili 10
        vrednosti.append(ocene[4]+ocene[5])
        boje = ['red', 'yellow', 'green']
        plt.title('Rezultati ispita')
        plt.pie(vrednosti, colors=boje, labels=labele)
        plt.show()
        
app=QtWidgets.QApplication([])
window=Gui()
app.exec_()