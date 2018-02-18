// Требуется максимизировать вероятность падения. Тогда вероятность успешного 
// прохождения всех тестов наименьшая. А это произведения (1 - p_i). Возьмем 
// логарифм. Теперь вместо произведения сумма каких-то отрицательных чисел.
// Задача свелась к задаче о рюкзаке, которую решим динамикой за O(NW), где
// N - число тестов, W - допустимое время работы. Надо быть аккуратным со
// взятием логарифма от 0. К счастью, в этом случае результат -inf, что 
// позволяет корректно найти ответ. В данном случаи им может быть любой тест,
// время работы которого не превышает ограничение. Если бы захотелось выбрать 
// как можно больше тестов с вероятностью падения 1, то это можно сделать жадно.

// Входные данные содержатся в файле "input.txt". Это строки по 2 разделенных
// пробелом числа -- время работы и вероятность падения. На последней строке 
// содержистя допустимое время работы. 
// Вывод производится в командную строку, это номера тестов, входящих в оптимальный 
// набор. В приведенном примере выгодно взять все тесты, кроме последнего, то есть 
// с номерами с первого по четвертый. 

// Sample input:                      Sample output:
// 3 0.5                              0.9375
// 2 0.5                              4
// 2 0.5                              3
// 3 0.5                              2
// 9 0.9                              1
// 10

package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
    "math"
)

func readFile(fname string) (W int, weight []int, prob []float64, err error) {
    b, err := ioutil.ReadFile(fname)
    if err != nil { return 0, nil, nil, err }
    
    lines := strings.Split(string(b), "\n")
    for _, l := range lines {
        if len(l) == 0 { continue }
        data := strings.Split(l, " ")
        // Если строка последняя, то чисел не два, а одно.
        if len(data) == 1 {
            W, err = strconv.Atoi(data[0])
            break
        }
        n, err := strconv.Atoi(data[0])
        if err != nil { return 0, nil, nil, err }
        weight = append(weight, n)
        
        f, err := strconv.ParseFloat(data[1], 64)
        if err != nil { return 0, nil, nil, err }
        prob = append(prob, f)
    }

    return W, weight, prob,  nil
}

func main() {
    W, weight, prob, err := readFile("input.txt")
    if err != nil { panic(err) }
    
    n := len(weight)
    for i := 0; i < n; i++ {
        prob[i] = math.Log2(1 - prob[i])
    }
    
    var res [][]float64
    var last [][]int
    res = make([][]float64, W + 1)
    last = make([][]int, W + 1)
    
    for i := 0; i <= W; i++ {
        res[i] = make([]float64, n + 1)
        last[i] = make([]int, n + 1)
    }
     
    
    for i := 0; i < n; i++ {
        for j := 0; j <= W; j++ {
            if (j + weight[i] <= W && res[j + weight[i]][i + 1] > res[j][i] + prob[i]) {
                last[j + weight[i]][i + 1] = i;
                res[j + weight[i]][i + 1] = res[j][i] + prob[i]
            }    
            if (res[j][i + 1] >= res[j][i]) {
                res[j][i + 1] = res[j][i]
                last[j][i + 1] = -1
            }
        }
    }
    
    min := 0.0
    var index int
    for i := 0; i <= W; i++ {
        if (res[i][n] < min) {
            min = res[i][n]
            index = i
        }
    }
    
    fmt.Println(1 - math.Pow(2, res[index][n]))
    for i := n; i > 0; i-- {
        if last[index][i] != -1 {
            fmt.Println(last[index][i] + 1)
            index -= weight[last[index][i]]   
        }
    }
}
