package main

//generics allows a func to take in different params rather than explicitly defining diff funcs for same purpose

func calcParam[T int | float64](slice []T) T {
    var sum T

    for _, v := range slice {
        sum += v
    }
    return sum
}

