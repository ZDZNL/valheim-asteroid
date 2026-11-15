package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.4" )

func BXaafJ9kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RlG7pMf7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9hVHx3yVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUHNaVxJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5uMq5OwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WPYpITRxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KXXnuYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func krD6S28JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhud47AQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCFOB7LYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yA5uLLxiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lEVX3HpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDPbiLP4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0la03L9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ogyxREnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5mvJcSxnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3uh1s4F9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ykrpsv8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aTII5dP0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WXoUkvirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXQ3hmYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UppX5WNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iwoVma2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tdt8f72IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VyVad2R9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UcnoynnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2eO3njFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijSoHqj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tC7lkrbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func doV2IQvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0ua9CRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pO9679nSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6EzqD1NPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mEkw4vGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GGwvbwCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SRlGI1R2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BQWi5u4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nyXRpU0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKrqFwJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xYBCZQECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H9w3cSxvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o243eurgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pcqoYF05Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WEINXN7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iwYCRdWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ep5nL3j8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stac4axSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ga8Uv8FIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rm9wkxviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIUoYko9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z5ofs2mQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhuWEMr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZmDFYV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Id3lh74XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ww6tw6m3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 73bVCG1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjGF8y7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNPxjE7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKNNofnbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXK5iDUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PrcqOUmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4sbpQEyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQcBoTlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6hrRNTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rc1dJmkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WBoCA0W1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UV42qYDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OmOsr1fuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qEqzHWbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCzhkzW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41lX6etDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8HXaCpKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nLHRZEB0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFJnIqLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjcVr4SLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bM4ddUzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zXx4cQqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XSPLk5b5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5GIZOwEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XKVFrnI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mpGur77eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func loCB7LQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RpQyFzjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EBnXqIhsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OehF90bCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Naq177LRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lQgpux77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VlcAeYKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzY12Y6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lcEFXvOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PGRCPcNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jrk9GARCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NuL2DkjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U9vkiegQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIeAaIlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vwqGV9DlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLgFd7e4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func McwVClkCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78VfXLv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X1ujZTElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cKGL9ot7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fKlLGah0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nDM4L0qzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyCWffTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CW5GdRJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ikXp5PINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INxzD8gyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GERiIGBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbmzLS47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QzL9pCDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xPVAKGsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQ1q7hAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSxlhMmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKiFKWR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U99Z5Nt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upjSjWpNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IbeACgNnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkGnW605Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOfogZcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FbI3vHOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EMq93e6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nwZgbtLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJXTmOhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fh7VX1GkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zPLBBQS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6r9IzMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xDnw33tiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5R9qvZEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKfDoDwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ctdlpZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0CTofaMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHhmRbL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzIDao5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BF7ZDcohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PGuGv1OIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9mn0rH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9pBROqSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owCGlzTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GvapFOwjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33jLc8fjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWnzZTM0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HAbtXWujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9A1aqffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P6SUHCpqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khFTvQzPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GPrFWwwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 09D8wYNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Thh8A4Z7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQnUHZUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2XwSUf0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wq1hr7u1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAlB0onwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmnbIjONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbplOKI1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1HsjNWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IeBosDIIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IcMkFjODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZvGXSN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0CmlA6HHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uibY84SDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NNsxzj3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x8OjIUZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q0hLd4vaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ILU5yHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCU8O8j7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cq1VwVlSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NVWDZG9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6uEXDn2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNCgtie8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSlqh8DRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BPV90AQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FubcFbYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxxE4zruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q6RVBnasWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J93D9w1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LszJn5dsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pDY2wTsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hey0iODAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dWTzltOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sregD0GZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8XFnCrVGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfwdM1bcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ca2d9ZUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SELEUYXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UOySnYjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XvCepiU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcPwNKEoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pdkNYoqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QqlSr9BwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41btV1s9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RAH72EF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPSHbewOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 899TPXbnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQTecEhEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZRsp3zuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XWVZegamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3HehykcBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4kqS3KF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2JWSxhYJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qNkwUvk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VNkcUCgLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lmkVCHXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gfzv2nUxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zyh6sKQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ML6Z5K8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ow6EnKjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RyfK4LknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boapCqCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IYb9LQtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pYjCPFXwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nuYMr8ELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YrlENQ0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HfYK3FkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 18vj9IZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HCLjmElOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EB5QIuo8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGcISC01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xKbsjTRsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vItXxgvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGw7Oh5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MWJL8HiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDjlMNrJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func weIuJ8llWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FsJYPuHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FRjjgn75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCSr8qBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhpyZShOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mwj2utUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jeMfHzivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t00HZQxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7nEHKL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8skmMS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qjh2bmkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oDq3vhuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBcsTvegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzUxECuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gjcWPjcWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LyeMkK0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dEQSVhkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9VAcIDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13GT86PhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RXu0TgN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZO8orM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dshu7pqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5ZkaTblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNAmEwIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BElRBUeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzkgzSAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NkdHSVY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOZ3Jro7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUF156bWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jexk8OLFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func irzcdeONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dKdIdk8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func COUoGeBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsBVPeWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C9o1glwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7hHLssyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVhFB8a2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBn0QtlCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3yoVPZW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W6JNHK13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dc4YkZdSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8HtPof6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wtNSBxaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iG2IgdKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4cDs5QBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M3s2IgUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q13suDrZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nqdVMM6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T6pGJZlzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjuokTfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hKq0z8N1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1MhEI8sSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VVXjXBApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6vo4r7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XM5WZcH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aoHoxVqAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbY7uCXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pcuHdSEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func brUE6aQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWyOVOrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wlSGNtXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SR95VlZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xME7F5DtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vb2uKPH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I1OOU67mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qiCzccMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5EibwN9xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lvfw5xgFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kSALRLQLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eck1EobdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5tWdRRwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJt2cIODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XOLVZW8RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X2W0arp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PfNQtZpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eAHs4IPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rlc2lQ1cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e78Q1FFZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dpui7DgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZI5qWGDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a5RDEJEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PSNjxI3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QakCbtpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWGI2a63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGxQPirAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N7U4tPgtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lMLXd7O7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IuSS6PcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3MF4ko5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4f7eeLbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBVirtKJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpdehuAtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i0WNJh5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LVLZ9VHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g8D4IclpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHev5ZeAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H7tqqyBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uBmIyomTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUS4VUL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ty7MHJsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fcYowIVGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dvel5KLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cR1Ku8OBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8LN7i82GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2dfvxctTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFYnioX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KEKGnIXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yqzhx3SyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNPvE8YxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9FU7qLfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3RFMftaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Azays3x6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sMRyTDygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GW5YIitiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0auerGtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PLOfZRljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H721uVDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SH3Mooe2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9UGpilRHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NecXU9TkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6EMCdoMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dtD3tGwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTsdvL0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bkNqTZEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tZrmTOvFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IvBOyJGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CkvM1WG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 234ez6RZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8L08uPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uad1atJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CoN1waCeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPf6k388Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QqJYZdTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AK8UiD3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyMXYAlqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZTcQVN4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8UTRNnWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZoY4U8YUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZMEiB28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x2QcPHhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s2HY1k3rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fvJkDROOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BL2rZ3EnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kql3DJv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPrE2VlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AWjVMCgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PuphDs7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8f99eFkiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tI4Yl6ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qLxgzuNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EGJDw5hAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJBsjPPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXyeFmcXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KDe2yYxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PoVQWxACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dk9VU4L9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z4iKsJU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WRoIHv67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FmoNi5BiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJMutEFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L2LrFFTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b2wNiEKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UdACLdXUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izW6RAxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l8YboIcLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRVhmhduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HYZAZhxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T0yGLb4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HQhZKUgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1gr2Qyf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qk6bLijOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gJkugEBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h9dcJpZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVroH9y3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KBvGsv9fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YoK9tf7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F61iR84PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1y5Hdh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYNvddB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aiNd5sUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MsU8D3XTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r99gZWL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KucOLggaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IWkOoWu7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HzROAyemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VOTOTiMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3GSAgFuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJnvJqv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zPDIaz1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76vvNZVsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V4IXX1oRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zhxewq6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 06R0aflCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MCU4xS0zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NkcuAwmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zRgoqRH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func simbVgUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2kaX7u4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WG5S95rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9v08cKH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cM1gc3BhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZvITZa9VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1kg2XWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78RGmr3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y58DaFkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76AtaT96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVdMCFnMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZGj5SZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSWdwaN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nMvrJLSDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6JE5wybKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LlDwwLARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S5fpKxcWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TDqTx4PEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wtZrHwnfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E8VsefsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmw5jwaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYvRncofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func micBqu93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LyZDb3lTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKz8OiRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NAG9Ewg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGgJ62jOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpRTQ089Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z58ET66VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OZrBS0h6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QpDFKVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZnpIkMJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8K5SaW5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bg4j1siWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTtFWrrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36PTZ7NxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y1pB7QIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lxUfiuhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PbUyXvMmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eg3a7NaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CelzSpgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jHNerXFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eEYfNKRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpbqL53WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 85v42HfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U1g1iULEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func faXH9mFRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HhF6IRceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZXuf8kMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLAge0dbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ww34Ta8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZUjg9nsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qn1dgmqWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3VI0enGRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rO6cLyuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hMtyIWp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6kRogOlqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LnesKK5cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZyobqV6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WyYBtistWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Chc9XjKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8p3UbrNdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IFU5tMHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LuQo9l7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bHnScHlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gfSz0PRsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RrSF1LQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1mBKCnhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h5SL4fBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v1OWxPUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YIY42ymlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uf9eFohCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7nxueOOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V2z1g3MqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q7sgof4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func spd6FKsBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FYpC7EEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3F7aAdv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hrlLRSwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RpdVFAYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U1rzHEchWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O2sJ3rtxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kP4ai8qMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3dZEErSjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5bTKcBGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AdiF43QLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqcYpGHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ewsj39Z5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nW2gGdmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cME1fxFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vHjFx1FcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ncyVAtvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 96VHyEMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func npkh8YxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TvYEgHIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mMyHC1ejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s55Y09NdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P89cmxThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2G8xgTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JoErupc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func smxQbl56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1J1QgGAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CTvJpWjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQCuStDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IxsXjEtZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O9FqPLZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D27s89ZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Us0rpZWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hlfHldAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ofW7trwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k7wThJuuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hzYEUYrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nLqN6DeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Y79RoT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func atjPnb7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8BbW8KjIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQXt5Xb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VU4jgsgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9h235ygaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X4z0ezSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XILLPqF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LitD3qcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cECULHcmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4S3s8w8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Qs6cclTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tWOQ3ypYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9DZoofxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3k38UDyhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9BH4INiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6T5Nk0eSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n41VMELdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJcT1L2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRUZAPkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dz0VjI24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7weC5dKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ms5ruuvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gs4khzdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lgKCjvOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4srQLUu5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dLGJOoYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mm3PDocaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6VOhFBmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ov1yrjRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXCraBuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z8bNFgN2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0BSXOcMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v1YRCOzWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgfLqkAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vq6kEV7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CgZdUb1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCXuFpacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xAH89sAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWqIJxqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ahLJtpxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2qWEidKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lATWni03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PstAX4PiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuVoVR6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0FR2QlpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5VSBibkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pvC7W9r2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jOQz5PPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dLs0tMViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLbk9sS0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bPa7IhUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJF03wULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TAEHHZ2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func siWRT6j0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AgCux9azWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ecU2sxF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RVD8L0nEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XyfoQCihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DVJRy5faWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EnhVG48ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XOV33vtXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wr0I7DReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ly9rZsclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1z2hAc0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzt9OdSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3M00wdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hm2ctfeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bVaMBu1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Wc5mkxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FBGZagWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9fX0PWEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z826qDrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4fPD76HoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLKC8jH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPLDi3BqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43lkcYi0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yAyhYjnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0T2edz0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VtLuo0KfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQuos3g7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8z4gWfPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k4IVA5gHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0yFR6GYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JbyVxgNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVHrJIKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rUrYL9IrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6icvwNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbkFSUMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VkxOxW2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OTfVC51iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pv2abz1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7dfhvNbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOF4rNJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUAc5gknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tX4fkngmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ifoCpHvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TX2rOircWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFeLcKgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B5MVah4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZi5cY7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ft71PKXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PQ5JYKJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MX9K5XapWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t9WRoI7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEiwsai5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6VuajDXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uKnwKVycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3WKOWS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8UBh6k9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nmSeAUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R75llyPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzMrKUM9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g6rTVsFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lATKla68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AjCQ7dlxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7E0gOQDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lcferZ99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VLMhM3AfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Vrp4hYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJXMfWXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qja8vZRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQMv953iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IuGJmP8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k7VZqif6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjPVGrN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lA3TXX8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jheD3zAxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPm4jIG3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hn7CE6c4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWpQmnZIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4CGO9hKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CspjvVJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VfLGCVr5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0t0vqP62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qErcQY5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 40kdXvMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ynFlZCcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GcqVCy3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wz9JwwEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JN1mHGkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJ0oWXwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RqiiuMaxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xs2SdrpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DD6NPNETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qAcuxxrMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wwXE83JqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJX0cbKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JYUtAoBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tsdNmkyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DI9HMauTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJG52F0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2V0AANIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wgUcHS9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mihi2GTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ia6GyUZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5qu1mVkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsxYzIQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d91fcrg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DY4InMfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7dAyMt49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FydLixyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GS9M7rFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fqjZnHpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jK9vX0ZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5qX2mUT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6TjEU2RKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8GZM7xMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jGUH9N6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y5AUIKJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RWiIEDTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9poUrrh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mxaZDldVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cmNJOLY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f6wo3RM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lkwbWBVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjlYzkghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9i0OfSE7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func deJF40jOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QRwYY3Y0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QstplOnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JpUY0c4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBZ6fr2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zieiAebdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3I5CDSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzS09f0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oYyZWSbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func snRMhFCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALiAXA8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AL1LQAZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4I85WA95Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ciRma2IGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSS6LlISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GexpzRPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wu1V0RBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNIPOlsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uWDSpxmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upVZYPsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7JZecX9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MMJgmiteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7F6umzqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nkE4hq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpvqLjkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3w3hno4SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 18gsURAGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7qlxWOqlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6VVgt3t4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1Yy8mawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SqzIOhWDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wEe6V6iVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGDqsKy0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6jTL8nOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tgno7NwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P9W86qXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8pR1ssreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0jS9y83kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJFGztOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tzR3ShPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CMIfMdR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwWDvQCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func euHioqkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WDDvXILKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WzOm3BfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cjre1hqYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MidS8RauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wACDfOtmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

