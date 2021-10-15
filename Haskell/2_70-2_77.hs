type Par a = (a,a)

pomnoi :: Par Int -> Int 
pomnoi (a,b) = a * b

data Bulovski = Tacno | Netacno

data Trougao a b c = Jednakostranicni Float 
    | Jednakokraki Float Float 
    | Raznostrani Float Float Float
    deriving (Show, Eq, Ord)

obim :: Trougao Float Float Float -> Float
obim (Jednakostranicni a) = 3*a
obim (Jednakokraki a b) = 2*a + b
obim (Raznostrani a b c) = a + b + c

listaTrougao :: [Float] -> [Trougao Float Float Float]
listaTrougao lista = map Jednakostranicni lista

data Zivotinja = Pas | Macka | Papagaj deriving Show

data Ljubimac = MojLjubimac {ime :: String, zivotinja :: Zivotinja, godine :: Int} deriving Show
stariji lj1 lj2 = godine lj1 > godine lj2

data Pravougaonik = MojPravougaonik {a :: Int, b :: Int}

instance Show Pravougaonik where
    show (MojPravougaonik a b) = "(" ++ show a ++ ", " ++ show b ++ ")"

instance Eq  Pravougaonik where
    (==) (MojPravougaonik a1 b1) (MojPravougaonik a2 b2) = a1 == a2 && b1 == b2 || a1 == b2 && a2 == b1

data Lista a = Null | Konstanta a (Lista a)

duzina :: Lista a -> Int
duzina Null = 0
duzina (Konstanta x xs) = 1 + duzina xs

glava :: [a] -> Maybe a
glava [] = Nothing 
glava (x:_) = Just x

rep :: [a] -> Maybe [a]
rep [] = Nothing
rep (_:xs) = Just xs

type Pos = (Int, Int)

type Pic = Pos -> Bool

empty :: Pic 
empty _ = False 

full :: Pic
full _ = True 

only :: Pos -> Pic
only p x = p == x

inverse :: Pic -> Pic
inverse f p = not (f p)

union :: Pic -> Pic -> Pic
union f1 f2 p = f1 p || f2 p

intersect :: Pic -> Pic -> Pic
intersect f1 f2 p = f1 p && f2 p

render :: Pic -> [[Bool]]
render p = [[p (x, y) | y <- [0..4]] | x <- [0..4]]

extract :: [[Bool]] -> Pic
extract l (x,y) = l !! x !! y