duzina [] = 0
duzina lista = 1 + duzina (tail lista)

duzina2 [] = 0
duzina2 (_:xs) = 1 + duzina2 xs

prvi (x:_) = x

ostatak (_:xs) = xs

parni a b = [x | x <- [a..b], mod x 2 == 0]

neparni a b = [x | x <- [a..b], not (mod x 2 == 0)]

parovi a b c d = [(x ,y) | x <- [a..b], y <- [c..d]]

parovi2 a b = [(x, y) | y <- [a..b], x <- [a..y]]

rep lista = if lista == [] then [] else tail lista

rep1 lista
    | lista == [] = []
    | otherwise = tail lista

rep2 [] = []
rep2 (x:xs) = xs

savrseni n = [x | x <- [1..n - 1], sum (faktori x) == x]
faktori x = [i | i <- [1..x - 1], (mod x i) == 0]

spoji listaListi = [e | lista <-listaListi, e <- lista]

sufiks [] = []
sufiks (x:xs) = (x:xs) : sufiks xs

izbaci _ [] = []
izbaci 0 (x:xs) = xs
izbaci p (x:xs) = x : (izbaci (p-1) xs)

ubaci e _ [] = [e]
ubaci e 0 lista = e : lista
ubaci e p (x:xs) = x : (ubaci e (p - 1) xs)