## Заметки из /src/runtime/map.go

Мапа - это хеш таблица. Данные в ней располагаются в массиве бакетов (вёдер).
Каждый бакет содержит 8 пар ключ/значение. Младшие биты хеша нужны, чтобы определить бакет. 
Каждый бакет содержит несколько старших битов хеша, чтобы отличать записи друг от друга. 

Если происходит переполнение бакета и в нем необходимо хранить больше 8 значений происходит связывание бакетов.
Когда хеш таблица растет, происходит аллокация нового массива бакетов, с удвоенным размером.
Бакеты постепенно копируются из старого массива бакетов в новый массив бакетов. 

Итерация по мапе возвращает ключи в порядке обхода, ключи никогда не перемещаются внутри бакета.
При увеличении таблицы итерация продолжается по старой таблице, а также происходит проверка, не был ли бакет, по которому выполняется итерация, перемещен (эвакуирован) в новую таблицу.

При компиляции операции с мапами перезаписываются методами из рантайма.
```go
  v := m["key"]       → runtime.mapaccess1(m, "key", &v)
  v, ok := m["key"]   → runtime.mapaccess2(m, "key", &v, &ok)
  for k, v := range m → runtime.mapaccessK(m, "key", &k, &v)
  m["key"] = 1        → runtime.mapassign(m, "key", 1)
  delete(m, "key")    → runtime.mapdelete(m, "key")
  clear(m)            → runtime.mapclear(m)
```

```go

// упрощенная реализация поиска по ключу
// 
func mapaccess(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
	if h == nil || h.count == 0 {
		return zeroPtr
	}
	
	hash := t.hasher(key, h.hash0) 
	bucketIdx := hash & (2^h.B-1) // hash % number of buckets
    bucket := add(h.buckets, bucketIdx*t.bucketsize) // jump to address
	
	top := hash >> (goarch.PtrSize*8 - 8) // верхний байт хеша
	for i:=0; i < 8; i++ {
	    if bucket.tophash[i] != top {
		    continue
		}

        k := add(unsafe.Pointer(b), i*uintptr(t.keysize))

        if t.key.equal(key, k) {
            return add(unsafe.Pointer(b), bucketCnt*uintptr(t.keysize)+i*uintptr(t.elemsize))
        }
	    
	}

    return unsafe.Pointer(&zeroVal[0])

```


```go
// Мап хедер
type hmap struct {
  count     int // Размер 
  flags     uint8 // битовые флаги: iterator итерация по бакетам, oldIterator итерация по старым бакетам, hashWriting запись, sameSizeGrow рост в тот же самый рамер 
  B         uint8  // log_2 бакетов (может содержать до 6.5 * 2^B записей)
  noverflow uint16 // примерное значение переполненных бакетов
  hash0     uint32 // сид хеша

  buckets    unsafe.Pointer // массив из 2^B бакетов. Может быть nil, если count==0.
  oldbuckets unsafe.Pointer // предыдущий массив бакетов (в два раза меньший), не nil только в процессе роста
  nevacuate  uintptr        // счетчик прогресса эвакуации (бакеты меньше этого уже были эвакуированы)

  extra *mapextra // опциональные поля, описаны ниже
}

// Содержит поля, которые есть не во всех мапах
type mapextra struct {
  overflow    *[]*bmap // бакеты переполнения
  oldoverflow *[]*bmap // бакеты переполнения для предыдущего массива бакетов
  nextOverflow *bmap // следующий свободный бакет для новых переполнений
}

// Вместо полной реализации мапы для каждого уникального объявления 
// компилятор создаст maptype 
type maptype struct {
  typ    _type
  key    *_type
  elem   *_type
  bucket *_type 
  hasher     func(unsafe.Pointer, uintptr) uintptr // функция для хеширования ключа
  keysize    uint8 // размер ключа для ариф. операций в бакете  
  elemsize   uint8  
  bucketsize uint16
  flags      uint32
}

type _type struct {
  size       uintptr
  equal func(unsafe.Pointer, unsafe.Pointer) bool
  // ...
}

```

### Как устроена функция хеширования ключа

В спецификации языка это не описано, значит, что принцип может легко изменится в будущем.
Сам алгоритм хеширования различается в зависимости от типа ключа и платформы.
Компилятор задаёт `maptype` и определяет в нем `hasher` для каждого типа `key`. 
Для интерфейсов в качестве ключей он работает медленнее, потому что приходится вычислять функцию в рантайме. 

```go
// src/runtime/alg.go
func typehash(t *_type, p unsafe.Pointer, h uintptr) uintptr

```

Хеш для структуры это хеш от всех её полей.  
Для архитектур с поддержкой AES-NI используется https://ru.wikipedia.org/wiki/Расширение_системы_команд_AES  
Для float хеш от NaN это рандом.  


### Стандартная библиотека

Go 1.21 принес нам пакет `maps` в котором собраны generic методы для работы с мапами

```go
func Equal[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool
func EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool
func Clone[M ~map[K]V, K comparable, V any](m M) M
func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2)
func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool)
```
