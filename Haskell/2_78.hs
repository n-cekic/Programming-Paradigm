data Map k v = Null | Node k v (Map k v) (Map k v) deriving Show

null Null = True
null _ = False

f k v1 v2 = k + v1 - v2

insertWithKey f k v Null = Node k v Null Null
insertWithKey f k v (Node k2 v2 l r)
    | k < k2 = Node k2 v2 (insertWithKey f k v l) r
    | k > k2 = Node k2 v2 l (insertWithKey f k v r)
    | otherwise = Node k2 (f k2 v2 v) l r

insertWith :: Ord k => (v -> v -> v) -> k -> v -> Map k v -> Map k v
insertWith f = insertWithKey (\ k v v2 -> f v v2)

insert :: Ord k => k -> v -> Map k v -> Map k v
insert = insertWith (\ v v2 -> v2)

fromList :: Ord k => [(k, v)] -> Map k v
fromList = foldl (\ m (k, v) -> insert k v m) Null

search :: Ord k => Map k v -> k -> Maybe v
search Null _ = Nothing 
search (Node k v l r) k2
    | k < k2 = search r k2
    | k > k2 = search l k2
    | otherwise = Just v

findWithDefault :: Ord k => v -> k -> Map k v -> v
findWithDefault d k m = case search m k of 
    Nothing -> d
    Just v -> v

member :: Ord k => Map k v -> (k -> Bool)
member m k = case search m k of
    Nothing -> False 
    Just _ -> True

adjustWithKey :: Ord k => (k -> v -> v) -> k -> Map k v -> Map k v
adjustWithKey f k Null = Null
adjustWithKey f k (Node k2 v2 l r)
    |k < k2 = Node k2 v2 (adjustWithKey f k l) r
    |k > k2 = Node k2 v2 l (adjustWithKey f k r)
    | otherwise  = Node k2 (f k v2) l r

adjustAll f Null = Null
adjustAll f (Node k v l r) = Node k (f k v) (adjustAll f l) (adjustAll f r)

deleteMin (Node k v Null r) = ((k, v), r)
deleteMin (Node k v l r) = ((k2, v2), Node k v o r)
    where ((k2, v2), o) = deleteMin l

delete k Null = Null
delete k (Node k2 v2 l r)
    | k < k2 = Node k2 v2 (delete k l) r
    | k > k2 = Node k2 v2 l (delete k r)
    | otherwise = case r of
        Null -> l
        _ -> Node k3 v3 l d
            where ((k3, v3), d) = deleteMin r

foldMapa :: (a -> k -> v -> a) -> a -> Map k v -> a
foldMapa _ i Null = i
foldMapa f i (Node k v l r) = foldMapa f (f (foldMapa f i l) k v) r

siye :: Map k v -> Int 
siye = foldMapa (\ a k v -> a + 1) 0

toList :: Ord k => Map k v -> [(k, v)]
toList = reverse . foldMapa (\a k v -> (k, v) : a) []

union :: Ord k => Map k v -> Map k v -> Map k v
union m m2 = foldMapa (\ a k v -> insert k v a) m m2