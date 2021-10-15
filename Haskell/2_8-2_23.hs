prost n = prostPom n 2

prostPom n k
    | n == k = True
    | mod n k == 0 = False
    | otherwise = prostPom n (k+1)

tipJne a b c 
    | a == 0 = print "degenerisana"
    | b * b - 4 * a * c == 0 = print "jedno resenje"
    | b * b - 4 * a * c > 0 = print "jedno resenje"
    | otherwise = print "kompleksna"

izDekadne a b =
    if a < b then print a
    else do 
        izDekadne (div a b) b
        print (mod a b)

delioci broj =
    nadjiDelioce broj 2

nadjiDelioce n k
    | k == n = []
    | mod n k == 0 = k : nadjiDelioce n (k + 1)
    | otherwise = nadjiDelioce n ( k + 1)

nadovezi l1 l2 n = 
    if n == 0 then l1
    else nadovezi (l1 ++ l2) l2 (n - 1)

brojDelilaca n = 
    prebrojDelioce n 2

prebrojDelioce n k
    | k == n = 0
    | (mod n k) == 0 = 1 + prebrojDelioce n (k + 1)
    | otherwise = prebrojDelioce n (k + 1)

fib n
    | n == 0 = 0
    | n == 1 = 1
    | otherwise = fib (n - 1) + fib (n - 2)

fibLista n =
    if n == 0 then []
    else do
        fibLista (n - 1) ++ [fib n]