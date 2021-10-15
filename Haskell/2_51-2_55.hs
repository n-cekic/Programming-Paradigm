proveraListe p l = if p == "<0" then and (map (<0) l) else and (map (>0) l)

provera :: (a -> Bool) -> [a] -> Bool
provera p =and . map p

obrni :: Foldable t => t a -> [a]
obrni = foldl (flip (:)) []

delioci n = filter (f n) [1..n]
f n x = mod n x == 0

prost n = length (delioci n) > 2
cifre n = if n < 10 then [n] else (mod n 10) : cifre (div n 10)

broj [] = 0
broj (x:xs) = broj xs * 10 + x