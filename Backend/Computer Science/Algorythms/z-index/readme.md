**Алгоритм Z-функции** (или алгоритм Z-индекса) — это эффективный способ обработки строк, используемый для решения задач поиска подстрок, сопоставления шаблонов, анализа повторов и других строковых операций. Он вычисляет Z-функцию строки, которая помогает быстро находить # Алгоритм Z-функции

Алгоритм **Z-функции** (или Z-индекса) — это эффективный способ обработки строк, используемый для поиска подстрок, сопоставления шаблонов, анализа повторов и других операций. Он вычисляет массив Z, который помогает быстро находить совпадения.

---


## Что такое Z-функция?

Для строки S длиной 𝑛, 𝑍[𝑖] — это длина наибольшего префикса строки 𝑆, начинающегося с позиции 𝑖, который одновременно является префиксом всей строки. Формально:

Z[i]=max{k∣S[0…k−1]=S[i…i+k−1]}

Пример: Для строки 𝑆 = "abcababc":

𝑍[0]= 𝑛 (длина всей строки, по соглашению).

𝑍[1]= 0 (с подстрокой 𝑆[1:]="𝑏𝑐𝑎𝑏𝑎𝑏𝑐" нет совпадения с началом 𝑆).

𝑍[3]= 2 (подстрока "𝑎𝑏" совпадает с началом 𝑆).

### Пример:

Для строки \( S = \text{"abcababc"} \):

| Индекс \( i \) | Подстрока \( S[i:] \) | \( Z[i] \) |
|----------------|------------------------|------------|
| 0              | abcababc              | 8          |
| 1              | bcababc               | 0          |
| 2              | cababc                | 0          |
| 3              | ababc                 | 2          |
| 4              | babc                  | 0          |
| 5              | abc                   | 3          |
| 6              | bc                    | 0          |
| 7              | c                     | 0          |

---

## Основная идея алгоритма

1. Поддерживается "окно" с границами \( [L, R] \), где \( S[L:R] \) совпадает с префиксом строки \( S \).
2. Если текущий индекс \( i \) попадает в окно (\( i \leq R \)), значение \( Z[i] \) может быть вычислено с использованием уже известных данных.
3. Если \( i > R \), \( Z[i] \) вычисляется с нуля, сопоставляя символы строки.

---

## Пошаговый процесс

1. Инициализация массива \( Z \):
   - \( Z[0] = n \) (длина строки по соглашению).
   - Все остальные элементы инициализируются нулями.

2. Для каждого индекса \( i \) от 1 до \( n-1 \):
   - Если \( i \leq R \), устанавливаем \( Z[i] = \min(R - i + 1, Z[i-L]) \).
   - Расширяем \( Z[i] \), сравнивая символы \( S[Z[i]] \) и \( S[i + Z[i]] \).
   - Обновляем границы \( [L, R] \), если \( i + Z[i] - 1 > R \).

---

## Псевдокод алгоритма

```go
package main

import "fmt"

// ComputeZ вычисляет Z-функцию для строки s
func ComputeZ(s string) []int {
    n := len(s)
    z := make([]int, n)
    z[0] = n // z[0] равно длине строки по соглашению

    l, r := 0, 0 // l и r определяют правую границу текущего окна

    for i := 1; i < n; i++ {
        if i <= r {
            z[i] = min(r-i+1, z[i-l])
        }
        for i+z[i] < n && s[z[i]] == s[i+z[i]] {
            z[i]++
        }
        if i+z[i]-1 > r {
            l, r = i, i+z[i]-1
        }
    }

    return z
}

// min возвращает минимальное из двух целых чисел
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    s := "aabcaabxaaaz"
    z := ComputeZ(s)
    fmt.Println("Z-функция:", z)
}
```
### Пример работы алгоритма
Для строки 𝑆 = "aabcaabxaaaz"
* Вход: 𝑆 = "aabcaabxaaaz"
* Выход: [12,1,0,0,3,1,0,0,2,1,0,0]

| **Индекс \( i \)** | **Подстрока \( S[i:] \)** | **\( Z[i] \)** |
|---------------------|---------------------------|----------------|
| 0                   | aabcaabxaaaz              | 12             |
| 1                   | abcaabxaaaz               | 1              |
| 2                   | bcaabxaaaz                | 0              |
| 3                   | caabxaaaz                 | 0              |
| 4                   | aabxaaaz                  | 3              |
| 5                   | abxaaaz                   | 1              |
| 6                   | bxaaaz                    | 0              |
| 7                   | xaaaz                     | 0              |
| 8                   | aaaz                      | 2              |
| 9                   | aaz                       | 1              |
| 10                  | az                        | 0              |
| 11                  | z                         | 0              |


### Применение Z-функции
1. **Поиск подстроки (Pattern Matching)** 
Конкатенируем шаблон 𝑃 и строку 𝑇 через специальный разделитель:
``` 
𝑆 = 𝑃 + # + 𝑇 
```
Вычисляем 𝑍[𝑖]. Все индексы 𝑖, где 𝑍[𝑖] = len(𝑃), указывают на вхождение 𝑃 в 𝑇.

2. **Нахождение повторов** 
Z-функция помогает найти повторяющиеся префиксы в строке.

3. **Определение периодичности строки**
Используя Z-функцию, можно быстро определить минимальный период строки.

### Сложность алгоритма
 * **Временная сложность**: 𝑂(𝑛), так как каждый символ строки обрабатывается не более двух раз.
 * **Пространственная сложность**: 𝑂(𝑛), для хранения массива Z.