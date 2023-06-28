/*
worker pool 指的是有許多的 goroutines 同步的進行一個工作。要建立 worker pool，
會先建立許多的 worker goroutine，這些 goroutine 中會：

 - 進行相同的 job
 - 有兩個 channel，一個用來接受任務（task channel），一個用來回傳結果（result channel）
 - 都等待 task channel 傳來要進行的 tasks
 - 一但收到 tasks 就可以做事並透過 result channel 回傳結果
*/