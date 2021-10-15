import Data.List 

spoji l1 l2 = zip l1 l2

uvecaj n lista = map (+n) lista

pozitivni lista = filter (>0) lista

pozitivniDeo lista = takeWhile (>0) lista

saberi lista = foldr (+) 0 lista

absSuma listaListi = map (abs.sum) listaListi

sledbenici l = map (+1) (filter (>0) l)

element _ [] = False
element k (x:xs) = x == k || element k xs

element1 k lista = or (map (==k) lista)

sortiraj :: Ord a => [(a,a)] -> [(a,a)]
sortiraj = sortBy (flip poredi)

poredi :: Ord a => (a,a) -> (a,a) -> Ordering
poredi a b = compare (snd a) (snd b)
{-
sortiraj3 :: Ord a => [(a,a)] -> [(a,a)]
sortiraj3 = sortBy poredi 

poredi :: Ord a => (a,a) -> (a,a) -> Ordering 
poredi a b = compare (snd a) (snd b)
-}

qsort [] = []
qsort (x:xs) = qsort manji ++ [x] ++ qsort veci
    where
        manji = filter (<x) xs
        veci = filter (>=x) xs