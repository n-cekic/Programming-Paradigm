razlozi n = [(x, y) | x <-[1..n], y <-[1..div n x], x * y == n]

brojRastava n = length (razlozi n)

par _ [] = (0,0)
par n (x:xs) =
    if (fst x) == (n - 1) && (snd x) == (n + 1) then x
    else par n xs

podeli p lista = [podeliPocetak p lista, podeliKraj p lista]

podeliKraj 0 lista = lista
podeliKraj p (x:xs) = [a | a <- podeliKraj (p-1) xs] 

podeliPocetak 0 llista = []
podeliPocetak p (x:xs) = x : podeliPocetak (p-1) xs