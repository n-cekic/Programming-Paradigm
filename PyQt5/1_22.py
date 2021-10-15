# Napisati program kojim se iscrtava grafik logaritamske funkcije na intervalu koji
# zadaje korisnik

from PyQt5 import QtWidgets
import numpy
from matplotlib import pyplot

class Gui(QtWidgets.QWidget):
    def __init__(self):
        super().__init__()
        self.initGui()
    
    def initGui(self):
        self.setGeometry(200,200,250,250)
        self.setWindowTitle("Logaritamska funkcija")

        grid=QtWidgets.QGridLayout()
        self.setLayout(grid)
        grid.setSpacing(20)

        self.labela1=QtWidgets.QLabel("leva granica",self)
        grid.addWidget(self.labela1,0,0)

        self.text1=QtWidgets.QLineEdit(self)
        grid.addWidget(self.text1,0,1)

        self.labela2=QtWidgets.QLabel("desna granica",self)
        grid.addWidget(self.labela2,0,2)

        self.text2=QtWidgets.QLineEdit(self)
        grid.addWidget(self.text2,0,3)

        # dugme postavljamo da se prostire u dve kolone (poslednji argument za addWidget)
        # na taj nacin ce biti u centru 
        self.dugme=QtWidgets.QPushButton("GRAFIK")
        grid.addWidget(self.dugme,1,1,1,2)
        self.dugme.clicked.connect(self.grafik)

        self.show()
    
    def grafik(self):
        leva=float(self.text1.text())
        desna=float(self.text2.text())

        x=numpy.arange(leva,desna+1)
        y=numpy.log(x)

        pyplot.title("logaritamska funkcija")

        pyplot.plot(x,y,"g",linewidth=3,label="log")
        pyplot.grid(True,c="#dddddd")
        pyplot.legend()
        pyplot.show()

app=QtWidgets.QApplication([])
window=Gui()
app.exec_()